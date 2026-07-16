package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	helpTitle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	helpCmd   = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
	helpDesc  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	helpGroup = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
)

type HelpEntry struct {
	Command     string
	Description string
}

type HelpText struct {
	title  string
	groups []HelpGroup
	width  int
}

type HelpGroup struct {
	Name    string
	Entries []HelpEntry
}

func NewHelpText(title string) *HelpText {
	return &HelpText{
		title: title,
		width: 80,
	}
}

func (h *HelpText) AddGroup(name string, entries []HelpEntry) {
	h.groups = append(h.groups, HelpGroup{Name: name, Entries: entries})
}

func (h *HelpText) SetWidth(w int) {
	h.width = w
}

func (h *HelpText) Render() string {
	var parts []string
	parts = append(parts, helpTitle.Render(h.title))
	parts = append(parts, "")

	for _, group := range h.groups {
		parts = append(parts, helpGroup.Render("── "+group.Name+" ──"))
		for _, entry := range group.Entries {
			padding := 20 - len(entry.Command)
			if padding < 1 {
				padding = 1
			}
			cmd := helpCmd.Render(entry.Command)
			desc := helpDesc.Render(entry.Description)
			parts = append(parts, "  "+cmd+strings.Repeat(" ", padding)+desc)
		}
		parts = append(parts, "")
	}

	return strings.Join(parts, "\n")
}
