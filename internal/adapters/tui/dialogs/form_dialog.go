package dialogs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var (
	formLabelStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	formInputStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	formActiveStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	formInactiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

type FormField struct {
	Name     string
	Label    string
	Value    string
	Password bool
}

type FormResult map[string]string

type FormDialog struct {
	title    string
	fields   []FormField
	focus    int
	width    int
	active   bool
	onSubmit func(FormResult) tea.Msg
}

func NewFormDialog(title string, fields []FormField, onSubmit func(FormResult) tea.Msg) *FormDialog {
	return &FormDialog{
		title:    title,
		fields:   fields,
		active:   true,
		width:    50,
		onSubmit: onSubmit,
	}
}

func (d *FormDialog) SetWidth(w int) {
	d.width = w
}

func (d *FormDialog) Init() tea.Cmd {
	return nil
}

func (d *FormDialog) currentField() *FormField {
	if d.focus >= 0 && d.focus < len(d.fields) {
		return &d.fields[d.focus]
	}
	return nil
}

func (d *FormDialog) Update(msg tea.Msg) (*FormDialog, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab", "down":
			d.focus = (d.focus + 1) % (len(d.fields) + 1)
		case "shift+tab", "up":
			d.focus--
			if d.focus < 0 {
				d.focus = len(d.fields)
			}
		case "enter":
			if d.focus == len(d.fields) {
				d.active = false
				result := make(FormResult)
				for _, f := range d.fields {
					result[f.Name] = f.Value
				}
				if d.onSubmit != nil {
					return d, func() tea.Msg {
						return d.onSubmit(result)
					}
				}
			}
		case "esc":
			d.active = false
		case "backspace":
			if f := d.currentField(); f != nil && len(f.Value) > 0 {
				f.Value = f.Value[:len(f.Value)-1]
			}
		default:
			if f := d.currentField(); f != nil {
				if len(msg.String()) == 1 && msg.String()[0] >= 32 {
					f.Value += msg.String()
				}
			}
		}
	}
	return d, nil
}

func (d *FormDialog) View() string {
	boxW := d.width
	if boxW < 40 {
		boxW = 40
	}
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Padding(1, 2).
		Width(boxW)

	var lines []string
	lines = append(lines, dialogTitle.Render(d.title))
	lines = append(lines, "")

	for i, f := range d.fields {
		label := formLabelStyle.Render(f.Label + ":")
		val := f.Value
		if f.Password {
			val = strings.Repeat("*", len(val))
		}
		if val == "" {
			val = " "
		}

		input := formInputStyle.Render(val)
		if i == d.focus {
			input = formActiveStyle.Render(val + "█")
		}
		if f.Password && val != " " {
			input = formActiveStyle.Render(strings.Repeat("*", len(f.Value)) + "█")
		}

		lines = append(lines, fmt.Sprintf("  %s %s", label, input))
		lines = append(lines, "")
	}

	submitBtn := formInactiveStyle.Render("[ Submit ]")
	if d.focus == len(d.fields) {
		submitBtn = formActiveStyle.Render("[ Submit ]")
	}
	lines = append(lines, "  "+submitBtn)

	content := strings.Join(lines, "\n")
	return style.Render(content)
}

func (d *FormDialog) Active() bool {
	return d.active
}

func (d *FormDialog) SetActive(a bool) {
	d.active = a
}

func (d *FormDialog) Focused() bool {
	return d.active
}

func (d *FormDialog) SetFocused(f bool) {
	d.active = f
}

var _ = fmt.Sprintf
