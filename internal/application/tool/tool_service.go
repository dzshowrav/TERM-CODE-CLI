package tool

import (
	"context"
	"fmt"
	"strings"
	"time"

	"termcode/internal/domain/tool"
	"termcode/internal/infrastructure/executor"
	git "termcode/internal/infrastructure/git"
)

type Service struct {
	shell *executor.ShellExecutor
	files *executor.FileExecutor
	git   *git.Service
}

func NewService() *Service {
	return &Service{
		shell: executor.NewShellExecutor(),
		files: executor.NewFileExecutor(),
		git:   git.NewService(),
	}
}

func (s *Service) AvailableTools() []tool.Tool {
	return []tool.Tool{
		{
			Name:        "read",
			Description: "Read the contents of a file",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "File path to read"},
				},
				"required": []string{"path"},
			},
		},
		{
			Name:        "write",
			Description: "Write content to a file (creates directories if needed)",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":    map[string]any{"type": "string", "description": "File path"},
					"content": map[string]any{"type": "string", "description": "File content"},
				},
				"required": []string{"path", "content"},
			},
		},
		{
			Name:        "edit",
			Description: "Edit a file by replacing exact text matches",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":  map[string]any{"type": "string", "description": "File path"},
					"edits": map[string]any{"type": "array", "items": map[string]any{"type": "object", "properties": map[string]any{"old_str": map[string]any{"type": "string"}, "new_str": map[string]any{"type": "string"}}}},
				},
				"required": []string{"path", "edits"},
			},
		},
		{
			Name:        "bash",
			Description: "Execute a shell command",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"command": map[string]any{"type": "string", "description": "Shell command to execute"},
				},
				"required": []string{"command"},
			},
		},
		{
			Name:        "search",
			Description: "Search for files by content using ripgrep",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"pattern": map[string]any{"type": "string", "description": "Search pattern"},
					"path":    map[string]any{"type": "string", "description": "Directory to search"},
				},
				"required": []string{"pattern"},
			},
		},
		{
			Name:        "git_status",
			Description: "Show the working tree status of a git repository",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "Repository path"},
				},
				"required": []string{"path"},
			},
		},
		{
			Name:        "git_log",
			Description: "Show commit logs",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":  map[string]any{"type": "string", "description": "Repository path"},
					"count": map[string]any{"type": "integer", "description": "Number of commits to show"},
				},
				"required": []string{"path"},
			},
		},
		{
			Name:        "git_diff",
			Description: "Show changes in the working tree",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path": map[string]any{"type": "string", "description": "Repository path"},
				},
				"required": []string{"path"},
			},
		},
		{
			Name:        "git_add",
			Description: "Stage files for commit",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":  map[string]any{"type": "string", "description": "Repository path"},
					"files": map[string]any{"type": "array", "items": map[string]any{"type": "string"}, "description": "Files to stage (omit for all)"},
				},
				"required": []string{"path"},
			},
		},
		{
			Name:        "git_commit",
			Description: "Commit staged changes",
			InputSchema: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"path":    map[string]any{"type": "string", "description": "Repository path"},
					"message": map[string]any{"type": "string", "description": "Commit message"},
				},
				"required": []string{"path", "message"},
			},
		},
	}
}

func (s *Service) Execute(ctx context.Context, toolName string, input map[string]any) *tool.Result {
	start := time.Now()

	result := &tool.Result{
		Tool:   toolName,
		Input:  fmt.Sprintf("%v", input),
		Status: tool.StatusRunning,
		Time:   start,
	}

	switch toolName {
	case "read":
		path, _ := input["path"].(string)
		content, err := s.files.Read(path)
		if err != nil {
			result.Status = tool.StatusFailed
			result.Error = err.Error()
		} else {
			result.Status = tool.StatusSuccess
			result.Output = content
		}

	case "write":
		path, _ := input["path"].(string)
		content, _ := input["content"].(string)
		if err := s.files.Write(path, content); err != nil {
			result.Status = tool.StatusFailed
			result.Error = err.Error()
		} else {
			result.Status = tool.StatusSuccess
			result.Output = fmt.Sprintf("Wrote %d bytes to %s", len(content), path)
		}

	case "edit":
		result = s.handleEdit(input, result)

	case "bash":
		command, _ := input["command"].(string)
		shellResult, err := s.shell.Execute(ctx, command)
		if err != nil {
			result.Status = tool.StatusFailed
			result.Error = err.Error()
		} else if shellResult.ExitCode != 0 {
			result.Status = tool.StatusFailed
			result.Output = shellResult.Stdout
			result.Error = shellResult.Stderr
		} else {
			result.Status = tool.StatusSuccess
			result.Output = shellResult.Stdout
		}

	case "search":
		pattern, _ := input["pattern"].(string)
		searchPath, _ := input["path"].(string)
		if searchPath == "" {
			searchPath = "."
		}
		searcher := executor.NewFileSearcher()
		results, err := searcher.Search(searchPath, pattern)
		if err != nil {
			result.Status = tool.StatusFailed
			result.Error = err.Error()
		} else {
			result.Status = tool.StatusSuccess
			result.Output = strings.Join(results, "\n")
		}

	case "git_status":
		path, _ := input["path"].(string)
		result = s.handleGitStatus(path, result)

	case "git_log":
		path, _ := input["path"].(string)
		count, _ := input["count"].(int)
		if count <= 0 {
			count = 10
		}
		result = s.handleGitLog(path, count, result)

	case "git_diff":
		path, _ := input["path"].(string)
		result = s.handleGitDiff(path, result)

	case "git_add":
		path, _ := input["path"].(string)
		filesRaw, _ := input["files"].([]any)
		var files []string
		for _, f := range filesRaw {
			files = append(files, fmt.Sprintf("%v", f))
		}
		result = s.handleGitAdd(path, files, result)

	case "git_commit":
		path, _ := input["path"].(string)
		message, _ := input["message"].(string)
		result = s.handleGitCommit(path, message, result)

	default:
		result.Status = tool.StatusFailed
		result.Error = fmt.Sprintf("unknown tool: %s", toolName)
	}

	result.Duration = time.Since(start).Milliseconds()
	return result
}

