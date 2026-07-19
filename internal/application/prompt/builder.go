package prompt

import (
	"fmt"
	"strings"
	"time"
)

type ContextLevel int

const (
	ContextMinimal ContextLevel = iota
	ContextNormal
	ContextDetailed
)

type Builder struct {
	agentName    string
	provider     string
	model        string
	workspace    string
	gitBranch    string
	tools        []ToolInfo
	systemPrompt string
	customRules  []string
	maxTokens    int
}

type ToolInfo struct {
	Name        string
	Description string
	Schema      string
}

func NewBuilder() *Builder {
	return &Builder{
		agentName: "General",
		maxTokens: 4096,
	}
}

func (b *Builder) SetAgent(name string)     { b.agentName = name }
func (b *Builder) SetProvider(p string)     { b.provider = p }
func (b *Builder) SetModel(m string)        { b.model = m }
func (b *Builder) SetWorkspace(w string)    { b.workspace = w }
func (b *Builder) SetGitBranch(br string)   { b.gitBranch = br }
func (b *Builder) SetSystemPrompt(p string) { b.systemPrompt = p }
func (b *Builder) SetMaxTokens(n int)       { b.maxTokens = n }
func (b *Builder) AddCustomRule(r string)   { b.customRules = append(b.customRules, r) }

func (b *Builder) SetTools(tools []ToolInfo) {
	b.tools = tools
}

func (b *Builder) Build(level ContextLevel) string {
	var parts []string

	role := b.buildRole()
	if role != "" {
		parts = append(parts, role)
	}

	if b.systemPrompt != "" {
		parts = append(parts, b.systemPrompt)
	}

	custom := b.buildCustomRules()
	if custom != "" {
		parts = append(parts, custom)
	}

	context := b.buildContext(level)
	if context != "" {
		parts = append(parts, context)
	}

	toolsSection := b.buildTools()
	if toolsSection != "" {
		parts = append(parts, toolsSection)
	}

	return strings.Join(parts, "\n\n")
}

func (b *Builder) buildRole() string {
	if b.agentName == "" || b.agentName == "General" {
		return ""
	}

	rolePrompts := map[string]string{
		"Expert":    "You are a senior software engineer. Provide expert-level guidance.",
		"Architect": "You are a software architect. Focus on system design and architecture.",
	}

	if p, ok := rolePrompts[b.agentName]; ok {
		return p
	}

	return fmt.Sprintf("You are acting as '%s'. Respond accordingly.", b.agentName)
}

func (b *Builder) buildCustomRules() string {
	if len(b.customRules) == 0 {
		return ""
	}
	return "Custom rules:\n- " + strings.Join(b.customRules, "\n- ")
}

func (b *Builder) buildContext(level ContextLevel) string {
	if level == ContextMinimal {
		return ""
	}

	var ctx []string
	ctx = append(ctx, fmt.Sprintf("Current time: %s", time.Now().Format(time.RFC1123)))

	if b.provider != "" {
		ctx = append(ctx, fmt.Sprintf("Provider: %s", b.provider))
	}
	if b.model != "" {
		ctx = append(ctx, fmt.Sprintf("Model: %s", b.model))
	}
	if b.workspace != "" {
		ctx = append(ctx, fmt.Sprintf("Workspace: %s", b.workspace))
	}
	if b.gitBranch != "" {
		ctx = append(ctx, fmt.Sprintf("Git branch: %s", b.gitBranch))
	}

	if level >= ContextDetailed && b.workspace != "" {
		ctx = append(ctx, fmt.Sprintf("All commands execute in workspace: %s", b.workspace))
		ctx = append(ctx, "Use relative paths when referring to files in the workspace.")
	}

	return strings.Join(ctx, "\n")
}

func (b *Builder) buildTools() string {
	if len(b.tools) == 0 {
		return ""
	}

	var parts []string
	parts = append(parts, "Available tools:")

	for _, t := range b.tools {
		desc := t.Name
		if t.Description != "" {
			desc += ": " + t.Description
		}
		parts = append(parts, "  - "+desc)
	}

	parts = append(parts, "")
	parts = append(parts, "To use a tool, include a tool_call in your response.")

	return strings.Join(parts, "\n")
}

func (b *Builder) BuildSystemMessage() string {
	return b.Build(ContextNormal)
}
