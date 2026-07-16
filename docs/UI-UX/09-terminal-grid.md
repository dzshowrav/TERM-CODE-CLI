# 09-terminal-grid.md

# Terminal Grid System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Terminal Grid System used throughout the Mobile AI CLI.

Every screen, component, overlay, message, dialog, code block, markdown renderer, and interactive element must align to this grid.

The grid is the foundation of the entire interface.

---

# Design Goals

The Terminal Grid must be

- Mobile First
- Termux Optimized
- Terminal Native
- Consistent
- Predictable
- Responsive
- Keyboard Safe
- Performance Friendly

---

# Supported Platform

Primary

- Android
- Termux

Orientation

- Portrait Only

---

# Design Philosophy

Everything is positioned using terminal cells.

Never use pixel-based positioning.

Every visible element occupies one or more terminal cells.

---

# Grid Unit

Base Unit

```
1 Cell
```

Every measurement derives from this unit.

---

# Grid Coordinate System

```
Column

↓

X

Row

↓

Y
```

Origin

```
(0,0)
```

Located at the upper-left corner of the terminal viewport.

---

# Grid Structure

```
Rows

↓

Dynamic

Columns

↓

Dynamic
```

The renderer calculates the available grid from the terminal size.

---

# Typical Mobile Terminal

Approximate portrait size

```
Columns

36–50

Rows

24–60
```

Never hardcode dimensions.

---

# Dynamic Grid

Grid dimensions update automatically when

- Terminal resizes
- Font changes
- Keyboard opens
- Keyboard closes

---

# Cell Rules

Every cell contains

- One Unicode character
- One Nerd Font icon
- One Powerline symbol
- One space

Never render more than one visible glyph per cell.

---

# Cell Width

Standard width

```
1 Cell
```

Wide Unicode characters must be handled correctly by the renderer.

---

# Cell Height

```
1 Row
```

No component may use fractional rows.

---

# Layout Grid

```
Conversation

↓

Input

↓

Status
```

All regions align to the grid.

---

# Horizontal Alignment

Supported

```
Left

Center

Right
```

Default

```
Left
```

---

# Vertical Alignment

Supported

```
Top

Center

Bottom
```

Default

```
Top
```

---

# Padding

Horizontal

```
1 Cell
```

Vertical

```
0 Cell
```

---

# Margin

Outer Margin

```
1 Cell
```

Only where appropriate.

---

# Grid Snapping

Every component snaps to

```
Grid Cell
```

Never allow half-cell positioning.

---

# Conversation Grid

Messages occupy complete rows.

Example

```
Assistant

Project initialized.

User

Create a React app.
```

---

# Input Grid

Input expands vertically.

Minimum

```
1 Row
```

Maximum

```
8 Rows
```

---

# Status Grid

Height

```
1 Row
```

Always fixed.

---

# Dialog Grid

Width

```
Maximum 90%
```

Height

```
Content Based
```

Position

```
Centered
```

---

# Overlay Grid

Every overlay aligns to

Grid

Safe Area

Never overlap protected regions.

---

# Table Grid

Columns align perfectly.

Example

```
Name      Status

src       Ready

README    Modified
```

Never use proportional spacing.

---

# Tree Grid

Indentation

```
4 Spaces
```

Example

```
workspace

    src

        components

        pages

    package.json
```

---

# Code Grid

Every character occupies one cell.

Preserve

- Indentation
- Blank lines
- Alignment

Never modify code spacing.

---

# Markdown Grid

Headings

Paragraphs

Lists

Tables

Quotes

Code

All align to the grid.

---

# Icon Grid

Every icon occupies

```
1 Cell
```

Spacing

```
1 Space

↓

Label
```

---

# Cursor Grid

Cursor occupies

```
1 Cell
```

Must never overlap another character.

---

# Selection Grid

Selection aligns exactly to character cells.

No partial selection rendering.

---

# Scroll Grid

Scrolling occurs

```
Row by Row
```

Never scroll partial rows.

---

# Resize Behavior

On resize

```
Measure Terminal

↓

Recalculate Grid

↓

Reflow Layout

↓

Render
```

---

# Keyboard Behavior

When keyboard opens

Grid height decreases.

Grid width remains unchanged.

Conversation shrinks.

Input remains fixed.

Status remains fixed.

---

# Safe Area Integration

Grid begins

After Top Safe Area.

Grid ends

Before Bottom Safe Area.

---

# Responsive Rules

Small terminal

Reduce

Visible content.

Never reduce

Grid size.

Never scale characters.

---

# Rendering Order

```
Grid

↓

Layout

↓

Components

↓

Text

↓

Cursor
```

---

# Layering

Layer 1

Conversation

Layer 2

Input

Layer 3

Status

Layer 4

Dialog

Layer 5

Notification

Maximum

```
5 Layers
```

---

# Performance

Render only

Changed Cells.

Avoid

Full screen redraw.

---

# Accessibility

Large fonts

Grid recalculates automatically.

Screen readers

Grid has no effect.

Reduced motion

No change.

---

# Restrictions

Never use

- Absolute pixel positioning
- Fractional rows
- Fractional columns
- Floating elements
- Freeform positioning
- Arbitrary offsets

---

# Grid Hierarchy

```
Terminal

↓

Viewport

↓

Grid

↓

Layout

↓

Component

↓

Character
```

---

# Example Grid

```
┌──────────────────────────────────┐
│ Assistant                        │
│                                  │
│ Project created successfully.    │
│                                  │
│ npm install completed.           │
├──────────────────────────────────┤
│ > npm run dev                    │
├──────────────────────────────────┤
│ GPT-5 │ Main │ 2.1K │ Ready      │
└──────────────────────────────────┘
```

Every visible element aligns to the terminal grid.

---

# Grid Checklist

Every screen must

- Align to terminal cells
- Snap every component to the grid
- Preserve code alignment
- Preserve table alignment
- Respect safe areas
- Support dynamic resizing
- Support keyboard-safe layouts
- Avoid pixel positioning
- Avoid freeform layouts

---

# Core Rules

1. One cell is the smallest layout unit.
2. Every component snaps to the grid.
3. Never use pixel coordinates.
4. Grid dimensions are dynamic.
5. Preserve character alignment.
6. Scroll only by rows.
7. Code formatting is never modified.
8. Grid recalculates automatically.
9. Rendering is incremental.
10. The terminal grid defines the entire interface.

---

# Summary

The Terminal Grid System provides a terminal-native layout foundation for the Mobile AI CLI. Every screen is constructed from dynamic terminal cells rather than pixels, ensuring consistent rendering across Android Termux environments. By aligning every component, message, code block, and interactive element to the same grid, the interface remains predictable, responsive, keyboard-safe, and optimized for one-handed mobile use.