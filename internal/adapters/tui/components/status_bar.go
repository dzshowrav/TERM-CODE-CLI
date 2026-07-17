package components

import (
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var (
	statusStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	modelStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	agentStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("141"))
	branchStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("83"))
	versionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	barEmpty     = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
	barFilled    = lipgloss.NewStyle().Foreground(lipgloss.Color("221"))
	sepStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
)

type StatusBar struct {
	frame     int
	working   bool
	modelName string
	agentName string
	version   string
	branch    string
	width     int
}

type workingTickMsg time.Time

const tickInterval = 150 * time.Millisecond

var barFrames = []string{
	"[⬝⬝⬝⬝⬝⬝]",
	"[■⬝⬝⬝⬝⬝]",
	"[■■⬝⬝⬝⬝]",
	"[■■■⬝⬝⬝]",
	"[■■■■⬝⬝]",
	"[■■■■■⬝]",
	"[■■■■■■]",
	"[■■■■■⬝]",
	"[■■■■⬝⬝]",
	"[■■■⬝⬝⬝]",
	"[■■⬝⬝⬝⬝]",
	"[■⬝⬝⬝⬝⬝]",
}

type StatusBarConfig struct {
	ModelName string
	AgentName string
	Version   string
}

func NewStatusBar(cfg StatusBarConfig) *StatusBar {
	return &StatusBar{
		frame:     0,
		working:   false,
		modelName: cfg.ModelName,
		agentName: cfg.AgentName,
		version:   cfg.Version,
		width:     80,
	}
}

func (b *StatusBar) SetWorking(v bool) {
	b.working = v
	if v {
		b.frame = 0
	}
}

func (b *StatusBar) SetModel(name string) {
	b.modelName = name
}

func (b *StatusBar) SetAgent(name string) {
	b.agentName = name
}

func (b *StatusBar) SetBranch(name string) {
	b.branch = name
}

func (b *StatusBar) SetWidth(w int) {
	b.width = w
}

func (b *StatusBar) Working() bool {
	return b.working
}

func (b *StatusBar) Init() tea.Cmd {
	return b.tick()
}

func (b *StatusBar) tick() tea.Cmd {
	if !b.working {
		return nil
	}
	return tea.Tick(tickInterval, func(t time.Time) tea.Msg {
		return workingTickMsg(t)
	})
}

func (b *StatusBar) Update(msg tea.Msg) tea.Cmd {
	switch msg.(type) {
	case workingTickMsg:
		if !b.working {
			return nil
		}
		b.frame = (b.frame + 1) % len(barFrames)
		return b.tick()
	}
	return nil
}

func (b *StatusBar) View() string {
	if !b.working {
		return b.idleView()
	}
	return b.workingView()
}

func (b *StatusBar) progressBar() string {
	frame := barFrames[b.frame]
	return barFilled.Render(frame[:4]) + barEmpty.Render(frame[4:])
}

func (b *StatusBar) idleView() string {
	model := modelStyle.Render(b.modelName)
	agent := agentStyle.Render(b.agentName)
	ver := versionStyle.Render(b.version)
	sep := sepStyle.Render(" │ ")

	if b.width < 25 {
		return model
	}

	var branch string
	if b.branch != "" {
		branch = branchStyle.Render(b.branch) + sep
	}

	if b.width < 40 {
		return statusStyle.Render(branch + model + sep + agent)
	}

	if b.width < 60 {
		return statusStyle.Render(branch + model + sep + agent)
	}

	return statusStyle.Render(branch + model + sep + agent + sep + ver)
}

func (b *StatusBar) workingView() string {
	bar := b.progressBar()
	model := modelStyle.Render(b.modelName)
	agent := agentStyle.Render(b.agentName)
	ver := versionStyle.Render(b.version)
	sep := sepStyle.Render(" │ ")

	if b.width < 25 {
		return bar + sep + model
	}

	var branch string
	if b.branch != "" {
		branch = branchStyle.Render(b.branch) + sep
	}

	if b.width < 40 {
		return bar + sep + branch + model + sep + agent
	}

	if b.width < 60 {
		return bar + sep + branch + model + sep + agent
	}

	return bar + sep + branch + model + sep + agent + sep + ver
}
