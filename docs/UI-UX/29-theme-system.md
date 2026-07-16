# 29-theme-system.md

# Theme System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Theme System used throughout the Mobile AI CLI.

The Theme System controls colors, visual hierarchy, terminal appearance, contrast levels, syntax presentation, and user customization while maintaining a professional terminal-native experience.

The system must provide a consistent visual language inspired by modern AI CLI tools while being optimized for Android Termux.

---

# Design Goals

The Theme System must be

- Mobile First
- Terminal Native
- Highly Customizable
- Accessible
- Lightweight
- Consistent
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

A theme is not only a color palette.

It defines

- Visual hierarchy
- Information priority
- Reading comfort
- Interaction clarity

Every UI element must receive its appearance from the theme engine.

---

# Theme Architecture

```
Theme Configuration

↓

Token Resolver

↓

Component Styles

↓

Terminal Renderer

↓

UI Output
```

---

# Theme Tokens

All visual values are controlled through tokens.

Example

```text
background

foreground

primary

secondary

success

warning

error

info
```

---

# Color System

## Background Colors

Used for

- Main screen
- Modal areas
- Panels

Example

```text
background
surface
overlay
```

---

# Foreground Colors

Used for

- Text
- Labels
- Descriptions

Example

```text
text-primary

text-secondary

text-muted
```

---

# Accent Colors

Used for

- Active items
- Selection
- Important actions

Example

```text
primary

accent
```

---

# Status Colors

## Success

Used for

- Completed actions
- Successful operations

---

## Warning

Used for

- Attention required
- Recoverable issues

---

## Error

Used for

- Failures
- Invalid actions

---

## Info

Used for

- General information

---

# Default Themes

The system supports built-in themes.

---

# Dark Theme

Primary default theme.

Characteristics

- Low brightness
- High readability
- Terminal friendly

Example

```text
Background

Black / Dark Gray

Text

Light Gray
```

---

# Light Theme

Optional.

Characteristics

- Bright background
- Dark text

---

# High Contrast Theme

Designed for accessibility.

Characteristics

- Strong separation
- Maximum readability

---

# Terminal Classic Theme

Inspired by traditional terminals.

Characteristics

- Monochrome
- Minimal colors

---

# Custom Themes

Users may create custom themes.

Supported customization

- Colors
- Contrast
- Text brightness
- Accent colors

---

# Theme File Format

Example

```json
{
  "name": "custom-dark",
  "background": "#000000",
  "foreground": "#ffffff",
  "primary": "#00ffff"
}
```

---

# Component Theming

Every component consumes theme tokens.

Examples

- Chat Message
- Command Input
- Status Bar
- Toast
- Dialog
- Progress UI
- File Picker

---

# Chat Theme

Controls

- User messages
- AI messages
- Code blocks
- Markdown

---

# Input Theme

Controls

- Cursor
- Placeholder
- Active state

---

# Status Bar Theme

Controls

- Model indicator
- Context usage
- System state

---

# Selection Theme

Controls

- Highlighted items
- Focused elements

---

# Code Theme

Controls

- Syntax highlighting
- Keywords
- Strings
- Comments

---

# Syntax Highlighting

Uses separate syntax tokens.

Example

```text
keyword

function

string

number

comment
```

---

# Icon Theme

Controls

- Icon appearance
- Icon brightness
- Icon alignment

---

# Icon Rules

Icons must be

- Professional
- Simple
- Terminal compatible
- Text-free

Never replace icons with emoji.

---

# Theme Switching

Supported methods

- Settings Screen
- Command Palette
- Configuration File

---

# Switching Flow

```
User Selects Theme

↓

Load Theme

↓

Validate Tokens

↓

Apply Globally

↓

Save Preference
```

---

# Live Preview

Optional.

Users can preview themes before applying.

---

# Theme Persistence

Selected theme is stored.

Restored after application restart.

---

# Auto Theme

Optional.

Can follow

- System Theme
- Time Based Theme

---

# Theme Validation

Every theme must validate

- Required tokens exist
- Colors are readable
- Contrast is acceptable

---

# Accessibility

Support

- High Contrast
- Color independent indicators
- Large fonts

Never rely only on color.

---

# Keyboard Behavior

Theme changes do not interrupt typing.

---

# Safe Area

Theme applies consistently across

- Chat Area
- Status Bar
- Input Area
- Dialogs

---

# Performance

Theme changes should update without full application reload.

Use cached tokens.

---

# Security

Theme files are treated as configuration only.

Never execute theme content.

---

# Error Handling

Invalid theme

Display

```text
Theme loading failed.

Using default theme.
```

---

# Restrictions

Never

- Hardcode colors inside components
- Use unreadable contrast
- Depend only on color meaning
- Use emoji as icons
- Break terminal compatibility

---

# Example Theme Mapping

```text
Primary

↓

Selected Command

↓

Active Model

↓

Focused Element
```

---

# Example Status Colors

```text
Success

Task completed


Warning

Attention needed


Error

Operation failed
```

---

# Theme Checklist

Every Theme System must

- Use design tokens
- Support multiple themes
- Support customization
- Support accessibility
- Support syntax themes
- Support persistence
- Avoid hardcoded colors
- Work in terminal environments
- Remain lightweight
- Work on Android Termux

---

# Core Rules

1. Every UI color comes from theme tokens.
2. Components never define their own colors.
3. Default theme must be readable.
4. Contrast must always be maintained.
5. Themes must survive restart.
6. Color is never the only indicator.
7. Icons remain professional and text-free.
8. Theme switching must be instant.
9. Terminal compatibility is mandatory.
10. Optimize for mobile environments.

---

# Summary

The Theme System provides the visual foundation of the Mobile AI CLI. Through token-based colors, component-level styling, syntax support, accessibility features, and customizable themes, it creates a consistent terminal-native interface optimized for Android Termux while maintaining a modern AI CLI experience.