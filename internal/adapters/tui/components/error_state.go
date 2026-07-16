package components

import "github.com/charmbracelet/lipgloss"

type ErrorState struct {
	message string
	width   int
}

func NewErrorState() *ErrorState {
	return &ErrorState{
		width: 80,
	}
}

func (e *ErrorState) SetWidth(w int) {
	e.width = w
}

func (e *ErrorState) SetMessage(msg string) {
	e.message = msg
}

func (e *ErrorState) Render() string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("196")).
		Bold(true).
		Width(e.width).
		Align(lipgloss.Center).
		Padding(1, 2)
	return style.Render("Error: " + e.message)
}
