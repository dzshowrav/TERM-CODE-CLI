package components

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	emptyTitle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	emptyHint  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
)

type EmptyState struct {
	title string
	hint  string
	width int
}

func NewEmptyState(title, hint string) *EmptyState {
	return &EmptyState{
		title: title,
		hint:  hint,
		width: 80,
	}
}

func (e *EmptyState) SetWidth(w int) {
	e.width = w
}

func (e *EmptyState) SetTitle(title string) {
	e.title = title
}

func (e *EmptyState) SetHint(hint string) {
	e.hint = hint
}

func (e *EmptyState) Render() string {
	lines := []string{
		"",
		emptyTitle.Render(e.title),
		"",
		emptyHint.Render(e.hint),
		"",
	}
	return fmt.Sprintf("%s\n%s\n%s", lines[0], lines[1], lines[3])
}
