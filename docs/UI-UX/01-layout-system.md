# 01-layout-system.md

# Layout System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the complete layout system used throughout the Mobile AI CLI.

Every screen must use this layout system.

No screen may introduce its own layout rules.

---

# Design Goals

The layout system must be

- Mobile First
- Terminal Native
- Thumb Friendly
- Keyboard Safe
- Consistent
- Predictable
- Fast
- Portrait Optimized

---

# Supported Devices

Supported

- Android Phone
- Termux
- Portrait Mode

Primary Width

```
320dp → 480dp
```

Primary Height

```
640dp → 960dp+
```

---

# Unsupported Layouts

Never support

- Desktop Layout
- Tablet Layout
- Landscape Layout
- Split Screen
- Multi Window Layout
- Sidebar Layout

---

# Global Screen Structure

Every screen follows exactly this order.

```
Conversation Area

↓

Input Area

↓

Status Bar
```

Nothing may appear below the Status Bar.

---

# Master Layout

```
┌────────────────────────────────┐
│                                │
│                                │
│                                │
│                                │
│                                │
│                                │
│        Conversation Area       │
│                                │
│                                │
│                                │
│                                │
├────────────────────────────────┤
│ > Command Input                │
├────────────────────────────────┤
│ Model │ Tokens │ Session │ CPU │
└────────────────────────────────┘
```

---

# Layout Zones

## Zone 1

Conversation

Largest area.

Contains

- AI messages
- User messages
- Code
- Markdown
- Tool output
- Search results

---

## Zone 2

Input

Contains

- Prompt input
- Cursor
- Inline suggestions

---

## Zone 3

Status

Contains

- Active Model
- Token Count
- Session
- Workspace
- Connection
- Git Branch

---

# Layout Priority

```
Conversation

↓

Input

↓

Status
```

Conversation always receives maximum available space.

---

# Safe Area

Always respect

Top

```
Status Bar Height
```

Bottom

```
Keyboard
Gesture Area
```

Left

```
Screen Padding
```

Right

```
Screen Padding
```

---

# Outer Margin

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

# Inner Padding

Conversation

```
1 Cell
```

Input

```
1 Cell
```

Status

```
1 Cell
```

---

# Grid System

Use a fixed terminal grid.

```
Columns

↓

Dynamic

Rows

↓

Dynamic
```

Never use pixel-perfect positioning.

---

# Component Alignment

Horizontal

```
Left
```

Vertical

```
Top
```

Center alignment only for loading or splash screens.

---

# Layout Rules

Always

- Stack vertically
- Fill available width
- Keep consistent spacing

Never

- Overlay components
- Float components
- Use absolute positioning unnecessarily

---

# Conversation Area

Must expand automatically.

```
Minimum Height

↓

Remaining Space
```

If messages overflow

```
Scroll
```

---

# Input Area

Always fixed.

Never scroll.

Height

```
Auto
```

Minimum

```
1 Line
```

Maximum

```
8 Lines
```

---

# Status Bar

Always fixed.

Height

```
1 Line
```

Never scroll.

---

# Keyboard Behavior

When keyboard opens

```
Conversation

↓

Shrink
```

Input remains visible.

Status Bar remains visible.

Nothing may move below keyboard.

---

# Keyboard Layout

Closed

```
Conversation

↓

Input

↓

Status
```

Opened

```
Conversation (Reduced)

↓

Input

↓

Status

↓

Keyboard
```

---

# Scrolling Rules

Conversation

```
Scrollable
```

Input

```
Not Scrollable
```

Status

```
Never Scrollable
```

---

# Auto Scroll

Automatically scroll only when

- User sends message
- AI streams response

Never interrupt manual scrolling.

---

# Touch Zones

Conversation

```
Tap

Long Press

Select
```

Input

```
Tap

Paste

Selection
```

Status

```
Tap
```

---

# One-Hand Usage

Primary controls must remain within the lower thumb zone.

Avoid placing important controls near the top edge.

---

# Maximum Visible Layers

Only three layers

```
Conversation

↓

Input

↓

Status
```

No additional floating layers.

---

# Overlay Policy

Allowed

- Dialog
- Command Palette
- Confirmation

Not Allowed

- Floating Buttons
- Floating Chat Heads
- Persistent Popups

---

# Sidebar Policy

Never use

```
Left Sidebar

Right Sidebar

Hidden Sidebar
```

Navigation occurs inside the main content.

---

# Split Layout

Not allowed.

Never display

```
Editor | Chat

Explorer | Chat

Sidebar | Content
```

Single-column layout only.

---

# Width Usage

Components use

```
100% Width
```

No centered narrow layouts.

---

# Height Usage

Conversation uses

```
Remaining Height
```

Input

```
Content Height
```

Status

```
Fixed Height
```

---

# Screen Resize

When terminal resizes

```
Recalculate Layout

↓

Reflow Components

↓

Redraw
```

---

# Orientation

Supported

```
Portrait
```

Unsupported

```
Landscape
```

---

# Loading Layout

```
┌──────────────────────────────┐
│                              │
│                              │
│         Spinner              │
│                              │
│      Loading Project         │
│                              │
├──────────────────────────────┤
│ >                            │
├──────────────────────────────┤
│ Loading...                   │
└──────────────────────────────┘
```

---

# Empty Layout

```
┌──────────────────────────────┐
│                              │
│      No Conversation         │
│                              │
│ Start by typing a prompt     │
│                              │
├──────────────────────────────┤
│ >                            │
├──────────────────────────────┤
│ Ready                        │
└──────────────────────────────┘
```

---

# Active Conversation Layout

```
┌──────────────────────────────┐
│ User                         │
│ Create React App             │
│                              │
│ Assistant                    │
│ Project initialized          │
│                              │
│ Assistant                    │
│ Generating files...          │
├──────────────────────────────┤
│ > npm run dev                │
├──────────────────────────────┤
│ GPT-5 │ 3.2K │ Main │ Ready  │
└──────────────────────────────┘
```

---

# Layout Constraints

Never

- Hide input
- Hide status
- Overlap keyboard
- Clip conversation
- Place controls outside safe area

Always

- Preserve layout order
- Keep consistent spacing
- Adapt to terminal resize
- Maintain readability

---

# Performance

Layout recalculation should occur only when

- Terminal size changes
- Keyboard state changes
- Input height changes
- New content arrives

Avoid unnecessary re-layout operations.

---

# Accessibility

Support

- Large fonts
- Screen readers
- High contrast
- Reduced motion

Maintain sufficient spacing for touch interaction.

---

# Layout Checklist

Every screen must

- Use a single column
- Fill available width
- Keep input fixed
- Keep status fixed
- Keep conversation scrollable
- Respect keyboard safe area
- Respect bottom gesture area
- Avoid sidebars
- Avoid floating UI
- Support one-hand operation

---

# Summary

The Layout System defines a strict single-column, mobile-first structure optimized for Android Termux.

Every screen consists of only three persistent zones:

- Conversation Area
- Input Area
- Status Bar

This approach ensures predictable navigation, efficient one-handed interaction, keyboard-safe behavior, and a clean terminal-native experience suitable for an AI-powered CLI on mobile devices.