package tui

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"

	"termcode/internal/adapters/tui/components"
	"termcode/internal/adapters/tui/keybindings"
	"termcode/internal/adapters/tui/screens"
	"termcode/internal/application/conversation"
	costapp "termcode/internal/application/cost"
	modelapp "termcode/internal/application/model"
	pluginapp "termcode/internal/application/plugin"
	"termcode/internal/application/provider"
	approuter "termcode/internal/application/router"
	toolapp "termcode/internal/application/tool"
	"termcode/internal/domain/session"
	domainworkspace "termcode/internal/domain/workspace"
	"termcode/internal/infrastructure/collab"
	sqliterepo "termcode/internal/infrastructure/database/sqlite"
	"termcode/internal/infrastructure/eventbus"
	git "termcode/internal/infrastructure/git"
	"termcode/pkg/helpers"
)

type screenType int

const (
	screenHome screenType = iota
	screenChat
)

type AppState int

const (
	StateIdle AppState = iota
	StateCallingTool
	StateWaitingForTool
	StateStreaming
	StatePaused
	StateCompleted
	StateCancelled
	StateError
)

func (s AppState) String() string {
	switch s {
	case StateIdle:
		return "idle"
	case StateCallingTool:
		return "calling tool"
	case StateWaitingForTool:
		return "waiting for tool"
	case StateStreaming:
		return "streaming"
	case StatePaused:
		return "paused"
	case StateCompleted:
		return "completed"
	case StateCancelled:
		return "cancelled"
	case StateError:
		return "error"
	}
	return "unknown"
}

type streamContentMsg string

type streamReasoningMsg string

type streamDoneMsg struct {
	content          string
	reasoningContent string
	inputTokens      int
	outputTokens     int
	err              error
}

type CancelStreamMsg struct{}

type toolStateMsg AppState

type permissionAskMsg struct {
	toolName string
	toolArgs string
	resultCh chan<- string
}

type partialResponseMsg struct {
	content          string
	reasoningContent string
}

type AppModel struct {
	width            int
	height           int
	ready            bool
	screen           screenType
	state            AppState
	homeScreen       *screens.HomeScreen
	chatScreen       *screens.ChatScreen
	commandInput     *components.CommandInput
	cmdPalette       *components.CommandPalette
	statusBar        *components.StatusBar
	commands         *commandRegistry
	providerSvc      *provider.Service
	modelSvc         *modelapp.Service
	chatSvc          *conversation.Service
	ctx              context.Context
	cancel           context.CancelFunc
	program          *tea.Program
	modelName        string
	providerName     string
	agentName        string
	workspace        string
	history          []session.Message
	streamBuf        string
	modelRepo        *sqliterepo.ModelRepo
	sessionRepo      *sqliterepo.SessionRepo
	messageRepo      *sqliterepo.MessageRepo
	settingsRepo     *sqliterepo.SettingsRepo
	branchRepo       *sqliterepo.BranchRepo
	currentSess      *session.Session
	gitBranch        string
	activeDialog     screens.DialogScreen
	dialogStack      []screens.DialogScreen
	dialogUpdated    bool
	backgroundAnim   *screens.BackgroundAnim
	keyMgr           *keybindings.Manager
	pendingPermCh    chan<- string
	eventBus         *eventbus.Bus
	toastManager     *components.ToastManager
	contextGauge     *components.ContextGauge
	partialContent   string
	partialReasoning string
	partialInputTok  int
	partialOutputTok int
	undoHistory      []ChatEvent
	redoHistory      []ChatEvent
	branches         []ConversationBranch
	currentBranch    int
	streamGen        int
	totalMsgCount    int
	collabServer     *collab.Server
	pluginManager    *pluginapp.Manager
}

type ConversationBranch struct {
	ID      string
	Name    string
	History []session.Message
}

type ChatEvent struct {
	Type   string // "add_message", "delete_message", "edit_message", "regenerate"
	Index  int
	OldMsg *session.Message
	NewMsg *session.Message
}

func NewApp() *AppModel {
	helpers.SetEncryptionKeyDir(os.Getenv("TERMCODE_CONFIG_DIR"))

	anim := screens.NewBackgroundAnim(80, 24)
	bus := eventbus.New()
	tm := components.NewToastManager()

	bus.SubscribeAsync(eventbus.EventToolComplete, func(e eventbus.Event) {
		if name, ok := e.Data.(string); ok {
			tm.Success("Tool completed: " + name)
		}
	})

	bus.SubscribeAsync(eventbus.EventToolFailed, func(e eventbus.Event) {
		if name, ok := e.Data.(string); ok {
			tm.Error("Tool failed: " + name)
		}
	})

	bus.SubscribeAsync(eventbus.EventError, func(e eventbus.Event) {
		if msg, ok := e.Data.(string); ok {
			tm.Error(msg)
		}
	})

	bus.SubscribeAsync(eventbus.EventAttention, func(e eventbus.Event) {
		if msg, ok := e.Data.(string); ok {
			tm.Warning(msg)
		}
	})

	bus.SubscribeAsync(eventbus.EventNotification, func(e eventbus.Event) {
		if msg, ok := e.Data.(string); ok {
			tm.Info(msg)
		}
	})

	return &AppModel{
		ready:     true,
		screen:    screenHome,
		agentName: "General",
		workspace: "~",
		homeScreen: screens.NewHomeScreen(screens.HomeScreenConfig{
			ProviderName:  "none",
			ModelName:     "none",
			AgentName:     "General",
			WorkspacePath: "~",
		}),
		chatScreen:     screens.NewChatScreen(),
		commandInput:   components.NewCommandInput(),
		cmdPalette:     components.NewCommandPalette(),
		backgroundAnim: anim,
		statusBar: components.NewStatusBar(components.StatusBarConfig{
			ModelName: "none",
			AgentName: "General",
			Version:   "v0.1.0",
		}),
		state:         StateIdle,
		ctx:           context.Background(),
		eventBus:      bus,
		toastManager:  tm,
		contextGauge:  components.NewContextGauge(),
		modelName:     "none",
		providerName:  "none",
		history:       []session.Message{},
		keyMgr:        keybindings.NewManager(),
		branches:      []ConversationBranch{},
		currentBranch: 0,
		pluginManager: pluginapp.NewManager(os.ExpandEnv("$HOME/.config/termcode/plugins")),
		totalMsgCount: 0,
	}
}

