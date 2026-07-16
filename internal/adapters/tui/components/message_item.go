package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	roleUser      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39")).Render("You")
	roleAssistant = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("83")).Render("AI")
	roleSystem    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("240")).Render("Sys")
	roleTool      = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214")).Render("Tool")
)

type MessageItemRenderer struct {
	width int
}

func NewMessageItemRenderer(width int) *MessageItemRenderer {
	return &MessageItemRenderer{width: width}
}

func (r *MessageItemRenderer) SetWidth(w int) {
	r.width = w
}

func (r *MessageItemRenderer) Render(role, content string) string {
	var label string
	switch role {
	case "user":
		label = roleUser
	case "assistant":
		label = roleAssistant
	case "system":
		label = roleSystem
	case "tool":
		label = roleTool
	default:
		label = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("245")).Render(role)
	}

	available := r.width - 4
	if available < 20 {
		available = 20
	}
	contentLines := wrapLines(content, available)

	var buf string
	for _, l := range contentLines {
		if buf != "" {
			buf += "\n"
		}
		buf += l
	}

	return fmt.Sprintf("%s\n%s", label, buf)
}
