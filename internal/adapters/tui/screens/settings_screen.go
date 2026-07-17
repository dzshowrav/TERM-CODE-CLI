package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	settingLabel = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	settingValue = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	settingGroup = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Bold(true)
)

type SettingEntry struct {
	Key   string
	Value string
	Group string
}

type SettingsScreen struct {
	width    int
	height   int
	settings []SettingEntry
}

func NewSettingsScreen() *SettingsScreen {
	return &SettingsScreen{
		width:  80,
		height: 24,
	}
}

func (s *SettingsScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *SettingsScreen) SetSettings(settings []SettingEntry) {
	s.settings = settings
}

func (s *SettingsScreen) grouped() map[string][]SettingEntry {
	groups := make(map[string][]SettingEntry)
	for _, set := range s.settings {
		groups[set.Group] = append(groups[set.Group], set)
	}
	return groups
}

func (s *SettingsScreen) View() string {
	header := styles.H1.Render("Settings")
	sep := styles.SeparatorLine(s.width)

	var lines []string
	for group, entries := range s.grouped() {
		lines = append(lines, settingGroup.Render("── "+group+" ──"))
		for _, e := range entries {
			lines = append(lines, fmt.Sprintf("  %s  %s", settingLabel.Render(e.Key+":"), settingValue.Render(e.Value)))
		}
		lines = append(lines, "")
	}

	if len(lines) == 0 {
		return styles.Content(s.width, fmt.Sprintf("%s\n%s\n%s", header, sep,
			lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("No settings available.")))
	}

	return styles.Content(s.width, fmt.Sprintf("%s\n%s\n%s", header, sep, strings.Join(lines, "\n")))
}
