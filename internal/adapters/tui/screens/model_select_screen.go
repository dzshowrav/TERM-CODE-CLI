package screens

import (
	"fmt"
	"sort"
	"strings"
	"unicode"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
	domainmodel "termcode/internal/domain/model"
)

type modelSelectItem struct {
	isHeader bool
	title    string
	model    domainmodel.Model
}

type sortMode int

const (
	sortByName sortMode = iota
	sortByProvider
	sortByContext
)

type ModelSelectScreen struct {
	width    int
	height   int
	models   []domainmodel.Model
	items    []modelSelectItem
	search   string
	cursor   int
	scroll   int
	done     bool
	result   string
	onSelect func(modelID string) string
	sort     sortMode
	showHelp bool
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

// isFreeByName checks ONLY the model ID / display name for the word "free".
// Pricing fields default to 0 for all models (not set), so they cannot be
// used as a free indicator — relying on them tagged everything as free.
func isFreeByName(m domainmodel.Model) bool {
	id := strings.ToLower(m.ModelID)
	name := strings.ToLower(m.DisplayName)
	return strings.Contains(id, "free") || strings.Contains(name, "free")
}

// nextSelectable moves in direction dir (+1 / -1) starting at start, skipping
// over section-header items so the cursor always lands on a model row.
func (s *ModelSelectScreen) nextSelectable(start, dir int) int {
	curr := start
	for curr >= 0 && curr < len(s.items) {
		if !s.items[curr].isHeader {
			return curr
		}
		curr += dir
	}
	return -1
}

func (s *ModelSelectScreen) applyFilter() {
	var filtered []domainmodel.Model
	if s.search == "" {
		filtered = s.models
	} else {
		lower := strings.ToLower(s.search)
		for _, m := range s.models {
			if strings.Contains(strings.ToLower(m.DisplayName), lower) ||
				strings.Contains(strings.ToLower(m.ModelID), lower) ||
				strings.Contains(strings.ToLower(m.ProviderID), lower) {
				filtered = append(filtered, m)
			}
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		switch s.sort {
		case sortByName:
			ni, nj := strings.ToLower(filtered[i].DisplayName), strings.ToLower(filtered[j].DisplayName)
			if ni != nj {
				return ni < nj
			}
			return filtered[i].ModelID < filtered[j].ModelID
		case sortByProvider:
			pi, pj := strings.ToLower(filtered[i].ProviderID), strings.ToLower(filtered[j].ProviderID)
			if pi != pj {
				return pi < pj
			}
			return filtered[i].ModelID < filtered[j].ModelID
		case sortByContext:
			if filtered[i].MaxContext != filtered[j].MaxContext {
				return filtered[i].MaxContext > filtered[j].MaxContext
			}
			return filtered[i].ModelID < filtered[j].ModelID
		default:
			return filtered[i].ModelID < filtered[j].ModelID
		}
	})

	// Split into free (by name) vs all others
	var freeModels []domainmodel.Model
	var otherModels []domainmodel.Model
	for _, m := range filtered {
		if isFreeByName(m) {
			freeModels = append(freeModels, m)
		} else {
			otherModels = append(otherModels, m)
		}
	}

	s.items = nil
	if len(freeModels) > 0 {
		s.items = append(s.items, modelSelectItem{isHeader: true, title: "Free Models"})
		for _, m := range freeModels {
			s.items = append(s.items, modelSelectItem{model: m})
		}
	}
	if len(otherModels) > 0 {
		// Only show "All Models" header when there are also free models shown
		if len(freeModels) > 0 {
			s.items = append(s.items, modelSelectItem{isHeader: true, title: "All Models"})
		}
		for _, m := range otherModels {
			s.items = append(s.items, modelSelectItem{model: m})
		}
	}

	// Keep cursor on a valid selectable item
	if len(s.items) > 0 {
		if s.cursor >= len(s.items) {
			s.cursor = len(s.items) - 1
		}
		if found := s.nextSelectable(s.cursor, 1); found != -1 {
			s.cursor = found
		} else if found := s.nextSelectable(s.cursor, -1); found != -1 {
			s.cursor = found
		} else {
			s.cursor = 0
		}
	} else {
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
			if len(s.items) > 0 && s.cursor >= 0 && s.cursor < len(s.items) {
				item := s.items[s.cursor]
				if !item.isHeader {
					s.result = s.onSelect(item.model.ModelID)
					s.done = true
				}
			}
			return s, nil
		case "s":
			s.sort = (s.sort + 1) % 3
			s.cursor = 0
			s.scroll = 0
			s.applyFilter()
		case "?":
			s.showHelp = !s.showHelp
		case "up", "k":
			if s.cursor > 0 {
				if found := s.nextSelectable(s.cursor-1, -1); found != -1 {
					s.cursor = found
					s.ensureVisible()
				}
			}
		case "down", "j":
			if s.cursor < len(s.items)-1 {
				if found := s.nextSelectable(s.cursor+1, 1); found != -1 {
					s.cursor = found
					s.ensureVisible()
				}
			}
		case "backspace":
			if len(s.search) > 0 {
				s.search = s.search[:len(s.search)-1]
				s.cursor = 0
				s.scroll = 0
				s.applyFilter()
			}
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
	if maxVis > 0 && s.cursor >= s.scroll+maxVis {
		s.scroll = s.cursor - maxVis + 1
	}
}

func (s *ModelSelectScreen) entryBounds() (arrowCount, maxEntries int) {
	// Fixed body lines: title(1) sep(1) search(1) sep(1) = 4
	// Plus dialog borders(2) + outer overhead(3) + hint bottom(2) = 7
	// Total fixed overhead = 11
	avail := s.height - 11
	if avail < 1 {
		return 0, 1
	}
	if s.scroll > 0 {
		arrowCount = 1
	}
	if s.scroll+avail-arrowCount < len(s.items) {
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
	sortLabels := map[sortMode]string{
		sortByName:     "Name",
		sortByProvider: "Provider",
		sortByContext:  "Context",
	}
	sortInfo := fmt.Sprintf(" [Sort: %s]", sortLabels[s.sort])
	searchLine := fmt.Sprintf("Search: %s%s", searchText, lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(sortInfo))

	_, maxEntries := s.entryBounds()
	if s.scroll+maxEntries > len(s.items) {
		s.scroll = max(0, len(s.items)-maxEntries)
	}
	end := min(s.scroll+maxEntries, len(s.items))

	var bodyLines []string
	bodyLines = append(bodyLines, title)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, searchLine)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))

	if s.scroll > 0 {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll))))
	}

	for i := s.scroll; i < end; i++ {
		item := s.items[i]
		if item.isHeader {
			bodyLines = append(bodyLines, styles.Subtitle.Render(" "+item.title))
		} else {
			bodyLines = append(bodyLines, s.renderModelLine(item.model, innerW, i == s.cursor))
		}
	}

	if end < len(s.items) {
		remaining := len(s.items) - end
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↓ %d more", remaining))))
	}

	if s.showHelp {
		helpLines := []string{
			"",
			styles.Subtitle.Render(" Keyboard Shortcuts"),
			" type: Filter models by name/ID/provider",
			" s:    Cycle sort (Name / Provider / Context)",
			" ?:    Toggle this help",
			" enter: Select model",
			" esc:  Close",
		}
		bodyLines = append(bodyLines, helpLines...)
	}

	bodyLines = append(bodyLines, "")

	bottomText := "type:search  s:sort  enter:select  ?:help  esc:close"
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
	if isFreeByName(m) {
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
