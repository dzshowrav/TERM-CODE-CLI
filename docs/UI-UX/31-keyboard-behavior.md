# 31-keyboard-behavior.md

# Keyboard Behavior
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Keyboard Behavior System used throughout the Mobile AI CLI.

The Keyboard Behavior System controls how the application responds to the Android virtual keyboard, physical keyboards, terminal input methods, shortcuts, focus changes, and text composition.

The system is designed specifically for mobile-first CLI usage where the keyboard is always considered an active part of the interface.

---

# Design Goals

The Keyboard Behavior System must be

- Mobile First
- Keyboard Aware
- Input Focused
- Terminal Native
- Fast
- Predictable
- Non-Blocking
- Accessible

---

# Supported Platform

Primary

- Android
- Termux

Input Devices

Supported

- Android Virtual Keyboard
- Physical Keyboard
- Bluetooth Keyboard
- Terminal Keyboard Input

---

# Design Philosophy

The keyboard is not an external accessory.

The keyboard is a primary navigation and interaction layer.

The entire UI must adapt around keyboard visibility.

---

# Keyboard States

Supported states

```
Hidden

Opening

Visible

Closing

Disconnected
```

---

# Keyboard Lifecycle

```
User Focuses Input

↓

Keyboard Opens

↓

Layout Adjusts

↓

User Types

↓

Keyboard Closes

↓

Layout Restores
```

---

# Main Rule

The Command Input area always remains connected with the keyboard.

---

# Layout Behavior

When keyboard is closed

```
┌─────────────────────┐
│ Conversation        │
│                     │
│                     │
├─────────────────────┤
│ Status Bar          │
├─────────────────────┤
│ Command Input       │
└─────────────────────┘
```

---

When keyboard is open

```
┌─────────────────────┐
│ Conversation        │
│                     │
├─────────────────────┤
│ Status Bar          │
├─────────────────────┤
│ Command Input       │
├─────────────────────┤
│ Android Keyboard    │
└─────────────────────┘
```

---

# Status Bar Position

The Status Bar always remains below the input area.

Required order

```
Conversation Area

↓

Command Input

↓

Status Bar

↓

Keyboard
```

---

# Input Focus

When application starts

Focus behavior:

Default

```
Command Input Ready
```

---

# Auto Focus Rules

Auto focus enabled for

- New Chat
- Command Execution
- Search Input

Disabled for

- Settings Screen
- Information Dialog

---

# Text Input Behavior

Supported

- Normal typing
- Multiline input
- Paste
- Copy
- Delete
- Selection
- Cursor movement

---

# Multiline Input

Supported.

Example

```text
Explain this architecture

with complete details
```

---

# Input Expansion

Command Input expands vertically when needed.

Maximum height must be limited.

---

# Long Message Input

For large text

Support

- Internal scrolling
- Line navigation

---

# Cursor Behavior

Cursor must remain visible.

When text scrolls

Cursor position is preserved.

---

# Selection Behavior

Supported

- Select Text
- Copy
- Cut
- Paste

---

# Paste Handling

Large pasted content should

- Avoid UI freeze
- Process incrementally
- Preserve formatting

---

# Enter Key Behavior

Default

```
Enter

↓

Send Message
```

---

# Multiline Enter

Supported shortcut

```
Shift + Enter

↓

New Line
```

---

# Command Mode

Commands start with

```
/
```

Example

```text
/model GPT-5
```

---

# Keyboard Shortcuts

Supported examples

```
Ctrl + C

Cancel Operation
```

```
Ctrl + L

Clear Input
```

```
Ctrl + P

Command Palette
```

```
Ctrl + K

Search
```

---

# Escape Key

Used for

- Close Dialog
- Cancel Command
- Clear Selection

---

# Arrow Navigation

Supported for

- Command History
- Menus
- File Picker
- Model Selector

---

# Command History

The input remembers previous commands.

Navigation

```
Arrow Up

Previous Command
```

```
Arrow Down

Next Command
```

---

# Input History Storage

Stores

- Previous commands
- Search queries
- Tool commands

Sensitive content must be protected.

---

# Keyboard With Dialogs

When a dialog opens

The keyboard behavior depends on the dialog type.

---

# Input Dialog

Keyboard opens automatically.

Example

```
Rename File

[filename]

Keyboard Active
```

---

# Confirmation Dialog

Keyboard closes.

Focus moves to actions.

---

# Search Interface

Search inputs receive immediate focus.

Examples

- Command Palette
- File Picker
- Model Selector

---

# Keyboard With Streaming

AI streaming continues while keyboard is open.

User can continue typing.

---

# Keyboard With Tools

Tool execution does not hide the keyboard unless required.

---

# Keyboard With File Picker

Search input remains accessible.

File list height adjusts.

---

# Keyboard With Settings

Input fields remain visible.

Screen scrolls when required.

---

# Safe Area Handling

Keyboard system respects

- Display Cutout
- Navigation Bar
- Input Area
- Status Bar

---

# Scroll Behavior

When keyboard opens

The application scrolls only when necessary.

Never force unwanted scrolling.

---

# Draft Preservation

If keyboard closes

Current input draft remains.

---

# Application Switching

When returning from another app

Restore

- Keyboard state
- Input text
- Cursor position

---

# Orientation Handling

Portrait optimized.

Keyboard rotation changes must not break layout.

---

# Accessibility

Support

- Voice Input
- Screen Readers
- Large Text
- Alternative Keyboards

---

# Performance

Keyboard events must be processed instantly.

Avoid

- Full UI redraw
- Input lag
- Cursor delay

---

# Error Handling

Input failure

Display

```text
Input unavailable.
```

---

# Security

Keyboard input may contain sensitive data.

Protect

- Passwords
- API Keys
- Tokens

Never store sensitive input history.

---

# Restrictions

Never

- Hide Command Input
- Cover input with keyboard
- Lose typed text
- Reset cursor unexpectedly
- Block typing during AI generation
- Move Status Bar above Input

---

# Example Mobile Flow

```
Open App

↓

Keyboard Opens

↓

Type Prompt

↓

Send

↓

AI Streams Response

↓

Continue Typing
```

---

# Example Command Flow

```
/

↓

Command Palette

↓

Select Command

↓

Execute
```

---

# Keyboard Checklist

Every Keyboard System must

- Detect keyboard state
- Preserve input
- Maintain focus
- Support shortcuts
- Support history
- Respect safe areas
- Work with Android keyboard
- Prevent input loss
- Stay responsive
- Optimize for Termux

---

# Core Rules

1. Keyboard is a primary interaction layer.
2. Command Input always remains accessible.
3. Status Bar stays below input.
4. Keyboard never covers important UI.
5. User text is never lost.
6. Streaming does not block typing.
7. Focus behavior must be predictable.
8. Shortcuts improve navigation.
9. Sensitive input remains protected.
10. Optimize for Android mobile usage.

---

# Summary

The Keyboard Behavior System ensures that the Mobile AI CLI feels natural on Android devices. By treating the keyboard as a core UI component, maintaining input stability, supporting shortcuts, preserving drafts, and adapting layouts dynamically, the CLI provides a reliable mobile coding experience optimized for Termux environments.