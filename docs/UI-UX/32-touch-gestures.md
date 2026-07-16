# 32-touch-gestures.md

# Touch Gestures
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Touch Gesture System used throughout the Mobile AI CLI.

The Touch Gesture System enables natural mobile interaction for terminal-based workflows, allowing users to navigate, select, scroll, edit, and control the application without depending only on keyboard input.

The system is designed for Android Termux and mobile-first CLI usage.

---

# Design Goals

The Touch Gesture System must be

- Mobile First
- Touch Friendly
- Terminal Native
- Precise
- Predictable
- Lightweight
- Keyboard Compatible
- Accessibility Friendly

---

# Supported Platform

Primary

- Android
- Termux

Input Methods

Supported

- Touch Screen
- Stylus
- Trackpad
- Mouse

---

# Design Philosophy

Touch interaction should enhance CLI workflows without replacing keyboard efficiency.

Every gesture must have a clear purpose.

---

# Gesture Categories

Supported

- Tap
- Double Tap
- Long Press
- Swipe
- Scroll
- Drag
- Pinch (Optional)
- Multi Touch (Optional)

---

# Gesture Priority

The system prioritizes

1. Accuracy
2. Predictability
3. Speed
4. Accessibility

---

# Tap

## Purpose

Primary selection action.

---

# Tap Behavior

Used for

- Select item
- Open item
- Activate button
- Focus input

---

# Examples

Command Palette

```
Tap Command

↓

Execute
```

File Picker

```
Tap File

↓

Select File
```

---

# Double Tap

## Purpose

Quick open action.

---

# Supported Areas

- File Picker
- Tree Renderer
- Workspace Browser

---

# Example

```
Double Tap Folder

↓

Open Folder
```

---

# Long Press

## Purpose

Display additional actions.

---

# Long Press Actions

Examples

- Copy
- Rename
- Delete
- Details
- Preview

---

# Example

```
Hold File

↓

Context Actions
```

---

# Swipe

## Purpose

Navigate and dismiss.

---

# Horizontal Swipe

Possible uses

- Close Modal
- Switch Panels
- Navigate History

---

# Vertical Swipe

Used for

- Scroll Content
- Navigate Lists

---

# Scroll

## Purpose

Move through large content.

---

# Scroll Areas

Supported

- Conversation
- File List
- Command Results
- Documentation
- Logs

---

# Scroll Behavior

Scrolling must be

- Smooth
- Incremental
- Position Preserving

---

# Auto Scroll

Used for

- AI Streaming
- Logs

Only when user is already near bottom.

---

# Manual Scroll Protection

If user scrolls upward

Disable automatic scrolling.

---

# Drag

## Purpose

Move or adjust elements.

---

# Supported Uses

Optional

- Resize Panels
- Reorder Items
- Select Text

---

# Text Selection

Supported gestures

- Long Press
- Drag Handles

---

# Copy Gesture

Flow

```
Select Text

↓

Copy

↓

Clipboard
```

---

# Paste Gesture

Supported

```
Long Press Input

↓

Paste
```

---

# Modal Gestures

Dialogs support

- Tap Outside (Optional)
- Swipe Down Close (Optional)

---

# Bottom Area Gestures

The bottom area contains

- Command Input
- Status Bar

Gestures must not accidentally trigger system navigation.

---

# Keyboard Compatibility

Touch gestures must work together with keyboard input.

Examples

```
Touch Select

+

Keyboard Type
```

---

# File Picker Gestures

Supported

```
Tap

Select File
```

```
Double Tap

Open Folder
```

```
Long Press

File Actions
```

---

# Tree Renderer Gestures

Supported

```
Tap Arrow

Expand Node
```

```
Tap Node

Select
```

```
Long Press

Actions
```

---

# Command Palette Gestures

Supported

```
Tap Result

Execute
```

```
Swipe

Browse Results
```

---

# Model Selector Gestures

Supported

```
Tap Model

Switch Model
```

```
Long Press

View Details
```

---

# Chat Screen Gestures

Supported

```
Scroll Up

View History
```

```
Scroll Down

Return Latest
```

---

# Code Block Gestures

Supported

- Scroll
- Select Text
- Copy Content

---

# Touch Target Size

Interactive elements must provide enough space.

Minimum recommended

```
44px equivalent
```

---

# Touch Feedback

Selection must provide feedback.

Examples

- Highlight
- Focus indicator
- State change

Never rely only on color.

---

# Gesture Conflicts

When gestures overlap

Priority:

```
Active Component

↓

Application Gesture

↓

System Gesture
```

---

# Accidental Touch Prevention

Avoid accidental actions.

Important actions require confirmation.

Examples

- Delete
- Reset
- Remove Workspace

---

# Accessibility

Support

- Screen Readers
- Touch Assistance
- Large Touch Targets

---

# Performance

Gesture handling must be

- Low latency
- Lightweight
- Non-blocking

---

# Safe Area

Gestures respect

- Display Cutout
- Keyboard Area
- Navigation Area
- Status Bar

---

# Error Handling

Invalid gesture

Ignore safely.

Never crash the interface.

---

# Security

Dangerous actions triggered by gestures require confirmation.

---

# Restrictions

Never

- Use hidden gestures
- Require complex gestures
- Trigger destructive actions accidentally
- Conflict with Android navigation
- Replace keyboard workflows completely

---

# Example Mobile Workflow

```
Tap Input

↓

Keyboard Opens

↓

Type Prompt

↓

Swipe Conversation

↓

Review Response
```

---

# Example File Workflow

```
Long Press File

↓

Context Menu

↓

Copy Path
```

---

# Touch Gesture Checklist

Every Touch System must

- Support basic gestures
- Provide feedback
- Avoid accidental actions
- Respect keyboard behavior
- Respect safe areas
- Support accessibility
- Maintain performance
- Work with terminal UI
- Optimize Android usage
- Remain predictable

---

# Core Rules

1. Every gesture has a clear purpose.
2. Tap is the primary action.
3. Long press reveals secondary actions.
4. Scrolling must remain smooth.
5. Destructive actions require confirmation.
6. Gestures never block typing.
7. Touch targets must be usable.
8. System gestures must be respected.
9. Accessibility is required.
10. Optimize for Android Termux.

---

# Summary

The Touch Gesture System transforms the Mobile AI CLI into a practical mobile-first terminal experience. Through simple interactions such as tapping, scrolling, long pressing, and selecting, users can efficiently manage files, commands, models, and conversations while maintaining the power and flexibility of a CLI environment on Android Termux.