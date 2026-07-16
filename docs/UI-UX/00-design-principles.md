# 00-design-principles.md

# Mobile AI CLI UI/UX Design Principles
## Version 1.0

---

# Purpose

This document defines the global UI/UX principles for the Mobile AI CLI.

Every screen, component, interaction, animation, layout, and rendering rule must follow this specification.

This document has the highest priority among all UI/UX specifications.

---

# Design Goals

The interface must be

- Mobile First
- Terminal Native
- Touch Friendly
- Keyboard Safe
- Minimal
- Fast
- Readable
- Reactive
- Consistent
- Accessible

The experience should feel like a professional AI coding terminal rather than a traditional mobile application.

---

# Target Platform

Primary platform

- Android
- Termux

Supported

- Android 10+
- Portrait Mode

Not supported

- Desktop-first layouts
- Landscape-first layouts
- Tablet-first layouts

---

# Design Philosophy

The interface follows these principles.

## Mobile First

Everything is designed for a phone before considering any larger display.

Every component must fit comfortably within one-hand operation.

---

## Terminal Native

The application behaves like a real terminal.

Users should always feel they are interacting with a CLI rather than a GUI application.

---

## AI First

The AI conversation is always the primary content.

Everything else supports the conversation.

Nothing should distract from it.

---

## Content First

Information has higher priority than decoration.

Avoid unnecessary visual elements.

---

## Minimal Interface

Display only what is necessary.

Hide advanced controls until needed.

---

## Fast Interaction

Every action should require as few touches as possible.

---

## Predictable Layout

The interface must never unexpectedly change position.

The user should always know where everything is.

---

# Screen Orientation

Supported

```
Portrait
```

Not Supported

```
Landscape
```

---

# Screen Hierarchy

```
Conversation

↓

Input

↓

Status Bar
```

Nothing may appear below the Status Bar.

---

# Sidebar Policy

Sidebars are prohibited.

Never use

- Left Sidebar
- Right Sidebar
- Floating Sidebar
- Collapsible Sidebar

Navigation must happen inside the main screen.

---

# Bottom First Layout

Every important interaction happens near the bottom.

Reason

- Thumb accessibility
- Mobile ergonomics
- Faster interaction

---

# Keyboard First Design

The software must assume

```
Keyboard

↓

Always Opens
```

Every layout must adapt automatically.

Nothing important may become hidden behind the keyboard.

---

# Keyboard Safe Area

When keyboard opens

```
Conversation

↓

Resizes

↓

Input

↓

Status Bar
```

The input must never move off-screen.

---

# Status Bar Position

The Status Bar always remains below the Input.

Layout

```
Conversation

↓

Input

↓

Status Bar
```

Never

```
Status

↓

Input
```

---

# Navigation Principles

Navigation must remain simple.

Preferred

- Tap
- Long Press
- Swipe
- Keyboard

Avoid

- Nested menus
- Multi-level navigation
- Floating windows

---

# Visual Style

Use

- Clean spacing
- Sharp borders
- Simple layout

Avoid

- Heavy decoration
- Large shadows
- Glass effects
- Blur

---

# Color Philosophy

Colors communicate state.

Examples

```
Primary

Success

Warning

Error

Info

Neutral
```

Never use random colors.

---

# Typography

Text must be

- Highly readable
- Monospaced where appropriate
- Consistent

Avoid decorative fonts.

---

# Icon Policy

Use only

- Terminal icons
- Nerd Font icons
- Codicons
- Powerline icons

Do not

- Replace icons with text
- Use emoji
- Use decorative icons

---

# Emoji Policy

Emoji are prohibited.

Never display

- Faces
- Objects
- Symbols represented as emoji

Only terminal-compatible icons are allowed.

---

# Component Principles

Every component must

- Have one purpose
- Be reusable
- Be lightweight
- Support keyboard
- Support touch

---

# Layout Rules

Always

- Align to grid
- Respect spacing
- Maintain hierarchy

Never

- Overlap components
- Hide controls
- Break alignment

---

# Safe Area Rules

Respect

- Screen edges
- Keyboard
- Bottom gesture area
- Notches

No component should touch screen edges directly.

---

# Terminal Grid

Every screen follows

```
Rows

Columns

Spacing
```

The grid remains consistent across the application.

---

# Animation Principles

Animations must

- Be fast
- Be subtle
- Never block interaction

Avoid

- Long animations
- Bounce effects
- Excessive motion

---

# Loading Philosophy

Loading must always communicate progress.

Allowed

- Spinner
- Progress Bar
- Skeleton

Never freeze the interface.

---

# Error Philosophy

Errors must

- Explain the problem
- Suggest recovery
- Keep the application usable

Avoid

- Technical dumps
- Unclear messages

---

# Empty States

Empty screens should explain

- Why empty
- What to do next

Never leave blank screens.

---

# Touch Targets

Minimum touch size

```
44 × 44 dp
```

Preferred

```
48 × 48 dp
```

---

# Thumb Zone

Frequently used controls belong in the lower portion of the screen.

Avoid placing primary actions near the top.

---

# Accessibility

Support

- Screen readers
- Large fonts
- High contrast
- Reduced motion

---

# Performance Goals

Target

```
60 FPS
```

Input latency

```
Less than 16 ms
```

Screen updates

```
Incremental
```

---

# Memory Usage

Avoid unnecessary allocations.

Reuse

- Components
- Buffers
- Layout objects

---

# Responsiveness

UI must react immediately to

- Keyboard
- Orientation changes
- Terminal resize
- Streaming output

---

# Consistency Rules

Every screen must share

- Same spacing
- Same typography
- Same icon system
- Same interaction model
- Same navigation pattern

---

# Conversation Priority

Highest priority

```
Conversation
```

Second

```
Input
```

Third

```
Status
```

Everything else is secondary.

---

# Rendering Rules

Render only changed content whenever possible.

Avoid full-screen redraws.

---

# Offline First

The interface must remain usable without internet access.

Network state should never block navigation.

---

# Security Principles

Never expose

- Secrets
- Tokens
- Credentials

Sensitive information must remain hidden unless explicitly requested.

---

# Design Restrictions

Never use

- Sidebars
- Floating action buttons
- Pop-up advertisements
- Emoji
- Desktop layouts
- Hidden navigation
- Multi-column layouts
- Mouse-dependent interactions

---

# Core UX Principles

1. Mobile First
2. Terminal Native
3. AI First
4. Content First
5. Keyboard Safe
6. Bottom Focused
7. Touch Friendly
8. Minimal
9. Consistent
10. Fast
11. Accessible
12. Predictable
13. Offline First
14. Incremental Rendering
15. One-Hand Operation

---

# Master Layout

```
┌──────────────────────────────┐
│                              │
│                              │
│                              │
│                              │
│        Conversation          │
│                              │
│                              │
│                              │
├──────────────────────────────┤
│ > Command Input              │
├──────────────────────────────┤
│ Model │ Tokens │ Session     │
└──────────────────────────────┘
```

---

# Summary

This document establishes the foundation for every screen and interaction in the Mobile AI CLI.

Every future UI/UX specification must follow these principles to ensure a consistent, terminal-native, mobile-first experience optimized for Android Termux, one-handed operation, and AI-assisted coding workflows.