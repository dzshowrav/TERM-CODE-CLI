package session_test

import (
	"testing"
	"time"

	"termcode/internal/domain/session"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		sessionName string
		providerID  string
		modelID     string
	}{
		{
			name:        "creates active session",
			sessionName: "Test Session",
			providerID:  "prov-1",
			modelID:     "model-1",
		},
		{
			name:        "creates session with empty name",
			sessionName: "",
			providerID:  "prov-2",
			modelID:     "model-2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := session.New(tt.sessionName, tt.providerID, tt.modelID)
			if s.ID == "" {
				t.Error("expected non-empty ID")
			}
			if s.Name != tt.sessionName {
				t.Errorf("expected Name=%q", tt.sessionName)
			}
			if s.ProviderID != tt.providerID {
				t.Errorf("expected ProviderID=%q", tt.providerID)
			}
			if s.ModelID != tt.modelID {
				t.Errorf("expected ModelID=%q", tt.modelID)
			}
			if s.Status != session.StatusActive {
				t.Errorf("expected Status=%q", session.StatusActive)
			}
		})
	}
}

func TestNew_Defaults(t *testing.T) {
	s := session.New("Test", "p1", "m1")
	if s.MessageCnt != 0 {
		t.Errorf("expected MessageCnt=0, got %d", s.MessageCnt)
	}
	if s.TokenIn != 0 {
		t.Errorf("expected TokenIn=0, got %d", s.TokenIn)
	}
	if s.TokenOut != 0 {
		t.Errorf("expected TokenOut=0, got %d", s.TokenOut)
	}
	if s.Workspace != "" {
		t.Errorf("expected Workspace=empty, got %q", s.Workspace)
	}
}

func TestNew_Timestamps(t *testing.T) {
	before := time.Now()
	s := session.New("Test", "p1", "m1")
	after := time.Now()
	if s.CreatedAt.Before(before) || s.CreatedAt.After(after) {
		t.Error("CreatedAt should be near current time")
	}
	if s.UpdatedAt.Before(before) || s.UpdatedAt.After(after) {
		t.Error("UpdatedAt should be near current time")
	}
}
