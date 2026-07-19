package screens

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/styles"
)

var (
	provConnected = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Render("●")
	provDisconn   = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Render("○")
	provAuthFail  = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("●")
	provNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	provURLStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	provDefStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render(" (active)")
)

type ProviderItem struct {
	Name     string
	URL      string
	Status   string
	Latency  int64
	IsActive bool
}

type ProviderListScreen struct {
	width     int
	height    int
	providers []ProviderItem
	cursor    int
	scroll    int
	done      bool
	result    string
	onDelete  func(name string) string
	onSelect  func(name string) string
	onEdit    func(name string) string
	onRefresh func()
}

func NewProviderListScreen() *ProviderListScreen {
	return &ProviderListScreen{
		width:  80,
		height: 24,
	}
}

func (s *ProviderListScreen) SetSize(w, h int) {
	s.width = w
	s.height = h
	if s.onRefresh != nil {
		s.onRefresh()
	}
}

func (s *ProviderListScreen) SetProviders(providers []ProviderItem) {
	s.providers = providers
	if s.cursor >= len(s.providers) {
		s.cursor = len(s.providers) - 1
	}
}

func (s *ProviderListScreen) OnDelete(fn func(name string) string) {
	s.onDelete = fn
}

func (s *ProviderListScreen) OnSelect(fn func(name string) string) {
	s.onSelect = fn
}

func (s *ProviderListScreen) OnEdit(fn func(name string) string) {
	s.onEdit = fn
}

func (s *ProviderListScreen) OnRefresh(fn func()) {
	s.onRefresh = fn
}

func (s *ProviderListScreen) Done() bool     { return s.done }
func (s *ProviderListScreen) Result() string { return s.result }

func (s *ProviderListScreen) Update(msg tea.Msg) (DialogScreen, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			s.done = true
		case "enter":
			if len(s.providers) > 0 && s.cursor >= 0 && s.cursor < len(s.providers) {
				p := s.providers[s.cursor]
				if s.onSelect != nil {
					s.result = s.onSelect(p.Name)
				} else {
					s.result = fmt.Sprintf("Selected provider '%s'", p.Name)
				}
				s.done = true
			}
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
				if s.cursor < s.scroll {
					s.scroll = s.cursor
				}
			}
		case "down", "j":
			if s.cursor < len(s.providers)-1 {
				s.cursor++
				maxVis := s.height - 10
				if maxVis < 1 {
					maxVis = 1
				}
				if s.cursor >= s.scroll+maxVis {
					s.scroll = s.cursor - maxVis + 1
				}
			}
		case "d", "D":
			if len(s.providers) > 0 && s.cursor >= 0 && s.cursor < len(s.providers) {
				p := s.providers[s.cursor]
				if s.onDelete != nil {
					s.onDelete(p.Name)
				}
				s.providers = append(s.providers[:s.cursor], s.providers[s.cursor+1:]...)
				if s.cursor >= len(s.providers) {
					s.cursor = max(0, len(s.providers)-1)
				}
				if s.scroll >= len(s.providers) {
					s.scroll = max(0, len(s.providers)-1)
				}
			}
		case "e", "E":
			if len(s.providers) > 0 && s.cursor >= 0 && s.cursor < len(s.providers) {
				p := s.providers[s.cursor]
				if s.onEdit != nil {
					s.result = s.onEdit(p.Name)
				}
			}
		}
	}
	return s, nil
}

func (s *ProviderListScreen) View() string {
	innerW := s.width - 2

	title := fmt.Sprintf("%-*s%s", innerW-4, "Provider Management", "esc")

	if len(s.providers) == 0 {
		body := fmt.Sprintf("%s\n%s\n%s",
			title,
			styles.DialogSep(innerW),
			styles.HintStyle.Render("No providers configured."),
		)
		return styles.DialogBox(s.width, body)
	}

	maxVis := s.height - 10
	if maxVis < 1 {
		maxVis = 1
	}
	if s.scroll+maxVis > len(s.providers) {
		s.scroll = max(0, len(s.providers)-maxVis)
	}
	end := min(s.scroll+maxVis, len(s.providers))

	var bodyLines []string
	bodyLines = append(bodyLines, title)
	bodyLines = append(bodyLines, styles.DialogSep(innerW))
	bodyLines = append(bodyLines, "")

	for i := s.scroll; i < end; i++ {
		p := s.providers[i]
		icon := provDisconn
		switch p.Status {
		case "connected":
			icon = provConnected
		case "auth_failed":
			icon = provAuthFail
		}

		active := ""
		if p.IsActive {
			active = provDefStyle
		}

		cursor := "  "
		if i == s.cursor {
			cursor = styles.Active.Render("> ")
		}

		latency := ""
		if p.Latency > 0 {
			latency = fmt.Sprintf(" %dms", p.Latency)
		}

		urlStr := p.URL
		urlMax := innerW - 10
		if len([]rune(urlStr)) > urlMax {
			urlStr = string([]rune(urlStr)[:urlMax]) + "..."
		}

		nameStr := p.Name
		nameMax := innerW - 8
		if len([]rune(nameStr)) > nameMax {
			nameStr = string([]rune(nameStr)[:nameMax]) + "..."
		}

		bodyLines = append(bodyLines, fmt.Sprintf(" %s%s%s%s", cursor, icon, provNameStyle.Render(nameStr), active))
		bodyLines = append(bodyLines, fmt.Sprintf("    %s%s", provURLStyle.Render(urlStr), latency))
		bodyLines = append(bodyLines, "")
	}

	if s.scroll > 0 {
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↑ %d more", s.scroll))))
	}
	if end < len(s.providers) {
		remaining := len(s.providers) - end
		bodyLines = append(bodyLines, fmt.Sprintf("  %s", styles.HintStyle.Render(fmt.Sprintf("↓ %d more", remaining))))
	}

	bodyLines = append(bodyLines, "")
	hintText := "esc: Close  ↵: Select  e: Edit  d: Delete"
	bodyLines = append(bodyLines, fmt.Sprintf("%*s", innerW, styles.HintStyle.Render(hintText)))

	body := strings.Join(bodyLines, "\n")
	return styles.DialogBox(s.width, body)
}
