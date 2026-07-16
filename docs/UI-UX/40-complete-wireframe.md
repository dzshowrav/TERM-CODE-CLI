# 40-complete-wireframe.md

# Complete Wireframe System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the complete wireframe architecture for the Mobile AI CLI.

The wireframe describes the complete screen structure, component placement, interaction zones, information hierarchy, and responsive behavior.

The design combines:

- Modern AI CLI workflow
- Terminal-native interface
- Mobile-first interaction
- Keyboard-aware layout
- Touch-friendly controls
- Android Termux optimization

---

# Design Goals

The Complete Wireframe System must be:

- Mobile First
- No Sidebar Architecture
- Keyboard Aware
- Terminal Native
- Fast Navigation
- Minimal Distraction
- Touch Friendly
- CLI Optimized
- AI Coding Agent Focused

---

# Platform

Primary:

- Android
- Termux

Screen:

- Portrait First

Input:

- Touch
- Virtual Keyboard
- Physical Keyboard

---

# Core Layout Philosophy

The application uses a single-column workspace.

No sidebar.

No permanent navigation panel.

Everything is controlled through:

- Command Input
- Command Palette
- Context Views
- Modal Screens
- Keyboard Shortcuts

---

# Global Application Structure

```
Application

│
├── Header Area
│
├── Main Content Area
│
├── Command Input Area
│
└── Status Bar
```

---

# Complete Screen Layout

```
┌──────────────────────────────┐
│ Header                       │
│ App Name | Session | Model   │
├──────────────────────────────┤
│                              │
│                              │
│ Conversation Area            │
│                              │
│ User Messages                │
│ AI Responses                 │
│ Tool Results                 │
│ Code Blocks                  │
│                              │
│                              │
├──────────────────────────────┤
│ Command Input                │
│ > Type message...            │
├──────────────────────────────┤
│ Status Bar                   │
│ Model | Context | Status     │
└──────────────────────────────┘
```

---

# Layout Zones

The application contains five primary zones:

1. Header Zone
2. Conversation Zone
3. Tool Execution Zone
4. Command Input Zone
5. Status Bar Zone

---

# 1. Header Zone

## Purpose

Provides global context.

---

# Header Layout

```
┌──────────────────────────────┐
│ Mobile AI CLI                │
│ Session: Project X           │
│ Model: GPT-5                 │
└──────────────────────────────┘
```

---

# Header Information

Contains:

- Application Name
- Current Session
- Active Model
- Workspace Name
- Connection Status

---

# Header Rules

Must:

- Remain compact
- Never become a sidebar
- Never hide important information

---

# 2. Conversation Zone

## Purpose

Main workspace.

Contains:

- User prompts
- AI responses
- Tool execution
- Thinking states
- Streaming output

---

# Conversation Wireframe

```
┌──────────────────────────────┐
│                              │
│ User                         │
│ > Create React component     │
│                              │
│ AI                           │
│ I will create the component. │
│                              │
│ Tool                         │
│ Reading files...             │
│                              │
│ AI                           │
│ Completed                    │
│                              │
└──────────────────────────────┘
```

---

# Message Structure

Each message contains:

```
Author

↓

Content

↓

Metadata
```

---

# User Message

Example:

```
User

> Build authentication system
```

---

# AI Message

Example:

```
Assistant

I will analyze the project.
```

---

# System Message

Example:

```
System

Model changed to GPT-5
```

---

# Tool Message

Example:

```
Tool

filesystem.read

Success
```

---

# 3. Thinking View

## Purpose

Shows AI processing state.

---

# Wireframe

```
┌──────────────────────────────┐
│ AI Thinking                  │
│                              │
│ Analyzing project structure  │
│                              │
└──────────────────────────────┘
```

---

# Rules

Must:

- Be collapsible
- Not block interaction
- Support streaming

---

# 4. Tool Execution Zone

## Purpose

Shows agent actions.

---

# Tool Layout

```
┌──────────────────────────────┐
│ Tool Execution               │
├──────────────────────────────┤
│ filesystem.search            │
│                              │
│ Searching src/               │
│                              │
│ Status: Running              │
└──────────────────────────────┘
```

---

# Tool Information

Shows:

- Tool Name
- Input
- Output
- Status
- Duration

---

# Tool States

```
Waiting

Running

Completed

Failed

Cancelled
```

---

# 5. Command Input Zone

## Most Important Area

The command input is always accessible.

---

# Closed Keyboard Layout

```
┌──────────────────────────────┐
│ Conversation                 │
│                              │
├──────────────────────────────┤
│ > Write command              │
├──────────────────────────────┤
│ Status Bar                   │
└──────────────────────────────┘
```

---

# Open Keyboard Layout

```
┌──────────────────────────────┐
│ Conversation                 │
│                              │
├──────────────────────────────┤
│ > Write command              │
├──────────────────────────────┤
│ Status Bar                   │
├──────────────────────────────┤
│ Android Keyboard             │
└──────────────────────────────┘
```

---

# Command Input Features

Supports:

- Text Input
- Commands
- Autocomplete
- History
- Multiline
- Paste
- File References

