package tool

import (
	"context"
	"database/sql"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"

	"termcode/internal/domain/permission"
	"termcode/internal/domain/tool"
	"termcode/internal/infrastructure/executor"
	git "termcode/internal/infrastructure/git"

	_ "modernc.org/sqlite"
)

const maxOutputBytes = 100 * 1024

type PermissionChecker interface {
	IsAllowed(toolName string) bool
	IsDenied(toolName string) bool
	Request(toolName, args string) string
}

type Service struct {
	registry       *tool.Registry
	shell          *executor.ShellExecutor
	files          *executor.FileExecutor
	git            *git.Service
	permChecker    PermissionChecker
	undo           *executor.UndoService
	memoryData     map[string]string
	mu             sync.Mutex
	outputCallback func(string)
}

func NewService() *Service {
	s := &Service{
		registry:   tool.NewRegistry(),
		shell:      executor.NewShellExecutor(),
		files:      executor.NewFileExecutor(),
		git:        git.NewService(),
		undo:       executor.NewUndoService(),
		memoryData: make(map[string]string),
	}
	if err := s.registerBuiltins(); err != nil {
		panic(fmt.Sprintf("register built-in tools: %v", err))
	}
	return s
}

func (s *Service) SetPermissionChecker(pc PermissionChecker) {
	s.permChecker = pc
}

func (s *Service) SetTrustManager(tm *TrustManager) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if pc, ok := s.permChecker.(*StorePermissionChecker); ok {
		pc.SetTrustManager(tm)
	}
}

func (s *Service) SetOutputCallback(cb func(string)) {
	s.mu.Lock()
	s.outputCallback = cb
	s.mu.Unlock()
}

func (s *Service) GetOutputCallback() func(string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.outputCallback
}

func (s *Service) PermissionChecker() PermissionChecker {
	return s.permChecker
}

func (s *Service) Registry() *tool.Registry {
	return s.registry
}

func (s *Service) AvailableTools() []tool.Tool {
	return s.registry.List()
}

func (s *Service) Lookup(name string) (tool.Tool, bool) {
	return s.registry.Lookup(name)
}

func (s *Service) Execute(ctx context.Context, toolName string, input map[string]any) *tool.Result {
	start := time.Now()

	result := &tool.Result{
		Tool:   toolName,
		Input:  fmt.Sprintf("%v", input),
		Status: tool.StatusRunning,
		Time:   start,
	}

	t, ok := s.registry.Lookup(toolName)
	if !ok {
		result.Status = tool.StatusFailed
		result.Error = fmt.Sprintf("unknown tool: %s", toolName)
		result.Duration = time.Since(start).Milliseconds()
		return result
	}

	// 1. Validate
	if val := tool.ValidateArgs(t.InputSchema, input); !val.Valid {
		result.Status = tool.StatusFailed
		result.Error = fmt.Sprintf("validation: %s", val.Error())
		result.Duration = time.Since(start).Milliseconds()
		return result
	}

	// 2. Permission check
	if s.permChecker != nil {
		if s.permChecker.IsDenied(toolName) {
			result.Status = tool.StatusDenied
			result.Error = fmt.Sprintf("tool %q is denied by permission policy", toolName)
			result.Duration = time.Since(start).Milliseconds()
			return result
		}
		if !s.permChecker.IsAllowed(toolName) {
			argsJSON, _ := json.Marshal(input)
			action := s.permChecker.Request(toolName, string(argsJSON))
			switch action {
			case "deny", "":
				result.Status = tool.StatusDenied
				result.Error = fmt.Sprintf("tool %q denied by user", toolName)
				result.Duration = time.Since(start).Milliseconds()
				return result
			case "allow_once":
				// Allow this one time; don't persist the decision
			case "always_allow":
				if pc, ok := s.permChecker.(*StorePermissionChecker); ok {
					pc.Set(toolName, permission.LevelAlwaysAllow)
				}
			}
		}
	}

	// 3. OnBefore hook
	if t.Hooks.OnBefore != nil {
		if err := t.Hooks.OnBefore(ctx, input); err != nil {
			result.Status = tool.StatusFailed
			result.Error = fmt.Sprintf("before hook: %s", err)
			result.Duration = time.Since(start).Milliseconds()
			return result
		}
	}

	s.executeWithRetry(ctx, toolName, input, result, t)

	return result
}

// ── executeTool ────────────────────────────────────────────
// Runs the tool executor in a goroutine with timeout, populates result in-place.
// Does NOT fire OnEnd/OnError hooks — that is the caller's responsibility.
func (s *Service) executeTool(ctx context.Context, toolName string, input map[string]any, result *tool.Result, t tool.Tool) {
	timeout := t.DefaultTimeout
	if timeout <= 0 {
		timeout = 30000
	}
	execCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Millisecond)
	defer cancel()

	done := make(chan struct{}, 1)
	var execErr error

	go func() {
		defer func() {
			if r := recover(); r != nil {
				execErr = fmt.Errorf("panic: %v", r)
			}
			done <- struct{}{}
		}()
		switch t.Name {
		case "read":
			execErr = s.execRead(input, result)
		case "write":
			execErr = s.execWrite(input, result)
		case "edit":
			execErr = s.execEdit(input, result)
		case "create_file":
			execErr = s.execCreateFile(input, result)
		case "delete_file":
			execErr = s.execDeleteFile(input, result)
		case "replace":
			execErr = s.execReplace(input, result)
		case "rename", "move":
			execErr = s.execRename(input, result)
		case "copy":
			execErr = s.execCopy(input, result)
		case "patch":
			execErr = s.execPatch(input, result)
		case "undo":
			execErr = s.execUndo(result)
		case "redo":
			execErr = s.execRedo(result)
		case "read_head":
			execErr = s.execReadHead(input, result)
		case "read_tail":
			execErr = s.execReadTail(input, result)
		case "read_range":
			execErr = s.execReadRange(input, result)
		case "list_dir":
			execErr = s.execListDir(input, result)
		case "file_info":
			execErr = s.execFileInfo(input, result)
		case "search":
			execErr = s.execSearch(execCtx, input, result)
		case "glob":
			execErr = s.execGlob(input, result)
		case "bash":
			execErr = s.execBash(execCtx, input, result)
		case "git_status":
			execErr = s.execGitStatus(input, result)
		case "git_log":
			execErr = s.execGitLog(input, result)
		case "git_diff":
			execErr = s.execGitDiff(input, result)
		case "git_add":
			execErr = s.execGitAdd(input, result)
		case "git_commit":
			execErr = s.execGitCommit(input, result)
		case "git_branches":
			execErr = s.execGitBranches(input, result)
		case "git_checkout":
			execErr = s.execGitCheckout(input, result)
		case "ask_user":
			execErr = fmt.Errorf("ask_user requires user interaction and is not available in tool mode")
		case "fetch_url":
			execErr = s.execFetchURL(input, result)
		case "calculator":
			execErr = s.execCalculate(input, result)
		case "json":
			execErr = s.execJSON(input, result)
		case "xml":
			execErr = s.execXML(input, result)
		case "memory":
			execErr = s.execMemory(input, result)
		case "database":
			execErr = s.execDatabase(input, result)
		case "image_analysis":
			execErr = s.execImageAnalysis(input, result)
		case "ocr":
			execErr = s.execOCR(input, result)
		case "browser":
			execErr = s.execBrowser(input, result)
		default:
			execErr = fmt.Errorf("unknown tool: %s", toolName)
		}
	}()

	select {
	case <-done:
		if execErr != nil {
			result.Status = tool.StatusFailed
			result.Error = execErr.Error()
		} else if result.Status == tool.StatusRunning {
			result.Status = tool.StatusSuccess
		}
	case <-execCtx.Done():
		if execCtx.Err() == context.DeadlineExceeded {
			result.Status = tool.StatusTimeout
			result.Error = fmt.Sprintf("timeout after %dms", timeout)
		} else {
			result.Status = tool.StatusAborted
			result.Error = "execution cancelled"
		}
		if t.Hooks.OnAbort != nil {
			t.Hooks.OnAbort(ctx)
		}
	}

	result.Duration = time.Since(result.Time).Milliseconds()
	result.Truncate(maxOutputBytes)
}

