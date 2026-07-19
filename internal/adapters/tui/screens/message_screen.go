package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"

	"termcode/internal/adapters/tui/styles"
)

type MessageScreen struct {
	width   int
	height  int
	title   string
	content string
	scroll  int
	done    bool
}

func NewMessageScreen(title, content string) *MessageScreen {
	return &MessageScreen{
		width:   80,
		height:  24,
		title:   title,
		content: content,
	}
}

func (s *MessageScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *MessageScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "up", "k":
			if s.scroll > 0 {
				s.scroll--
			}
		case "down", "j":
			if s.scroll < len(strings.Split(s.content, "\n"))-1 {
				s.scroll++
			}
		}
	}
	return s, nil
}

func (s *MessageScreen) Done() bool     { return s.done }
func (s *MessageScreen) Result() string { return "" }

func (s *MessageScreen) View() string {
	innerW := s.width - 2
	contentLines := strings.Split(s.content, "\n")

	maxVis := s.height - 6
	if maxVis < 1 {
		maxVis = 1
	}

	needsTop := s.scroll > 0
	needsBottom := s.scroll+maxVis < len(contentLines)
	if needsTop {
		maxVis--
	}
	if needsBottom {
		maxVis--
	}
	if maxVis < 1 {
		maxVis = 1
	}

	if s.scroll+maxVis > len(contentLines) {
		s.scroll = max(0, len(contentLines)-maxVis)
	}
	end := s.scroll + maxVis
	if end > len(contentLines) {
		end = len(contentLines)
	}

	var bodyLines []string
	bodyLines = append(bodyLines, fmt.Sprintf("%-*s%s", innerW-4, s.title, "esc"))
	bodyLines = append(bodyLines, styles.DialogSep(innerW))

	for i := s.scroll; i < end; i++ {
		bodyLines = append(bodyLines, fmt.Sprintf(" %s", contentLines[i]))
	}

	if needsTop {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll))))
	}
	if needsBottom {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↓ %d more", len(contentLines)-end))))
	}

	bodyLines = append(bodyLines, "")
	bodyLines = append(bodyLines, fmt.Sprintf("%*s", innerW, styles.HintStyle.Render("arrows: Scroll  esc: Close")))

	return styles.DialogBox(s.width, strings.Join(bodyLines, "\n"))
}
