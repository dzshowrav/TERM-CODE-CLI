package dialogs

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var (
	notifyInfo = lipgloss.NewStyle().Background(lipgloss.Color("39")).Foreground(lipgloss.Color("255")).Padding(0, 1).Render("i")
	notifyOK   = lipgloss.NewStyle().Background(lipgloss.Color("83")).Foreground(lipgloss.Color("235")).Padding(0, 1).Render("✓")
	notifyErr  = lipgloss.NewStyle().Background(lipgloss.Color("196")).Foreground(lipgloss.Color("255")).Padding(0, 1).Render("!")
)

type NotificationLevel int

const (
	LevelInfo NotificationLevel = iota
	LevelSuccess
	LevelError
)

type Notification struct {
	message string
	level   NotificationLevel
	endTime time.Time
	width   int
}

func NewNotification(message string, level NotificationLevel, duration time.Duration) *Notification {
	return &Notification{
		message: message,
		level:   level,
		endTime: time.Now().Add(duration),
		width:   80,
	}
}

func (n *Notification) SetWidth(w int) {
	n.width = w
}

func (n *Notification) Expired() bool {
	return time.Now().After(n.endTime)
}

func (n *Notification) Render() string {
	icon := notifyInfo
	switch n.level {
	case LevelSuccess:
		icon = notifyOK
	case LevelError:
		icon = notifyErr
	}

	msg := n.message
	maxW := n.width - 6
	if len(msg) > maxW {
		msg = msg[:maxW]
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("255")).
		Background(lipgloss.Color("236")).
		Padding(0, 2).
		Width(n.width)
	return style.Render(fmt.Sprintf("%s %s", icon, msg))
}

func (n *Notification) View() string {
	return n.Render()
}
