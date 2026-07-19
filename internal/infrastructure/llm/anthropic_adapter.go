package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"termcode/pkg/apitypes"
)

type AnthropicAdapter struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewAnthropicAdapter(baseURL, apiKey string) *AnthropicAdapter {
	u := strings.TrimRight(baseURL, "/")
	u = strings.TrimSuffix(u, "/v1")
	return &AnthropicAdapter{
		baseURL: u,
		apiKey:  apiKey,
		client: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

type anthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicRequest struct {
	Model       string             `json:"model"`
	Messages    []anthropicMessage `json:"messages"`
	System      string             `json:"system,omitempty"`
	MaxTokens   int                `json:"max_tokens"`
	Stream      bool               `json:"stream"`
	Temperature float64            `json:"temperature,omitempty"`
}

type anthropicContent struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type anthropicResponse struct {
	ID         string             `json:"id"`
	Type       string             `json:"type"`
	Role       string             `json:"role"`
	Content    []anthropicContent `json:"content"`
	Model      string             `json:"model"`
	StopReason string             `json:"stop_reason"`
	Usage      *struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage,omitempty"`
}

func (a *AnthropicAdapter) Chat(ctx context.Context, req *apitypes.ChatRequest) (*apitypes.ChatResponse, error) {
	antReq := a.toAnthropicRequest(req)
	body, err := json.Marshal(antReq)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, a.baseURL+"/v1/messages", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", a.apiKey)
	httpReq.Header.Set("anthropic-version", "2023-06-01")

	resp, err := a.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		msg := string(respBody)
		if isHTML(msg) {
			msg = "server returned non-JSON response — check your Anthropic URL"
		}
		return nil, fmt.Errorf("Anthropic error (HTTP %d): %s", resp.StatusCode, msg)
	}

	var antResp anthropicResponse
	if err := json.NewDecoder(resp.Body).Decode(&antResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return a.fromAnthropicResponse(&antResp, req.Model), nil
}

func (a *AnthropicAdapter) ChatStream(ctx context.Context, req *apitypes.ChatRequest) (<-chan apitypes.StreamChunk, <-chan error) {
	// Delegate to OpenAI-compatible streaming for now
	adapter := NewOpenAIAdapter(a.baseURL, a.apiKey)
	return adapter.ChatStream(ctx, req)
}

func (a *AnthropicAdapter) toAnthropicRequest(req *apitypes.ChatRequest) *anthropicRequest {
	var system string
	messages := make([]anthropicMessage, 0, len(req.Messages))
	for _, m := range req.Messages {
		if m.Role == "system" {
			system = m.Content
			continue
		}
		messages = append(messages, anthropicMessage{Role: m.Role, Content: m.Content})
	}

	maxTokens := req.MaxTokens
	if maxTokens <= 0 {
		maxTokens = 4096
	}

	return &anthropicRequest{
		Model:       req.Model,
		Messages:    messages,
		System:      system,
		MaxTokens:   maxTokens,
		Temperature: req.Temperature,
	}
}

func (a *AnthropicAdapter) fromAnthropicResponse(resp *anthropicResponse, model string) *apitypes.ChatResponse {
	content := ""
	for _, c := range resp.Content {
		if c.Type == "text" {
			content += c.Text
		}
	}

	chatResp := &apitypes.ChatResponse{
		ID:      resp.ID,
		Model:   model,
		Created: time.Now().Unix(),
		Choices: []apitypes.Choice{
			{
				Message: apitypes.ChatMessage{
					Role:    resp.Role,
					Content: content,
				},
				FinishReason: resp.StopReason,
			},
		},
	}

	if resp.Usage != nil {
		chatResp.Usage = &apitypes.Usage{
			PromptTokens:     resp.Usage.InputTokens,
			CompletionTokens: resp.Usage.OutputTokens,
			TotalTokens:      resp.Usage.InputTokens + resp.Usage.OutputTokens,
		}
	}

	return chatResp
}
