package executor

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type ShellResult struct {
	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exit_code"`
	Duration int64  `json:"duration_ms"`
}

type ShellExecutor struct {
	allowedCommands []string
	deniedCommands  []string
	timeout         time.Duration
}

var dangerousCommands = []string{
	"rm -rf /", "mkfs", "dd if=", "format", "shutdown", "reboot",
	":(){ :|:& };:", "> /dev/sda", "wget", "curl -s",
}

func NewShellExecutor() *ShellExecutor {
	return &ShellExecutor{
		timeout: 30 * time.Second,
	}
}

func (e *ShellExecutor) Execute(ctx context.Context, command string) (*ShellResult, error) {
	if reason, ok := e.isDangerous(command); ok {
		return &ShellResult{
			Stderr:   fmt.Sprintf("Command blocked: %s", reason),
			ExitCode: -1,
		}, nil
	}

	ctx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	start := time.Now()

	var stdout, stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, "bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	duration := time.Since(start).Milliseconds()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return &ShellResult{
				Stderr:   err.Error(),
				ExitCode: -1,
				Duration: duration,
			}, nil
		}
	}

	return &ShellResult{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: exitCode,
		Duration: duration,
	}, nil
}

func (e *ShellExecutor) ExecuteInDir(ctx context.Context, command, dir string) (*ShellResult, error) {
	if reason, ok := e.isDangerous(command); ok {
		return &ShellResult{
			Stderr:   fmt.Sprintf("Command blocked: %s", reason),
			ExitCode: -1,
		}, nil
	}

	ctx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	start := time.Now()

	var stdout, stderr bytes.Buffer
	cmd := exec.CommandContext(ctx, "bash", "-c", command)
	cmd.Dir = dir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	duration := time.Since(start).Milliseconds()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return &ShellResult{
				Stderr:   err.Error(),
				ExitCode: -1,
				Duration: duration,
			}, nil
		}
	}

	return &ShellResult{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		ExitCode: exitCode,
		Duration: duration,
	}, nil
}

func (e *ShellExecutor) isDangerous(command string) (string, bool) {
	lower := strings.ToLower(command)
	for _, dangerous := range dangerousCommands {
		if strings.Contains(lower, dangerous) {
			return fmt.Sprintf("command contains dangerous pattern: %s", dangerous), true
		}
	}
	return "", false
}
