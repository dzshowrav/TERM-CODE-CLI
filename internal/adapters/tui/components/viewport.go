package components

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

type Viewport struct {
	width   int
	height  int
	content []string
	offset  int
}

func NewViewport(width, height int) *Viewport {
	return &Viewport{
		width:   width,
		height:  height,
		content: []string{},
		offset:  0,
	}
}

func (v *Viewport) SetSize(w, h int) {
	v.width = w
	v.height = h
	v.clampOffset()
}

func (v *Viewport) SetContent(lines []string) {
	v.content = lines
	v.clampOffset()
	v.scrollToBottom()
}

func (v *Viewport) AddLine(line string) {
	v.content = append(v.content, line)
	v.scrollToBottom()
}

func (v *Viewport) scrollToBottom() {
	totalLines := len(v.content)
	if totalLines > v.height {
		v.offset = totalLines - v.height
	} else {
		v.offset = 0
	}
}

func (v *Viewport) ScrollUp(n int) {
	v.offset -= n
	v.clampOffset()
}

func (v *Viewport) ScrollDown(n int) {
	v.offset += n
	v.clampOffset()
}

func (v *Viewport) ScrollToTop() {
	v.offset = 0
}

func (v *Viewport) ScrollToBottom() {
	v.scrollToBottom()
}

func (v *Viewport) clampOffset() {
	totalLines := len(v.content)
	maxOffset := totalLines - v.height
	if maxOffset < 0 {
		maxOffset = 0
	}
	if v.offset > maxOffset {
		v.offset = maxOffset
	}
	if v.offset < 0 {
		v.offset = 0
	}
}

func (v *Viewport) AtBottom() bool {
	return v.offset >= len(v.content)-v.height
}

func (v *Viewport) Update(msg tea.Msg) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "pgup":
			v.ScrollUp(v.height)
		case "pgdown":
			v.ScrollDown(v.height)
		case "up":
			v.ScrollUp(1)
		case "down":
			v.ScrollDown(1)
		}
	}
}

func (v *Viewport) View() string {
	if len(v.content) == 0 {
		return strings.Repeat("\n", v.height-1)
	}

	visible := v.content
	if len(visible) > v.height {
		end := v.offset + v.height
		if end > len(visible) {
			end = len(visible)
		}
		visible = visible[v.offset:end]
	}

	for len(visible) < v.height {
		visible = append(visible, "")
	}

	return strings.Join(visible, "\n")
}
