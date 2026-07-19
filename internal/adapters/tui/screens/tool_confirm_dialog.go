package screens

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type ToolConfirmDialog struct {
	width      int
	height     int
	toolName   string
	toolArgs   string
	reason     string
	done       bool
	result     string
	focusField int
}

func NewToolConfirmDialog(name, args, reason string) *ToolConfirmDialog {
	return &ToolConfirmDialog{
		toolName:   name,
		toolArgs:   args,
		reason:     reason,
		focusField: 0,
	}
}

func (d *ToolConfirmDialog) SetSize(w, h int) {
	d.width = w
	d.height = h
}

func (d *ToolConfirmDialog) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		d.SetSize(msg.Width, msg.Height)
		return d, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			d.done = true
			d.result = ""
			return d, nil

		case "tab", "right":
			d.focusField = (d.focusField + 1) % 3
		case "shift+tab", "left":
			d.focusField = (d.focusField - 1 + 3) % 3
			return d, nil

		case "enter":
			switch d.focusField {
			case 0:
				d.done = true
				d.result = "allow_once"
			case 1:
				d.done = true
				d.result = "always_allow"
			case 2:
				d.done = true
				d.result = "deny"
			}
			return d, nil
		}
	}
	return d, nil
}

func (d *ToolConfirmDialog) View() string {
	var body strings.Builder

	body.WriteString(styles.Warning.Render("Permission Required") + "\n\n")

	body.WriteString(styles.Accent.Render("Tool: ") + d.toolName + "\n")
	if d.toolArgs != "" {
		body.WriteString(styles.Dim.Render("Args: ") + d.toolArgs + "\n")
	}
	if d.reason != "" {
		body.WriteString(styles.Dim.Render("Reason: ") + d.reason + "\n")
	}

	body.WriteString("\n" + styles.Dim.Render("Choose permission:") + "\n\n")

	onceBtn := styles.FormBtnNormal.Render(" Allow Once ")
	alwaysBtn := styles.FormBtnNormal.Render(" Always Allow ")
	denyBtn := styles.FormBtnNormal.Render(" Deny ")
	switch d.focusField {
	case 0:
		onceBtn = styles.FormBtnActive.Render(" Allow Once ")
	case 1:
		alwaysBtn = styles.FormBtnActive.Render(" Always Allow ")
	case 2:
		denyBtn = styles.FormBtnActive.Render(" Deny ")
	}

	innerW := d.width - 2
	body.WriteString(lipgloss.NewStyle().Width(innerW).Align(lipgloss.Center).Render(
		onceBtn + "  " + alwaysBtn + "  " + denyBtn,
	))

	fullBody := body.String() + "\n" +
		styles.DialogSep(innerW) + "\n" +
		styles.Help.Render(" tab/arrows: navigate  enter: confirm  esc: cancel")

	return styles.DialogBox(d.width, fullBody)
}

func (d *ToolConfirmDialog) Done() bool {
	return d.done
}

func (d *ToolConfirmDialog) Result() string {
	return d.result
}
