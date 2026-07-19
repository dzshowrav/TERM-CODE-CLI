package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type AgentEntry struct {
	Name        string
	Description string
}

type AgentSelectScreen struct {
	width    int
	height   int
	agents   []AgentEntry
	cursor   int
	scroll   int
	done     bool
	result   string
	onSelect func(name string) string
}

func NewAgentSelectScreen(onSelect func(name string) string) *AgentSelectScreen {
	return &AgentSelectScreen{
		width:    80,
		height:   24,
		onSelect: onSelect,
		agents: []AgentEntry{
			{Name: "General", Description: "General-purpose coding assistant"},
			{Name: "Expert", Description: "Expert-level code review and debugging"},
			{Name: "Architect", Description: "System design and architecture"},
		},
	}
}

func (s *AgentSelectScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *AgentSelectScreen) Done() bool     { return s.done }
func (s *AgentSelectScreen) Result() string { return s.result }

func (s *AgentSelectScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "enter":
			if len(s.agents) > 0 && s.cursor >= 0 && s.cursor < len(s.agents) {
				if s.onSelect != nil {
					s.result = s.onSelect(s.agents[s.cursor].Name)
				}
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
			if s.cursor < len(s.agents)-1 {
				s.cursor++
				maxVis := s.height - 8
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

func (s *AgentSelectScreen) View() string {
	innerW := s.width - 2

	title := fmt.Sprintf("%-*s%s", innerW-4, "Agent Selection", "esc")

	maxVis := s.height - 8
	if maxVis < 1 {
		maxVis = 1
	}
	if s.scroll+maxVis > len(s.agents) {
		s.scroll = max(0, len(s.agents)-maxVis)
	}
	end := min(s.scroll+maxVis, len(s.agents))

	var bodyLines []string
	bodyLines = append(bodyLines, title)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, "")

	for i := s.scroll; i < end; i++ {
		a := s.agents[i]
		cursor := "  "
		nameStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
		if i == s.cursor {
			cursor = styles.Active.Render("> ")
			nameStyle = styles.ValueStyle
		}
		bodyLines = append(bodyLines, fmt.Sprintf(" %s%s", cursor, nameStyle.Render(a.Name)))
		bodyLines = append(bodyLines, fmt.Sprintf("    %s", lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render(a.Description)))
		bodyLines = append(bodyLines, "")
	}

	if s.scroll > 0 {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll))))
	}
	if end < len(s.agents) {
		remaining := len(s.agents) - end
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↓ %d more", remaining))))
	}

	bodyLines = append(bodyLines, "")
	hintText := "esc: Close  ↵: Select"
	bodyLines = append(bodyLines, fmt.Sprintf("%*s", innerW, styles.HintStyle.Render(hintText)))

	body := strings.Join(bodyLines, "\n")
	return styles.DialogBox(s.width, body)
}
