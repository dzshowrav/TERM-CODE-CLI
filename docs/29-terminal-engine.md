# 29-terminal-engine.md

# Terminal Engine Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Terminal Engine?

A **Terminal Engine** is the subsystem responsible for executing shell commands, managing terminal sessions, capturing output, streaming results, handling processes, and providing a secure interface between the AI Coding Agent and the operating system.

It enables the AI Agent to interact with the system just like a developer using a terminal.

Instead of allowing every module to execute shell commands directly, all terminal operations pass through the Terminal Engine.

---

# Why Terminal Engine?

Without Terminal Engine

```
Agent

↓

Execute Shell

↓

Unknown State
```

Problems

- Unsafe execution
- No process tracking
- No output streaming
- Difficult cancellation
- Poor security

---

With Terminal Engine

```
Application

↓

Terminal Engine

↓

Process Manager

↓

Operating System
```

---

# Goals

A production Terminal Engine should provide

- Secure command execution
- Interactive terminal sessions
- Process management
- Output streaming
- Command history
- Cancellation support
- Environment isolation
- Working directory management
- Timeout control
- Cross-platform compatibility

---

# High-Level Architecture

```
             User Request

                   │

                   ▼

           Terminal Engine

                   │

      ┌────────────┼────────────┐

      ▼            ▼            ▼

 Command      Process      Session

      ▼            ▼            ▼

 Stream      History      Security

      └────────────┼────────────┘

                   ▼

           Operating System
```

---

# Folder Structure

```
src/

terminal/

    TerminalEngine.ts

    TerminalSession.ts

    CommandExecutor.ts

    ProcessManager.ts

    StreamHandler.ts

    WorkingDirectory.ts

    EnvironmentManager.ts

    CommandHistory.ts

    TerminalSecurity.ts

    TerminalEvents.ts

    TerminalMetrics.ts

    TerminalValidator.ts
```

---

# Core Components

## Terminal Engine

Central controller.

Responsibilities

- Execute commands
- Manage sessions
- Stream output
- Track processes

---

## Terminal Session

Represents an interactive shell session.

Stores

```
Session ID

Current Directory

Environment

Running Processes
```

---

## Command Executor

Executes

```
Shell Commands

Scripts

Programs
```

through the operating system.

---

## Process Manager

Tracks

```
PID

Status

CPU

Memory

Exit Code
```

for every process.

---

## Stream Handler

Streams

```
stdout

stderr

stdin
```

between the process and the AI Agent.

---

## Working Directory Manager

Maintains the active project directory for each terminal session.

---

## Environment Manager

Loads and manages

```
Environment Variables

PATH

Shell Configuration
```

---

## Command History

Stores

```
Executed Commands

Timestamps

Exit Codes
```

for auditing and convenience.

---

## Terminal Security

Enforces

- Permission checks
- Command validation
- Sandbox rules
- Restricted operations

---

## Terminal Validator

Checks

- Command syntax
- Allowed operations
- Working directory
- Execution policy

---

# Terminal Lifecycle

```
Create Session

↓

Validate Command

↓

Execute

↓

Stream Output

↓

Finish

↓

Store History
```

---

# Terminal Session Object

Contains

```
Session ID

Workspace

Working Directory

Environment

Running Processes

Metadata
```

---

# Process States

```
Created

↓

Running

↓

Paused

↓

Completed

↓

Failed

↓

Killed
```

---

# Command Flow

```
Command

↓

Validation

↓

Execution

↓

Streaming

↓

Exit Code

↓

History
```

---

# Interactive Session

```
User

↓

Shell

↓

stdin

↓

stdout

↓

stderr
```

Supports long-running commands.

---

# Output Streaming

```
Command

↓

stdout

↓

Stream Handler

↓

Conversation Manager

↓

User
```

Real-time updates improve responsiveness.

---

# Command Cancellation

```
User

↓

Cancel

↓

Process Manager

↓

Terminate Process
```

---

# Timeout Control

```
Command

↓

Timer

↓

Timeout

↓

Kill Process
```

Prevents runaway processes.

