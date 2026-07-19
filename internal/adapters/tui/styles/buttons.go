package styles

import "github.com/charmbracelet/lipgloss"

var (
	Warning = lipgloss.NewStyle().
		Foreground(lipgloss.Color("208")).
		Bold(true)

	Accent = lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Bold(true)

	Help = lipgloss.NewStyle().
		Foreground(lipgloss.Color("240")).
		Italic(true)

	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("39"))

	FormBtnNormal = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))

	FormBtnActive = lipgloss.NewStyle().
			Foreground(lipgloss.Color("39")).
			Bold(true)

	DiffAdded = lipgloss.NewStyle().
			Foreground(lipgloss.Color("83"))

	DiffRemoved = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196"))
)