// ── executeWithRetry ───────────────────────────────────────
// Calls executeTool and retries on transient errors (timeout, connection, etc.).
// Fires OnStart before each attempt, OnEnd/OnError after all attempts.
func (s *Service) executeWithRetry(ctx context.Context, toolName string, input map[string]any, result *tool.Result, t tool.Tool) {
	if t.Hooks.OnStart != nil {
		t.Hooks.OnStart(ctx, input)
	}

	retryErr := retry.Do(
		func() error {
			result.Status = tool.StatusRunning
			result.Error = ""
			result.Output = ""
			result.Truncated = false

			s.executeTool(ctx, toolName, input, result, t)

			if cb := s.GetOutputCallback(); cb != nil && result.Output != "" {
				cb(result.Output)
			}

			if result.Error == "" && result.Status == tool.StatusSuccess {
				return nil
			}
			if result.Error == "" {
				return fmt.Errorf("tool %s failed with status %s", toolName, result.Status)
			}
			return fmt.Errorf("%s", result.Error)
		},
		retry.Attempts(2),
		retry.Delay(1*time.Second),
		retry.RetryIf(func(err error) bool {
			errStr := err.Error()
			return strings.Contains(errStr, "timeout") ||
				strings.Contains(errStr, "connection") ||
				strings.Contains(errStr, "temporary") ||
				strings.Contains(errStr, "rate limit")
		}),
		retry.LastErrorOnly(true),
	)

	if retryErr != nil && result.Error == "" {
		result.Error = retryErr.Error()
		result.Status = tool.StatusFailed
	}

	if result.Status == tool.StatusFailed || result.Status == tool.StatusTimeout || result.Status == tool.StatusAborted {
		if t.Hooks.OnError != nil && result.Error != "" {
			t.Hooks.OnError(ctx, fmt.Errorf("%s", result.Error))
		}
	} else if t.Hooks.OnEnd != nil {
		t.Hooks.OnEnd(ctx, result)
	}
}

// ── Built-in Registration ──────────────────────────────────

