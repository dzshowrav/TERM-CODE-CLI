package screens

import (
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"

	"termcode/internal/adapters/tui/components"
)

type scrollApplyMsg struct{}

const (
	scrollDebounceInterval = 50 * time.Millisecond
	scrollSingleGap        = 200 * time.Millisecond
	autoCollapseDelay      = 2 * time.Second
	toolAnimTickInterval   = 150 * time.Millisecond
)

type ToolQueuedMsg struct {
	Index int
	Name  string
	Args  string
}

type ToolInitializingMsg struct {
	Index int
	Name  string
}

type ToolConnectingMsg struct {
	Index int
	Name  string
}

type ToolStartedMsg struct {
	Index int
	Name  string
}

type ToolOutputMsg struct {
	Index  int
	Name   string
	Output string
}

type ToolCompletedMsg struct {
	Index    int
	Name     string
	Output   string
	Status   string
	Error    string
	Duration int64
}

type AutoCollapseMsg struct {
	Index int
}

type ToolAnimTick struct{}

type ChatScreen struct {
	width            int
	height           int
	viewport         *components.Viewport
	modelName        string
	messages         []ChatMessage
	scrolledAway     bool
	scrollDelta      int
	scrollPending    bool
	lastScrollTime   time.Time
	toolCards        []*components.ToolCard
	focusedCard      int
	animTick         int
	showThinking     bool
	thinkingExpanded bool
	thinkingStart    time.Time
	thinkingDur      time.Duration
	toolCount        int
	searchMode       bool
	searchQuery      string
	filterMode       int // 0=all 1=running 2=completed 3=failed
	totalDuration    int64
	inputTokens      int
	outputTokens     int
}

type ChatMessage struct {
	Role    string
	Content string
	ToolID  string
}

var (
	userStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).SetString("You")
	assistantStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).SetString("AI")
	systemStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).SetString("System")
	thinkingStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("141")).SetString("Thinking")
	metricsStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	timelineStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
)

func NewChatScreen() *ChatScreen {
	return &ChatScreen{
		width:       80,
		height:      24,
		viewport:    components.NewViewport(76, 1),
		modelName:   "none",
		messages:    []ChatMessage{},
		focusedCard: -1,
	}
}

func (s *ChatScreen) SetViewportSize(w, h int) {
	s.width = w + 4
	s.viewport.SetSize(w, h)
	s.viewport.GotoBottom()
}

func (s *ChatScreen) SetModel(name string) {
	s.modelName = name
}

func (s *ChatScreen) AddMessage(role, content string) {
	s.messages = append(s.messages, ChatMessage{Role: role, Content: content})
	s.renderMessages()
}

func (s *ChatScreen) AddToolMessage(role, content, toolID string) {
	s.messages = append(s.messages, ChatMessage{Role: role, Content: content, ToolID: toolID})
	s.renderMessages()
}

func (s *ChatScreen) ShowThinking() {
	s.showThinking = true
	s.thinkingStart = time.Now()
	s.renderMessages()
}

func (s *ChatScreen) HideThinking() {
	s.showThinking = false
	s.renderMessages()
}

func (s *ChatScreen) AddToolCard(name, args string) {
	card := components.NewToolCard(name)
	card.SetArgs(args)
	card.SetWidth(s.viewport.Height())
	s.toolCards = append(s.toolCards, card)
	s.renderMessages()
}

func (s *ChatScreen) UpdateToolCard(index int, fn func(*components.ToolCard)) {
	if index >= 0 && index < len(s.toolCards) {
		fn(s.toolCards[index])
		s.renderMessages()
	}
}

func (s *ChatScreen) ClearToolCards() {
	s.toolCards = nil
	s.showThinking = false
	s.thinkingStart = time.Time{}
	s.thinkingExpanded = false
	s.toolCount = 0
	s.totalDuration = 0
	s.inputTokens = 0
	s.outputTokens = 0
	s.focusedCard = -1
	s.renderMessages()
}

func (s *ChatScreen) persistToolResults() {
	for _, card := range s.toolCards {
		dur := fmt.Sprintf("%dms", card.Duration())
		summary := fmt.Sprintf("  %s  %s  %s",
			card.Name(),
			card.Status(),
			dur,
		)
		s.messages = append(s.messages, ChatMessage{
			Role:    "tool",
			Content: summary,
		})
	}
}

func (s *ChatScreen) SetToolMetrics(count int, totalDuration int64, inputTokens, outputTokens int) {
	s.toolCount = count
	s.totalDuration = totalDuration
	s.inputTokens = inputTokens
	s.outputTokens = outputTokens
	s.persistToolResults()
}

func (s *ChatScreen) ToolCount() int {
	return s.toolCount
}