func (m *AppModel) SetProviderService(svc *provider.Service, modelRepo *sqliterepo.ModelRepo, sessionRepo *sqliterepo.SessionRepo, messageRepo *sqliterepo.MessageRepo, settingsRepo *sqliterepo.SettingsRepo, branchRepo *sqliterepo.BranchRepo) {
	m.providerSvc = svc
	m.modelRepo = modelRepo
	m.sessionRepo = sessionRepo
	m.messageRepo = messageRepo
	m.settingsRepo = settingsRepo
	m.branchRepo = branchRepo
	m.modelSvc = modelapp.NewService(modelRepo)
	m.commands = newCommandRegistry(svc, m.modelSvc, m)
	runtimeRouter := approuter.NewRuntimeRouter(svc, m.modelSvc)
	m.chatSvc = conversation.NewService(runtimeRouter)
	m.chatSvc.SetCostEngine(costapp.New())
	m.chatSvc.SetCheckpointPath(filepath.Join(os.TempDir(), "termcode", "checkpoint.json"))
	if settingsRepo != nil {
		tm := toolapp.NewTrustManager(settingsRepo)
		m.chatSvc.ToolService().SetTrustManager(tm)
	}

	m.chatSvc.SetPermissionRequestFunc(func(toolName, args string, resultCh chan<- string) {
		if m.program != nil {
			m.program.Send(permissionAskMsg{toolName: toolName, toolArgs: args, resultCh: resultCh})
		} else {
			resultCh <- "deny"
		}
	})

	m.homeScreen.SetAnim(m.backgroundAnim)

	p, err := svc.GetDefault(m.ctx)
	if err == nil && p != nil {
		m.providerName = p.Name
		m.homeScreen.UpdateConfig(screens.HomeScreenConfig{
			ProviderName: p.Name,
			ProviderURL:  p.BaseURL,
		})
		m.statusBar.SetModel(p.Name)
		m.chatScreen.SetAgent(m.agentName)
	}

	models, _ := m.modelSvc.List(m.ctx)
	if len(models) > 0 {
		m.modelName = models[0].ModelID
		saved, _ := m.settingsRepo.Get(m.ctx, "model_id")
		if saved != "" {
			for _, mo := range models {
				if mo.ModelID == saved {
					m.modelName = saved
					break
				}
			}
		}
		m.chatSvc.SetModelID(m.modelName)
		m.chatScreen.SetModel(m.modelName)
		m.homeScreen.UpdateConfig(screens.HomeScreenConfig{
			ModelName: m.modelName,
		})
		m.statusBar.SetModel(m.providerName + "/" + m.modelName)
	}

	savedAgent, _ := m.settingsRepo.Get(m.ctx, "agent_name")
	if savedAgent != "" {
		m.agentName = savedAgent
		m.chatScreen.SetAgent(m.agentName)
		m.homeScreen.UpdateConfig(screens.HomeScreenConfig{AgentName: m.agentName})
		m.statusBar.SetAgent(m.agentName)
	}

	m.detectGitBranch()
	m.ensureSession()
	m.screen = screenHome
}

func (m *AppModel) SetMessageRepo(repo *sqliterepo.MessageRepo) {
	m.messageRepo = repo
}

func (m *AppModel) loadSessionByID(ctx context.Context, id string) {
	if m.sessionRepo == nil {
		return
	}
	sess, err := m.sessionRepo.GetByID(ctx, id)
	if err != nil || sess == nil {
		return
	}
	m.currentSess = sess
	m.history = nil
	m.chatScreen = screens.NewChatScreen()
	if m.messageRepo != nil {
		msgs, err := m.messageRepo.ListBySession(ctx, sess.ID)
		if err == nil {
			m.history = make([]session.Message, len(msgs))
			for i, msg := range msgs {
				m.history[i] = *msg
				m.chatScreen.AddMessage(string(msg.Role), msg.Content, msg.Reasoning)
			}
		}
	}
	m.loadBranches()
	m.screen = screenChat
}

func (m *AppModel) loadLastSession() {
	if m.sessionRepo == nil {
		return
	}
	sessions, err := m.sessionRepo.ListActive(m.ctx)
	if err != nil || len(sessions) == 0 {
		return
	}
	m.currentSess = sessions[0]
	m.loadBranches()

	if m.messageRepo != nil {
		msgs, err := m.messageRepo.ListBySession(m.ctx, m.currentSess.ID)
		if err == nil {
			m.history = make([]session.Message, len(msgs))
			for i, msg := range msgs {
				m.history[i] = *msg
				m.chatScreen.AddMessage(string(msg.Role), msg.Content, msg.Reasoning)
			}
			if len(msgs) > 0 {
				m.screen = screenChat
			}
		}
	}
}

