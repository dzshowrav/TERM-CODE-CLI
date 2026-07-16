# 28-settings-screen.md

# Settings Screen
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Settings Screen system used throughout the Mobile AI CLI.

The Settings Screen provides a centralized location for configuring application behavior, AI preferences, interface options, tools, providers, permissions, storage, and developer settings.

The system must remain terminal-native while providing a mobile-first configuration experience optimized for Android Termux.

---

# Design Goals

The Settings Screen must be

- Mobile First
- Terminal Native
- Organized
- Searchable
- Safe
- Accessible
- Keyboard Friendly
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

Settings should be discoverable but never overwhelming.

Users should quickly find what they need while advanced options remain available.

---

# Display Location

The Settings Screen opens as a dedicated application view.

Conversation state remains preserved.

---

# Layout

```text
┌──────────────────────────────┐
│ Settings                     │
├──────────────────────────────┤
│ Search Settings...           │
├──────────────────────────────┤
│ AI                          │
│ Appearance                  │
│ Models                      │
│ Tools                       │
│ Permissions                 │
│ Workspace                   │
│ Storage                     │
│ Advanced                    │
└──────────────────────────────┘
```

---

# Settings Categories

Supported categories

- General
- Appearance
- AI Configuration
- Model Management
- Tools
- MCP
- Permissions
- Workspace
- Storage
- Notifications
- Keyboard
- Advanced

---

# Search Settings

Always available.

Supports

- Setting Name Search
- Category Search
- Keyword Search

Example

```text
Search: theme
```

Results

```text
Appearance

Theme Settings
```

---

# General Settings

Contains

- Application Name
- Default Session
- Startup Behavior
- Language
- Default Workspace

---

# Appearance Settings

Contains

- Theme
- Color Scheme
- Font Size
- Terminal Density
- Animation Level

---

# Theme Settings

Supported

```
Dark

Light

High Contrast

Custom
```

---

# Font Settings

Options

- Font Size
- Line Height
- Terminal Font Preference

---

# Density Settings

Options

```
Compact

Normal

Comfortable
```

---

# AI Settings

Contains

- Default Model
- Response Style
- Streaming
- Context Management

---

# Streaming Settings

Options

```
Enabled

Disabled
```

---

# Context Settings

Contains

- Context Limit
- Auto Compression
- Memory Usage

---

# Model Settings

Contains

- Providers
- Models
- API Configuration
- Default Model

---

# Provider Settings

Examples

```text
OpenAI

Anthropic

Google

Local
```

---

# Tool Settings

Contains

- Enabled Tools
- Tool Permissions
- Execution Rules

---

# MCP Settings

Contains

- MCP Servers
- Connection Status
- Server Permissions

---

# Permission Settings

Controls

- File Access
- Shell Execution
- Network Access
- External Tools

---

# Workspace Settings

Contains

- Default Folder
- Indexing
- Ignore Rules
- File Watching

---

# Storage Settings

Contains

- Cache
- Session Data
- Logs
- Temporary Files

---

# Notification Settings

Controls

- Toast Notifications
- Error Alerts
- Completion Alerts

---

# Keyboard Settings

Contains

- Shortcuts
- Key Bindings
- Input Behavior

---

# Advanced Settings

Contains

- Debug Mode
- Logging Level
- Developer Options

---

# Setting Types

Supported controls

- Toggle
- Select
- Input
- Slider
- Action Button

---

# Toggle Example

```text
Streaming

[ ON ]
```

---

# Select Example

```text
Theme

Dark
```

---

# Input Example

```text
Workspace

/home/project
```

---

# Slider Example

```text
Font Size

██████░░░░
```

---

# Action Button Example

```text
Clear Cache
```

---

# Change Confirmation

Important changes require confirmation.

Examples

- Delete data
- Reset settings
- Change permissions

---

# Save Behavior

Supported modes

```
Auto Save

Manual Save
```

Default

```
Auto Save
```

---

# Reset Settings

Available.

Requires confirmation.

Example

```text
Reset Settings?

All preferences will be restored.

Cancel     Reset
```

---

# Import Settings

Supported.

Formats

- JSON
- Config File

---

# Export Settings

Supported.

Allows backup of preferences.

---

# Settings Sync

Optional.

Can synchronize across devices.

---

# Security

Never display

- API Keys
- Tokens
- Passwords

Sensitive values must be hidden.

---

# Password Fields

Display

```text
••••••••
```

---

# Keyboard Behavior

When keyboard opens

- Inputs remain visible
- Screen adjusts automatically
- Scrolling remains available

---

# Safe Area

Settings Screen respects

- Status Bar
- Command Input
- Keyboard
- Navigation Area

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

---

# Performance

Settings should load instantly.

Large configuration sections load lazily.

---

# Error Handling

Invalid setting

Display

```text
Invalid value.
```

---

# Failed Save

Display

```text
Unable to save settings.
```

---

# Restrictions

Never

- Apply dangerous changes silently
- Expose secrets
- Reset without confirmation
- Block navigation
- Lose configuration data

---

# Example Settings Flow

```text
Settings

↓

AI

↓

Model

↓

GPT-5

↓

Save
```

---

# Example Permission Flow

```text
Settings

↓

Permissions

↓

Shell Access

↓

Enable
```

---

# Settings Checklist

Every Settings Screen must

- Support search
- Group related options
- Protect sensitive data
- Support reset
- Support import/export
- Save reliably
- Respect safe areas
- Support keyboard input
- Remain accessible
- Stay terminal-native

---

# Core Rules

1. Settings are grouped logically.
2. Search is always available.
3. Sensitive data remains hidden.
4. Important changes require confirmation.
5. Settings save safely.
6. Reset actions require approval.
7. Keyboard never blocks inputs.
8. Configuration remains recoverable.
9. UI remains lightweight.
10. Optimize for Android Termux.

---

# Summary

The Settings Screen provides complete control over the Mobile AI CLI environment, including AI models, tools, appearance, permissions, storage, and advanced configuration. With a searchable structure, safe modification flow, and terminal-native mobile-first design, it allows users to customize the application while maintaining reliability and performance on Android Termux.