package components

import (
	"strings"
)

type Viewport struct {
	content  []string
	offset   int
	width    int
	height   int
	maxLines int
}

func NewViewport(width, height int) *Viewport {
	return &Viewport{
		width:  width,
		height: height,
	}
}

func (v *Viewport) SetContent(content any) {
	var lines []string
	switch c := content.(type) {
	case string:
		lines = strings.Split(c, "\n")
	case []string:
		lines = c
	default:
		return
	}
	var wrapped []string
	for _, line := range lines {
		if v.width > 0 && len([]rune(line)) > v.width {
			runes := []rune(line)
			for len(runes) > 0 {
				if len(runes) <= v.width {
					wrapped = append(wrapped, string(runes))
					break
				}
				wrapped = append(wrapped, string(runes[:v.width]))
				runes = runes[v.width:]
			}
		} else {
			wrapped = append(wrapped, line)
		}
	}
	v.content = wrapped
	v.maxLines = len(wrapped)
	if v.offset > v.maxLines-v.height {
		v.offset = v.maxLines - v.height
	}
	if v.offset < 0 {
		v.offset = 0
	}
}

func (v *Viewport) SetWidth(w int) {
	v.width = w
}

func (v *Viewport) SetHeight(h int) {
	v.height = h
	if v.maxLines > 0 && v.offset+v.height > v.maxLines {
		v.offset = v.maxLines - v.height
	}
	if v.offset < 0 {
		v.offset = 0
	}
}

func (v *Viewport) SetSize(w, h int) {
	v.width = w
	v.height = h
	if v.maxLines > 0 && v.offset+v.height > v.maxLines {
		v.offset = v.maxLines - v.height
	}
	if v.offset < 0 {
		v.offset = 0
	}
}

func (v *Viewport) GotoBottom() {
	if v.maxLines > v.height {
		v.offset = v.maxLines - v.height
	} else {
		v.offset = 0
	}
}

func (v *Viewport) AddLine(line string) {
	v.content = append(v.content, line)
	v.maxLines = len(v.content)
	if len(v.content) > v.height {
		v.offset = len(v.content) - v.height
	}
}

func (v *Viewport) ScrollUp(n int) {
	v.offset -= n
	if v.offset < 0 {
		v.offset = 0
	}
}

func (v *Viewport) ScrollDown(n int) {
	v.offset += n
	if v.offset > v.maxLines-v.height {
		v.offset = v.maxLines - v.height
	}
	if v.offset < 0 {
		v.offset = 0
	}
}

func (v *Viewport) Height() int {
	return v.height
}

func (v *Viewport) AtTop() bool {
	return v.offset <= 0
}

func (v *Viewport) AtBottom() bool {
	return v.offset >= v.maxLines-v.height
}

func (v *Viewport) Update(msg any) {
}

func (v *Viewport) View() string {
	if len(v.content) == 0 {
		return strings.Repeat("\n", v.height-1)
	}

	end := v.offset + v.height
	if end > len(v.content) {
		end = len(v.content)
	}

	visible := v.content[v.offset:end]
	result := strings.Join(visible, "\n")
	if pad := v.height - len(visible); pad > 0 {
		result += strings.Repeat("\n", pad)
	}
	return result
}
