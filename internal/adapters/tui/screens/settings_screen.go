package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
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
	done     bool
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

func (s *SettingsScreen) Done() bool {
	return s.done
}

func (s *SettingsScreen) Result() string {
	return ""
}

func (s *SettingsScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "enter":
			s.done = true
		}
	}
	return s, nil
}

func (s *SettingsScreen) grouped() map[string][]SettingEntry {
	groups := make(map[string][]SettingEntry)
	for _, set := range s.settings {
		groups[set.Group] = append(groups[set.Group], set)
	}
	return groups
}

func (s *SettingsScreen) View() string {
	innerW := s.width - 2
	sep := styles.DialogSep(innerW)

	var lines []string
	lines = append(lines, fmt.Sprintf("%-*s%s", innerW-4, "Settings", "esc"))
	lines = append(lines, sep)
	lines = append(lines, "")

	for group, entries := range s.grouped() {
		lines = append(lines, settingGroup.Render("── "+group+" ──"))
		for _, e := range entries {
			lines = append(lines, fmt.Sprintf("  %s  %s", settingLabel.Render(e.Key+":"), settingValue.Render(e.Value)))
		}
		lines = append(lines, "")
	}

	if len(s.settings) == 0 {
		lines = append(lines, lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("No settings available."))
		lines = append(lines, "")
	}

	hintText := "esc: Close"
	hintLine := lipgloss.NewStyle().Width(innerW).Align(lipgloss.Center).Render(styles.HintStyle.Render(hintText))
	lines = append(lines, hintLine)
	lines = append(lines, "")

	body := strings.Join(lines, "\n")
	return styles.DialogBox(s.width, body)
}
