package session_test

import (
	"testing"

	"termcode/internal/domain/session"
)

func TestNewMessage(t *testing.T) {
	tests := []struct {
		name      string
		sessionID string
		role      session.Role
		content   string
	}{
		{name: "user message", sessionID: "s1", role: session.RoleUser, content: "hello"},
		{name: "assistant message", sessionID: "s1", role: session.RoleAssistant, content: "hi"},
		{name: "system message", sessionID: "s1", role: session.RoleSystem, content: "be helpful"},
		{name: "tool message", sessionID: "s1", role: session.RoleTool, content: "result"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := session.NewMessage(tt.sessionID, tt.role, tt.content)
			if m.ID == "" {
				t.Error("expected non-empty ID")
			}
			if m.SessionID != tt.sessionID {
				t.Errorf("expected SessionID=%q", tt.sessionID)
			}
			if m.Role != tt.role {
				t.Errorf("expected Role=%q", tt.role)
			}
			if m.Content != tt.content {
				t.Errorf("expected Content=%q", tt.content)
			}
		})
	}
}

func TestMessage_RoleChecks(t *testing.T) {
	tests := []struct {
		name   string
		role   session.Role
		isTool bool
		isUser bool
		isAsst bool
	}{
		{name: "user", role: session.RoleUser, isUser: true},
		{name: "assistant", role: session.RoleAssistant, isAsst: true},
		{name: "tool", role: session.RoleTool, isTool: true},
		{name: "system", role: session.RoleSystem},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := session.NewMessage("s1", tt.role, "")
			if got := m.IsTool(); got != tt.isTool {
				t.Errorf("IsTool() = %v, want %v", got, tt.isTool)
			}
			if got := m.IsUser(); got != tt.isUser {
				t.Errorf("IsUser() = %v, want %v", got, tt.isUser)
			}
			if got := m.IsAssistant(); got != tt.isAsst {
				t.Errorf("IsAssistant() = %v, want %v", got, tt.isAsst)
			}
		})
	}
}
