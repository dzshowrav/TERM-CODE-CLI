# 11-command-input.md

# Command Input
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

The Command Input is the primary interaction component of the Mobile AI CLI.

Every prompt, command, slash command, file reference, AI instruction, and chat message originates from this component.

The Command Input must remain permanently accessible throughout the user's session.

---

# Design Goals

The Command Input must be

- Mobile First
- Terminal Native
- Keyboard First
- Thumb Friendly
- One-Hand Optimized
- Fast
- Predictable
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

The input area is the command center of the application.

It should behave like a modern AI coding terminal rather than a messaging application.

Typing must always feel immediate.

---

# Position

Always located above the Status Bar.

```
Conversation

↓

Command Input

↓

Status Bar
```

The input never changes position.

---

# Layout

```
┌────────────────────────────────────┐
│ >                                  │
├────────────────────────────────────┤
│ Model │ Tokens │ Session │ Ready   │
└────────────────────────────────────┘
```

---

# Input Area

Contains

- Prompt
- Cursor
- Inline Completion
- Attachment Indicator
- Slash Commands
- Mention System

---

# Prompt Symbol

Always display a prompt prefix.

Example

```
>
```

or

```
$
```

The prefix remains fixed.

---

# Input Width

```
100%
```

Use the entire available width.

---

# Input Height

Minimum

```
1 Line
```

Maximum

```
8 Lines
```

When the maximum height is reached

Conversation continues to shrink.

---

# Multiline Support

Supported

User may type

- Long prompts
- Markdown
- Code
- JSON
- YAML
- SQL

---

# Cursor

Always visible.

Never hidden behind the keyboard.

Cursor occupies one terminal cell.

---

# Cursor Behavior

Blinking

Optional

Movement

Immediate

Never animate cursor movement.

---

# Placeholder

Displayed only when input is empty.

Example

```
Ask anything...
```

Placeholder disappears immediately after typing begins.

---

# Typing Behavior

Every key press

Immediately updates the input.

No delayed rendering.

---

# Input Expansion

As text grows

```
Input Height

↑

Conversation Height

↓
```

Status Bar never moves.

---

# Keyboard Behavior

Keyboard Closed

```
Conversation

↓

Input

↓

Status
```

Keyboard Open

```
Conversation Reduced

↓

Input

↓

Status

↓

Keyboard
```

The input is never hidden.

---

# Input Focus

Application startup

↓

Input receives focus.

After sending

↓

Input receives focus again.

---

# Sending Messages

Default

Enter

Alternative

Send Action

Input clears only after successful submission.

---

# New Line

Supported

When multiline mode is enabled.

Never insert unintended line breaks.

---

# Slash Commands

Supported.

Example

```
/help

/model

/theme

/clear

/session

/workspace

/search
```

Suggestions appear while typing.

---

# Mention Support

Supported.

Examples

```
@workspace

@file

@folder

@git

@terminal
```

Autocomplete is available.

---

# File References

Supported.

Examples

```
src/App.tsx

README.md

package.json
```

Autocomplete when possible.

---

# Path Completion

Supported.

Suggestions appear as the user types.

---

# Command History

Supported.

Previous prompts remain accessible.

History persists across the session.

---

# History Navigation

Supported via

- Keyboard
- Touch

History never overwrites the current draft without confirmation.

---

# Draft Recovery

Unsent text is preserved automatically.

After reopening the application

Draft remains available.

---

# Input Validation

Validate

- Empty prompt
- Invalid command
- Unsupported syntax

Validation never blocks typing.

---

# Auto Complete

Supported for

- Commands
- Paths
- Files
- Models
- Tools
- Slash Commands

Suggestions update while typing.

---

# Inline Suggestions

Displayed inside the input.

User may

Accept

Ignore

Dismiss

Suggestions never overwrite user text automatically.

---

# Paste

Supported.

Large pasted content should remain responsive.

---

# Copy

Supported.

Formatting preserved.

---

# Undo

Supported.

Redo

Supported.

---

# Selection

Supported.

Selection remains inside the input.

---

# Scrolling

When multiline exceeds visible height

Internal scrolling begins.

Conversation remains fixed.

---

# Input States

Supported

```
Idle

Focused

Typing

Streaming

Disabled

Error
```

---

# Disabled State

Input temporarily disabled only when absolutely required.

Conversation remains readable.

---

# Streaming State

While AI streams

User may continue typing.

The next prompt is prepared independently.

---

# Error State

If sending fails

Input content remains.

Never delete unsent text.

---

# Attachment Support

Supported.

Files

Folders

Images

Logs

Multiple attachments may be queued.

---

# Drag and Drop

Not required.

Termux primarily relies on file picker and paste.

---

# Character Limit

No practical limit.

Large prompts should scroll smoothly.

---

# Performance

Typing latency target

```
Less than 16 ms
```

Rendering should update only changed characters.

---

# Safe Area

Input always remains inside the protected area.

Never overlap

- Keyboard
- Navigation Gesture Area

---

# Accessibility

Support

- Large fonts
- Screen readers
- High contrast
- Reduced motion

Focus should always be announced correctly.

---

# Restrictions

Never

- Hide the input
- Overlay the input
- Move the input above the conversation
- Replace the input with a dialog
- Clear text unexpectedly
- Block typing while AI streams

---

# Layout Example

```
┌────────────────────────────────────┐
│ Assistant                          │
│                                    │
│ Project created successfully.      │
│                                    │
│ User                               │
│                                    │
│ Create a React project.            │
│                                    │
├────────────────────────────────────┤
│ > npm create vite@latest           │
├────────────────────────────────────┤
│ GPT-5 │ Main │ 2.8K │ Ready        │
└────────────────────────────────────┘
```

---

# Input Checklist

Every Command Input must

- Stay visible
- Stay keyboard safe
- Support multiline
- Support slash commands
- Support mentions
- Support autocomplete
- Preserve drafts
- Preserve failed prompts
- Expand smoothly
- Remain responsive

---

# Core Rules

1. The Command Input is always visible.
2. It remains above the Status Bar.
3. Input expands vertically up to eight lines.
4. Conversation shrinks before the input moves.
5. Keyboard never hides the input.
6. Drafts are automatically preserved.
7. AI streaming never blocks typing.
8. Autocomplete is non-destructive.
9. Input updates immediately on every keystroke.
10. The Command Input is optimized for one-handed Android Termux use.

---

# Summary

The Command Input is the central interaction surface of the Mobile AI CLI. It is designed as a persistent, terminal-native command interface that supports natural language prompts, slash commands, code, file references, and AI workflows. Its fixed bottom position, keyboard-safe behavior, multiline support, and responsive rendering provide a reliable mobile-first experience for AI-assisted development on Android Termux.