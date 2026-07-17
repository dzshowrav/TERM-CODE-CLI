package dialogs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type ModelFormResult struct {
	ModelID      string
	DisplayName  string
	ProviderID   string
	ProviderName string
	Category     string
	MaxContext   int
}

type modelAddField int

const (
	modelFieldModelID modelAddField = iota
	modelFieldDisplayName
	modelFieldProvider
	modelFieldCategory
	modelFieldContext
	modelFieldSaveBtn
	modelFieldCount
)

type ModelAddDialog struct {
	width     int
	height    int
	screenW   int
	screenH   int
	active    bool
	focus     modelAddField
	fields    map[modelAddField]string
	providers []ProviderSelectItem
	provIdx   int
	onAdd     func(ModelFormResult) tea.Msg
}

func NewModelAddDialog(providers []ProviderSelectItem, onAdd func(ModelFormResult) tea.Msg) *ModelAddDialog {
	return &ModelAddDialog{
		width:     56,
		height:    20,
		active:    true,
		focus:     modelFieldModelID,
		providers: providers,
		fields: map[modelAddField]string{
			modelFieldModelID:     "",
			modelFieldDisplayName: "",
			modelFieldProvider:    "",
			modelFieldCategory:    "general",
			modelFieldContext:     "4096",
		},
		onAdd: onAdd,
	}
}

func (d *ModelAddDialog) SetSize(w, h int) {
	d.screenW = w
	d.screenH = h
}

func (d *ModelAddDialog) Active() bool {
	return d.active
}

func (d *ModelAddDialog) Init() tea.Cmd {
	return nil
}

var categories = []string{"general", "coding", "reasoning", "vision", "embedding", "audio", "custom"}