func (m *AppModel) loadBranches() {
	if m.branchRepo == nil || m.currentSess == nil {
		return
	}
	records, err := m.branchRepo.ListBySession(m.ctx, m.currentSess.ID)
	if err != nil {
		return
	}
	m.branches = nil
	for _, r := range records {
		m.branches = append(m.branches, ConversationBranch{
			ID:      r.ID,
			Name:    r.Name,
			History: r.History,
		})
	}
	if len(m.branches) > 0 {
		m.currentBranch = len(m.branches) - 1
	}
}

func (m *AppModel) ensureSession() {
	if m.currentSess != nil {
		return
	}

	now := time.Now()
	sessName := "Session " + now.Format("2006-01-02 15:04")
	s := session.New(sessName, m.providerName, m.modelName)
	m.currentSess = s

	if m.sessionRepo != nil {
		if err := m.sessionRepo.Create(m.ctx, s); err != nil {
			m.toastManager.Error("Failed to create session: " + err.Error())
		}
	}
}

func (m *AppModel) saveMessage(role session.Role, content string, reasoning ...string) {
	if m.currentSess == nil {
		return
	}
	msg := session.NewMessage(m.currentSess.ID, role, content)
	if len(reasoning) > 0 {
		msg.Reasoning = reasoning[0]
	}
	m.history = append(m.history, *msg)
	m.totalMsgCount++

	if m.messageRepo != nil {
		if err := m.messageRepo.Create(m.ctx, msg); err != nil {
			m.toastManager.Error("Failed to save message: " + err.Error())
		}
	}

	if m.sessionRepo != nil && m.currentSess != nil {
		m.currentSess.MessageCnt = m.totalMsgCount
		m.currentSess.UpdatedAt = time.Now()
		if err := m.sessionRepo.Update(m.ctx, m.currentSess); err != nil {
			m.toastManager.Error("Failed to update session: " + err.Error())
		}
	}
}

func (m *AppModel) SetSelectedModel(modelID string) {
	m.modelName = modelID
	m.chatSvc.SetModelID(modelID)
	m.homeScreen.UpdateConfig(screens.HomeScreenConfig{
		ModelName: modelID,
	})
	m.statusBar.SetModel(m.providerName + "/" + modelID)
	m.chatScreen.SetModel(modelID)
	if m.settingsRepo != nil {
		m.settingsRepo.Set(m.ctx, "model_id", modelID)
	}
	m.eventBus.Emit(eventbus.EventModelChanged, modelID)
	m.toastManager.Info("Model: " + modelID)
}

func (m *AppModel) SetProgram(p *tea.Program) {
	m.program = p
}

func (m *AppModel) startChat(ctx context.Context) string {
	lastUserIdx := -1
	for i := len(m.history) - 1; i >= 0; i-- {
		if m.history[i].Role == session.RoleUser {
			lastUserIdx = i
			break
		}
	}
	if lastUserIdx < 0 {
		return "No user message to retry."
	}

	m.state = StateStreaming
	m.streamBuf = ""
	m.partialContent = ""
	m.partialReasoning = ""

	if m.cancel != nil {
		m.cancel()
	}
	m.ctx, m.cancel = context.WithCancel(context.Background())

	historyCopy := make([]session.Message, len(m.history))
	copy(historyCopy, m.history)
	input := m.history[lastUserIdx].Content

	m.eventBus.Emit(eventbus.EventStreamStarted, input)

	m.streamGen++
	gen := m.streamGen
	go m.runStream(input, historyCopy, m.program, gen)

	return "__startchat__"
}

func (m *AppModel) runStream(input string, historyCopy []session.Message, prog *tea.Program, gen int) {
	onToolEvent := func(ev conversation.ToolEvent) {
		if gen != m.streamGen {
			return
		}
		if prog == nil {
			return
		}
		switch ev.Type {
		case conversation.ToolQueued:
			prog.Send(screens.ToolQueuedMsg{Index: ev.Index, Name: ev.Name, Args: ev.Args})
			prog.Send(toolStateMsg(StateCallingTool))
		case conversation.ToolInitializing:
			prog.Send(screens.ToolInitializingMsg{Index: ev.Index, Name: ev.Name})
			prog.Send(toolStateMsg(StateWaitingForTool))
		case conversation.ToolConnecting:
			prog.Send(screens.ToolConnectingMsg{Index: ev.Index, Name: ev.Name})
			prog.Send(toolStateMsg(StateWaitingForTool))
		case conversation.ToolStarted:
			prog.Send(screens.ToolStartedMsg{Index: ev.Index, Name: ev.Name})
			prog.Send(toolStateMsg(StateCallingTool))
			m.eventBus.Emit(eventbus.EventToolStarted, ev.Name)
		case conversation.ToolOutput:
			prog.Send(screens.ToolOutputMsg{Index: ev.Index, Name: ev.Name, Output: ev.Output})
			prog.Send(toolStateMsg(StateWaitingForTool))
		case conversation.ToolCompleted:
			prog.Send(screens.ToolCompletedMsg{Index: ev.Index, Name: ev.Name, Output: ev.Output, Status: ev.Status, Error: ev.Error, Duration: ev.Duration})
			if ev.Error != "" {
				m.eventBus.Emit(eventbus.EventToolFailed, ev.Name+": "+ev.Error)
			} else {
				m.eventBus.Emit(eventbus.EventToolComplete, ev.Name)
			}
		}
	}

	ctx := m.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	result, err := m.chatSvc.SendMessage(ctx, input, historyCopy, func(chunk string, reasoning bool, done bool) {
		if prog != nil {
			if reasoning {
				prog.Send(streamReasoningMsg(chunk))
			} else {
				prog.Send(streamContentMsg(chunk))
			}
		}
	}, onToolEvent)
	if prog != nil {
		tokIn, tokOut := 0, 0
		if result != nil {
			tokIn = result.InputTokens
			tokOut = result.OutputTokens
		}
		content := ""
		reasoningContent := ""
		if result != nil {
			content = result.Content
			reasoningContent = result.ReasoningContent
		}
		prog.Send(streamDoneMsg{content: content, reasoningContent: reasoningContent, inputTokens: tokIn, outputTokens: tokOut, err: err})
	}
}

