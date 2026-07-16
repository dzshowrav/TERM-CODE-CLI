# 13-tool-execution.md

# Tool Execution
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines how tool execution is displayed and managed inside the Mobile AI CLI.

A "tool" is any external action performed by the AI, including reading files, editing code, searching the workspace, executing shell commands, calling MCP servers, accessing APIs, or interacting with Git.

Tool execution must always feel transparent, predictable, and non-blocking.

---

# Design Goals

The Tool Execution system must be

- Mobile First
- Terminal Native
- Transparent
- Real-Time
- Non-Blocking
- Streaming Friendly
- Keyboard Safe
- Performance Optimized

---

# Design Philosophy

The AI should never perform hidden actions.

Every tool invocation must be visible to the user.

Users should always know

- Which tool is running
- What it is doing
- Whether it succeeded
- Whether it failed

---

# Execution Flow

```
User Prompt

↓

AI Planning

↓

Tool Request

↓

Permission Check

↓

Tool Execution

↓

Streaming Output

↓

Result

↓

AI Response
```

---

# Tool Types

Supported

- File System
- Shell
- Git
- Search
- Workspace Indexer
- MCP
- HTTP API
- Database
- Package Manager
- AI Sub-Agent

---

# Display Location

Tool execution appears **inside the Conversation Area**.

It never opens a new page.

It never replaces the conversation.

---

# Execution Block

Each tool execution is displayed as a structured block.

Example

```text
Tool

Reading Files

Status: Running
```

---

# Tool Header

Every execution block contains

- Tool Name
- Current Status
- Timestamp (optional)

---

# Tool States

Supported

```
Pending

Waiting Permission

Running

Streaming

Completed

Cancelled

Failed

Timed Out
```

Only one state is active.

---

# Pending State

Displayed immediately after the AI decides to use a tool.

Example

```text
Tool

Preparing...
```

---

# Permission State

If permission is required

Display

```text
Waiting for Permission...
```

Execution pauses until the user responds.

---

# Running State

Display

```text
Reading Files...
```

or

```text
Searching Workspace...
```

The user can continue reading previous messages.

---

# Streaming State

If the tool produces incremental output

Display updates continuously.

Example

```text
Searching...

12 files scanned

34 files scanned

91 files scanned

Completed
```

---

# Completed State

Display

```text
Completed

18 files indexed.
```

---

# Failed State

Display

```text
Failed

Permission denied.
```

A readable explanation must always be included.

---

# Cancelled State

Display

```text
Cancelled by User
```

---

# Timeout State

Display

```text
Operation Timed Out
```

Retry may be offered.

---

# Inline Output

Small outputs remain inline.

Example

```text
Tool

Git Status

3 modified files
```

---

# Large Output

Large outputs become collapsible.

Collapsed example

```text
Tool

Shell Output

Show Output
```

Expanded example

```text
Tool

npm install

...

Complete output
```

---

# Code Output

Rendered using

- Monospaced text
- Preserved formatting
- Horizontal scrolling

Never reformat tool output.

---

# File Changes

When files are modified

Display

```text
Modified

src/App.tsx

package.json
```

---

# Created Files

Display

```text
Created

README.md
```

---

# Deleted Files

Display

```text
Deleted

old-config.json
```

---

# Search Results

Display

```text
Searching...

42 matches found.
```

Results remain expandable.

---

# Shell Commands

Display

```text
Running

npm install
```

Command output streams live.

---

# Git Operations

Examples

```text
Git

Commit Created
```

```text
Git

Branch Switched
```

```text
Git

Merge Completed
```

---

# MCP Operations

Display

```text
MCP

Connected
```

or

```text
MCP

Fetching Context...
```

---

# API Calls

Display

```text
HTTP

GET /projects
```

Sensitive information must never be displayed.

---

# Database Operations

Display

```text
Database

Running Query...
```

Results remain expandable.

---

# Progress Indicator

Long-running operations may display

- Percentage
- Item Count
- Current Task

Example

```text
Indexing

214 / 830 files
```

---

# Multiple Tools

If multiple tools execute

Display them in chronological order.

Never overlap outputs.

---

# Nested Tool Calls

Nested execution is allowed.

Example

```text
Workspace Indexer

↓

Search

↓

File Read
```

Indent nested tools.

---

# Conversation Behavior

Tool execution never interrupts

- Typing
- Scrolling
- Reading

The user may continue interacting with the application.

---

# Status Bar Integration

During execution

Status Bar displays

```text
Executing
```

or

```text
Searching
```

---

# Keyboard Behavior

Keyboard remains usable.

Users may type the next prompt while tools execute.

---

# Safe Area

Tool blocks remain inside the Conversation Area.

They never overlap

- Command Input
- Status Bar
- Keyboard

---

# Accessibility

Support

- High Contrast
- Screen Readers
- Large Fonts

Tool states must always include readable text.

---

# Performance

Render only changed output.

Avoid full conversation redraws.

Stream incrementally.

---

# Error Handling

Every failure must include

- Tool Name
- Failure State
- Human-readable Reason

Never expose internal stack traces by default.

---

# Retry

When appropriate

Display

```text
Retry Available
```

The user controls retries.

---

# Security

Never execute tools silently.

Never hide

- File changes
- Shell commands
- Network requests

Every action must be visible.

---

# Restrictions

Never

- Hide running tools
- Freeze the interface
- Block scrolling
- Block typing
- Open separate execution pages
- Display raw internal errors without context

---

# Example Layout

```text
Assistant

Creating React project...

Tool

Running

npm create vite@latest

Completed

Project created successfully.

Assistant

The project is ready.
```

---

# Long Operation Example

```text
Tool

Workspace Indexer

Running

214 / 830 files indexed

527 / 830 files indexed

830 / 830 files indexed

Completed
```

---

# Tool Execution Checklist

Every tool execution must

- Be visible
- Show current state
- Stream progress
- Preserve formatting
- Stay inside the conversation
- Respect the safe area
- Support retry
- Report failures clearly
- Never block typing
- Never block scrolling

---

# Core Rules

1. Every tool execution is visible.
2. Tool execution stays inside the Conversation Area.
3. Running tools stream output incrementally.
4. Users may continue typing while tools execute.
5. Every execution has a clear state.
6. File changes are always reported.
7. Sensitive information is never exposed.
8. Long outputs are collapsible.
9. Failures always include readable explanations.
10. Tool execution remains transparent and terminal-native.

---

# Summary

The Tool Execution system provides a transparent, real-time view of every action performed by the AI. Whether reading files, executing shell commands, searching the workspace, communicating with MCP servers, or modifying code, every operation is displayed inline within the conversation. This design keeps users informed, preserves trust, supports multitasking, and delivers a responsive, mobile-first terminal experience optimized for Android Termux.