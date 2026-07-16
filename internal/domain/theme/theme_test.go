package theme_test

import (
	"testing"

	"termcode/internal/domain/theme"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name      string
		themeName string
		palette   string
		isDark    bool
	}{
		{name: "creates dark theme", themeName: "dracula", palette: "#282a36", isDark: true},
		{name: "creates light theme", themeName: "solarized", palette: "#fdf6e3", isDark: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			th := theme.New(tt.themeName, tt.palette, tt.isDark)
			if th.ID == "" {
				t.Error("expected non-empty ID")
			}
			if th.Name != tt.themeName {
				t.Errorf("expected Name=%q", tt.themeName)
			}
			if th.Palette != tt.palette {
				t.Errorf("expected Palette=%q", tt.palette)
			}
			if th.IsDark != tt.isDark {
				t.Errorf("expected IsDark=%v", tt.isDark)
			}
		})
	}
}

func TestNew_Defaults(t *testing.T) {
	th := theme.New("test", "#000", false)
	if th.Version != "1.0" {
		t.Errorf("expected Version=1.0, got %q", th.Version)
	}
	if th.IsActive {
		t.Error("expected IsActive=false")
	}
}
