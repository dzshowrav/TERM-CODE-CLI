package styles

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	App          = lipgloss.NewStyle()
	Screen       = lipgloss.NewStyle()
	Separator    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	TitleStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	Subtitle     = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	LabelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	ValueStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	HintStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	Active       = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	Inactive     = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	contentStyle = lipgloss.NewStyle().Padding(0, 2)

	DialogBorder   = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	DialogSepStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
)

func Content(width int, body string) string {
	return contentStyle.Copy().Width(width).Render(body)
}

func SeparatorLine(width int) string {
	if width < 1 {
		return ""
	}
	return Separator.Render(strings.Repeat("─", width))
}

func DialogSep(width int) string {
	if width < 1 {
		return ""
	}
	return DialogSepStyle.Render(strings.Repeat("─", width))
}

func DialogBox(totalWidth int, body string) string {
	if totalWidth < 3 {
		return body
	}
	innerWidth := totalWidth - 2

	bH := DialogBorder.Render("─")
	bV := DialogBorder.Render("│")
	bTL := DialogBorder.Render("╭")
	bTR := DialogBorder.Render("╮")
	bBL := DialogBorder.Render("╰")
	bBR := DialogBorder.Render("╯")

	var b strings.Builder

	b.WriteString(bTL)
	b.WriteString(strings.Repeat(bH, innerWidth))
	b.WriteString(bTR)
	b.WriteByte('\n')

	for _, line := range strings.Split(body, "\n") {
		b.WriteString(bV)
		vw := lipgloss.Width(line)
		if vw < innerWidth {
			b.WriteString(line)
			b.WriteString(strings.Repeat(" ", innerWidth-vw))
		} else {
			b.WriteString(lipgloss.NewStyle().Width(innerWidth).Render(line))
		}
		b.WriteString(bV)
		b.WriteByte('\n')
	}

	b.WriteString(bBL)
	b.WriteString(strings.Repeat(bH, innerWidth))
	b.WriteString(bBR)

	return b.String()
}