func (m *AppModel) ShowDialog(d screens.DialogScreen) {
	if m.activeDialog != nil {
		if len(m.dialogStack) >= 10 {
			m.toastManager.Error("Dialog stack full")
			return
		}
		m.dialogStack = append(m.dialogStack, m.activeDialog)
	}
	m.dialogUpdated = true
	if m.width > 0 && m.height > 0 {
		d.SetSize(m.width, m.height)
	}
	m.activeDialog = d
	m.keyMgr.PushScope(keybindings.ScopeDialog)
	m.layout()
}

func (m *AppModel) closeDialog() {
	if m.activeDialog != nil {
		m.activeDialog = nil
		m.keyMgr.PopScope()
	}
	if len(m.dialogStack) > 0 {
		m.activeDialog = m.dialogStack[len(m.dialogStack)-1]
		m.dialogStack = m.dialogStack[:len(m.dialogStack)-1]
		if m.width > 0 && m.height > 0 {
			m.activeDialog.SetSize(m.width, m.height)
		}
	}
	m.layout()
}

func (m *AppModel) SetWorkspace(path string) {
	m.workspace = path
	m.homeScreen.UpdateConfig(screens.HomeScreenConfig{WorkspacePath: path})

	wt := domainworkspace.NewTrustManager()
	if !wt.IsTrusted(path) {
		d := screens.NewWorkspaceTrustDialog(path, wt.IsDangerous(path))
		m.ShowDialog(d)
	}
}

func (m *AppModel) IsWorkspaceTrusted() bool {
	wt := domainworkspace.NewTrustManager()
	return wt.IsTrusted(m.workspace)
}

// layout is the single source of truth for all component sizing.
// model(1) + sep(1) + input top(1) + input content(1) + input bottom(1) + status(1) + safe(1) + palette(h) = 7 + paletteHeight
func (m *AppModel) layout() {
	if m.width == 0 || m.height == 0 {
		return
	}

	// Palette max items based on available space
	maxPalette := m.height - 8
	if maxPalette < 3 {
		maxPalette = 3
	}
	maxItems := maxPalette - 3
	if maxItems < 1 {
		maxItems = 1
	}
	if maxItems > 6 {
		maxItems = 6
	}
	m.cmdPalette.SetMaxItems(maxItems)

	paletteH := m.cmdPalette.Height()

	// Chat screen viewport: remaining after fixed + palette overhead
	// AppModel.View() = vpH + paletteH + state(1) + input(3) + gauge(1) + status(1) + safe(1) = vpH + paletteH + 8
	extraLines := 0
	if m.toastManager.HasToast() {
		extraLines += 2
	}
	if m.contextGauge.IsActive() {
		extraLines++
	}
	if m.state != StateIdle {
		extraLines++
	}
	overheadVp := 7 + paletteH + extraLines
	vpW := m.width - 4
	if vpW < 10 {
		vpW = 10
	}
	vpH := m.height - overheadVp
	if vpH < 1 {
		vpH = 1
	}
	m.chatScreen.SetViewportSize(vpW, vpH)

	// Home screen: content fills avilable space
	// View() = mainContent (homeH-4 lines) + palette (paletteH) + input (3) + status (1) = homeH + paletteH
	// homeH + paletteH = m.height  =>  homeH = m.height - paletteH
	homeH := m.height - paletteH
	if homeH < 8 {
		homeH = 8
	}
	m.homeScreen.SetSize(m.width, homeH)

	m.commandInput.SetWidth(m.width)
	m.cmdPalette.SetWidth(m.width)
	m.statusBar.SetWidth(m.width)
	m.contextGauge.SetWidth(m.width)
	m.toastManager.SetSize(m.width)
	if m.activeDialog != nil {
		m.activeDialog.SetSize(m.width, m.height)
	}
}