func (s *ChatScreen) TotalDuration() int64 {
	return s.totalDuration
}

func (s *ChatScreen) renderMessages() {
	var lines []string
	for _, msg := range s.messages {
		roleLabel := ""
		switch msg.Role {
		case "user":
			roleLabel = userStyle.String()
		case "assistant":
			roleLabel = assistantStyle.String()
		case "tool":
			roleLabel = lipgloss.NewStyle().Foreground(lipgloss.Color("141")).Bold(true).Render("Tool")
		default:
			roleLabel = systemStyle.String()
		}
		lines = append(lines, roleLabel)
		lines = append(lines, msg.Content)
		lines = append(lines, "")
	}

	if s.showThinking {
		if !s.thinkingStart.IsZero() {
			elapsed := time.Since(s.thinkingStart)
			dur := fmt.Sprintf("%.1fs", elapsed.Seconds())
			thinkingLine := fmt.Sprintf("  ◐ thinking... (%s)", dur)
			lines = append(lines, thinkingStyle.String())
			lines = append(lines, thinkingLine)
			lines = append(lines, "")
		} else {
			lines = append(lines, thinkingStyle.String())
			lines = append(lines, "  ◐ thinking...")
			lines = append(lines, "")
		}
	} else if s.thinkingDur > 0 {
		dur := fmt.Sprintf("%.1fs", s.thinkingDur.Seconds())
		expandIcon := "▶"
		if s.thinkingExpanded {
			expandIcon = "▼"
		}
		focusIndicator := ""
		if s.focusedCard < 0 {
			focusIndicator = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("▸") + " "
		}
		thinkingLine := fmt.Sprintf("  %s%s %s  %s", focusIndicator, expandIcon, thinkingStyle.String(), dur)
		lines = append(lines, thinkingLine)
		if s.thinkingExpanded {
			lines = append(lines, "  ◐ thinking...")
		}
		lines = append(lines, "")
	}

	filterLabel := ""
	switch s.filterMode {
	case 1:
		filterLabel = " [running]"
	case 2:
		filterLabel = " [completed]"
	case 3:
		filterLabel = " [failed]"
	}
	searchLabel := ""
	if s.searchMode {
		if s.searchQuery != "" {
			searchLabel = fmt.Sprintf(" search:/%s/", s.searchQuery)
		} else {
			searchLabel = " search:_"
		}
	}
	if filterLabel != "" || (s.searchMode && s.searchQuery != "") {
		infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
		lines = append(lines, infoStyle.Render(fmt.Sprintf(" Filter:%s%s", filterLabel, searchLabel)))
	} else if s.searchMode {
		infoStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
		lines = append(lines, infoStyle.Render(" Press any key to search by name..."))
	}

	renderedAny := false
	for i, card := range s.toolCards {
		if s.filterMode > 0 {
			status := card.Status()
			switch s.filterMode {
			case 1:
				if !card.IsRunning() {
					continue
				}
			case 2:
				if status != components.ToolCompleted {
					continue
				}
			case 3:
				if status != components.ToolFailed {
					continue
				}
			}
		}
		if s.searchMode && s.searchQuery != "" {
			if !strings.Contains(strings.ToLower(card.Name()), strings.ToLower(s.searchQuery)) {
				continue
			}
		}
		if renderedAny || s.thinkingDur > 0 {
			lines = append(lines, timelineStyle.Render("  ↓"))
		}
		card.SetFocused(i == s.focusedCard)
		cardLines := strings.Split(card.View(), "\n")
		lines = append(lines, cardLines...)
		lines = append(lines, "")
		renderedAny = true
	}

	if s.toolCount > 0 && len(s.toolCards) > 0 {
		last := s.toolCards[len(s.toolCards)-1]
		if last.Status() == components.ToolCompleted || last.Status() == components.ToolFailed {
			dur := fmt.Sprintf("%.1fs", float64(s.totalDuration)/1000)
			metrics := fmt.Sprintf("  Tools: %d  Duration: %s", s.toolCount, dur)
			if s.inputTokens > 0 || s.outputTokens > 0 {
				metrics += fmt.Sprintf("  Tokens: %d (in: %d, out: %d)", s.inputTokens+s.outputTokens, s.inputTokens, s.outputTokens)
			}
			lines = append(lines, metricsStyle.Render(metrics))
			lines = append(lines, "")
		}
	}

	s.viewport.SetContent(lines)
	if !s.scrolledAway {
		s.viewport.GotoBottom()
	}
}

