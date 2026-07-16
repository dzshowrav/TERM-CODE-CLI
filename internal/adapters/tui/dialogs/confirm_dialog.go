package dialogs

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var (
	dialogStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("39")).
			Padding(1, 2).
			Width(40)
	dialogTitle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	dialogText     = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	dialogActive   = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	dialogInactive = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

type ConfirmResult bool

type ConfirmDialog struct {
	title    string
	message  string
	yes      bool
	active   bool
	onResult func(ConfirmResult) tea.Msg
}

func NewConfirmDialog(title, message string, onResult func(ConfirmResult) tea.Msg) *ConfirmDialog {
	return &ConfirmDialog{
		title:    title,
		message:  message,
		yes:      true,
		active:   true,
		onResult: onResult,
	}
}

func (d *ConfirmDialog) Init() tea.Cmd {
	return nil
}

func (d *ConfirmDialog) Update(msg tea.Msg) (*ConfirmDialog, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			d.active = false
			if d.onResult != nil {
				return d, func() tea.Msg {
					return d.onResult(ConfirmResult(d.yes))
				}
			}
		case "left", "right", "tab":
			d.yes = !d.yes
		case "esc":
			d.active = false
			if d.onResult != nil {
				return d, func() tea.Msg {
					return d.onResult(ConfirmResult(false))
				}
			}
		}
	}
	return d, nil
}

func (d *ConfirmDialog) View() string {
	yesBtn := dialogActive.Render("Yes")
	noBtn := dialogActive.Render("No")
	if d.yes {
		yesBtn = dialogActive.Render("> Yes <")
	} else {
		noBtn = dialogActive.Render("> No <")
	}

	content := fmt.Sprintf(
		"%s\n\n%s\n\n%s  %s",
		dialogTitle.Render(d.title),
		dialogText.Render(d.message),
		yesBtn, noBtn,
	)

	return dialogStyle.Render(content)
}

func (d *ConfirmDialog) Active() bool {
	return d.active
}

func (d *ConfirmDialog) SetActive(a bool) {
	d.active = a
}

func (d *ConfirmDialog) Focused() bool {
	return d.active
}

func (d *ConfirmDialog) SetFocused(f bool) {
	d.active = f
}