func (s *Service) registerBuiltins() error {
	return s.registry.RegisterAll([]tool.Tool{
		// ── FileSystem ──
		func() tool.Tool {
			t := tool.New("read", "Read the complete contents of a file", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "File path to read"},
				},
				"required": []any{"path"},
			}, tool.CatFileSystem, tool.CapRead, tool.CapFiles)
			t.DisplayName = "Read File"
			t.AllowedContexts = []tool.AllowedContext{tool.CtxFilesystem}
			t.Aliases = []string{"read_file", "cat"}
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("write", "Write content to a file (creates directories if needed)", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":    map[string]any{"type": "string", "description": "File path"},
					"content": map[string]any{"type": "string", "description": "File content"},
				},
				"required": []any{"path", "content"},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.Aliases = []string{"write_file"}
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("edit", "Edit a file by replacing exact text matches", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "File path"},
					"edits": map[string]any{"type": "array", "items": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"old_str": map[string]any{"type": "string", "description": "Text to replace"},
							"new_str": map[string]any{"type": "string", "description": "Replacement text"},
						},
						"required": []any{"old_str", "new_str"},
					}, "description": "Array of edit operations"},
				},
				"required": []any{"path", "edits"},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.Aliases = []string{"search_replace"}
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("create_file", "Create a new empty file or directory", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "File path"},
					"type": map[string]any{"type": "string", "enum": []any{"file", "dir"}, "description": "Type to create"},
				},
				"required": []any{"path"},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.DefaultTimeout = 5000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("delete_file", "Delete a file or empty directory", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "File path to delete"},
				},
				"required": []any{"path"},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.DefaultTimeout = 5000
			t.Dangerous = true
			return t
		}(),
		func() tool.Tool {
			t := tool.New("list_dir", "List files and directories in a path", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "Directory path"},
				},
				"required": []any{"path"},
			}, tool.CatFileSystem, tool.CapRead, tool.CapFiles)
			t.Aliases = []string{"ls", "dir"}
			t.DefaultTimeout = 5000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("file_info", "Get metadata about a file or directory", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "File path"},
				},
				"required": []any{"path"},
			}, tool.CatFileSystem, tool.CapRead, tool.CapFiles)
			t.Aliases = []string{"stat"}
			t.DefaultTimeout = 5000
			return t
		}(),

		func() tool.Tool {
			t := tool.New("replace", "Replace text in a file using exact match or regex", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":     map[string]any{"type": "string", "description": "File path"},
					"pattern":  map[string]any{"type": "string", "description": "Text or regex pattern to find"},
					"new_text": map[string]any{"type": "string", "description": "Replacement text"},
					"is_regex": map[string]any{"type": "boolean", "description": "Treat pattern as regex (default: false)"},
				},
				"required": []any{"path", "pattern", "new_text"},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.Aliases = []string{"regex_replace", "replace_text"}
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("rename", "Rename or move a file or directory", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"source": map[string]any{"type": "string", "description": "Current path"},
					"dest":   map[string]any{"type": "string", "description": "New path"},
				},
				"required": []any{"source", "dest"},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.Aliases = []string{"move", "mv", "rename_file"}
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("copy", "Copy a file or directory", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"source": map[string]any{"type": "string", "description": "Source path"},
					"dest":   map[string]any{"type": "string", "description": "Destination path"},
				},
				"required": []any{"source", "dest"},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.Aliases = []string{"cp", "copy_file"}
			t.DefaultTimeout = 30000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("patch", "Apply a unified diff/patch file to files", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":     map[string]any{"type": "string", "description": "Path to patch file"},
					"content":  map[string]any{"type": "string", "description": "Inline patch content (alternative to path)"},
					"base_dir": map[string]any{"type": "string", "description": "Base directory for relative paths in patch"},
				},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.Aliases = []string{"apply_patch", "diff_apply"}
			t.DefaultTimeout = 30000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("undo", "Undo the last file operation", map[string]any{
				"type":       "object",
				"properties": map[string]any{},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.Aliases = []string{"revert"}
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("redo", "Redo the last undone file operation", map[string]any{
				"type":       "object",
				"properties": map[string]any{},
			}, tool.CatFileSystem, tool.CapWrite, tool.CapFiles)
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("read_head", "Read the first N lines of a file (efficient for large files)", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":      map[string]any{"type": "string", "description": "File path"},
					"max_lines": map[string]any{"type": "integer", "description": "Max lines to read (default: 100)"},
				},
				"required": []any{"path"},
			}, tool.CatFileSystem, tool.CapRead, tool.CapFiles)
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("read_tail", "Read the last N lines of a file (efficient for large files)", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":      map[string]any{"type": "string", "description": "File path"},
					"max_lines": map[string]any{"type": "integer", "description": "Max lines to read (default: 100)"},
				},
				"required": []any{"path"},
			}, tool.CatFileSystem, tool.CapRead, tool.CapFiles)
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("read_range", "Read a specific line range from a file", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":       map[string]any{"type": "string", "description": "File path"},
					"start_line": map[string]any{"type": "integer", "description": "Start line (1-indexed, default: 1)"},
					"end_line":   map[string]any{"type": "integer", "description": "End line (inclusive)"},
				},
				"required": []any{"path", "end_line"},
			}, tool.CatFileSystem, tool.CapRead, tool.CapFiles)
			t.DefaultTimeout = 10000
			return t
		}(),

		// ── Search ──
		func() tool.Tool {
			t := tool.New("search", "Search for files by content using ripgrep", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"pattern": map[string]any{"type": "string", "description": "Search pattern (regex)"},
					"path":    map[string]any{"type": "string", "description": "Directory to search (default: workspace)"},
					"include": map[string]any{"type": "string", "description": "File glob to include (e.g. *.go)"},
				},
				"required": []any{"pattern"},
			}, tool.CatSearch, tool.CapSearch, tool.CapRead)
			t.Aliases = []string{"grep", "find_text"}
			t.DefaultTimeout = 15000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("glob", "Find files by name pattern using glob", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"pattern": map[string]any{"type": "string", "description": "Glob pattern (e.g. **/*.go)"},
					"path":    map[string]any{"type": "string", "description": "Directory to search"},
				},
				"required": []any{"pattern"},
			}, tool.CatSearch, tool.CapSearch, tool.CapRead)
			t.Aliases = []string{"find"}
			t.DefaultTimeout = 10000
			return t
		}(),

		// ── Terminal ──
		func() tool.Tool {
			t := tool.New("bash", "Execute a shell command", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"command": map[string]any{"type": "string", "description": "Shell command to execute"},
				},
				"required": []any{"command"},
			}, tool.CatTerminal, tool.CapExecute)
			t.Aliases = []string{"shell", "sh", "run"}
			t.DisplayName = "Execute Command"
			t.AllowedContexts = []tool.AllowedContext{tool.CxTerminal}
			t.PermissionLevel = tool.PermAsk
			t.DefaultTimeout = 30000
			t.Dangerous = true
			return t
		}(),

		// ── Git ──
		func() tool.Tool {
			t := tool.New("git_status", "Show the working tree status", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "Repository path"},
				},
				"required": []any{"path"},
			}, tool.CatGit, tool.CapGit, tool.CapRead)
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("git_log", "Show commit logs", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":  map[string]any{"type": "string", "description": "Repository path"},
					"count": map[string]any{"type": "integer", "description": "Number of commits"},
				},
				"required": []any{"path"},
			}, tool.CatGit, tool.CapGit, tool.CapRead)
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("git_diff", "Show unstaged changes in working tree", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "Repository path"},
				},
				"required": []any{"path"},
			}, tool.CatGit, tool.CapGit, tool.CapRead)
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("git_add", "Stage files for commit", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":  map[string]any{"type": "string", "description": "Repository path"},
					"files": map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "Files to stage (omit for all)"},
				},
				"required": []any{"path"},
			}, tool.CatGit, tool.CapGit, tool.CapWrite)
			t.DefaultTimeout = 15000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("git_commit", "Commit staged changes", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":    map[string]any{"type": "string", "description": "Repository path"},
					"message": map[string]any{"type": "string", "description": "Commit message"},
				},
				"required": []any{"path", "message"},
			}, tool.CatGit, tool.CapGit, tool.CapWrite)
			t.DefaultTimeout = 15000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("git_branches", "List branches in repository", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "Repository path"},
				},
				"required": []any{"path"},
			}, tool.CatGit, tool.CapGit, tool.CapRead)
			t.Aliases = []string{"git_branch"}
			t.DefaultTimeout = 10000
			return t
		}(),
		func() tool.Tool {
			t := tool.New("git_checkout", "Switch or create branches", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":   map[string]any{"type": "string", "description": "Repository path"},
					"branch": map[string]any{"type": "string", "description": "Branch name"},
					"create": map[string]any{"type": "boolean", "description": "Create new branch"},
				},
				"required": []any{"path", "branch"},
			}, tool.CatGit, tool.CapGit, tool.CapWrite)
			t.DefaultTimeout = 15000
			return t
		}(),

		// ── Utility ──
		func() tool.Tool {
			t := tool.New("ask_user", "Ask the user a question and wait for response", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"question": map[string]any{"type": "string", "description": "Question to ask the user"},
					"options":  map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "Multiple choice options"},
				},
				"required": []any{"question"},
			}, tool.CatUtility, tool.CapRead)
			t.DefaultTimeout = 60000
			return t
		}(),

		// ── Web / HTTP ──
		func() tool.Tool {
			t := tool.New("fetch_url", "Fetch a URL and return its contents. Supports GET and POST requests.", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"url":    map[string]any{"type": "string", "description": "URL to fetch"},
					"method": map[string]any{"type": "string", "enum": []any{"GET", "POST"}, "description": "HTTP method (default: GET)"},
					"body":   map[string]any{"type": "string", "description": "Request body for POST requests"},
				},
				"required": []any{"url"},
			}, tool.CatNetwork, tool.CapNetwork)
			t.Aliases = []string{"http_get", "http_post", "web_fetch", "fetch"}
			t.DefaultTimeout = 30000
			return t
		}(),

		// ── Calculator ──
		func() tool.Tool {
			t := tool.New("calculator", "Evaluate a mathematical expression and return the result. Supports +, -, *, /, parentheses, and functions: sqrt, sin, cos, tan, pow, abs, round, floor, ceil, log, ln.", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"expression": map[string]any{"type": "string", "description": "Mathematical expression to evaluate (e.g. '2 + 2', 'sqrt(16)', 'sin(pi/2)')"},
				},
				"required": []any{"expression"},
			}, tool.CatUtility, tool.CapRead)
			t.Aliases = []string{"calc", "math", "evaluate"}
			t.DefaultTimeout = 5000
			return t
		}(),

		// ── JSON ──
		func() tool.Tool {
			t := tool.New("json", "Process JSON data: validate, format, or query using dot-separated paths.", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"operation": map[string]any{"type": "string", "enum": []any{"validate", "format", "query"}, "description": "Operation to perform"},
					"data":      map[string]any{"type": "string", "description": "JSON data to process"},
					"path":      map[string]any{"type": "string", "description": "Dot-separated path for query (e.g. 'foo.bar.0')"},
				},
				"required": []any{"operation", "data"},
			}, tool.CatCode, tool.CapRead)
			t.Aliases = []string{"json_validate", "json_format", "json_query"}
			t.DefaultTimeout = 5000
			return t
		}(),

		// ── XML ──
		func() tool.Tool {
			t := tool.New("xml", "Process XML data: validate or format.", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"operation": map[string]any{"type": "string", "enum": []any{"validate", "format"}, "description": "Operation to perform"},
					"data":      map[string]any{"type": "string", "description": "XML data to process"},
				},
				"required": []any{"operation", "data"},
			}, tool.CatCode, tool.CapRead)
			t.Aliases = []string{"xml_validate", "xml_format"}
			t.DefaultTimeout = 5000
			return t
		}(),

		// ── Memory (Key-Value Store) ──
		func() tool.Tool {
			t := tool.New("memory", "Store and retrieve key-value pairs in the session memory. Operations: get, set, delete, list, search.", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"operation": map[string]any{"type": "string", "enum": []any{"get", "set", "delete", "list", "search"}, "description": "Memory operation"},
					"key":       map[string]any{"type": "string", "description": "Key to get/set/delete"},
					"value":     map[string]any{"type": "string", "description": "Value to store (for set operation)"},
					"query":     map[string]any{"type": "string", "description": "Search query (for search operation)"},
				},
				"required": []any{"operation"},
			}, tool.CatUtility, tool.CapRead, tool.CapWrite)
			t.Aliases = []string{"kv", "store", "remember"}
			t.DefaultTimeout = 5000
			return t
		}(),

		// ── Database (SQLite) ──
		func() tool.Tool {
			t := tool.New("database", "Execute SQL queries against a SQLite database file. Supports SELECT, INSERT, UPDATE, DELETE, CREATE TABLE.", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "Path to SQLite database file"},
					"sql":  map[string]any{"type": "string", "description": "SQL query to execute"},
				},
				"required": []any{"sql"},
			}, tool.CatDatabase, tool.CapDatabase, tool.CapRead)
			t.Aliases = []string{"sql", "sqlite", "query", "db"}
			t.DefaultTimeout = 30000
			t.Dangerous = true
			return t
		}(),

		// ── Image Analysis (stub) ──
		func() tool.Tool {
			t := tool.New("image_analysis", "Analyze an image file. This environment does not have image analysis capabilities.", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "Path to the image file"},
				},
				"required": []any{"path"},
			}, tool.CatUtility, tool.CapRead)
			t.Aliases = []string{"analyze_image", "image"}
			t.DefaultTimeout = 5000
			return t
		}(),

		// ── OCR (stub) ──
		func() tool.Tool {
			t := tool.New("ocr", "Extract text from an image using OCR. This environment does not have OCR capabilities.", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "Path to the image file"},
				},
				"required": []any{"path"},
			}, tool.CatUtility, tool.CapRead)
			t.Aliases = []string{"image_to_text", "extract_text"}
			t.DefaultTimeout = 5000
			return t
		}(),

		// ── Browser Automation (stub) ──
		func() tool.Tool {
			t := tool.New("browser", "Automate a web browser. This environment does not have browser automation capabilities.", map[string]any{
				"type": "object",
				"properties": map[string]any{
					"action": map[string]any{"type": "string", "description": "Browser action (not available)"},
					"url":    map[string]any{"type": "string", "description": "URL to navigate to"},
				},
				"required": []any{"action"},
			}, tool.CatBrowser, tool.CapBrowser)
			t.Aliases = []string{"browser_automation", "web_automation"}
			t.DefaultTimeout = 5000
			return t
		}(),
	})
}

