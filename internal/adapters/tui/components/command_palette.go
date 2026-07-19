package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type CommandEntry struct {
	Command     string
	Description string
}

var DefaultCommands = []CommandEntry{
	{Command: "help", Description: "Show available commands"},
	{Command: "about", Description: "Show about TermCode"},
	{Command: "network", Description: "Show network status"},
	{Command: "providers", Description: "List, select, edit, delete providers"},
	{Command: "provider add", Description: "Add a new provider"},
	{Command: "provider sync", Description: "Sync models from provider"},
	{Command: "models", Description: "List & select models"},
	{Command: "addmodel", Description: "Add a custom model"},
	{Command: "agents", Description: "List & select agents"},
	{Command: "workspace", Description: "Show/set workspace path"},
	{Command: "sessions list", Description: "List saved sessions"},
	{Command: "sessions new", Description: "Start a new session"},
	{Command: "sessions delete", Description: "Delete a session"},
	{Command: "git status", Description: "Show git status"},
	{Command: "git log", Description: "Show commit log"},
	{Command: "git diff", Description: "Show working tree changes"},
	{Command: "git add", Description: "Stage files"},
	{Command: "git commit", Description: "Commit staged changes"},
	{Command: "git branches", Description: "List branches"},
	{Command: "clear", Description: "Clear session / return home"},
	{Command: "home", Description: "Return to home screen"},
	{Command: "exit", Description: "Close TermCode"},
}

var (
	paletteBorder  = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("240"))
	paletteSel     = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	paletteItem    = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
	paletteDesc    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	paletteDim     = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
	paletteScroll  = lipgloss.NewStyle().Foreground(lipgloss.Color("238")).Italic(true)
	paletteFilter  = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Bold(true)
	paletteNoMatch = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
)

type CommandPalette struct {
	visible  bool
	entries  []CommandEntry
	filtered []CommandEntry
	selected int
	filter   string
	maxItems int
	scroll   int
	width    int
}

func NewCommandPalette() *CommandPalette {
	return &CommandPalette{
		entries:  DefaultCommands,
		maxItems: 6,
		width:    80,
	}
}

func (p *CommandPalette) SetWidth(w int) {
	p.width = w
}

func (p *CommandPalette) SetMaxItems(n int) {
	p.maxItems = n
	if p.maxItems < 1 {
		p.maxItems = 1
	}
}

func (p *CommandPalette) Visible() bool {
	return p.visible
}

func (p *CommandPalette) Show() {
	p.visible = true
	p.selected = 0
	p.scroll = 0
	p.filter = ""
	p.filtered = p.entries
}

func (p *CommandPalette) Hide() {
	p.visible = false
}

func fuzzyScore(input, target string) int {
	input = strings.ToLower(input)
	target = strings.ToLower(target)

	if strings.Contains(target, input) {
		return 100 + len(input)
	}

	score := 0
	ti := 0
	for _, ch := range input {
		for ti < len(target) {
			ti++
			if rune(target[ti-1]) == ch {
				score += 10
				break
			}
		}
		if ti >= len(target) {
			score -= 5
		}
	}
	if score > 0 && strings.HasPrefix(target, input) {
		score += 20
	}
	return score
}

func (p *CommandPalette) SetFilter(filter string) {
	p.filter = filter
	p.selected = 0
	p.scroll = 0

	if filter == "" {
		p.filtered = p.entries
		return
	}

	type scoredEntry struct {
		entry CommandEntry
		score int
	}

	var scored []scoredEntry
	for _, e := range p.entries {
		s := fuzzyScore(filter, e.Command)
		if s > 0 {
			scored = append(scored, scoredEntry{entry: e, score: s})
		}
	}

	for i := 0; i < len(scored); i++ {
		for j := i + 1; j < len(scored); j++ {
			if scored[j].score > scored[i].score {
				scored[i], scored[j] = scored[j], scored[i]
			}
		}
	}

	p.filtered = make([]CommandEntry, len(scored))
	for i, se := range scored {
		p.filtered[i] = se.entry
	}
}

func (p *CommandPalette) Filter() string {
	return p.filter
}

func (p *CommandPalette) SelectedCommand() string {
	if len(p.filtered) == 0 {
		return ""
	}
	return "/" + p.filtered[p.selected].Command
}

func (p *CommandPalette) SelectUp() {
	if len(p.filtered) == 0 {
		return
	}
	p.selected--
	if p.selected < 0 {
		p.selected = len(p.filtered) - 1
	}
	p.ensureVisible()
}

func (p *CommandPalette) SelectDown() {
	if len(p.filtered) == 0 {
		return
	}
	p.selected++
	if p.selected >= len(p.filtered) {
		p.selected = 0
	}
	p.ensureVisible()
}

func (p *CommandPalette) ensureVisible() {
	if p.selected < p.scroll {
		p.scroll = p.selected
	}
	if p.selected >= p.scroll+p.maxItems {
		p.scroll = p.selected - p.maxItems + 1
	}
}

func (p *CommandPalette) Count() int {
	return len(p.entries)
}

func (p *CommandPalette) FilteredCount() int {
	return len(p.filtered)
}

func (p *CommandPalette) Height() int {
	if !p.visible {
		return 0
	}
	if len(p.filtered) == 0 {
		return 3 // border top + message + border bottom
	}
	h := 2 // borders
	start := p.scroll
	end := start + p.maxItems
	if end > len(p.filtered) {
		end = len(p.filtered)
	}
	h += end - start
	if start > 0 {
		h++
	}
	if end < len(p.filtered) {
		h++
	}
	return h
}

func (p *CommandPalette) View() string {
	if !p.visible {
		return ""
	}

	if len(p.filtered) == 0 {
		content := paletteNoMatch.Render(fmt.Sprintf("No commands match '%s'", p.filter))
		return paletteBorder.Render(content)
	}

	var lines []string

	if p.scroll > 0 {
		lines = append(lines, paletteScroll.Render(fmt.Sprintf("  ↑ %d more", p.scroll)))
	}

	end := p.scroll + p.maxItems
	if end > len(p.filtered) {
		end = len(p.filtered)
	}

	displayItems := p.filtered[p.scroll:end]
	for i, entry := range displayItems {
		idx := p.scroll + i
		prefix := "  "
		cmdStyle := paletteItem
		if idx == p.selected {
			prefix = " >"
			cmdStyle = paletteSel
		}

		cmd := "/" + entry.Command
		desc := entry.Description

		cmdW := p.width / 3
		if cmdW < 16 {
			cmdW = 16
		}
		if cmdW > 22 {
			cmdW = 22
		}
		paddedCmd := cmd
		if len([]rune(cmd)) > cmdW {
			paddedCmd = string([]rune(cmd)[:cmdW])
		} else {
			paddedCmd = cmd + strings.Repeat(" ", cmdW-len([]rune(cmd)))
		}

		descMax := p.width - 10 - cmdW
		if descMax < 0 {
			descMax = 0
		}
		if len([]rune(desc)) > descMax {
			desc = string([]rune(desc)[:descMax])
		}

		lines = append(lines, fmt.Sprintf("%s %s%s %s",
			paletteDim.Render(prefix),
			cmdStyle.Render(paddedCmd),
			paletteDim.Render(" "),
			paletteDesc.Render(desc)))
	}

	if end < len(p.filtered) {
		remaining := len(p.filtered) - end
		lines = append(lines, paletteScroll.Render(fmt.Sprintf("  ↓ %d more", remaining)))
	}

	content := strings.Join(lines, "\n")
	return paletteBorder.Width(p.width - 4).Render(content)
}
