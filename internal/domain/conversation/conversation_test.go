package conversation_test

import (
	"testing"

	"termcode/internal/domain/conversation"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		sessionID string
	}{
		{name: "creates conversation", sessionID: "session-1"},
		{name: "creates with empty session id", sessionID: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := conversation.New(tt.sessionID)
			if c.ID == "" {
				t.Error("expected non-empty ID")
			}
			if c.SessionID != tt.sessionID {
				t.Errorf("expected SessionID=%q", tt.sessionID)
			}
			if c.Status != conversation.StatusActive {
				t.Errorf("expected Status=%q", conversation.StatusActive)
			}
		})
	}
}

func TestNew_Defaults(t *testing.T) {
	c := conversation.New("s1")
	if c.MsgCount != 0 {
		t.Errorf("expected MsgCount=0, got %d", c.MsgCount)
	}
	if c.TokenIn != 0 {
		t.Errorf("expected TokenIn=0, got %d", c.TokenIn)
	}
	if c.TokenOut != 0 {
		t.Errorf("expected TokenOut=0, got %d", c.TokenOut)
	}
}

func TestAddTokens(t *testing.T) {
	c := conversation.New("s1")
	c.AddTokens(100, 50)
	if c.TokenIn != 100 {
		t.Errorf("expected TokenIn=100, got %d", c.TokenIn)
	}
	if c.TokenOut != 50 {
		t.Errorf("expected TokenOut=50, got %d", c.TokenOut)
	}
	if c.MsgCount != 1 {
		t.Errorf("expected MsgCount=1, got %d", c.MsgCount)
	}

	c.AddTokens(200, 100)
	if c.TokenIn != 300 {
		t.Errorf("expected TokenIn=300, got %d", c.TokenIn)
	}
	if c.TokenOut != 150 {
		t.Errorf("expected TokenOut=150, got %d", c.TokenOut)
	}
	if c.MsgCount != 2 {
		t.Errorf("expected MsgCount=2, got %d", c.MsgCount)
	}
}

func TestIsEmpty(t *testing.T) {
	t.Run("new conversation is empty", func(t *testing.T) {
		c := conversation.New("s1")
		if !c.IsEmpty() {
			t.Error("expected IsEmpty=true for new conversation")
		}
	})

	t.Run("after adding tokens is not empty", func(t *testing.T) {
		c := conversation.New("s1")
		c.AddTokens(10, 5)
		if c.IsEmpty() {
			t.Error("expected IsEmpty=false after adding tokens")
		}
	})
}
