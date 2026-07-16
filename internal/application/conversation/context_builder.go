package conversation

import (
	"termcode/internal/domain/session"
	"termcode/pkg/apitypes"
)

type ContextBuilder struct{}

func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{}
}

func (b *ContextBuilder) Build(systemPrompt string, history []session.Message, newPrompt string) []apitypes.ChatMessage {
	msgs := make([]apitypes.ChatMessage, 0)

	if systemPrompt != "" {
		msgs = append(msgs, apitypes.ChatMessage{
			Role:    "system",
			Content: systemPrompt,
		})
	}

	for _, m := range history {
		msgs = append(msgs, apitypes.ChatMessage{
			Role:    string(m.Role),
			Content: m.Content,
		})
	}

	msgs = append(msgs, apitypes.ChatMessage{
		Role:    "user",
		Content: newPrompt,
	})

	return msgs
}
