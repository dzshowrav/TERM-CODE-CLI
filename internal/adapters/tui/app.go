package tui

import (
	"context"
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"

	"termcode/internal/adapters/tui/components"
	"termcode/internal/adapters/tui/screens"
	"termcode/internal/application/conversation"
	modelapp "termcode/internal/application/model"
	"termcode/internal/application/provider"
	approuter "termcode/internal/application/router"
	"termcode/internal/domain/session"
	sqliterepo "termcode/internal/infrastructure/database/sqlite"
	git "termcode/internal/infrastructure/git"
)

type screenType int

const (
	screenHome screenType = iota
	screenChat
)

type streamContentMsg string

type streamDoneMsg struct {
	content      string
	inputTokens  int
	outputTokens int
	err          error
}

type AppModel struct {
	width        int
	height       int
	ready        bool
	screen       screenType
	homeScreen   *screens.HomeScreen
	chatScreen   *screens.ChatScreen
	commandInput *components.CommandInput
	cmdPalette   *components.CommandPalette
	statusBar    *components.StatusBar
	working      bool
	commands     *commandRegistry
	providerSvc  *provider.Service
	modelSvc     *modelapp.Service
	chatSvc      *conversation.Service
	ctx          context.Context
	program      *tea.Program
	modelName    string
	providerName string
	agentName    string
	workspace    string
	history      []session.Message
	streamBuf    string
	modelRepo    *sqliterepo.ModelRepo
	sessionRepo  *sqliterepo.SessionRepo
	messageRepo  *sqliterepo.MessageRepo
	currentSess  *session.Session
	gitBranch    string
	activeDialog screens.DialogScreen
}

func NewApp() *AppModel {
	return &AppModel{
		screen:    screenHome,
		ctx:       context.Background(),
		agentName: "General",
		workspace: "~",
		homeScreen: screens.NewHomeScreen(screens.HomeScreenConfig{
			ProviderName:  "none",
			ModelName:     "none",
			AgentName:     "General",
			WorkspacePath: "~",
		}),
		chatScreen:   screens.NewChatScreen(),
		commandInput: components.NewCommandInput(),
		cmdPalette:   components.NewCommandPalette(),
		statusBar: components.NewStatusBar(components.StatusBarConfig{
			ModelName: "none",
			AgentName: "General",
			Version:   "v0.1.0",
		}),
		working:      false,
		modelName:    "none",
		providerName: "none",
		history:      []session.Message{},
	}
}

func (m *AppModel) SetProviderService(svc *provider.Service, modelRepo *sqliterepo.ModelRepo, sessionRepo *sqliterepo.SessionRepo, messageRepo *sqliterepo.MessageRepo) {
	m.providerSvc = svc
	m.modelRepo = modelRepo
	m.sessionRepo = sessionRepo
	m.messageRepo = messageRepo
	m.modelSvc = modelapp.NewService(modelRepo)
	m.commands = newCommandRegistry(svc, m.modelSvc, m)
	runtimeRouter := approuter.NewRuntimeRouter(svc, m.modelSvc)
	m.chatSvc = conversation.NewService(runtimeRouter)

	p, err := svc.GetDefault(m.ctx)
	if err == nil && p != nil {
		m.providerName = p.Name
		m.homeScreen.UpdateConfig(screens.HomeScreenConfig{
			ProviderName: p.Name,
			ProviderURL:  p.BaseURL,
		})
		m.statusBar.SetModel(p.Name)
		m.chatScreen.SetModel(p.Name)
	}

	models, _ := m.modelSvc.List(m.ctx)
	if len(models) > 0 {
		m.modelName = models[0].ModelID
		m.chatSvc.SetModelID(m.modelName)
		m.homeScreen.UpdateConfig(screens.HomeScreenConfig{
			ModelName: m.modelName,
		})
		m.statusBar.SetModel(m.providerName + "/" + m.modelName)
	}

	m.detectGitBranch()
	m.loadLastSession()
}

func (m *AppModel) SetMessageRepo(repo *sqliterepo.MessageRepo) {
	m.messageRepo = repo
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

	if m.messageRepo != nil {
		msgs, err := m.messageRepo.ListBySession(m.ctx, m.currentSess.ID)
		if err == nil {
			m.history = make([]session.Message, len(msgs))
			for i, msg := range msgs {
				m.history[i] = *msg
				m.chatScreen.AddMessage(string(msg.Role), msg.Content)
			}
			if len(msgs) > 0 {
				m.screen = screenChat
			}
		}
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
		m.sessionRepo.Create(m.ctx, s)
	}
}

func (m *AppModel) saveMessage(role session.Role, content string) {
	if m.currentSess == nil {
		return
	}
	msg := session.NewMessage(m.currentSess.ID, role, content)
	m.history = append(m.history, *msg)

	if m.messageRepo != nil {
		m.messageRepo.Create(m.ctx, msg)
	}

	if m.sessionRepo != nil && m.currentSess != nil {
		m.currentSess.MessageCnt = len(m.history)
		m.currentSess.UpdatedAt = time.Now()
		m.sessionRepo.Update(m.ctx, m.currentSess)
	}
}

func (m *AppModel) SetSelectedModel(modelID string) {
	m.modelName = modelID
	m.chatSvc.SetModelID(modelID)
	m.homeScreen.UpdateConfig(screens.HomeScreenConfig{
		ModelName: modelID,
	})
	m.statusBar.SetModel(m.providerName + "/" + modelID)
	m.chatScreen.SetModel(m.providerName + "/" + modelID)
}

func (m *AppModel) SetProgram(p *tea.Program) {
	m.program = p
}

