package screens

import (
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	sessNameStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	sessInfoStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	sessActiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	sessPinnedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
)

type SessionListItem struct {
	ID        string
	Name      string
	MsgCount  int
	TokenIn   int
	TokenOut  int
	IsActive  bool
	IsPinned  bool
	UpdatedAt time.Time
}

type SessionScreen struct {
	width    int
	height   int
	sessions []SessionListItem
	cursor   int
	scroll   int
	done     bool
	result   string
	onDelete func(id string) string
	showAll  bool
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
	if s.cursor >= len(s.sessions) {
		s.cursor = len(s.sessions) - 1
	}
}

func (s *SessionScreen) OnDelete(fn func(id string) string) {
	s.onDelete = fn
}

func (s *SessionScreen) Done() bool     { return s.done }
func (s *SessionScreen) Result() string { return s.result }

func (s *SessionScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "enter":
			if len(s.sessions) > 0 && s.cursor >= 0 && s.cursor < len(s.sessions) {
				s.result = s.sessions[s.cursor].ID
				s.done = true
			}
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
				if s.cursor < s.scroll {
					s.scroll = s.cursor
				}
			}
		case "down", "j":
			if s.cursor < len(s.sessions)-1 {
				s.cursor++
				maxVis := s.height - 10
				if maxVis < 1 {
					maxVis = 1
				}
				if s.cursor >= s.scroll+maxVis {
					s.scroll = s.cursor - maxVis + 1
				}
			}
		case "d", "D":
			if len(s.sessions) > 0 && s.cursor >= 0 && s.cursor < len(s.sessions) {
				ses := s.sessions[s.cursor]
				if s.onDelete != nil {
					s.result = s.onDelete(ses.ID)
				} else {
					s.result = fmt.Sprintf("Session '%s' deleted.", ses.ID)
				}
				s.sessions = append(s.sessions[:s.cursor], s.sessions[s.cursor+1:]...)
				if s.cursor >= len(s.sessions) {
					s.cursor = len(s.sessions) - 1
				}
				if s.cursor < 0 {
					s.cursor = 0
				}
				if len(s.sessions) == 0 {
					s.done = true
				}
			}
		case "r", "R":
			if len(s.sessions) > 0 && s.cursor >= 0 && s.cursor < len(s.sessions) {
				ses := s.sessions[s.cursor]
				s.result = "__rename__:" + ses.ID
				s.done = true
			}
		case "p", "P":
			if len(s.sessions) > 0 && s.cursor >= 0 && s.cursor < len(s.sessions) {
				ses := s.sessions[s.cursor]
				if ses.IsPinned {
					s.result = "__unpin__:" + ses.ID
				} else {
					s.result = "__pin__:" + ses.ID
				}
				s.done = true
			}
		case "e", "E":
			if len(s.sessions) > 0 && s.cursor >= 0 && s.cursor < len(s.sessions) {
				ses := s.sessions[s.cursor]
				s.result = "__export__:" + ses.ID
				s.done = true
			}
		}
	}
	return s, nil
}

func (s *SessionScreen) View() string {
	innerW := s.width - 2

	title := fmt.Sprintf("%-*s%s", innerW-4, "Session Manager", "esc")

	if len(s.sessions) == 0 {
		body := fmt.Sprintf("%s\n%s\n%s",
			title,
			styles.DialogSep(innerW),
			styles.HintStyle.Render("No sessions."),
		)
		return styles.DialogBox(s.width, body)
	}

	needsTop := s.scroll > 0
	needsBottom := false

	maxVis := s.height - 10
	if maxVis < 1 {
		maxVis = 1
	}
	arrowCount := 0
	if needsTop {
		arrowCount++
	}
	if s.scroll+maxVis-arrowCount < len(s.sessions) {
		needsBottom = true
		arrowCount++
		if arrowCount > 2 {
			arrowCount = 2
		}
	}
	maxVis -= arrowCount
	if maxVis < 1 {
		maxVis = 1
	}

	if s.scroll+maxVis > len(s.sessions) {
		s.scroll = max(0, len(s.sessions)-maxVis)
	}
	end := min(s.scroll+maxVis, len(s.sessions))

	var bodyLines []string
	bodyLines = append(bodyLines, title)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, "")

	for i := s.scroll; i < end; i++ {
		ses := s.sessions[i]
		active := ""
		if ses.IsActive {
			active = sessActiveStyle.Render(" [active]")
		}
		pinned := ""
		if ses.IsPinned {
			pinned = sessPinnedStyle.Render(" [P]")
		}

		cursor := "  "
		nameStyle := sessNameStyle
		if i == s.cursor {
			cursor = styles.Active.Render("> ")
			nameStyle = styles.ValueStyle
		}

		info := sessInfoStyle.Render(fmt.Sprintf(
			"%d msgs · %s",
			ses.MsgCount,
			ses.UpdatedAt.Format("2006-01-02 15:04"),
		))
		bodyLines = append(bodyLines, fmt.Sprintf(" %s%s%s%s  %s", cursor, nameStyle.Render(ses.Name), active, pinned, info))
	}

	if needsTop {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll))))
	}
	if needsBottom {
		remaining := len(s.sessions) - end
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↓ %d more", remaining))))
	}

	bodyLines = append(bodyLines, "")
	hintText := "esc: Close  ↵: Select  d:Delete  r:Rename  p:Pin  e:Export"
	bodyLines = append(bodyLines, fmt.Sprintf("%*s", innerW, styles.HintStyle.Render(hintText)))

	body := strings.Join(bodyLines, "\n")
	return styles.DialogBox(s.width, body)
}
