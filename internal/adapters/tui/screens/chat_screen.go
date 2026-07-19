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

type ThinkingTickMsg struct{}

type thoughtEntry struct {
	text      string
	dur       time.Duration
	collapsed bool
}

type tlType int

const (
	tlThought tlType = iota
	tlToolCard
)

type timelineItem struct {
	typ        tlType
	thoughtIdx int
	cardIdx    int
}

type ChatScreen struct {
	width          int
	height         int
	viewport       *components.Viewport
	modelName      string
	messages       []ChatMessage
	scrolledAway   bool
	scrollDelta    int
	scrollPending  bool
	lastScrollTime time.Time
	toolCards      []*components.ToolCard
	timeline       []timelineItem
	focusedCard    int
	animTick       int
	showThinking   bool
	thinkingStart  time.Time
	thinkingDur    time.Duration
	toolCount      int
	totalDuration  int64
	inputTokens    int
	outputTokens   int
	responseText   string
	thoughts       []thoughtEntry
	agentName      string
	streamActive   bool
	cachedBase     []string
	mdRenderer     *components.MarkdownRenderer
}

type ChatMessage struct {
	Role      string
	Content   string
	Reasoning string
	ToolID    string
}

var (
	userStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("83")).SetString("You")
	systemStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).SetString("System")
	metricsStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
)

func assistantLabel(model, agent string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Render("\u25cf "+model) +
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(" │ ") +
		lipgloss.NewStyle().Foreground(lipgloss.Color("83")).Render(agent)
}

func NewChatScreen() *ChatScreen {
	return &ChatScreen{
		width:       80,
		height:      24,
		viewport:    components.NewViewport(76, 1),
		modelName:   "none",
		messages:    []ChatMessage{},
		focusedCard: -1,
		mdRenderer:  components.NewMarkdownRenderer(),
	}
}

func (s *ChatScreen) SetViewportSize(w, h int) {
	s.width = w + 4
	s.viewport.SetSize(w, h)
	s.renderMessages()
	if !s.scrolledAway {
		s.viewport.GotoBottom()
	}
}

func (s *ChatScreen) SetModel(name string) {
	s.modelName = name
}

func (s *ChatScreen) SetAgent(name string) {
	s.agentName = name
}

func (s *ChatScreen) AddMessage(role, content string, reasoning ...string) {
	msg := ChatMessage{Role: role, Content: content}
	if len(reasoning) > 0 {
		msg.Reasoning = reasoning[0]
	}
	s.messages = append(s.messages, msg)
	s.cachedBase = nil
	s.renderMessages()
}

func (s *ChatScreen) UpdateMessage(idx int, role, content, reasoning string) {
	if idx < 0 || idx >= len(s.messages) {
		return
	}
	s.messages[idx].Role = role
	s.messages[idx].Content = content
	s.messages[idx].Reasoning = reasoning
	s.cachedBase = nil
	s.renderMessages()
}

func (s *ChatScreen) RemoveMessage(idx int) {
	if idx < 0 || idx >= len(s.messages) {
		return
	}
	s.messages = append(s.messages[:idx], s.messages[idx+1:]...)
	s.cachedBase = nil
	s.renderMessages()
}

func (s *ChatScreen) InsertMessage(idx int, role, content, reasoning string) {
	if idx < 0 || idx > len(s.messages) {
		return
	}
	msg := ChatMessage{Role: role, Content: content, Reasoning: reasoning}
	s.messages = append(s.messages, ChatMessage{})
	copy(s.messages[idx+1:], s.messages[idx:])
	s.messages[idx] = msg
	s.cachedBase = nil
	s.renderMessages()
}

func (s *ChatScreen) AddToolMessage(role, content, toolID string) {
	s.messages = append(s.messages, ChatMessage{Role: role, Content: content, ToolID: toolID})
	s.cachedBase = nil
	s.renderMessages()
}

func (s *ChatScreen) ShowThinking() tea.Cmd {
	s.showThinking = true
	s.thinkingStart = time.Now()
	s.streamActive = false
	s.cachedBase = nil
	s.renderMessages()
	return tea.Tick(100*time.Millisecond, func(time.Time) tea.Msg {
		return ThinkingTickMsg{}
	})
}

func (s *ChatScreen) HideThinking() {
	s.showThinking = false
	s.thinkingDur = 0
	s.streamActive = false
	s.cachedBase = nil
	s.renderMessages()
}

