package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	gaugeEmpty  = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
	gaugeFilled = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	gaugeWarn   = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
	gaugeFull   = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	gaugeLabel  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

type TokenGauge struct {
	width   int
	current int
	max     int
}

func NewTokenGauge(width int) *TokenGauge {
	return &TokenGauge{
		width: width,
	}
}

func (g *TokenGauge) SetWidth(w int) {
	g.width = w
}

func (g *TokenGauge) SetUsage(current, max int) {
	g.current = current
	g.max = max
}

func (g *TokenGauge) Percent() float64 {
	if g.max <= 0 {
		return 0
	}
	pct := float64(g.current) / float64(g.max)
	if pct > 1.0 {
		pct = 1.0
	}
	return pct
}

func (g *TokenGauge) View() string {
	pct := g.Percent()

	gaugeWidth := g.width - 12
	if gaugeWidth < 5 {
		gaugeWidth = 5
	}

	filled := int(pct * float64(gaugeWidth))
	if filled > gaugeWidth {
		filled = gaugeWidth
	}
	empty := gaugeWidth - filled

	fillStyle := gaugeFilled
	if pct > 0.85 {
		fillStyle = gaugeFull
	} else if pct > 0.7 {
		fillStyle = gaugeWarn
	}

	bar := fillStyle.Render(repeat("■", filled)) + gaugeEmpty.Render(repeat("■", empty))

	return gaugeLabel.Render(fmt.Sprintf("ctx %s %3.0f%%", bar, pct*100))
}

func repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	result := make([]byte, n*len(s))
	for i := 0; i < n; i++ {
		copy(result[i*len(s):], s)
	}
	return string(result)
}
