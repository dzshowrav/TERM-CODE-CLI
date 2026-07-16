package llm

import (
	"termcode/internal/domain/session"
	"termcode/pkg/apitypes"
)

type RequestBuilder struct{}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}

type BuildOptions struct {
	Model        string
	Messages     []session.Message
	SystemPrompt string
	Stream       bool
	Temperature  float64
	MaxTokens    int
	Tools        []map[string]any
}

func (b *RequestBuilder) Build(opts BuildOptions) *apitypes.ChatRequest {
	msgs := make([]apitypes.ChatMessage, 0)

	if opts.SystemPrompt != "" {
		msgs = append(msgs, apitypes.ChatMessage{
			Role:    "system",
			Content: opts.SystemPrompt,
		})
	}

	for _, m := range opts.Messages {
		msgs = append(msgs, apitypes.ChatMessage{
			Role:    string(m.Role),
			Content: m.Content,
		})
	}

	req := &apitypes.ChatRequest{
		Model:       opts.Model,
		Messages:    msgs,
		Stream:      opts.Stream,
		Temperature: opts.Temperature,
		MaxTokens:   opts.MaxTokens,
	}

	if len(opts.Tools) > 0 {
		req.Tools = opts.Tools
	}

	return req
}