func (m *AppModel) Init() tea.Cmd {
	return tea.Batch(
		m.commandInput.Init(),
		tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
			return screens.AnimTickMsg{}
		}),
	)
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if wmsg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = wmsg.Width
		m.height = wmsg.Height
		m.ready = true

		m.homeScreen.SetSize(wmsg.Width, wmsg.Height)
		m.backgroundAnim.SetSize(wmsg.Width, wmsg.Height)

		var cmd tea.Cmd
		if m.activeDialog != nil {
			m.activeDialog.SetSize(wmsg.Width, wmsg.Height)
			var d screens.DialogScreen
			d, cmd = m.activeDialog.Update(wmsg)
			m.activeDialog = d
		}

		m.layout()
		return m, cmd
	}

	// Global CancelStreamMsg — cancels any active streaming
	if _, ok := msg.(CancelStreamMsg); ok {
		if m.cancel != nil && m.state != StateIdle && m.state != StatePaused && m.state != StateCompleted && m.state != StateCancelled {
			m.cancel()
			m.cancel = nil
			m.partialContent = m.streamBuf
			m.state = StateCancelled
			m.statusBar.SetWorking(false)
			m.chatScreen.HideThinking()
			m.chatScreen.SetStreamActive(false)
			m.eventBus.Emit(eventbus.EventStreamComplete, nil)
			m.toastManager.Warning("Generation cancelled")
			return m, nil
		}
	}

	switch msg := msg.(type) {
	case permissionAskMsg:
		m.pendingPermCh = msg.resultCh
		d := screens.NewToolConfirmDialog(msg.toolName, msg.toolArgs, "Allow this tool?")
		m.ShowDialog(d)
		return m, nil
	}

	if m.activeDialog != nil {
		switch msg := msg.(type) {
		case tea.KeyMsg, tea.PasteMsg:
			m.dialogUpdated = false
			d, cmd := m.activeDialog.Update(msg)
			if m.dialogUpdated {
				return m, nil
			}
			m.activeDialog = d
			if cmd != nil {
				return m, cmd
			}
			if m.activeDialog.Done() {
				result := m.activeDialog.Result()
				m.closeDialog()

				if m.pendingPermCh != nil {
					m.pendingPermCh <- result
					m.pendingPermCh = nil
					return m, nil
				}

				switch {
				case result == "__trust__":
					domainworkspace.NewTrustManager().MarkTrusted(m.workspace)
				case strings.HasPrefix(result, "__rename__"):
					parts := strings.SplitN(result, ":", 2)
					if len(parts) == 2 && m.sessionRepo != nil {
						m.currentSess.Name = parts[1]
						m.sessionRepo.Update(m.ctx, m.currentSess)
						m.toastManager.Info("Session renamed: " + parts[1])
					}
				case strings.HasPrefix(result, "__pin__"):
					if m.sessionRepo != nil {
						m.currentSess.Status = session.StatusPinned
						m.sessionRepo.Update(m.ctx, m.currentSess)
						m.toastManager.Info("Session pinned")
					}
				case strings.HasPrefix(result, "__unpin__"):
					if m.sessionRepo != nil {
						m.currentSess.Status = session.StatusActive
						m.sessionRepo.Update(m.ctx, m.currentSess)
						m.toastManager.Info("Session unpinned")
					}
				case strings.HasPrefix(result, "__branch__"):
					m.switchToBranch(strings.TrimPrefix(result, "__branch__:"))
				case strings.HasPrefix(result, "__editmsg__"):
					parts := strings.SplitN(result, ":", 3)
					if len(parts) == 3 {
						idx := 0
						fmt.Sscanf(parts[1], "%d", &idx)
						m.editMessage(idx, parts[2])
					}
				case strings.HasPrefix(result, "__delmsg__"):
					idx := 0
					fmt.Sscanf(strings.TrimPrefix(result, "__delmsg__:"), "%d", &idx)
					m.deleteMessage(idx)
				case result != "" && !strings.HasPrefix(result, "__"):
					if m.sessionRepo != nil {
						_, err := m.sessionRepo.GetByID(m.ctx, result)
						if err == nil {
							m.loadSessionByID(m.ctx, result)
							m.layout()
							return m, tea.Tick(100*time.Millisecond, func(time.Time) tea.Msg {
								return screens.ThinkingTickMsg{}
							})
						}
					}
					if m.screen == screenHome {
						m.screen = screenChat
					}
					m.layout()
				}
			}
			return m, nil
		}
	}

	var statusCmd tea.Cmd
	switch msg := msg.(type) {

	case screens.AnimTickMsg:
		if m.activeDialog == nil {
			if m.screen == screenHome {
				m.backgroundAnim.Tick()
			}
			return m, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
				return screens.AnimTickMsg{}
			})
		}
		return m, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
			return screens.AnimTickMsg{}
		})

	case tea.PasteMsg:
		if m.activeDialog != nil {
			d, cmd := m.activeDialog.Update(msg)
			m.activeDialog = d
			return m, cmd
		}

	case tea.KeyMsg:
		if m.cmdPalette.Visible() {
			switch msg.String() {
			case "up":
				m.cmdPalette.SelectUp()
				m.layout()
				return m, nil
			case "down":
				m.cmdPalette.SelectDown()
				m.layout()
				return m, nil
			case "enter", "tab":
				cmd := m.cmdPalette.SelectedCommand()
				if cmd != "" {
					m.commandInput.SetValue(cmd + " ")
					m.commandInput.SetFocused(true)
				}
				m.cmdPalette.Hide()
				m.layout()
				return m, nil
			case "esc":
				m.cmdPalette.Hide()
				m.layout()
				return m, nil
			}
		}
		if m.activeDialog == nil {
			switch msg.String() {
			case "tab":
				agents := []string{"General", "Expert", "Architect"}
				idx := 0
				for i, a := range agents {
					if a == m.agentName {
						idx = (i + 1) % len(agents)
						break
					}
				}
				m.agentName = agents[idx]
				m.homeScreen.UpdateConfig(screens.HomeScreenConfig{AgentName: agents[idx]})
				m.statusBar.SetAgent(agents[idx])
				m.chatScreen.SetAgent(agents[idx])
				if m.settingsRepo != nil {
					m.settingsRepo.Set(m.ctx, "agent_name", agents[idx])
				}
				return m, nil
			}
		}

		switch msg.String() {
		case "ctrl+c":
			if m.state != StateIdle && m.state != StateCompleted && m.state != StateCancelled && m.cancel != nil {
				return m, func() tea.Msg { return CancelStreamMsg{} }
			}
			return m, tea.Quit
		}

		if m.state == StateStreaming && msg.String() == "esc" {
			m.state = StatePaused
			if m.cancel != nil {
				m.cancel()
				m.cancel = nil
			}
			m.partialContent = m.streamBuf
			m.chatScreen.HideThinking()
			m.chatScreen.ClearResponseText()
			return m, nil
		}

		_ = m.keyMgr.Resolve(msg)

	case components.SubmitMsg:
		input := string(msg)
		if input == "" {
			return m, nil
		}

		if input[0] == '/' {
			result := m.commands.execute(m.ctx, input)

			switch result {
			case "__exit__":
				return m, tea.Quit
			case "__clear__":
				m.screen = screenHome
				m.state = StateIdle
				m.statusBar.SetWorking(false)
				m.layout()
				return m, nil
			case "__home__":
				m.screen = screenHome
				m.state = StateIdle
				m.statusBar.SetWorking(false)
				m.layout()
				return m, nil
			case "__chat__":
				m.layout()
				return m, nil
			case "__dialog__":
				return m, nil
			case "__startchat__":
				m.layout()
				statusCmd = m.statusBar.SetWorking(true)
				cmd := m.chatScreen.ShowThinking()
				if statusCmd != nil {
					cmd = tea.Batch(cmd, statusCmd)
				}
				return m, cmd
			default:
				if result != "" {
					m.toastManager.Info(result)
				}
			}
			return m, nil
		}

		if m.screen == screenHome {
			m.screen = screenChat
			m.layout()
		}

		m.ensureSession()
		m.chatScreen.FlushResponse()
		m.chatScreen.AddMessage("user", input)
		m.saveMessage(session.RoleUser, input)

		m.state = StateStreaming
		m.streamBuf = ""
		m.partialContent = ""
		m.partialReasoning = ""
		statusCmd = m.statusBar.SetWorking(true)

		// Create cancellable context
		if m.cancel != nil {
			m.cancel()
		}
		m.ctx, m.cancel = context.WithCancel(context.Background())

		historyCopy := make([]session.Message, len(m.history))
		copy(historyCopy, m.history)

		m.chatScreen.ClearToolCards()
		cmd := m.chatScreen.ShowThinking()
		if statusCmd != nil {
			cmd = tea.Batch(cmd, statusCmd)
		}

		m.eventBus.Emit(eventbus.EventStreamStarted, input)

		m.streamGen++
		gen := m.streamGen
		go m.runStream(input, historyCopy, m.program, gen)

		return m, cmd

	case streamContentMsg:
		text := string(msg)
		m.streamBuf += text
		if len(m.streamBuf) > 1_048_576 {
			m.streamBuf = m.streamBuf[len(m.streamBuf)-524_288:]
		}
		return m, m.chatScreen.AppendResponseText(text)

	case streamReasoningMsg:
		m.chatScreen.AppendThought(string(msg))
		return m, nil

	case toolStateMsg:
		m.state = AppState(msg)
		return m, nil

	case streamDoneMsg:
		if m.state == StatePaused {
			if m.cancel != nil {
				m.cancel()
				m.cancel = nil
			}
			return m, nil
		}
		m.state = StateCompleted
		m.statusBar.SetWorking(false)
		m.chatScreen.HideThinking()
		m.eventBus.Emit(eventbus.EventStreamComplete, nil)
		if m.cancel != nil {
			m.cancel()
			m.cancel = nil
		}
		if msg.err != nil {
			m.state = StateError
			if errors.Is(msg.err, context.Canceled) {
				m.state = StateCancelled
				m.toastManager.Warning("Generation cancelled")
				if m.partialContent != "" {
					m.chatScreen.SetResponse(m.partialContent)
				}
				return m, nil
			}
			m.eventBus.Emit(eventbus.EventError, msg.err.Error())
			m.chatScreen.ClearToolCards()
			m.chatScreen.AddMessage("system", "Error: "+msg.err.Error())
			m.saveMessage(session.RoleSystem, "Error: "+msg.err.Error())
		} else {
			m.chatScreen.SetToolMetrics(m.chatScreen.ToolCount(), m.chatScreen.TotalDuration(), msg.inputTokens, msg.outputTokens)
			m.chatScreen.PersistToolResults()
			if msg.inputTokens > 0 || msg.outputTokens > 0 {
				ctxMax := m.getContextLimit()
				m.contextGauge.SetUsage(msg.inputTokens+msg.outputTokens, ctxMax)
			}
			if msg.content != "" {
				m.chatScreen.SetResponse(msg.content)
				m.saveMessage(session.RoleAssistant, msg.content, msg.reasoningContent)
				// Auto-compression check
				if m.shouldCompressContext() {
					m.compressContext()
				}
			}
		}
		m.cancel = nil
		// Auto-save branch
		if m.currentBranch > 0 && len(m.branches) > 0 {
			b := &m.branches[m.currentBranch]
			b.History = make([]session.Message, len(m.history))
			copy(b.History, m.history)
			if m.branchRepo != nil && m.currentSess != nil {
				if err := m.branchRepo.Update(m.ctx, &sqliterepo.BranchRecord{
					ID:      b.ID,
					Name:    b.Name,
					History: b.History,
				}); err != nil {
					m.toastManager.Error("Failed to save branch: " + err.Error())
				}
			}
		}
		return m, nil

	default:
		statusCmd = m.statusBar.Update(msg)
	}

	chatCmd := m.chatScreen.Update(msg)

	var cmd tea.Cmd
	m.commandInput, cmd = m.commandInput.Update(msg)
	if chatCmd != nil {
		cmd = tea.Batch(chatCmd, cmd)
	}
	if statusCmd != nil {
		cmd = tea.Batch(cmd, statusCmd)
	}

	val := m.commandInput.Value()
	if strings.HasPrefix(val, "/") {
		filter := ""
		if len(val) > 1 {
			filter = val[1:]
		}
		m.cmdPalette.SetFilter(filter)
		if !m.cmdPalette.Visible() {
			m.cmdPalette.Show()
		}
		m.layout()
	} else {
		if m.cmdPalette.Visible() {
			m.cmdPalette.Hide()
			m.layout()
		}
	}

	return m, cmd
}

