package screens

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	sessNameStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	sessInfoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	sessActiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
)

type SessionListItem struct {
	ID        string
	Name      string
	MsgCount  int
	TokenIn   int
	TokenOut  int
	IsActive  bool
	UpdatedAt time.Time
}

type SessionScreen struct {
	width    int
	height   int
	sessions []SessionListItem
}

func NewSessionScreen() *SessionScreen {
	return &SessionScreen{
		width:  80,
		height: 24,
	}
}

func (s *SessionScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *SessionScreen) SetSessions(sessions []SessionListItem) {
	s.sessions = sessions
}

func (s *SessionScreen) View() string {
	header := styles.H1.Render("Session Manager")
	sep := styles.SeparatorLine(s.width)

	if len(s.sessions) == 0 {
		empty := lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("No sessions.")
		return fmt.Sprintf("%s\n%s\n%s", header, sep, empty)
	}

	var lines []string
	for _, ses := range s.sessions {
		active := ""
		if ses.IsActive {
			active = sessActiveStyle.Render(" [active]")
		}
		lines = append(lines, fmt.Sprintf(" %s%s", sessNameStyle.Render(ses.Name), active))
		lines = append(lines, fmt.Sprintf("   %s", sessInfoStyle.Render(
			fmt.Sprintf(
				"%d msgs · %d in / %d out · %s",
				ses.MsgCount, ses.TokenIn, ses.TokenOut,
				ses.UpdatedAt.Format("2006-01-02 15:04"),
			),
		)))
		lines = append(lines, "")
	}

	return fmt.Sprintf("%s\n%s\n%s", header, sep, strings.Join(lines, "\n"))
}
