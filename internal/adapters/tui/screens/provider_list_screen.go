package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	provConnected = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Render("●")
	provDisconn   = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("○")
	provAuthFail  = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("●")
	provNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	provURLStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	provDefStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render(" (active)")
)

type ProviderItem struct {
	Name     string
	URL      string
	Status   string
	Latency  int64
	IsActive bool
}

type ProviderListScreen struct {
	width     int
	height    int
	providers []ProviderItem
}

func NewProviderListScreen() *ProviderListScreen {
	return &ProviderListScreen{
		width:  80,
		height: 24,
	}
}

func (s *ProviderListScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ProviderListScreen) SetProviders(providers []ProviderItem) {
	s.providers = providers
}

func (s *ProviderListScreen) View() string {
	header := styles.H1.Render("Provider Management")
	sep := styles.SeparatorLine(s.width)

	if len(s.providers) == 0 {
		empty := lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("No providers configured.")
		hint := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Use /provider add to add one.")
		return fmt.Sprintf("%s\n%s\n%s\n%s", header, sep, empty, hint)
	}

	var lines []string
	for _, p := range s.providers {
		icon := provDisconn
		switch p.Status {
		case "connected":
			icon = provConnected
		case "auth_failed":
			icon = provAuthFail
		}

		active := ""
		if p.IsActive {
			active = provDefStyle
		}

		latency := ""
		if p.Latency > 0 {
			latency = fmt.Sprintf(" %dms", p.Latency)
		}

		lines = append(lines, fmt.Sprintf(" %s %s%s", icon, provNameStyle.Render(p.Name), active))
		lines = append(lines, fmt.Sprintf("   %s%s", provURLStyle.Render(p.URL), latency))
		lines = append(lines, "")
	}

	return fmt.Sprintf("%s\n%s\n%s", header, sep, strings.Join(lines, "\n"))
}
