package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	gitFileNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	gitStatusStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	gitAddedStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("83"))
	gitModifiedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
	gitDeletedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	gitSelectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
)

type GitFileItem struct {
	Name   string
	Status string
}

type GitAddScreen struct {
	width    int
	height   int
	files    []GitFileItem
	selected map[int]bool
	cursor   int
	scroll   int
	done     bool
	result   string
	onStage  func(files []string) string
}

func NewGitAddScreen(files []GitFileItem, onStage func(files []string) string) *GitAddScreen {
	return &GitAddScreen{
		width:    80,
		height:   24,
		files:    files,
		selected: make(map[int]bool),
		onStage:  onStage,
	}
}

func (s *GitAddScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *GitAddScreen) Done() bool     { return s.done }
func (s *GitAddScreen) Result() string { return s.result }

func (s *GitAddScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
				if s.cursor < s.scroll {
					s.scroll = s.cursor
				}
			}
		case "down", "j":
			if s.cursor < len(s.files)-1 {
				s.cursor++
				maxVis := s.height - 9
				if maxVis < 1 {
					maxVis = 1
				}
				if s.cursor >= s.scroll+maxVis {
					s.scroll = s.cursor - maxVis + 1
				}
			}
		case " ", "enter":
			if len(s.files) > 0 && s.cursor >= 0 && s.cursor < len(s.files) {
				s.selected[s.cursor] = !s.selected[s.cursor]
			}
		case "a":
			if len(s.files) == 0 {
				break
			}
			allSel := true
			for i := range s.files {
				if !s.selected[i] {
					allSel = false
					break
				}
			}
			for i := range s.files {
				s.selected[i] = !allSel
			}
		case "s":
			if s.onStage != nil {
				var files []string
				for i := range s.files {
					if s.selected[i] {
						files = append(files, s.files[i].Name)
					}
				}
				if len(files) == 0 {
					for i := range s.files {
						files = append(files, s.files[i].Name)
					}
				}
				s.result = s.onStage(files)
				s.done = true
			}
		}
	}
	return s, nil
}

func (s *GitAddScreen) View() string {
	innerW := s.width - 2

	title := fmt.Sprintf("%-*s%s", innerW-4, "Stage Files", "esc")

	if len(s.files) == 0 {
		body := fmt.Sprintf("%s\n%s\n%s",
			title,
			styles.DialogSep(innerW),
			styles.HintStyle.Render("No unstaged files."),
		)
		return styles.DialogBox(s.width, body)
	}

	needsTop := s.scroll > 0
	needsBottom := false

	maxVis := s.height - 9
	if maxVis < 1 {
		maxVis = 1
	}
	arrowCount := 0
	if needsTop {
		arrowCount++
	}
	if s.scroll+maxVis-arrowCount < len(s.files) {
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

	if s.scroll+maxVis > len(s.files) {
		s.scroll = max(0, len(s.files)-maxVis)
	}
	end := min(s.scroll+maxVis, len(s.files))

	var bodyLines []string
	bodyLines = append(bodyLines, title)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, "")

	for i := s.scroll; i < end; i++ {
		f := s.files[i]
		sel := " "
		if s.selected[i] {
			sel = styles.Active.Render("✓")
		}
		statusClr := gitModifiedStyle
		switch f.Status {
		case "A":
			statusClr = gitAddedStyle
		case "D":
			statusClr = gitDeletedStyle
		case "M":
			statusClr = gitModifiedStyle
		case "??":
			statusClr = gitStatusStyle
		}
		status := statusClr.Render(f.Status)

		cursor := " "
		nameStyle := gitFileNameStyle
		if i == s.cursor {
			if s.selected[i] {
				cursor = gitSelectedStyle.Render(">")
			} else {
				cursor = styles.Active.Render(">")
			}
			nameStyle = styles.ValueStyle
		}

		bodyLines = append(bodyLines, fmt.Sprintf(" %s %s %s%s", cursor, sel, status, nameStyle.Render(" "+f.Name)))
	}

	if needsTop {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll))))
	}
	if needsBottom {
		remaining := len(s.files) - end
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↓ %d more", remaining))))
	}

	selectedCount := 0
	for _, sel := range s.selected {
		if sel {
			selectedCount++
		}
	}

	bodyLines = append(bodyLines, "")
	hintText := fmt.Sprintf("arrows: Navigate  SPACE: Select  a: All/none  s: Stage (%d)  esc: Cancel", selectedCount)
	bodyLines = append(bodyLines, fmt.Sprintf("%*s", innerW, styles.HintStyle.Render(hintText)))

	body := strings.Join(bodyLines, "\n")
	return styles.DialogBox(s.width, body)
}
