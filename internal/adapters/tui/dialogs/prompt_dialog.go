package dialogs

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var (
	promptLabel = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	promptInput = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
)

type PromptResult string

type PromptDialog struct {
	title    string
	label    string
	value    string
	cursor   int
	active   bool
	width    int
	onResult func(PromptResult) tea.Msg
}

func NewPromptDialog(title, label string, onResult func(PromptResult) tea.Msg) *PromptDialog {
	return &PromptDialog{
		title:    title,
		label:    label,
		active:   true,
		width:    50,
		onResult: onResult,
	}
}

func (d *PromptDialog) SetWidth(w int) {
	d.width = w
}

func (d *PromptDialog) Init() tea.Cmd {
	return nil
}

func (d *PromptDialog) Update(msg tea.Msg) (*PromptDialog, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			d.active = false
			if d.onResult != nil {
				return d, func() tea.Msg {
					return d.onResult(PromptResult(d.value))
				}
			}
		case "esc":
			d.active = false
			if d.onResult != nil {
				return d, func() tea.Msg {
					return d.onResult(PromptResult(""))
				}
			}
		case "backspace":
			if d.cursor > 0 {
				d.value = d.value[:d.cursor-1] + d.value[d.cursor:]
				d.cursor--
			}
		case "delete":
			if d.cursor < len(d.value) {
				d.value = d.value[:d.cursor] + d.value[d.cursor+1:]
			}
		case "left":
			if d.cursor > 0 {
				d.cursor--
			}
		case "right":
			if d.cursor < len(d.value) {
				d.cursor++
			}
		default:
			if len(msg.String()) == 1 && msg.String()[0] >= 32 {
				r := msg.String()
				d.value = d.value[:d.cursor] + r + d.value[d.cursor:]
				d.cursor++
			}
		}
	}
	return d, nil
}

func (d *PromptDialog) View() string {
	display := d.value
	cursor := ""
	if d.active {
		cursor = "█"
	}

	content := fmt.Sprintf(
		"%s\n\n%s\n> %s",
		dialogTitle.Render(d.title),
		promptLabel.Render(d.label),
		promptInput.Render(display[:d.cursor]+cursor+display[d.cursor:]),
	)

	boxW := d.width
	if boxW < 40 {
		boxW = 40
	}
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Padding(1, 2).
		Width(boxW)
	return style.Render(content)
}

func (d *PromptDialog) Active() bool {
	return d.active
}

func (d *PromptDialog) SetActive(a bool) {
	d.active = a
}

func (d *PromptDialog) Focused() bool {
	return d.active
}

func (d *PromptDialog) SetFocused(f bool) {
	d.active = f
}
