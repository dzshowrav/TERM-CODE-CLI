package screens

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	selHeader  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("39"))
	selActive  = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	selItem    = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	selSub     = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	selSection = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

type ModelSelectorItem struct {
	ID           string
	DisplayName  string
	Category     string
	ProviderName string
	ContextSize  int
	IsActive     bool
	Section      string
}

type ModelSelector struct {
	width    int
	height   int
	items    []ModelSelectorItem
	search   string
	selected int
}

func NewModelSelector() *ModelSelector {
	return &ModelSelector{
		width:  80,
		height: 20,
	}
}

func (s *ModelSelector) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ModelSelector) SetItems(items []ModelSelectorItem) {
	s.items = items
}

func (s *ModelSelector) SetSearch(q string) {
	s.search = q
}

func (s *ModelSelector) SetSelected(idx int) {
	s.selected = idx
}

func (s *ModelSelector) filtered() []ModelSelectorItem {
	if s.search == "" {
		return s.items
	}
	q := strings.ToLower(s.search)
	var result []ModelSelectorItem
	for _, item := range s.items {
		if strings.Contains(strings.ToLower(item.DisplayName), q) ||
			strings.Contains(strings.ToLower(item.ID), q) {
			result = append(result, item)
		}
	}
	return result
}

func (s *ModelSelector) View() string {
	header := fmt.Sprintf("Select Model\n%s", styles.SeparatorLine(s.width))
	search := fmt.Sprintf(" %s ", formInputStyle.Render(s.search+"█"))

	var lines []string
	items := s.filtered()

	if len(items) == 0 {
		lines = append(lines, selSub.Render("  No matching models."))
	} else {
		var lastSection string
		for i, item := range items {
			if item.Section != "" && item.Section != lastSection {
				lines = append(lines, "")
				lines = append(lines, selSection.Render("── "+item.Section+" ──"))
				lastSection = item.Section
			}

			prefix := "  "
			style := selItem
			if i == s.selected {
				prefix = "> "
				style = selActive
			}

			activeMark := ""
			if item.IsActive {
				activeMark = selActive.Render(" ✓")
			}

			sub := fmt.Sprintf("%s · %s", item.ProviderName, item.Category)
			if item.ContextSize > 0 {
				sub += fmt.Sprintf(" · %dK", item.ContextSize/1024)
			}

			lines = append(lines, prefix+style.Render(item.DisplayName)+activeMark)
			lines = append(lines, "   "+selSub.Render(sub))
		}
	}

	content := strings.Join(lines, "\n")
	return fmt.Sprintf("%s\n%s\n%s", header, search, content)
}