func (s *ChatScreen) AddToolCard(name, args string) {
	card := components.NewToolCard(name)
	card.SetArgs(args)
	card.SetWidth(s.width - 8)
	s.toolCards = append(s.toolCards, card)
	s.cachedBase = nil
	s.renderMessages()
}

func (s *ChatScreen) UpdateToolCard(index int, fn func(*components.ToolCard)) {
	if index >= 0 && index < len(s.toolCards) {
		fn(s.toolCards[index])
		s.cachedBase = nil
		s.renderMessages()
	}
}

func (s *ChatScreen) AppendResponseText(text string) tea.Cmd {
	s.streamActive = true
	s.responseText += text
	s.renderMessages()
	if !s.scrolledAway {
		s.viewport.GotoBottom()
	}
	return nil
}

func (s *ChatScreen) SetResponse(text string) {
	s.responseText = text
	s.cachedBase = nil
	s.renderMessages()
}

func (s *ChatScreen) AppendThought(text string) {
	dur := s.thinkingDur
	if dur == 0 && !s.thinkingStart.IsZero() {
		dur = time.Since(s.thinkingStart)
	}
	s.thoughts = append(s.thoughts, thoughtEntry{text: text, dur: dur})
	s.timeline = append(s.timeline, timelineItem{
		typ:        tlThought,
		thoughtIdx: len(s.thoughts) - 1,
	})
	s.showThinking = false
	s.cachedBase = nil
	s.renderMessages()
}

func (s *ChatScreen) FlushResponse() {
	if s.responseText != "" {
		reasoning := ""
		if len(s.thoughts) > 0 {
			parts := make([]string, len(s.thoughts))
			for i, t := range s.thoughts {
				parts[i] = t.text
			}
			reasoning = strings.Join(parts, "\n")
		}
		s.messages = append(s.messages, ChatMessage{Role: "assistant", Content: s.responseText, Reasoning: reasoning})
		s.responseText = ""
	}
	s.streamActive = false
	s.cachedBase = nil
}

func (s *ChatScreen) SetStreamActive(active bool) {
	s.streamActive = active
	if !active {
		s.cachedBase = nil
	}
}

func (s *ChatScreen) ClearResponseText() {
	s.responseText = ""
	s.renderMessages()
}

func (s *ChatScreen) ClearToolCards() {
	s.toolCards = nil
	s.timeline = nil
	s.showThinking = false
	s.thinkingStart = time.Time{}
	s.thinkingDur = 0
	s.toolCount = 0
	s.totalDuration = 0
	s.inputTokens = 0
	s.outputTokens = 0
	s.focusedCard = -1
	s.thoughts = nil
	s.responseText = ""
	s.cachedBase = nil
	s.streamActive = false
	s.renderMessages()
}

