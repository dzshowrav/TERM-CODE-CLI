# 22-progress-ui.md

# Progress UI
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Progress UI system used throughout the Mobile AI CLI.

The Progress UI communicates the status of long-running operations such as AI generation, file indexing, repository cloning, package installation, workspace scanning, MCP synchronization, downloads, uploads, and tool execution.

The interface must provide clear feedback without interrupting the user's workflow.

---

# Design Goals

The Progress UI must be

- Mobile First
- Terminal Native
- Real-Time
- Lightweight
- Non-Blocking
- Streaming Friendly
- Accessible
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

Users should always understand

- What is happening
- How much is completed
- What remains
- Whether interaction is still possible

Progress indicators must be informative rather than decorative.

---

# Display Location

Progress indicators appear inside the Conversation Area.

Long-running operations may also update the Status Bar.

The Chat Screen always remains visible.

---

# Progress Lifecycle

```
Queued

↓

Preparing

↓

Running

↓

Completing

↓

Completed
```

Alternative paths

```
Running

↓

Cancelled
```

or

```
Running

↓

Failed
```

---

# Progress States

Supported

```
Queued

Preparing

Running

Paused

Resuming

Completed

Cancelled

Failed

Timed Out
```

Only one state is active.

---

# Basic Layout

```text
Indexing Workspace

████████░░░░░░░░ 45%

214 / 480 Files
```

---

# Components

A Progress UI may contain

- Title
- Current State
- Progress Bar
- Percentage
- Current Task
- Item Count
- Elapsed Time
- Remaining Time (optional)

---

# Title

Describes the operation.

Examples

```text
Installing Packages

Searching Workspace

Generating Response

Downloading Model
```

---

# Current State

Examples

```text
Preparing

Running

Paused

Completed
```

---

# Progress Bar

Recommended style

```text
████████░░░░░░░░
```

Alternative

```text
########--------
```

Bar width adapts to terminal size.

---

# Percentage

Displayed when measurable.

Example

```text
72%
```

---

# Item Count

Example

```text
214 / 480 Files
```

---

# Current Task

Example

```text
Scanning src/components
```

Updates dynamically.

---

# Elapsed Time

Example

```text
00:01:42
```

---

# Remaining Time

Optional.

Example

```text
Estimated

00:00:18
```

Hide if estimation is unreliable.

---

# Indeterminate Progress

For unknown duration

Display

```text
Working...
```

with an animated indicator.

No percentage shown.

---

# Spinner

Supported styles

```text
|

/

-

\
```

or

```text
⠁

⠂

⠄

⠂
```

Animation must remain subtle.

---

# Multiple Operations

Supported.

Display sequentially in the Conversation Area.

Status Bar shows only the highest-priority task.

---

# Nested Progress

Example

```text
Workspace Index

↓

Search

↓

File Read
```

Nested operations remain visually distinct.

---

# Streaming Integration

Progress updates occur while AI responses continue streaming.

Neither blocks the other.

---

# Tool Integration

Examples

- Git Clone
- npm Install
- Workspace Index
- MCP Sync
- Search
- Build
- Test

Each tool updates its own progress independently.

---

# Status Bar Integration

Possible values

```text
Running

Indexing

Downloading

Searching

Ready
```

---

# Keyboard Behavior

Users may continue typing during progress updates.

The Command Input remains fully interactive.

---

# Auto Scroll

Automatically scroll only if the user is already at the bottom.

Otherwise, preserve the current scroll position.

---

# Safe Area

Progress indicators remain inside the Conversation Area.

Never overlap

- Command Input
- Status Bar
- Keyboard

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Progress must always include readable text, not only graphics.

---

# Performance

Update only changed values.

Avoid redrawing the full progress block.

---

# Completion

Display

```text
Completed

480 Files Indexed
```

Progress bar remains visible briefly before becoming part of the conversation history.

---

# Failure

Display

```text
Failed

Permission denied.
```

Provide a clear explanation.

---

# Cancellation

Display

```text
Cancelled by User
```

Partial progress remains visible.

---

# Timeout

Display

```text
Operation Timed Out
```

Retry may be offered.

---

# Security

Never expose

- Hidden system paths
- Internal prompts
- Sensitive credentials

Only user-relevant progress information is displayed.

---

# Restrictions

Never

- Freeze the interface
- Hide progress unexpectedly
- Remove partial progress after failure
- Display misleading percentages
- Block typing
- Block scrolling

---

# Example File Indexing

```text
Workspace Index

██████████░░░░░░ 63%

315 / 500 Files

Scanning components
```

---

# Example Download

```text
Downloading Model

██████████████░░ 88%

642 MB / 730 MB
```

---

# Example Package Installation

```text
Installing Packages

Working...

Resolving dependencies
```

---

# Example Completion

```text
Completed

500 Files Indexed
```

---

# Progress UI Checklist

Every Progress UI must

- Display the current state
- Update incrementally
- Support determinate progress
- Support indeterminate progress
- Respect safe areas
- Remain responsive
- Support streaming
- Preserve history
- Report failures clearly
- Stay terminal-native

---

# Core Rules

1. Progress updates are always visible.
2. Users may continue typing.
3. Display percentages only when meaningful.
4. Use indeterminate indicators when necessary.
5. Preserve partial progress after interruptions.
6. Update incrementally.
7. Never freeze the Chat Screen.
8. Respect mobile safe areas.
9. Keep progress understandable.
10. Optimize for Android Termux.

---

# Summary

The Progress UI provides a consistent, terminal-native way to communicate the status of long-running operations within the Mobile AI CLI. By combining progress bars, textual status, incremental updates, and non-blocking interaction, it keeps users informed without interrupting coding or conversation. The design is optimized for Android Termux, ensuring responsive performance and a mobile-first user experience.