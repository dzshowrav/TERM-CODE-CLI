# 14-streaming-response.md

# Streaming Response
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Streaming Response system for the Mobile AI CLI.

Streaming allows AI responses to appear incrementally as they are generated, rather than waiting for the complete response.

The system must provide immediate feedback, maintain smooth rendering, and preserve a responsive user experience on Android Termux.

---

# Design Goals

The Streaming Response system must be

- Mobile First
- Terminal Native
- Real-Time
- Low Latency
- Smooth
- Non-Blocking
- Incremental
- Performance Optimized

---

# Design Philosophy

Users should see progress immediately.

The interface should feel alive.

Streaming must never block typing, scrolling, or tool execution.

---

# Response Lifecycle

```
User Prompt

↓

AI Processing

↓

Thinking

↓

Streaming Starts

↓

Streaming Continues

↓

Streaming Complete

↓

Idle
```

---

# Response States

Supported

```
Queued

Thinking

Streaming

Paused

Resumed

Completed

Cancelled

Failed
```

Only one state may be active.

---

# Queued

Displayed immediately after prompt submission.

Example

```text
Queued...
```

---

# Thinking

Displayed while the AI prepares a response.

Example

```text
Thinking...
```

This state disappears automatically when streaming begins.

---

# Streaming

Characters, words, or lines appear incrementally.

Rendering begins immediately after the first token is received.

---

# Paused

If streaming is temporarily interrupted

Display

```text
Streaming Paused
```

---

# Resumed

If streaming continues

Display

```text
Streaming Resumed
```

---

# Completed

Display

```text
Response Complete
```

The cursor returns to the Command Input.

---

# Cancelled

Display

```text
Response Cancelled
```

Partially generated content remains visible.

---

# Failed

Display

```text
Streaming Failed
```

Provide a readable explanation.

---

# Rendering Strategy

Streaming updates only newly received content.

Previously rendered text must never be redrawn unnecessarily.

---

# Rendering Order

```
Receive Token

↓

Append Text

↓

Render Updated Rows

↓

Scroll If Needed
```

---

# Token Handling

Tokens are appended sequentially.

Never reorder received content.

---

# Character Rendering

Characters appear in the exact order received.

Never skip characters.

Never duplicate characters.

---

# Markdown Streaming

Markdown is rendered progressively.

Supported

- Headings
- Lists
- Tables
- Quotes
- Code Blocks
- Links

Incomplete markdown remains readable until completed.

---

# Code Block Streaming

Code formatting is preserved during streaming.

Example

```text
function hello() {
```

The code block continues growing until complete.

---

# Table Streaming

Rows appear incrementally.

Columns remain aligned.

If alignment is temporarily incomplete, update only affected rows.

---

# List Streaming

List items appear as they are generated.

Existing items never move unexpectedly.

---

# Tool Output Integration

Streaming may continue while tool output is displayed.

Tool blocks and AI responses remain visually distinct.

---

# Conversation Behavior

Conversation continues to scroll normally.

Users may read older messages while streaming continues.

---

# Auto Scroll

Automatically scroll only if

The user is already at the bottom.

If the user scrolls upward

Disable automatic scrolling.

---

# Manual Scroll

Supported during streaming.

Streaming continues without interruption.

---

# Typing During Streaming

Users may type a new prompt while the current response streams.

The draft remains independent.

---

# Input Behavior

Command Input remains fully interactive.

Never disable typing during streaming.

---

# Status Bar Integration

Possible Status values

```
Thinking

Streaming

Ready
```

Updates immediately.

---

# Progress Indicator

Optional.

Display

- Token Count
- Elapsed Time
- Current State

Avoid unnecessary visual noise.

---

# Cursor Behavior

Streaming cursor appears only at the end of the active response.

It disappears when streaming completes.

---

# Selection

Users may select previously streamed text.

Selection must not interrupt rendering.

---

# Copy

Users may copy completed or partially streamed text.

Formatting is preserved.

---

# Pause Handling

If the connection pauses

Keep partial content visible.

Do not erase generated text.

---

# Resume Handling

Continue appending from the last received token.

Never restart the response automatically.

---

# Error Recovery

If streaming fails

Keep partial response.

Offer retry when appropriate.

---

# Long Responses

Very long responses continue streaming naturally.

Conversation scrolls efficiently.

Only visible rows are rendered.

---

# Performance

Render only changed rows.

Avoid full-screen redraws.

Target latency

```
Less than 50 ms
```

---

# Memory Management

Large conversations

Render visible content first.

Keep older content cached.

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Streaming updates must remain readable.

---

# Safe Area

Streaming always occurs inside the Conversation Area.

Never overlap

- Command Input
- Status Bar
- Keyboard

---

# Keyboard Behavior

Keyboard opening

Reduces Conversation height.

Streaming continues without interruption.

---

# Responsive Behavior

Terminal resize

```
Pause Layout

↓

Recalculate Grid

↓

Continue Streaming
```

No content is lost.

---

# Security

Never stream

- Sensitive system information
- Hidden prompts
- Internal debugging data

Only user-visible content is streamed.

---

# Restrictions

Never

- Freeze the UI
- Restart streaming automatically
- Duplicate text
- Reorder tokens
- Hide partial responses
- Block user input

---

# Example Layout

```text
Assistant

Thinking...

Creating a React application...

Installing dependencies...

Project created successfully.
```

---

# Streaming Example

```text
Assistant

The application architecture consists of

Frontend

Backend

Database

Authentication

Deployment
```

Each line appears progressively.

---

# Streaming Checklist

Every streamed response must

- Begin immediately
- Preserve order
- Preserve formatting
- Preserve markdown
- Preserve code blocks
- Support scrolling
- Support typing
- Respect safe areas
- Recover from interruptions
- Finish cleanly

---

# Core Rules

1. Stream responses incrementally.
2. Never block typing.
3. Never redraw unchanged content.
4. Preserve formatting during streaming.
5. Auto-scroll only when appropriate.
6. Keep partial responses after interruptions.
7. Resume from the last received token.
8. Render efficiently.
9. Always update the Status Bar.
10. Optimize for Android Termux.

---

# Summary

The Streaming Response system delivers AI output in real time, allowing users to read responses as they are generated. By combining incremental rendering, efficient updates, uninterrupted typing, and terminal-native behavior, the Mobile AI CLI provides a fast, responsive, and transparent AI experience optimized for Android Termux and mobile-first development workflows.