# Shell Execution

## Overview

Shell execution lets the agent run commands in the user's terminal environment. This is one of the most powerful and most dangerous tools — it can start servers, install packages, run tests, and more.

## Command Execution

### Execution Model
- Commands run in a persistent shell session
- Each command is executed in the workspace directory
- The shell environment inherits the user's PATH, env vars, and shell config
- Commands have a configurable timeout (default: 30 seconds, adjustable per-command)
- Output is captured (stdout + stderr combined or separated)

### Lifecycle
1. **Queued** — command is submitted, waiting for execution slot
2. **Executing** — command is running
3. **Streaming** — output is being produced (real-time display)
4. **Completed** — command exited with status code 0
5. **Failed** — command exited with non-zero status
6. **Timed Out** — command exceeded timeout
7. **Aborted** — user or system cancelled the command
8. **Killed** — command was forcibly terminated

### Output Capture
- Stdout and stderr are captured separately but displayed merged
- Output is streamed in real-time (every line as it's produced)
- After completion, the full output is available for the model
- Long output is truncated (head + tail) for return to model
- Full output is stored in the session

### Timeout Behavior
When a command times out:
1. The command receives a graceful termination signal
2. If it doesn't respond within a grace period, it's forcefully killed
3. Partial output up to the timeout point is captured
4. The model is informed of the timeout and partial results

### Interactive Commands
Commands that require input (read from stdin) are detected:
- The agent prompts the user to provide input
- Input can be provided as a string or file
- Commands expecting interactive input time out if no input is provided

## Long-Running Commands

Commands expected to run for a long time (servers, watchers, dev servers):
- Warn the user about expected duration
- Can be started in "background" mode (non-blocking)
- Background commands are tracked — their output continues streaming
- The agent can check on background processes later
- Background processes can be stopped by the user

## Inline Bash (!)

The user can execute shell commands inline during input:
- A command starting with `!` is treated as a shell command, not AI input
- The command is executed directly in the shell
- Output is displayed in the conversation
- The result is NOT sent to the AI model (unless configured otherwise)
- Inline execution is instantaneous — no AI round-trip
- Supports pipes, redirects, and compound commands

## Security & Permissions

Shell execution follows the permission system:
- All shell commands require approval by default (unless auto-approved by policy)
- Commands matching certain patterns (install, delete, format, destroy) always require approval
- The command string is visible to the user before approval
- The actual command is displayed during execution

## Process Management

The system tracks all child processes:
- Process ID, command, start time, duration
- Current status (running, completed, failed, timed out)
- Process tree (parent-child relationships for spawned processes)
- Resource usage (CPU, memory) if available
- User can list and kill processes

## Environment Management
- Environment variables can be set per-session or per-command
- Sensitive variables (tokens, keys) are marked and not displayed
- The working directory can be changed for specific commands
- Shell type (bash, zsh, sh, PowerShell) is configurable

## Key Design Decisions

- Commands run in the user's shell, not a sandbox — the user's environment is preserved
- All commands are visible to the user — no hidden execution
- Long-running commands are non-blocking by explicit choice, not by default
- Interactive commands are detected and handled explicitly
- Inline execution (!) bypasses the AI entirely — it's a UX feature, not a tool
- Command timeout prevents runaway processes
- Background process management gives the user control over long-running tasks
