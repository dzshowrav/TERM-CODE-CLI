package tool

import (
	"context"
	"time"
)

type Status string

const (
	StatusPending  Status = "pending"
	StatusRunning  Status = "running"
	StatusApproved Status = "approved"
	StatusDenied   Status = "denied"
	StatusSuccess  Status = "success"
	StatusFailed   Status = "failed"
	StatusTimeout  Status = "timeout"
	StatusAborted  Status = "aborted"
)

type Category string

const (
	CatFileSystem  Category = "filesystem"
	CatDevelopment Category = "development"
	CatGit         Category = "git"
	CatSearch      Category = "search"
	CatNetwork     Category = "network"
	CatCode        Category = "code_analysis"
	CatTerminal    Category = "terminal"
	CatUtility     Category = "utility"
	CatPlugin      Category = "plugin"
	CatDatabase    Category = "database"
	CatBrowser     Category = "browser"
	CatDangerous   Category = "dangerous"
)

type Capability string

const (
	CapRead     Capability = "read"
	CapWrite    Capability = "write"
	CapExecute  Capability = "execute"
	CapNetwork  Capability = "network"
	CapFiles    Capability = "files"
	CapSearch   Capability = "search"
	CapGit      Capability = "git"
	CapDatabase Capability = "database"
	CapBrowser  Capability = "browser"
	CapDanger   Capability = "dangerous"
)

type AllowedContext string

const (
	CtxFilesystem  AllowedContext = "filesystem"
	CtxNetwork     AllowedContext = "network"
	CtxEnvironment AllowedContext = "environment"
	CxTerminal     AllowedContext = "terminal"
)

type PermissionLevel string

const (
	PermAlwaysAllow PermissionLevel = "always_allow"
	PermAllowOnce   PermissionLevel = "allow_once"
	PermAsk         PermissionLevel = "ask"
	PermDeny        PermissionLevel = "deny"
)

type HookFunc func(ctx context.Context, args map[string]any)

type ToolHooks struct {
	OnBefore func(ctx context.Context, args map[string]any) error
	OnStart  func(ctx context.Context, args map[string]any)
	OnDelta  func(ctx context.Context, delta string)
	OnEnd    func(ctx context.Context, result *Result)
	OnError  func(ctx context.Context, err error)
	OnAbort  func(ctx context.Context)
}

type Tool struct {
	Name            string           `json:"name"`
	DisplayName     string           `json:"display_name,omitempty"`
	Description     string           `json:"description"`
	InputSchema     any              `json:"input_schema"`
	Category        Category         `json:"category"`
	Aliases         []string         `json:"aliases,omitempty"`
	Version         string           `json:"version,omitempty"`
	Author          string           `json:"author,omitempty"`
	Capabilities    []Capability     `json:"capabilities,omitempty"`
	DefaultTimeout  int              `json:"default_timeout_ms,omitempty"`
	AllowedContexts []AllowedContext `json:"allowed_contexts,omitempty"`
	PermissionLevel PermissionLevel  `json:"permission_level,omitempty"`
	RateLimit       int              `json:"rate_limit_per_minute,omitempty"`
	Dangerous       bool             `json:"dangerous,omitempty"`
	Hooks           ToolHooks        `json:"-"`
	Source          string           `json:"source,omitempty"`
}

func New(name, description string, schema any, category Category, capabilities ...Capability) Tool {
	return Tool{
		Name:           name,
		DisplayName:    name,
		Description:    description,
		InputSchema:    schema,
		Category:       category,
		Capabilities:   capabilities,
		Version:        "1.0.0",
		Author:         "built-in",
		Source:         "built-in",
		DefaultTimeout: 30000,
		RateLimit:      60,
	}
}

func (t Tool) MatchName(s string) bool {
	if t.Name == s {
		return true
	}
	for _, a := range t.Aliases {
		if a == s {
			return true
		}
	}
	return false
}

type Result struct {
	Tool      string    `json:"tool"`
	Input     string    `json:"input"`
	Output    string    `json:"output"`
	Error     string    `json:"error,omitempty"`
	Status    Status    `json:"status"`
	Duration  int64     `json:"duration_ms"`
	Time      time.Time `json:"time"`
	Truncated bool      `json:"truncated,omitempty"`
	RawSize   int       `json:"raw_size,omitempty"`
}

func (r *Result) Truncate(maxBytes int) {
	if len(r.Output) > maxBytes {
		r.RawSize = len(r.Output)
		r.Output = r.Output[:maxBytes] + "... (truncated)"
		r.Truncated = true
	}
}