func (m *AppModel) View() tea.View {
	statusView := m.statusBar.View()

	var mainContent string
	if m.activeDialog != nil {
		mainContent = m.activeDialog.View()
		dialogLines := strings.Count(mainContent, "\n") + 1
		availableSpace := m.height - dialogLines - 1
		if availableSpace > 0 {
			topPad := availableSpace / 2
			bottomPad := availableSpace - topPad
			mainContent = strings.Repeat("\n", topPad) + mainContent + strings.Repeat("\n", bottomPad)
		} else {
			mainContent = mainContent + "\n"
		}
	} else if m.screen == screenHome {
		mainContent = m.homeScreen.View()
	} else {
		mainContent = m.chatScreen.View()
	}

	paletteView := m.cmdPalette.View()
	inputView := m.commandInput.View()
	if m.activeDialog != nil {
		paletteView = ""
		inputView = ""
	}

	toastView := m.toastManager.View()
	if toastView != "" {
		toastView = toastView + "\n"
	}

	var result string
	gaugeView := m.contextGauge.View()
	if gaugeView != "" {
		gaugeView = "\n" + gaugeView
	}

	// State indicator
	stateLine := ""
	if m.state != StateIdle && m.state != StateCompleted {
		stateColor := "245"
		switch m.state {
		case StateStreaming:
			stateColor = "39"
		case StateCallingTool, StateWaitingForTool:
			stateColor = "221"
		case StatePaused:
			stateColor = "214"
		case StateCancelled:
			stateColor = "196"
		case StateError:
			stateColor = "196"
		}
		stateLine = "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color(stateColor)).Render("  ● "+m.state.String())
	}

	if m.activeDialog != nil {
		result = toastView + mainContent + statusView
	} else if paletteView != "" {
		result = toastView + mainContent + "\n" + paletteView + "\n" + inputView + gaugeView + stateLine + "\n" + statusView
	} else {
		result = toastView + mainContent + "\n" + inputView + gaugeView + stateLine + "\n" + statusView
	}
	v := tea.NewView(result)
	v.AltScreen = true
	v.MouseMode = tea.MouseModeAllMotion
	return v
}

