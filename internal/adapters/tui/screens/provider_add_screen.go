package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	formFieldStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	formLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	formInputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	formHintStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
	formBtnNormal  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	formBtnActive  = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
)

type ProviderAddScreen struct {
	width      int
	height     int
	name       string
	baseURL    string
	apiKey     string
	desc       string
	focusField int
	done       bool
	result     string
	testStatus string
	onSubmit   func(name, baseURL, apiKey, desc string) string
	onTest     func(name, baseURL, apiKey, desc string) string
}

func NewProviderAddScreen(onSubmit func(name, baseURL, apiKey, desc string) string, onTest func(name, baseURL, apiKey, desc string) string) *ProviderAddScreen {
	return &ProviderAddScreen{
		width:      80,
		height:     24,
		focusField: 0,
		onSubmit:   onSubmit,
		onTest:     onTest,
	}
}

func (s *ProviderAddScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ProviderAddScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.PasteMsg:
		paste := msg.String()
		switch s.focusField {
		case 0:
			s.name += paste
		case 1:
			s.baseURL += paste
		case 2:
			s.apiKey += paste
		case 3:
			s.desc += paste
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "enter":
			if s.focusField < 4 {
				s.focusField++
			} else if s.focusField == 4 {
				if s.onTest != nil {
					s.testStatus = s.onTest(s.name, s.baseURL, s.apiKey, s.desc)
				}
			} else if s.focusField == 5 {
				if s.onSubmit != nil {
					s.result = s.onSubmit(s.name, s.baseURL, s.apiKey, s.desc)
				}
				s.done = true
			}

		case "tab", "down":
			s.focusField = (s.focusField + 1) % 6
		case "shift+tab", "up":
			if s.focusField == 0 {
				s.focusField = 5
			} else {
				s.focusField--
			}
		case "backspace":
			if s.focusField < 4 {
				s.deleteChar()
			}
		default:
			if s.focusField < 4 && len(msg.String()) == 1 {
				s.addChar(msg.String())
			}
		}
	}
	return s, nil
}

func (s *ProviderAddScreen) addChar(ch string) {
	switch s.focusField {
	case 0:
		s.name += ch
	case 1:
		s.baseURL += ch
	case 2:
		s.apiKey += ch
	case 3:
		s.desc += ch
	}
}

func (s *ProviderAddScreen) deleteChar() {
	switch s.focusField {
	case 0:
		if len(s.name) > 0 {
			s.name = s.name[:len(s.name)-1]
		}
	case 1:
		if len(s.baseURL) > 0 {
			s.baseURL = s.baseURL[:len(s.baseURL)-1]
		}
	case 2:
		if len(s.apiKey) > 0 {
			s.apiKey = s.apiKey[:len(s.apiKey)-1]
		}
	case 3:
		if len(s.desc) > 0 {
			s.desc = s.desc[:len(s.desc)-1]
		}
	}
}

func (s *ProviderAddScreen) Done() bool     { return s.done }
func (s *ProviderAddScreen) Result() string { return s.result }

func (s *ProviderAddScreen) View() string {
	header := styles.H1.Render("Add Provider")

	innerWidth := s.width - 2
	sep := styles.DialogSep(innerWidth)

	fields := []struct {
		label string
		value string
		hint  string
	}{
		{"Name", s.name, "e.g. OpenCode Zen"},
		{"URL", s.baseURL, "https://api.openai.com"},
		{"Key", maskKey(s.apiKey), "sk-..."},
		{"Desc", s.desc, "Optional notes"},
	}

	var lines []string
	lines = append(lines, "")
	lines = append(lines, header)
	lines = append(lines, sep)

	for i, f := range fields {
		val := f.value
		if val == "" && i != s.focusField {
			val = formHintStyle.Render(f.hint)
		}
		if i == s.focusField {
			val += "█"
		}
		maxVal := innerWidth - len(f.label) - 7
		if maxVal < 5 {
			maxVal = 5
		}
		if len([]rune(val)) > maxVal {
			val = string([]rune(val)[:maxVal]) + "..."
		}
		fieldLine := fmt.Sprintf(" %s:  %s", formLabelStyle.Render(f.label), formInputStyle.Render(val))
		lines = append(lines, fieldLine)
		lines = append(lines, sep)
	}

	testBtn := formBtnNormal.Render("[ Test Connection ]")
	saveBtn := formBtnNormal.Render("[ Save Provider ]")
	if s.focusField == 4 {
		testBtn = formBtnActive.Render("[ Test Connection ]")
	}
	if s.focusField == 5 {
		saveBtn = formBtnActive.Render("[ Save Provider ]")
	}
	btnLine := lipgloss.JoinHorizontal(lipgloss.Center, testBtn, saveBtn)
	btnLine = lipgloss.NewStyle().Width(innerWidth).Align(lipgloss.Center).Render(btnLine)

	lines = append(lines, "")
	if s.testStatus != "" {
		statusClr := lipgloss.Color("83")
		if strings.HasPrefix(s.testStatus, "Error") || strings.HasPrefix(s.testStatus, "error") {
			statusClr = lipgloss.Color("196")
		}
		statusLine := lipgloss.NewStyle().Foreground(statusClr).Render(" " + s.testStatus)
		lines = append(lines, statusLine)
		lines = append(lines, "")
	}
	lines = append(lines, btnLine)
	lines = append(lines, "")

	hintText := formHintStyle.Render("Arrows: Navigate \u2022 ESC: Cancel")
	hintLine := lipgloss.Place(innerWidth, 1, lipgloss.Center, lipgloss.Center, hintText)
	lines = append(lines, hintLine)
	lines = append(lines, "")

	return styles.DialogBox(s.width, strings.Join(lines, "\n"))
}

func maskKey(key string) string {
	if len(key) <= 8 {
		return strings.Repeat("*", len(key))
	}
	return key[:4] + strings.Repeat("*", len(key)-8) + key[len(key)-4:]
}
