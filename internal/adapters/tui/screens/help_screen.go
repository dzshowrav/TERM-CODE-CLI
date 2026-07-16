package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	helpSec  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Bold(true)
	helpCmd  = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
	helpDesc = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

type HelpScreen struct {
	width  int
	height int
}

func NewHelpScreen() *HelpScreen {
	return &HelpScreen{
		width:  80,
		height: 24,
	}
}

func (s *HelpScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *HelpScreen) View() string {
	header := styles.H1.Render("Help - Slash Commands")
	sep := styles.SeparatorLine(s.width)

	sections := []struct {
		name    string
		entries []struct{ cmd, desc string }
	}{
		{"Conversation", []struct{ cmd, desc string }{
			{"/clear", "Clear current conversation"},
			{"/home", "Return to home screen"},
			{"/retry", "Retry last AI response"},
			{"/cancel", "Cancel current operation"},
		}},
		{"Provider & Model", []struct{ cmd, desc string }{
			{"/provider", "View/switch provider"},
			{"/provider list", "List all providers"},
			{"/provider add", "Add a new provider"},
			{"/model", "Open model selector"},
			{"/model list", "List all models"},
			{"/all models", "List all models grouped"},
		}},
		{"Session", []struct{ cmd, desc string }{
			{"/session", "Session management"},
			{"/sessions", "List saved sessions"},
		}},
		{"Workspace & Git", []struct{ cmd, desc string }{
			{"/workspace", "Show/set workspace"},
			{"/search", "Search workspace"},
			{"/git", "Git operations"},
		}},
		{"Settings", []struct{ cmd, desc string }{
			{"/settings", "Open settings"},
			{"/theme", "Change theme"},
		}},
		{"System", []struct{ cmd, desc string }{
			{"/help", "Show this help"},
			{"/quit", "Quit application"},
		}},
	}

	var lines []string
	for _, sec := range sections {
		lines = append(lines, helpSec.Render("── "+sec.name+" ──"))
		for _, e := range sec.entries {
			padding := 18 - len(e.cmd)
			if padding < 1 {
				padding = 1
			}
			lines = append(lines, fmt.Sprintf(
				"  %s%s%s",
				helpCmd.Render(e.cmd),
				strings.Repeat(" ", padding),
				helpDesc.Render(e.desc),
			))
		}
		lines = append(lines, "")
	}

	return fmt.Sprintf("%s\n%s\n%s", header, sep, strings.Join(lines, "\n"))
}