func (m *AppModel) ShowDialog(d screens.DialogScreen) {
	d.SetSize(m.width, m.height)
	m.activeDialog = d
	m.layout()
}

func (m *AppModel) SetWorkspace(path string) {
	m.workspace = path
	m.homeScreen.UpdateConfig(screens.HomeScreenConfig{WorkspacePath: path})
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
	// AppModel.View() = vpH + paletteH + input(3) + status(1) + safe(1) = vpH + paletteH + 5
	overheadVp := 5 + paletteH
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
}

func (m *AppModel) Init() tea.Cmd {
	return tea.Batch(
		m.commandInput.Init(),
	)
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.activeDialog != nil {
		d, cmd := m.activeDialog.Update(msg)
		m.activeDialog = d
		if cmd != nil {
			return m, cmd
		}
		if m.activeDialog.Done() {
			result := m.activeDialog.Result()
			m.activeDialog = nil
			if result != "" && !strings.HasPrefix(result, "__") {
				if m.screen == screenHome {
					m.screen = screenChat
				}
				m.chatScreen.AddMessage("system", result)
				m.saveMessage(session.RoleSystem, result)
			}
			m.layout()
		}
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true

		m.homeScreen.SetSize(msg.Width, msg.Height)
		m.layout()

		return m, nil

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
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

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
			case "__clear__", "__home__":
				m.screen = screenHome
				m.working = false
				m.statusBar.SetWorking(false)
			case "__dialog__":
				return m, nil
			default:
				if m.screen == screenHome {
					m.screen = screenChat
				}
				m.chatScreen.AddMessage("system", result)
				m.saveMessage(session.RoleSystem, result)
			}
			return m, nil
		}

		if m.screen == screenHome {
			m.screen = screenChat
			m.layout()
		}

		m.ensureSession()
		m.chatScreen.AddMessage("user", input)
		m.saveMessage(session.RoleUser, input)

		m.working = true
		m.streamBuf = ""
		m.statusBar.SetWorking(true)

		historyCopy := make([]session.Message, len(m.history))
		copy(historyCopy, m.history)
		prog := m.program

		m.chatScreen.ClearToolCards()
		m.chatScreen.ShowThinking()

		go func() {
			onToolEvent := func(ev conversation.ToolEvent) {
				if prog == nil {
					return
				}
				switch ev.Type {
				case conversation.ToolQueued:
					prog.Send(screens.ToolQueuedMsg{Index: ev.Index, Name: ev.Name, Args: ev.Args})
				case conversation.ToolInitializing:
					prog.Send(screens.ToolInitializingMsg{Index: ev.Index, Name: ev.Name})
				case conversation.ToolConnecting:
					prog.Send(screens.ToolConnectingMsg{Index: ev.Index, Name: ev.Name})
				case conversation.ToolStarted:
					prog.Send(screens.ToolStartedMsg{Index: ev.Index, Name: ev.Name})
				case conversation.ToolOutput:
					prog.Send(screens.ToolOutputMsg{Index: ev.Index, Name: ev.Name, Output: ev.Output})
				case conversation.ToolCompleted:
					prog.Send(screens.ToolCompletedMsg{Index: ev.Index, Name: ev.Name, Output: ev.Output, Status: ev.Status, Error: ev.Error, Duration: ev.Duration})
				}
			}

			result, err := m.chatSvc.SendMessage(m.ctx, input, historyCopy, func(chunk string, done bool) {
				if prog != nil {
					prog.Send(streamContentMsg(chunk))
				}
			}, onToolEvent)
			if prog != nil {
				tokIn, tokOut := 0, 0
				if result != nil {
					tokIn = result.InputTokens
					tokOut = result.OutputTokens
				}
				content := ""
				if result != nil {
					content = result.Content
				}
				prog.Send(streamDoneMsg{content: content, inputTokens: tokIn, outputTokens: tokOut, err: err})
			}
		}()

		return m, nil

	case streamContentMsg:
		m.streamBuf += string(msg)
		return m, nil

	case streamDoneMsg:
		m.working = false
		m.statusBar.SetWorking(false)
		m.chatScreen.HideThinking()
		if msg.err != nil {
			m.chatScreen.ClearToolCards()
			m.chatScreen.AddMessage("system", "Error: "+msg.err.Error())
			m.saveMessage(session.RoleSystem, "Error: "+msg.err.Error())
		} else {
			m.chatScreen.SetToolMetrics(m.chatScreen.ToolCount(), m.chatScreen.TotalDuration(), msg.inputTokens, msg.outputTokens)
			if msg.content != "" {
				m.chatScreen.AddMessage("assistant", msg.content)
				m.saveMessage(session.RoleAssistant, msg.content)
			}
		}
		return m, nil

	default:
		m.statusBar.Update(msg)
	}

	chatCmd := m.chatScreen.Update(msg)

	var cmd tea.Cmd
	m.commandInput, cmd = m.commandInput.Update(msg)
	if chatCmd != nil {
		cmd = tea.Batch(chatCmd, cmd)
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
	if !m.ready {
		return tea.NewView("Initializing...")
	}

	statusView := m.statusBar.View()

	var mainContent string
	if m.activeDialog != nil {
		mainContent = m.activeDialog.View()
		dialogLines := strings.Count(mainContent, "\n") + 1
		overhead := 3
		if topPad := (m.height - dialogLines - overhead) / 2; topPad > 0 {
			mainContent = strings.Repeat("\n", topPad) + mainContent
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

	var result string
	if paletteView != "" {
		result = mainContent + "\n" + paletteView + "\n" + inputView + "\n" + statusView
	} else {
		result = mainContent + "\n" + inputView + "\n" + statusView
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

var _ = fmt.Sprintf