func (s *Service) handleEdit(input map[string]any, result *tool.Result) *tool.Result {
	path, _ := input["path"].(string)
	editsRaw, ok := input["edits"].([]any)
	if !ok {
		result.Status = tool.StatusFailed
		result.Error = "edits must be an array"
		return result
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
	editResult, err := s.files.Edit(path, edits)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = err.Error()
	} else if !editResult.Success {
		result.Status = tool.StatusFailed
		result.Error = editResult.Message
	} else {
		result.Status = tool.StatusSuccess
		result.Output = editResult.Message
	}
	return result
}

func (s *Service) handleGitStatus(path string, result *tool.Result) *tool.Result {
	repo, err := s.git.Open(path)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = fmt.Sprintf("open repo: %v", err)
		return result
	}
	status, err := s.git.Status(repo)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = err.Error()
		return result
	}
	if status.Clean {
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("On branch %s. Clean working tree.", status.Branch)
		return result
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
	return result
}

func (s *Service) handleGitLog(path string, count int, result *tool.Result) *tool.Result {
	repo, err := s.git.Open(path)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = fmt.Sprintf("open repo: %v", err)
		return result
	}
	entries, err := s.git.Log(repo, count)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = err.Error()
		return result
	}
	var lines []string
	for _, e := range entries {
		msg := strings.SplitN(e.Message, "\n", 2)[0]
		lines = append(lines, fmt.Sprintf("%s %s - %s", e.Hash, e.When.Format("2006-01-02"), msg))
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(lines, "\n")
	return result
}

func (s *Service) handleGitDiff(path string, result *tool.Result) *tool.Result {
	repo, err := s.git.Open(path)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = fmt.Sprintf("open repo: %v", err)
		return result
	}
	diff, err := s.git.Diff(repo)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = err.Error()
		return result
	}
	if len(diff.Files) == 0 {
		result.Status = tool.StatusSuccess
		result.Output = "No changes."
		return result
	}
	var lines []string
	for _, f := range diff.Files {
		lines = append(lines, fmt.Sprintf("%s (+%d/-%d)", f.Name, f.Added, f.Removed))
	}
	result.Status = tool.StatusSuccess
	result.Output = strings.Join(lines, "\n")
	return result
}

func (s *Service) handleGitAdd(path string, files []string, result *tool.Result) *tool.Result {
	repo, err := s.git.Open(path)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = fmt.Sprintf("open repo: %v", err)
		return result
	}
	if len(files) == 0 {
		if err := s.git.AddAll(repo); err != nil {
			result.Status = tool.StatusFailed
			result.Error = err.Error()
			return result
		}
		result.Status = tool.StatusSuccess
		result.Output = "All files staged."
	} else {
		if err := s.git.Add(repo, files); err != nil {
			result.Status = tool.StatusFailed
			result.Error = err.Error()
			return result
		}
		result.Status = tool.StatusSuccess
		result.Output = fmt.Sprintf("Staged: %s", strings.Join(files, ", "))
	}
	return result
}

func (s *Service) handleGitCommit(path, message string, result *tool.Result) *tool.Result {
	repo, err := s.git.Open(path)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = fmt.Sprintf("open repo: %v", err)
		return result
	}
	hash, err := s.git.Commit(repo, message)
	if err != nil {
		result.Status = tool.StatusFailed
		result.Error = err.Error()
		return result
	}
	result.Status = tool.StatusSuccess
	result.Output = fmt.Sprintf("Committed as %s: %s", hash, message)
	return result
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
