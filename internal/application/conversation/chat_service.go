package conversation

import (
	"context"
	"encoding/json"
	"fmt"

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
	Content      string
	InputTokens  int
	OutputTokens int
	TotalTokens  int
}

const maxToolRounds = 10

type Service struct {
	router  domainllm.Router
	toolSvc *toolapp.Service
	modelID string
}

func NewService(router domainllm.Router) *Service {
	return &Service{
		router:  router,
		toolSvc: toolapp.NewService(),
		modelID: "gpt-4o",
	}
}

func (s *Service) SetModelID(modelID string) {
	s.modelID = modelID
}

func (s *Service) ModelID() string {
	return s.modelID
}

func (s *Service) Tools() []tool.Tool {
	return s.toolSvc.AvailableTools()
}

func (s *Service) SendMessage(ctx context.Context, prompt string, history []session.Message, onChunk func(string, bool), onToolEvent ToolEventCallback) (*ChatResult, error) {
	endpoint, err := s.router.Resolve(ctx, s.modelID)
	if err != nil {
		return nil, fmt.Errorf("resolve endpoint: %w", err)
	}

	adapter := llm.NewOpenAIAdapter(endpoint.BaseURL, endpoint.APIKey)

	messages := buildMessages(history, prompt)
	toolDefs := buildToolDefs(s.toolSvc.AvailableTools())

	return s.converse(ctx, adapter, messages, toolDefs, onChunk, onToolEvent, 0)
}

func (s *Service) converse(ctx context.Context, adapter *llm.OpenAIAdapter, messages []apitypes.ChatMessage, toolDefs []map[string]any, onChunk func(string, bool), onToolEvent ToolEventCallback, round int) (*ChatResult, error) {
	if round > maxToolRounds {
		return nil, fmt.Errorf("too many tool call rounds (%d)", maxToolRounds)
	}

	req := &apitypes.ChatRequest{
		Model:       s.modelID,
		Messages:    messages,
		Stream:      false,
		Temperature: 0.7,
		MaxTokens:   4096,
	}

	if len(toolDefs) > 0 {
		req.Tools = toolDefs
	}

	resp, err := adapter.Chat(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("chat request: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no response")
	}

	choice := resp.Choices[0]
	content := choice.Message.Content
	finishReason := choice.FinishReason

	inputTokens := 0
	outputTokens := 0
	if resp.Usage != nil {
		inputTokens = resp.Usage.PromptTokens
		outputTokens = resp.Usage.CompletionTokens
	}

	if finishReason == "tool_calls" && len(choice.Message.ToolCalls) > 0 {
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

			result := s.toolSvc.Execute(ctx, tc.Function.Name, args)

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

		next, err := s.converse(ctx, adapter, messages, toolDefs, onChunk, onToolEvent, round+1)
		if err != nil {
			return nil, err
		}
		next.InputTokens += inputTokens
		next.OutputTokens += outputTokens
		next.TotalTokens = next.InputTokens + next.OutputTokens
		return next, nil
	}

	if onChunk != nil {
		onChunk(content, true)
	}

	return &ChatResult{
		Content:      content,
		InputTokens:  inputTokens,
		OutputTokens: outputTokens,
		TotalTokens:  inputTokens + outputTokens,
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
