package agent_test

import (
	"testing"

	"termcode/internal/domain/agent"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name         string
		agentName    string
		description  string
		systemPrompt string
	}{
		{
			name:         "creates agent with all fields",
			agentName:    "coder",
			description:  "Coding assistant",
			systemPrompt: "You are a coding assistant.",
		},
		{
			name:      "creates agent with minimal fields",
			agentName: "helper",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := agent.New(tt.agentName, tt.description, tt.systemPrompt)
			if a.ID == "" {
				t.Error("expected non-empty ID")
			}
			if a.Name != tt.agentName {
				t.Errorf("expected Name=%q", tt.agentName)
			}
			if a.Description != tt.description {
				t.Errorf("expected Description=%q", tt.description)
			}
			if a.SystemPrompt != tt.systemPrompt {
				t.Errorf("expected SystemPrompt=%q", tt.systemPrompt)
			}
			if a.IsDefault {
				t.Error("expected IsDefault=false")
			}
		})
	}
}

func TestNew_Tools(t *testing.T) {
	a := agent.New("test", "", "")
	if a.Tools != nil {
		t.Errorf("expected Tools=nil, got %v", a.Tools)
	}
}
