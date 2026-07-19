package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type HomeScreen struct {
	width         int
	height        int
	providerName  string
	modelName     string
	agentName     string
	workspacePath string
	providerURL   string
	anim          *BackgroundAnim
}

type HomeScreenConfig struct {
	ProviderName  string
	ProviderURL   string
	ModelName     string
	AgentName     string
	WorkspacePath string
}

func NewHomeScreen(cfg HomeScreenConfig) *HomeScreen {
	return &HomeScreen{
		width:         80,
		height:        24,
		providerName:  cfg.ProviderName,
		providerURL:   cfg.ProviderURL,
		modelName:     cfg.ModelName,
		agentName:     cfg.AgentName,
		workspacePath: cfg.WorkspacePath,
	}
}

func (s *HomeScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *HomeScreen) UpdateConfig(cfg HomeScreenConfig) {
	if cfg.ProviderName != "" {
		s.providerName = cfg.ProviderName
	}
	if cfg.ProviderURL != "" {
		s.providerURL = cfg.ProviderURL
	}
	if cfg.ModelName != "" {
		s.modelName = cfg.ModelName
	}
	if cfg.AgentName != "" {
		s.agentName = cfg.AgentName
	}
	if cfg.WorkspacePath != "" {
		s.workspacePath = cfg.WorkspacePath
	}
}

func (s *HomeScreen) SetAnim(a *BackgroundAnim) {
	s.anim = a
}

func (s *HomeScreen) View() string {
	availH := s.height - 4
	if availH < 3 {
		return ""
	}

	cw := s.width

	var titleStyle = styles.TitleStyle
	if s.anim != nil {
		gs := s.anim.GlowStyle()
		titleStyle = styles.TitleStyle.Foreground(gs.GetForeground())
	}
	title := titleStyle.Width(cw).Align(lipgloss.Center).Render("TERM CODE CLI")
	tagline := styles.Subtitle.Width(cw).Align(lipgloss.Center).Render("Universal Coding Agent")

	line1 := styles.Content(cw, "Provider : "+
		styles.ValueStyle.Render(s.providerName))

	line2 := styles.Content(cw, "Model    : "+
		styles.ValueStyle.Render(s.modelName))

	line3 := styles.Content(cw, "Agent    : "+
		styles.ValueStyle.Render(s.agentName))

	line4 := styles.Content(cw, "Workspace: "+
		styles.ValueStyle.Render(s.workspacePath))

	sep := styles.SeparatorLine(cw)
	hint := styles.HintStyle.Width(cw).Align(lipgloss.Center).Render("Press / for commands")

	body := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n\n%s",
		title, tagline, sep,
		line1, line2, line3, line4,
		sep, hint,
	)

	lines := strings.Split(body, "\n")
	padding := availH - len(lines)
	if padding > 0 {
		tp := padding / 2
		bp := padding - tp
		p := make([]string, 0, availH)
		for i := 0; i < tp; i++ {
			p = append(p, "")
		}
		p = append(p, lines...)
		for i := 0; i < bp; i++ {
			p = append(p, "")
		}
		lines = p
	}
	if len(lines) > availH {
		lines = lines[:availH]
	}

	if s.anim != nil {
		lines = s.anim.RenderParticles(lines)
	}

	return strings.Join(lines, "\n")
}
