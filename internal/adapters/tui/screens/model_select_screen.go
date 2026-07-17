package screens

import (
	"fmt"
	"strings"
	"unicode"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
	domainmodel "termcode/internal/domain/model"
)

type ModelSelectScreen struct {
	width    int
	height   int
	models   []domainmodel.Model
	filtered []domainmodel.Model
	search   string
	cursor   int
	scroll   int
	done     bool
	result   string
	onSelect func(modelID string) string
}

func NewModelSelectScreen(onSelect func(modelID string) string) *ModelSelectScreen {
	return &ModelSelectScreen{
		width:    80,
		height:   24,
		onSelect: onSelect,
	}
}

func (s *ModelSelectScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ModelSelectScreen) SetModels(models []domainmodel.Model) {
	s.models = models
	s.applyFilter()
}

func (s *ModelSelectScreen) Done() bool {
	return s.done
}

func (s *ModelSelectScreen) Result() string {
	return s.result
}

func (s *ModelSelectScreen) applyFilter() {
	s.filtered = make([]domainmodel.Model, 0, len(s.models))
	if s.search == "" {
		s.filtered = append(s.filtered, s.models...)
	} else {
		lower := strings.ToLower(s.search)
		for _, m := range s.models {
			if strings.Contains(strings.ToLower(m.DisplayName), lower) ||
				strings.Contains(strings.ToLower(m.ModelID), lower) {
				s.filtered = append(s.filtered, m)
			}
		}
	}
	if s.cursor >= len(s.filtered) {
		s.cursor = max(0, len(s.filtered)-1)
	}
	if s.cursor < 0 {
		s.cursor = 0
	}
}

func (s *ModelSelectScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
			s.result = ""
			return s, nil
		case "enter":
			if len(s.filtered) > 0 && s.cursor >= 0 && s.cursor < len(s.filtered) {
				m := s.filtered[s.cursor]
				s.result = s.onSelect(m.ModelID)
				s.done = true
			}
			return s, nil
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
				s.ensureVisible()
			}
		case "down", "j":
			if s.cursor < len(s.filtered)-1 {
				s.cursor++
				s.ensureVisible()
			}
		case "backspace":
			if len(s.search) > 0 {
				s.search = s.search[:len(s.search)-1]
				s.cursor = 0
				s.scroll = 0
				s.applyFilter()
			}
		case "ctrl+f":
			return s, nil
		case "ctrl+p":
			return s, nil
		default:
			r := []rune(msg.String())
			if len(r) == 1 && !unicode.IsControl(r[0]) {
				s.search += string(r)
				s.cursor = 0
				s.scroll = 0
				s.applyFilter()
			}
		}
	case tea.PasteMsg:
		s.search += msg.String()
		s.cursor = 0
		s.scroll = 0
		s.applyFilter()
	}
	return s, nil
}

func (s *ModelSelectScreen) ensureVisible() {
	_, maxVis := s.entryBounds()
	if s.cursor < s.scroll {
		s.scroll = s.cursor
	}
	if s.cursor >= s.scroll+maxVis && maxVis > 0 {
		s.scroll = s.cursor - maxVis + 1
	}
}

func (s *ModelSelectScreen) entryBounds() (arrowCount, maxEntries int) {
	bodyFixed := 7
	avail := s.height - bodyFixed - 5
	if avail < 1 {
		return 0, 0
	}
	if s.scroll > 0 {
		arrowCount = 1
	}
	if s.scroll+avail-arrowCount < len(s.filtered) {
		arrowCount++
		if arrowCount > 2 {
			arrowCount = 2
		}
	}
	maxEntries = avail - arrowCount
	if maxEntries < 0 {
		maxEntries = 0
	}
	return arrowCount, maxEntries
}

func (s *ModelSelectScreen) View() string {
	innerW := s.width - 2

	title := fmt.Sprintf("%-*s%s", innerW-4, "Select model", "esc")

	searchText := s.search
	if searchText == "" {
		searchText = styles.HintStyle.Render("search your model")
	}
	searchLine := fmt.Sprintf("Search: %s", searchText)

	header := styles.Subtitle.Render(" All Models")

	_, maxEntries := s.entryBounds()
	if s.scroll+maxEntries > len(s.filtered) {
		s.scroll = max(0, len(s.filtered)-maxEntries)
	}

	end := min(s.scroll+maxEntries, len(s.filtered))

	var bodyLines []string
	bodyLines = append(bodyLines, title)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, searchLine)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, header)

	for i := s.scroll; i < end; i++ {
		bodyLines = append(bodyLines, s.renderModelLine(s.filtered[i], innerW, i == s.cursor))
	}

	if s.scroll > 0 {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll))))
	}
	if end < len(s.filtered) {
		remaining := len(s.filtered) - end
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↓ %d more", remaining))))
	}

	bodyLines = append(bodyLines, "")

	bottomText := "add provider ctrl+p  Favorite ctrl+f"
	bottomW := lipgloss.Width(bottomText)
	pad := (innerW - bottomW) / 2
	if pad < 0 {
		pad = 0
	}
	bodyLines = append(bodyLines, fmt.Sprintf("%s%s", strings.Repeat(" ", pad), styles.HintStyle.Render(bottomText)))

	body := strings.Join(bodyLines, "\n")
	return styles.DialogBox(s.width, body)
}

func (s *ModelSelectScreen) renderModelLine(m domainmodel.Model, width int, isCursor bool) string {
	name := m.DisplayName
	if name == "" {
		name = m.ModelID
	}

	var freeTag string
	if m.PricingInput == 0 && m.PricingOut == 0 {
		freeTag = styles.Active.Render("Free")
	}

	var cursor string
	if isCursor {
		cursor = styles.Active.Render(">")
		name = styles.ValueStyle.Render(name)
	} else {
		cursor = " "
		name = lipgloss.NewStyle().Foreground(lipgloss.Color("250")).Render(name)
	}

	nameW := lipgloss.Width(name)
	tagW := lipgloss.Width(freeTag)
	padW := width - nameW - tagW - 2
	if padW < 0 {
		padW = 0
	}

	return fmt.Sprintf("%s %s%s%s", cursor, name, strings.Repeat(" ", padW), freeTag)
}
