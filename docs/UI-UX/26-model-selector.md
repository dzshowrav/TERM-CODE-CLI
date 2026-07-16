# 26-model-selector.md

# Model Selector
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Model Selector system used throughout the Mobile AI CLI.

The Model Selector allows users to view, search, switch, configure, and manage AI models from different providers without leaving the current conversation workflow.

The system must provide a fast, terminal-native, mobile-first model switching experience optimized for Android Termux.

---

# Design Goals

The Model Selector must be

- Mobile First
- Terminal Native
- Fast
- Searchable
- Provider Independent
- Keyboard Friendly
- Touch Friendly
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

Model switching should be simple and transparent.

Users should always understand

- Which model is active
- Which provider it belongs to
- What capabilities it supports
- What limitations exist

---

# Display Location

The Model Selector appears as a modal overlay.

The Conversation Area remains visible behind it.

---

# Layout

```text
┌────────────────────────────────┐
│ Select Model                   │
├────────────────────────────────┤
│ Search models...               │
├────────────────────────────────┤
│ ✓ GPT-5                        │
│   OpenAI                       │
│   Coding · Reasoning           │
│                                │
│   Claude                       │
│   Anthropic                    │
│   Long Context                 │
│                                │
│   Gemini                       │
│   Google                       │
├────────────────────────────────┤
│ Cancel                         │
└────────────────────────────────┘
```

---

# Components

The Model Selector contains

- Header
- Search Input
- Provider List
- Model List
- Model Details
- Action Area

---

# Header

Displays

```text
Select Model
```

---

# Search Input

Always visible.

Supports

- Model Name Search
- Provider Search
- Capability Search
- Fuzzy Search

---

# Provider Grouping

Models may be grouped by provider.

Example

```text
OpenAI

Anthropic

Google

Local Models
```

---

# Model Item

Each model entry contains

- Model Name
- Provider
- Capability
- Status
- Context Size (optional)

---

# Example

```text
GPT-5

OpenAI

Coding + Reasoning

Active
```

---

# Active Model

The current model is clearly marked.

Example

```text
✓ GPT-5
```

---

# Model Status

Supported states

```
Available

Active

Loading

Unavailable

Error
```

---

# Available State

Example

```text
GPT-5

Available
```

---

# Active State

Example

```text
✓ GPT-5

Current Model
```

---

# Loading State

Example

```text
Loading Model...
```

---

# Unavailable State

Example

```text
Model unavailable.
```

---

# Error State

Example

```text
Connection failed.
```

---

# Model Details

Optional expandable information.

Contains

- Provider
- Version
- Context Limit
- Features
- Cost Information
- Local/Cloud Status

---

# Capability Tags

Examples

```
Coding

Reasoning

Vision

Fast

Long Context

Local
```

---

# Local Models

Local models display separately.

Example

```text
Local Models

Llama

Qwen

Mistral
```

---

# Cloud Models

Cloud models display provider information.

Example

```text
Cloud Models

GPT

Claude

Gemini
```

---

# Selection

Supported

- Keyboard Navigation
- Touch Selection

Only one active model at a time.

---

# Confirmation

Model switching requires explicit selection.

---

# Switching Flow

```
User Selects Model

↓

Loading

↓

Connection Check

↓

Model Activated

↓

Toast Notification
```

---

# Success Notification

Example

```text
Model switched to GPT-5.
```

---

# Failure Notification

Example

```text
Unable to switch model.
```

---

# Search Ranking

Results ranked by

1. Exact Match
2. Provider Match
3. Capability Match
4. Recent Usage

---

# Recent Models

Frequently used models appear first.

---

# Favorite Models

Optional.

Users may pin preferred models.

---

# Keyboard Navigation

Supported

- Move Selection
- Search
- Confirm
- Cancel

---

# Touch Navigation

Supported

- Tap Model
- Scroll List
- Expand Details

---

# Long Press

Optional actions

- View Details
- Set Default
- Remove Provider

---

# Streaming Integration

Model switching does not interrupt completed conversations.

If streaming is active

The system may require confirmation before switching.

---

# Context Integration

Changing models may affect

- Context Limit
- Tool Support
- Response Style

The user should be informed.

---

# Tool Integration

Different models may support different tools.

Example

```text
GPT-5

Tools Supported

MCP

Terminal

File Access
```

---

# Permission Integration

Model providers requiring credentials must request permission.

---

# Safe Area

The Model Selector respects

- Display Cutout
- Keyboard
- Status Bar
- Command Input

Never overlap protected areas.

---

# Keyboard Behavior

When keyboard opens

Search remains active.

Model list height adjusts.

---

# Accessibility

Support

- High Contrast
- Large Fonts
- Screen Readers

Model information must remain readable.

---

# Performance

Large model lists

Use

- Incremental filtering
- Virtual rendering
- Cached provider data

---

# Error Handling

Provider unavailable

Display

```text
Provider unavailable.
```

Invalid configuration

Display

```text
Configuration required.
```

---

# Security

Never display

- API Keys
- Tokens
- Private Credentials

Only safe provider information is shown.

---

# Restrictions

Never

- Switch models silently
- Hide active model
- Remove context without warning
- Block user interaction
- Display unavailable models as active

---

# Example Model Selection

```text
Select Model

Search...

✓ GPT-5
  OpenAI
  Coding

  Claude
  Anthropic
  Long Context

  Gemini
  Google

Cancel
```

---

# Example Switching

```text
GPT-5

↓

Loading...

↓

Active
```

---

# Model Selector Checklist

Every Model Selector must

- Show active model
- Support search
- Support providers
- Support local models
- Show capabilities
- Handle switching safely
- Respect context changes
- Protect credentials
- Support keyboard and touch
- Remain terminal-native

---

# Core Rules

1. Active model is always visible.
2. Model switching requires user action.
3. Search is always available.
4. Providers are clearly identified.
5. Capabilities are displayed.
6. Context changes are handled safely.
7. Credentials remain hidden.
8. Loading states are visible.
9. Never interrupt user workflow.
10. Optimize for Android Termux.

---

# Summary

The Model Selector provides a fast and transparent way to manage AI models inside the Mobile AI CLI. With provider grouping, fuzzy search, capability information, safe switching, and terminal-native rendering, it enables users to choose the best AI model for coding, reasoning, and automation tasks while maintaining a mobile-first Android Termux experience.