func (m *AppModel) detectGitBranch() {
	svc := git.NewService()
	if svc.IsRepo(m.workspace) {
		repo, err := svc.Open(m.workspace)
		if err == nil {
			branch, err := svc.GetBranch(repo)
			if err == nil && branch != "" {
				m.gitBranch = branch
				m.statusBar.SetBranch(branch)
			}
		}
	}
}

func (m *AppModel) getContextLimit() int {
	if m.modelRepo == nil {
		return 128000
	}
	models, err := m.modelSvc.List(m.ctx)
	if err != nil || len(models) == 0 {
		return 128000
	}
	for _, mo := range models {
		if mo.ModelID == m.modelName {
			if mo.MaxContext > 0 {
				return mo.MaxContext
			}
			break
		}
	}
	return 128000
}

func (m *AppModel) shouldCompressContext() bool {
	total := 0
	for _, h := range m.history {
		total += len(h.Content)
	}
	return total > m.getContextLimit()*3
}

func (m *AppModel) compressContext() {
	if len(m.history) < 4 {
		return
	}
	var kept []session.Message
	kept = append(kept, m.history[0])
	for i := len(m.history) - 3; i < len(m.history); i++ {
		if i >= 0 {
			kept = append(kept, m.history[i])
		}
	}
	summary := fmt.Sprintf("[Context compressed: %d messages summarized]", len(m.history)-len(kept))
	kept = append([]session.Message{{
		Role:    session.RoleSystem,
		Content: summary,
	}}, kept...)
	m.history = kept
	m.toastManager.Info(summary)
}

func (m *AppModel) switchToBranch(branchID string) {
	for i, b := range m.branches {
		if b.ID == branchID {
			m.currentBranch = i
			m.history = make([]session.Message, len(b.History))
			copy(m.history, b.History)
			m.chatScreen = screens.NewChatScreen()
			for _, msg := range m.history {
				m.chatScreen.AddMessage(string(msg.Role), msg.Content, msg.Reasoning)
			}
			m.currentSess = nil
			m.toastManager.Info("Switched to branch: " + b.Name)
			return
		}
	}
}