func (s *ChatScreen) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case scrollApplyMsg:
		if s.scrollDelta > 0 {
			s.viewport.ScrollDown(s.scrollDelta)
		} else if s.scrollDelta < 0 {
			s.viewport.ScrollUp(-s.scrollDelta)
		}
		s.scrolledAway = !s.viewport.AtBottom()
		s.scrollDelta = 0
		s.scrollPending = false
		return nil

	case AutoCollapseMsg:
		if msg.Index >= 0 && msg.Index < len(s.toolCards) {
			s.toolCards[msg.Index].SetExpanded(false)
			s.renderMessages()
		}
		return nil

	case ToolQueuedMsg:
		card := components.NewToolCard(msg.Name)
		card.SetArgs(msg.Args)
		card.SetWidth(s.width - 8)
		s.toolCards = append(s.toolCards, card)
		if s.showThinking && !s.thinkingStart.IsZero() {
			s.thinkingDur = time.Since(s.thinkingStart)
		}
		s.showThinking = false
		s.renderMessages()
		return tea.Tick(toolAnimTickInterval, func(time.Time) tea.Msg {
			return ToolAnimTick{}
		})

	case ToolInitializingMsg:
		idx := msg.Index
		if idx < 0 || idx >= len(s.toolCards) {
			idx = len(s.toolCards) - 1
		}
		s.UpdateToolCard(idx, func(c *components.ToolCard) {
			c.SetStatus(components.ToolInitializing)
		})
		return nil

	case ToolConnectingMsg:
		idx := msg.Index
		if idx < 0 || idx >= len(s.toolCards) {
			idx = len(s.toolCards) - 1
		}
		s.UpdateToolCard(idx, func(c *components.ToolCard) {
			c.SetStatus(components.ToolConnecting)
		})
		return nil

	case ToolStartedMsg:
		idx := len(s.toolCards) - 1
		if msg.Index >= 0 && msg.Index < len(s.toolCards) {
			idx = msg.Index
		}
		s.UpdateToolCard(idx, func(c *components.ToolCard) {
			c.SetStatus(components.ToolRunning)
		})
		return tea.Tick(toolAnimTickInterval, func(time.Time) tea.Msg {
			return ToolAnimTick{}
		})

	case ToolOutputMsg:
		s.UpdateToolCard(len(s.toolCards)-1, func(c *components.ToolCard) {
			c.SetOutput(msg.Output)
		})
		return nil

	case ToolCompletedMsg:
		idx := len(s.toolCards) - 1
		s.UpdateToolCard(idx, func(c *components.ToolCard) {
			if msg.Status == "completed" {
				c.SetStatus(components.ToolCompleted)
				c.SetOutput(msg.Output)
			} else {
				c.SetStatus(components.ToolFailed)
				c.SetOutput(msg.Output)
				c.SetError(msg.Error)
			}
			c.SetDuration(msg.Duration)
		})
		s.toolCount++
		s.totalDuration += msg.Duration
		s.renderMessages()
		return tea.Tick(autoCollapseDelay, func(time.Time) tea.Msg {
			return AutoCollapseMsg{Index: idx}
		})

	case ToolAnimTick:
		s.animTick++
		animated := false
		for _, c := range s.toolCards {
			if c.IsRunning() {
				c.SetAnimFrame(s.animTick)
				animated = true
			}
		}
		if animated {
			s.renderMessages()
			return tea.Tick(toolAnimTickInterval, func(time.Time) tea.Msg {
				return ToolAnimTick{}
			})
		}
		return nil

	case tea.MouseClickMsg:
		m := msg.Mouse()
		if m.Button == tea.MouseLeft && len(s.toolCards) > 0 {
			if s.focusedCard >= 0 {
				card := s.toolCards[s.focusedCard]
				if card.Expanded() && card.OutputLines() > 15 {
					card.ToggleShowAll()
				} else {
					card.ToggleExpanded()
				}
			} else {
				card := s.toolCards[len(s.toolCards)-1]
				card.ToggleExpanded()
			}
			s.renderMessages()
		}
		return nil

	case tea.MouseWheelMsg:
		m := msg.Mouse()
		if m.Button == tea.MouseWheelUp {
			if s.applyImmediate() {
				s.viewport.ScrollUp(3)
				s.scrolledAway = !s.viewport.AtBottom()
				return nil
			}
			s.scrollDelta -= 3
		} else if m.Button == tea.MouseWheelDown {
			if s.applyImmediate() {
				s.viewport.ScrollDown(3)
				s.scrolledAway = !s.viewport.AtBottom()
				return nil
			}
			s.scrollDelta += 3
		}
		if !s.scrollPending {
			s.scrollPending = true
			return tea.Tick(scrollDebounceInterval, func(time.Time) tea.Msg {
				return scrollApplyMsg{}
			})
		}
		return nil

	case tea.KeyMsg:
		if s.focusedCard >= 0 && len(s.toolCards) > 0 {
			return s.handleCardFocusKey(msg.String())
		}
		return s.handleScrollKey(msg.String())
	}
	return nil
}

