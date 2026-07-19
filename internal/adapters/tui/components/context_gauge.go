package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type ContextGauge struct {
	width int
	used  int
	total int
	label string
}

func NewContextGauge() *ContextGauge {
	return &ContextGauge{
		width: 80,
		label: "Context",
	}
}

func (g *ContextGauge) SetWidth(w int) {
	g.width = w
}

func (g *ContextGauge) SetUsage(used, total int) {
	g.used = used
	g.total = total
}

func (g *ContextGauge) IsActive() bool {
	return g.total > 0
}

func (g *ContextGauge) SetLabel(label string) {
	g.label = label
}

func (g *ContextGauge) Percentage() float64 {
	if g.total <= 0 {
		return 0
	}
	pct := float64(g.used) / float64(g.total) * 100
	if pct > 100 {
		pct = 100
	}
	return pct
}

func (g *ContextGauge) View() string {
	if g.total <= 0 {
		return ""
	}

	pct := g.Percentage()

	barLen := g.width - 30
	if barLen < 10 {
		barLen = 10
	}

	filled := int(float64(barLen) * pct / 100)
	if filled > barLen {
		filled = barLen
	}
	empty := barLen - filled

	barColor := "39"
	switch {
	case pct >= 90:
		barColor = "196"
	case pct >= 70:
		barColor = "214"
	case pct >= 50:
		barColor = "220"
	}

	filledBar := lipgloss.NewStyle().
		Background(lipgloss.Color(barColor)).
		Render(strings.Repeat(" ", filled))

	emptyBar := lipgloss.NewStyle().
		Background(lipgloss.Color("236")).
		Render(strings.Repeat(" ", empty))

	bar := filledBar + emptyBar

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Width(8).
		Render(g.label)

	pctStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(barColor)).
		Bold(true).
		Width(5).
		Align(lipgloss.Right).
		Render(fmt.Sprintf("%.0f%%", pct))

	usageStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("250")).
		Render(fmt.Sprintf("%d/%d", g.used, g.total))

	return fmt.Sprintf("%s %s %s %s", labelStyle, bar, pctStyle, usageStyle)
}
