# 02-tui-rendering-pipeline.md

# TUI Rendering Pipeline
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a TUI Rendering Pipeline?

A **TUI (Terminal User Interface) Rendering Pipeline** is the complete process that converts application state into what the user sees in the terminal.

It is similar to a browser rendering HTML into pixels, but instead renders text, colors, borders, progress bars, input boxes, and layouts into terminal cells.

```
Application State
        │
        ▼
Rendering Pipeline
        │
        ▼
Terminal Screen
```

---

# Why is a Rendering Pipeline Needed?

Without a pipeline:

```
Agent

↓

print()

↓

print()

↓

print()

↓

print()
```

Problems

- Screen flickers
- Layout breaks
- Slow updates
- Cursor jumps
- Difficult state management

---

With a Rendering Pipeline

```
State

↓

Diff Engine

↓

Layout Engine

↓

Renderer

↓

Terminal
```

Only changed parts are redrawn.

---

# Goals

A good rendering pipeline should provide

- Fast rendering
- Flicker-free updates
- Stable cursor
- Efficient redraw
- Responsive layout
- Theme support
- Keyboard interaction
- Animation support
- Scroll support
- Resize support

---

# High-Level Architecture

```
User Input
      │
      ▼
Input Manager
      │
      ▼
State Manager
      │
      ▼
Render Scheduler
      │
      ▼
Virtual Screen
      │
      ▼
Diff Engine
      │
      ▼
Terminal Renderer
      │
      ▼
Real Terminal
```

---

# Folder Structure

```
src/

tui/

    App.ts

    Renderer.ts

    RenderScheduler.ts

    LayoutEngine.ts

    VirtualScreen.ts

    DiffEngine.ts

    CursorManager.ts

    ScreenBuffer.ts

    TerminalWriter.ts

    Viewport.ts

    Components/

    Widgets/

    Input/

    Animation/

    Theme/

    Hooks/

    Utils/
```

---

# Core Components

## App

Root UI component.

Responsibilities

- Build screen
- Hold layout
- Connect state

---

## State Manager

Stores

```
Messages

Cursor

Selection

Theme

Scroll

Loading

Progress

Status
```

---

## Render Scheduler

Controls

- When rendering starts
- Frame timing
- Batched updates
- Prevent unnecessary redraw

---

## Virtual Screen

Stores an in-memory copy of the terminal.

Example

```
----------------------------------

OpenCode

Hello

Thinking...

----------------------------------
```

Nothing is written to the real terminal yet.

---

## Diff Engine

Compares

```
Old Screen

↓

New Screen
```

Only changed cells are updated.

Example

Old

```
Thinking...
```

New

```
Done
```

Only one line changes.

---

## Terminal Renderer

Writes ANSI escape sequences.

Responsibilities

- Cursor movement
- Color output
- Line updates
- Clear operations

---

# Rendering Lifecycle

```
State Change

↓

Schedule Render

↓

Create Virtual Screen

↓

Layout

↓

Diff

↓

Write Terminal

↓

Done
```

---

# Rendering Pipeline

```
Application State

↓

Component Tree

↓

Layout Calculation

↓

Virtual Screen

↓

Diff Engine

↓

ANSI Commands

↓

Terminal
```

---

# Component Tree

Example

```
App

├── Header

├── Sidebar

├── Chat

├── StatusBar

└── Input
```

Every component returns visual content.

---

# Layout Engine

Calculates

```
Width

Height

Padding

Margin

Alignment

Position
```

---

# Layout Example

```
+------------------------------+

Header

+------------------------------+

Chat Window

Chat Window

Chat Window

+------------------------------+

Input

+------------------------------+
```

---

# Virtual Screen

Represents terminal as rows and columns.

Example

```
Cell

Row

Column

Character

Foreground

Background

Style
```

Every terminal position becomes one cell.

---

# Screen Buffer

Stores

```
Character

Foreground

Background

Bold

Italic

Underline

Position
```

---

# Diff Algorithm

Compare

```
Previous Buffer

↓

Current Buffer
```

If identical

```
Skip
```

If changed

