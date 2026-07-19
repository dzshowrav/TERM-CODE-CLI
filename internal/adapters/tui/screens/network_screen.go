package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type NetworkStatus int

const (
	NetworkUnknown NetworkStatus = iota
	NetworkOnline
	NetworkOffline
	NetworkLimited
)

type NetworkState struct {
	Status      NetworkStatus
	ProviderURL string
	ProviderOK  bool
	Latency     string
	MCPCount    int
	MCPOnline   int
	GitOK       bool
}

type NetworkScreen struct {
	width  int
	height int
	done   bool
	state  NetworkState
}

func NewNetworkScreen(state NetworkState) *NetworkScreen {
	return &NetworkScreen{
		state: state,
	}
}

func (s *NetworkScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *NetworkScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "enter", "q":
			s.done = true
		}
	}
	return s, nil
}

func (s *NetworkScreen) View() string {
	if s.width < 10 {
		return "Network Status"
	}

	body := s.renderBody()
	return styles.DialogBox(s.width-4, body)
}

func (s *NetworkScreen) renderBody() string {
	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	valueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	statusStyle := func(ok bool, label string) string {
		if ok {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Render("●") +
				" " + label
		}
		return lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("●") +
			" " + label
	}

	lines := []string{
		lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true).Render("Network Status"),
		"",
	}

	switch s.state.Status {
	case NetworkOnline:
		lines = append(lines, statusStyle(true, "Connected"))
	case NetworkOffline:
		lines = append(lines, statusStyle(false, "Offline"))
	case NetworkLimited:
		lines = append(lines, lipgloss.NewStyle().Foreground(lipgloss.Color("214")).Render("● Limited"))
	default:
		lines = append(lines, lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("○ Unknown"))
	}

	lines = append(lines, "")
	lines = append(lines, labelStyle.Render("Provider URL:   ")+valueStyle.Render(s.state.ProviderURL))
	if s.state.ProviderOK {
		lines = append(lines, labelStyle.Render("Provider:       ")+statusStyle(true, fmt.Sprintf("OK (%s)", s.state.Latency)))
	} else {
		lines = append(lines, labelStyle.Render("Provider:       ")+statusStyle(false, "Unreachable"))
	}

	lines = append(lines, "")
	mcpStatus := fmt.Sprintf("%d / %d online", s.state.MCPOnline, s.state.MCPCount)
	lines = append(lines, labelStyle.Render("MCP Servers:    ")+valueStyle.Render(mcpStatus))
	lines = append(lines, labelStyle.Render("Git:            ")+statusStyle(s.state.GitOK, fmt.Sprintf("%v", s.state.GitOK)))
	lines = append(lines, "")
	lines = append(lines, styles.HintStyle.Render("Press ESC or q to close"))

	innerW := s.width - 8
	if innerW < 20 {
		innerW = 20
	}

	var b strings.Builder
	for _, line := range lines {
		vw := lipgloss.Width(line)
		if vw < innerW {
			b.WriteString(line)
			b.WriteString(strings.Repeat(" ", innerW-vw))
		} else {
			b.WriteString(lipgloss.NewStyle().Width(innerW).Render(line))
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (s *NetworkScreen) Done() bool         { return s.done }
func (s *NetworkScreen) Result() string     { return "" }
func (s *NetworkScreen) Title() string      { return "Network Status" }
func (s *NetworkScreen) IsFullScreen() bool { return false }
