package conversation

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	cost "termcode/internal/application/cost"
	toolapp "termcode/internal/application/tool"
	domainllm "termcode/internal/domain/llm"
	"termcode/internal/domain/session"
	"termcode/internal/domain/tool"
	"termcode/internal/infrastructure/llm"
	"termcode/pkg/apitypes"
)

type ToolEventType int

const (
	ToolQueued ToolEventType = iota
	ToolInitializing
	ToolConnecting
	ToolStarted
	ToolOutput
	ToolCompleted
)

type ToolEvent struct {
	Type     ToolEventType
	Index    int
	Name     string
	Args     string
	Output   string
	Status   string
	Duration int64
	Error    string
}

type ToolEventCallback func(ToolEvent)

type ChatResult struct {
	Content          string
	ReasoningContent string
	InputTokens      int
	OutputTokens     int
	TotalTokens      int
}

type Service struct {
	router         domainllm.Router
	toolSvc        *toolapp.Service
	modelID        string
	costEngine     *cost.Engine
	checkpointPath string
	mu             sync.Mutex
}

func NewService(router domainllm.Router) *Service {
	svc := toolapp.NewService()
	svc.SetPermissionChecker(toolapp.NewPermissionChecker())
	return &Service{
		router:  router,
		toolSvc: svc,
		modelID: "gpt-4o",
	}
}

func (s *Service) ToolService() *toolapp.Service {
	return s.toolSvc
}

func (s *Service) SetPermissionRequestFunc(f func(toolName, args string, resultCh chan<- string)) {
	if pc, ok := s.toolSvc.PermissionChecker().(*toolapp.StorePermissionChecker); ok {
		pc.SetRequestFunc(f)
	}
}

func (s *Service) SetCostEngine(ce *cost.Engine) {
	s.costEngine = ce
}

func (s *Service) SetModelID(modelID string) {
	s.modelID = modelID
}

func (s *Service) ModelID() string {
	return s.modelID
}

func (s *Service) SetCheckpointPath(path string) {
	s.checkpointPath = path
}

func (s *Service) SaveCheckpoint(state interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.checkpointPath == "" {
		return nil
	}
	data, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("marshal checkpoint: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(s.checkpointPath), 0o755); err != nil {
		return fmt.Errorf("mkdir checkpoint: %w", err)
	}
	tmpPath := s.checkpointPath + ".tmp"
	if err := os.WriteFile(tmpPath, data, 0o644); err != nil {
		return fmt.Errorf("write checkpoint: %w", err)
	}
	return os.Rename(tmpPath, s.checkpointPath)
}

func (s *Service) LoadCheckpoint(state interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.checkpointPath == "" {
		return fmt.Errorf("no checkpoint path set")
	}
	data, err := os.ReadFile(s.checkpointPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("no checkpoint found")
		}
		return fmt.Errorf("read checkpoint: %w", err)
	}
	return json.Unmarshal(data, state)
}

func (s *Service) ClearCheckpoint() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.checkpointPath != "" {
		return os.Remove(s.checkpointPath)
	}
	return nil
}

func (s *Service) Tools() []tool.Tool {
	return s.toolSvc.AvailableTools()
}

func (s *Service) SendMessage(ctx context.Context, prompt string, history []session.Message, onChunk func(content string, reasoning bool, done bool), onToolEvent ToolEventCallback) (*ChatResult, error) {
	endpoint, err := s.router.Resolve(ctx, s.modelID)
	if err != nil {
		return nil, fmt.Errorf("resolve endpoint: %w", err)
	}

	adapter := llm.NewOpenAIAdapter(endpoint.BaseURL, endpoint.APIKey)

	messages := buildMessages(history, prompt)
	toolDefs := buildToolDefs(s.toolSvc.AvailableTools())

	return s.converse(ctx, adapter, messages, toolDefs, onChunk, onToolEvent, 0)
}

