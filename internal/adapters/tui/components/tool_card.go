package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	toolCardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(0, 1)
	toolNameStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Bold(true)
	toolStatusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	toolSuccessIcon = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Render("✓")
	toolFailIcon    = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("✗")
	toolRunIcon     = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("◐")
)

type ToolCard struct {
	name    string
	status  string
	content string
	width   int
}

func NewToolCard(name, status, content string) *ToolCard {
	return &ToolCard{
		name:    name,
		status:  status,
		content: content,
		width:   80,
	}
}

func (c *ToolCard) SetWidth(w int) {
	c.width = w
}

func (c *ToolCard) SetStatus(status string) {
	c.status = status
}

func (c *ToolCard) Render() string {
	icon := toolRunIcon
	switch c.status {
	case "success", "completed":
		icon = toolSuccessIcon
	case "failed", "error":
		icon = toolFailIcon
	}

	header := fmt.Sprintf("%s %s (%s)", icon, toolNameStyle.Render(c.name), toolStatusStyle.Render(c.status))

	cardWidth := c.width - 4
	if cardWidth < 20 {
		cardWidth = 20
	}

	content := header
	if c.content != "" {
		content += "\n" + c.content
	}

	return toolCardStyle.Width(cardWidth).Render(content)
}

func (c *ToolCard) Lines() []string {
	return strings.Split(c.Render(), "\n")
}