---

# Working Directory Flow

```
Workspace

↓

Working Directory

↓

Execute Command
```

Each session may have its own directory.

---

# Environment Variables

Examples

```
PATH

HOME

SHELL

NODE_ENV

CUSTOM_VARIABLES
```

Loaded per session.

---

# Event Bus Integration

Common events

```
terminal:start

terminal:command

terminal:stream

terminal:exit

terminal:error

terminal:kill
```

---

# Workflow Engine Integration

Workflow steps may invoke terminal commands.

---

# Task Planner Integration

Planner decides

- Which commands to execute
- Execution order
- Dependencies

---

# Permission Engine Integration

Every command is validated before execution.

---

# Conversation Manager Integration

Streams terminal output to the user in real time.

---

# Logging Integration

Logs

```
Commands

Exit Codes

Errors

Execution Time
```

---

# State Manager Integration

Stores

```
Session State

Current Directory

Running Processes
```

---

# Plugin Integration

Plugins may add

- Shell integrations
- Command aliases
- Custom executors
- Environment providers

---

# Skills Integration

Skills may define reusable command workflows such as

- Project initialization
- Dependency installation
- Build pipelines
- Deployment scripts

---

# Cache Strategy

Cache

```
Environment

Working Directories

Command Metadata
```

Avoid caching command output unless explicitly required.

---

# Error Handling

```
Command Failed

↓

Capture stderr

↓

Log Error

↓

Return Exit Code

↓

Continue
```

Never lose diagnostic information.

---

# Security

Always

- Validate commands
- Restrict dangerous operations
- Isolate environments
- Respect permissions
- Audit execution

Never

- Execute unvalidated commands
- Allow arbitrary privilege escalation
- Expose secrets through output
- Ignore exit codes

---

# Performance Optimizations

Use

- Persistent sessions
- Streaming output
- Process reuse where appropriate
- Background execution
- Incremental logging

Avoid

- Creating a new shell for every command
- Blocking output streams
- Running unnecessary commands

---

# Best Practices

Always

- Validate commands
- Track processes
- Stream output
- Record history
- Handle cancellations
- Respect workspace boundaries

Never

- Ignore command failures
- Mix unrelated sessions
- Leave orphaned processes
- Skip security validation

---

# Common Mistakes

Bad

```
Agent

↓

Execute Shell

↓

Done
```

No tracking or validation.

---

Good

```
Terminal Engine

↓

Validate

↓

Execute

↓

Stream

↓

History

↓

Complete
```

Reliable and secure.

---

# Testing Checklist

- Command execution
- Interactive sessions
- Output streaming
- Process management
- Cancellation
- Timeout
- Environment handling
- Security validation
- History
- Error recovery

---

# Advantages

- Secure shell access
- Real-time output
- Process management
- Better debugging
- Cross-platform abstraction
- Improved automation

---

# Disadvantages

- Process management complexity
- Platform differences
- Security risks if misconfigured
- Resource consumption

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Cursor
- Continue.dev
- GitHub Copilot Workspace
- Enterprise AI Platforms

---

# Complete Terminal Flow

```
User Request

↓

Task Planner

↓

Workflow Engine

↓

Permission Engine

↓

Terminal Engine

↓

Command Validator

↓

Command Executor

↓

Process Manager

↓

Stream Handler

↓

Conversation Manager

↓

User
```

---

# Summary

The **Terminal Engine** is the execution layer responsible for securely managing shell commands, terminal sessions, processes, output streams, and operating system interactions within an AI Coding Agent.

A production-grade Terminal Engine should include:

- Terminal Engine
- Terminal Session
- Command Executor
- Process Manager
- Stream Handler
- Working Directory Manager
- Environment Manager
- Command History
- Terminal Security
- Terminal Validator
- Event Bus Integration

By providing validated command execution, interactive sessions, real-time streaming, process supervision, and strong security controls, the Terminal Engine enables AI Coding Agents such as OpenCode, Antigravity CLI, Cursor, Claude Code, and enterprise AI platforms to safely automate complex development workflows through the command line.