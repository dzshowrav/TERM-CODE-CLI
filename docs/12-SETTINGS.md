# OpenChat CLI

# Settings System Specification

Version: 1.0

---

# Overview

The Settings System controls every configurable behavior of OpenChat CLI.

Users should never need to manually edit configuration files.

All settings must be manageable through an interactive fullscreen interface and slash commands.

The system supports

• Global Settings

• Workspace Settings

• Session Settings

• Runtime Overrides

---

# Design Goals

✓ Zero JSON Editing

✓ Searchable

✓ Mobile Friendly

✓ Instant Apply

✓ Workspace Override

✓ Session Override

✓ Import / Export

✓ Reset Support

---

# Configuration Priority

Default Values

↓

Global Settings

↓

Workspace Settings

↓

Session Settings

↓

Runtime Changes

Higher levels override lower levels.

---

# Database

settings

Columns

key

value

scope

workspace_id

session_id

updated_at

---

# Settings Categories

General

Appearance

AI

Providers

Models

Agents

Skills

Chat

Streaming

Workspace

Tools

Permissions

MCP

Plugins

Security

Performance

Advanced

Developer

---

# Settings Screen

Command

/settings

Layout

Search Settings...

────────────────────────

General

Appearance

AI

Providers

Models

Agents

Skills

Chat

Streaming

Workspace

Tools

Permissions

Security

Advanced

Developer

────────────────────────

Enter

Open

ESC

Back

---

# Search

Users can search

Setting Name

Description

Category

Keyword

Example

theme

↓

Theme

stream

↓

Streaming

permission

↓

Permissions

---

# General

Language

English

বাংলা (Future)

Timezone

Auto

Startup Screen

Home

Resume Last Session

Check Updates

Auto

Notifications

Enabled

Confirm Exit

Enabled

---

# Appearance

Theme

Dark

Light

System

Accent Color

Cyan

Green

Blue

Purple

Orange

Status Bar

Visible

Animations

Enabled

Compact Mode

Disabled

Line Numbers

Enabled

Markdown Rendering

Enhanced

Syntax Highlighting

Enabled

---

# AI

Default Provider

Selectable

Default Model

Selectable

Default Agent

Selectable

Auto Load Skills

Enabled

Auto Detect Project

Enabled

Reasoning

Balanced

Temperature

0.7

Max Output Tokens

Auto

Context Compression

Enabled

---

# Providers

Auto Test Connection

Enabled

Health Check Interval

5 Minutes

Default Timeout

30 Seconds

Retry Attempts

3

Auto Reconnect

Enabled

Fallback Provider

Optional

---

# Models

Auto Discover

Disabled

Show Favorites First

Enabled

Remember Last Model

Enabled

Validate Before Use

Enabled

---

# Agents

Remember Last Agent

Enabled

Auto Switch

Disabled

Preferred Agent

General

---

# Skills

Auto Detect

Enabled

Load Dependencies

Enabled

Lazy Loading

Enabled

Maximum Skills

20

---

# Chat

Auto Title

Enabled

Auto Save

Enabled

Message Timestamp

Visible

Markdown

Enabled

Code Highlighting

Enabled

Conversation Summary

Enabled

Show Token Usage

Enabled

---

# Streaming

Streaming

Enabled

Typing Cursor

Enabled

Thinking Indicator

Enabled

Tool Animation

Enabled

Render Speed

Real Time

First Token Timeout

15 Seconds

---

# Workspace

Auto Scan

Enabled

Git Integration

Enabled

Index Files

Enabled

Watch File Changes

Enabled

Ignored Directories

node_modules

vendor

.git

build

dist

coverage

tmp

.cache

---

# Tools

Allow Read Tools

Always

Allow Write Tools

Ask

Allow Bash

Ask

Allow Git

Ask

Allow Delete

Ask

Parallel Execution

Enabled

Execution Timeout

60 Seconds

---

# Permissions

Default Mode

Ask

Remember Decisions

Enabled

Reset Permissions

Button

Dangerous Actions

Require Confirmation

Always

---

# MCP

Auto Connect

Enabled

Reconnect

Enabled

Health Checks

Enabled

Discover Tools

Enabled

Discover Prompts

Enabled

Discover Resources

Enabled

---

# Plugins

Auto Load

Enabled

Plugin Updates

Manual

Safe Mode

Enabled

Allow Third-Party Plugins

Disabled

---

# Security

Mask API Keys

Enabled

Encrypt Local Secrets

Future

Hide Logs

Enabled

Secure Clipboard

Enabled

Clear Clipboard on Exit

Disabled

Session Lock

Disabled

---

# Performance

Memory Optimization

Enabled

Database Cache

Enabled

Parallel Requests

Enabled

Maximum Concurrent Tools

5

Workspace Cache

Enabled

Provider Cache

Enabled

---

# Advanced

Debug Logging

Disabled

Verbose Mode

Disabled

Experimental Features

Disabled

Database Maintenance

Manual

Clear Cache

Button

Vacuum Database

Button

---

# Developer

Developer Mode

Disabled

Show Internal Events

Disabled

Show Raw API Requests

Disabled

Show Raw Responses

Disabled

Performance Overlay

Disabled

Export Debug Logs

Button

---

# Theme Selection

/theme

Displays

Dark

Light

Midnight

OLED

Solarized

Nord

Dracula

GitHub Dark

Monokai

Custom

---

# Reset Settings

/reset settings

Options

Reset Current Category

Reset Workspace

Reset Global Settings

Factory Reset

Confirmation Required

---

# Import Settings

/settings import

Supports

JSON

Future

YAML

TOML

---

# Export Settings

/settings export

Formats

JSON

Markdown

Encrypted Backup (Future)

---

# Validation

Every setting is validated before saving.

Invalid values are rejected with a clear explanation.

---

# Live Updates

Most settings apply instantly.

Some settings require

Workspace Reload

or

Application Restart

The UI must clearly indicate which behavior applies.

---

# Performance Goals

Settings Open

<30ms

Search

Instant

Save

<10ms

Theme Switch

Immediate

---

# Future Features

Cloud Settings Sync

Multiple Profiles

Workspace Templates

Organization Policies

Settings Marketplace

Import from OpenCode

Import from Claude Code

Import from Gemini CLI

---

# Design Principles

The Settings System should make OpenChat CLI fully customizable without exposing users to configuration files.

Every option should be searchable, well documented, and immediately understandable.

Users should feel confident changing behavior from the CLI itself, whether they are on Android Termux, Linux, macOS, or Windows, while advanced users retain fine-grained control over every aspect of the application.