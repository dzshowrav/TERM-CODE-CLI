# 20-file-picker.md

# File Picker
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the File Picker used throughout the Mobile AI CLI.

The File Picker allows users to browse, search, preview, and select files or directories for AI interactions, tool execution, uploads, workspace management, and project navigation.

The File Picker must be fully optimized for Android Termux and mobile-first workflows.

---

# Design Goals

The File Picker must be

- Mobile First
- Terminal Native
- Keyboard Friendly
- Touch Friendly
- Fast
- Searchable
- Accessible
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

File selection should require the fewest possible interactions.

Browsing and searching should work together seamlessly.

The File Picker never replaces the Chat Screen.

---

# Display Location

The File Picker appears as a modal overlay above the Conversation Area.

The Chat Screen remains visible in the background.

---

# Layout

```text
┌────────────────────────────────────┐
│ Select File                        │
├────────────────────────────────────┤
│ Search...                          │
├────────────────────────────────────┤
│ 📁 src                             │
│ 📁 public                          │
│ 📄 package.json                    │
│ 📄 README.md                       │
├────────────────────────────────────┤
│ Cancel                  Select     │
└────────────────────────────────────┘
```

---

# Components

The File Picker contains

- Header
- Search Field
- File List
- Directory List
- Preview (optional)
- Footer Actions

---

# Header

Displays

- Current Directory
- Navigation Path

Example

```text
workspace/src
```

---

# Search Field

Always located below the header.

Supports

- File Name
- Folder Name
- Extension
- Partial Match
- Fuzzy Match

---

# File List

Displays

- Files
- Folders

Directories appear before files.

---

# Sorting

Default order

1. Directories
2. Files

Alphabetical sorting within each group.

---

# Hidden Files

Optional.

Visibility controlled by user settings.

Example

```text
.gitignore
```

---

# File Types

Supported

- Source Code
- Markdown
- JSON
- YAML
- Images
- Archives
- Logs
- Configuration Files
- Text Files

---

# Directory Navigation

Supported.

Selecting a directory opens it.

---

# Parent Directory

Displayed as

```text
..
```

Allows navigation upward.

---

# Selection Mode

Supported

- Single File
- Multiple Files
- Directory
- Multiple Directories

---

# Multi-Selection

Users may select multiple items before confirmation.

Selection remains visible.

---

# File Preview

Optional.

Supported for

- Text
- Markdown
- JSON
- Source Code

Large files display only a preview.

---

# Unsupported Preview

Display

```text
Preview unavailable.
```

---

# Search Behavior

Search updates results immediately while typing.

No confirmation required.

---

# Search Scope

Current directory by default.

Recursive search optional.

---

# Empty Search Result

Display

```text
No matching files found.
```

---

# File Information

Optional metadata

- Size
- Modified Date
- File Type

---

# Keyboard Navigation

Supported

- Move Selection
- Open Directory
- Confirm Selection
- Cancel

---

# Touch Navigation

Supported

- Tap
- Double Tap
- Long Press

---

# Long Press

May display

- File Information
- Preview
- Context Actions

---

# Context Actions

Examples

- Open
- Select
- Copy Path
- View Details

---

# Confirmation

Selection confirmed only after explicit action.

Example

```text
Select
```

---

# Cancellation

Cancel closes the File Picker.

No changes are applied.

---

# Streaming Integration

The File Picker remains responsive during AI streaming.

---

# Tool Integration

Selected files may be passed to

- AI
- Search
- Workspace Indexer
- MCP
- Git
- Shell
- Code Editor

---

# Workspace Integration

Default location

Current workspace.

Users may navigate outside the workspace if permitted.

---

# Safe Area

The File Picker respects

- Display Cutout
- Keyboard
- Gesture Navigation

It never overlaps protected regions.

---

# Keyboard Behavior

Keyboard opening

Reduces file list height.

Search field remains visible.

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Selection state must always be readable.

---

# Performance

Large directories

Load incrementally.

Virtual rendering recommended.

---

# Error Handling

Unreadable directory

Display

```text
Access denied.
```

Missing file

Display

```text
File not found.
```

---

# Security

Respect workspace permissions.

Never display inaccessible files without authorization.

---

# Restrictions

Never

- Replace the Chat Screen
- Hide the search field
- Automatically select files
- Block typing
- Freeze during directory loading
- Display duplicate entries

---

# Example File Picker

```text
Select File

Search...

📁 src
📁 public
📄 package.json
📄 README.md

Cancel            Select
```

---

# Example Search

```text
Search: app

📄 App.tsx
📄 App.test.ts
📄 app.config.json
```

---

# Example Preview

```text
README.md

# Mobile AI CLI

Installation guide...
```

---

# File Picker Checklist

Every File Picker must

- Support search
- Support directory navigation
- Support file preview
- Support multi-selection
- Respect safe areas
- Support touch input
- Support keyboard input
- Render efficiently
- Remain responsive
- Stay terminal-native

---

# Core Rules

1. The File Picker appears as a modal overlay.
2. Directories appear before files.
3. Search updates instantly.
4. Selection requires explicit confirmation.
5. File preview is optional.
6. Large directories load incrementally.
7. Respect workspace permissions.
8. Never replace the Chat Screen.
9. Remain keyboard-safe.
10. Optimize for Android Termux.

---

# Summary

The File Picker provides a fast, searchable, and terminal-native interface for browsing and selecting files within the Mobile AI CLI. Designed specifically for Android Termux, it combines responsive search, directory navigation, optional previews, and efficient rendering while preserving workspace safety and maintaining a smooth mobile-first user experience.