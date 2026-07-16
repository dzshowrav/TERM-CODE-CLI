package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	toolExIcon    = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("◐")
	toolExDone    = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Render("✓")
	toolExFail    = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("✗")
	toolExName    = lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Bold(true)
	toolExContent = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	toolExTime    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

type ToolExecutionItem struct {
	Name     string
	Input    string
	Output   string
	Error    string
	Status   string // running, success, failed
	Duration int64  // ms
}

type ToolExecutionScreen struct {
	width  int
	height int
	tools  []ToolExecutionItem
}

func NewToolExecutionScreen() *ToolExecutionScreen {
	return &ToolExecutionScreen{
		width:  80,
		height: 24,
	}
}

func (s *ToolExecutionScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ToolExecutionScreen) AddExecution(item ToolExecutionItem) {
	s.tools = append(s.tools, item)
}

func (s *ToolExecutionScreen) Clear() {
	s.tools = nil
}

func (s *ToolExecutionScreen) View() string {
	header := styles.H1.Render("Tool Execution")
	sep := styles.SeparatorLine(s.width)

	if len(s.tools) == 0 {
		empty := lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("No tools executed.")
		return fmt.Sprintf("%s\n%s\n%s", header, sep, empty)
	}

	var lines []string
	for _, t := range s.tools {
		icon := toolExIcon
		switch t.Status {
		case "success", "completed":
			icon = toolExDone
		case "failed", "error":
			icon = toolExFail
		}

		duration := ""
		if t.Duration > 0 {
			duration = fmt.Sprintf(" (%dms)", t.Duration)
		}

		lines = append(lines, fmt.Sprintf(" %s %s%s", icon, toolExName.Render(t.Name), toolExTime.Render(duration)))

		if t.Input != "" {
			lines = append(lines, fmt.Sprintf("   %s", toolExContent.Render(t.Input)))
		}
		if t.Output != "" {
			maxW := s.width - 6
			if maxW < 20 {
				maxW = 20
			}
			for _, line := range strings.Split(t.Output, "\n") {
				runes := []rune(line)
				for len(runes) > maxW {
					lines = append(lines, fmt.Sprintf("   %s", toolExContent.Render(string(runes[:maxW]))))
					runes = runes[maxW:]
				}
				if len(runes) > 0 {
					lines = append(lines, fmt.Sprintf("   %s", toolExContent.Render(string(runes))))
				}
			}
		}
		if t.Error != "" {
			maxW := s.width - 6
			if maxW < 20 {
				maxW = 20
			}
			errRunes := []rune(t.Error)
			for len(errRunes) > maxW {
				lines = append(lines, fmt.Sprintf("   %s", lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(string(errRunes[:maxW]))))
				errRunes = errRunes[maxW:]
			}
			if len(errRunes) > 0 {
				lines = append(lines, fmt.Sprintf("   %s", lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render(string(errRunes))))
			}
		}
		lines = append(lines, "")
	}

	return fmt.Sprintf("%s\n%s\n%s", header, sep, strings.Join(lines, "\n"))
}
