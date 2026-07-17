package dialogs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type ModelListItem struct {
	DisplayName  string
	ModelID      string
	Category     string
	ProviderName string
	MaxContext   int
	IsFavorite   bool
	IsLocal      bool
}

type ModelListDialog struct {
	models  []ModelListItem
	scroll  int
	screenW int
	screenH int
	active  bool
}

func NewModelListDialog(models []ModelListItem) *ModelListDialog {
	return &ModelListDialog{
		models: models,
		active: true,
	}
}

func (d *ModelListDialog) SetSize(w, h int) {
	d.screenW = w
	d.screenH = h
}

func (d *ModelListDialog) Active() bool {
	return d.active
}

func (d *ModelListDialog) Init() tea.Cmd {
	return nil
}

func (d *ModelListDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "enter", "q":
			d.active = false
		case "up", "k":
			if d.scroll > 0 {
				d.scroll--
			}
		case "down", "j":
			if d.scroll < len(d.models)-1 {
				d.scroll++
			}
		}
	}
	return d, nil
}

func (d *ModelListDialog) View() tea.View {
	boxW := d.screenW - 8
	if boxW < 60 {
		boxW = 60
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
	lines = append(lines, dlgTitle.Render(fmt.Sprintf("Models (%d)", len(d.models))))
	lines = append(lines, "")

	if len(d.models) == 0 {
		lines = append(lines, "  No models configured.")
		lines = append(lines, "")
		lines = append(lines, dlgHint.Render("  Use /provider sync or /addmodel to add models."))
	} else {
		modelPad := contentW - 4

		header := fmt.Sprintf("  %-*s  %-6s  %s", modelPad-12, "Model", "Tokens", "Provider")
		lines = append(lines, dlgHint.Render(header))
		lines = append(lines, dlgSep.Render(strings.Repeat("─", contentW)))

		start := d.scroll
		visible := d.models
		if start > len(visible) {
			start = 0
		}
		visible = visible[start:]

		for _, m := range visible {
			fav := ""
			if m.IsFavorite {
				fav = " *"
			}

			nameDisplay := m.DisplayName + fav
			nameW := modelPad - 12 - 2
			if len(nameDisplay) > nameW {
				nameDisplay = nameDisplay[:nameW-1] + "…"
			}

			ctxStr := fmt.Sprintf("%dK", m.MaxContext/1000)
			if m.MaxContext >= 1000000 {
				ctxStr = fmt.Sprintf("%dM", m.MaxContext/1000000)
			}

			provDisplay := m.ProviderName
			provW := 20
			if len(provDisplay) > provW {
				provDisplay = provDisplay[:provW-1] + "…"
			}

			catColor := dlgValue
			switch m.Category {
			case "coding":
				catColor = dlgStatusOK
			case "reasoning":
				catColor = dlgCursor
			case "vision":
				catColor = dlgStatusWarn
			}

			lines = append(lines, fmt.Sprintf("  %s  %s  %s  %s",
				dlgValue.Render(fmt.Sprintf("%-*s", nameW+2, nameDisplay)),
				catColor.Render(fmt.Sprintf("%-6s", ctxStr)),
				dlgHint.Render(m.Category),
				dlgHint.Render(provDisplay),
			))
		}
	}

	lines = append(lines, "")
	lines = append(lines, dlgHint.Render(fmt.Sprintf("%s%s%s",
		strings.Repeat(" ", (contentW-27)/2),
		"Arrows: Scroll  •  ESC: Close",
		strings.Repeat(" ", (contentW-27)/2),
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
