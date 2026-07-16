package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	progressFilled = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	progressEmpty  = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
	progressLabel  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

type ProgressBar struct {
	width   int
	current int
	total   int
	label   string
}

func NewProgressBar(width int) *ProgressBar {
	return &ProgressBar{
		width: width,
	}
}

func (p *ProgressBar) SetWidth(w int) {
	p.width = w
}

func (p *ProgressBar) SetProgress(current, total int) {
	p.current = current
	p.total = total
}

func (p *ProgressBar) SetLabel(label string) {
	p.label = label
}

func (p *ProgressBar) Percent() float64 {
	if p.total <= 0 {
		return 0
	}
	pct := float64(p.current) / float64(p.total)
	if pct > 1.0 {
		pct = 1.0
	}
	return pct
}

func (p *ProgressBar) View() string {
	labelW := lipgloss.Width(p.label)
	barWidth := p.width - labelW - 10
	if barWidth < 5 {
		barWidth = 5
	}

	pct := p.Percent()
	filled := int(pct * float64(barWidth))
	empty := barWidth - filled

	bar := progressFilled.Render(repeat("■", filled)) + progressEmpty.Render(repeat("■", empty))

	return progressLabel.Render(fmt.Sprintf("%s %s %3.0f%%", p.label, bar, pct*100))
}
