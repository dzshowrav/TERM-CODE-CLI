# OpenChat CLI

# User Interface & User Experience Specification

Version: 1.0

---

# Design Philosophy

OpenChat CLI is **NOT** a traditional terminal program.

It behaves like a native terminal application.

The terminal itself becomes the application's window.

The interface must feel modern while remaining lightweight and keyboard-driven.

Primary platform:

Android + Termux

Secondary:

Linux

macOS

Windows

---

# UI Principles

The UI follows these rules.

✓ Chat First

✓ Mobile First

✓ Keyboard First

✓ Minimal

✓ Fast

✓ Zero Clutter

✓ Zero Sidebar

✓ Zero Mouse Dependency

✓ Everything Accessible via Slash Commands

---

# Screen Architecture

The application contains only three primary screens.

1. Home Screen

2. Chat Screen

3. Full Screen Dialogs

There are no sidebars.

There are no tabs.

There are no split layouts.

---

# Home Screen

Displayed when no active conversation exists.

Purpose

• Welcome user

• Show active configuration

• Ready to start chatting

Layout

╭──────────────────────────────────────╮
│                                      │
│           OpenChat CLI               │
│      Universal Coding Agent          │
│                                      │
│  Provider : OpenCode Zen             │
│  Model    : DeepSeek V4              │
│  Agent    : General                  │
│  Workspace: ~/projects               │
│                                      │
│  >                                  │
│                                      │
│  Press "/" for commands             │
│                                      │
╰──────────────────────────────────────╯

---

# Chat Screen

Once conversation starts.

Layout

╭──────────────────────────────────────╮
│ AI                                  │
│                                      │
│ Response...                          │
│                                      │
│ ✓ Reading files                      │
│ ✓ Editing routes                     │
│                                      │
│──────────────────────────────────────│
│ >                                    │
│──────────────────────────────────────│
│ Zen │ DeepSeek │ General │ Ready     │
╰──────────────────────────────────────╯

---

# Layout Zones

The screen has four zones.

Conversation Area

↓

Tool Activity

↓

Prompt Input

↓

Status Bar

Nothing else.

---

# Conversation Design

No bubbles.

No cards.

Plain markdown.

Example

You

Create login page.

AI

I'll generate authentication.

Reading routes...

Updating controller...

Done.

Cleaner.

Faster.

Perfect for terminals.

---

# Markdown Rendering

Support

# Heading

## Heading

Lists

Tables

Code Blocks

Bold

Italic

Links

Quotes

Task Lists

Horizontal Rules

Nested Lists

---

# Code Blocks

Example

```php

Route::get('/login');

```

Requirements

Syntax Highlighting

Line Numbers (optional)

Copy Shortcut

Scrollable

---

# Streaming UX

The response streams token by token.

Example

Thinking...

↓

Analyzing...

↓

Generating...

↓

Done

Never freeze.

Always show progress.

---

# Tool Activity

Displayed inline.

Example

✓ Read composer.json

✓ Read routes/web.php

✓ Updated AuthController.php

✓ Ran composer install

✓ Tests Passed

Every action is visible.

---

# Thinking Indicator

When AI is thinking.

Display

● Thinking...

or

Generating...

or

Analyzing Workspace...

Never show a blank screen.

---

# Prompt Input

Always fixed at bottom.

Example

>

Create Laravel API.

Supports

Multi-line

Paste

History

Autocomplete

Slash Commands

Keyboard Navigation

---

# Command Palette

Opened by typing

/

Example

Search Commands...

/provider api

/add model

/all models

/settings

/history

/help

Arrow Keys

↓

Select

↓

Enter

Command executes.

---

# Dialog System

Every management screen opens as a fullscreen dialog.

Never use floating windows.

Examples

Provider Manager

Model Manager

Settings

History

Permissions

Help

Agents

Skills

---

# Provider Dialog

/provider api

Layout

Add Provider

Provider Name

[____________]

Base URL

[____________]

API Key

[************]

Save

Cancel

---

# Provider Manager

/providers

Layout

Providers

OpenCode Zen

✓ Connected

OpenAI

Offline

Groq

Connected

Actions

Add

Edit

Delete

Test

Export

Import

---

# Model Dialog

/add model

Layout

Model ID

[deepseek-v4]

Provider

▼ OpenCode Zen

Save

---

# Model Manager

/all models

Layout

Models

● DeepSeek V4

OpenCode Zen

○ GPT-5.5

OpenAI

○ Gemini

OpenRouter

Enter

↓

Active Model changes instantly.

---

# Agent Dialog

/agents

General

Laravel

React

Flutter

Python

Security

Reviewer

Architect

Debugger

---

# Skills Dialog

/skills

Laravel

Bootstrap

React

Docker

Git

Security

Performance

Testing

---

# Session Dialog

/history

Today

Laravel Login

React Project

Yesterday

Docker Setup

Archive

Rename

Delete

Export

Fork

---

# Settings Dialog

/settings

Appearance

Theme

Streaming

Context

Permissions

Default Agent

Default Model

Notifications

Language

Advanced

---

# Permission Dialog

Whenever a dangerous tool is requested.

Example

Execute Bash?

composer install

Allow Once

Always Allow

Deny

Never execute automatically.

---

# Search UX

Search appears everywhere.

Example

Search Provider...

Search Models...

Search Sessions...

Search Skills...

Instant filtering.

---

# Empty States

No Providers

No provider configured.

Run

/provider api

No Models

No models available.

Run

/add model

No Sessions

Start a new conversation.

---

# Error States

Connection Failed

Provider Offline

Invalid API Key

Model Not Found

Tool Failed

Every error should explain

Why

How to Fix

---

# Notifications

Small inline notifications.

✓ Provider Saved

✓ Model Added

✓ Session Exported

✓ Connected

Avoid intrusive popups.

---

# Status Bar

Always visible.

Example

Workspace

Git Branch

Provider

Model

Agent

Context

Version

Example

~/dev-store

main

Zen

DeepSeek

Laravel

46%

v1.0

---

# Keyboard Shortcuts

/

Open Command Palette

ESC

Close Dialog

Enter

Confirm

Ctrl+C

Interrupt AI

Ctrl+L

Clear Screen

Ctrl+R

Retry

Ctrl+S

Save Session

Ctrl+N

New Session

Ctrl+H

History

---

# Colors

Dark Theme

Primary

Cyan

Success

Green

Warning

Yellow

Error

Red

Selection

Blue

Disabled

Gray

High contrast for AMOLED displays.

---

# Typography

Use terminal monospace only.

Never mix fonts.

Keep spacing consistent.

---

# Animations

Very subtle.

Loading spinner

Streaming cursor

Progress dots

No heavy animations.

---

# Accessibility

Readable colors

Keyboard only

Screen reader friendly

Large terminal compatibility

No color-only indicators

Icons always have text.

---

# Mobile Optimization

Designed for

80 columns

100 columns

Portrait Mode

Landscape Mode

No horizontal scrolling.

---

# UX Goals

Users should never need documentation to perform common tasks.

Everything should be discoverable through the "/" command palette.

The interface should remain clean, distraction-free, and transparent, showing exactly what the AI is doing while keeping the user's focus on the conversation and code.