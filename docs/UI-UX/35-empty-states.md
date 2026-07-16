# 35-empty-states.md

# Empty States
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

This document defines the Empty State System used throughout the Mobile AI CLI.

The Empty State System describes how the application communicates when no data, content, result, or resource is currently available.

The system must guide users toward the next meaningful action instead of showing blank screens.

---

# Design Goals

The Empty State System must be

- Mobile First
- Terminal Native
- Informative
- Action Oriented
- Minimal
- Fast
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

An empty state is not a missing state.

It is an opportunity to guide the user.

Every empty screen should answer:

- What is empty?
- Why is it empty?
- What can the user do next?

---

# Empty State Architecture

```
No Data Detected

↓

Empty State Resolver

↓

Message Selection

↓

Action Suggestions

↓

User Action
```

---

# Empty State Types

Supported

- New Conversation
- No Sessions
- No Files
- No Search Results
- No Models
- No Tools
- No Workspace
- No History
- No Network Data
- No Permissions

---

# Empty State Structure

Every empty state contains

- Icon
- Title
- Description
- Action
- Optional Help

---

# Icon Rules

Icons must be

- Professional
- Simple
- Terminal compatible
- Text-free

Never use emoji as icons.

---

# Title Rules

Title must be short.

Example

Good

```text
No Sessions
```

Bad

```text
There are currently no sessions available in this application
```

---

# Description Rules

Description explains the situation.

Example

```text
Create a new session to start working with AI.
```

---

# Action Rules

Every useful empty state should provide a next action.

Examples

```
Create Session

Add Workspace

Connect Provider

Search Again
```

---

# New Conversation Empty State

Displayed when no messages exist.

Example

```text
No Conversation

Start a new AI conversation.
```

Action

```text
Start Chat
```

---

# First Launch Empty State

Displayed when application starts for the first time.

Example

```text
Welcome

Configure your AI environment.
```

Actions

```
Setup

Continue
```

---

# Session Empty State

When no saved sessions exist.

Example

```text
No Sessions

Create your first workspace session.
```

Action

```text
New Session
```

---

# Workspace Empty State

When no project folder is connected.

Example

```text
No Workspace

Connect a project folder to begin.
```

Action

```text
Select Workspace
```

---

# File Picker Empty State

When a folder contains no files.

Example

```text
No Files

This folder is empty.
```

Action

```text
Choose Another Folder
```

---

# Search Empty State

When search returns no results.

Example

```text
No Results

Try another search term.
```

Action

```text
Clear Search
```

---

# Model Empty State

When no AI model is configured.

Example

```text
No Models

Add an AI provider first.
```

Action

```text
Configure Provider
```

---

# Tool Empty State

When no tools are available.

Example

```text
No Tools

Enable MCP tools to continue.
```

Action

```text
Manage Tools
```

---

# MCP Empty State

When no MCP servers exist.

Example

```text
No MCP Servers

Connect an MCP server.
```

Action

```text
Add Server
```

---

# History Empty State

When conversation history is empty.

Example

```text
No History

Completed conversations appear here.
```

---

# Notification Empty State

When no notifications exist.

Example

```text
No Notifications

Everything is up to date.
```

---

# Error History Empty State

When no errors are stored.

Example

```text
No Errors

No problems recorded.
```

---

# Layout

Example

```text
┌──────────────────────────┐
│                          │
│          [Icon]          │
│                          │
│       No Sessions        │
│                          │
│ Create your first        │
│ AI workspace session.    │
│                          │
│     [New Session]        │
│                          │
└──────────────────────────┘
```

---

# Positioning

Empty states should appear

- Centered vertically
- Centered horizontally
- Above keyboard area

---

# Keyboard Behavior

If an empty state contains input

Keyboard opens normally.

Example

```
No Results

[Search Input]
```

---

# Safe Area

Empty states respect

- Status Bar
- Command Input
- Keyboard
- Navigation Area

---

# Loading Transition

Empty states should transition smoothly.

Example

```
No Data

↓

Loading

↓

Content Available
```

---

# Dynamic Updates

When data appears

Empty state disappears automatically.

---

# Animation

Allowed

- Fade transition
- Minimal appearance animation

Avoid

- Large movement
- Decorative effects

---

# Accessibility

Support

- Screen Readers
- High Contrast
- Large Text

All empty states must remain understandable without icons.

---

# Performance

Empty states should

- Load instantly
- Use minimal rendering
- Require no heavy resources

---

# Offline Behavior

When offline

Show useful guidance.

Example

```text
No Connection

Reconnect to load remote data.
```

---

# Permission Empty State

When permission is missing.

Example

```text
No File Access

Grant permission to view files.
```

---

# Security

Never show sensitive information in empty messages.

---

# Restrictions

Never

- Show blank screens
- Use confusing messages
- Hide possible actions
- Display fake content
- Use emoji as visual indicators
- Block user progress

---

# Example Workflow

```
Open App

↓

No Workspace

↓

Select Folder

↓

Workspace Loaded
```

---

# Example Search Workflow

```
Search

↓

No Results

↓

Change Query

↓

Results Found
```

---

# Empty State Checklist

Every Empty State System must

- Explain missing content
- Provide next action
- Use clear language
- Support accessibility
- Respect safe areas
- Work offline
- Remain lightweight
- Support keyboard interaction
- Work in terminal UI
- Optimize Android Termux

---

# Core Rules

1. Never show a completely blank screen.
2. Every empty state explains the situation.
3. Every important empty state provides an action.
4. Messages must be short and clear.
5. Icons support text, not replace it.
6. Empty states must load instantly.
7. User should always know the next step.
8. Keyboard behavior must remain consistent.
9. Privacy information must stay hidden.
10. Optimize for mobile CLI workflows.

---

# Summary

The Empty State System ensures that the Mobile AI CLI remains helpful even when no data exists. Through clear explanations, useful actions, terminal-friendly layouts, and mobile-first behavior, empty screens become guidance points that help users continue their workflow efficiently on Android Termux.