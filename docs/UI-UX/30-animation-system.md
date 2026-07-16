# 30-animation-system.md

# Animation System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Animation System used throughout the Mobile AI CLI.

The Animation System controls transitions, loading effects, state changes, streaming indicators, and visual feedback while maintaining terminal performance and a distraction-free coding environment.

The system must provide subtle, meaningful animations inspired by modern AI CLI experiences while remaining fully optimized for Android Termux.

---

# Design Goals

The Animation System must be

- Mobile First
- Terminal Native
- Lightweight
- Purpose Driven
- Performance Optimized
- Accessible
- Battery Friendly

---

# Supported Platform

Primary

- Android
- Termux

Orientation

- Portrait Only

---

# Design Philosophy

Animation exists to explain change.

It must never exist only for decoration.

Every animation should answer one of these questions:

- What changed?
- What is happening?
- What should the user notice?

---

# Animation Architecture

```
UI Event

↓

Animation Controller

↓

Transition State

↓

Terminal Renderer

↓

Display Update
```

---

# Animation Types

Supported

- Transition Animation
- Loading Animation
- Streaming Animation
- Selection Animation
- Progress Animation
- Notification Animation
- State Change Animation

---

# Animation Principles

## Meaningful

Animation must communicate information.

---

## Fast

Animations should complete quickly.

---

## Smooth

Avoid sudden visual jumps.

---

## Minimal

Avoid unnecessary movement.

---

# Default Timing

Recommended durations

```
Instant

0ms - 50ms
```

```
Fast

50ms - 150ms
```

```
Normal

150ms - 300ms
```

```
Slow

300ms - 500ms
```

---

# Transition Animation

Used for

- Opening dialogs
- Closing dialogs
- Switching screens
- Changing panels

---

# Screen Transition

Example

```
Chat

↓

Settings
```

The transition should feel immediate.

---

# Modal Animation

Example

```
Dialog Opening

↓

Dialog Visible
```

Recommended

Fast fade transition.

---

# Loading Animation

Used when waiting for

- AI response
- Tool execution
- File operation
- Network request

---

# Spinner Animation

Supported styles

```
|

/

-

\
```

or

```
⠋
⠙
⠹
⠸
```

---

# Spinner Rules

Spinner must

- Use low CPU
- Remain readable
- Stop when complete

---

# Streaming Animation

Used during AI response generation.

Example

```text
Generating response...
```

↓

```text
Generating response.
```

↓

```text
Generating response..
```

---

# Cursor Animation

For active input.

Example

```text
_
```

or

```text
█
```

---

# Selection Animation

Used for

- Command Palette
- File Picker
- Model Selector

Example

```text
> GPT-5
```

Selected state must remain visible.

---

# Progress Animation

Used with progress indicators.

Example

```text
████░░░░
```

Updates smoothly.

---

# Toast Animation

Used for

- Success
- Warning
- Error

Recommended

```
Appear

↓

Visible

↓

Disappear
```

---

# No Blocking Animation

Animations must never block

- Typing
- Scrolling
- Commands
- Tool execution

---

# Keyboard Behavior

When keyboard opens

Animations continue normally.

No layout reset.

---

# Safe Area

Animations must respect

- Command Input
- Status Bar
- Keyboard
- Display boundaries

---

# Streaming Integration

Animations and AI streaming must work together.

Example

```
Thinking Indicator

+

Generated Text Streaming
```

---

# Performance Rules

Animation updates should

- Use minimal redraws
- Avoid full screen refresh
- Pause when inactive

---

# Battery Optimization

On mobile devices

Reduce animation frequency when

- Battery saver enabled
- Device performance is low

---

# Accessibility

Support

- Reduced Motion Mode
- High Contrast
- Static Alternatives

---

# Reduced Motion Mode

When enabled

Animations become instant.

Example

Before

```
Fade In
```

After

```
Immediate Display
```

---

# Animation Settings

Available in Settings.

Options

```
Full Animation

Reduced Animation

Disabled
```

---

# Full Animation

Default mode.

All supported animations enabled.

---

# Reduced Animation

Only essential feedback remains.

---

# Disabled Animation

No movement.

Only state changes are displayed.

---

# Terminal Compatibility

Animations must work with

- ANSI terminals
- Termux
- Mobile terminal emulators

---

# Unsupported Animation

Never use

- Large screen movement
- Flashing effects
- Continuous unnecessary motion
- Heavy graphical effects

---

# Error Handling

Animation failure must not affect functionality.

Fallback

```
Static UI State
```

---

# Security

Animation data is visual only.

Never execute external animation content.

---

# Example Loading State

```text
AI Thinking

⠋ Processing
```

---

# Example Streaming State

```text
Generating response...

The application architecture
is being created...
```

---

# Example Transition

```text
Chat

↓

Settings
```

---

# Example Progress

```text
Indexing Workspace

████████░░ 80%
```

---

# Animation Checklist

Every Animation System must

- Improve understanding
- Remain lightweight
- Support reduced motion
- Avoid blocking interaction
- Work in terminal
- Respect safe areas
- Optimize CPU usage
- Optimize battery usage
- Support Android Termux
- Remain accessible

---

# Core Rules

1. Animation explains change.
2. Animation never delays user actions.
3. Keep animations short.
4. Avoid unnecessary movement.
5. Support reduced motion.
6. Never consume excessive CPU.
7. Streaming remains smooth.
8. Terminal compatibility is required.
9. Accessibility comes first.
10. Optimize for mobile environments.

---

# Summary

The Animation System provides subtle visual feedback throughout the Mobile AI CLI. By focusing on meaningful transitions, efficient loading indicators, streaming feedback, and accessibility support, it enhances usability without sacrificing terminal performance, battery efficiency, or Android Termux compatibility.