# 24-dialog-system.md

# Dialog System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Dialog System used throughout the Mobile AI CLI.

Dialogs provide focused interaction windows for important decisions, confirmations, permissions, configuration changes, destructive actions, and detailed information.

The Dialog System must remain terminal-native while providing a mobile-friendly interaction model optimized for Android Termux.

---

# Design Goals

The Dialog System must be

- Mobile First
- Terminal Native
- Focused
- Accessible
- Keyboard Safe
- Touch Friendly
- Non-Blocking
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

Dialogs should appear only when user attention is required.

They must provide enough context for a decision without creating unnecessary interruptions.

---

# Display Location

Dialogs appear as centered modal overlays above the Conversation Area.

Background content remains visible but inactive.

---

# Layout

```text
┌──────────────────────────────┐
│ Confirm Action                │
├──────────────────────────────┤
│ Delete selected file?         │
│                               │
│ This action cannot be undone. │
├──────────────────────────────┤
│ Cancel              Confirm   │
└──────────────────────────────┘
```

---

# Dialog Types

Supported

- Confirmation Dialog
- Input Dialog
- Selection Dialog
- Permission Dialog
- Information Dialog
- Error Dialog
- Progress Dialog

---

# Confirmation Dialog

Purpose

Used for actions requiring approval.

Examples

- Delete file
- Reset settings
- Remove workspace
- Execute dangerous command

---

# Example

```text
Delete File?

README.md will be removed.

Cancel        Delete
```

---

# Input Dialog

Purpose

Collect user input.

Examples

- Rename file
- Enter API key
- Create session name

---

# Example

```text
Session Name

[ mobile-cli ]

Cancel        Save
```

---

# Selection Dialog

Purpose

Choose one option from multiple choices.

Examples

- Select model
- Select theme
- Select workspace

---

# Example

```text
Select Model

GPT-5

Claude

Gemini
```

---

# Permission Dialog

Purpose

Request authorization.

Examples

- File access
- Tool execution
- Network access

---

# Example

```text
Permission Required

Allow terminal execution?

Deny          Allow
```

---

# Information Dialog

Purpose

Display important information.

Examples

- About application
- Version details
- Help information

---

# Example

```text
Mobile AI CLI

Version 1.0

OK
```

---

# Error Dialog

Purpose

Display critical failures.

---

# Example

```text
Operation Failed

Unable to connect.

Close
```

---

# Progress Dialog

Purpose

Display unavoidable blocking operations.

Only used when background execution is impossible.

---

# Example

```text
Installing Package

███████░░░ 70%

Cancel
```

---

# Components

Each Dialog contains

- Title
- Content
- Actions
- Optional Input
- Optional Description

---

# Title

Must clearly describe the purpose.

Good

```text
Delete File
```

Bad

```text
Warning
```

---

# Content

Should explain

- What happened
- What will happen
- What user needs to decide

---

# Actions

Actions appear at the bottom.

Examples

```
Cancel

Confirm
```

---

# Action Priority

Primary action

Displayed last/right side.

Secondary action

Displayed first/left side.

---

# Destructive Actions

Destructive actions require clear labeling.

Example

```text
Delete Permanently
```

Never use vague labels.

---

# Dialog Width

Adaptive.

Based on terminal width.

Must remain readable on small screens.

---

# Dialog Height

Minimum required height.

Avoid unnecessary empty space.

---

# Scrolling

Large dialog content may scroll.

Buttons remain accessible.

---

# Keyboard Behavior

When keyboard opens

Dialog moves above keyboard.

Input fields remain visible.

---

# Input Handling

Supported

- Text Input
- Password Input
- Search Input

---

# Focus Management

When opened

First interactive element receives focus.

---

# Keyboard Navigation

Supported

- Move Between Actions
- Confirm
- Cancel

---

# Touch Navigation

Supported

- Tap Actions
- Scroll Content
- Select Options

---

# Cancellation

Supported methods

- Back Action
- Escape Key
- Cancel Button

---

# Confirmation

Requires explicit user action.

No automatic confirmation.

---

# Nested Dialogs

Avoid multiple dialogs.

Maximum

```
One Active Dialog
```

---

# Streaming Integration

Dialogs may appear while AI streaming continues.

Background streaming remains paused visually but not necessarily stopped.

---

# Tool Integration

Examples

```text
Allow Shell Command?

npm install
```

```text
Allow File Modification?

src/App.tsx
```

---

# Safe Area

Dialogs respect

- Display Cutout
- Keyboard
- Status Bar
- Command Input

Never overlap protected areas.

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Every action must have readable labels.

---

# Performance

Dialog rendering should be lightweight.

Avoid unnecessary redraws.

---

# Error Handling

Invalid dialog state

Fallback to simple notification.

Never crash the application.

---

# Security

Sensitive inputs must use protected fields.

Never display

- Passwords
- Tokens
- Private keys

---

# Restrictions

Never

- Open multiple dialogs simultaneously
- Block the entire application unnecessarily
- Hide important actions
- Use unclear button labels
- Confirm destructive actions automatically
- Cover the keyboard

---

# Example Confirmation Flow

```text
User

Delete project

↓

Dialog

Delete Project?

Cancel       Delete

↓

User confirms

↓

Toast

Project deleted.
```

---

# Example Permission Flow

```text
AI

Need to run command

↓

Dialog

Allow?

npm install

Deny        Allow
```

---

# Example Input Flow

```text
Rename File

[ new-name.ts ]

Cancel        Save
```

---

# Dialog Checklist

Every Dialog must

- Have clear purpose
- Require intentional action
- Support keyboard input
- Support touch input
- Respect safe areas
- Handle cancellation
- Remain accessible
- Avoid unnecessary interruption
- Protect sensitive information
- Stay terminal-native

---

# Core Rules

1. Dialogs appear only when needed.
2. User decisions must be explicit.
3. One dialog at a time.
4. Actions must be clearly labeled.
5. Destructive actions require confirmation.
6. Keyboard must never hide inputs.
7. Dialogs must respect mobile safe areas.
8. Never expose sensitive information.
9. Keep dialogs small and focused.
10. Optimize for Android Termux.

---

# Summary

The Dialog System provides focused, terminal-native interactions for important decisions and user input inside the Mobile AI CLI. Through confirmation dialogs, permission requests, selections, and input forms, it creates a safe and predictable experience while maintaining mobile-first behavior, keyboard compatibility, and Android Termux optimization.