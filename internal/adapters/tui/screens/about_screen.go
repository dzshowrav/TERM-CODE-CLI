package screens

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type AboutScreen struct {
	width  int
	height int
	done   bool
}

func NewAboutScreen() *AboutScreen {
	return &AboutScreen{}
}

func (s *AboutScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *AboutScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "enter", "q":
			s.done = true
		}
	}
	return s, nil
}

func (s *AboutScreen) View() string {
	if s.width < 10 || s.height < 10 {
		return "About TermCode"
	}

	body := s.renderBody()
	return styles.DialogBox(s.width-4, body)
}

func (s *AboutScreen) renderBody() string {
	infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("255"))

	lines := []string{
		infoStyle.Render("TermCode"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("AI-Powered Terminal Coding Assistant"),
		"",
		labelStyle.Render("Version:    ") + valueStyle.Render("v0.1.0"),
		labelStyle.Render("Runtime:    ") + valueStyle.Render("Go 1.26"),
		labelStyle.Render("Framework:  ") + valueStyle.Render("Bubble Tea v2"),
		labelStyle.Render("License:    ") + valueStyle.Render("MIT"),
		labelStyle.Render("Storage:    ") + valueStyle.Render("SQLite"),
		"",
		labelStyle.Render("Architecture:"),
		valueStyle.Render("  Clean Architecture + DDD"),
		valueStyle.Render("  Event-Driven with EventBus"),
		valueStyle.Render("  MCP-native Tool Integration"),
		"",
		labelStyle.Render("Key Features:"),
		valueStyle.Render("  Multi-provider LLM support"),
		valueStyle.Render("  MCP server management"),
		valueStyle.Render("  Git integration"),
		valueStyle.Render("  Tool execution with safety"),
		valueStyle.Render("  Session management"),
		"",
		labelStyle.Render("TermCode is designed for Termux"),
		labelStyle.Render("with mobile-first, offline-first principles."),
		"",
		styles.HintStyle.Render("Press ESC or q to close"),
	}

	innerW := s.width - 8
	if innerW < 20 {
		innerW = 20
	}

	var b strings.Builder
	for _, line := range lines {
		vw := lipgloss.Width(line)
		if vw < innerW {
			b.WriteString(line)
			b.WriteString(strings.Repeat(" ", innerW-vw))
		} else {
			b.WriteString(lipgloss.NewStyle().Width(innerW).Render(line))
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (s *AboutScreen) Done() bool         { return s.done }
func (s *AboutScreen) Result() string     { return "" }
func (s *AboutScreen) Title() string      { return "About TermCode" }
func (s *AboutScreen) IsFullScreen() bool { return false }
