package screens

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"

	"termcode/internal/adapters/tui/components"
	"termcode/internal/adapters/tui/styles"
)

type BatchFile struct {
	Path      string
	Content   string
	OldStr    string
	NewStr    string
	MatchLine int
	MatchCol  int
	Selected  bool
	Applied   bool
	Error     string
}

type BatchEditScreen struct {
	width       int
	height      int
	files       []BatchFile
	cursor      int
	scroll      int
	viewport    *components.Viewport
	done        bool
	result      string
	pattern     string
	replacement string
	searchMode  bool
	searchInput string
	confirming  bool
	changed     int
	failed      int
}

func NewBatchEditScreen() *BatchEditScreen {
	vp := components.NewViewport(76, 10)
	return &BatchEditScreen{
		viewport:   vp,
		searchMode: true,
	}
}

func (s *BatchEditScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
	headerH := 8
	vpH := h - headerH
	if vpH < 3 {
		vpH = 3
	}
	s.viewport.SetSize(w-4, vpH)
}

func (s *BatchEditScreen) AddFile(path, oldStr, newStr string, matchLine, matchCol int) {
	data, _ := os.ReadFile(path)
	s.files = append(s.files, BatchFile{
		Path:      path,
		Content:   string(data),
		OldStr:    oldStr,
		NewStr:    newStr,
		MatchLine: matchLine,
		MatchCol:  matchCol,
		Selected:  true,
	})
	s.renderContent()
}

func (s *BatchEditScreen) SetFiles(files []BatchFile) {
	s.files = files
	s.renderContent()
}

func (s *BatchEditScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.SetSize(msg.Width, msg.Height)
		return s, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if s.confirming {
				s.confirming = false
				s.renderContent()
				return s, nil
			}
			s.done = true
			return s, nil

		case "enter":
			if s.searchMode {
				s.pattern = s.searchInput
				s.searchInput = ""
				s.searchMode = false
				s.renderContent()
				return s, nil
			}
			if s.confirming {
				s.applyEdits()
				s.done = true
				s.result = fmt.Sprintf("Applied %d changes, %d failed", s.changed, s.failed)
				return s, nil
			}
			if len(s.files) > 0 {
				f := &s.files[s.cursor]
				f.Selected = !f.Selected
				s.renderContent()
			}
			return s, nil

		case "a":
			if !s.searchMode && !s.confirming {
				allSel := true
				for i := range s.files {
					if !s.files[i].Selected {
						allSel = false
						break
					}
				}
				sel := !allSel
				for i := range s.files {
					s.files[i].Selected = sel
				}
				s.renderContent()
			}
			return s, nil

		case "s":
			if !s.searchMode && len(s.files) > 0 {
				s.confirming = true
				s.renderContent()
			}
			return s, nil

		case "up":
			if s.searchMode || s.confirming {
				return s, nil
			}
			if s.cursor > 0 {
				s.cursor--
				s.ensureVisible()
				s.renderContent()
			}
			return s, nil

		case "down":
			if s.searchMode || s.confirming {
				return s, nil
			}
			if s.cursor < len(s.files)-1 {
				s.cursor++
				s.ensureVisible()
				s.renderContent()
			}
			return s, nil

		case "backspace":
			if s.searchMode && len(s.searchInput) > 0 {
				s.searchInput = s.searchInput[:len(s.searchInput)-1]
				return s, nil
			}
			return s, nil

		default:
			if s.searchMode {
				key := msg.String()
				if len(key) == 1 && key[0] >= 32 && key[0] <= 126 {
					s.searchInput += key
				}
				return s, nil
			}
		}
	}
	return s, nil
}

func (s *BatchEditScreen) ensureVisible() {
	maxVisible := s.viewport.Height() - 1
	if maxVisible < 1 {
		maxVisible = 1
	}
	if s.cursor < s.scroll {
		s.scroll = s.cursor
	}
	if s.cursor >= s.scroll+maxVisible {
		s.scroll = s.cursor - maxVisible + 1
	}
}

