package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	codeBlockStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("235")).
			Foreground(lipgloss.Color("83")).
			Padding(0, 2)
	codeLangStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("236")).
			Foreground(lipgloss.Color("245")).
			Padding(0, 2)
)

type CodeBlock struct {
	language string
	code     string
	width    int
}

func NewCodeBlock(language, code string) *CodeBlock {
	return &CodeBlock{
		language: language,
		code:     code,
		width:    80,
	}
}

func (b *CodeBlock) SetWidth(w int) {
	b.width = w
}

func (b *CodeBlock) Render() string {
	var parts []string

	if b.language != "" {
		parts = append(parts, codeLangStyle.Render(b.language))
	}

	codeWidth := b.width - 4
	if codeWidth < 10 {
		codeWidth = 10
	}

	lines := strings.Split(b.code, "\n")
	var formatted []string
	for _, line := range lines {
		if len(line) > codeWidth {
			line = line[:codeWidth]
		}
		formatted = append(formatted, line)
	}

	parts = append(parts, codeBlockStyle.Width(codeWidth).Render(strings.Join(formatted, "\n")))

	return strings.Join(parts, "\n")
}
