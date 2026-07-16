# 08-safe-area.md

# Safe Area
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Safe Area System for the Mobile AI CLI.

The Safe Area ensures that every UI component remains visible, accessible, and interactive regardless of Android device characteristics, terminal size, keyboard state, gesture navigation, or display cutouts.

Every screen must respect these rules.

---

# Design Goals

The Safe Area system must be

- Mobile First
- Termux Optimized
- Keyboard Safe
- Touch Friendly
- One-Hand Friendly
- Terminal Native
- Predictable
- Consistent

---

# Supported Platform

Primary

- Android
- Termux

Orientation

- Portrait Only

---

# Safe Area Philosophy

The user should never lose access to important UI because of

- Software Keyboard
- Gesture Navigation
- Rounded Corners
- Camera Cutout
- Display Notch
- Waterdrop Notch
- Punch Hole Camera
- Curved Display
- Terminal Resize

---

# Protected Areas

The application must always protect

```
Top Safe Area

↓

Conversation Area

↓

Input Area

↓

Status Bar

↓

Bottom Safe Area

↓

System Navigation
```

---

# Safe Area Priority

```
Keyboard

↓

Input

↓

Status Bar

↓

Conversation
```

The keyboard has the highest layout priority.

---

# Top Safe Area

Protect against

- Camera Cutout
- Status Bar
- Rounded Corner

Top padding

```
Auto
```

Provided by the terminal.

Never manually increase unless required.

---

# Bottom Safe Area

Protect against

- Gesture Navigation
- Navigation Bar
- Keyboard

The application must never render below the safe area.

---

# Left Safe Area

Protect

- Rounded Corner
- Curved Display

Minimum padding

```
1 Cell
```

---

# Right Safe Area

Protect

- Rounded Corner
- Curved Display

Minimum padding

```
1 Cell
```

---

# Conversation Safe Area

Conversation may shrink.

Conversation may scroll.

Conversation must never overlap

- Input
- Status Bar
- Keyboard

---

# Input Safe Area

The Input Area always remains fully visible.

Never allow

```
Keyboard

↓

Input Hidden
```

This is prohibited.

---

# Status Bar Safe Area

The Status Bar always remains below the Input Area.

Layout

```
Conversation

↓

Input

↓

Status

↓

Keyboard
```

Never

```
Conversation

↓

Status

↓

Input
```

---

# Keyboard Safe Layout

Keyboard Closed

```
┌──────────────────────────────┐
│                              │
│                              │
│        Conversation          │
│                              │
├──────────────────────────────┤
│ > Input                      │
├──────────────────────────────┤
│ Status                       │
└──────────────────────────────┘
```

---

Keyboard Open

```
┌──────────────────────────────┐
│                              │
│      Conversation            │
│                              │
├──────────────────────────────┤
│ > Input                      │
├──────────────────────────────┤
│ Status                       │
├──────────────────────────────┤
│                              │
│      Keyboard                │
│                              │
└──────────────────────────────┘
```

---

# Keyboard Rules

When keyboard opens

Conversation

```
Shrink
```

Input

```
Remain Visible
```

Status

```
Remain Visible
```

Keyboard must never cover either.

---

# Gesture Navigation

Protect

Bottom Gesture Area

Never place

- Buttons
- Input
- Dialog Actions

inside the gesture zone.

---

# Terminal Resize

On resize

```
Detect

↓

Measure Safe Area

↓

Recalculate Layout

↓

Render
```

---

# Dynamic Safe Area

Safe Area updates automatically when

- Keyboard opens
- Keyboard closes
- Terminal resizes
- Orientation changes
- Font size changes

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

Landscape optimization is intentionally excluded.

---

# Dialog Safe Area

Dialogs must

Remain centered.

Never touch screen edges.

Minimum margin

```
2 Cells
```

---

# Overlay Safe Area

Every overlay respects

- Top Safe Area
- Bottom Safe Area
- Left Safe Area
- Right Safe Area

---

# Command Palette

Maximum height

```
70%
```

Never overlap keyboard.

If keyboard opens

Command Palette moves upward automatically.

---

# File Picker

Scrollable

Must always remain above

Input

Status

Keyboard

---

# Notification Safe Area

Toast messages

Appear above

Input Area

Never cover

Input

Status

Keyboard

---

# Search Overlay

Must resize dynamically.

Conversation remains partially visible.

---

# Tool Execution

Tool output appears

Inside Conversation

Only

Never inside the Safe Area reserved for Input.

---

# Selection Area

Text selection must remain visible.

Auto-scroll if necessary.

---

# Cursor Visibility

Cursor must always remain visible while typing.

If hidden

Automatically scroll input.

---

# Accessibility

Large Font

Safe Area recalculates automatically.

Reduced Motion

No Safe Area behavior changes.

Screen Reader

Safe Area remains identical.

---

# Minimum Visible Areas

Conversation

```
At least 30% Height
```

Input

```
100% Visible
```

Status

```
100% Visible
```

---

# Prohibited Layouts

Never allow

```
Keyboard

↓

Input Hidden
```

---

Never allow

```
Status Hidden
```

---

Never allow

```
Conversation Behind Keyboard
```

---

Never allow

```
Dialog Behind Keyboard
```

---

Never allow

```
Overlay Outside Safe Area
```

---

# Display Cutouts

Support

- Waterdrop
- Punch Hole
- Center Notch
- Corner Notch
- Rounded Corner

No UI element may be clipped.

---

# Curved Displays

Protect

Left Edge

Right Edge

No important content may touch the curved display.

---

# Touch Safe Area

Minimum touch distance from edge

```
1 Cell
```

Prevents accidental touches.

---

# Performance

Safe Area calculation should occur only when

- Keyboard changes
- Terminal resizes
- Device metrics change

Never calculate every frame.

---

# Safe Area Hierarchy

```
Display

↓

Terminal

↓

Safe Area

↓

Layout

↓

Components
```

---

# Example Layout

```
┌──────────────────────────────┐
│                              │
│ Assistant                    │
│                              │
│ Project created successfully │
│                              │
│ npm install completed        │
│                              │
├──────────────────────────────┤
│ > npm run dev                │
├──────────────────────────────┤
│ GPT-5 │ Ready │ Main         │
└──────────────────────────────┘
```

Every component remains inside the protected area.

---

# Safe Area Checklist

Every screen must

- Respect display cutouts
- Respect rounded corners
- Respect keyboard
- Respect gesture navigation
- Keep Input visible
- Keep Status visible
- Keep Conversation scrollable
- Recalculate on resize
- Preserve one-hand usability

---

# Core Rules

1. Input is never hidden.
2. Status Bar is always visible.
3. Conversation shrinks before Input moves.
4. Keyboard has highest layout priority.
5. Safe Area recalculates automatically.
6. Overlays respect all safe boundaries.
7. Nothing important touches screen edges.
8. Portrait mode only.
9. One-handed interaction is always preserved.
10. Safe Area behavior must remain predictable.

---

# Summary

The Safe Area System guarantees that every interaction remains accessible on Android Termux regardless of keyboard state, display shape, or terminal size.

By prioritizing the keyboard, protecting the Input Area and Status Bar, and dynamically resizing the Conversation Area, the Mobile AI CLI maintains a stable, terminal-native, mobile-first experience optimized for continuous AI-assisted coding workflows.