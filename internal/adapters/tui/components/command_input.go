package components

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var (
	inputPrompt = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("> ")
	inputStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	inputWrap   = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderTop(true).BorderBottom(true).BorderForeground(lipgloss.Color("236"))
)

type SubmitMsg string

type CommandInput struct {
	content strings.Builder
	value   string
	focused bool
	width   int
	cursor  int
}

func NewCommandInput() *CommandInput {
	return &CommandInput{
		focused: true,
		width:   80,
	}
}

func (c *CommandInput) SetWidth(w int) {
	c.width = w
}

func (c *CommandInput) Focused() bool {
	return c.focused
}

func (c *CommandInput) SetFocused(f bool) {
	c.focused = f
}

func (c *CommandInput) Value() string {
	return c.value
}

func (c *CommandInput) SetValue(v string) {
	c.value = v
	c.cursor = len(v)
}

func (c *CommandInput) Init() tea.Cmd {
	return nil
}

func (c *CommandInput) Update(msg tea.Msg) (*CommandInput, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.PasteMsg:
		c.value = c.value[:c.cursor] + msg.String() + c.value[c.cursor:]
		c.cursor += len(msg.String())

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			val := c.value
			if val != "" {
				c.value = ""
				c.cursor = 0
				return c, func() tea.Msg {
					return SubmitMsg(val)
				}
			}

		case "backspace":
			if c.cursor > 0 {
				c.value = c.value[:c.cursor-1] + c.value[c.cursor:]
				c.cursor--
			}

		case "delete", "ctrl+d":
			if c.cursor < len(c.value) {
				c.value = c.value[:c.cursor] + c.value[c.cursor+1:]
			}

		case "left", "ctrl+b":
			if c.cursor > 0 {
				c.cursor--
			}

		case "right", "ctrl+f":
			if c.cursor < len(c.value) {
				c.cursor++
			}

		case "home", "ctrl+a":
			c.cursor = 0

		case "end", "ctrl+e":
			c.cursor = len(c.value)

		case "ctrl+w":
			spaceIdx := strings.LastIndex(c.value[:c.cursor], " ")
			if spaceIdx == -1 {
				c.value = c.value[c.cursor:]
				c.cursor = 0
			} else {
				c.value = c.value[:spaceIdx+1] + c.value[c.cursor:]
				c.cursor = spaceIdx + 1
			}

		case "ctrl+u":
			c.value = c.value[c.cursor:]
			c.cursor = 0

		case "ctrl+k":
			c.value = c.value[:c.cursor]

		default:
			var r string
			if msg.String() == "space" {
				r = " "
			} else if len(msg.String()) == 1 && msg.String()[0] >= 32 {
				r = msg.String()
			}
			if r != "" {
				c.value = c.value[:c.cursor] + r + c.value[c.cursor:]
				c.cursor++
			}
		}
	}

	return c, nil
}

func (c *CommandInput) View() string {
	if !c.focused {
		return ""
	}

	display := c.value
	var content string
	if display == "" {
		display = "Type a message or / for commands..."
		runes := []rune(display)
		maxLen := c.width - 4
		if len(runes) > maxLen {
			runes = runes[:maxLen]
		}
		content = inputPrompt + "█" + string(runes)
	} else {
		runes := []rune(display)
		maxLen := c.width - 4
		if len(runes) > maxLen {
			runes = runes[len(runes)-maxLen:]
		}
		disp := string(runes)

		cRunes := []rune(disp)
		cursorPos := min(c.cursor, len(cRunes))
		if cursorPos > len(cRunes) {
			cursorPos = len(cRunes)
		}

		cursor := "█"

		before := string(cRunes[:cursorPos])
		after := string(cRunes[cursorPos:])
		content = inputPrompt + before + cursor + after
	}

	return inputWrap.Copy().Width(c.width).Render(inputStyle.Render(content))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
