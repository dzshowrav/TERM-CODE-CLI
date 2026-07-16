# 10-chat-screen.md

# Chat Screen
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

The Chat Screen is the primary workspace of the Mobile AI CLI.

Everything the user does begins and ends on this screen.

Unlike traditional applications, the Chat Screen is never replaced by another page. Dialogs, overlays, and temporary interfaces appear above it and always return to it.

This screen must remain visible throughout the entire session.

---

# Design Goals

The Chat Screen must be

- Mobile First
- Terminal Native
- Keyboard Safe
- One-Hand Friendly
- Streaming Optimized
- Touch Friendly
- Fast
- Minimal

---

# Screen Hierarchy

```
Conversation Area

↓

Input Area

↓

Status Bar
```

Nothing may exist below the Status Bar.

---

# Screen Layout

```
┌────────────────────────────────────┐
│                                    │
│ Conversation                       │
│                                    │
│ Assistant                          │
│                                    │
│ User                               │
│                                    │
│ Assistant                          │
│                                    │
│                                    │
├────────────────────────────────────┤
│ >                                  │
├────────────────────────────────────┤
│ Model │ Session │ Tokens │ Ready   │
└────────────────────────────────────┘
```

---

# Conversation Area

Purpose

Display

- User messages
- Assistant messages
- Tool output
- Markdown
- Code blocks
- Search results
- File previews

Conversation occupies every available row above the Input Area.

---

# Conversation Rules

Always

- Scroll vertically
- Preserve message order
- Preserve formatting
- Preserve markdown
- Preserve code formatting

Never

- Reverse message order
- Overlay messages
- Re-render unchanged messages

---

# Message Flow

```
User

↓

Assistant Thinking

↓

Assistant Streaming

↓

Assistant Complete
```

---

# Message Order

Always chronological.

Oldest

↓

Newest

Newest message always appears at the bottom.

---

# Message Width

Maximum

```
100%
```

Messages never appear inside bubbles.

Terminal style only.

---

# Message Alignment

All messages

```
Left Aligned
```

Never center conversation.

---

# User Message

Display

```
User

Message
```

No decorative container.

---

# Assistant Message

Display

```
Assistant

Response
```

Streaming supported.

---

# System Message

Purpose

Warnings

Updates

Notifications

Displayed using muted colors.

---

# Tool Message

Purpose

Show tool execution.

Examples

```
Reading Files

Searching

Editing

Writing

Completed
```

Tool messages remain inside the conversation.

---

# Streaming Message

While generating

Display continuously.

Never wait for completion before rendering.

---

# Thinking State

Display

```
Thinking...
```

or equivalent status.

Thinking disappears automatically when generation begins.

---

# Code Blocks

Must preserve

- Indentation
- Spacing
- Blank lines
- Syntax highlighting

Horizontal scrolling is allowed.

---

# Markdown

Supported

- Headings
- Lists
- Tables
- Quotes
- Links
- Code
- Horizontal Rules

Markdown renders inline.

---

# Tables

Must remain aligned.

If wider than viewport

Enable horizontal scrolling.

---

# Tree Views

Preserve indentation.

Example

```
workspace

    src

        components

        pages

    package.json
```

---

# Search Results

Display

Filename

↓

Matched Line

↓

Preview

↓

Path

---

# File Preview

Supported

- Text
- Markdown
- JSON
- Source Code

Large files

Preview only.

---

# Conversation Scrolling

Conversation

Scrollable.

Input

Fixed.

Status

Fixed.

---

# Auto Scroll

Automatically scroll only if

User is already at the bottom.

If user scrolls upward

Do not force scrolling.

---

# Manual Scroll

User may inspect older messages without interruption.

Streaming continues.

---

# Message Selection

Supported

- Long Press
- Keyboard Selection

Selection must preserve formatting.

---

# Copy

Supported

- Message
- Code Block
- Selection

Formatting preserved.

---

# Retry

Assistant responses may support retry.

Retry appears only when appropriate.

---

# Edit Prompt

User messages may be edited.

Editing creates a new conversation branch.

Original remains available.

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

---

# Input Visibility

Input remains visible at all times.

Never hidden.

---

# Status Visibility

Status Bar always visible.

---

# Loading

During startup

Conversation displays

```
Loading Session...
```

---

# Empty State

Display

```
No Conversation

Type a prompt below.
```

---

# Error State

Display

```
Operation Failed

Retry Available
```

Conversation remains visible.

---

# Notification

Temporary messages appear

Above Input.

Never cover

Conversation

Input

Status

---

# Focus

Default

Input

Only one component receives focus.

---

# Touch

Supported

- Tap
- Long Press
- Scroll

Unsupported

- Swipe Navigation
- Multi-touch Navigation

---

# Safe Area

Conversation respects

- Display Cutout
- Keyboard
- Gesture Navigation

---

# Rendering

Render incrementally.

Never redraw the full conversation unnecessarily.

---

# Performance

Target

```
60 FPS
```

Streaming latency

```
Minimal
```

Only changed rows are rendered.

---

# Accessibility

Support

- Large fonts
- High contrast
- Screen readers
- Reduced motion

---

# Restrictions

Never use

- Chat bubbles
- Floating buttons
- Sidebars
- Bottom navigation
- Tabs
- Avatar images
- Decorative backgrounds

---

# Screen States

Supported

```
Empty

Loading

Conversation

Streaming

Tool Execution

Error
```

State transitions must never replace the Chat Screen.

---

# Layout Example

```
┌────────────────────────────────────┐
│ Assistant                          │
│                                    │
│ Welcome to the workspace.          │
│                                    │
│ User                               │
│                                    │
│ Create a React project.            │
│                                    │
│ Assistant                          │
│                                    │
│ Running command...                 │
│ npm create vite@latest             │
│                                    │
├────────────────────────────────────┤
│ >                                  │
├────────────────────────────────────┤
│ GPT-5 │ Main │ 2.6K │ Ready        │
└────────────────────────────────────┘
```

---

# Chat Screen Checklist

Every Chat Screen must

- Keep Conversation visible
- Keep Input fixed
- Keep Status fixed
- Support streaming
- Preserve code formatting
- Preserve markdown
- Support incremental rendering
- Respect Safe Area
- Support one-hand usage
- Remain keyboard safe

---

# Core Rules

1. The Chat Screen is the permanent workspace.
2. Conversation is always the primary content.
3. Messages are never displayed in bubbles.
4. Input remains fixed at the bottom.
5. Status Bar remains below the Input.
6. Streaming updates incrementally.
7. Only the Conversation scrolls.
8. Overlays never replace the Chat Screen.
9. Code formatting is preserved exactly.
10. The Chat Screen is always optimized for Android Termux.

---

# Summary

The Chat Screen is the heart of the Mobile AI CLI. It provides a persistent, terminal-native workspace where conversations, code generation, tool execution, and AI interaction occur in a single continuous view. The design emphasizes readability, streaming performance, one-handed mobile interaction, and complete compatibility with Android Termux while avoiding desktop-oriented navigation patterns.