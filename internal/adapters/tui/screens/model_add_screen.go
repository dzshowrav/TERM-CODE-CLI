package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type ModelAddScreen struct {
	width       int
	height      int
	modelID     string
	displayName string
	provider    string
	contextSize int
	maxOutput   int
	focusField  int
}

func NewModelAddScreen() *ModelAddScreen {
	return &ModelAddScreen{
		width:       80,
		height:      24,
		contextSize: 4096,
		maxOutput:   4096,
	}
}

func (s *ModelAddScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ModelAddScreen) View() string {
	header := styles.H1.Render("Add Model")
	sep := styles.SeparatorLine(s.width)

	fields := []struct {
		label string
		value string
	}{
		{"Model ID", s.modelID},
		{"Display Name", s.displayName},
		{"Provider", s.provider},
		{"Context Size", fmt.Sprintf("%d", s.contextSize)},
		{"Max Output", fmt.Sprintf("%d", s.maxOutput)},
	}

	var lines []string
	for i, f := range fields {
		mark := " "
		if i == s.focusField {
			mark = ">"
		}
		lines = append(lines, fmt.Sprintf(" %s %s", mark, formLabelStyle.Render(f.label)))
		lines = append(lines, fmt.Sprintf("   %s", formInputStyle.Render(f.value)))
		lines = append(lines, "")
	}

	submit := lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("[ Save ]  [ Cancel ]")

	return fmt.Sprintf("%s\n%s\n%s\n%s", header, sep, strings.Join(lines, "\n"), submit)
}

var _ = strings.Builder{}