```
Render
```

---

# Rendering Modes

## Full Render

Entire screen.

```
Clear

↓

Render Everything
```

Used

- Startup
- Resize

---

## Partial Render

Only changed regions.

```
Old

↓

New

↓

Update Differences
```

Used most of the time.

---

# Cursor Manager

Controls

```
Cursor Position

Visibility

Blink

Restore
```

Never allow rendering to leave the cursor in the wrong place.

---

# ANSI Renderer

Outputs commands like

```
Move Cursor

Clear Line

Clear Screen

Set Color

Hide Cursor

Show Cursor
```

---

# Input Pipeline

```
Keyboard

↓

Parser

↓

Action

↓

State Update

↓

Render
```

Example

User presses

```
Arrow Down
```

↓

Selection changes

↓

UI rerenders

---

# Resize Handling

```
Terminal Resize

↓

Viewport Update

↓

Layout Recalculate

↓

Full Render
```

---

# Scroll Pipeline

```
Wheel

↓

Scroll Manager

↓

Viewport

↓

Render Visible Area
```

Never render hidden content.

---

# Animation Pipeline

```
Timer

↓

Frame Update

↓

State

↓

Render
```

Examples

- Spinner
- Progress Bar
- Typing Cursor

---

# Theme Engine

Provides

```
Foreground

Background

Accent

Border

Highlight

Selection

Status
```

Renderer never hardcodes colors.

---

# Event Integration

Example

```
agent:thinking

↓

State.loading=true

↓

Render

↓

Spinner Appears
```

---

# File Watcher Integration

```
File Changed

↓

State Updated

↓

Render

↓

Explorer Refresh
```

---

# MCP Integration

```
Tool Running

↓

Progress Updated

↓

Status Bar Refresh
```

---

# Agent Integration

```
LLM Token

↓

Append Message

↓

Render

↓

Streaming Output
```

---

# Performance Optimizations

Use

- Double buffering
- Virtual screen
- Dirty rectangle rendering
- Cell diffing
- Batched rendering
- Render throttling
- Frame scheduling

Avoid

- Full screen redraw
- Clearing screen every frame
- Printing line-by-line repeatedly

---

# Dirty Rectangles

Instead of

```
Render Entire Screen
```

Render only

```
Changed Region
```

Example

```
Header

Chat

Status

Input
```

Only

```
Chat
```

changes.

---

# Frame Timing

Typical target

```
30 FPS

or

60 FPS
```

Terminal usually requires much less.

---

# Error Handling

```
Render Error

↓

Catch

↓

Log

↓

Restore Cursor

↓

Continue
```

Never leave the terminal broken.

---

# Best Practices

Always

- Separate state from rendering
- Use virtual buffers
- Batch updates
- Keep renderer stateless
- Render only visible content
- Hide cursor during render
- Restore cursor after render

Never

- Mix business logic with UI
- Print directly from Agent
- Redraw unchanged content
- Block rendering with long tasks

---

# Common Mistakes

Bad

```
Agent

↓

console.log()

↓

console.log()

↓

console.log()
```

Produces flicker.

---

Good

```
Agent

↓

State

↓

Renderer

↓

Diff

↓

Terminal
```

---

# Testing Checklist

- Initial render
- Partial update
- Full redraw
- Resize
- Scroll
- Cursor movement
- Theme switching
- Animation
- Long messages
- Streaming output
- Error recovery

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Gemini CLI
- Lazygit
- k9s
- btop
- GitUI
- Bubble Tea applications
- Ink-based React CLI applications

---

# Summary

The **TUI Rendering Pipeline** is responsible for transforming application state into an efficient, flicker-free terminal interface.

A production-grade AI coding agent should use:

- State-driven rendering
- Virtual screen buffering
- Layout calculation
- Diff-based updates
- ANSI terminal rendering
- Efficient cursor management
- Responsive resize handling
- Streaming-friendly rendering

By separating rendering from business logic and updating only changed portions of the screen, the application achieves smooth performance, scalability, and a professional user experience comparable to modern tools like OpenCode and Antigravity CLI.