func (d *ModelAddDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
			d.focus = (d.focus + 1) % modelFieldCount

		case "shift+tab", "up":
			d.focus--
			if d.focus < 0 {
				d.focus = modelFieldCount - 1
			}

		case "enter":
			if d.focus == modelFieldSaveBtn {
				d.active = false
				if d.onAdd != nil {
					ctx := d.fields[modelFieldContext]
					maxCtx := 4096
					if ctx != "" {
						fmt.Sscanf(ctx, "%d", &maxCtx)
					}
					provID := ""
					provName := ""
					if d.provIdx >= 0 && d.provIdx < len(d.providers) {
						provID = d.providers[d.provIdx].ID
						provName = d.providers[d.provIdx].Name
					}
					return d, func() tea.Msg {
						return d.onAdd(ModelFormResult{
							ModelID:      d.fields[modelFieldModelID],
							DisplayName:  d.fields[modelFieldDisplayName],
							ProviderID:   provID,
							ProviderName: provName,
							Category:     d.fields[modelFieldCategory],
							MaxContext:   maxCtx,
						})
					}
				}
			}

		case "left":
			if d.focus == modelFieldCategory {
				idx := indexOf(categories, d.fields[modelFieldCategory])
				if idx > 0 {
					d.fields[modelFieldCategory] = categories[idx-1]
				}
			}
			if d.focus == modelFieldProvider && len(d.providers) > 0 {
				if d.provIdx > 0 {
					d.provIdx--
				}
			}

		case "right":
			if d.focus == modelFieldCategory {
				idx := indexOf(categories, d.fields[modelFieldCategory])
				if idx >= 0 && idx < len(categories)-1 {
					d.fields[modelFieldCategory] = categories[idx+1]
				}
			}
			if d.focus == modelFieldProvider && len(d.providers) > 0 {
				if d.provIdx < len(d.providers)-1 {
					d.provIdx++
				}
			}

		case "backspace":
			if f := d.currentField(); f != nil {
				s := d.fields[*f]
				if len(s) > 0 {
					d.fields[*f] = s[:len(s)-1]
				}
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

func (d *ModelAddDialog) currentField() *modelAddField {
	if d.focus < modelFieldSaveBtn {
		return &d.focus
	}
	return nil
}

func indexOf(slice []string, val string) int {
	for i, s := range slice {
		if s == val {
			return i
		}
	}
	return -1
}

func (d *ModelAddDialog) View() tea.View {
	boxW := d.width
	if boxW < 50 {
		boxW = 50
	}
	if d.screenW > 0 && boxW > d.screenW-4 {
		boxW = d.screenW - 4
	}
	contentW := boxW - 4

	var lines []string
	lines = append(lines, dlgTitle.Render("Add Custom Model"))
	lines = append(lines, "")

	sep := dlgSep.Render(strings.Repeat("─", contentW))

	editFields := []struct {
		label       string
		value       string
		placeholder string
		focus       modelAddField
	}{
		{"Model ID", d.fields[modelFieldModelID], "e.g. gpt-4o-mini", modelFieldModelID},
		{"Name", d.fields[modelFieldDisplayName], "e.g. GPT-4o Mini", modelFieldDisplayName},
	}

	for i, f := range editFields {
		if i > 0 {
			lines = append(lines, sep)
		}
		label := dlgLabel.Render(f.label + ":")
		val := dlgValue.Render(f.value)
		if f.value == "" {
			val = dlgPlaceholder.Render(f.placeholder)
		}
		cursor := ""
		if d.focus == f.focus {
			cursor = dlgCursor.Render("█")
			if f.value == "" {
				val = dlgPlaceholder.Render(f.placeholder)
			}
		}
		lines = append(lines, fmt.Sprintf("  %s %s%s", label, val, cursor))
	}

	lines = append(lines, sep)

	provLabel := dlgLabel.Render("Provider:")
	provVal := dlgValue.Render("No providers")
	if len(d.providers) > 0 {
		p := d.providers[d.provIdx]
		pName := p.Name
		if p.IsDefault {
			pName += " (active)"
		}
		provVal = dlgValue.Render(pName)
	}
	provCursor := ""
	if d.focus == modelFieldProvider {
		provCursor = dlgCursor.Render(" ◄ ►")
	}
	lines = append(lines, fmt.Sprintf("  %s %s%s", provLabel, provVal, provCursor))

	lines = append(lines, sep)

	catLabel := dlgLabel.Render("Category:")
	catVal := dlgValue.Render(d.fields[modelFieldCategory])
	catCursor := ""
	if d.focus == modelFieldCategory {
		catCursor = dlgCursor.Render(" ◄ ►")
	}
	lines = append(lines, fmt.Sprintf("  %s %s%s", catLabel, catVal, catCursor))

	lines = append(lines, sep)

	ctxLabel := dlgLabel.Render("Max Context:")
	ctxVal := dlgValue.Render(d.fields[modelFieldContext])
	if d.fields[modelFieldContext] == "" {
		ctxVal = dlgPlaceholder.Render("4096")
	}
	ctxCursor := ""
	if d.focus == modelFieldContext {
		ctxCursor = dlgCursor.Render("█")
		if d.fields[modelFieldContext] == "" {
			ctxVal = dlgPlaceholder.Render("4096")
		}
	}
	lines = append(lines, fmt.Sprintf("  %s %s%s", ctxLabel, ctxVal, ctxCursor))

	lines = append(lines, sep)
	lines = append(lines, "")

	saveBtn := dlgInactiveBtn.Render("[ Save Model ]")
	if d.focus == modelFieldSaveBtn {
		saveBtn = dlgActiveBtn.Render("[ Save Model ]")
	}
	btnPad := contentW - 16
	if btnPad < 2 {
		btnPad = 2
	}
	lines = append(lines, fmt.Sprintf("  %s%s", strings.Repeat(" ", btnPad/2), saveBtn))

	lines = append(lines, "")
	lines = append(lines, dlgHint.Render(fmt.Sprintf("%s%s%s",
		strings.Repeat(" ", (contentW-40)/2),
		"Tab: Navigate  ◄ ►: Cycle  •  ESC: Cancel",
		strings.Repeat(" ", (contentW-40)/2),
	)))

	box := dlgBorder.Width(boxW).Render(strings.Join(lines, "\n"))

	if d.screenW > 0 && d.screenH > 0 {
		return tea.NewView(lipgloss.Place(d.screenW, d.screenH,
			lipgloss.Center, lipgloss.Center,
			box))
	}
	return tea.NewView(box)
}
