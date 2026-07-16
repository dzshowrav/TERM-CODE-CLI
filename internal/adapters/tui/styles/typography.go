package styles

import "github.com/charmbracelet/lipgloss"

var (
	H1 = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("255"))

	H2 = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("39"))

	Body = lipgloss.NewStyle().
		Foreground(lipgloss.Color("250"))

	Dim = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240"))

	Err = lipgloss.NewStyle().
		Foreground(lipgloss.Color("196"))

	Success = lipgloss.NewStyle().
		Foreground(lipgloss.Color("83"))
)
