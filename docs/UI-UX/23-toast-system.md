# 23-toast-system.md

# Toast System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Toast System used throughout the Mobile AI CLI.

Toasts provide short, non-blocking notifications that inform users about completed actions, errors, warnings, permissions, and system events without interrupting the current workflow.

The Toast System must be optimized for Android Termux and mobile-first interaction.

---

# Design Goals

The Toast System must be

- Mobile First
- Terminal Native
- Non-Blocking
- Lightweight
- Readable
- Accessible
- Keyboard Safe
- Performance Optimized

---

# Supported Platform

Primary

- Android
- Termux

Orientation

- Portrait Only

---

# Design Philosophy

Notifications should communicate useful information without distracting the user.

Toasts must never require user interaction unless an action is explicitly provided.

---

# Display Location

Toast notifications appear directly above the Status Bar.

They never cover

- Command Input
- Keyboard
- System Navigation Area

---

# Display Order

```
Conversation Area

↓

Toast

↓

Status Bar

↓

Command Input

↓

Android Keyboard
```

---

# Toast Lifecycle

```
Created

↓

Visible

↓

Dismissed
```

Alternative

```
Created

↓

Action Selected

↓

Dismissed
```

---

# Toast Types

Supported

- Success
- Information
- Warning
- Error
- Permission
- Network
- System

---

# Success Toast

Purpose

Confirms successful completion.

Example

```text
Workspace indexed successfully.
```

---

# Information Toast

Purpose

Provides neutral information.

Example

```text
Theme updated.
```

---

# Warning Toast

Purpose

Warns about recoverable situations.

Example

```text
Low storage space.
```

---

# Error Toast

Purpose

Reports failures.

Example

```text
Unable to read file.
```

---

# Permission Toast

Purpose

Requests user attention.

Example

```text
Permission required.
```

---

# Network Toast

Purpose

Reports connectivity changes.

Example

```text
Network unavailable.
```

---

# System Toast

Purpose

Displays internal application events.

Example

```text
Workspace refreshed.
```

---

# Layout

```text
────────────────────────────────

Workspace indexed successfully.

────────────────────────────────
```

---

# Components

Each Toast may contain

- Notification Type
- Message
- Optional Action
- Optional Timeout

---

# Message Length

Recommended

One short sentence.

Maximum

Two lines.

---

# Actions

Optional.

Examples

```text
Retry
```

```text
Undo
```

```text
Open
```

---

# Timeout

Default

```
3 Seconds
```

Longer duration allowed for warnings and errors.

---

# Manual Dismiss

Supported.

Users may dismiss a toast immediately.

---

# Queue

Multiple Toasts

Displayed sequentially.

Never overlap.

---

# Duplicate Handling

Duplicate messages

Merge into one notification when possible.

---

# Priority

Highest

- Error
- Permission
- Warning
- Success
- Information

Higher priority replaces lower priority when necessary.

---

# Animation

Allowed

Simple terminal-friendly transition.

Examples

```
Fade In

Fade Out
```

No decorative effects.

---

# Prohibited Animation

Never use

- Bounce
- Rotate
- Scale
- Flash
- Shake

---

# Streaming Integration

Streaming continues while Toasts are visible.

Toast rendering never interrupts AI responses.

---

# Tool Integration

Examples

```text
Git Commit Complete.
```

```text
Workspace Indexed.
```

```text
Package Installed.
```

---

# Keyboard Behavior

When the Android keyboard is open

Toast moves upward automatically.

It never overlaps the keyboard.

---

# Safe Area

The Toast System respects

- Display Cutout
- Gesture Navigation
- Keyboard
- Bottom Insets

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Toast messages must always be readable.

---

# Performance

Render only active Toast.

Avoid unnecessary updates.

---

# Error Handling

If the Toast system fails

The application continues normally.

Notifications may fall back to the Conversation Area.

---

# Security

Never display

- Passwords
- Tokens
- API Keys
- Sensitive file paths

Only user-safe information is shown.

---

# Restrictions

Never

- Block user input
- Cover the Command Input
- Cover the Status Bar
- Cover the Keyboard
- Stay visible indefinitely
- Display duplicate notifications repeatedly

---

# Example Success

```text
Workspace indexed successfully.
```

---

# Example Error

```text
Unable to connect to MCP server.
```

---

# Example Warning

```text
Workspace contains unsaved changes.
```

---

# Example Action

```text
File deleted.

Undo
```

---

# Example Queue

```text
Theme updated.

↓

Workspace refreshed.

↓

Git status updated.
```

---

# Toast Checklist

Every Toast must

- Display concise information
- Auto-dismiss
- Respect safe areas
- Support optional actions
- Support queueing
- Support priorities
- Remain readable
- Never block typing
- Never interrupt streaming
- Remain terminal-native

---

# Core Rules

1. Toasts are non-blocking.
2. Toasts appear above the Status Bar.
3. Auto-dismiss after a short duration.
4. Queue multiple notifications.
5. Higher priority replaces lower priority when necessary.
6. Never overlap the keyboard.
7. Never expose sensitive information.
8. Keep messages concise.
9. Preserve streaming performance.
10. Optimize for Android Termux.

---

# Summary

The Toast System provides lightweight, terminal-native notifications that keep users informed without interrupting their workflow. Whether reporting successful actions, warnings, errors, permission requests, or system events, Toasts remain concise, accessible, and non-blocking while respecting mobile safe areas and maintaining a responsive Android Termux experience.