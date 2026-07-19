package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type SSETransport struct {
	url       string
	client    *http.Client
	body      io.ReadCloser
	scanner   *bufio.Scanner
	mu        sync.Mutex
	connected bool
	events    chan SSEEvent
	done      chan struct{}
}

type SSEEvent struct {
	Event string
	Data  string
	ID    string
}

func NewSSETransport(url string) *SSETransport {
	return &SSETransport{
		url: url,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		events: make(chan SSEEvent, 64),
		done:   make(chan struct{}),
	}
}

func (t *SSETransport) Connect(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, t.url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "text/event-stream")

	resp, err := t.client.Do(req)
	if err != nil {
		return fmt.Errorf("http connect: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("SSE connect: HTTP %d", resp.StatusCode)
	}

	t.body = resp.Body
	t.scanner = bufio.NewScanner(resp.Body)
	t.connected = true

	go t.readLoop()

	return nil
}

func (t *SSETransport) readLoop() {
	defer close(t.events)
	defer close(t.done)

	var current SSEEvent
	for t.scanner.Scan() {
		line := t.scanner.Text()
		if line == "" {
			if current.Data != "" {
				t.events <- current
			}
			current = SSEEvent{}
			continue
		}

		if strings.HasPrefix(line, "event: ") {
			current.Event = strings.TrimPrefix(line, "event: ")
		} else if strings.HasPrefix(line, "data: ") {
			current.Data = strings.TrimPrefix(line, "data: ")
		} else if strings.HasPrefix(line, "id: ") {
			current.ID = strings.TrimPrefix(line, "id: ")
		}
	}
}

func (t *SSETransport) Send(data []byte) error {
	return fmt.Errorf("SSE transport does not support sending")
}

func (t *SSETransport) Receive() ([]byte, error) {
	event, ok := <-t.events
	if !ok {
		return nil, io.EOF
	}
	return []byte(event.Data), nil
}

func (t *SSETransport) Close() error {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.connected = false
	if t.body != nil {
		return t.body.Close()
	}
	return nil
}

type WebSocketTransport struct {
	url       string
	conn      io.ReadWriteCloser
	mu        sync.Mutex
	connected bool
	buf       *bufio.ReadWriter
}

func NewWebSocketTransport(url string) *WebSocketTransport {
	return &WebSocketTransport{
		url: url,
	}
}

func (t *WebSocketTransport) Connect(ctx context.Context) error {
	return fmt.Errorf("WebSocket transport not yet implemented")
}

func (t *WebSocketTransport) Send(data []byte) error {
	return fmt.Errorf("WebSocket transport not yet implemented")
}

func (t *WebSocketTransport) Receive() ([]byte, error) {
	return nil, fmt.Errorf("WebSocket transport not yet implemented")
}

func (t *WebSocketTransport) Close() error {
	return nil
}

type MCPManager struct {
	clients map[string]*MCPClient
	mu      sync.RWMutex
}

func NewMCPManager() *MCPManager {
	return &MCPManager{
		clients: make(map[string]*MCPClient),
	}
}

func (m *MCPManager) ConnectStdio(ctx context.Context, id, cmd string, args, env []string) (*MCPClient, error) {
	transport := NewStdioTransport(cmd, args, env)
	client := NewMCPClient(transport)

	if err := client.Connect(ctx); err != nil {
		return nil, fmt.Errorf("connect stdio: %w", err)
	}
	if err := client.Initialize(ctx); err != nil {
		client.Close()
		return nil, fmt.Errorf("initialize: %w", err)
	}

	m.mu.Lock()
	m.clients[id] = client
	m.mu.Unlock()

	return client, nil
}

func (m *MCPManager) ConnectSSE(ctx context.Context, id, url string) (*MCPClient, error) {
	transport := NewSSETransport(url)
	client := NewMCPClient(transport)

	if err := client.Connect(ctx); err != nil {
		return nil, fmt.Errorf("connect SSE: %w", err)
	}
	if err := client.Initialize(ctx); err != nil {
		client.Close()
		return nil, fmt.Errorf("initialize: %w", err)
	}

	m.mu.Lock()
	m.clients[id] = client
	m.mu.Unlock()

	return client, nil
}

func (m *MCPManager) Get(id string) (*MCPClient, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.clients[id]
	return c, ok
}

func (m *MCPManager) Disconnect(id string) error {
	m.mu.Lock()
	client, ok := m.clients[id]
	if ok {
		delete(m.clients, id)
	}
	m.mu.Unlock()

	if !ok {
		return fmt.Errorf("client %s not found", id)
	}
	return client.Close()
}

func (m *MCPManager) DisconnectAll() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var lastErr error
	for id, client := range m.clients {
		if err := client.Close(); err != nil {
			lastErr = err
		}
		delete(m.clients, id)
	}
	return lastErr
}

func NewMCPClientFromServer(srv domainServer) (*MCPClient, error) {
	switch srv.Transport {
	case "stdio":
		return NewMCPClient(NewStdioTransport(srv.Command, srv.Args, srv.Env)), nil
	case "sse":
		return NewMCPClient(NewSSETransport(srv.URL)), nil
	default:
		return nil, fmt.Errorf("unsupported transport: %s", srv.Transport)
	}
}

type domainServer struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Transport string   `json:"transport"`
	Command   string   `json:"command,omitempty"`
	Args      []string `json:"args,omitempty"`
	URL       string   `json:"url,omitempty"`
	Env       []string `json:"env,omitempty"`
	Status    string   `json:"status"`
	Enabled   bool     `json:"enabled"`
}

func toRawMessage(v any) json.RawMessage {
	data, _ := json.Marshal(v)
	return data
}
