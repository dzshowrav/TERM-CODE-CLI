package dialogs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

type ProviderSyncDialog struct {
	providerName string
	result       string
	err          string
	screenW      int
	screenH      int
	active       bool
}

func NewProviderSyncDialog(providerName string) *ProviderSyncDialog {
	return &ProviderSyncDialog{
		providerName: providerName,
		active:       true,
	}
}

func (d *ProviderSyncDialog) SetResult(n int) {
	d.result = fmt.Sprintf("Synced %d models from '%s'.", n, d.providerName)
}

func (d *ProviderSyncDialog) SetError(err string) {
	d.err = err
}

func (d *ProviderSyncDialog) SetSize(w, h int) {
	d.screenW = w
	d.screenH = h
}

func (d *ProviderSyncDialog) Active() bool {
	return d.active
}

func (d *ProviderSyncDialog) Init() tea.Cmd {
	return nil
}

func (d *ProviderSyncDialog) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "enter", "q":
			d.active = false
		}
	}
	return d, nil
}

func (d *ProviderSyncDialog) View() tea.View {
	boxW := 56
	if d.screenW > 0 && boxW > d.screenW-4 {
		boxW = d.screenW - 4
	}

	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Width(boxW)

	var lines []string
	lines = append(lines, dlgTitle.Render("Sync Models"))
	lines = append(lines, "")

	if d.err != "" {
		lines = append(lines, dlgStatusFail.Render("  Error: "+d.err))
	} else if d.result != "" {
		lines = append(lines, dlgStatusOK.Render("  "+d.result))
	} else {
		lines = append(lines, dlgStatusWarn.Render("  Syncing models from '"+d.providerName+"'..."))
	}

	lines = append(lines, "")
	lines = append(lines, dlgHint.Render(fmt.Sprintf("%s%s%s",
		strings.Repeat(" ", (boxW-30)/2),
		"ESC: Close",
		strings.Repeat(" ", (boxW-30)/2),
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
