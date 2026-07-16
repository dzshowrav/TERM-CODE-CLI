package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	formFieldStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	formLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	formInputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	formHintStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
)

type ProviderAddScreen struct {
	width      int
	height     int
	name       string
	baseURL    string
	apiKey     string
	desc       string
	focusField int
}

func NewProviderAddScreen() *ProviderAddScreen {
	return &ProviderAddScreen{
		width:      80,
		height:     24,
		focusField: 0,
	}
}

func (s *ProviderAddScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ProviderAddScreen) View() string {
	header := styles.H1.Render("Add Provider")
	sep := styles.SeparatorLine(s.width)

	fields := []struct {
		label string
		value string
		hint  string
	}{
		{"Provider Name", s.name, "e.g., My OpenAI"},
		{"Base URL", s.baseURL, "e.g., https://api.openai.com/v1"},
		{"API Key", maskKey(s.apiKey), "Will be encrypted at rest"},
		{"Description (optional)", s.desc, "e.g., My primary provider"},
	}

	var lines []string
	for i, f := range fields {
		mark := " "
		if i == s.focusField {
			mark = ">"
		}
		val := f.value
		if val == "" {
			val = formHintStyle.Render(f.hint)
		}
		lines = append(lines, fmt.Sprintf(" %s %s", mark, formLabelStyle.Render(f.label)))
		lines = append(lines, fmt.Sprintf("   %s", formInputStyle.Render(val)))
		lines = append(lines, "")
	}

	submit := lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("[ Save & Test ]  [ Cancel ]")

	return fmt.Sprintf("%s\n%s\n%s\n%s", header, sep, strings.Join(lines, "\n"), submit)
}

func maskKey(key string) string {
	if len(key) <= 8 {
		return strings.Repeat("*", len(key))
	}
	return key[:4] + strings.Repeat("*", len(key)-8) + key[len(key)-4:]
}
