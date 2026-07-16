package components

import "github.com/charmbracelet/lipgloss"

type EmptyState struct {
	message string
	width   int
}

func NewEmptyState() *EmptyState {
	return &EmptyState{
		message: "No items to display.",
		width:   80,
	}
}

func (e *EmptyState) SetWidth(w int) {
	e.width = w
}

func (e *EmptyState) SetMessage(msg string) {
	e.message = msg
}

func (e *EmptyState) Render() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true).
		Width(e.width).
		Align(lipgloss.Center).
		Padding(1, 2)
	return style.Render(e.message)
}
