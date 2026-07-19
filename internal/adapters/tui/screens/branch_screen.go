package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type BranchItem struct {
	ID       string
	Name     string
	MsgCount int
	IsActive bool
}

type BranchScreen struct {
	width    int
	height   int
	branches []BranchItem
	cursor   int
	scroll   int
	done     bool
	result   string
}

func NewBranchScreen() *BranchScreen {
	return &BranchScreen{
		width:  80,
		height: 24,
	}
}

func (s *BranchScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *BranchScreen) SetBranches(branches []BranchItem) {
	s.branches = branches
	if s.cursor >= len(s.branches) {
		s.cursor = len(s.branches) - 1
	}
}

func (s *BranchScreen) Done() bool     { return s.done }
func (s *BranchScreen) Result() string { return s.result }

func (s *BranchScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "enter":
			if len(s.branches) > 0 && s.cursor >= 0 && s.cursor < len(s.branches) {
				s.result = "__branch__:" + s.branches[s.cursor].ID
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
			if s.cursor < len(s.branches)-1 {
				s.cursor++
				maxVis := s.height - 10
				if maxVis < 1 {
					maxVis = 1
				}
				if s.cursor >= s.scroll+maxVis {
					s.scroll = s.cursor - maxVis + 1
				}
			}
		}
	}
	return s, nil
}

func (s *BranchScreen) View() string {
	innerW := s.width - 2

	title := fmt.Sprintf("%-*s%s", innerW-4, "Conversation Branches", "esc")

	if len(s.branches) == 0 {
		body := fmt.Sprintf("%s\n%s\n%s",
			title,
			styles.DialogSep(innerW),
			styles.HintStyle.Render("No branches. Use /branch <name> to create one."),
		)
		return styles.DialogBox(s.width, body)
	}

	maxVis := s.height - 10
	if maxVis < 1 {
		maxVis = 1
	}
	if s.scroll+maxVis > len(s.branches) {
		s.scroll = max(0, len(s.branches)-maxVis)
	}
	end := min(s.scroll+maxVis, len(s.branches))

	var bodyLines []string
	bodyLines = append(bodyLines, title)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, "")

	activeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	inactiveStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	cursorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("214"))

	for i := s.scroll; i < end; i++ {
		b := s.branches[i]
		mark := "  "
		nameStyle := inactiveStyle
		if i == s.cursor {
			mark = cursorStyle.Render("▸ ")
			nameStyle = styles.ValueStyle
		}
		active := ""
		if b.IsActive {
			active = activeStyle.Render(" [active]")
		}
		bodyLines = append(bodyLines, fmt.Sprintf(" %s%s%s  %d msgs", mark, nameStyle.Render(b.Name), active, b.MsgCount))
	}

	bodyLines = append(bodyLines, "")
	bodyLines = append(bodyLines, fmt.Sprintf("%*s", innerW, styles.HintStyle.Render("esc: Close  ↵: Switch to branch")))

	body := strings.Join(bodyLines, "\n")
	return styles.DialogBox(s.width, body)
}
