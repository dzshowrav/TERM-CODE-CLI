package screens

import (
	"fmt"
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type ModelAddScreen struct {
	width       int
	height      int
	modelID     string
	displayName string
	provider    string
	contextSize string
	maxOutput   string
	focusField  int
	done        bool
	result      string
	onSubmit    func(id, display, provider string, ctxSize, maxOut int) string
}

func NewModelAddScreen(onSubmit func(id, display, provider string, ctxSize, maxOut int) string) *ModelAddScreen {
	return &ModelAddScreen{
		width:       80,
		height:      24,
		contextSize: "4096",
		maxOutput:   "4096",
		focusField:  0,
		onSubmit:    onSubmit,
	}
}

func (s *ModelAddScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ModelAddScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.PasteMsg:
		paste := msg.String()
		switch s.focusField {
		case 0:
			s.modelID += paste
		case 1:
			s.displayName += paste
		case 2:
			s.provider += paste
		case 3:
			s.contextSize += paste
		case 4:
			s.maxOutput += paste
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "enter":
			if s.focusField == 5 {
				ctx, _ := strconv.Atoi(s.contextSize)
				maxOut, _ := strconv.Atoi(s.maxOutput)
				if s.onSubmit != nil {
					s.result = s.onSubmit(s.modelID, s.displayName, s.provider, ctx, maxOut)
				}
				s.done = true
			} else if s.focusField == 6 {
				s.done = true
				s.result = ""
			}

		case "tab", "down":
			s.focusField = (s.focusField + 1) % 7
		case "shift+tab", "up":
			s.focusField--
			if s.focusField < 0 {
				s.focusField = 6
			}
		case "backspace":
			if s.focusField < 5 {
				s.deleteChar()
			}
		default:
			if s.focusField < 5 && len(msg.String()) == 1 {
				s.addChar(msg.String())
			}
		}
	}
	return s, nil
}

func (s *ModelAddScreen) addChar(ch string) {
	switch s.focusField {
	case 0:
		s.modelID += ch
	case 1:
		s.displayName += ch
	case 2:
		s.provider += ch
	case 3:
		s.contextSize += ch
	case 4:
		s.maxOutput += ch
	}
}

func (s *ModelAddScreen) deleteChar() {
	switch s.focusField {
	case 0:
		if len(s.modelID) > 0 {
			s.modelID = s.modelID[:len(s.modelID)-1]
		}
	case 1:
		if len(s.displayName) > 0 {
			s.displayName = s.displayName[:len(s.displayName)-1]
		}
	case 2:
		if len(s.provider) > 0 {
			s.provider = s.provider[:len(s.provider)-1]
		}
	case 3:
		if len(s.contextSize) > 0 {
			s.contextSize = s.contextSize[:len(s.contextSize)-1]
		}
	case 4:
		if len(s.maxOutput) > 0 {
			s.maxOutput = s.maxOutput[:len(s.maxOutput)-1]
		}
	}
}

func (s *ModelAddScreen) Done() bool     { return s.done }
func (s *ModelAddScreen) Result() string { return s.result }

func (s *ModelAddScreen) View() string {
	header := styles.H1.Render("Add Model")

	innerWidth := s.width - 2
	sep := styles.DialogSep(innerWidth)

	fields := []struct {
		label string
		value string
		hint  string
	}{
		{"Model ID", s.modelID, "e.g. gpt-4o"},
		{"Display Name", s.displayName, "e.g. GPT-4o"},
		{"Provider", s.provider, "e.g. openai"},
		{"Context Size", s.contextSize, "e.g. 4096"},
		{"Max Output", s.maxOutput, "e.g. 4096"},
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
		fieldLine := fmt.Sprintf(" %s:  %s", formLabelStyle.Render(f.label), formInputStyle.Render(val))
		lines = append(lines, fieldLine)
		lines = append(lines, sep)
	}

	saveBtn := formBtnNormal.Render("[ Save ]")
	cancelBtn := formBtnNormal.Render("[ Cancel ]")
	if s.focusField == 5 {
		saveBtn = formBtnActive.Render("[ Save ]")
	}
	if s.focusField == 6 {
		cancelBtn = formBtnActive.Render("[ Cancel ]")
	}
	btnLine := lipgloss.JoinHorizontal(lipgloss.Center, saveBtn, cancelBtn)
	btnLine = lipgloss.NewStyle().Width(innerWidth).Align(lipgloss.Center).Render(btnLine)

	lines = append(lines, "")
	lines = append(lines, btnLine)
	lines = append(lines, "")

	hintText := formHintStyle.Render("Arrows: Navigate \u2022 ESC: Cancel")
	hintLine := lipgloss.Place(innerWidth, 1, lipgloss.Center, lipgloss.Center, hintText)
	lines = append(lines, hintLine)
	lines = append(lines, "")

	return styles.DialogBox(s.width, strings.Join(lines, "\n"))
}

var _ = strings.Builder{}
