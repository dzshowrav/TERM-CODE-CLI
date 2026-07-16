package screens

import (
	"fmt"

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

func (s *HomeScreen) View() string {
	title := styles.TitleStyle.Render("TERM CODE CLI")
	tagline := styles.Subtitle.Render("Universal Coding Agent")

	providerLabel := styles.LabelStyle.Render("Provider : ")
	providerVal := styles.ValueStyle.Render(fmt.Sprintf("%s (%s)", s.providerName, s.providerURL))

	modelLabel := styles.LabelStyle.Render("Model    : ")
	modelVal := styles.ValueStyle.Render(s.modelName)

	agentLabel := styles.LabelStyle.Render("Agent    : ")
	agentVal := styles.ValueStyle.Render(s.agentName)

	wsLabel := styles.LabelStyle.Render("Workspace: ")
	wsVal := styles.ValueStyle.Render(s.workspacePath)

	hint := styles.HintStyle.Render("Type a message or / for commands")

	availableLines := s.height - 10
	if availableLines < 1 {
		availableLines = 1
	}
	spacing := ""
	for i := 0; i < availableLines/2; i++ {
		spacing += "\n"
	}

	separator := styles.Separator.Render("─")

	content := fmt.Sprintf(
		"%s\n%s\n%s\n%s%s\n%s%s\n%s%s\n%s%s\n%s\n%s%s",
		title,
		tagline,
		separator,
		providerLabel, providerVal,
		modelLabel, modelVal,
		agentLabel, agentVal,
		wsLabel, wsVal,
		separator,
		spacing,
		hint,
	)

	return styles.Screen.Copy().
		Width(s.width).
		Height(s.height).
		Render(content)
}