func (s *ChatScreen) PersistToolResults() {
	for _, card := range s.toolCards {
		dur := fmt.Sprintf("%dms", card.Duration())
		output := card.Output()
		if output != "" {
			parts := strings.SplitN(output, "\n", 2)
			if len(parts) > 1 {
				firstLines := parts[0]
				rest := parts[1]
				if len(rest) > 2000 {
					rest = rest[:2000] + "... (truncated)"
				}
				output = firstLines + "\n" + rest
			}
			if len(output) > 2500 {
				output = output[:2500] + "... (truncated)"
			}
		}
		summary := fmt.Sprintf("  %s  %s  %s",
			card.Name(),
			card.Status(),
			dur,
		)
		if output != "" {
			summary += "\n" + output
		}
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
	s.renderMessages()
}

func (s *ChatScreen) ToolCount() int {
	return s.toolCount
}

func (s *ChatScreen) TotalDuration() int64 {
	return s.totalDuration
}

func (s *ChatScreen) hasActiveToolCards() bool {
	for _, c := range s.toolCards {
		if c.IsRunning() {
			return true
		}
	}
	return false
}

func (s *ChatScreen) renderMessages() {
	if len(s.messages) == 0 && !s.showThinking {
		vpW := s.viewport.Width()
		vpH := s.viewport.Height()
		if vpH < 3 {
			return
		}
		welcome := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true).Render("Start a conversation")
		centered := lipgloss.NewStyle().Width(vpW).Align(lipgloss.Center).Render(welcome)
		topPad := (vpH - 3) / 2
		if topPad < 0 {
			topPad = 0
		}
		lines := make([]string, vpH)
		for i := range lines {
			lines[i] = ""
		}
		lines[topPad] = centered
		s.viewport.SetContent(lines)
		s.viewport.GotoBottom()
		return
	}

	var lines []string
	for _, msg := range s.messages {
		roleLabel := ""
		switch msg.Role {
		case "user":
			roleLabel = userStyle.String()
			barStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("83"))
			for _, contentLine := range strings.Split(msg.Content, "\n") {
				lines = append(lines, "  "+barStyle.Render("┃")+"  "+contentLine)
			}
			lines = append(lines, "")
			continue
		case "assistant":
			roleLabel = assistantLabel(s.modelName, s.agentName)
		case "tool":
			roleLabel = lipgloss.NewStyle().Foreground(lipgloss.Color("141")).Bold(true).Render("Tool")
		default:
			roleLabel = systemStyle.String()
		}
		lines = append(lines, roleLabel)
		if msg.Role == "assistant" && msg.Reasoning != "" {
			reasonDur := fmt.Sprintf("%.0fs", float64(len(msg.Reasoning))/40)
			tokenEst := len(msg.Reasoning) / 4
			if tokenEst < 1 {
				tokenEst = 1
			}
			lines = append(lines, fmt.Sprintf("  \u25be  Thought for %s, %d tokens", reasonDur, tokenEst))
			for _, line := range strings.Split(msg.Reasoning, "\n") {
				lines = append(lines, "    "+line)
			}
			lines = append(lines, "")
		}
		if msg.Role == "assistant" && msg.Content != "" {
			s.mdRenderer.SetWidth(s.width - 6)
			rendered := s.mdRenderer.Render(msg.Content)
			for _, line := range strings.Split(rendered, "\n") {
				if components.IsDiffLine(line) {
					lines = append(lines, "  "+components.ColorizeDiff(line))
				} else {
					lines = append(lines, "  "+line)
				}
			}
		} else {
			for _, contentLine := range strings.Split(msg.Content, "\n") {
				if msg.Role == "tool" && strings.HasPrefix(contentLine, "  ") {
					trimmed := contentLine[2:]
					if components.IsDiffLine(trimmed) {
						lines = append(lines, "  "+components.ColorizeDiff(trimmed))
						continue
					}
				}
				lines = append(lines, "  "+contentLine)
			}
		}
		lines = append(lines, "")
	}

	spinnerChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinner := spinnerChars[s.animTick%10]

	for _, item := range s.timeline {
		switch item.typ {
		case tlThought:
			t := s.thoughts[item.thoughtIdx]
			dur := fmt.Sprintf("%.0fs", t.dur.Seconds())
			tokenEst := len(t.text) / 4
			if tokenEst < 1 {
				tokenEst = 1
			}
			if t.collapsed {
				lines = append(lines, fmt.Sprintf("  \u25b8  Thought for %s, %d tokens  (ctrl+o to expand)", dur, tokenEst))
			} else {
				lines = append(lines, fmt.Sprintf("  \u25be  Thought for %s, %d tokens", dur, tokenEst))
				for _, line := range strings.Split(t.text, "\n") {
					lines = append(lines, "    "+line)
				}
			}
			lines = append(lines, "")

		case tlToolCard:
			card := s.toolCards[item.cardIdx]
			card.SetFocused(item.cardIdx == s.focusedCard)
			cardLines := strings.Split(card.View(), "\n")
			lines = append(lines, cardLines...)
			lines = append(lines, "")
		}
	}

	if len(s.thoughts) == 0 && s.showThinking {
		if !s.thinkingStart.IsZero() {
			elapsed := time.Since(s.thinkingStart)
			dur := fmt.Sprintf("%.1fs", elapsed.Seconds())
			lines = append(lines, fmt.Sprintf("  %s  Working... %s", spinner, dur))
			lines = append(lines, "")
		} else {
			lines = append(lines, "  ⠋  Working...")
			lines = append(lines, "")
		}
	} else if s.thinkingDur > 0 && s.responseText == "" {
		dur := fmt.Sprintf("%.1fs", s.thinkingDur.Seconds())
		lines = append(lines, fmt.Sprintf("  %s  Working...  %s", spinner, dur))
		lines = append(lines, "")
	}

	if s.responseText != "" {
		lines = append(lines, assistantLabel(s.modelName, s.agentName))
		s.mdRenderer.SetWidth(s.width - 6)
		rendered := s.mdRenderer.Render(s.responseText)
		for _, line := range strings.Split(rendered, "\n") {
			if components.IsDiffLine(line) {
				lines = append(lines, "  "+components.ColorizeDiff(line))
			} else {
				lines = append(lines, "  "+line)
			}
		}
		lines = append(lines, "")
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
	if s.streamActive && s.cachedBase == nil && s.responseText != "" {
		labelIdx := len(lines)
		for i := len(lines) - 1; i >= 0; i-- {
			if strings.Contains(lines[i], s.modelName) && strings.Contains(lines[i], s.agentName) {
				labelIdx = i
				break
			}
		}
		s.cachedBase = make([]string, labelIdx)
		copy(s.cachedBase, lines[:labelIdx])
	}
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

	case ToolQueuedMsg:
		card := components.NewToolCard(msg.Name)
		card.SetArgs(msg.Args)
		card.SetWidth(s.width - 8)
		s.toolCards = append(s.toolCards, card)
		s.timeline = append(s.timeline, timelineItem{
			typ:     tlToolCard,
			cardIdx: len(s.toolCards) - 1,
		})
		if s.showThinking && !s.thinkingStart.IsZero() {
			s.thinkingDur = time.Since(s.thinkingStart)
		}
		s.showThinking = false
		s.renderMessages()
		return nil

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
		return nil

	case ToolOutputMsg:
		idx := len(s.toolCards) - 1
		if msg.Index >= 0 && msg.Index < len(s.toolCards) {
			idx = msg.Index
		}
		s.UpdateToolCard(idx, func(c *components.ToolCard) {
			c.SetOutput(msg.Output)
		})
		return nil

	case ToolCompletedMsg:
		idx := len(s.toolCards) - 1
		if msg.Index >= 0 && msg.Index < len(s.toolCards) {
			idx = msg.Index
		}
		s.UpdateToolCard(idx, func(c *components.ToolCard) {
			if msg.Status == "completed" {
				c.SetStatus(components.ToolCompleted)
				c.SetOutput(msg.Output)
				c.SetExpanded(false)
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
		return nil

	case ThinkingTickMsg:
		if s.showThinking || s.thinkingDur > 0 || s.hasActiveToolCards() || s.streamActive {
			s.animTick++
			for _, c := range s.toolCards {
				if c.IsRunning() {
					c.SetAnimFrame(s.animTick)
				}
			}
			s.renderMessages()
			return tea.Tick(100*time.Millisecond, func(time.Time) tea.Msg {
				return ThinkingTickMsg{}
			})
		}
		return nil

	case tea.MouseClickMsg:
		m := msg.Mouse()
		if m.Button == tea.MouseLeft && len(s.toolCards) > 0 {
			if s.focusedCard >= 0 {
				s.toolCards[s.focusedCard].ToggleExpanded()
			} else {
				s.toolCards[len(s.toolCards)-1].ToggleExpanded()
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
		if msg.String() == "ctrl+o" {
			anyExp := false
			for _, t := range s.thoughts {
				if !t.collapsed {
					anyExp = true
					break
				}
			}
			if !anyExp {
				for _, c := range s.toolCards {
					if c.Expanded() {
						anyExp = true
						break
					}
				}
			}
			if anyExp {
				for i := range s.thoughts {
					s.thoughts[i].collapsed = true
				}
				for _, c := range s.toolCards {
					c.SetExpanded(false)
				}
			} else {
				for i := range s.thoughts {
					s.thoughts[i].collapsed = false
				}
			}
			s.renderMessages()
			return nil
		}
		if s.focusedCard >= 0 && len(s.toolCards) > 0 {
			return s.handleCardFocusKey(msg.String())
		}
		return s.handleScrollKey(msg.String())
	}
	return nil
}

func (s *ChatScreen) handleCardFocusKey(key string) tea.Cmd {
	switch key {
	case "esc", "tab":
		s.focusedCard = -1
		s.renderMessages()
		return nil

	case "up":
		s.focusedCard--
		if s.focusedCard < -1 {
			s.focusedCard = len(s.toolCards) - 1
		} else if s.focusedCard < 0 && s.thinkingDur == 0 {
			s.focusedCard = len(s.toolCards) - 1
		}
		s.renderMessages()
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
		s.renderMessages()
		return nil

	case "enter":
		if s.focusedCard >= 0 {
			card := s.toolCards[s.focusedCard]
			if card.Expanded() && card.OutputLines() > 15 {
				card.ToggleShowAll()
			} else {
				card.ToggleExpanded()
			}
		}
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
			card := s.toolCards[lastIdx]
			if card.Expanded() && card.OutputLines() > 15 {
				card.ToggleShowAll()
			} else {
				card.ToggleExpanded()
			}
			s.renderMessages()
		}
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
