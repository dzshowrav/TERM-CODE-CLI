package dialogs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type ModelSelectItem struct {
	ModelID      string
	DisplayName  string
	ProviderName string
	Category     string
	IsFavorite   bool
}

type ModelSelectDialog struct {
	models   []ModelSelectItem
	selected int
	screenW  int
	screenH  int
	active   bool
	onSelect func(modelID string) tea.Msg
}

func NewModelSelectDialog(models []ModelSelectItem, onSelect func(modelID string) tea.Msg) *ModelSelectDialog {
	return &ModelSelectDialog{
		models:   models,
		active:   true,
		onSelect: onSelect,
	}
}

func (d *ModelSelectDialog) SetSize(w, h int) {
	d.screenW = w
	d.screenH = h
}

func (d *ModelSelectDialog) Active() bool {
	return d.active
}

func (d *ModelSelectDialog) Init() tea.Cmd {
	return nil
}

func (d *ModelSelectDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if len(d.models) == 0 {
				d.active = false
				return d, nil
			}
			d.active = false
			m := d.models[d.selected]
			if d.onSelect != nil {
				return d, func() tea.Msg {
					return d.onSelect(m.ModelID)
				}
			}
		case "esc":
			d.active = false
		case "up", "k":
			if d.selected > 0 {
				d.selected--
			}
		case "down", "j":
			if d.selected < len(d.models)-1 {
				d.selected++
			}
		}
	}
	return d, nil
}

func (d *ModelSelectDialog) View() tea.View {
	boxW := d.screenW - 8
	if boxW < 56 {
		boxW = 56
	}
	if d.screenW > 0 && boxW > d.screenW-4 {
		boxW = d.screenW - 4
	}
	contentW := boxW - 4

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Width(boxW)

	var lines []string
	lines = append(lines, dlgTitle.Render("Select Model"))
	lines = append(lines, "")

	if len(d.models) == 0 {
		lines = append(lines, "  No models available.")
		lines = append(lines, "")
		lines = append(lines, dlgHint.Render("  Use /provider sync or /addmodel first."))
	} else {
		for i, m := range d.models {
			prefix := "  "
			itemStyle := dlgValue
			if i == d.selected {
				prefix = "> "
				itemStyle = dlgCursor
			}

			label := m.DisplayName
			if m.IsFavorite {
				label += dlgStatusOK.Render(" *")
			}

			catColor := dlgHint
			switch m.Category {
			case "coding":
				catColor = dlgStatusOK
			case "reasoning":
				catColor = dlgCursor
			case "vision":
				catColor = dlgStatusWarn
			}

			provW := 20
			provDisplay := m.ProviderName
			if len(provDisplay) > provW {
				provDisplay = provDisplay[:provW-1] + "…"
			}

			nameW := contentW - 22 - provW
			nameDisplay := label
			if len(nameDisplay) > nameW {
				nameDisplay = nameDisplay[:nameW-1] + "…"
			}

			lines = append(lines, fmt.Sprintf("%s%s %s %s",
				prefix,
				itemStyle.Render(fmt.Sprintf("%-*s", nameW, nameDisplay)),
				catColor.Render(fmt.Sprintf("%-8s", m.Category)),
				dlgHint.Render(provDisplay),
			))
		}
	}

	lines = append(lines, "")
	lines = append(lines, dlgHint.Render(fmt.Sprintf("%s%s%s",
		strings.Repeat(" ", (contentW-31)/2),
		"Arrows: Navigate  •  Enter: Select",
		strings.Repeat(" ", (contentW-31)/2),
	)))

	content := strings.Join(lines, "\n")
	box := style.Render(content)

	if d.screenW > 0 && d.screenH > 0 {
		return tea.NewView(lipgloss.Place(d.screenW, d.screenH,
			lipgloss.Center, lipgloss.Center,
			box))
	}
	return tea.NewView(box)
}
