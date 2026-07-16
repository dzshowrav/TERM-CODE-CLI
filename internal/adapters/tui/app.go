package tui

import (
	"context"
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"

	"termcode/internal/adapters/tui/components"
	"termcode/internal/adapters/tui/screens"
	"termcode/internal/application/conversation"
	modelapp "termcode/internal/application/model"
	"termcode/internal/application/provider"
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
	content string
	err     error
}

type AppModel struct {
	width        int
	height       int
	ready        bool
	screen       screenType
	homeScreen   *screens.HomeScreen
	chatScreen   *screens.ChatScreen
	commandInput *components.CommandInput
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
	m.chatSvc = conversation.NewService(svc)

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

func (m *AppModel) Init() tea.Cmd {
	return tea.Batch(
		m.commandInput.Init(),
	)
}

func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true

		m.homeScreen.SetSize(msg.Width, msg.Height)
		m.chatScreen.SetSize(msg.Width, msg.Height)
		m.commandInput.SetWidth(msg.Width)
		m.statusBar.SetWidth(msg.Width)

		return m, nil

	case tea.KeyMsg:
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
			case "__clear__", "__home__":
				m.screen = screenHome
				m.working = false
				m.statusBar.SetWorking(false)
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

		go func() {
			content, err := m.chatSvc.SendMessage(m.ctx, input, historyCopy, func(chunk string, done bool) {
				if prog != nil {
					prog.Send(streamContentMsg(chunk))
				}
			})
			if prog != nil {
				prog.Send(streamDoneMsg{content: content, err: err})
			}
		}()

		return m, nil

	case streamContentMsg:
		m.streamBuf += string(msg)
		return m, nil

	case streamDoneMsg:
		m.working = false
		m.statusBar.SetWorking(false)
		if msg.err != nil {
			m.chatScreen.AddMessage("system", "Error: "+msg.err.Error())
			m.saveMessage(session.RoleSystem, "Error: "+msg.err.Error())
		} else {
			m.chatScreen.AddMessage("assistant", msg.content)
			m.saveMessage(session.RoleAssistant, msg.content)
		}
		return m, nil

	default:
		m.statusBar.Update(msg)
	}

	var cmd tea.Cmd
	m.commandInput, cmd = m.commandInput.Update(msg)
	return m, cmd
}

func (m *AppModel) View() tea.View {
	if !m.ready {
		return tea.NewView("Initializing...")
	}

	inputView := m.commandInput.View()
	statusView := m.statusBar.View()

	var mainContent string
	if m.screen == screenHome {
		mainContent = m.homeScreen.View()
	} else {
		mainContent = m.chatScreen.View()
	}

	result := mainContent + "\n" + inputView + "\n" + statusView
	return tea.NewView(result)
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
