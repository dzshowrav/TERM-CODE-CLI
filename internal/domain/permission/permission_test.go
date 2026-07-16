package permission_test

import (
	"testing"

	"termcode/internal/domain/permission"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		toolName string
		level    permission.Level
	}{
		{name: "always allow", toolName: "bash", level: permission.LevelAlwaysAllow},
		{name: "ask", toolName: "write", level: permission.LevelAsk},
		{name: "deny", toolName: "delete", level: permission.LevelDeny},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := permission.New(tt.toolName, tt.level)
			if e.ToolName != tt.toolName {
				t.Errorf("expected ToolName=%q", tt.toolName)
			}
			if e.Permission != tt.level {
				t.Errorf("expected Permission=%q", tt.level)
			}
			if e.UpdatedAt.IsZero() {
				t.Error("expected non-zero UpdatedAt")
			}
		})
	}
}

func TestIsAllowed(t *testing.T) {
	tests := []struct {
		name    string
		level   permission.Level
		allowed bool
		denied  bool
	}{
		{name: "always allow", level: permission.LevelAlwaysAllow, allowed: true, denied: false},
		{name: "allow once", level: permission.LevelAllowOnce, allowed: true, denied: false},
		{name: "ask", level: permission.LevelAsk, allowed: false, denied: false},
		{name: "deny", level: permission.LevelDeny, allowed: false, denied: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := permission.New("test", tt.level)
			if got := e.IsAllowed(); got != tt.allowed {
				t.Errorf("IsAllowed() = %v, want %v", got, tt.allowed)
			}
			if got := e.IsDenied(); got != tt.denied {
				t.Errorf("IsDenied() = %v, want %v", got, tt.denied)
			}
		})
	}
}
