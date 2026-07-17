package dialogs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type providerField int

const (
	fieldName providerField = iota
	fieldURL
	fieldKey
	fieldDesc
	fieldTestBtn
	fieldSaveBtn
	fieldCount
)

type ProviderFormResult struct {
	Name        string
	BaseURL     string
	APIKey      string
	Description string
}

type ProviderAddDialog struct {
	width    int
	height   int
	screenW  int
	screenH  int
	active   bool
	focus    providerField
	fields   map[providerField]string
	onAdd    func(ProviderFormResult) tea.Msg
	onTest   func(ProviderFormResult) tea.Msg
	testing  bool
	status   string
	statusOK bool
}

func (d *ProviderAddDialog) SetSize(w, h int) {
	d.screenW = w
	d.screenH = h
}

var (
	dlgBorder      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("39"))
	dlgTitle       = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39")).Padding(0, 2)
	dlgSep         = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	dlgLabel       = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Bold(true)
	dlgValue       = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	dlgPlaceholder = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Italic(true)
	dlgCursor      = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	dlgHint        = lipgloss.NewStyle().Foreground(lipgloss.Color("238"))
	dlgActiveBtn   = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	dlgInactiveBtn = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	dlgStatusOK    = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Bold(true)
	dlgStatusFail  = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	dlgStatusWarn  = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
)

func NewProviderAddDialog(onAdd func(ProviderFormResult) tea.Msg, onTest func(ProviderFormResult) tea.Msg) *ProviderAddDialog {
	return &ProviderAddDialog{
		width:  56,
		height: 22,
		active: true,
		focus:  fieldName,
		fields: map[providerField]string{
			fieldName: "",
			fieldURL:  "",
			fieldKey:  "",
			fieldDesc: "",
		},
		onAdd:  onAdd,
		onTest: onTest,
	}
}

func (d *ProviderAddDialog) Active() bool {
	return d.active
}

func (d *ProviderAddDialog) Init() tea.Cmd {
	return nil
}

type ProviderTestResultMsg struct {
	Success bool
	Message string
}

func (d *ProviderAddDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ProviderTestResultMsg:
		d.testing = false
		d.status = msg.Message
		d.statusOK = msg.Success
		return d, nil

	case tea.PasteMsg:
		if f := d.currentField(); f != nil {
			d.fields[*f] += msg.String()
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			d.active = false
			return d, nil

		case "tab", "down":
			d.focus = (d.focus + 1) % fieldCount

		case "shift+tab", "up":
			d.focus--
			if d.focus < 0 {
				d.focus = fieldCount - 1
			}

		case "enter":
			switch d.focus {
			case fieldSaveBtn:
				d.active = false
				if d.onAdd != nil {
					return d, func() tea.Msg {
						return d.onAdd(ProviderFormResult{
							Name:        d.fields[fieldName],
							BaseURL:     d.fields[fieldURL],
							APIKey:      d.fields[fieldKey],
							Description: d.fields[fieldDesc],
						})
					}
				}

			case fieldTestBtn:
				if d.onTest != nil {
					d.testing = true
					d.status = ""
					return d, func() tea.Msg {
						return d.onTest(ProviderFormResult{
							Name:        d.fields[fieldName],
							BaseURL:     d.fields[fieldURL],
							APIKey:      d.fields[fieldKey],
							Description: d.fields[fieldDesc],
						})
					}
				}
			}

		case "backspace":
			if f := d.currentField(); f != nil && len(d.fields[*f]) > 0 {
				s := d.fields[*f]
				d.fields[*f] = s[:len(s)-1]
			}

		default:
			if msg.String() == "space" {
				if f := d.currentField(); f != nil {
					d.fields[*f] += " "
				}
			} else if len(msg.String()) == 1 && msg.String()[0] >= 32 {
				if f := d.currentField(); f != nil {
					d.fields[*f] += msg.String()
				}
			}
		}
	}

	return d, nil
}

func (d *ProviderAddDialog) currentField() *providerField {
	if d.focus >= fieldName && d.focus <= fieldDesc {
		return &d.focus
	}
	return nil
}

func (d *ProviderAddDialog) View() tea.View {
	boxW := d.width
	if boxW < 50 {
		boxW = 50
	}
	if d.screenW > 0 && boxW > d.screenW-4 {
		boxW = d.screenW - 4
	}

	contentW := boxW - 4

	var lines []string

	lines = append(lines, dlgTitle.Render("Add OpenAI-Compatible Provider"))
	lines = append(lines, "")

	sep := dlgSep.Render(strings.Repeat("─", contentW))

	fields := []struct {
		label       string
		value       string
		placeholder string
		password    bool
		focus       providerField
	}{
		{"Name", d.fields[fieldName], "e.g. OpenCode Zen", false, fieldName},
		{"URL", d.fields[fieldURL], "https://api.openai.com", false, fieldURL},
		{"Key", d.fields[fieldKey], "sk-...", true, fieldKey},
		{"Desc", d.fields[fieldDesc], "Optional notes", false, fieldDesc},
	}

	for i, f := range fields {
		if i > 0 {
			lines = append(lines, sep)
		}

		label := dlgLabel.Render(f.label + ":")
		var val string
		if f.value == "" {
			val = dlgPlaceholder.Render(f.placeholder)
		} else {
			display := f.value
			if f.password {
				display = strings.Repeat("•", len(display))
			}
			val = dlgValue.Render(display)
		}

		cursor := ""
		if d.focus == f.focus {
			cursor = dlgCursor.Render("█")
			if f.value == "" {
				val = dlgPlaceholder.Render(f.placeholder)
			}
		}

		line := fmt.Sprintf("  %s %s%s", label, val, cursor)
		lines = append(lines, line)
	}

	lines = append(lines, sep)

	if d.status != "" {
		var statusLine string
		if d.testing {
			statusLine = dlgStatusWarn.Render("  Testing...")
		} else if d.statusOK {
			statusLine = dlgStatusOK.Render("  " + d.status)
		} else {
			statusLine = dlgStatusFail.Render("  " + d.status)
		}
		lines = append(lines, statusLine)
		lines = append(lines, "")
	} else {
		lines = append(lines, "")
		lines = append(lines, "")
	}

	testBtn := dlgInactiveBtn.Render("[ Test Connection ]")
	saveBtn := dlgInactiveBtn.Render("[ Save Provider ]")
	if d.focus == fieldTestBtn {
		testBtn = dlgActiveBtn.Render("[ Test Connection ]")
	}
	if d.focus == fieldSaveBtn {
		saveBtn = dlgActiveBtn.Render("[ Save Provider ]")
	}

	btnPad := contentW - 39
	if btnPad < 2 {
		btnPad = 2
	}
	btnLine := fmt.Sprintf("  %s%s%s", testBtn, strings.Repeat(" ", btnPad), saveBtn)
	lines = append(lines, btnLine)

	lines = append(lines, "")
	lines = append(lines, dlgHint.Render(fmt.Sprintf("%s%s%s",
		strings.Repeat(" ", (contentW-32)/2),
		"Arrows: Navigate  •  ESC: Cancel",
		strings.Repeat(" ", (contentW-32)/2),
	)))
	lines = append(lines, "")

	box := dlgBorder.Width(boxW).Render(strings.Join(lines, "\n"))

	if d.screenW > 0 && d.screenH > 0 {
		return tea.NewView(lipgloss.Place(d.screenW, d.screenH,
			lipgloss.Center, lipgloss.Center,
			box))
	}

	return tea.NewView(box)
}