func (s *BatchEditScreen) renderContent() {
	if s.searchMode {
		s.viewport.SetContent([]string{
			styles.Dim.Render("Enter search pattern:"),
			"",
			"> " + s.searchInput + "█",
			"",
			styles.Help.Render("Type pattern, Enter to confirm, Esc to cancel"),
		})
		return
	}

	if s.confirming {
		var lines []string
		lines = append(lines, styles.Title.Render("Confirm Batch Edit?"))
		lines = append(lines, "")
		var selected int
		for _, f := range s.files {
			if f.Selected {
				selected++
			}
		}
		lines = append(lines, styles.Dim.Render(fmt.Sprintf("Files: %d selected / %d total", selected, len(s.files))))
		lines = append(lines, "")
		for _, f := range s.files {
			if f.Selected {
				lines = append(lines, styles.Body.Render("  ✓ "+filepath.Base(f.Path)))
			}
		}
		lines = append(lines, "")
		lines = append(lines, styles.Help.Render(" Enter: apply  Esc: cancel"))
		s.viewport.SetContent(lines)
		return
	}

	if len(s.files) == 0 {
		s.viewport.SetContent([]string{
			styles.Dim.Render("No files matched"),
			"",
			styles.Help.Render("Esc to close"),
		})
		return
	}

	maxVisible := s.viewport.Height()
	var rendered []string
	selectedCount := 0
	for _, f := range s.files {
		if f.Selected {
			selectedCount++
		}
	}
	rendered = append(rendered, styles.Title.Render(fmt.Sprintf("Batch Edit: %d files (%d selected)", len(s.files), selectedCount)))
	rendered = append(rendered, "")

	for i := s.scroll; i < len(s.files); i++ {
		f := s.files[i]
		if len(rendered) >= maxVisible+1 {
			break
		}
		prefix := " "
		if f.Selected {
			prefix = "✓"
		}
		line := fmt.Sprintf("%s %s:%d", prefix, filepath.Base(f.Path), f.MatchLine)
		if i == s.cursor {
			rendered = append(rendered, styles.Active.Render("▸ "+line))
		} else if f.Selected {
			rendered = append(rendered, styles.ValueStyle.Render("  "+line))
		} else {
			rendered = append(rendered, styles.Inactive.Render("  "+line))
		}
		rendered = append(rendered, styles.Dim.Render("    "+f.OldStr+" → "+f.NewStr))
	}

	s.viewport.SetContent(rendered)
}

func (s *BatchEditScreen) applyEdits() {
	s.changed = 0
	s.failed = 0
	for i := range s.files {
		f := &s.files[i]
		if !f.Selected {
			continue
		}
		data, err := os.ReadFile(f.Path)
		if err != nil {
			f.Error = err.Error()
			s.failed++
			continue
		}
		content := string(data)
		n := strings.Count(content, f.OldStr)
		if n == 0 {
			f.Error = "pattern not found"
			s.failed++
			continue
		}
		content = strings.ReplaceAll(content, f.OldStr, f.NewStr)
		if err := os.WriteFile(f.Path, []byte(content), 0o644); err != nil {
			f.Error = err.Error()
			s.failed++
			continue
		}
		f.Applied = true
		s.changed += n
	}
}

func (s *BatchEditScreen) View() string {
	header := styles.Title.Render("Batch Edit")

	var status string
	if s.searchMode {
		status = styles.Dim.Render("Search mode: enter pattern to find")
	} else if s.confirming {
		status = styles.Dim.Render("Confirm batch edit")
	} else {
		sel := 0
		for _, f := range s.files {
			if f.Selected {
				sel++
			}
		}
		status = styles.Dim.Render(fmt.Sprintf("%d files, %d selected", len(s.files), sel))
	}

	sep := styles.DialogSep(s.width)
	body := header + "\n" +
		status + "\n" +
		sep + "\n" +
		s.viewport.View() + "\n" +
		sep + "\n"

	var help string
	if s.searchMode {
		help = styles.Help.Render(" Type pattern + Enter  Esc:cancel")
	} else if s.confirming {
		help = styles.Help.Render(" Enter: apply  Esc: cancel")
	} else {
		help = styles.Help.Render(" ↑/↓:nav  Enter:toggle  a:all  s:apply  Esc:cancel")
	}
	body += help

	return styles.DialogBox(s.width, body)
}

func (s *BatchEditScreen) Done() bool {
	return s.done
}

func (s *BatchEditScreen) Result() string {
	return s.result
}