// ── Executor Functions ─────────────────────────────────────
// Each returns an error if the tool failed, nil on success.
// The result struct is populated in-place.

func (s *Service) execRead(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	content, err := s.files.Read(path)
	if err != nil {
		result.Error = err.Error()
		return nil // error already in result
	}
	result.Status = tool.StatusSuccess
	result.Output = content
	return nil
}

func (s *Service) execWrite(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	content, _ := input["content"].(string)

	oldData, _ := os.ReadFile(path)

	if err := s.files.Write(path, content); err != nil {
		result.Error = err.Error()
		return nil
	}

	s.undo.Record(executor.UndoEntry{
		Type:    executor.OpWrite,
		Path:    path,
		OldData: oldData,
		NewData: []byte(content),
	})

	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Wrote %d bytes to %s", len(content), path)
	return nil
}

func (s *Service) execEdit(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	editsRaw, ok := input["edits"].([]any)
	if !ok {
		result.Error = "edits must be an array"
		return nil
	}
	var edits []executor.EditOp
	for _, er := range editsRaw {
		em, ok := er.(map[string]any)
		if !ok {
			continue
		}
		edits = append(edits, executor.EditOp{
			OldStr: toString(em["old_str"]),
			NewStr: toString(em["new_str"]),
		})
	}

	oldData, _ := os.ReadFile(path)

	editResult, err := s.files.Edit(path, edits)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	if !editResult.Success {
		result.Error = editResult.Message
		return nil
	}

	s.undo.Record(executor.UndoEntry{
		Type:    executor.OpEdit,
		Path:    path,
		OldData: oldData,
	})

	result.Status = tool.StatusSuccess
	result.Output = editResult.Message
	return nil
}

