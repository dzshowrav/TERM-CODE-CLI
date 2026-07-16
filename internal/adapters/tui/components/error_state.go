package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	errorIcon  = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("!")
	errorTitle = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	errorMsg   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	errorHint  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
)

type ErrorState struct {
	message string
	hint    string
	width   int
}

func NewErrorState(message, hint string) *ErrorState {
	return &ErrorState{
		message: message,
		hint:    hint,
		width:   80,
	}
}

func (e *ErrorState) SetWidth(w int) {
	e.width = w
}

func (e *ErrorState) SetMessage(msg string) {
	e.message = msg
}

func (e *ErrorState) Render() string {
	return fmt.Sprintf(
		"%s %s\n%s\n%s",
		errorIcon,
		errorTitle.Render("Error"),
		errorMsg.Render(e.message),
		errorHint.Render(e.hint),
	)
}
