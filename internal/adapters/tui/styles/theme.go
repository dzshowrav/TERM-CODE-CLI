package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	App        = lipgloss.NewStyle()
	Screen     = lipgloss.NewStyle()
	Separator  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	TitleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	Subtitle   = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	LabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	ValueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	HintStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	Active     = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	Inactive   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

func SeparatorLine(width int) string {
	return Separator.Copy().Width(width).Render("─")
}
