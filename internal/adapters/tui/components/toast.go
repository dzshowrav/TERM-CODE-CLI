package components

import (
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type ToastLevel int

const (
	ToastInfo ToastLevel = iota
	ToastSuccess
	ToastWarning
	ToastError
)

type Toast struct {
	ID        string
	Level     ToastLevel
	Message   string
	CreatedAt time.Time
	Duration  time.Duration
}

type ToastManager struct {
	mu        sync.RWMutex
	toasts    []Toast
	width     int
	maxAge    time.Duration
	maxToasts int
}

func NewToastManager() *ToastManager {
	return &ToastManager{
		maxAge:    5 * time.Second,
		maxToasts: 5,
	}
}

func (m *ToastManager) SetSize(w int) {
	m.mu.Lock()
	m.width = w
	m.mu.Unlock()
}

func (m *ToastManager) Show(level ToastLevel, msg string) {
	m.mu.Lock()
	id := time.Now().Format("150405.000000")
	dur := m.maxAge
	switch level {
	case ToastError:
		dur = 8 * time.Second
	case ToastSuccess:
		dur = 4 * time.Second
	}
	m.toasts = append(m.toasts, Toast{
		ID:        id,
		Level:     level,
		Message:   msg,
		CreatedAt: time.Now(),
		Duration:  dur,
	})
	if len(m.toasts) > m.maxToasts {
		m.toasts = m.toasts[len(m.toasts)-m.maxToasts:]
	}
	m.mu.Unlock()
}

func (m *ToastManager) HasToast() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.toasts) > 0
}

func (m *ToastManager) Info(msg string)    { m.Show(ToastInfo, msg) }
func (m *ToastManager) Success(msg string) { m.Show(ToastSuccess, msg) }
func (m *ToastManager) Warning(msg string) { m.Show(ToastWarning, msg) }
func (m *ToastManager) Error(msg string)   { m.Show(ToastError, msg) }

func (m *ToastManager) View() string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	now := time.Now()
	lines := make([]string, 0, len(m.toasts))
	for _, t := range m.toasts {
		if now.Sub(t.CreatedAt) > t.Duration {
			continue
		}
		elapsed := now.Sub(t.CreatedAt)
		alpha := 1.0
		fadeStart := t.Duration * 8 / 10
		if elapsed > fadeStart {
			remaining := t.Duration - elapsed
			alpha = float64(remaining) / float64(t.Duration-fadeStart)
			if alpha < 0 {
				alpha = 0
			}
		}

		var icon, bgColor, fgColor string
		switch t.Level {
		case ToastInfo:
			icon, bgColor, fgColor = " i ", "39", "15"
		case ToastSuccess:
			icon, bgColor, fgColor = " ✓ ", "42", "15"
		case ToastWarning:
			icon, bgColor, fgColor = " ! ", "43", "0"
		case ToastError:
			icon, bgColor, fgColor = " ✗ ", "41", "15"
		}

		dimStyle := lipgloss.NewStyle()
		if alpha < 1 {
			dimStyle = dimStyle.Faint(true)
		}
		remainWidth := m.width - 2
		if remainWidth < 10 {
			remainWidth = 10
		}
		shortMsg := t.Message
		if len(shortMsg) > remainWidth-4 {
			shortMsg = shortMsg[:remainWidth-5] + "…"
		}

		badge := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(fgColor)).
			Background(lipgloss.Color(bgColor)).
			Render(icon)

		text := lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			Width(remainWidth - 2).
			Render(shortMsg)

		line := dimStyle.Render(badge + " " + text)
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
