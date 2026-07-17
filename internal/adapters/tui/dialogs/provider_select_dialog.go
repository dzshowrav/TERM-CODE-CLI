package dialogs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type ProviderSelectItem struct {
	ID        string
	Name      string
	BaseURL   string
	IsDefault bool
}

type ProviderSelectDialog struct {
	providers []ProviderSelectItem
	selected  int
	screenW   int
	screenH   int
	active    bool
	onSelect  func(id string) tea.Msg
}

func NewProviderSelectDialog(providers []ProviderSelectItem, onSelect func(id string) tea.Msg) *ProviderSelectDialog {
	return &ProviderSelectDialog{
		providers: providers,
		active:    true,
		onSelect:  onSelect,
	}
}

func (d *ProviderSelectDialog) SetSize(w, h int) {
	d.screenW = w
	d.screenH = h
}

func (d *ProviderSelectDialog) Active() bool {
	return d.active
}

func (d *ProviderSelectDialog) Init() tea.Cmd {
	return nil
}

type providerSelectResultMsg struct {
	id string
}

func (d *ProviderSelectDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if len(d.providers) == 0 {
				d.active = false
				return d, nil
			}
			d.active = false
			p := d.providers[d.selected]
			if d.onSelect != nil {
				return d, func() tea.Msg {
					return d.onSelect(p.ID)
				}
			}
		case "esc":
			d.active = false
		case "up", "k":
			if d.selected > 0 {
				d.selected--
			}
		case "down", "j":
			if d.selected < len(d.providers)-1 {
				d.selected++
			}
		}
	}
	return d, nil
}

func (d *ProviderSelectDialog) View() tea.View {
	boxW := 56
	if d.screenW > 0 && boxW > d.screenW-4 {
		boxW = d.screenW - 4
	}
	contentW := boxW - 4

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Width(boxW)

	var lines []string
	lines = append(lines, dlgTitle.Render("Select Provider"))
	lines = append(lines, "")

	if len(d.providers) == 0 {
		lines = append(lines, "  No providers configured.")
		lines = append(lines, "")
		lines = append(lines, dlgHint.Render("  Use /provider add to add one."))
	} else {
		for i, p := range d.providers {
			prefix := "  "
			itemStyle := dlgValue
			if i == d.selected {
				prefix = "> "
				itemStyle = dlgCursor
			}

			label := p.Name
			if p.IsDefault {
				label += "  " + dlgStatusOK.Render("(active)")
			}

			url := p.BaseURL
			maxURL := contentW - 8
			if len(url) > maxURL {
				url = url[:maxURL-3] + "..."
			}

			lines = append(lines, prefix+itemStyle.Render(label))
			lines = append(lines, "    "+dlgHint.Render(url))

			if i < len(d.providers)-1 {
				lines = append(lines, "")
			}
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
