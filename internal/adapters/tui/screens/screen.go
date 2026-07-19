package screens

import tea "charm.land/bubbletea/v2"

type Screen interface {
	SetSize(w, h int)
	Update(msg tea.Msg) (Screen, tea.Cmd)
	View() string
}

type FullScreen interface {
	Screen
	IsFullScreen() bool
	Title() string
}
