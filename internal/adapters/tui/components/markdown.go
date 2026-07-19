package components

import (
	"strings"
	"unicode"

	"github.com/charmbracelet/lipgloss"
)

type MarkdownRenderer struct {
	width       int
	codeStyle   lipgloss.Style
	h1, h2, h3  lipgloss.Style
	quietStyle  lipgloss.Style
	bulletStyle lipgloss.Style
	inlineCode  lipgloss.Style
	sepStyle    lipgloss.Style
}

func NewMarkdownRenderer() *MarkdownRenderer {
	return &MarkdownRenderer{
		width: 80,
		codeStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			Background(lipgloss.Color("234")).
			Padding(0, 2),
		h1: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39")).
			Underline(true),
		h2: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39")),
		h3: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("45")),
		quietStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("242")).
			Italic(true),
		bulletStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("250")),
		inlineCode: lipgloss.NewStyle().
			Foreground(lipgloss.Color("43")).
			Background(lipgloss.Color("236")),
		sepStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")),
	}
}

func (r *MarkdownRenderer) SetWidth(w int) {
	r.width = w
}

func (r *MarkdownRenderer) Render(markdown string) string {
	lines := strings.Split(markdown, "\n")
	var out []string
	inBlock := false

	for _, line := range lines {
		trim := strings.TrimSpace(line)

		if strings.HasPrefix(trim, "```") {
			if inBlock {
				out = append(out, r.codeStyle.Render(strings.Repeat(" ", r.width-4)))
				inBlock = false
				continue
			}
			inBlock = true
			out = append(out, r.codeStyle.Render(strings.Repeat(" ", r.width-4)))
			continue
		}

		if inBlock {
			out = append(out, r.codeStyle.Render(trim))
			continue
		}

		if trim == "" {
			out = append(out, "")
			continue
		}

		out = append(out, r.renderLine(trim))
	}

	return strings.Join(out, "\n")
}

func (r *MarkdownRenderer) renderLine(line string) string {
	for i := 6; i >= 1; i-- {
		p := strings.Repeat("#", i) + " "
		if strings.HasPrefix(line, p) {
			switch i {
			case 1:
				return r.h1.Render(line[len(p):])
			case 2:
				return r.h2.Render(line[len(p):])
			case 3:
				return r.h3.Render(line[len(p):])
			default:
				return r.h3.Render(line[len(p):])
			}
		}
	}

	if strings.HasPrefix(line, "> ") {
		return r.quietStyle.Render("│ " + line[2:])
	}

	if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
		return r.bulletStyle.Render("• " + line[2:])
	}

	if len(line) > 2 && unicode.IsDigit(rune(line[0])) && line[1] == '.' && line[2] == ' ' {
		return r.bulletStyle.Render(line)
	}

	if line == "---" || line == "***" || line == "___" {
		return r.sepStyle.Render(strings.Repeat("─", r.width))
	}

	return line
}
