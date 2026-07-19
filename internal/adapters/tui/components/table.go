package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Column struct {
	Name  string
	Width int
	Align lipgloss.Position
}

type TableRenderer struct {
	columns     []Column
	rows        [][]string
	width       int
	headerStyle lipgloss.Style
	rowStyle    lipgloss.Style
	altRowStyle lipgloss.Style
	borderStyle lipgloss.Style
	selected    int
}

func NewTableRenderer() *TableRenderer {
	return &TableRenderer{
		headerStyle: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39")),
		rowStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")),
		altRowStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("250")),
		borderStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")),
		selected: -1,
	}
}

func (t *TableRenderer) SetWidth(w int) {
	t.width = w
}

func (t *TableRenderer) SetColumns(cols []Column) {
	t.columns = cols
}

func (t *TableRenderer) SetRows(rows [][]string) {
	t.rows = rows
}

func (t *TableRenderer) SetSelected(row int) {
	t.selected = row
}

func (t *TableRenderer) Render() string {
	if len(t.columns) == 0 {
		return ""
	}

	totalW := 0
	for _, c := range t.columns {
		totalW += c.Width + 3
	}
	if totalW < 1 {
		totalW = 1
	}

	var b strings.Builder

	header := t.renderHeader()
	b.WriteString(t.borderStyle.Render(strings.Repeat("─", totalW)))
	b.WriteByte('\n')
	b.WriteString(header)
	b.WriteByte('\n')
	b.WriteString(t.borderStyle.Render(strings.Repeat("─", totalW)))
	b.WriteByte('\n')

	for i, row := range t.rows {
		rowStr := t.renderRow(row, i)
		b.WriteString(rowStr)
		b.WriteByte('\n')
	}

	b.WriteString(t.borderStyle.Render(strings.Repeat("─", totalW)))

	return b.String()
}

func (t *TableRenderer) renderHeader() string {
	var cells []string
	for _, col := range t.columns {
		name := col.Name
		if len([]rune(name)) > col.Width {
			name = string([]rune(name)[:col.Width])
		}
		cell := lipgloss.NewStyle().
			Width(col.Width).
			Align(col.Align).
			Render(name)
		cells = append(cells, cell)
	}
	return " " + strings.Join(cells, " ┃ ")
}

func (t *TableRenderer) renderRow(row []string, idx int) string {
	var cells []string
	for ci, col := range t.columns {
		val := ""
		if ci < len(row) {
			val = row[ci]
		}
		if len([]rune(val)) > col.Width {
			val = string([]rune(val)[:col.Width-1]) + "…"
		}
		style := t.rowStyle
		if idx%2 == 1 {
			style = t.altRowStyle
		}
		cell := style.Copy().
			Width(col.Width).
			Align(col.Align).
			Render(val)

		if idx == t.selected {
			cell = lipgloss.NewStyle().
				Foreground(lipgloss.Color("0")).
				Background(lipgloss.Color("39")).
				Width(col.Width).
				Align(col.Align).
				Render(val)
		}

		cells = append(cells, cell)
	}

	prefix := " "
	if idx == t.selected {
		prefix = ">"
	}
	return prefix + strings.Join(cells, " │ ")
}

func (t *TableRenderer) RenderSimple(headers []string, rows [][]string) string {
	cols := make([]Column, len(headers))
	for i, h := range headers {
		cols[i] = Column{Name: h, Width: 20, Align: lipgloss.Left}
	}
	t.columns = cols
	t.rows = rows
	return t.Render()
}

func Table(headers []string, rows [][]string) string {
	r := NewTableRenderer()
	return r.RenderSimple(headers, rows)
}
