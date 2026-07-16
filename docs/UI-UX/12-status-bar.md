# 12-status-bar.md

# Status Bar
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

The Status Bar is the permanent information panel located at the bottom of the Mobile AI CLI.

It provides real-time application status without interrupting the conversation or command input.

The Status Bar must remain visible at all times.

---

# Design Goals

The Status Bar must be

- Mobile First
- Terminal Native
- Always Visible
- Compact
- Readable
- Real-Time
- Non-Intrusive
- Performance Friendly

---

# Supported Platform

Primary

- Android
- Termux

Orientation

- Portrait Only

---

# Design Philosophy

The Status Bar communicates system state rather than controls.

It should answer the question:

"What is happening right now?"

without distracting the user.

---

# Position

Always below the Command Input.

```
Conversation

↓

Command Input

↓

Status Bar
```

Nothing may appear below the Status Bar except the Android keyboard.

---

# Height

Fixed

```
1 Row
```

The height never changes.

---

# Width

```
100%
```

Uses the entire terminal width.

---

# Layout

```
┌────────────────────────────────────┐
│ GPT-5 │ Main │ Ready │ 2.8K │ UTF-8 │
└────────────────────────────────────┘
```

---

# Information Priority

Highest Priority

- Model
- Current State

Medium Priority

- Session
- Workspace
- Git Branch

Lower Priority

- Token Count
- Encoding
- Line Ending

---

# Standard Sections

The Status Bar may contain

- AI Model
- Workspace
- Session
- Git Branch
- Current Mode
- Token Count
- Context Size
- Connection Status
- Tool Status
- Memory Usage
- CPU Usage
- File Encoding
- Line Ending
- Time

Only show relevant items.

---

# AI Model

Displays

Current active model.

Examples

```
GPT-5

Claude

Gemini

Qwen

DeepSeek
```

Updates immediately after switching.

---

# Workspace

Displays

Current workspace name.

Example

```
my-project
```

---

# Session

Displays

Current conversation session.

Example

```
Session 3
```

---

# Git Branch

Displays

Current Git branch.

Example

```
main
```

Hidden if Git is unavailable.

---

# Current State

Possible values

```
Ready

Thinking

Streaming

Executing

Searching

Indexing

Waiting

Offline

Error
```

Only one state is active.

---

# Token Count

Displays

Approximate context usage.

Example

```
2.8K
```

Updates automatically.

---

# Context Size

Optional.

Shows

```
48K

128K

256K
```

Useful for long conversations.

---

# Connection Status

Possible values

```
Online

Offline

Connecting

Disconnected
```

Displayed only when applicable.

---

# Tool Status

Displays

Currently running tool.

Examples

```
Git

Search

MCP

Terminal

Indexer
```

Hidden when idle.

---

# Memory Usage

Optional.

Example

```
420 MB
```

Useful for local models.

---

# CPU Usage

Optional.

Example

```
12%
```

Shown only during heavy processing.

---

# File Encoding

Examples

```
UTF-8

UTF-16
```

Shown only in editor-related contexts.

---

# Line Ending

Examples

```
LF

CRLF
```

Hidden unless editing files.

---

# Time

Optional.

Format

```
HH:MM
```

Displayed only if enabled.

---

# Item Order

Recommended order

```
Model

↓

Workspace

↓

Branch

↓

State

↓

Tokens

↓

Connection
```

Less important items are removed first when space is limited.

---

# Dynamic Visibility

Small terminal width

Hide

- Time
- Encoding
- Memory
- CPU

Keep

- Model
- State
- Workspace

Always prioritize essential information.

---

# Color Rules

Model

Primary

Ready

Success

Thinking

Accent

Streaming

Accent

Searching

Info

Executing

Primary

Warning

Warning

Error

Error

Offline

Muted

---

# Status Updates

The Status Bar updates immediately when

- Model changes
- Session changes
- Workspace changes
- Tool starts
- Tool ends
- AI starts streaming
- AI finishes streaming
- Connection changes

---

# Animation

Allowed

Minimal text updates.

Not Allowed

- Flashing
- Bouncing
- Sliding
- Decorative animations

---

# Keyboard Behavior

Keyboard Closed

```
Conversation

↓

Command Input

↓

Status Bar
```

Keyboard Open

```
Conversation Reduced

↓

Command Input

↓

Status Bar

↓

Keyboard
```

Status Bar always remains visible.

---

# Safe Area

The Status Bar must always remain inside the protected safe area.

Never overlap

- Keyboard
- Gesture Navigation
- Display Cutout

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Every status item must have readable text.

---

# Performance

Update only changed fields.

Avoid redrawing the entire Status Bar.

---

# Error Handling

If a value cannot be determined

Display

```
Unknown
```

Never leave corrupted or partially rendered text.

---

# Responsive Behavior

When width decreases

Hide optional fields.

Never truncate

- Current State
- Model

---

# Restrictions

Never

- Make the Status Bar scroll
- Increase height
- Hide during streaming
- Hide during typing
- Display pop-up controls
- Replace text with icons only

---

# Example Layout

```
┌────────────────────────────────────┐
│ Assistant                          │
│                                    │
│ Project created successfully.      │
│                                    │
├────────────────────────────────────┤
│ > npm run dev                      │
├────────────────────────────────────┤
│ GPT-5 │ my-app │ main │ Ready │ 2.8K │
└────────────────────────────────────┘
```

---

# Compact Layout

Small terminal

```
GPT-5 │ Ready │ 2.8K
```

Medium terminal

```
GPT-5 │ my-app │ Ready │ 2.8K
```

Large terminal

```
GPT-5 │ my-app │ main │ Ready │ 2.8K │ Online │ UTF-8
```

---

# Status Bar Checklist

Every Status Bar must

- Stay visible
- Stay one row high
- Display the active model
- Display the current state
- Update in real time
- Respect safe areas
- Adapt to screen width
- Preserve readability
- Avoid unnecessary redraws

---

# Core Rules

1. The Status Bar is always visible.
2. It is always below the Command Input.
3. Height is permanently one row.
4. Display only meaningful information.
5. Update fields independently.
6. Hide optional items on small screens.
7. Preserve the current state indicator.
8. Never interrupt user interaction.
9. Remain keyboard-safe.
10. Optimize for Android Termux and mobile-first usage.

---

# Summary

The Status Bar is the permanent system information layer of the Mobile AI CLI. It provides concise, real-time feedback about the AI model, workspace, session, execution state, tokens, and connection without distracting from the conversation. Its fixed one-row layout, responsive behavior, and terminal-native design ensure that users always understand the application's current state while maintaining maximum screen space for coding and AI interaction.