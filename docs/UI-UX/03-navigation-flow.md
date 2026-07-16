# 03-navigation-flow.md

# Navigation Flow
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the complete navigation flow for the Mobile AI CLI.

Navigation in this application is fundamentally different from a traditional mobile application.

The application behaves like a terminal workspace instead of multiple pages.

---

# Design Goals

Navigation must be

- Mobile First
- Terminal Native
- Fast
- Predictable
- Keyboard Safe
- One-Hand Friendly
- Minimal
- Consistent

---

# Navigation Philosophy

Users should never feel they are moving between pages.

Instead,

the workspace changes,

the content changes,

the context changes,

but the application remains on one primary screen.

---

# Navigation Principles

Navigation is based on

- Context
- Commands
- Temporary Overlays
- Focus

Navigation is NOT based on

- Pages
- Sidebars
- Bottom Navigation
- Drawer Menus
- Tab Bars

---

# Navigation Model

```
Workspace

↓

Conversation

↓

Temporary Overlay

↓

Conversation
```

Every temporary UI returns to the Conversation Screen.

---

# Primary Navigation

The Conversation Screen is always the root.

```
Conversation

↓

Input

↓

Status
```

Everything begins here.

---

# Navigation Layers

Maximum depth

```
3 Layers
```

Example

```
Conversation

↓

Command Palette

↓

Confirmation Dialog
```

No additional layers are allowed.

---

# Navigation Types

Supported

- Command Navigation
- Overlay Navigation
- Context Navigation
- Keyboard Navigation
- Touch Navigation

Not Supported

- Sidebar Navigation
- Multi-Page Navigation
- Bottom Tabs
- Drawer Navigation
- Wizard Navigation

---

# Conversation Flow

```
Open Workspace

↓

Conversation

↓

Prompt

↓

AI Response

↓

Next Prompt

↓

Continue
```

The Conversation Screen never changes.

---

# Overlay Navigation

Example

```
Conversation

↓

Model Selector

↓

Conversation
```

The overlay disappears after selection.

---

# Command Palette Flow

```
Conversation

↓

Command Palette

↓

Run Command

↓

Conversation
```

The Command Palette never becomes a permanent screen.

---

# File Picker Flow

```
Conversation

↓

File Picker

↓

Select File

↓

Conversation
```

---

# Permission Flow

```
Conversation

↓

Permission Dialog

↓

Allow

↓

Conversation
```

or

```
Conversation

↓

Permission Dialog

↓

Deny

↓

Conversation
```

---

# Confirmation Flow

```
Conversation

↓

Confirmation Dialog

↓

Confirm

↓

Conversation
```

---

# Error Flow

```
Conversation

↓

Error Dialog

↓

Dismiss

↓

Conversation
```

Never redirect to a separate error page.

---

# Search Flow

```
Conversation

↓

Search Overlay

↓

Results

↓

Select

↓

Conversation
```

---

# Model Selection

```
Conversation

↓

Model Selector

↓

Choose Model

↓

Conversation
```

---

# Settings Flow

```
Conversation

↓

Settings Overlay

↓

Save

↓

Conversation
```

Settings never replace the primary workspace.

---

# Session Flow

```
Launch

↓

Welcome

↓

Workspace

↓

Conversation

↓

Exit
```

---

# Tool Execution Flow

```
Conversation

↓

Execute Tool

↓

Streaming Output

↓

Complete

↓

Conversation
```

Tool execution always appears inline.

---

# Keyboard Navigation

Supported

```
Arrow Keys

Tab

Enter

Escape

Ctrl Shortcuts
```

Keyboard remains a first-class navigation method.

---

# Touch Navigation

Supported

- Tap
- Long Press
- Vertical Scroll

Not Supported

- Multi-touch navigation
- Swipe between pages
- Gesture-based back stack

---

# Focus Navigation

Focus order

```
Input

↓

Dialog

↓

Command Palette

↓

Conversation Selection
```

Only one component may have focus.

---

# Back Navigation

Android Back Button

Rules

If dialog exists

```
Close Dialog
```

If overlay exists

```
Close Overlay
```

Otherwise

```
Ask Before Exit
```

Never immediately terminate the application.

---

# Deep Navigation

Never create navigation like

```
Page

↓

Sub Page

↓

Sub Page

↓

Sub Page
```

Maximum depth

```
3
```

---

# Workspace Navigation

Changing project

```
Workspace A

↓

Workspace B
```

Conversation resets only if required.

---

# Navigation Memory

Remember

- Last cursor position
- Last scroll position
- Active workspace
- Active model
- Session state

Restore automatically.

---

# Screen Transition Rules

Allowed

- Fade
- Instant
- Minimal Slide

Avoid

- Bounce
- Zoom
- Rotate
- Complex animations

---

# Navigation Timing

Target

```
<100ms
```

Opening overlays should feel immediate.

---

# Interruptions

Incoming events

```
Tool Finished

Permission Request

Model Changed
```

Must never interrupt typing unexpectedly.

---

# Notification Flow

```
Background Event

↓

Toast

↓

Dismiss Automatically
```

Conversation remains visible.

---

# Terminal Resize

When terminal changes

```
Resize

↓

Recalculate Layout

↓

Stay On Current Screen
```

Navigation state must remain unchanged.

---

# Keyboard Open

```
Conversation

↓

Shrink

↓

Input

↓

Status
```

Navigation must remain stable.

---

# Navigation Restrictions

Never use

- Sidebars
- Bottom Navigation
- Floating Navigation
- Drawer Menus
- Multi-column Navigation
- Nested Navigation
- Hidden Navigation

---

# Navigation Diagram

```
                    Launch
                       │
                       ▼
                  Welcome Screen
                       │
                       ▼
                Conversation Screen
                       │
      ┌──────────┬───────────┬───────────┐
      ▼          ▼           ▼           ▼
 Command     File Picker  Settings   Model Selector
 Palette         │            │             │
      └──────────┴───────────┴─────────────┘
                       │
                       ▼
                Conversation Screen
                       │
                 Tool Execution
                       │
                       ▼
                Conversation Screen
                       │
                  Exit Dialog
                       │
             ┌─────────┴─────────┐
             ▼                   ▼
          Cancel               Exit
             │                   │
             ▼                   ▼
      Conversation            Application
                               Closed
```

---

# Navigation Checklist

Every navigation action must

- Preserve workspace
- Preserve conversation
- Preserve input state
- Preserve status bar
- Support keyboard
- Support touch
- Return to conversation
- Be interrupt-safe
- Respect one-hand usage

---

# Core Navigation Rules

1. Conversation is always the home screen.
2. Navigation uses overlays instead of pages.
3. Sidebars are prohibited.
4. Bottom navigation is prohibited.
5. Maximum navigation depth is three layers.
6. Keyboard remains active whenever possible.
7. Back button closes overlays before exiting.
8. Conversation is never hidden during normal navigation.
9. All temporary UI returns to the conversation.
10. Navigation must remain fast, predictable, and terminal-native.

---

# Summary

The Mobile AI CLI follows a **single-workspace navigation architecture**.

Users remain inside one persistent Conversation Screen while temporary interactions such as model selection, file picking, settings, permissions, and command execution appear as lightweight overlays.

This approach minimizes navigation complexity, preserves context, supports one-handed mobile use, and delivers a consistent terminal-native experience optimized for Android Termux.
