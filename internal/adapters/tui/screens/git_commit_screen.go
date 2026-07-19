package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	commitLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	commitInputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
)

type GitCommitScreen struct {
	width       int
	height      int
	message     string
	description string
	focusField  int
	done        bool
	result      string
	onCommit    func(msg string) string
}

func NewGitCommitScreen(onCommit func(msg string) string) *GitCommitScreen {
	return &GitCommitScreen{
		width:      80,
		height:     24,
		focusField: 0,
		onCommit:   onCommit,
	}
}

func (s *GitCommitScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *GitCommitScreen) Done() bool     { return s.done }
func (s *GitCommitScreen) Result() string { return s.result }

func (s *GitCommitScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.PasteMsg:
		paste := msg.String()
		if s.focusField == 0 {
			s.message += paste
		} else if s.focusField == 1 {
			s.description += paste
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "enter":
			if s.focusField == 0 {
				if s.message != "" {
					if s.onCommit != nil {
						msg := s.message
						if s.description != "" {
							msg = s.message + "\n\n" + s.description
						}
						s.result = s.onCommit(msg)
					}
					s.done = true
				}
			} else if s.focusField == 2 {
				if s.message != "" {
					if s.onCommit != nil {
						msg := s.message
						if s.description != "" {
							msg = s.message + "\n\n" + s.description
						}
						s.result = s.onCommit(msg)
					}
					s.done = true
				}
			}
		case "tab", "down":
			s.focusField = (s.focusField + 1) % 3
		case "shift+tab", "up":
			s.focusField--
			if s.focusField < 0 {
				s.focusField = 2
			}
		case "backspace":
			if s.focusField == 0 && len(s.message) > 0 {
				s.message = s.message[:len(s.message)-1]
			} else if s.focusField == 1 && len(s.description) > 0 {
				s.description = s.description[:len(s.description)-1]
			}
		default:
			if len(msg.String()) == 1 {
				if s.focusField == 0 {
					s.message += msg.String()
				} else if s.focusField == 1 {
					s.description += msg.String()
				}
			}
		}
	}
	return s, nil
}

func (s *GitCommitScreen) View() string {
	innerW := s.width - 2
	sep := styles.DialogSep(innerW)

	var lines []string
	lines = append(lines, "")
	lines = append(lines, styles.H1.Render("Commit Changes"))
	lines = append(lines, sep)

	msgVal := s.message
	descVal := s.description

	if msgVal == "" && s.focusField != 0 {
		msgVal = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Brief summary of changes")
	}
	if s.focusField == 0 {
		msgVal += "█"
	}

	lines = append(lines, fmt.Sprintf(" %s  %s", commitLabelStyle.Render("Message:"), commitInputStyle.Render(msgVal)))
	lines = append(lines, sep)

	if descVal == "" && s.focusField != 1 {
		descVal = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Optional detailed description")
	}
	if s.focusField == 1 {
		descVal += "█"
	}

	lines = append(lines, fmt.Sprintf(" %s  %s", commitLabelStyle.Render("Details:"), commitInputStyle.Render(descVal)))
	lines = append(lines, sep)
	lines = append(lines, "")

	commitBtn := formBtnNormal.Render("[ Commit ]")
	if s.focusField == 2 {
		commitBtn = formBtnActive.Render("[ Commit ]")
	}
	btnLine := lipgloss.NewStyle().Width(innerW).Align(lipgloss.Center).Render(commitBtn)
	lines = append(lines, btnLine)
	lines = append(lines, "")

	hintText := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Tab: Navigate  ESC: Cancel")
	hintLine := lipgloss.Place(innerW, 1, lipgloss.Center, lipgloss.Center, hintText)
	lines = append(lines, hintLine)
	lines = append(lines, "")

	return styles.DialogBox(s.width, strings.Join(lines, "\n"))
}