func (s *Service) converse(ctx context.Context, adapter *llm.OpenAIAdapter, messages []apitypes.ChatMessage, toolDefs []map[string]any, onChunk func(content string, reasoning bool, done bool), onToolEvent ToolEventCallback, round int) (*ChatResult, error) {
	if len(toolDefs) == 0 {
		return s.converseStream(ctx, adapter, messages, onChunk)
	}

	req := &apitypes.ChatRequest{
		Model:       s.modelID,
		Messages:    messages,
		Stream:      false,
		Temperature: 0.7,
		MaxTokens:   4096,
		Tools:       toolDefs,
	}

	resp, err := adapter.Chat(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("chat request: %w", err)
	}

	if len(resp.Choices) == 0 {
		return &ChatResult{}, nil
	}

	choice := resp.Choices[0]
	content := choice.Message.Content
	reasoningContent := choice.ReasoningText()
	if reasoningContent == "" {
		reasoningContent = choice.Message.ReasoningText()
	}
	finishReason := choice.FinishReason

	inputTokens := 0
	outputTokens := 0
	if resp.Usage != nil {
		inputTokens = resp.Usage.PromptTokens
		outputTokens = resp.Usage.CompletionTokens
	}

	if finishReason == "tool_calls" && len(choice.Message.ToolCalls) > 0 {
		if onChunk != nil {
			if reasoningContent != "" {
				onChunk(reasoningContent, true, false)
			}
			if content != "" {
				onChunk(content, false, false)
			}
		}

		messages = append(messages, apitypes.ChatMessage{
			Role:      "assistant",
			Content:   content,
			ToolCalls: choice.Message.ToolCalls,
		})

		buildArgs := func(tc apitypes.ToolCall) (map[string]any, string) {
			var args map[string]any
			if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err != nil {
				args = map[string]any{}
			}
			argsJSON, _ := json.Marshal(args)
			return args, string(argsJSON)
		}

		for i, tc := range choice.Message.ToolCalls {
			_, argsJSON := buildArgs(tc)
			if onToolEvent != nil {
				onToolEvent(ToolEvent{
					Type:  ToolQueued,
					Index: i,
					Name:  tc.Function.Name,
					Args:  argsJSON,
				})
			}
		}

		for i, tc := range choice.Message.ToolCalls {
			args, argsJSON := buildArgs(tc)

			if onToolEvent != nil {
				onToolEvent(ToolEvent{
					Type:  ToolInitializing,
					Index: i,
					Name:  tc.Function.Name,
				})
			}

			if onToolEvent != nil {
				onToolEvent(ToolEvent{
					Type:  ToolConnecting,
					Index: i,
					Name:  tc.Function.Name,
				})
			}

			if onToolEvent != nil {
				onToolEvent(ToolEvent{
					Type:  ToolStarted,
					Index: i,
					Name:  tc.Function.Name,
					Args:  argsJSON,
				})
			}

			s.toolSvc.SetOutputCallback(func(output string) {
				if onToolEvent != nil {
					onToolEvent(ToolEvent{
						Type:   ToolOutput,
						Index:  i,
						Name:   tc.Function.Name,
						Args:   argsJSON,
						Output: output,
					})
				}
			})
			result := s.toolSvc.Execute(ctx, tc.Function.Name, args)
			s.toolSvc.SetOutputCallback(nil)

			if onToolEvent != nil {
				ev := ToolEvent{
					Type:     ToolCompleted,
					Index:    i,
					Name:     tc.Function.Name,
					Args:     argsJSON,
					Duration: result.Duration,
				}
				if result.Status == tool.StatusSuccess {
					ev.Status = "completed"
					ev.Output = result.Output
				} else {
					ev.Status = "failed"
					ev.Output = result.Output
					ev.Error = result.Error
				}
				onToolEvent(ev)
			}

			resultJSON, _ := json.Marshal(result)
			messages = append(messages, apitypes.ChatMessage{
				Role:       "tool",
				Content:    string(resultJSON),
				ToolCallID: tc.ID,
			})
		}

		if err := s.SaveCheckpoint(messages); err != nil {
			return nil, fmt.Errorf("save checkpoint: %w", err)
		}

		next, err := s.converse(ctx, adapter, messages, toolDefs, onChunk, onToolEvent, round+1)
		if err != nil {
			return nil, err
		}
		next.InputTokens += inputTokens
		next.OutputTokens += outputTokens
		next.TotalTokens = next.InputTokens + next.OutputTokens
		return next, nil
	}

	if err := s.SaveCheckpoint(messages); err != nil {
		return nil, fmt.Errorf("save checkpoint: %w", err)
	}

	if onChunk != nil {
		if reasoningContent != "" {
			onChunk(reasoningContent, true, false)
		}
		onChunk(content, false, true)
	}

	return &ChatResult{
		Content:          content,
		ReasoningContent: reasoningContent,
		InputTokens:      inputTokens,
		OutputTokens:     outputTokens,
		TotalTokens:      inputTokens + outputTokens,
	}, nil
}