func (m *AppModel) createBranch(name string) {
	branch := ConversationBranch{
		ID:      uuid.New().String(),
		Name:    name,
		History: make([]session.Message, len(m.history)),
	}
	copy(branch.History, m.history)
	m.branches = append(m.branches, branch)
	m.currentBranch = len(m.branches) - 1
	if m.branchRepo != nil && m.currentSess != nil {
		rec := &sqliterepo.BranchRecord{
			ID:        branch.ID,
			SessionID: m.currentSess.ID,
			Name:      branch.Name,
			History:   branch.History,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		m.branchRepo.Create(m.ctx, rec)
	}
	m.toastManager.Info("Created branch: " + name)
}

func (m *AppModel) editMessage(idx int, newContent string) {
	if idx < 0 || idx >= len(m.history) {
		return
	}
	oldMsg := m.history[idx]
	event := ChatEvent{Type: "edit_message", Index: idx, OldMsg: &oldMsg}
	newMsg := oldMsg
	newMsg.Content = newContent
	event.NewMsg = &newMsg
	m.undoHistory = append(m.undoHistory, event)
	m.redoHistory = nil

	m.history[idx] = newMsg
	if m.messageRepo != nil {
		m.history[idx].Content = newContent
		if err := m.messageRepo.Update(m.ctx, &m.history[idx]); err != nil {
			m.toastManager.Error("Failed to update message: " + err.Error())
		}
	}
	if m.sessionRepo != nil && m.currentSess != nil {
		if err := m.sessionRepo.Update(m.ctx, m.currentSess); err != nil {
			m.toastManager.Error("Failed to update session: " + err.Error())
		}
	}
	m.chatScreen.UpdateMessage(idx, string(newMsg.Role), newMsg.Content, newMsg.Reasoning)
	m.toastManager.Info("Message edited")
}

func (m *AppModel) deleteMessage(idx int) {
	if idx < 0 || idx >= len(m.history) {
		return
	}
	oldMsg := m.history[idx]
	event := ChatEvent{Type: "delete_message", Index: idx, OldMsg: &oldMsg}
	m.undoHistory = append(m.undoHistory, event)
	m.redoHistory = nil

	m.history = append(m.history[:idx], m.history[idx+1:]...)
	if m.messageRepo != nil {
		if err := m.messageRepo.Delete(m.ctx, oldMsg.ID); err != nil {
			m.toastManager.Error("Failed to delete message: " + err.Error())
		}
	}
	if m.currentSess != nil {
		m.currentSess.MessageCnt = len(m.history)
		if m.sessionRepo != nil {
			if err := m.sessionRepo.Update(m.ctx, m.currentSess); err != nil {
				m.toastManager.Error("Failed to update session: " + err.Error())
			}
		}
	}
	m.chatScreen.RemoveMessage(idx)
	m.toastManager.Info("Message deleted")
}

func (m *AppModel) undoLastEvent() {
	if len(m.undoHistory) == 0 {
		return
	}
	event := m.undoHistory[len(m.undoHistory)-1]
	m.undoHistory = m.undoHistory[:len(m.undoHistory)-1]
	m.redoHistory = append(m.redoHistory, event)

	switch event.Type {
	case "edit_message":
		if event.OldMsg != nil && event.Index >= 0 && event.Index < len(m.history) {
			m.history[event.Index] = *event.OldMsg
			m.chatScreen.UpdateMessage(event.Index, string(event.OldMsg.Role), event.OldMsg.Content, event.OldMsg.Reasoning)
			if m.messageRepo != nil {
				if err := m.messageRepo.Update(m.ctx, &m.history[event.Index]); err != nil {
					m.toastManager.Error("Failed to update message: " + err.Error())
				}
			}
		}
	case "delete_message":
		if event.OldMsg != nil {
			m.history = append(m.history, *event.OldMsg)
			copy(m.history[event.Index+1:], m.history[event.Index:])
			m.history[event.Index] = *event.OldMsg
			m.chatScreen.InsertMessage(event.Index, string(event.OldMsg.Role), event.OldMsg.Content, event.OldMsg.Reasoning)
			if m.messageRepo != nil {
				if err := m.messageRepo.Create(m.ctx, &m.history[event.Index]); err != nil {
					m.toastManager.Error("Failed to restore message: " + err.Error())
				}
			}
		}
	}
	m.toastManager.Info("Undone")
}

func (m *AppModel) redoLastEvent() {
	if len(m.redoHistory) == 0 {
		return
	}
	event := m.redoHistory[len(m.redoHistory)-1]
	m.redoHistory = m.redoHistory[:len(m.redoHistory)-1]
	m.undoHistory = append(m.undoHistory, event)

	switch event.Type {
	case "edit_message":
		if event.NewMsg != nil && event.Index >= 0 && event.Index < len(m.history) {
			m.history[event.Index] = *event.NewMsg
			m.chatScreen.UpdateMessage(event.Index, string(event.NewMsg.Role), event.NewMsg.Content, event.NewMsg.Reasoning)
			if m.messageRepo != nil {
				if err := m.messageRepo.Update(m.ctx, &m.history[event.Index]); err != nil {
					m.toastManager.Error("Failed to update message: " + err.Error())
				}
			}
		}
	case "delete_message":
		if event.Index >= 0 && event.Index < len(m.history) {
			if m.messageRepo != nil {
				if err := m.messageRepo.Delete(m.ctx, m.history[event.Index].ID); err != nil {
					m.toastManager.Error("Failed to delete message: " + err.Error())
				}
			}
			m.history = append(m.history[:event.Index], m.history[event.Index+1:]...)
			m.chatScreen.RemoveMessage(event.Index)
		}
	}
	m.toastManager.Info("Redone")
}