func (s *ChatScreen) handleCardFocusKey(key string) tea.Cmd {
	if s.searchMode {
		if key == "esc" || key == "enter" {
			s.searchMode = false
			s.searchQuery = ""
			s.renderMessages()
			return nil
		}
		if len(key) == 1 && key >= " " {
			s.searchQuery = key
			s.renderMessages()
		}
		return nil
	}

	if key == "esc" {
		if s.searchMode || s.filterMode > 0 {
			s.searchMode = false
			s.searchQuery = ""
			s.filterMode = 0
			s.renderMessages()
		}
		s.focusedCard = -1
		s.renderMessages()
		return nil
	}

	switch key {
	case "up":
		s.focusedCard--
		if s.focusedCard < -1 {
			s.focusedCard = len(s.toolCards) - 1
		} else if s.focusedCard < 0 && s.thinkingDur == 0 {
			s.focusedCard = len(s.toolCards) - 1
		}
		return nil

	case "down":
		s.focusedCard++
		if s.focusedCard >= len(s.toolCards) {
			if s.thinkingDur > 0 {
				s.focusedCard = -1
			} else {
				s.focusedCard = 0
			}
		}
		return nil

	case "enter":
		if s.focusedCard < 0 {
			s.thinkingExpanded = !s.thinkingExpanded
		} else {
			card := s.toolCards[s.focusedCard]
			if card.Expanded() && card.OutputLines() > 15 {
				card.ToggleShowAll()
			} else {
				card.ToggleExpanded()
			}
		}
		s.renderMessages()
		return nil

	case " ":
		if s.focusedCard < 0 {
			s.thinkingExpanded = false
		} else {
			s.toolCards[s.focusedCard].SetExpanded(false)
		}
		s.renderMessages()
		return nil

	case "tab":
		s.focusedCard = -1
		s.renderMessages()
		return nil

	default:
		return nil
	}
}

func (s *ChatScreen) handleScrollKey(key string) tea.Cmd {
	switch key {
	case "tab":
		if len(s.toolCards) > 0 {
			s.focusedCard = 0
			s.renderMessages()
		}
		return nil

	case "enter":
		if len(s.toolCards) > 0 {
			lastIdx := len(s.toolCards) - 1
			s.toolCards[lastIdx].ToggleExpanded()
			s.renderMessages()
		}
		return nil

	case " ":
		if len(s.toolCards) > 0 {
			for _, c := range s.toolCards {
				c.SetExpanded(false)
			}
			s.renderMessages()
		}
		return nil

	case "esc":
		if s.searchMode || s.filterMode > 0 {
			s.searchMode = false
			s.searchQuery = ""
			s.filterMode = 0
			s.renderMessages()
		}
		return nil

	case "/":
		s.searchMode = !s.searchMode
		if !s.searchMode {
			s.searchQuery = ""
		}
		s.focusedCard = -1
		s.renderMessages()
		return nil

	case "f":
		s.filterMode = (s.filterMode + 1) % 4
		s.focusedCard = -1
		s.renderMessages()
		return nil

	case "pgup", "ctrl+b":
		if s.applyImmediate() {
			s.viewport.ScrollUp(s.viewport.Height() / 2)
			s.scrolledAway = !s.viewport.AtBottom()
			return nil
		}
		s.scrollDelta -= s.viewport.Height() / 2

	case "pgdown", "ctrl+f":
		if s.applyImmediate() {
			s.viewport.ScrollDown(s.viewport.Height() / 2)
			s.scrolledAway = !s.viewport.AtBottom()
			return nil
		}
		s.scrollDelta += s.viewport.Height() / 2

	case "up":
		if s.applyImmediate() {
			s.viewport.ScrollUp(1)
			s.scrolledAway = !s.viewport.AtBottom()
			return nil
		}
		s.scrollDelta--

	case "down":
		if s.applyImmediate() {
			s.viewport.ScrollDown(1)
			s.scrolledAway = !s.viewport.AtBottom()
			return nil
		}
		s.scrollDelta++

	case "home":
		for !s.viewport.AtTop() {
			s.viewport.ScrollUp(1)
		}
		s.scrolledAway = true
		return nil

	case "end":
		s.viewport.GotoBottom()
		s.scrolledAway = false
		return nil

	default:
		return nil
	}

	if !s.scrollPending {
		s.scrollPending = true
		return tea.Tick(scrollDebounceInterval, func(time.Time) tea.Msg {
			return scrollApplyMsg{}
		})
	}
	return nil
}

func (s *ChatScreen) applyImmediate() bool {
	now := time.Now()
	gap := now.Sub(s.lastScrollTime)
	s.lastScrollTime = now
	return gap > scrollSingleGap
}

func (s *ChatScreen) View() string {
	return s.viewport.View()
}
