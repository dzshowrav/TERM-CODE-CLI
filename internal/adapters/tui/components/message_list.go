package components

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type MessageItem struct {
	Role    string
	Content string
}

var (
	userLabel      = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).SetString(">")
	assistantLabel = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).SetString("AI")
	systemLabel    = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).SetString("*")
	toolLabel      = lipgloss.NewStyle().Foreground(lipgloss.Color("221")).SetString("tool")
)

type MessageList struct {
	viewport *Viewport
	messages []MessageItem
	width    int
	height   int
}

func NewMessageList(width, height int) *MessageList {
	return &MessageList{
		viewport: NewViewport(width, height),
		width:    width,
		height:   height,
	}
}

func (l *MessageList) SetSize(w, h int) {
	l.width = w
	l.height = h
	l.viewport.SetSize(w, h)
}

func (l *MessageList) Add(msg MessageItem) {
	l.messages = append(l.messages, msg)
	l.render()
}

func (l *MessageList) render() {
	var lines []string
	for _, msg := range l.messages {
		roleStr := formatRole(msg.Role)
		for _, line := range wrapLines(msg.Content, l.width-2) {
			lines = append(lines, roleStr+" "+line)
		}
		lines = append(lines, "")
	}
	l.viewport.SetContent(lines)
}

func (l *MessageList) View() string {
	return l.viewport.View()
}

func (l *MessageList) Update(msg tea.Msg) {
	l.viewport.Update(msg)
}

func formatRole(role string) string {
	switch role {
	case "user":
		return userLabel.String()
	case "assistant":
		return assistantLabel.String()
	case "tool":
		return toolLabel.String()
	default:
		return systemLabel.String()
	}
}

func wrapLines(s string, max int) []string {
	if max <= 0 {
		return []string{s}
	}
	var lines []string
	runes := []rune(s)
	for len(runes) > max {
		lines = append(lines, string(runes[:max]))
		runes = runes[max:]
	}
	if len(runes) > 0 {
		lines = append(lines, string(runes))
	}
	if len(lines) == 0 {
		lines = []string{""}
	}
	return lines
}
