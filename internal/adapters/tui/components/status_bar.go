package components

import (
	"strings"
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

type StatusBarField int

const (
	FieldModel StatusBarField = iota
	FieldAgent
	FieldBranch
	FieldVersion
)

type StatusBar struct {
	frame     int
	working   bool
	modelName string
	agentName string
	version   string
	branch    string
	width     int
	fields    []StatusBarField
}

var defaultFields = []StatusBarField{FieldBranch, FieldModel, FieldAgent, FieldVersion}

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
	fields := make([]StatusBarField, len(defaultFields))
	copy(fields, defaultFields)
	return &StatusBar{
		frame:     0,
		working:   false,
		modelName: cfg.ModelName,
		agentName: cfg.AgentName,
		version:   cfg.Version,
		width:     80,
		fields:    fields,
	}
}

func (b *StatusBar) SetFields(fields []StatusBarField) {
	if len(fields) == 0 {
		b.fields = make([]StatusBarField, len(defaultFields))
		copy(b.fields, defaultFields)
		return
	}
	b.fields = fields
}

func (b *StatusBar) SetWorking(v bool) tea.Cmd {
	b.working = v
	if v {
		b.frame = 0
		return b.tick()
	}
	return nil
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
	runes := []rune(barFrames[b.frame])
	var sb strings.Builder
	for _, r := range runes {
		if r == '■' {
			sb.WriteString(barFilled.Render(string(r)))
		} else {
			sb.WriteString(barEmpty.Render(string(r)))
		}
	}
	return sb.String()
}

func (b *StatusBar) idleView() string {
	sep := sepStyle.Render(" │ ")
	parts := b.renderFields(sep)
	if len(parts) == 0 {
		return modelStyle.Render(b.modelName)
	}
	return statusStyle.Render(strings.Join(parts, sep))
}

func (b *StatusBar) workingView() string {
	bar := b.progressBar()
	sep := sepStyle.Render(" │ ")
	parts := b.renderFields(sep)
	if len(parts) == 0 {
		return bar + sep + modelStyle.Render(b.modelName)
	}
	return bar + sep + statusStyle.Render(strings.Join(parts, sep))
}

func (b *StatusBar) renderFields(sep string) []string {
	var parts []string
	for _, f := range b.fields {
		switch f {
		case FieldModel:
			parts = append(parts, modelStyle.Render(b.modelName))
		case FieldAgent:
			parts = append(parts, agentStyle.Render(b.agentName))
		case FieldBranch:
			if b.branch != "" {
				parts = append(parts, branchStyle.Render(b.branch))
			}
		case FieldVersion:
			parts = append(parts, versionStyle.Render(b.version))
		}

	}
	return parts
}