func (s *Service) execReplace(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	pattern, _ := input["pattern"].(string)
	newText, _ := input["new_text"].(string)
	isRegex, _ := input["is_regex"].(bool)

	res, err := s.files.Replace(path, []executor.ReplaceOp{{
		Pattern: pattern,
		NewText: newText,
		IsRegex: isRegex,
	}})
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	if !res.Success {
		result.Error = res.Message
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = res.Message
	return nil
}

func (s *Service) execRename(input map[string]any, result *tool.Result) error {
	source, _ := input["source"].(string)
	dest, _ := input["dest"].(string)
	if err := s.files.Rename(source, dest); err != nil {
		result.Error = err.Error()
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Renamed %s -> %s", source, dest)
	return nil
}

func (s *Service) execCopy(input map[string]any, result *tool.Result) error {
	source, _ := input["source"].(string)
	dest, _ := input["dest"].(string)
	if err := s.files.Copy(source, dest); err != nil {
		result.Error = err.Error()
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Copied %s -> %s", source, dest)
	return nil
}

func (s *Service) execPatch(input map[string]any, result *tool.Result) error {
	patchFile, _ := input["path"].(string)
	content, _ := input["content"].(string)
	baseDir, _ := input["base_dir"].(string)

	if patchFile != "" {
		if err := s.applyPatchFromFile(patchFile); err != nil {
			result.Error = err.Error()
			return nil
		}
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("Applied patch: %s", patchFile)
		return nil
	}

	if content != "" {
		if baseDir == "" {
			baseDir = "."
		}
		patches := executor.ParsePatchContent(content)
		var applied []string
		for _, p := range patches {
			fullPath := filepath.Join(baseDir, p.File)
			if err := s.applyPatchToFile(fullPath, p.Hunks); err != nil {
				result.Error = fmt.Sprintf("apply to %s: %s", p.File, err)
				return nil
			}
			applied = append(applied, p.File)
		}
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("Applied patch to %d files: %s", len(applied), strings.Join(applied, ", "))
		return nil
	}

	result.Error = "either path or content must be provided"
	return nil
}

func (s *Service) applyPatchFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read patch file: %w", err)
	}
	patches := executor.ParsePatchContent(string(data))
	for _, p := range patches {
		absPath, err := filepath.Abs(p.File)
		if err != nil {
			return fmt.Errorf("resolve %s: %w", p.File, err)
		}
		if err := s.applyPatchToFile(absPath, p.Hunks); err != nil {
			return fmt.Errorf("%s: %w", p.File, err)
		}
	}
	return nil
}

func (s *Service) applyPatchToFile(filePath string, hunks []executor.Hunk) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	content := strings.Split(string(data), "\n")
	if len(content) > 0 && content[len(content)-1] == "" {
		content = content[:len(content)-1]
	}

	for i := len(hunks) - 1; i >= 0; i-- {
		h := hunks[i]
		var newLines []string
		for _, line := range h.Lines {
			if len(line) == 0 {
				continue
			}
			switch line[0] {
			case '+':
				newLines = append(newLines, line[1:])
			case '-':
				continue
			default:
				newLines = append(newLines, line[1:])
			}
		}

		oldLines := make([]string, 0)
		for _, line := range h.Lines {
			if len(line) == 0 {
				continue
			}
			if line[0] != '+' {
				oldLines = append(oldLines, line[1:])
			}
		}

		start := h.OldStart - 1
		if start < 0 {
			start = 0
		}
		if start > len(content) {
			start = len(content)
		}

		end := start + len(oldLines)
		if end > len(content) {
			end = len(content)
		}

		var newContent []string
		newContent = append(newContent, content[:start]...)
		newContent = append(newContent, newLines...)
		if end < len(content) {
			newContent = append(newContent, content[end:]...)
		}
		content = newContent
	}

	return os.WriteFile(filePath, []byte(strings.Join(content, "\n")), 0o644)
}

func (s *Service) execUndo(result *tool.Result) error {
	entry, ok := s.undo.PopUndo()
	if !ok {
		result.Error = "nothing to undo"
		return nil
	}

	var err error
	switch entry.Type {
	case executor.OpWrite:
		if len(entry.OldData) == 0 {
			err = os.Remove(entry.Path)
		} else {
			err = os.WriteFile(entry.Path, entry.OldData, 0o644)
		}
	case executor.OpEdit, executor.OpReplace:
		err = os.WriteFile(entry.Path, entry.OldData, 0o644)
	case executor.OpDelete:
		dir := filepath.Dir(entry.Path)
		if mkErr := os.MkdirAll(dir, 0o755); mkErr != nil {
			result.Error = mkErr.Error()
			return nil
		}
		err = os.WriteFile(entry.Path, entry.OldData, 0o644)
	case executor.OpMove:
		err = os.Rename(entry.Path, entry.OldPath)
	case executor.OpCreate:
		err = os.Remove(entry.Path)
	case executor.OpCopy:
		err = nil
	}

	if err != nil {
		result.Error = err.Error()
		return nil
	}

	s.undo.PushRedo(entry)

	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Undone: %s %s", entry.Type, entry.Path)
	return nil
}

func (s *Service) execRedo(result *tool.Result) error {
	entry, ok := s.undo.PopRedo()
	if !ok {
		result.Error = "nothing to redo"
		return nil
	}

	var err error
	switch entry.Type {
	case executor.OpWrite:
		err = os.WriteFile(entry.Path, entry.NewData, 0o644)
	case executor.OpEdit, executor.OpReplace:
		err = os.WriteFile(entry.Path, entry.NewData, 0o644)
	case executor.OpDelete:
		err = os.WriteFile(entry.Path, entry.OldData, 0o644)
	case executor.OpMove:
		err = os.Rename(entry.Path, entry.OldPath)
	case executor.OpCreate:
		err = os.Remove(entry.Path)
	case executor.OpCopy:
		err = os.Remove(entry.Path)
	}

	if err != nil {
		result.Error = err.Error()
		return nil
	}

	s.undo.PushUndo(entry)

	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Redone: %s %s", entry.Type, entry.Path)
	return nil
}

func (s *Service) execReadHead(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	maxLines, _ := input["max_lines"].(int)
	if maxLines <= 0 {
		maxLines = 100
	}
	lines, err := s.files.ReadHead(path, maxLines)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(lines, "\n")
	return nil
}

func (s *Service) execReadTail(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	maxLines, _ := input["max_lines"].(int)
	if maxLines <= 0 {
		maxLines = 100
	}
	lines, err := s.files.ReadTail(path, maxLines)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(lines, "\n")
	return nil
}

func (s *Service) execReadRange(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	startLine, _ := input["start_line"].(int)
	endLine, _ := input["end_line"].(int)
	if startLine <= 0 {
		startLine = 1
	}
	if endLine < startLine {
		endLine = startLine + 100
	}
	lines, err := s.files.ReadLinesRange(path, startLine-1, endLine)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(lines, "\n")
	return nil
}

func (s *Service) execCreateFile(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	fileType, _ := input["type"].(string)
	if fileType == "" {
		fileType = "file"
	}
	switch fileType {
	case "dir":
		if err := os.MkdirAll(path, 0o755); err != nil {
			result.Error = err.Error()
			return nil
		}
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("Directory created: %s", path)
	default:
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			result.Error = err.Error()
			return nil
		}
		f, err := os.Create(path)
		if err != nil {
			result.Error = err.Error()
			return nil
		}
		f.Close()
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("File created: %s", path)
	}
	return nil
}

func (s *Service) execDeleteFile(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	if err := os.Remove(path); err != nil {
		result.Error = err.Error()
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Deleted: %s", path)
	return nil
}

func (s *Service) execListDir(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	if path == "" {
		path = "."
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	var lines []string
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			continue
		}
		size := ""
		if !info.IsDir() {
			size = fmt.Sprintf(" (%d bytes)", info.Size())
		}
		prefix := " "
		if info.IsDir() {
			prefix = "d"
		}
		lines = append(lines, fmt.Sprintf("%s %s%s", prefix, e.Name(), size))
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(lines, "\n")
	return nil
}

func (s *Service) execFileInfo(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	info, err := os.Stat(path)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	kind := "file"
	if info.IsDir() {
		kind = "directory"
	}
	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Name: %s\nType: %s\nSize: %d bytes\nMode: %s\nModified: %s",
		path, kind, info.Size(), info.Mode(), info.ModTime().Format(time.RFC3339))
	return nil
}

func (s *Service) execSearch(ctx context.Context, input map[string]any, result *tool.Result) error {
	pattern, _ := input["pattern"].(string)
	searchPath, _ := input["path"].(string)
	if searchPath == "" {
		searchPath = "."
	}
	searcher := executor.NewFileSearcher()
	results, err := searcher.Search(searchPath, pattern)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(results, "\n")
	return nil
}

