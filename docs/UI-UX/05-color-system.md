# 05-color-system.md

# Color System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the complete color system used throughout the Mobile AI CLI.

The color system establishes a consistent visual language for terminal rendering while maintaining readability, accessibility, and compatibility with Android Termux.

Every screen, component, and state must follow this specification.

---

# Design Goals

The color system must be

- Mobile First
- Terminal Native
- High Contrast
- Accessible
- Consistent
- Themeable
- Performance Friendly

---

# Platform

Primary Platform

- Android
- Termux

Rendering

- ANSI Terminal Colors
- True Color (24-bit) when supported
- ANSI 256 fallback
- ANSI 16 fallback

---

# Design Philosophy

Color exists to communicate information.

Color is **never decoration**.

Every color must have a semantic meaning.

---

# Theme Philosophy

Every theme uses the same semantic color tokens.

Only color values change.

The UI structure never changes.

---

# Color Hierarchy

```
Background

↓

Surface

↓

Primary Content

↓

Secondary Content

↓

Accent

↓

Status Colors
```

---

# Semantic Color Tokens

Core semantic tokens

```
Background
Surface
Border
Foreground
Muted
Primary
Secondary
Accent
Selection
Cursor
```

Status tokens

```
Success
Warning
Error
Info
Loading
Disabled
```

Conversation tokens

```
User
Assistant
System
Tool
Thinking
Streaming
Code
Markdown
```

---

# Background

Purpose

Main terminal background.

Rules

- Always darkest layer.
- Never use gradients.
- Never use textures.

---

# Surface

Purpose

Panels

Dialogs

Input

Command Palette

Rules

Slightly different from Background.

Must remain visually subtle.

---

# Foreground

Purpose

Primary readable text.

Must always provide excellent contrast.

---

# Muted

Purpose

Secondary information.

Examples

- Timestamp
- Hint
- Placeholder
- Metadata

Never use for important content.

---

# Primary

Purpose

Primary actions.

Examples

- Selected item
- Active component
- Current command

---

# Secondary

Purpose

Less important actions.

Examples

- Optional controls
- Secondary labels

---

# Accent

Purpose

Highlight interactive elements.

Examples

- Active cursor
- Focus ring
- Selection

Accent should never dominate the screen.

---

# Border

Purpose

Separate logical sections.

Examples

Input

Dialogs

Tables

Status Bar

Borders should remain subtle.

---

# Cursor

Purpose

Indicate text insertion position.

Must remain highly visible.

---

# Selection

Purpose

Selected text.

Selected file.

Selected command.

Selection must remain readable.

---

# Success

Used for

- Completed
- Saved
- Finished
- Connected

Never use for warnings.

---

# Warning

Used for

- Confirmation required
- Validation
- Missing information

Warnings should attract attention without appearing dangerous.

---

# Error

Used for

- Failed operations
- Invalid input
- Permission denied

Errors must remain readable.

---

# Info

Used for

- General information
- Notifications
- Progress

---

# Loading

Used during

- Streaming
- Processing
- Searching
- Indexing

---

# Disabled

Used for

Unavailable actions.

Disabled items must remain readable.

---

# User Message

Purpose

Differentiate user messages from assistant messages.

Must remain subtle.

---

# Assistant Message

Purpose

Primary conversation output.

Highest readability priority.

---

# Tool Output

Purpose

Differentiate tool execution from conversation.

Examples

- Git
- Search
- File operations

---

# Thinking

Purpose

Represent internal reasoning state.

Should not distract from the conversation.

---

# Streaming

Purpose

Indicate live generation.

Must appear dynamic without excessive animation.

---

# Markdown

Markdown preserves semantic highlighting.

Examples

Headings

Links

Lists

Quotes

---

# Code

Purpose

Syntax highlighting.

Color comes from the syntax highlighting engine.

The surrounding layout still follows semantic colors.

---

# Input Area

Background

Surface

Text

Foreground

Cursor

Accent

Placeholder

Muted

---

# Status Bar

Background

Surface

Foreground

Metadata

Muted

Warnings

Warning

Errors

Error

---

# Dialog

Background

Surface

Border

Border

Primary Action

Primary

Cancel

Secondary

---

# Command Palette

Background

Surface

Selected Item

Primary

Search Text

Foreground

Hints

Muted

---

# Tables

Header

Primary

Rows

Foreground

Border

Border

Selection

Accent

---

# Tree View

Folder

Primary

File

Foreground

Selection

Accent

Metadata

Muted

---

# Search Results

Matched text

Accent

File name

Primary

Path

Muted

---

# Terminal Output

Standard output

Foreground

Warnings

Warning

Errors

Error

Success

Success

---

# Focus

Focused components

Accent

Unfocused components

Default colors

---

# Color Priority

```
Error

↓

Warning

↓

Success

↓

Accent

↓

Primary

↓

Foreground

↓

Muted
```

Higher priority colors override lower priority colors.

---

# Accessibility

Minimum contrast

```
WCAG AA
```

Preferred

```
WCAG AAA
```

Do not rely on color alone.

Always combine

Color

+

Label

+

Icon

---

# Theme Support

Supported

Light

Dark

High Contrast

OLED Black

Custom

Every theme must preserve semantic meaning.

---

# Performance

Colors should be

- Cached
- Reused
- Theme-driven

Avoid unnecessary recalculation.

---

# Responsive Behavior

Changing theme must

```
Update Colors

↓

Keep Layout

↓

Keep State

↓

No Flicker
```

---

# Restrictions

Never use

- Random colors
- Decorative gradients
- Rainbow text
- Flashing colors
- Neon effects
- Heavy shadows
- Glass effects

---

# Recommended Visual Hierarchy

```
Background

↓

Surface

↓

Conversation

↓

Input

↓

Status

↓

Dialogs
```

---

# Example Screen

```
┌──────────────────────────────┐
│ Assistant                    │
│                              │
│ Project created successfully │
│                              │
│ Tool Output                  │
│ Reading files...             │
├──────────────────────────────┤
│ > npm run dev                │
├──────────────────────────────┤
│ GPT-5 │ 2.4K │ Ready         │
└──────────────────────────────┘
```

Semantic colors determine appearance.

No component hardcodes colors.

---

# Color Rules

1. Background is always the lowest visual layer.
2. Conversation has the highest readability priority.
3. Input always remains clearly visible.
4. Status Bar uses subdued colors.
5. Errors always use Error color.
6. Warnings always use Warning color.
7. Success always uses Success color.
8. Accent is reserved for focus and selection.
9. Code highlighting is handled separately.
10. Themes only change values, never semantics.

---

# Summary

The Color System establishes a semantic, theme-driven color architecture for the Mobile AI CLI.

Every color communicates meaning rather than decoration, ensuring a clean terminal-native interface optimized for Android Termux, mobile-first interaction, accessibility, and high-performance rendering across multiple terminal color capabilities.
