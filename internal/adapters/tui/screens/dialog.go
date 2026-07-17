package screens

import tea "charm.land/bubbletea/v2"

type DialogScreen interface {
	SetSize(w, h int)
	Update(msg tea.Msg) (DialogScreen, tea.Cmd)
	View() string
	Done() bool
	Result() string
}
