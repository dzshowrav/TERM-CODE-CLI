package screens

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

type MCPAuthDialog struct {
	width      int
	height     int
	serverName string
	serverURL  string
	done       bool
	result     string
	focusField int
}

func NewMCPAuthDialog(serverName, serverURL string) *MCPAuthDialog {
	return &MCPAuthDialog{
		serverName: serverName,
		serverURL:  serverURL,
		focusField: 1,
	}
}

func (d *MCPAuthDialog) SetSize(w, h int) {
	d.width = w
	d.height = h
}

func (d *MCPAuthDialog) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
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

		case "tab", "shift+tab", "up", "down":
			d.focusField = 1 - d.focusField
			return d, nil

		case "enter":
			if d.focusField == 1 {
				d.done = true
				d.result = "allow"
			} else {
				d.done = true
				d.result = ""
			}
			return d, nil
		}
	}
	return d, nil
}

func (d *MCPAuthDialog) View() string {
	var body strings.Builder

	body.WriteString(styles.Warning.Render("MCP Server Authorization") + "\n\n")
	body.WriteString(styles.Accent.Render("Server: ") + d.serverName + "\n")
	body.WriteString(styles.Dim.Render("URL: ") + d.serverURL + "\n\n")
	body.WriteString(styles.Dim.Render("This server requested access to:"))
	body.WriteString("\n" + styles.Dim.Render("  \u2022 Read files in the workspace"))
	body.WriteString("\n" + styles.Dim.Render("  \u2022 Execute tools and commands"))
	body.WriteString("\n" + styles.Dim.Render("  \u2022 Access workspace resources"))
	body.WriteString("\n\n")
	body.WriteString(styles.Dim.Render("Allow this MCP server to connect?") + "\n\n")

	innerW := d.width - 2
	allowBtn := styles.FormBtnActive.Render(" Allow ")
	denyBtn := styles.FormBtnNormal.Render(" Deny ")
	if d.focusField == 0 {
		denyBtn = styles.FormBtnActive.Render(" Deny ")
		allowBtn = styles.FormBtnNormal.Render(" Allow ")
	}

	body.WriteString(lipgloss.NewStyle().Width(innerW).Align(lipgloss.Center).Render(
		allowBtn + "  " + denyBtn,
	))

	fullBody := body.String() + "\n" +
		styles.DialogSep(innerW) + "\n" +
		styles.Help.Render(" tab: navigate  enter: confirm  esc: cancel")

	return styles.DialogBox(d.width, fullBody)
}

func (d *MCPAuthDialog) Done() bool {
	return d.done
}

func (d *MCPAuthDialog) Result() string {
	return d.result
}
