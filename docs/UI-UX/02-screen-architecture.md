# 02-screen-architecture.md

# Screen Architecture
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the complete screen architecture for the Mobile AI CLI.

It specifies how every screen is organized, how components communicate, and how the interface behaves during user interaction.

Every screen in the application must follow this architecture.

---

# Design Goals

The architecture must be

- Mobile First
- Terminal Native
- One-Hand Friendly
- Keyboard Safe
- Fast
- Predictable
- Consistent
- Modular
- Reactive

---

# Supported Platform

Primary Platform

- Android
- Termux

Orientation

- Portrait Only

---

# Architecture Philosophy

The application behaves like a terminal instead of a traditional mobile app.

Users never navigate through multiple pages.

Instead, the interface updates the current screen.

---

# Screen Philosophy

Every screen represents one active workspace.

The workspace changes its content instead of navigating to another page.

---

# Global Screen Structure

```
Application

↓

Workspace

↓

Screen

↓

Layout

↓

Components

↓

Terminal Renderer
```

---

# Screen Hierarchy

```
AI CLI

↓

Workspace

↓

Conversation Screen

↓

Conversation Area

↓

Input Area

↓

Status Bar
```

---

# Screen Types

The application contains the following logical screens.

```
Splash

Welcome

Conversation

Tool Execution

Command Palette

File Picker

Model Selector

Settings

Permission Dialog

Confirmation Dialog

Error Screen
```

Only one primary screen is active at a time.

---

# Navigation Model

Navigation never depends on sidebars.

Navigation occurs by

- Commands
- Dialogs
- Temporary overlays

---

# Root Screen

```
┌──────────────────────────────┐
│                              │
│      Conversation Area       │
│                              │
│                              │
│                              │
├──────────────────────────────┤
│ > Input                      │
├──────────────────────────────┤
│ Status Bar                   │
└──────────────────────────────┘
```

---

# Conversation Screen

Purpose

Primary AI interaction.

Contains

- User messages
- AI messages
- Markdown
- Code
- Tool output
- Search results

---

# Input Area

Purpose

Collect user commands.

Supports

- Multi-line input
- Cursor
- Selection
- Paste
- Keyboard shortcuts

---

# Status Bar

Purpose

Display runtime information.

Examples

```
Model

Session

Workspace

Tokens

Branch

Connection

Time
```

Always remains visible.

---

# Overlay Architecture

Temporary UI appears as overlays.

Allowed

```
Dialog

Command Palette

Permission Request

Confirmation

Model Selector

File Picker
```

After closing

```
Return

↓

Conversation Screen
```

---

# Overlay Rules

Overlay width

```
90%
```

Maximum height

```
70%
```

Background

```
Dim Terminal
```

Never cover the entire screen unless absolutely required.

---

# Screen Stack

```
Conversation

↓

Dialog

↓

Confirmation
```

Maximum stack depth

```
3
```

---

# Workspace

A workspace represents

- Current project
- Current directory
- Session state
- Active AI context

Only one workspace is active.

---

# Workspace Layout

```
Workspace

↓

Conversation

↓

Input

↓

Status
```

---

# Rendering Flow

```
Application State

↓

Screen State

↓

Layout

↓

Renderer

↓

Terminal
```

---

# Screen State

Every screen has

```
Visible

Hidden

Loading

Streaming

Error

Empty
```

---

# State Transitions

```
Loading

↓

Ready

↓

Streaming

↓

Ready

↓

Completed
```

---

# Loading Screen

```
┌──────────────────────────────┐
│                              │
│                              │
│        Loading...            │
│                              │
│                              │
├──────────────────────────────┤
│                              │
├──────────────────────────────┤
│ Starting Session             │
└──────────────────────────────┘
```

---

# Empty Screen

```
┌──────────────────────────────┐
│                              │
│     No Conversation          │
│                              │
│ Type a command below         │
│                              │
├──────────────────────────────┤
│ >                            │
├──────────────────────────────┤
│ Ready                        │
└──────────────────────────────┘
```

---

# Streaming Screen

```
Assistant

Generating...

Generating...

Generating...

Cursor
```

Updates incrementally.

Never redraw the full screen.

---

# Tool Execution Screen

When tools execute

Conversation remains visible.

Tool output appears inline.

Example

```
Reading Files...

Searching...

Editing...

Done
```

---

# Error Screen

Purpose

Explain the problem.

Show

- Error summary
- Recovery action

Never expose internal stack traces by default.

---

# Screen Persistence

Conversation

Persistent

Input

Persistent

Status

Persistent

Dialog

Temporary

Command Palette

Temporary

---

# Keyboard Behavior

When keyboard opens

Conversation height decreases.

Input remains fixed.

Status remains fixed.

Nothing is hidden.

---

# Scroll Behavior

Conversation

Scrollable

Input

Static

Status

Static

---

# Focus Management

Only one component has focus.

Priority

```
Input

↓

Dialog

↓

Command Palette

↓

Conversation Selection
```

---

# Gesture Support

Supported

- Tap
- Long Press
- Vertical Scroll

Unsupported

- Swipe Navigation
- Multi-finger Gestures
- Pinch Zoom

---

# Touch Targets

Minimum

```
44dp
```

Preferred

```
48dp
```

---

# Screen Communication

Screens communicate through

```
Event Bus
```

Never communicate directly.

---

# Component Tree

```
Application

└── Workspace

    └── Screen

        ├── Conversation

        ├── Input

        └── Status
```

---

# Resize Handling

On terminal resize

```
Detect

↓

Recalculate Layout

↓

Re-render

↓

Maintain Scroll Position
```

---

# Performance Rules

Never

- Rebuild entire screen
- Re-render unchanged components

Always

- Update incrementally
- Cache layouts
- Reuse components

---

# Accessibility

Support

- High contrast
- Large fonts
- Reduced motion
- Screen readers

---

# Restrictions

Never use

- Sidebars
- Floating Action Buttons
- Bottom Navigation
- Tab Bars
- Drawer Menus
- Desktop Windows
- Split Panels

---

# Master Screen Diagram

```
┌────────────────────────────────┐
│                                │
│                                │
│                                │
│        Conversation Area       │
│                                │
│                                │
│                                │
├────────────────────────────────┤
│ > User Input                   │
├────────────────────────────────┤
│ Model │ Tokens │ Git │ Ready   │
└────────────────────────────────┘
```

---

# Architecture Principles

1. One Active Screen
2. Single Column Layout
3. Bottom Input
4. Bottom Status Bar
5. Keyboard Safe
6. Conversation First
7. Overlay Instead of Navigation
8. Event Driven
9. Incremental Rendering
10. Mobile First

---

# Summary

The Screen Architecture defines a single-screen, workspace-oriented interface optimized for Android Termux.

The application never relies on sidebars or desktop navigation patterns. Instead, it uses a persistent conversation area, a fixed command input, and a fixed status bar, with lightweight overlays for temporary interactions.

This architecture ensures fast rendering, predictable navigation, one-handed usability, and a terminal-native experience suitable for an AI-powered mobile CLI.