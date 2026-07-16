package screens

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/components"
	"termcode/internal/adapters/tui/styles"
)

type ChatScreen struct {
	width     int
	height    int
	viewport  *components.Viewport
	modelName string
	messages  []ChatMessage
}

type ChatMessage struct {
	Role    string
	Content string
}

var (
	userStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).SetString("You")
	assistantStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).SetString("AI")
	systemStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).SetString("System")
)

func NewChatScreen() *ChatScreen {
	return &ChatScreen{
		width:     80,
		height:    24,
		viewport:  components.NewViewport(80, 20),
		modelName: "none",
		messages:  []ChatMessage{},
	}
}

func (s *ChatScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
	vpHeight := h - 1
	if vpHeight < 1 {
		vpHeight = 1
	}
	s.viewport.SetSize(w, vpHeight)
}

func (s *ChatScreen) SetModel(name string) {
	s.modelName = name
}

func (s *ChatScreen) AddMessage(role, content string) {
	s.messages = append(s.messages, ChatMessage{Role: role, Content: content})
	s.renderMessages()
}

func (s *ChatScreen) renderMessages() {
	var lines []string
	for _, msg := range s.messages {
		roleLabel := ""
		switch msg.Role {
		case "user":
			roleLabel = userStyle.String()
		case "assistant":
			roleLabel = assistantStyle.String()
		default:
			roleLabel = systemStyle.String()
		}
		lines = append(lines, roleLabel)
		lines = append(lines, msg.Content)
		lines = append(lines, "")
	}
	s.viewport.SetContent(lines)
}

func (s *ChatScreen) View() string {
	modelLine := styles.H2.Copy().Width(s.width).Render(s.modelName)
	separator := styles.SeparatorLine(s.width)
	vp := s.viewport.View()

	return fmt.Sprintf("%s\n%s\n%s", modelLine, separator, vp)
}
