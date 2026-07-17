package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	modFav       = lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Render("★")
	modNoFav     = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render("☆")
	modActive    = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render(" ✓")
	modNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	modCatStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	modCtxStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	modProvStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("141"))
	modSepStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
)

type ModelListItem struct {
	ModelID      string
	DisplayName  string
	ProviderName string
	Category     string
	ContextSize  int
	IsFavorite   bool
	IsActive     bool
}

type ModelListScreen struct {
	width       int
	height      int
	models      []ModelListItem
	groupByProv bool
	done        bool
	scroll      int
}

func NewModelListScreen() *ModelListScreen {
	return &ModelListScreen{
		width:       80,
		height:      24,
		groupByProv: true,
	}
}

func (s *ModelListScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
}

func (s *ModelListScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "up":
			if s.scroll > 0 {
				s.scroll--
			}
		case "down":
			s.scroll++
		}
	}
	return s, nil
}

func (s *ModelListScreen) Done() bool     { return s.done }
func (s *ModelListScreen) Result() string { return "" }

func (s *ModelListScreen) SetModels(models []ModelListItem) {
	s.models = models
}

func (s *ModelListScreen) grouped() map[string][]ModelListItem {
	groups := make(map[string][]ModelListItem)
	for _, m := range s.models {
		groups[m.ProviderName] = append(groups[m.ProviderName], m)
	}
	return groups
}

func (s *ModelListScreen) View() string {
	header := styles.H1.Render("All Models")
	sep := styles.SeparatorLine(s.width)

	if len(s.models) == 0 {
		empty := lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("No models configured.")
		hint := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render("Use /provider sync to load models.")
		return styles.Content(s.width, fmt.Sprintf("%s\n%s\n%s\n%s", header, sep, empty, hint))
	}

	var lines []string

	if s.groupByProv {
		for prov, models := range s.grouped() {
			lines = append(lines, modSepStyle.Render("── "+prov+" ──"))
			for _, m := range models {
				lines = append(lines, s.renderLine(m))
			}
			lines = append(lines, "")
		}
	} else {
		for _, m := range s.models {
			lines = append(lines, s.renderLine(m))
		}
	}

	return styles.Content(s.width, fmt.Sprintf("%s\n%s\n%s", header, sep, strings.Join(lines, "\n")))
}

func (s *ModelListScreen) renderLine(m ModelListItem) string {
	fav := modNoFav
	if m.IsFavorite {
		fav = modFav
	}

	active := ""
	if m.IsActive {
		active = modActive
	}

	ctx := ""
	if m.ContextSize > 0 {
		ctx = fmt.Sprintf(" · %dK", m.ContextSize/1024)
	}

	return fmt.Sprintf(
		" %s %s%s %s%s%s",
		fav,
		modNameStyle.Render(m.DisplayName),
		active,
		modProvStyle.Render(m.ProviderName),
		modCatStyle.Render(" · "+m.Category),
		modCtxStyle.Render(ctx),
	)
}
