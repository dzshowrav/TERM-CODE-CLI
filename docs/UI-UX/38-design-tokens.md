# 38-design-tokens.md

# Design Tokens
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Design Token System used throughout the Mobile AI CLI.

Design Tokens provide a centralized system for controlling visual, interaction, spacing, typography, animation, and component behavior.

The system ensures consistency across all terminal UI components while supporting themes, customization, and future scalability.

---

# Design Goals

The Design Token System must be

- Consistent
- Theme Driven
- Scalable
- Maintainable
- Terminal Native
- Mobile First
- Performance Optimized

---

# Supported Platform

Primary

- Android
- Termux

Rendering Environment

- Terminal UI
- ANSI Compatible Output

---

# Design Philosophy

Design tokens act as the single source of truth.

Components should consume tokens instead of defining values directly.

---

# Token Architecture

```
Design Tokens

↓

Theme Resolver

↓

Component System

↓

Renderer

↓

Terminal Output
```

---

# Token Categories

The system contains:

- Color Tokens
- Typography Tokens
- Spacing Tokens
- Layout Tokens
- Border Tokens
- Icon Tokens
- Animation Tokens
- State Tokens
- Accessibility Tokens

---

# Token Naming Convention

Format:

```
category-name-purpose
```

Example:

```
color-background-primary

spacing-md

text-size-small
```

---

# Color Tokens

Controls all visual colors.

---

# Background Tokens

```
background-primary

background-secondary

background-surface

background-overlay
```

---

# Text Tokens

```
text-primary

text-secondary

text-muted

text-disabled
```

---

# Accent Tokens

```
accent-primary

accent-secondary
```

---

# Status Tokens

Success:

```
status-success
```

Warning:

```
status-warning
```

Error:

```
status-error
```

Info:

```
status-info
```

---

# Selection Tokens

Used for focused elements.

```
selection-background

selection-text

selection-border
```

---

# Code Highlight Tokens

Used for syntax rendering.

```
code-keyword

code-function

code-string

code-number

code-comment
```

---

# Typography Tokens

Controls text appearance.

---

# Font Family Tokens

```
font-terminal

font-code

font-ui
```

---

# Font Size Tokens

```
text-xs

text-sm

text-md

text-lg

text-xl
```

---

# Font Weight Tokens

```
weight-normal

weight-medium

weight-bold
```

---

# Line Height Tokens

```
line-tight

line-normal

line-relaxed
```

---

# Spacing Tokens

Controls distance between elements.

---

# Base Spacing Scale

Example:

```
spacing-0

spacing-1

spacing-2

spacing-3

spacing-4

spacing-5

spacing-6
```

---

# Component Spacing

Examples:

```
chat-padding

input-padding

dialog-padding

panel-gap
```

---

# Layout Tokens

Controls structure.

---

# Screen Tokens

```
screen-padding

screen-margin

screen-height
```

---

# Component Size Tokens

```
input-height

button-height

icon-size
```

---

# Terminal Grid Tokens

Controls terminal layout.

```
grid-column

grid-row

grid-gap
```

---

# Border Tokens

Controls separators.

---

# Border Style

```
border-solid

border-dashed

border-none
```

---

# Border Width

```
border-thin

border-normal

border-heavy
```

---

# Border Characters

Terminal compatible.

Examples:

```
│
─
┌
┐
└
┘
```

---

# Radius Tokens

For compatible renderers.

```
radius-small

radius-medium

radius-large
```

---

# Icon Tokens

Controls icon behavior.

---

# Icon Size

```
icon-xs

icon-sm

icon-md

icon-lg
```

---

# Icon Spacing

```
icon-gap-small

icon-gap-medium
```

---

# Icon State

```
icon-active

icon-disabled

icon-selected
```

---

# Animation Tokens

Controls movement.

---

# Duration Tokens

```
animation-fast

animation-normal

animation-slow
```

---

# Transition Tokens

```
transition-open

transition-close

transition-update
```

---

# Loading Tokens

```
spinner-speed

progress-update-rate
```

---

# State Tokens

Controls component states.

---

# Interactive States

```
state-default

state-hover

state-focus

state-active

state-disabled
```

---

# System States

```
state-loading

state-success

state-error

state-warning
```

---

# Accessibility Tokens

Controls accessibility preferences.

---

# Contrast Tokens

```
contrast-normal

contrast-high
```

---

# Motion Tokens

```
motion-enabled

motion-reduced

motion-disabled
```

---

# Text Accessibility Tokens

```
text-scale-normal

text-scale-large
```

---

# Theme Integration

All tokens must support themes.

Example:

Dark Theme:

```
background-primary

↓

#000000
```

Light Theme:

```
background-primary

↓

#FFFFFF
```

---

# Component Usage

Components never define raw values.

Example:

Wrong:

```text
background: black
```

Correct:

```text
background: color-background-primary
```

---

# Token Resolution

Flow:

```
Component Request

↓

Token Name

↓

Theme Value

↓

Renderer Output
```

---

# Runtime Theme Switching

When theme changes:

```
New Theme

↓

Update Tokens

↓

Refresh Components
```

---

# Custom Token Support

Users may create custom themes.

Example:

```json
{
 "color-background-primary": "#000000",
 "accent-primary": "#00FFFF"
}
```

---

# Token Validation

Every token set must validate:

- Required tokens exist
- Values are readable
- No invalid references exist

---

# Performance Rules

Tokens should be:

- Cached
- Resolved once
- Reused

Avoid:

- Recalculating values frequently

---

# Security

Theme and token files must never execute code.

Only configuration values are allowed.

---

# Restrictions

Never:

- Hardcode visual values inside components
- Create inconsistent spacing
- Ignore theme tokens
- Use unsupported terminal colors
- Depend only on color meaning

---

# Design Token Checklist

Every Design Token System must:

- Centralize design values
- Support themes
- Support accessibility
- Support customization
- Work in terminal UI
- Improve consistency
- Reduce duplication
- Support mobile layouts
- Remain lightweight
- Work on Android Termux

---

# Core Rules

1. Tokens are the single source of truth.
2. Components consume tokens only.
3. Themes modify token values.
4. Accessibility uses token overrides.
5. Hardcoded styles are prohibited.
6. Tokens must be reusable.
7. Naming must remain consistent.
8. Terminal compatibility is required.
9. Performance must be considered.
10. Design consistency is mandatory.

---

# Summary

The Design Token System creates the foundation of the Mobile AI CLI design architecture. By centralizing colors, typography, spacing, layout, states, and accessibility values, the system enables consistent UI development, dynamic themes, easier maintenance, and scalable terminal-native experiences optimized for Android Termux.