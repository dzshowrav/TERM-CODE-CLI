package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"

	"termcode/internal/adapters/tui/components"
	"termcode/internal/adapters/tui/styles"
)

type DiffFile struct {
	Path    string
	Status  string
	Added   int
	Removed int
	Content string
}

type DiffScreen struct {
	width      int
	height     int
	files      []DiffFile
	fileCursor int
	scroll     int
	viewport   *components.Viewport
	done       bool
	result     string
}

func NewDiffScreen(files []DiffFile) *DiffScreen {
	vp := components.NewViewport(76, 10)
	return &DiffScreen{
		files:    files,
		viewport: vp,
		done:     false,
	}
}

func (s *DiffScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
	headerH := 4
	vpH := h - headerH
	if vpH < 3 {
		vpH = 3
	}
	s.viewport.SetSize(w-4, vpH)
}

func (s *DiffScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.SetSize(msg.Width, msg.Height)
		return s, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q":
			s.done = true
			return s, nil

		case "right", "n", "tab":
			if s.fileCursor < len(s.files)-1 {
				s.fileCursor++
				s.scroll = 0
				s.renderContent()
			}
			return s, nil

		case "left", "p", "shift+tab":
			if s.fileCursor > 0 {
				s.fileCursor--
				s.scroll = 0
				s.renderContent()
			}
			return s, nil

		case "up":
			if s.scroll > 0 {
				s.scroll--
				s.renderContent()
			}
			return s, nil

		case "down":
			if len(s.files) > 0 {
				contentLines := strings.Split(s.files[s.fileCursor].Content, "\n")
				maxScroll := len(contentLines) - s.viewport.Height()
				if maxScroll < 0 {
					maxScroll = 0
				}
				if s.scroll < maxScroll {
					s.scroll++
					s.renderContent()
				}
			}
			return s, nil

		case "home":
			s.scroll = 0
			s.renderContent()
			return s, nil

		case "end":
			contentLines := strings.Split(s.files[s.fileCursor].Content, "\n")
			s.scroll = len(contentLines) - s.viewport.Height()
			if s.scroll < 0 {
				s.scroll = 0
			}
			s.renderContent()
			return s, nil

		case "enter":
			s.done = true
			s.result = "diff"
			return s, nil
		}
	}
	return s, nil
}

func (s *DiffScreen) renderContent() {
	if len(s.files) == 0 {
		s.viewport.SetContent([]string{styles.Dim.Render("No changes")})
		return
	}
	if s.fileCursor >= len(s.files) {
		s.fileCursor = len(s.files) - 1
	}
	f := s.files[s.fileCursor]
	lines := strings.Split(f.Content, "\n")

	rendered := make([]string, 0, len(lines))
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "+++") || strings.HasPrefix(line, "---"):
			rendered = append(rendered, styles.Dim.Render(line))
		case strings.HasPrefix(line, "@@"):
			rendered = append(rendered, styles.Accent.Render(line))
		case strings.HasPrefix(line, "+"):
			rendered = append(rendered, styles.DiffAdded.Render(line))
		case strings.HasPrefix(line, "-"):
			rendered = append(rendered, styles.DiffRemoved.Render(line))
		default:
			rendered = append(rendered, line)
		}
	}

	maxLines := s.viewport.Height()
	maxScroll := len(rendered) - maxLines
	if maxScroll < 0 {
		maxScroll = 0
	}
	if s.scroll > maxScroll {
		s.scroll = maxScroll
	}

	if s.scroll > 0 {
		rendered = rendered[s.scroll:]
	}
	if len(rendered) > maxLines {
		rendered = rendered[:maxLines]
	}

	s.viewport.SetContent(rendered)
}

func (s *DiffScreen) View() string {
	if len(s.files) == 0 {
		return styles.Dim.Render("No changes to display")
	}

	header := styles.Title.Render(fmt.Sprintf("Diff View (%d file(s))", len(s.files)))

	totalAdd := 0
	totalDel := 0
	for _, f := range s.files {
		totalAdd += f.Added
		totalDel += f.Removed
	}

	var fileNav strings.Builder
	for i, f := range s.files {
		sep := " "
		if i == s.fileCursor {
			fileNav.WriteString(styles.Active.Render(fmt.Sprintf(" [%s]", f.Path)))
		} else {
			fileNav.WriteString(styles.Dim.Render(fmt.Sprintf(" %s", f.Path)))
		}
		if i < len(s.files)-1 {
			fileNav.WriteString(sep)
		}
	}

	stats := styles.Dim.Render(fmt.Sprintf("+%d -%d  (%d/%d)",
		totalAdd, totalDel, s.fileCursor+1, len(s.files)))

	sep := strings.Repeat("─", s.width)
	body := header + "\n" +
		stats + "\n" +
		styles.Dim.Render(sep) + "\n" +
		s.viewport.View() + "\n" +
		styles.Dim.Render(sep) + "\n" +
		fileNav.String() + "\n" +
		styles.Help.Render(" esc:close  ←/→:files  ↑/↓:scroll  enter:accept")

	return styles.DialogBox(s.width, body)
}

func (s *DiffScreen) Done() bool {
	return s.done
}

func (s *DiffScreen) Result() string {
	return s.result
}
