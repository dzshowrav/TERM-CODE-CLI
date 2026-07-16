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
	{Command: "provider list", Description: "List configured providers"},
	{Command: "provider add", Description: "Add a new provider"},
	{Command: "provider select", Description: "Select active provider"},
	{Command: "provider sync", Description: "Sync models from provider"},
	{Command: "model list", Description: "List available models"},
	{Command: "model select", Description: "Select active model"},
	{Command: "agent list", Description: "List available agents"},
	{Command: "agent select", Description: "Select active agent"},
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

func (p *CommandPalette) SetFilter(filter string) {
	p.filter = filter
	p.selected = 0
	p.scroll = 0

	if filter == "" {
		p.filtered = p.entries
		return
	}

	lower := strings.ToLower(filter)
	var matched []CommandEntry
	for _, e := range p.entries {
		if strings.Contains(strings.ToLower(e.Command), lower) {
			matched = append(matched, e)
		}
	}
	p.filtered = matched
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
