package dialogs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type ProviderListItem struct {
	Name      string
	BaseURL   string
	Status    string
	IsDefault bool
	Latency   int64
}

type ProviderListDialog struct {
	providers []ProviderListItem
	scroll    int
	screenW   int
	screenH   int
	active    bool
}

func NewProviderListDialog(providers []ProviderListItem) *ProviderListDialog {
	return &ProviderListDialog{
		providers: providers,
		active:    true,
	}
}

func (d *ProviderListDialog) SetSize(w, h int) {
	d.screenW = w
	d.screenH = h
}

func (d *ProviderListDialog) Active() bool {
	return d.active
}

func (d *ProviderListDialog) Init() tea.Cmd {
	return nil
}

func (d *ProviderListDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if d.scroll < len(d.providers)-1 {
				d.scroll++
			}
		}
	}
	return d, nil
}

func (d *ProviderListDialog) View() tea.View {
	boxW := 60
	if d.screenW > 0 && boxW > d.screenW-4 {
		boxW = d.screenW - 4
	}
	contentW := boxW - 4

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Width(boxW)

	var lines []string
	lines = append(lines, dlgTitle.Render("Configured Providers"))
	lines = append(lines, "")

	if len(d.providers) == 0 {
		lines = append(lines, "  No providers configured.")
		lines = append(lines, "")
		lines = append(lines, dlgHint.Render("  Use /provider add to add one."))
	} else {
		visible := d.providers
		start := d.scroll
		if start > len(visible) {
			start = 0
		}
		visible = visible[start:]

		for i, p := range visible {
			if i > 0 {
				lines = append(lines, dlgSep.Render(strings.Repeat("─", contentW)))
			}

			activeMark := "  "
			if p.IsDefault {
				activeMark = "> "
			}

			nameLine := activeMark + dlgLabel.Render(p.Name)
			if p.IsDefault {
				nameLine += "  " + dlgStatusOK.Render("(active)")
			}
			lines = append(lines, nameLine)

			url := p.BaseURL
			if len(url) > contentW-6 {
				url = url[:contentW-9] + "..."
			}
			lines = append(lines, "    "+dlgValue.Render(url))

			statusColor := dlgStatusWarn
			switch p.Status {
			case "connected":
				statusColor = dlgStatusOK
			case "auth_failed", "offline":
				statusColor = dlgStatusFail
			}
			statusStr := "Status: " + p.Status
			if p.Status == "connected" && p.Latency > 0 {
				statusStr += fmt.Sprintf("  %dms", p.Latency)
			}
			lines = append(lines, "    "+statusColor.Render(statusStr))
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
