package mcp_test

import (
	"testing"

	"termcode/internal/domain/mcp"
)

func TestNewStdio(t *testing.T) {
	tests := []struct {
		name    string
		svrName string
		command string
		args    []string
	}{
		{name: "with args", svrName: "server1", command: "node", args: []string{"index.js"}},
		{name: "without args", svrName: "server2", command: "python", args: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := mcp.NewStdio(tt.svrName, tt.command, tt.args)
			if s.ID == "" {
				t.Error("expected non-empty ID")
			}
			if s.Name != tt.svrName {
				t.Errorf("expected Name=%q", tt.svrName)
			}
			if s.Command != tt.command {
				t.Errorf("expected Command=%q", tt.command)
			}
			if s.Transport != mcp.TransportStdio {
				t.Errorf("expected Transport=stdio, got %q", s.Transport)
			}
			if s.Status != mcp.StatusDisconnected {
				t.Errorf("expected Status=disconnected, got %q", s.Status)
			}
			if !s.Enabled {
				t.Error("expected Enabled=true")
			}
		})
	}
}

func TestNewSSE(t *testing.T) {
	tests := []struct {
		name    string
		svrName string
		url     string
	}{
		{name: "with url", svrName: "remote", url: "https://example.com/sse"},
		{name: "with localhost", svrName: "local", url: "http://localhost:8080/sse"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := mcp.NewSSE(tt.svrName, tt.url)
			if s.Transport != mcp.TransportSSE {
				t.Errorf("expected Transport=sse, got %q", s.Transport)
			}
			if s.URL != tt.url {
				t.Errorf("expected URL=%q", tt.url)
			}
			if s.Command != "" {
				t.Errorf("expected empty Command for SSE, got %q", s.Command)
			}
		})
	}
}
