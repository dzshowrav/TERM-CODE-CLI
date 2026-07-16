package conversation

import (
	"context"
	"encoding/json"
	"fmt"

	toolapp "termcode/internal/application/tool"
	"termcode/internal/domain/provider"
	"termcode/internal/domain/session"
	"termcode/internal/domain/tool"
	"termcode/internal/infrastructure/llm"
	"termcode/pkg/apitypes"
)

const maxToolRounds = 10

type Service struct {
	providerSvc ProviderService
	toolSvc     *toolapp.Service
	modelID     string
}

type ProviderService interface {
	GetDefault(ctx context.Context) (*provider.Provider, error)
	GetByID(ctx context.Context, id string) (*provider.Provider, error)
	DecryptAPIKey(encrypted string) (string, error)
}

func NewService(providerSvc ProviderService) *Service {
	return &Service{
		providerSvc: providerSvc,
		toolSvc:     toolapp.NewService(),
		modelID:     "gpt-4o",
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

func (s *Service) SendMessage(ctx context.Context, prompt string, history []session.Message, onChunk func(string, bool)) (string, error) {
	p, err := s.providerSvc.GetDefault(ctx)
	if err != nil {
		return "", fmt.Errorf("no provider: %w", err)
	}

	apiKey := p.APIKey
	if apiKey != "" {
		decrypted, err := s.providerSvc.DecryptAPIKey(apiKey)
		if err == nil {
			apiKey = decrypted
		}
	}

	adapter := llm.NewOpenAIAdapter(p.BaseURL, apiKey)

	messages := buildMessages(history, prompt)
	toolDefs := buildToolDefs(s.toolSvc.AvailableTools())

	return s.converse(ctx, adapter, messages, toolDefs, onChunk, 0)
}

func (s *Service) converse(ctx context.Context, adapter *llm.OpenAIAdapter, messages []apitypes.ChatMessage, toolDefs []map[string]any, onChunk func(string, bool), round int) (string, error) {
	if round > maxToolRounds {
		return "", fmt.Errorf("too many tool call rounds (%d)", maxToolRounds)
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
		return "", fmt.Errorf("chat request: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response")
	}

	choice := resp.Choices[0]
	content := choice.Message.Content
	finishReason := choice.FinishReason

	if finishReason == "tool_calls" && len(choice.Message.ToolCalls) > 0 {
		messages = append(messages, apitypes.ChatMessage{
			Role:      "assistant",
			Content:   content,
			ToolCalls: choice.Message.ToolCalls,
		})

		for _, tc := range choice.Message.ToolCalls {
			var args map[string]any
			if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err != nil {
				args = map[string]any{}
			}

			result := s.toolSvc.Execute(ctx, tc.Function.Name, args)

			resultJSON, _ := json.Marshal(result)
			messages = append(messages, apitypes.ChatMessage{
				Role:       "tool",
				Content:    string(resultJSON),
				ToolCallID: tc.ID,
			})
		}

		return s.converse(ctx, adapter, messages, toolDefs, onChunk, round+1)
	}

	if onChunk != nil {
		onChunk(content, true)
	}

	return content, nil
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
