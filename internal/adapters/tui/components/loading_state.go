package components

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type loadingTickMsg time.Time

var loadingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Italic(true)

type LoadingState struct {
	message string
	frames  []string
	frame   int
	active  bool
}

func NewLoadingState(message string) *LoadingState {
	return &LoadingState{
		message: message,
		frames:  []string{"◐", "◓", "◑", "◒"},
		active:  true,
	}
}

func (l *LoadingState) SetMessage(msg string) {
	l.message = msg
}

func (l *LoadingState) Start() {
	l.active = true
	l.frame = 0
}

func (l *LoadingState) Stop() {
	l.active = false
}

func (l *LoadingState) IsActive() bool {
	return l.active
}

func (l *LoadingState) Init() tea.Cmd {
	return l.tick()
}

func (l *LoadingState) tick() tea.Cmd {
	if !l.active {
		return nil
	}
	return tea.Tick(200*time.Millisecond, func(t time.Time) tea.Msg {
		return loadingTickMsg(t)
	})
}

func (l *LoadingState) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case loadingTickMsg:
		if !l.active {
			return nil
		}
		l.frame = (l.frame + 1) % len(l.frames)
		return l.tick()
	}
	return nil
}

func (l *LoadingState) View() string {
	if !l.active {
		return ""
	}
	spinner := l.frames[l.frame]
	return loadingStyle.Render(fmt.Sprintf("%s %s", spinner, l.message))
}
