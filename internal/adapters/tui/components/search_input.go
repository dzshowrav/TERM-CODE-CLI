package components

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var (
	searchPrompt = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("search: ")
	searchStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
)

type SearchInput struct {
	value    string
	cursor   int
	focused  bool
	width    int
	onSubmit func(string) tea.Msg
}

func NewSearchInput(onSubmit func(string) tea.Msg) *SearchInput {
	return &SearchInput{
		focused:  true,
		width:    80,
		onSubmit: onSubmit,
	}
}

func (s *SearchInput) SetWidth(w int) {
	s.width = w
}

func (s *SearchInput) Value() string {
	return s.value
}

func (s *SearchInput) Focused() bool {
	return s.focused
}

func (s *SearchInput) SetFocused(f bool) {
	s.focused = f
}

func (s *SearchInput) Init() tea.Cmd {
	return nil
}

func (s *SearchInput) Update(msg tea.Msg) (*SearchInput, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if s.onSubmit != nil {
				return s, func() tea.Msg {
					return s.onSubmit(s.value)
				}
			}
		case "backspace":
			if s.cursor > 0 {
				s.value = s.value[:s.cursor-1] + s.value[s.cursor:]
				s.cursor--
			}
		case "delete":
			if s.cursor < len(s.value) {
				s.value = s.value[:s.cursor] + s.value[s.cursor+1:]
			}
		case "left":
			if s.cursor > 0 {
				s.cursor--
			}
		case "right":
			if s.cursor < len(s.value) {
				s.cursor++
			}
		default:
			if len(msg.String()) == 1 && msg.String()[0] >= 32 {
				r := msg.String()
				s.value = s.value[:s.cursor] + r + s.value[s.cursor:]
				s.cursor++
			}
		}
	}
	return s, nil
}

func (s *SearchInput) View() string {
	display := s.value
	dispLen := len(display)
	maxLen := s.width - len(searchPrompt) - 2
	if dispLen > maxLen {
		display = display[dispLen-maxLen:]
	}

	cursor := ""
	if s.focused {
		cursor = "█"
	}

	return searchPrompt + searchStyle.Render(display[:s.cursor]+cursor+display[s.cursor:])
}
