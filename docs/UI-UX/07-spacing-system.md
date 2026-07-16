# 07-spacing-system.md

# Spacing System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the complete spacing system used throughout the Mobile AI CLI.

Spacing establishes visual hierarchy, improves readability, separates logical groups, and ensures a consistent terminal-native experience across all screens.

Every layout, component, dialog, overlay, and interaction must follow this specification.

---

# Design Goals

The spacing system must be

- Mobile First
- Terminal Native
- Consistent
- Predictable
- Minimal
- Readable
- Performance Friendly

---

# Platform

Primary Platform

- Android
- Termux

Orientation

- Portrait Only

---

# Design Philosophy

Whitespace is an interface element.

Spacing should organize information rather than decorate it.

Empty space is intentional.

---

# Spacing Unit

Base Unit

```
1 Cell
```

Every spacing value is derived from this unit.

---

# Spacing Scale

```
0
1
2
3
4
6
8
12
16
```

Never create arbitrary spacing values.

---

# Layout Spacing

Between major layout sections

```
0
```

Conversation

↓

Input

↓

Status

No empty rows.

---

# Screen Padding

Left

```
1 Cell
```

Right

```
1 Cell
```

Top

```
0
```

Bottom

```
0
```

---

# Conversation Padding

Inside message container

```
1 Cell
```

---

# Input Padding

Left

```
1 Cell
```

Right

```
1 Cell
```

Top

```
0
```

Bottom

```
0
```

---

# Status Bar Padding

Horizontal

```
1 Cell
```

Vertical

```
0
```

---

# Dialog Padding

Internal

```
2 Cells
```

External Margin

```
2 Cells
```

---

# Overlay Padding

Internal

```
2 Cells
```

Outer Margin

```
2 Cells
```

---

# Card Padding

Internal

```
1 Cell
```

Never exceed

```
2 Cells
```

---

# List Spacing

Between list items

```
0
```

Related items should remain compact.

---

# Section Spacing

Between logical sections

```
1 Empty Line
```

Never use multiple empty lines.

---

# Paragraph Spacing

Between paragraphs

```
1 Empty Line
```

---

# Heading Spacing

Before Heading

```
1 Empty Line
```

After Heading

```
1 Empty Line
```

---

# Message Spacing

Between conversation messages

```
1 Empty Line
```

Never stack messages without separation.

---

# Code Block Spacing

Before

```
1 Empty Line
```

After

```
1 Empty Line
```

Inside code

```
Preserve Original Formatting
```

---

# Table Spacing

Above Table

```
1 Empty Line
```

Below Table

```
1 Empty Line
```

Cell Padding

```
1 Space
```

---

# Tree View

Indentation

```
4 Spaces
```

Example

```
Workspace
    src
        components
        pages
```

---

# Markdown Spacing

Heading

```
1 Empty Line
```

Paragraph

```
1 Empty Line
```

List

```
0
```

Quote

```
1 Empty Line
```

Code

```
1 Empty Line
```

---

# Command Palette

Item Padding

```
1 Cell
```

Search Box Margin

```
1 Cell
```

---

# Search Results

Between results

```
0
```

Between result groups

```
1 Empty Line
```

---

# Toolbar

Between icons

```
1 Space
```

Maximum

```
2 Spaces
```

---

# Status Items

Between items

```
2 Spaces
```

Example

```
Model  Tokens  Session  Branch
```

---

# Icon Spacing

Icon

↓

Label

```
1 Space
```

Never remove this spacing.

---

# Button Padding

Horizontal

```
2 Spaces
```

Vertical

```
0
```

---

# Badge Padding

Horizontal

```
1 Space
```

Vertical

```
0
```

---

# Notification Spacing

Internal

```
1 Cell
```

External

```
1 Cell
```

---

# Loading Indicator

Spinner

↓

Label

```
1 Space
```

---

# Empty State

Illustration (if any)

↓

Message

```
1 Empty Line
```

Message

↓

Action

```
1 Empty Line
```

---

# Error State

Title

↓

Description

```
1 Empty Line
```

Description

↓

Action

```
1 Empty Line
```

---

# Keyboard Safe Area

When keyboard opens

Conversation shrinks.

Spacing values never change.

---

# Terminal Resize

Spacing scale remains identical.

Only available layout space changes.

---

# Alignment Rules

Always align components to the spacing grid.

Never manually offset components.

---

# Scroll Area

Top Padding

```
0
```

Bottom Padding

```
1 Cell
```

Allows the final message to remain readable above the input area.

---

# Touch Targets

Minimum touch area

```
44 × 44 dp
```

Preferred

```
48 × 48 dp
```

Spacing must support comfortable touch interaction.

---

# Accessibility

Spacing must remain sufficient for

- Large fonts
- Screen readers
- High contrast themes

Never compress spacing to increase density.

---

# Performance

Spacing values are constants.

Never calculate spacing dynamically.

Cache layout measurements whenever possible.

---

# Responsive Rules

Small screens

Reduce content width.

Do not reduce spacing scale.

Maintain consistent rhythm.

---

# Restrictions

Never use

- Random spacing
- Double empty lines
- Negative margins
- Overlapping components
- Manual alignment using repeated spaces
- Decorative whitespace

---

# Spacing Hierarchy

```
Layout

↓

Section

↓

Message

↓

Paragraph

↓

Inline Content
```

Each level becomes progressively more compact.

---

# Example Layout

```
Conversation

Assistant

Project initialized successfully.

User

Create a React application.

Assistant

Running command...

npm create vite@latest
```

Every message is separated by one empty line.

---

# Example Status Bar

```
Model  Tokens  Session  Branch  Ready
```

Consistent spacing improves readability.

---

# Example Dialog

```
+--------------------------+

  Delete Workspace

  This action cannot
  be undone.

  Cancel   Delete

+--------------------------+
```

Internal padding remains consistent.

---

# Spacing Checklist

Every screen must

- Use the base spacing scale
- Maintain one-column rhythm
- Preserve message spacing
- Preserve code formatting
- Keep icons separated from labels
- Maintain touch-friendly spacing
- Respect keyboard-safe layout
- Avoid unnecessary whitespace

---

# Core Rules

1. Base spacing unit is one terminal cell.
2. Every spacing value uses the predefined scale.
3. Never insert arbitrary whitespace.
4. Separate messages with one empty line.
5. Preserve original code spacing.
6. Align everything to the layout grid.
7. Maintain spacing consistency across all screens.
8. Optimize for one-handed mobile use.
9. Never sacrifice readability for density.
10. Spacing is part of the information architecture.

---

# Summary

The Spacing System defines a consistent rhythm for the Mobile AI CLI.

Using a fixed terminal-cell scale ensures predictable layouts, improved readability, and stable rendering across Android Termux environments. Every component, message, dialog, overlay, and interaction follows the same spacing principles, creating a clean, mobile-first, terminal-native experience.