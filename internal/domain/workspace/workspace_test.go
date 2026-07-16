package workspace_test

import (
	"testing"

	"termcode/internal/domain/workspace"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		nameArg string
		path    string
	}{
		{name: "creates workspace", nameArg: "my-project", path: "/home/user/project"},
		{name: "creates with root path", nameArg: "root", path: "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := workspace.New(tt.nameArg, tt.path)
			if w.ID == "" {
				t.Error("expected non-empty ID")
			}
			if w.Name != tt.nameArg {
				t.Errorf("expected Name=%q", tt.nameArg)
			}
			if w.Path != tt.path {
				t.Errorf("expected Path=%q", tt.path)
			}
			if w.IsDefault {
				t.Error("expected IsDefault=false")
			}
		})
	}
}
