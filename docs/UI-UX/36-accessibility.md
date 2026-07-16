# 36-accessibility.md

# Accessibility System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Accessibility System used throughout the Mobile AI CLI.

The Accessibility System ensures that the CLI interface remains usable for users with different visual, motor, hearing, and interaction needs while maintaining a terminal-native experience.

The system must provide equal access to all features without reducing performance or functionality.

---

# Design Goals

The Accessibility System must be

- Inclusive
- Mobile First
- Terminal Native
- Keyboard Friendly
- Screen Reader Compatible
- High Contrast Supported
- Customizable
- Performance Optimized

---

# Supported Platform

Primary

- Android
- Termux

Accessibility Support

- Android Accessibility Services
- Screen Readers
- Large Text Settings
- Alternative Input Methods

---

# Design Philosophy

Accessibility is not an additional feature.

It is a core part of the interface architecture.

Every component must remain understandable through:

- Visual feedback
- Text information
- Keyboard interaction
- Assistive technology

---

# Accessibility Architecture

```
User Interaction

↓

Accessibility Layer

↓

Component Semantics

↓

Terminal Renderer

↓

Output
```

---

# Accessibility Principles

## Perceivable

Information must be visible and understandable.

---

## Operable

Users must be able to control the application.

---

## Understandable

Actions and states must be clear.

---

## Robust

Features must work with assistive technologies.

---

# Text Accessibility

All important information must have readable text.

Examples:

Good

```text
File loading failed.
```

Bad

```text
!
```

---

# Color Accessibility

The interface must never depend only on color.

Example

Bad

```text
Red = Error
```

Good

```text
Error: Permission denied
```

---

# High Contrast Mode

Supported.

Improves visibility by increasing separation between

- Background
- Text
- Active Elements
- Status Indicators

---

# High Contrast Example

Normal

```text
Warning
```

High Contrast

```text
WARNING: Permission Required
```

---

# Font Accessibility

Supported settings

- Font Size
- Text Density
- Line Height

---

# Large Text Mode

The interface must adapt when text size increases.

Requirements

- No clipped text
- No broken layouts
- Scroll support

---

# Screen Reader Support

The system should expose meaningful information.

Examples

Component:

```text
Model Selector
```

Screen Reader:

```text
Current model GPT-5. Double tap to change.
```

---

# Semantic Labels

Interactive elements require labels.

Examples

Button

```text
Send Message
```

Input

```text
Command Input
```

---

# Focus Management

Focus must be predictable.

Supported

- Keyboard Focus
- Touch Focus
- Screen Reader Focus

---

# Focus Order

Recommended order

```
Header

↓

Content

↓

Input

↓

Status Bar

↓

Actions
```

---

# Keyboard Accessibility

All important actions must work without touch.

Supported

- Navigation
- Selection
- Commands
- Dialog Actions

---

# Keyboard Navigation

Supported keys

```
Arrow Keys

Tab

Enter

Escape
```

---

# Touch Accessibility

Touch targets must be large enough.

Recommended minimum

```
44px equivalent
```

---

# Gesture Accessibility

Avoid requiring complex gestures.

Required actions must support alternatives.

Example

Instead of:

```
Swipe Only
```

Provide:

```
Swipe

or

Keyboard Shortcut
```

---

# Motion Accessibility

Support reduced motion mode.

---

# Reduced Motion

When enabled

Disable:

- Decorative animation
- Fast transitions
- Continuous movement

Keep:

- Status changes
- Progress information

---

# Animation Accessibility

Animations must never be required to understand state.

Example

Bad

```
Only spinner shows completion
```

Good

```
Completed successfully
```

---

# Error Accessibility

Errors must include text descriptions.

Example

Bad

```
!
```

Good

```
Error: File not found
```

---

# Loading Accessibility

Loading states must provide text.

Example

Bad

```
Spinner only
```

Good

```
Loading workspace...
```

---

# Empty State Accessibility

Empty states must explain

- Current situation
- Available action

Example

```text
No Workspace

Select a folder to start.
```

---

# Status Bar Accessibility

Status information must remain readable.

Contains

- Current Model
- Connection State
- Context Usage
- System Status

---

# Command Input Accessibility

The command input must support

- Cursor visibility
- Text navigation
- Copy/Paste
- Voice Input

---

# Voice Input Support

Compatible with mobile voice typing.

Requirements

- Preserve formatting
- Handle long input
- Maintain cursor position

---

# Color Blind Support

Use additional indicators.

Example

Instead of:

```text
Green
```

Use:

```text
Success: Completed
```

---

# Icon Accessibility

Icons require meaning.

Rules:

- Icons must have labels
- Icons cannot replace important text
- Emoji are not used as icons

---

# Code Block Accessibility

Code rendering must support

- Text selection
- Copy
- Scrolling
- Clear contrast

---

# Markdown Accessibility

Markdown renderer must preserve

- Heading hierarchy
- Lists
- Tables
- Code structure

---

# Table Accessibility

Tables should provide

- Clear headers
- Readable columns
- Horizontal scrolling when needed

---

# File Tree Accessibility

Tree navigation must expose

- File name
- Folder state
- Selection state

Example

```text
src folder expanded
```

---

# Dialog Accessibility

Dialogs must provide

- Title
- Description
- Available actions

Example

```text
Delete Session

This action cannot be undone.

Cancel

Delete
```

---

# Permission Accessibility

Permission requests must explain why access is needed.

Example

```text
File access is required to read project files.
```

---

# Offline Accessibility

Offline states must explain limitations.

Example

```text
Offline mode enabled.

Local features available.
```

---

# Accessibility Settings

Available options

```
High Contrast

Large Text

Reduced Motion

Screen Reader Support

Keyboard Navigation
```

---

# Privacy

Accessibility features must not expose

- Private files
- API keys
- Tokens
- Hidden data

---

# Performance

Accessibility features must remain lightweight.

Avoid:

- Heavy processing
- Delayed responses
- Blocking interaction

---

# Testing Requirements

Every feature should be tested with

- Keyboard only
- Screen reader
- Large text
- High contrast
- Reduced motion

---

# Restrictions

Never

- Depend only on color
- Hide important information inside icons
- Require touch-only actions
- Use animation as the only feedback
- Break layouts with large text
- Ignore keyboard navigation

---

# Accessibility Checklist

Every Accessibility System must

- Support text descriptions
- Support keyboard control
- Support screen readers
- Support high contrast
- Support reduced motion
- Support large text
- Preserve usability
- Avoid color dependency
- Respect privacy
- Work on Android Termux

---

# Core Rules

1. Every action must have understandable feedback.
2. Text is the primary information source.
3. Color is never the only indicator.
4. Keyboard navigation is required.
5. Touch alternatives should exist.
6. Animations must support reduced motion.
7. Screen readers must receive useful information.
8. Layout must adapt to user settings.
9. Privacy must always be protected.
10. Accessibility is part of the core design.

---

# Summary

The Accessibility System ensures that the Mobile AI CLI remains usable, understandable, and powerful for all users. Through semantic information, keyboard support, screen reader compatibility, high contrast modes, reduced motion support, and adaptable layouts, the CLI delivers an inclusive terminal experience optimized for Android Termux.