package components

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type thinkingTickMsg time.Time

var thinkingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Italic(true)

type ThinkingIndicator struct {
	frames []string
	frame  int
	active bool
}

func NewThinkingIndicator() *ThinkingIndicator {
	return &ThinkingIndicator{
		frames: []string{"◐", "◓", "◑", "◒"},
		frame:  0,
		active: false,
	}
}

func (t *ThinkingIndicator) Start() {
	t.active = true
	t.frame = 0
}

func (t *ThinkingIndicator) Stop() {
	t.active = false
}

func (t *ThinkingIndicator) IsActive() bool {
	return t.active
}

func (t *ThinkingIndicator) Init() tea.Cmd {
	return nil
}

func (t *ThinkingIndicator) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case thinkingTickMsg:
		if !t.active {
			return nil
		}
		t.frame = (t.frame + 1) % len(t.frames)
		return tea.Tick(200*time.Millisecond, func(ti time.Time) tea.Msg {
			return thinkingTickMsg(ti)
		})
	}
	return nil
}

func (t *ThinkingIndicator) View() string {
	if !t.active {
		return ""
	}
	return thinkingStyle.Render(t.frames[t.frame] + " thinking...")
}
