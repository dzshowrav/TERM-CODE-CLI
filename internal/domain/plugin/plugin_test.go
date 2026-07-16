package plugin_test

import (
	"testing"

	"termcode/internal/domain/plugin"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		pluginName  string
		version     string
		author      string
		description string
	}{
		{
			name:        "creates plugin with all fields",
			pluginName:  "formatter",
			version:     "1.0.0",
			author:      "Alice",
			description: "Code formatter plugin",
		},
		{
			name:       "creates plugin with minimal fields",
			pluginName: "linter",
			version:    "0.1.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := plugin.New(tt.pluginName, tt.version, tt.author, tt.description)
			if p.ID == "" {
				t.Error("expected non-empty ID")
			}
			if p.Name != tt.pluginName {
				t.Errorf("expected Name=%q", tt.pluginName)
			}
			if p.Version != tt.version {
				t.Errorf("expected Version=%q", tt.version)
			}
			if p.Author != tt.author {
				t.Errorf("expected Author=%q", tt.author)
			}
			if p.Description != tt.description {
				t.Errorf("expected Description=%q", tt.description)
			}
		})
	}
}

func TestNew_Defaults(t *testing.T) {
	p := plugin.New("test", "1.0", "", "")
	if p.Status != plugin.StatusInactive {
		t.Errorf("expected Status=inactive, got %q", p.Status)
	}
	if !p.Enabled {
		t.Error("expected Enabled=true")
	}
}
