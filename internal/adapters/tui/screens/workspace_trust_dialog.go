package screens

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type WorkspaceTrustDialog struct {
	width         int
	height        int
	workspacePath string
	done          bool
	result        string
	focusField    int
	dangerous     bool
}

func NewWorkspaceTrustDialog(path string, dangerous bool) *WorkspaceTrustDialog {
	return &WorkspaceTrustDialog{
		workspacePath: path,
		dangerous:     dangerous,
		focusField:    0,
	}
}

func (d *WorkspaceTrustDialog) SetSize(w, h int) {
	d.width = w
	d.height = h
}

func (d *WorkspaceTrustDialog) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
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

		case "tab", "down":
			d.focusField = (d.focusField + 1) % 2
			return d, nil
		case "shift+tab", "up":
			d.focusField = (d.focusField - 1 + 2) % 2
			return d, nil

		case "enter":
			if d.focusField == 0 {
				d.done = true
				d.result = ""
			} else {
				d.done = true
				d.result = "__trust__"
			}
			return d, nil
		}
	}
	return d, nil
}

func (d *WorkspaceTrustDialog) View() string {
	var body strings.Builder

	body.WriteString(styles.Warning.Render("Unverified Workspace") + "\n\n")
	body.WriteString(styles.Dim.Render("Path: ") + d.workspacePath + "\n\n")

	if d.dangerous {
		body.WriteString(styles.Warning.Render("This workspace contains files that could execute code.") + "\n")
		body.WriteString(styles.Dim.Render("Always verify untrusted workspaces before allowing") + "\n")
		body.WriteString(styles.Dim.Render("file operations, git commands, or code execution.") + "\n\n")
	} else {
		body.WriteString(styles.Dim.Render("Allow this workspace to access project files and") + "\n")
		body.WriteString(styles.Dim.Render("run commands?") + "\n\n")
	}

	innerW := d.width - 2
	if innerW < 1 {
		innerW = 40
	}
	trustBtn := styles.FormBtnActive.Render(" Trust ")
	denyBtn := styles.FormBtnNormal.Render(" Deny ")
	if d.focusField == 0 {
		trustBtn = styles.FormBtnNormal.Render(" Trust ")
		denyBtn = styles.FormBtnActive.Render(" Deny ")
	}

	body.WriteString(lipgloss.NewStyle().Width(innerW).Align(lipgloss.Center).Render(
		trustBtn + "  " + denyBtn,
	))

	fullBody := body.String() + "\n" +
		styles.DialogSep(innerW) + "\n" +
		styles.Help.Render(" tab/shift+tab: navigate  enter: confirm  esc: cancel")

	return styles.DialogBox(d.width, fullBody)
}

func (d *WorkspaceTrustDialog) Done() bool {
	return d.done
}

func (d *WorkspaceTrustDialog) Result() string {
	return d.result
}
