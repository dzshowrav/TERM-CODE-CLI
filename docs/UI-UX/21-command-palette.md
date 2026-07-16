# 21-command-palette.md

# Command Palette
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Command Palette used throughout the Mobile AI CLI.

The Command Palette provides a fast, searchable interface for executing commands, switching models, opening workspaces, navigating files, changing themes, accessing tools, and triggering application actions without leaving the Chat Screen.

The Command Palette must be optimized for Android Termux and mobile-first interaction.

---

# Design Goals

The Command Palette must be

- Mobile First
- Terminal Native
- Search Driven
- Keyboard Friendly
- Touch Friendly
- Fast
- Non-Intrusive
- Accessible

---

# Supported Platform

Primary

- Android
- Termux

Orientation

- Portrait Only

---

# Design Philosophy

Everything should be searchable.

Users should never need to remember where a feature is located.

The Command Palette acts as the universal launcher for the application.

---

# Display Location

The Command Palette appears as a modal overlay above the Chat Screen.

The conversation remains visible in the background.

---

# Layout

```text
┌────────────────────────────────────┐
│ Command Palette                    │
├────────────────────────────────────┤
│ Search commands...                 │
├────────────────────────────────────┤
│ /new-session                       │
│ /clear                             │
│ /theme                             │
│ /model                             │
│ /workspace                         │
├────────────────────────────────────┤
│ Esc                    Enter       │
└────────────────────────────────────┘
```

---

# Components

The Command Palette contains

- Header
- Search Field
- Result List
- Optional Category
- Footer Actions

---

# Header

Displays

```text
Command Palette
```

---

# Search Field

Always visible.

Receives focus immediately after opening.

Supports

- Command Search
- Fuzzy Search
- Prefix Match
- Partial Match

---

# Result List

Displays matching actions.

Updates immediately while typing.

---

# Categories

Optional grouping

Examples

- Commands
- Models
- Workspace
- Files
- Git
- Tools
- Settings
- Themes
- Sessions

---

# Commands

Examples

```text
/new-session

/clear

/help

/history

/settings

/theme

/model

/workspace

/search

/index

/git

/tools
```

---

# Models

Examples

```text
GPT-5

Claude

Gemini

Qwen

DeepSeek
```

Selecting a model switches the active provider.

---

# Workspace Actions

Examples

```text
Open Workspace

Recent Workspace

Close Workspace

Refresh Workspace
```

---

# File Actions

Examples

```text
Open File

Recent Files

Search Files

Reveal File
```

---

# Git Actions

Examples

```text
Git Status

Commit

Pull

Push

Checkout Branch
```

---

# Tool Actions

Examples

```text
Run Terminal

Run Search

Run Indexer

Open MCP
```

---

# Settings

Examples

```text
Theme

Font Size

Renderer

Permissions

Shortcuts
```

---

# Theme Actions

Examples

```text
Dark

Light

High Contrast

Terminal Classic
```

---

# Session Actions

Examples

```text
New Session

Switch Session

Rename Session

Delete Session
```

---

# Search Behavior

Search updates instantly.

No confirmation required.

---

# Fuzzy Search

Supported.

Example

```text
mdl

↓

Model
```

---

# Ranking

Results ranked by

1. Exact Match
2. Prefix Match
3. Fuzzy Match
4. Recent Usage

---

# Recent Commands

Frequently used commands appear near the top.

---

# Selection

Supported

- Keyboard
- Touch

Only one active item at a time.

---

# Confirmation

Selecting an item executes the associated action.

---

# Cancellation

Cancel closes the Command Palette.

No changes are applied.

---

# Keyboard Navigation

Supported

- Move Selection
- Confirm
- Cancel
- Scroll Results

---

# Touch Navigation

Supported

- Tap
- Scroll
- Long Press (optional)

---

# Long Press

May display

- Description
- Shortcut
- Usage Information

---

# Shortcuts

Optional display

Example

```text
/theme

Change application theme.
```

---

# Empty Result

Display

```text
No matching commands.
```

---

# Dynamic Results

Results update while

- Typing
- Workspace changes
- Model changes
- Tool availability changes

---

# Safe Area

The Command Palette respects

- Display Cutout
- Keyboard
- Gesture Navigation

It never overlaps protected regions.

---

# Keyboard Behavior

Keyboard opening

Reduces result list height.

Search field remains visible.

---

# Streaming Integration

The Command Palette remains usable while AI responses stream.

Streaming continues in the background.

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Search results remain readable.

---

# Performance

Large command sets

Filter incrementally.

Only visible rows render.

---

# Error Handling

Unknown command

Display

```text
Command not found.
```

Unavailable command

Display

```text
Currently unavailable.
```

---

# Security

Only commands permitted by the current session and permissions are displayed.

Restricted actions remain hidden.

---

# Restrictions

Never

- Replace the Chat Screen
- Freeze while searching
- Hide the search field
- Execute commands automatically
- Display duplicate results
- Block typing

---

# Example Command Palette

```text
Command Palette

Search...

/new-session
/model
/theme
/settings
/workspace

Esc            Enter
```

---

# Example Search

```text
Search: git

Git Status

Git Commit

Git Push

Git Pull
```

---

# Example Model Search

```text
Search: g

GPT-5

Gemini
```

---

# Command Palette Checklist

Every Command Palette must

- Support fuzzy search
- Support command execution
- Support model switching
- Support workspace actions
- Respect safe areas
- Support touch navigation
- Support keyboard navigation
- Update dynamically
- Remain responsive
- Stay terminal-native

---

# Core Rules

1. The Command Palette appears as a modal overlay.
2. Search receives focus immediately.
3. Results update while typing.
4. Commands require explicit confirmation.
5. Fuzzy search is supported.
6. Recent commands influence ranking.
7. Respect workspace permissions.
8. Never replace the Chat Screen.
9. Remain keyboard-safe.
10. Optimize for Android Termux.

---

# Summary

The Command Palette is the universal action launcher of the Mobile AI CLI. It provides a fast, searchable interface for commands, models, tools, workspaces, files, Git operations, and settings without interrupting the current conversation. Its terminal-native design, incremental search, and mobile-first behavior make it an essential navigation component for Android Termux workflows.