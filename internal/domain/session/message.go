package session

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleSystem    Role = "system"
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleTool      Role = "tool"
)

type Message struct {
	ID        string    `json:"id"`
	SessionID string    `json:"session_id"`
	Role      Role      `json:"role"`
	Content   string    `json:"content"`
	Reasoning string    `json:"reasoning,omitempty"`
	ToolCall  string    `json:"tool_call,omitempty"`
	ToolRes   string    `json:"tool_result,omitempty"`
	TokenIn   int       `json:"tokens_in"`
	TokenOut  int       `json:"tokens_out"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMessage(sessionID string, role Role, content string) *Message {
	return &Message{
		ID:        uuid.New().String(),
		SessionID: sessionID,
		Role:      role,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (m *Message) IsTool() bool {
	return m.Role == RoleTool
}

func (m *Message) IsUser() bool {
	return m.Role == RoleUser
}

func (m *Message) IsAssistant() bool {
	return m.Role == RoleAssistant
}
