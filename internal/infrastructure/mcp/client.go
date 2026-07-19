package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
	"time"
)

type TransportType string

const (
	TransportStdio TransportType = "stdio"
	TransportSSE   TransportType = "sse"
	TransportWS    TransportType = "websocket"
)

type Transport interface {
	Connect(ctx context.Context) error
	Send(data []byte) error
	Receive() ([]byte, error)
	Close() error
}

type StdioTransport struct {
	cmd     *exec.Cmd
	stdin   io.WriteCloser
	stdout  io.ReadCloser
	stderr  io.ReadCloser
	mu      sync.Mutex
	scanner *bufio.Scanner
	cmdStr  string
	args    []string
	env     []string
}

func NewStdioTransport(cmdStr string, args, env []string) *StdioTransport {
	return &StdioTransport{
		cmdStr: cmdStr,
		args:   args,
		env:    env,
	}
}

func (t *StdioTransport) Connect(ctx context.Context) error {
	t.cmd = exec.CommandContext(ctx, t.cmdStr, t.args...)
	if len(t.env) > 0 {
		t.cmd.Env = append(t.cmd.Env, t.env...)
	}

	stdin, err := t.cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("stdin pipe: %w", err)
	}
	t.stdin = stdin

	stdout, err := t.cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("stdout pipe: %w", err)
	}
	t.stdout = stdout
	t.scanner = bufio.NewScanner(t.stdout)
	t.scanner.Buffer(make([]byte, 0, 1024*1024), 1024*1024)

	stderr, err := t.cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("stderr pipe: %w", err)
	}
	t.stderr = stderr

	if err := t.cmd.Start(); err != nil {
		return fmt.Errorf("start process: %w", err)
	}

	go func() {
		errBuf := new(strings.Builder)
		io.Copy(errBuf, t.stderr)
	}()

	return nil
}

func (t *StdioTransport) Send(data []byte) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	msg := string(data) + "\n"
	_, err := fmt.Fprint(t.stdin, msg)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

func (t *StdioTransport) Receive() ([]byte, error) {
	for t.scanner.Scan() {
		line := strings.TrimSpace(t.scanner.Text())
		if line == "" {
			continue
		}
		return []byte(line), nil
	}
	if err := t.scanner.Err(); err != nil {
		return nil, fmt.Errorf("receive: %w", err)
	}
	return nil, io.EOF
}

func (t *StdioTransport) Close() error {
	if t.cmd != nil && t.cmd.Process != nil {
		t.cmd.Process.Kill()
		t.cmd.Wait()
	}
	return nil
}

type MCPClient struct {
	transport    Transport
	connected    bool
	mu           sync.RWMutex
	timeout      time.Duration
	serverInfo   map[string]any
	capabilities map[string]any
}

func NewMCPClient(transport Transport) *MCPClient {
	return &MCPClient{
		transport: transport,
		timeout:   30 * time.Second,
	}
}

func (c *MCPClient) Connect(ctx context.Context) error {
	if err := c.transport.Connect(ctx); err != nil {
		return fmt.Errorf("transport connect: %w", err)
	}
	c.mu.Lock()
	c.connected = true
	c.mu.Unlock()
	return nil
}

func (c *MCPClient) IsConnected() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.connected
}

func (c *MCPClient) Call(ctx context.Context, method string, params map[string]any) (map[string]any, error) {
	req := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  params,
	}

	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	if err := c.transport.Send(reqData); err != nil {
		c.mu.Lock()
		c.connected = false
		c.mu.Unlock()
		return nil, fmt.Errorf("send: %w", err)
	}

	respData, err := c.transport.Receive()
	if err != nil {
		c.mu.Lock()
		c.connected = false
		c.mu.Unlock()
		return nil, fmt.Errorf("receive: %w", err)
	}

	var resp struct {
		JSONRPC string         `json:"jsonrpc"`
		ID      int            `json:"id"`
		Result  map[string]any `json:"result,omitempty"`
		Error   *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error,omitempty"`
	}
	if err := json.Unmarshal(respData, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("MCP error (code %d): %s", resp.Error.Code, resp.Error.Message)
	}

	return resp.Result, nil
}

func (c *MCPClient) Initialize(ctx context.Context) error {
	result, err := c.Call(ctx, "initialize", map[string]any{
		"protocolVersion": "2024-11-05",
		"clientInfo": map[string]any{
			"name":    "termcode",
			"version": "0.1.0",
		},
	})
	if err != nil {
		return fmt.Errorf("initialize: %w", err)
	}

	c.mu.Lock()
	if sv, ok := result["serverInfo"].(map[string]any); ok {
		c.serverInfo = sv
	}
	if caps, ok := result["capabilities"].(map[string]any); ok {
		c.capabilities = caps
	}
	c.mu.Unlock()

	return nil
}

func (c *MCPClient) ListTools(ctx context.Context) ([]ToolInfo, error) {
	result, err := c.Call(ctx, "tools/list", nil)
	if err != nil {
		return nil, fmt.Errorf("list tools: %w", err)
	}

	toolsRaw, ok := result["tools"].([]any)
	if !ok {
		return nil, fmt.Errorf("unexpected tools/list response format")
	}

	tools := make([]ToolInfo, 0, len(toolsRaw))
	for _, tRaw := range toolsRaw {
		tMap, ok := tRaw.(map[string]any)
		if !ok {
			continue
		}
		tool := ToolInfo{
			Name:        toString(tMap["name"]),
			Description: toString(tMap["description"]),
		}
		if schema, ok := tMap["inputSchema"]; ok {
			tool.InputSchema = schema
		}
		tools = append(tools, tool)
	}
	return tools, nil
}

func (c *MCPClient) CallTool(ctx context.Context, name string, args map[string]any) (*ToolResult, error) {
	result, err := c.Call(ctx, "tools/call", map[string]any{
		"name":      name,
		"arguments": args,
	})
	if err != nil {
		return nil, fmt.Errorf("call tool: %w", err)
	}

	toolResult := &ToolResult{
		Success: true,
	}

	if content, ok := result["content"].([]any); ok {
		var parts []string
		for _, cRaw := range content {
			if cMap, ok := cRaw.(map[string]any); ok {
				if text, ok := cMap["text"].(string); ok {
					parts = append(parts, text)
				}
			}
		}
		toolResult.Output = strings.Join(parts, "\n")
	}

	if isError, ok := result["isError"].(bool); ok && isError {
		toolResult.Success = false
		toolResult.Error = toolResult.Output
		toolResult.Output = ""
	}

	return toolResult, nil
}

func (c *MCPClient) Close() error {
	c.mu.Lock()
	c.connected = false
	c.mu.Unlock()
	return c.transport.Close()
}

func (c *MCPClient) ServerInfo() map[string]any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.serverInfo
}

func (c *MCPClient) Capabilities() map[string]any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.capabilities
}

type ToolInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema any    `json:"inputSchema,omitempty"`
}

type ToolResult struct {
	Success bool   `json:"success"`
	Output  string `json:"output,omitempty"`
	Error   string `json:"error,omitempty"`
}

func toString(v any) string {
	if v == nil {
		return ""
	}
	s, ok := v.(string)
	if !ok {
		return fmt.Sprintf("%v", v)
	}
	return s
}