func (s *Service) execGlob(input map[string]any, result *tool.Result) error {
	pattern, _ := input["pattern"].(string)
	searchPath, _ := input["path"].(string)
	if searchPath != "" {
		pattern = filepath.Join(searchPath, pattern)
	}
	matches, err := filepath.Glob(pattern)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	if len(matches) == 0 {
		result.Status = tool.StatusSuccess
		result.Output = "No matches."
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(matches, "\n")
	return nil
}

func (s *Service) execBash(ctx context.Context, input map[string]any, result *tool.Result) error {
	command, _ := input["command"].(string)
	shellResult, err := s.shell.Execute(ctx, command)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	if shellResult.ExitCode != 0 {
		result.Output = shellResult.Stdout
		result.Error = shellResult.Stderr
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = shellResult.Stdout
	return nil
}

func (s *Service) execGitStatus(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	repo, err := s.git.Open(path)
	if err != nil {
		result.Error = fmt.Sprintf("open repo: %v", err)
		return nil
	}
	status, err := s.git.Status(repo)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	if status.Clean {
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("On branch %s. Clean working tree.", status.Branch)
		return nil
	}
	var parts []string
	parts = append(parts, fmt.Sprintf("On branch %s (%s)", status.Branch, status.Hash))
	if len(status.Staged) > 0 {
		parts = append(parts, "Staged: "+strings.Join(status.Staged, ", "))
	}
	if len(status.Modified) > 0 {
		parts = append(parts, "Modified: "+strings.Join(status.Modified, ", "))
	}
	if len(status.Added) > 0 {
		parts = append(parts, "Added: "+strings.Join(status.Added, ", "))
	}
	if len(status.Deleted) > 0 {
		parts = append(parts, "Deleted: "+strings.Join(status.Deleted, ", "))
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(parts, "\n")
	return nil
}

func (s *Service) execGitLog(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	count, _ := input["count"].(int)
	if count <= 0 {
		count = 10
	}
	repo, err := s.git.Open(path)
	if err != nil {
		result.Error = fmt.Sprintf("open repo: %v", err)
		return nil
	}
	entries, err := s.git.Log(repo, count)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	var lines []string
	for _, e := range entries {
		msg := strings.SplitN(e.Message, "\n", 2)[0]
		lines = append(lines, fmt.Sprintf("%s %s - %s", e.Hash, e.When.Format("2006-01-02"), msg))
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(lines, "\n")
	return nil
}

func (s *Service) execGitDiff(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	repo, err := s.git.Open(path)
	if err != nil {
		result.Error = fmt.Sprintf("open repo: %v", err)
		return nil
	}
	diff, err := s.git.Diff(repo)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	if len(diff.Files) == 0 {
		result.Status = tool.StatusSuccess
		result.Output = "No changes."
		return nil
	}
	var lines []string
	for _, f := range diff.Files {
		lines = append(lines, fmt.Sprintf("%s (+%d/-%d)", f.Name, f.Added, f.Removed))
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(lines, "\n")
	return nil
}

func (s *Service) execGitAdd(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	filesRaw, _ := input["files"].([]any)
	repo, err := s.git.Open(path)
	if err != nil {
		result.Error = fmt.Sprintf("open repo: %v", err)
		return nil
	}
	if len(filesRaw) == 0 {
		if err := s.git.AddAll(repo); err != nil {
			result.Error = err.Error()
			return nil
		}
		result.Status = tool.StatusSuccess
		result.Output = "All files staged."
		return nil
	}
	var files []string
	for _, f := range filesRaw {
		files = append(files, fmt.Sprintf("%v", f))
	}
	if err := s.git.Add(repo, files); err != nil {
		result.Error = err.Error()
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Staged: %s", strings.Join(files, ", "))
	return nil
}

func (s *Service) execGitCommit(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	message, _ := input["message"].(string)
	repo, err := s.git.Open(path)
	if err != nil {
		result.Error = fmt.Sprintf("open repo: %v", err)
		return nil
	}
	hash, err := s.git.Commit(repo, message)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Committed as %s: %s", hash, message)
	return nil
}

func (s *Service) execGitBranches(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	repo, err := s.git.Open(path)
	if err != nil {
		result.Error = fmt.Sprintf("open repo: %v", err)
		return nil
	}
	branches, err := s.git.Branches(repo)
	if err != nil {
		result.Error = err.Error()
		return nil
	}
	currentBranch, err := s.git.GetBranch(repo)
	if err != nil {
		currentBranch = ""
	}
	var lines []string
	for _, b := range branches {
		if b == currentBranch {
			lines = append(lines, fmt.Sprintf("* %s (current)", b))
		} else {
			lines = append(lines, fmt.Sprintf("  %s", b))
		}
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(lines, "\n")
	return nil
}

func (s *Service) execGitCheckout(input map[string]any, result *tool.Result) error {
	path, _ := input["path"].(string)
	branch, _ := input["branch"].(string)
	create, _ := input["create"].(bool)
	repo, err := s.git.Open(path)
	if err != nil {
		result.Error = fmt.Sprintf("open repo: %v", err)
		return nil
	}
	if err := s.git.Checkout(repo, branch, create); err != nil {
		result.Error = err.Error()
		return nil
	}
	verb := "Switched to"
	if create {
		verb = "Created and switched to"
	}
	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("%s branch '%s'.", verb, branch)
	return nil
}

func (s *Service) execFetchURL(input map[string]any, result *tool.Result) error {
	rawURL, _ := input["url"].(string)
	method, _ := input["method"].(string)
	body, _ := input["body"].(string)

	if method == "" {
		method = "GET"
	}
	method = strings.ToUpper(method)

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		result.Error = fmt.Sprintf("invalid URL: %v", err)
		result.Status = tool.StatusFailed
		return nil
	}
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		result.Error = fmt.Sprintf("unsupported URL scheme: %s", parsedURL.Scheme)
		result.Status = tool.StatusFailed
		return nil
	}

	client := &http.Client{Timeout: 30 * time.Second}

	var req *http.Request
	if method == "POST" && body != "" {
		req, err = http.NewRequest(method, rawURL, strings.NewReader(body))
		if err != nil {
			result.Error = fmt.Sprintf("create request: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, rawURL, nil)
		if err != nil {
			result.Error = fmt.Sprintf("create request: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		result.Error = fmt.Sprintf("request failed: %v", err)
		result.Status = tool.StatusFailed
		return nil
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Error = fmt.Sprintf("read response: %v", err)
		result.Status = tool.StatusFailed
		return nil
	}

	headers := ""
	for k, v := range resp.Header {
		headers += fmt.Sprintf("%s: %s\n", k, strings.Join(v, ", "))
	}

	bodyStr := string(data)
	if strings.Contains(bodyStr, "<!DOCTYPE") || strings.Contains(bodyStr, "<html") {
		bodyStr = bodyStr[:min(len(bodyStr), 500)] + "\n... (HTML content truncated)"
	}

	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("HTTP %d %s\n\n%s\n\n%s", resp.StatusCode, resp.Status, headers, bodyStr)
	return nil
}

func (s *Service) execCalculate(input map[string]any, result *tool.Result) error {
	expr, _ := input["expression"].(string)
	if expr == "" {
		result.Error = "empty expression"
		result.Status = tool.StatusFailed
		return nil
	}

	node, err := parser.ParseExpr(expr)
	if err != nil {
		result.Error = fmt.Sprintf("parse error: %v", err)
		result.Status = tool.StatusFailed
		return nil
	}

	val, err := s.evalExpr(node)
	if err != nil {
		result.Error = err.Error()
		result.Status = tool.StatusFailed
		return nil
	}

	result.Status = tool.StatusSuccess
	if val == math.Trunc(val) && !math.IsInf(val, 0) && math.Abs(val) < 1e15 {
		result.Output = fmt.Sprintf("%d", int64(val))
	} else {
		result.Output = fmt.Sprintf("%v", val)
	}
	return nil
}

func (s *Service) evalExpr(node ast.Node) (float64, error) {
	switch n := node.(type) {
	case *ast.Ident:
		switch n.Name {
		case "pi":
			return math.Pi, nil
		case "e":
			return math.E, nil
		default:
			return 0, fmt.Errorf("undefined variable: %s", n.Name)
		}
	case *ast.BasicLit:
		if n.Kind == token.INT || n.Kind == token.FLOAT {
			return strconv.ParseFloat(n.Value, 64)
		}
		return 0, fmt.Errorf("unsupported literal: %s", n.Value)
	case *ast.BinaryExpr:
		x, err := s.evalExpr(n.X)
		if err != nil {
			return 0, err
		}
		y, err := s.evalExpr(n.Y)
		if err != nil {
			return 0, err
		}
		switch n.Op {
		case token.ADD:
			return x + y, nil
		case token.SUB:
			return x - y, nil
		case token.MUL:
			return x * y, nil
		case token.QUO:
			if y == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return x / y, nil
		default:
			return 0, fmt.Errorf("unsupported operator: %s", n.Op.String())
		}
	case *ast.UnaryExpr:
		x, err := s.evalExpr(n.X)
		if err != nil {
			return 0, err
		}
		switch n.Op {
		case token.SUB:
			return -x, nil
		case token.ADD:
			return x, nil
		default:
			return 0, fmt.Errorf("unsupported unary operator: %s", n.Op.String())
		}
	case *ast.ParenExpr:
		return s.evalExpr(n.X)
	case *ast.CallExpr:
		fn := ""
		if ident, ok := n.Fun.(*ast.Ident); ok {
			fn = ident.Name
		}
		args := make([]float64, len(n.Args))
		for i, arg := range n.Args {
			v, err := s.evalExpr(arg)
			if err != nil {
				return 0, err
			}
			args[i] = v
		}
		switch fn {
		case "sqrt":
			if len(args) != 1 {
				return 0, fmt.Errorf("sqrt requires 1 argument")
			}
			return math.Sqrt(args[0]), nil
		case "sin":
			if len(args) != 1 {
				return 0, fmt.Errorf("sin requires 1 argument")
			}
			return math.Sin(args[0]), nil
		case "cos":
			if len(args) != 1 {
				return 0, fmt.Errorf("cos requires 1 argument")
			}
			return math.Cos(args[0]), nil
		case "tan":
			if len(args) != 1 {
				return 0, fmt.Errorf("tan requires 1 argument")
			}
			return math.Tan(args[0]), nil
		case "pow":
			if len(args) != 2 {
				return 0, fmt.Errorf("pow requires 2 arguments")
			}
			return math.Pow(args[0], args[1]), nil
		case "abs":
			if len(args) != 1 {
				return 0, fmt.Errorf("abs requires 1 argument")
			}
			return math.Abs(args[0]), nil
		case "round":
			if len(args) != 1 {
				return 0, fmt.Errorf("round requires 1 argument")
			}
			return math.Round(args[0]), nil
		case "floor":
			if len(args) != 1 {
				return 0, fmt.Errorf("floor requires 1 argument")
			}
			return math.Floor(args[0]), nil
		case "ceil":
			if len(args) != 1 {
				return 0, fmt.Errorf("ceil requires 1 argument")
			}
			return math.Ceil(args[0]), nil
		case "log":
			if len(args) != 1 {
				return 0, fmt.Errorf("log requires 1 argument")
			}
			return math.Log10(args[0]), nil
		case "ln":
			if len(args) != 1 {
				return 0, fmt.Errorf("ln requires 1 argument")
			}
			return math.Log(args[0]), nil
		default:
			return 0, fmt.Errorf("unsupported function: %s", fn)
		}
	case *ast.SelectorExpr:
		x, err := s.evalExpr(n.X)
		if err != nil {
			return 0, err
		}
		switch n.Sel.Name {
		case "pi":
			return x * math.Pi, nil
		case "e":
			return x * math.E, nil
		default:
			return 0, fmt.Errorf("unsupported selector: %s", n.Sel.Name)
		}
	default:
		return 0, fmt.Errorf("unsupported expression type: %T", node)
	}
}

func (s *Service) execJSON(input map[string]any, result *tool.Result) error {
	operation, _ := input["operation"].(string)
	data, _ := input["data"].(string)
	path, _ := input["path"].(string)

	switch operation {
	case "validate":
		var v any
		if err := json.Unmarshal([]byte(data), &v); err != nil {
			result.Error = fmt.Sprintf("invalid JSON: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		result.Status = tool.StatusSuccess
		result.Output = "valid JSON"

	case "format":
		var v any
		if err := json.Unmarshal([]byte(data), &v); err != nil {
			result.Error = fmt.Sprintf("invalid JSON: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		formatted, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			result.Error = fmt.Sprintf("format error: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		result.Status = tool.StatusSuccess
		result.Output = string(formatted)

	case "query":
		var v any
		if err := json.Unmarshal([]byte(data), &v); err != nil {
			result.Error = fmt.Sprintf("invalid JSON: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		val, err := s.queryJSON(v, path)
		if err != nil {
			result.Error = err.Error()
			result.Status = tool.StatusFailed
			return nil
		}
		out, _ := json.MarshalIndent(val, "", "  ")
		result.Status = tool.StatusSuccess
		result.Output = string(out)

	default:
		result.Error = fmt.Sprintf("unknown operation: %s", operation)
		result.Status = tool.StatusFailed
	}
	return nil
}

func (s *Service) queryJSON(v any, path string) (any, error) {
	if path == "" {
		return v, nil
	}
	parts := strings.Split(path, ".")
	current := v
	for _, part := range parts {
		if m, ok := current.(map[string]any); ok {
			val, exists := m[part]
			if !exists {
				return nil, fmt.Errorf("key %q not found", part)
			}
			current = val
		} else if arr, ok := current.([]any); ok {
			idx, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("expected array index, got %q", part)
			}
			if idx < 0 || idx >= len(arr) {
				return nil, fmt.Errorf("index %d out of bounds (len %d)", idx, len(arr))
			}
			current = arr[idx]
		} else if m, ok := current.(map[string]string); ok {
			val, exists := m[part]
			if !exists {
				return nil, fmt.Errorf("key %q not found", part)
			}
			current = val
		} else {
			return nil, fmt.Errorf("cannot navigate into %T", current)
		}
	}
	return current, nil
}

func (s *Service) execXML(input map[string]any, result *tool.Result) error {
	operation, _ := input["operation"].(string)
	data, _ := input["data"].(string)

	switch operation {
	case "validate":
		var v any
		if err := xml.Unmarshal([]byte(data), &v); err != nil {
			result.Error = fmt.Sprintf("invalid XML: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		result.Status = tool.StatusSuccess
		result.Output = "valid XML"

	case "format":
		var v any
		if err := xml.Unmarshal([]byte(data), &v); err != nil {
			result.Error = fmt.Sprintf("invalid XML: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		out, err := xml.MarshalIndent(v, "", "  ")
		if err != nil {
			result.Error = fmt.Sprintf("format error: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		result.Status = tool.StatusSuccess
		result.Output = string(out)

	default:
		result.Error = fmt.Sprintf("unknown operation: %s", operation)
		result.Status = tool.StatusFailed
	}
	return nil
}

func (s *Service) execMemory(input map[string]any, result *tool.Result) error {
	operation, _ := input["operation"].(string)
	key, _ := input["key"].(string)
	value, _ := input["value"].(string)
	query, _ := input["query"].(string)

	s.mu.Lock()
	defer s.mu.Unlock()

	switch operation {
	case "set":
		if key == "" {
			result.Error = "key is required for set"
			result.Status = tool.StatusFailed
			return nil
		}
		s.memoryData[key] = value
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("stored %d bytes under key %q", len(value), key)

	case "get":
		if key == "" {
			result.Error = "key is required for get"
			result.Status = tool.StatusFailed
			return nil
		}
		val, exists := s.memoryData[key]
		if !exists {
			result.Error = fmt.Sprintf("key %q not found", key)
			result.Status = tool.StatusFailed
			return nil
		}
		result.Status = tool.StatusSuccess
		result.Output = val

	case "delete":
		if key == "" {
			result.Error = "key is required for delete"
			result.Status = tool.StatusFailed
			return nil
		}
		delete(s.memoryData, key)
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("deleted key %q", key)

	case "list":
		if len(s.memoryData) == 0 {
			result.Status = tool.StatusSuccess
			result.Output = "(empty)"
			return nil
		}
		var lines []string
		for k := range s.memoryData {
			val := s.memoryData[k]
			display := val
			if len(display) > 60 {
				display = display[:60] + "..."
			}
			lines = append(lines, fmt.Sprintf("%s = %s", k, display))
		}
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("%d entries:\n%s", len(s.memoryData), strings.Join(lines, "\n"))

	case "search":
		if query == "" {
			result.Error = "query is required for search"
			result.Status = tool.StatusFailed
			return nil
		}
		var lines []string
		q := strings.ToLower(query)
		for k, v := range s.memoryData {
			if strings.Contains(strings.ToLower(k), q) || strings.Contains(strings.ToLower(v), q) {
				display := v
				if len(display) > 60 {
					display = display[:60] + "..."
				}
				lines = append(lines, fmt.Sprintf("%s = %s", k, display))
			}
		}
		if len(lines) == 0 {
			result.Status = tool.StatusSuccess
			result.Output = fmt.Sprintf("no entries matching %q", query)
			return nil
		}
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("%d matches:\n%s", len(lines), strings.Join(lines, "\n"))

	default:
		result.Error = fmt.Sprintf("unknown operation: %s", operation)
		result.Status = tool.StatusFailed
	}
	return nil
}

func (s *Service) execDatabase(input map[string]any, result *tool.Result) error {
	sqlQuery, _ := input["sql"].(string)
	dbPath, _ := input["path"].(string)

	if sqlQuery == "" {
		result.Error = "SQL query is required"
		result.Status = tool.StatusFailed
		return nil
	}

	if dbPath == "" {
		dbPath = ":memory:"
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		result.Error = fmt.Sprintf("open database: %v", err)
		result.Status = tool.StatusFailed
		return nil
	}
	defer db.Close()

	upper := strings.ToUpper(strings.TrimSpace(sqlQuery))
	isQuery := strings.HasPrefix(upper, "SELECT") || strings.HasPrefix(upper, "PRAGMA") || strings.HasPrefix(upper, "EXPLAIN")

	if isQuery {
		rows, err := db.Query(sqlQuery)
		if err != nil {
			result.Error = fmt.Sprintf("query error: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			result.Error = fmt.Sprintf("get columns: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}

		var output strings.Builder
		output.WriteString(strings.Join(columns, " | "))
		output.WriteString("\n")
		for i := range columns {
			if i > 0 {
				output.WriteString("---")
			}
			output.WriteString(strings.Repeat("-", len(columns[i])))
		}
		output.WriteString("\n")

		rowCount := 0
		for rows.Next() {
			vals := make([]any, len(columns))
			valPtrs := make([]any, len(columns))
			for i := range vals {
				valPtrs[i] = &vals[i]
			}
			if err := rows.Scan(valPtrs...); err != nil {
				result.Error = fmt.Sprintf("scan row: %v", err)
				result.Status = tool.StatusFailed
				return nil
			}
			for i, v := range vals {
				if i > 0 {
					output.WriteString(" | ")
				}
				if v == nil {
					output.WriteString("NULL")
				} else {
					output.WriteString(fmt.Sprintf("%v", v))
				}
			}
			output.WriteString("\n")
			rowCount++
			if rowCount >= 100 {
				output.WriteString(fmt.Sprintf("... (%d+ rows truncated)", rowCount))
				break
			}
		}

		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("%d row(s) returned:\n%s", rowCount, output.String())
	} else {
		res, err := db.Exec(sqlQuery)
		if err != nil {
			result.Error = fmt.Sprintf("exec error: %v", err)
			result.Status = tool.StatusFailed
			return nil
		}
		affected, _ := res.RowsAffected()
		lastID, _ := res.LastInsertId()
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("OK (rows affected: %d, last insert ID: %d)", affected, lastID)
	}
	return nil
}

func (s *Service) execImageAnalysis(input map[string]any, result *tool.Result) error {
	result.Status = tool.StatusSuccess
	result.Output = "Image analysis is not available in this environment. The terminal-based AI coding CLI does not support image processing. Consider using an external image analysis tool or API."
	return nil
}

func (s *Service) execOCR(input map[string]any, result *tool.Result) error {
	result.Status = tool.StatusSuccess
	result.Output = "OCR (optical character recognition) is not available in this environment. The terminal-based AI coding CLI does not support image processing. Consider using an external OCR tool or API."
	return nil
}

func (s *Service) execBrowser(input map[string]any, result *tool.Result) error {
	result.Status = tool.StatusSuccess
	result.Output = "Browser automation is not available in this environment. The terminal-based AI coding CLI does not have browser automation capabilities. Consider using a headless browser tool like Playwright or Puppeteer externally."
	return nil
}

func toString(v any) string {
	if v == nil {
		return ""
	}
	s, ok := v.(string)
	if !ok {
		return fmt.Sprintf("%v", v)
	}
	return s
}
