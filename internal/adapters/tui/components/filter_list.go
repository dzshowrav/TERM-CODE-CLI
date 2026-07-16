package components

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var (
	listItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	listSelStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	listSubStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

type FilterList struct {
	items    []FilterItem
	selected int
	filter   string
	width    int
	height   int
	offset   int
	onSelect func(item FilterItem) tea.Msg
}

type FilterItem struct {
	ID       string
	Label    string
	Subtitle string
	Active   bool
}

func NewFilterList(onSelect func(item FilterItem) tea.Msg) *FilterList {
	return &FilterList{
		selected: 0,
		width:    80,
		height:   20,
		onSelect: onSelect,
	}
}

func (l *FilterList) SetItems(items []FilterItem) {
	l.items = items
	if l.selected >= len(items) {
		l.selected = 0
	}
}

func (l *FilterList) SetSize(w, h int) {
	l.width = w
	l.height = h
}

func (l *FilterList) SetFilter(f string) {
	l.filter = f
	l.selected = 0
}

func (l *FilterList) filtered() []FilterItem {
	if l.filter == "" {
		return l.items
	}
	f := strings.ToLower(l.filter)
	var result []FilterItem
	for _, item := range l.items {
		if strings.Contains(strings.ToLower(item.Label), f) {
			result = append(result, item)
		}
	}
	return result
}

func (l *FilterList) SelectedItem() *FilterItem {
	items := l.filtered()
	if len(items) == 0 {
		return nil
	}
	return &items[l.selected]
}

func (l *FilterList) Init() tea.Cmd {
	return nil
}

func (l *FilterList) Update(msg tea.Msg) (*FilterList, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if item := l.SelectedItem(); item != nil && l.onSelect != nil {
				return l, func() tea.Msg {
					return l.onSelect(*item)
				}
			}
		case "up", "ctrl+p":
			if l.selected > 0 {
				l.selected--
			}
		case "down", "ctrl+n":
			items := l.filtered()
			if l.selected < len(items)-1 {
				l.selected++
			}
		case "pgup":
			l.selected -= l.height
			if l.selected < 0 {
				l.selected = 0
			}
		case "pgdown":
			l.selected += l.height
			items := l.filtered()
			if l.selected >= len(items) {
				l.selected = len(items) - 1
			}
		case "home":
			l.selected = 0
		case "end":
			items := l.filtered()
			l.selected = len(items) - 1
		}
	}
	return l, nil
}

func (l *FilterList) View() string {
	items := l.filtered()
	if len(items) == 0 {
		return listSubStyle.Render("No items.")
	}

	var lines []string
	for i, item := range items {
		prefix := "  "
		if i == l.selected {
			prefix = "> "
		}

		var style lipgloss.Style
		if i == l.selected {
			style = listSelStyle
		} else {
			style = listItemStyle
		}

		activeMark := ""
		if item.Active {
			activeMark = listSelStyle.Render(" ✓")
		}

		lines = append(lines, prefix+style.Render(item.Label)+activeMark)
		if item.Subtitle != "" {
			lines = append(lines, "   "+listSubStyle.Render(item.Subtitle))
		}
	}

	visible := lines
	if len(visible) > l.height {
		visible = visible[:l.height]
	}

	return strings.Join(visible, "\n")
}
