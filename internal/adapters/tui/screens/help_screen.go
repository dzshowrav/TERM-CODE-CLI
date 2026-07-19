package screens

import (
	"fmt"
	"strings"
	"unicode"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	helpSec  = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Bold(true)
	helpCmd  = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
	helpDesc = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
)

type helpEntry struct {
	cmd  string
	desc string
}

type HelpScreen struct {
	width    int
	height   int
	all      []helpEntry
	filtered []helpEntry
	search   string
	scroll   int
	done     bool
}

func NewHelpScreen() *HelpScreen {
	entries := []helpEntry{
		{"/help", "Show this help screen"},
		{"/clear", "Clear current conversation"},
		{"/home", "Return to home screen"},
		{"/retry", "Retry last AI response"},
		{"/cancel", "Cancel current operation"},
		{"/stop", "Stop active generation"},
		{"/continue", "Continue from paused/stopped state"},
		{"/provider", "Provider management (add, sync)"},
		{"/providers", "Provider management (list, select, edit, delete)"},
		{"/models", "List, select & manage models"},
		{"/addmodel", "Add a custom model"},
		{"/sessions", "Session management (list, new, delete)"},
		{"/workspace", "Show/set workspace"},
		{"/agents", "Agent selection"},
		{"/git", "Git operations (status, log, diff, add, commit, branches)"},
		{"/settings", "Open settings panel"},
		{"/theme", "Change color theme"},
		{"/tools", "List and manage available tools"},
		{"/collab", "Collaboration (start, connect, sync, push, stop)"},
		{"/edit", "Edit a message by index"},
		{"/delete", "Delete a message by index"},
		{"/rename", "Rename current session"},
		{"/branch", "Branch management (list or create)"},
		{"/undo", "Undo last edit/delete action"},
		{"/redo", "Redo last undone action"},
		{"/export", "Export current session to JSON"},
		{"/import", "Import a session from JSON file"},
		{"/pin", "Pin current session"},
		{"/search", "Search sessions by name"},
		{"/status", "Show network/connection status"},
		{"/network", "Show network/connection status"},
		{"/about", "Show application information"},
		{"/batch", "Open batch edit mode"},
		{"/exit", "Quit the application"},
	}
	return &HelpScreen{
		width:  80,
		height: 24,
		all:    entries,
	}
}

func (s *HelpScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *HelpScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "up", "k":
			if s.scroll > 0 {
				s.scroll--
			}
		case "down", "j":
			avail := s.height - 10
			if avail < 1 {
				avail = 1
			}
			arrowCount := 0
			if s.scroll > 0 {
				arrowCount++
			}
			if s.scroll+avail-arrowCount < len(s.filtered) {
				arrowCount++
			}
			maxVis := avail - arrowCount
			if maxVis < 1 {
				maxVis = 1
			}
			if s.scroll < len(s.filtered)-maxVis {
				s.scroll++
			}
		case "backspace":
			if len(s.search) > 0 {
				s.search = s.search[:len(s.search)-1]
				s.applyFilter()
			}
		default:
			r := []rune(msg.String())
			if len(r) == 1 && !unicode.IsControl(r[0]) {
				s.search += string(r)
				s.scroll = 0
				s.applyFilter()
			}
		}
	case tea.PasteMsg:
		s.search += msg.String()
		s.scroll = 0
		s.applyFilter()
	}
	return s, nil
}

func (s *HelpScreen) applyFilter() {
	s.filtered = nil
	if s.search == "" {
		s.filtered = append(s.filtered, s.all...)
		return
	}
	lower := strings.ToLower(s.search)
	for _, e := range s.all {
		if strings.Contains(strings.ToLower(e.cmd), lower) ||
			strings.Contains(strings.ToLower(e.desc), lower) {
			s.filtered = append(s.filtered, e)
		}
	}
}

func (s *HelpScreen) Done() bool     { return s.done }
func (s *HelpScreen) Result() string { return "" }

func (s *HelpScreen) View() string {
	innerW := s.width - 2

	title := fmt.Sprintf("%-*s%s", innerW-4, "Help", "esc")

	var searchLine string
	if s.search == "" {
		searchLine = fmt.Sprintf("Search: %s%s", "█", styles.HintStyle.Render("search commands"))
	} else {
		searchLine = fmt.Sprintf("Search: %s%s", s.search, "█")
	}

	if s.filtered == nil {
		s.applyFilter()
	}

	maxVis := s.height - 10
	if maxVis < 1 {
		maxVis = 1
	}

	needsTop := s.scroll > 0
	needsBottom := s.scroll+maxVis < len(s.filtered)
	if needsTop {
		maxVis--
	}
	if needsBottom {
		maxVis--
	}
	if maxVis < 1 {
		maxVis = 1
	}

	if s.scroll+maxVis > len(s.filtered) {
		s.scroll = max(0, len(s.filtered)-maxVis)
	}
	end := s.scroll + maxVis
	if end > len(s.filtered) {
		end = len(s.filtered)
	}

	var bodyLines []string
	bodyLines = append(bodyLines, title)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, searchLine)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))

	if len(s.filtered) == 0 {
		bodyLines = append(bodyLines, styles.HintStyle.Render(fmt.Sprintf(" No commands match '%s'", s.search)))
	} else {
		cmdW := innerW / 3
		if cmdW < 16 {
			cmdW = 16
		}
		if cmdW > 22 {
			cmdW = 22
		}

		for i := s.scroll; i < end; i++ {
			e := s.filtered[i]
			padding := cmdW - len(e.cmd)
			if padding < 1 {
				padding = 1
			}
			bodyLines = append(bodyLines, fmt.Sprintf(
				"  %s%s%s",
				helpCmd.Render(e.cmd),
				strings.Repeat(" ", padding),
				helpDesc.Render(e.desc),
			))
		}
	}

	if needsTop {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll))))
	}
	if needsBottom {
		remaining := len(s.filtered) - end
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↓ %d more", remaining))))
	}

	bodyLines = append(bodyLines, "")
	hintText := "esc: Close"
	bodyLines = append(bodyLines, fmt.Sprintf("%*s", innerW, styles.HintStyle.Render(hintText)))

	return styles.DialogBox(s.width, strings.Join(bodyLines, "\n"))
}
