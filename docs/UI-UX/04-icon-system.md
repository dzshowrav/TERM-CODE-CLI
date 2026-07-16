# 04-icon-system.md

# Icon System
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the complete icon system used throughout the Mobile AI CLI.

Icons provide fast visual recognition while preserving a clean, terminal-native experience.

Every icon used by the application must follow this specification.

---

# Design Goals

The icon system must be

- Mobile First
- Terminal Native
- Unicode Friendly
- Termux Compatible
- Consistent
- Minimal
- Readable
- Accessible

---

# Platform

Primary

- Android
- Termux

Rendering Environment

- Terminal Emulator
- Monospace Font

---

# Design Philosophy

Icons should communicate meaning instantly.

Icons support text.

Icons never replace important text.

---

# Icon Sources

Allowed

- Nerd Fonts
- Codicons
- Powerline Symbols
- Unicode Terminal Symbols

Preferred

- Nerd Fonts

---

# Prohibited

Never use

- Emoji
- Decorative Icons
- Animated Emoji
- Bitmap Images
- PNG Icons
- SVG Icons inside the terminal
- Colored Emoji Fonts

---

# Icon Style

Icons must be

- Monochrome
- Lightweight
- Sharp
- Consistent
- Monospaced Friendly

---

# Icon Size

Standard

```
1 Cell
```

Large

```
2 Cells
```

Avoid wider icons.

---

# Icon Position

Icons always appear

```
Icon Space Label
```

Example

```
[icon] Open Workspace
```

Never

```
Label Icon
```

---

# Icon Alignment

Horizontal

```
Left
```

Vertical

```
Center
```

---

# Icon Spacing

Between icon and label

```
1 Space
```

Never place multiple spaces.

---

# Icon Categories

Application

Files

Folders

Search

Git

Terminal

AI

Tools

Status

Warnings

Errors

Success

Navigation

Sessions

Workspace

Models

Settings

Permissions

Network

---

# Application Icons

Used for

- Application
- Workspace
- Project

One icon per object.

---

# File Icons

Used for

- Files
- Documents
- Markdown
- Images
- JSON
- Source Code

Different file types may have different icons.

---

# Folder Icons

Used for

- Folder
- Open Folder
- Workspace
- Project Root

---

# Search Icons

Used for

- Search
- Find
- Replace
- Filter

---

# Git Icons

Used for

- Branch
- Commit
- Merge
- Push
- Pull
- Conflict

---

# AI Icons

Used for

- Assistant
- Thinking
- Streaming
- Context
- Memory
- Prompt

---

# Tool Icons

Used for

- File Tool
- Terminal Tool
- Search Tool
- Git Tool
- Network Tool

---

# Status Icons

Used for

- Ready
- Busy
- Loading
- Idle
- Connected
- Disconnected

---

# Warning Icons

Used for

- Validation
- Warning
- Attention

Should never look aggressive.

---

# Error Icons

Used for

- Failed
- Error
- Permission Denied
- Crash

Errors must remain readable.

---

# Success Icons

Used for

- Completed
- Success
- Saved
- Finished

---

# Navigation Icons

Used for

- Back
- Forward
- Expand
- Collapse

---

# Session Icons

Used for

- Session
- History
- Resume

---

# Model Icons

Used for

- AI Models
- Providers
- Local Models
- Cloud Models

---

# Settings Icons

Used for

- Preferences
- Configuration
- Theme
- Appearance

---

# Permission Icons

Used for

- File Access
- Terminal Access
- Network Access

---

# Network Icons

Used for

- Online
- Offline
- Sync
- Download
- Upload

---

# Icon Usage Rules

Every interactive element may contain

```
One Icon

+

One Label
```

Avoid multiple icons for a single action.

---

# Conversation Rules

Assistant messages

May display one assistant icon.

User messages

May display one user icon.

System messages

May display one system icon.

---

# Input Area

Allowed

- Prompt icon
- Attachment icon
- Send icon

Maximum

```
3 Icons
```

---

# Status Bar

Allowed

- Model
- Session
- Connection
- Git
- Token

Icons should remain compact.

---

# Toolbar Policy

Toolbars should contain only essential icons.

Maximum

```
5 Icons
```

---

# Overlay Icons

Dialogs

One leading icon.

Confirmation

One icon.

Warnings

One icon.

Errors

One icon.

---

# Icon Colors

Icons inherit

- Foreground Color

Exceptions

Status indicators may use semantic colors.

---

# Disabled Icons

Disabled icons should use reduced contrast.

Never hide them completely.

---

# Active Icons

Active icons should use

- Primary Color
- Accent Color

---

# Loading Icons

Use terminal-compatible animated indicators.

Avoid decorative animation.

---

# Icon Consistency

The same action must always use the same icon.

Never assign multiple icons to one action.

---

# Accessibility

Icons never replace readable text.

Screen readers must receive labels.

---

# Performance

Icons should

- Render quickly
- Occupy one terminal cell whenever possible
- Avoid expensive rendering

---

# Responsive Behavior

Small screens

Reduce labels before removing icons.

Never remove essential icons.

---

# Restrictions

Never use

- Emoji
- Decorative illustrations
- Multi-color graphics
- Large banner icons
- Floating icons

---

# Component Examples

Conversation Item

```
[Icon] Assistant

Message
```

File Item

```
[Icon] package.json
```

Git Branch

```
[Icon] main
```

Tool Execution

```
[Icon] Reading Files
```

Status

```
[Icon] Connected
```

---

# Icon Layout Example

```
┌──────────────────────────────┐
│ [Icon] Assistant             │
│                              │
│ Project created successfully │
├──────────────────────────────┤
│ > npm run dev                │
├──────────────────────────────┤
│ [Icon] GPT │ [Icon] Main     │
└──────────────────────────────┘
```

---

# Icon Checklist

Every icon must

- Be terminal compatible
- Be monochrome
- Be readable
- Have a text label
- Use one-cell width whenever possible
- Follow semantic meaning
- Respect spacing
- Remain consistent

---

# Core Rules

1. Emoji are prohibited.
2. Terminal-compatible icons only.
3. One icon per action.
4. Icons support text, never replace it.
5. Maintain consistent meaning across the application.
6. Prefer Nerd Fonts when available.
7. Always preserve readability.
8. Never overload the interface with icons.

---

# Summary

The Icon System establishes a clean, terminal-native visual language optimized for Android Termux.

Icons enhance recognition without distracting from the conversation, maintain semantic consistency, and remain fully compatible with mobile terminal environments while supporting accessibility, readability, and one-handed interaction.