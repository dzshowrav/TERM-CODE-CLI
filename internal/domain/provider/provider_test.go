package provider_test

import (
	"testing"

	"termcode/internal/domain/provider"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		nameArg     string
		baseURL     string
		apiKey      string
		description string
	}{
		{
			name:        "creates provider with all fields",
			nameArg:     "OpenAI",
			baseURL:     "https://api.openai.com/v1",
			apiKey:      "sk-test",
			description: "Test provider",
		},
		{
			name:    "creates provider with minimal fields",
			nameArg: "Local",
			baseURL: "http://localhost:8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := provider.New(tt.nameArg, tt.baseURL, tt.apiKey, tt.description)
			if p == nil {
				t.Fatal("expected non-nil provider")
			}
			if p.ID == "" {
				t.Error("expected non-empty ID")
			}
			if p.Name != tt.nameArg {
				t.Errorf("expected Name=%q, got %q", tt.nameArg, p.Name)
			}
			if p.BaseURL != tt.baseURL {
				t.Errorf("expected BaseURL=%q, got %q", tt.baseURL, p.BaseURL)
			}
			if p.Status != "disconnected" {
				t.Errorf("expected Status=disconnected, got %q", p.Status)
			}
			if p.CreatedAt.IsZero() {
				t.Error("expected non-zero CreatedAt")
			}
			if p.UpdatedAt.IsZero() {
				t.Error("expected non-zero UpdatedAt")
			}
		})
	}
}

func TestNew_APIKey(t *testing.T) {
	p := provider.New("Test", "http://localhost", "secret-key", "")
	if p.APIKey != "secret-key" {
		t.Errorf("expected APIKey=secret-key, got %q", p.APIKey)
	}
}

func TestNew_Defaults(t *testing.T) {
	p := provider.New("Test", "http://localhost", "", "")
	if p.Priority != 0 {
		t.Errorf("expected Priority=0, got %d", p.Priority)
	}
	if p.IsDefault {
		t.Error("expected IsDefault=false")
	}
}
