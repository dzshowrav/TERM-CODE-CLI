package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type HelpEntry struct {
	Command     string
	Description string
}

type HelpText struct {
	entries []HelpEntry
	width   int
}

func NewHelpText() *HelpText {
	return &HelpText{
		width: 80,
	}
}

func (h *HelpText) SetWidth(w int) {
	h.width = w
}

func (h *HelpText) SetEntries(entries []HelpEntry) {
	h.entries = entries
}

func (h *HelpText) Render() string {
	var lines []string
	for _, entry := range h.entries {
		cmdW := h.width / 3
		if cmdW < 16 {
			cmdW = 16
		}
		if cmdW > 24 {
			cmdW = 24
		}
		padding := cmdW - len(entry.Command)
		if padding < 1 {
			padding = 1
		}
		line := fmt.Sprintf("  %s%s%s",
			lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39")).Render("/"+entry.Command),
			strings.Repeat(" ", padding),
			lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Render(entry.Description))
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
