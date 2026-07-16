package dialogs

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var (
	selItemStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	selActiveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
)

type SelectItem struct {
	ID    string
	Label string
}

type SelectResult string

type SelectDialog struct {
	title    string
	items    []SelectItem
	selected int
	width    int
	height   int
	active   bool
	onSelect func(SelectResult) tea.Msg
}

func NewSelectDialog(title string, items []SelectItem, onSelect func(SelectResult) tea.Msg) *SelectDialog {
	return &SelectDialog{
		title:    title,
		items:    items,
		width:    50,
		height:   10,
		active:   true,
		onSelect: onSelect,
	}
}

func (d *SelectDialog) SetWidth(w int) {
	d.width = w
}

func (d *SelectDialog) Init() tea.Cmd {
	return nil
}

func (d *SelectDialog) Update(msg tea.Msg) (*SelectDialog, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			d.active = false
			if d.selected < len(d.items) && d.onSelect != nil {
				item := d.items[d.selected]
				return d, func() tea.Msg {
					return d.onSelect(SelectResult(item.ID))
				}
			}
		case "esc":
			d.active = false
			if d.onSelect != nil {
				return d, func() tea.Msg {
					return d.onSelect(SelectResult(""))
				}
			}
		case "up", "ctrl+p":
			if d.selected > 0 {
				d.selected--
			}
		case "down", "ctrl+n":
			if d.selected < len(d.items)-1 {
				d.selected++
			}
		}
	}
	return d, nil
}

func (d *SelectDialog) View() string {
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

	for i, item := range d.items {
		prefix := "  "
		itemStyle := selItemStyle
		if i == d.selected {
			prefix = "> "
			itemStyle = selActiveStyle
		}

		label := item.Label
		maxLabel := boxW - 6
		if len(label) > maxLabel {
			label = label[:maxLabel]
		}
		lines = append(lines, prefix+itemStyle.Render(label))
	}

	content := strings.Join(lines, "\n")
	return style.Render(content)
}

func (d *SelectDialog) Active() bool {
	return d.active
}

func (d *SelectDialog) SetActive(a bool) {
	d.active = a
}

func (d *SelectDialog) Focused() bool {
	return d.active
}

func (d *SelectDialog) SetFocused(f bool) {
	d.active = f
}

var _ = fmt.Sprintf
