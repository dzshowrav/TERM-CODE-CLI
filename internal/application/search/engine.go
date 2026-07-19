package search

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type Match struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

type Options struct {
	MaxResults    int
	CaseSensitive bool
	Regex         bool
	IncludeExt    string
	ExcludePath   string
}

type Engine struct {
	workDir string
}

func New(workDir string) *Engine {
	return &Engine{workDir: workDir}
}

func (e *Engine) SearchText(ctx context.Context, query string, opts Options) ([]Match, error) {
	if opts.MaxResults <= 0 {
		opts.MaxResults = 50
	}

	args := []string{
		"--line-number",
		"--with-filename",
		"--color", "never",
	}

	if !opts.CaseSensitive {
		args = append(args, "-i")
	}
	if opts.Regex {
		args = append(args, "-E")
	} else {
		args = append(args, "-F")
	}
	if opts.IncludeExt != "" {
		args = append(args, "-g", "*."+opts.IncludeExt)
	}
	if opts.ExcludePath != "" {
		args = append(args, "--glob", "!"+opts.ExcludePath)
	}

	args = append(args, query, e.workDir)

	cmd := exec.CommandContext(ctx, "rg", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				return nil, nil
			}
		}
		return nil, fmt.Errorf("rg search: %w\n%s", err, stderr.String())
	}

	return parseMatches(stdout.String(), opts.MaxResults), nil
}

func (e *Engine) SearchFiles(ctx context.Context, pattern string) ([]string, error) {
	args := []string{"--color", "never", "-g", pattern, e.workDir}
	cmd := exec.CommandContext(ctx, "fd", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				return nil, nil
			}
		}
		return nil, fmt.Errorf("fd search: %w\n%s", err, stderr.String())
	}

	lines := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return nil, nil
	}
	return lines, nil
}

func (e *Engine) SearchCode(ctx context.Context, query string, opts Options) ([]Match, error) {
	return e.SearchText(ctx, query, opts)
}

func parseMatches(output string, max int) []Match {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		return nil
	}

	var matches []Match
	for _, line := range lines {
		if len(matches) >= max {
			break
		}

		parts := strings.SplitN(line, ":", 3)
		if len(parts) < 3 {
			continue
		}

		file := parts[0]
		lineNum := 0
		fmt.Sscanf(parts[1], "%d", &lineNum)
		content := parts[2]

		matches = append(matches, Match{
			File:    file,
			Line:    lineNum,
			Content: strings.TrimSpace(content),
			Type:    "text",
		})
	}
	return matches
}
