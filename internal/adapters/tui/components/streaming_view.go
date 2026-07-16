package components

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	streamingStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	cursorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
)

type StreamingView struct {
	content string
	cursor  string
	width   int
}

func NewStreamingView(width int) *StreamingView {
	return &StreamingView{
		cursor: "▌",
		width:  width,
	}
}

func (s *StreamingView) Append(text string) {
	s.content += text
}

func (s *StreamingView) SetContent(text string) {
	s.content = text
}

func (s *StreamingView) Content() string {
	return s.content
}

func (s *StreamingView) SetWidth(w int) {
	s.width = w
}

func (s *StreamingView) Finish() {
	s.cursor = ""
}

func (s *StreamingView) View() string {
	if s.content == "" {
		return cursorStyle.Render(s.cursor)
	}

	display := s.content
	if s.width > 0 && len(display) > s.width {
		display = display[:s.width]
	}

	return streamingStyle.Render(display) + cursorStyle.Render(s.cursor)
}