---

# Command Prefix

Commands start with:

```
/
```

Examples:

```
/model

/settings

/help

/search
```

---

# 6. Status Bar

## Position Rule

Always:

```
Command Input

↓

Status Bar
```

---

# Status Bar Layout

```
┌──────────────────────────────┐
│ GPT-5 | 60% Context | Ready  │
└──────────────────────────────┘
```

---

# Status Information

Contains:

- Active Model
- Token Usage
- Connection
- Tool State
- Mode

---

# Command Palette Wireframe

Opened by:

```
Ctrl + P
```

or

```
/command
```

---

# Layout

```
┌──────────────────────────────┐
│ Command Palette               │
├──────────────────────────────┤
│ Search command...             │
├──────────────────────────────┤
│ > Create File                 │
│   Search Project              │
│   Change Model                │
└──────────────────────────────┘
```

---

# Model Selector Wireframe

```
┌──────────────────────────────┐
│ Select Model                 │
├──────────────────────────────┤
│ Search...                    │
├──────────────────────────────┤
│ ✓ GPT-5                      │
│   Claude                     │
│   Gemini                     │
│   Local Models               │
└──────────────────────────────┘
```

---

# File Picker Wireframe

```
┌──────────────────────────────┐
│ Select File                  │
├──────────────────────────────┤
│ Search...                    │
├──────────────────────────────┤
│ src/                         │
│ ├── App.tsx                  │
│ ├── main.ts                  │
│ └── index.css                │
└──────────────────────────────┘
```

---

# Settings Wireframe

```
┌──────────────────────────────┐
│ Settings                     │
├──────────────────────────────┤
│ Search Settings              │
├──────────────────────────────┤
│ AI                           │
│ Models                       │
│ Tools                        │
│ Appearance                   │
│ Permissions                  │
│ Storage                      │
└──────────────────────────────┘
```

---

# Dialog Wireframe

```
┌──────────────────────────────┐
│ Confirm Action               │
├──────────────────────────────┤
│ Delete Session?              │
│ This cannot be undone.       │
├──────────────────────────────┤
│ Cancel       Confirm         │
└──────────────────────────────┘
```

---

# Error Wireframe

```
┌──────────────────────────────┐
│ Error                        │
├──────────────────────────────┤
│ File not found               │
│                              │
│ Retry                        │
└──────────────────────────────┘
```

---

# Loading Wireframe

```
┌──────────────────────────────┐
│ Processing                   │
│                              │
│ Analyzing files...           │
│                              │
└──────────────────────────────┘
```

---

# Empty State Wireframe

```
┌──────────────────────────────┐
│                              │
│ No Workspace                 │
│                              │
│ Select a project folder      │
│                              │
│ [Select Folder]              │
│                              │
└──────────────────────────────┘
```

---

# Mobile Responsive Rules

The interface must adapt to:

- Small screens
- Keyboard visibility
- Long messages
- Large code blocks

---

# Keyboard Priority

When keyboard opens:

Priority:

```
Input

↓

Latest Message

↓

Context
```

---

# Touch Priority

Touch targets:

- Minimum 44px
- Clear selection
- No hidden actions

---

# Navigation Model

No sidebar.

Navigation uses:

```
Command Palette

+

Modal Views

+

Keyboard Shortcuts
```

---

# State Management

Global states:

```
Idle

Thinking

Streaming

Executing Tool

Error

Complete
```

---

# Performance Rules

The wireframe must support:

- Virtual scrolling
- Partial rendering
- Lazy loading
- Cached views

---

# Accessibility

Must support:

- Screen readers
- High contrast
- Keyboard navigation
- Reduced motion

---

# Security

Never display:

- API keys
- Tokens
- Private credentials

---

# Restrictions

Never:

- Add sidebar
- Hide command input
- Move status bar above input
- Use emoji icons
- Replace icons with text
- Block keyboard interaction

---

# Complete User Flow

```
Open Application

↓

Restore Session

↓

Show Chat

↓

User Types Prompt

↓

AI Processing

↓

Tool Execution

↓

Streaming Response

↓

Complete
```

---

# Wireframe Checklist

Every screen must:

- Follow single-column layout
- Support keyboard
- Support touch
- Maintain status bar position
- Use design tokens
- Support themes
- Handle loading states
- Handle errors
- Work on Termux
- Remain mobile-first

---

# Core Rules

1. No sidebar architecture.
2. Command input is always accessible.
3. Status bar stays below input.
4. Keyboard is part of the layout.
5. Terminal rendering is primary.
6. Every action has feedback.
7. Every screen supports mobile usage.
8. Touch and keyboard coexist.
9. Performance is prioritized.
10. Design must remain CLI-native.

---

# Summary

The Complete Wireframe System defines the entire Mobile AI CLI interface structure. It creates a focused, single-column, keyboard-aware, terminal-native workspace inspired by modern AI coding agents while maintaining strict mobile-first compatibility for Android Termux.

The wireframe provides the foundation for implementing every screen, component, interaction, and workflow in the AI CLI ecosystem.