package components

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	userContentStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	assistantStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	systemContentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Italic(true)
	codeStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Background(lipgloss.Color("235"))
)

type MessageItemRenderer struct {
	width int
}

func NewMessageItemRenderer(width int) *MessageItemRenderer {
	return &MessageItemRenderer{width: width}
}

func (r *MessageItemRenderer) Render(msg MessageItem) string {
	var style lipgloss.Style
	switch msg.Role {
	case "user":
		style = userContentStyle
	case "assistant":
		style = assistantStyle
	default:
		style = systemContentStyle
	}

	roleStr := formatRole(msg.Role)
	return roleStr + " " + style.Render(msg.Content)
}
