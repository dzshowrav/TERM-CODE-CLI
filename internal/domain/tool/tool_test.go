package tool_test

import (
	"testing"

	"termcode/internal/domain/tool"
)

func TestTool_Fields(t *testing.T) {
	tests := []struct {
		name        string
		toolName    string
		description string
	}{
		{name: "bash tool", toolName: "bash", description: "Execute shell commands"},
		{name: "read tool", toolName: "read", description: "Read file contents"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := tool.Tool{
				Name:        tt.toolName,
				Description: tt.description,
				InputSchema: map[string]any{"type": "object"},
			}
			if tr.Name != tt.toolName {
				t.Errorf("expected Name=%q", tt.toolName)
			}
			if tr.Description != tt.description {
				t.Errorf("expected Description=%q", tt.description)
			}
			if tr.InputSchema == nil {
				t.Error("expected non-nil InputSchema")
			}
		})
	}
}

func TestResult_Defaults(t *testing.T) {
	r := tool.Result{
		Tool:   "bash",
		Input:  "ls",
		Output: "file1.txt",
		Status: tool.StatusSuccess,
	}
	if r.Tool != "bash" {
		t.Errorf("expected Tool=bash, got %q", r.Tool)
	}
	if r.Output != "file1.txt" {
		t.Errorf("expected Output=file1.txt, got %q", r.Output)
	}
	if r.Status != tool.StatusSuccess {
		t.Errorf("expected Status=success, got %q", r.Status)
	}
}

func TestResult_Error(t *testing.T) {
	r := tool.Result{
		Tool:   "bash",
		Input:  "invalid",
		Error:  "command not found",
		Status: tool.StatusFailed,
	}
	if r.Error != "command not found" {
		t.Errorf("expected Error=command not found, got %q", r.Error)
	}
	if r.Status != tool.StatusFailed {
		t.Errorf("expected Status=failed, got %q", r.Status)
	}
}

func TestStatusValues(t *testing.T) {
	statuses := []tool.Status{
		tool.StatusPending,
		tool.StatusRunning,
		tool.StatusApproved,
		tool.StatusDenied,
		tool.StatusSuccess,
		tool.StatusFailed,
	}
	for _, s := range statuses {
		t.Run(string(s), func(t *testing.T) {
			if s == "" {
				t.Error("expected non-empty status")
			}
		})
	}
}