func (s *Service) converseStream(ctx context.Context, adapter *llm.OpenAIAdapter, messages []apitypes.ChatMessage, onChunk func(content string, reasoning bool, done bool)) (*ChatResult, error) {
	streamCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	req := &apitypes.ChatRequest{
		Model:       s.modelID,
		Messages:    messages,
		Temperature: 0.7,
		MaxTokens:   4096,
	}

	if onChunk == nil {
		return &ChatResult{}, nil
	}

	chunks, errs := adapter.ChatStream(streamCtx, req)

	var contentBuf, reasoningBuf strings.Builder
	inputTokens, outputTokens := 0, 0
	budgetWarningIssued := false

	streamDone := false
	for !streamDone {
		select {
		case chunk, ok := <-chunks:
			if !ok {
				streamDone = true
				continue
			}
			for _, c := range chunk.Choices {
				rt := c.Delta.ReasoningText()
				if rt != "" {
					reasoningBuf.WriteString(rt)
					onChunk(rt, true, false)
				}
				if c.Delta.Content != "" {
					contentBuf.WriteString(c.Delta.Content)
					outputTokens++
					onChunk(c.Delta.Content, false, false)
				}
				if c.FinishReason == "stop" || c.FinishReason == "length" {
					streamDone = true
				}
			}

			if s.costEngine != nil && !budgetWarningIssued && outputTokens > 0 && outputTokens%100 == 0 {
				remaining := s.costEngine.BudgetRemaining()
				if remaining > 0 && remaining < 0.001 {
					budgetWarningIssued = true
					warning := fmt.Sprintf("\n\n[budget-warning: $%.4f remaining]", remaining)
					contentBuf.WriteString(warning)
					onChunk(warning, false, false)
				}
			}

		case err := <-errs:
			if err != nil {
				return nil, fmt.Errorf("stream error: %w", err)
			}
			streamDone = true

		case <-streamCtx.Done():
			return nil, streamCtx.Err()
		}
	}

	content := contentBuf.String()
	reasoningContent := reasoningBuf.String()

	if onChunk != nil {
		onChunk("", false, true)
	}

	return &ChatResult{
		Content:          content,
		ReasoningContent: reasoningContent,
		InputTokens:      inputTokens,
		OutputTokens:     outputTokens,
		TotalTokens:      inputTokens + outputTokens,
	}, nil
}

func buildMessages(history []session.Message, prompt string) []apitypes.ChatMessage {
	msgs := make([]apitypes.ChatMessage, 0, len(history)+1)
	for _, m := range history {
		msgs = append(msgs, apitypes.ChatMessage{
			Role:    string(m.Role),
			Content: m.Content,
		})
	}
	msgs = append(msgs, apitypes.ChatMessage{
		Role:    "user",
		Content: prompt,
	})
	return msgs
}

func buildToolDefs(tools []tool.Tool) []map[string]any {
	defs := make([]map[string]any, 0, len(tools))
	for _, t := range tools {
		defs = append(defs, map[string]any{
			"type": "function",
			"function": map[string]any{
				"name":        t.Name,
				"description": t.Description,
				"parameters":  t.InputSchema,
			},
		})
	}
	return defs
}
