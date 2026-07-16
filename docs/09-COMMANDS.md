# OpenChat CLI

# Slash Command System Specification

Version: 1.0

---

# Overview

The Slash Command System is the primary navigation system of OpenChat CLI.

Instead of navigating menus, users simply type "/".

A searchable command palette instantly appears.

Every feature inside OpenChat CLI must be accessible through slash commands.

No feature should require editing configuration files.

---

# Philosophy

Everything is a command.

Everything is searchable.

Everything is keyboard accessible.

Everything should execute within one or two keystrokes.

---

# Command Flow

User Types

/

↓

Command Palette Opens

↓

Search

↓

Arrow Keys

↓

Enter

↓

Execute

↓

Return to Chat

---

# Command Palette UI

┌──────────────────────────────────────────────┐

Search Command...

──────────────────────────────────────────────

/provider api

/add model

/all models

/agents

/skills

/session

/settings

/help

──────────────────────────────────────────────

↑↓

Navigate

Enter

Run

ESC

Close

└──────────────────────────────────────────────┘

---

# Categories

AI

Provider

Model

Agent

Skill

Workspace

Files

Session

History

Git

MCP

Tools

Settings

Theme

Developer

System

Help

---

# AI Commands

/chat

Start Chat

/new

New Conversation

/continue

Continue Session

/clear

Clear Conversation

/stop

Stop Generation

/retry

Retry Last Prompt

/summarize

Summarize Chat

/export chat

Export Conversation

---

# Provider Commands

/provider api

Add Provider

/providers

Provider Manager

/provider edit

Edit Provider

/provider delete

Delete Provider

/provider switch

Switch Provider

/provider test

Test Connection

/provider export

Export Provider

/provider import

Import Provider

/provider info

Provider Details

/provider health

Health Status

---

# Model Commands

/add model

Add Model

/all models

Model Library

/model

Quick Switch

/model edit

Edit Model

/model delete

Delete Model

/model favorite

Favorite

/model info

Details

/model export

Export

/model import

Import

---

# Agent Commands

/agents

Agent Manager

/new agent

Create Agent

/agent edit

Edit Agent

/agent clone

Clone Agent

/agent delete

Delete Agent

/agent export

Export

/agent import

Import

---

# Skill Commands

/skills

Skill Manager

/new skill

Create Skill

/skill enable

Enable

/skill disable

Disable

/skill export

Export

/skill import

Import

---

# Session Commands

/session

Current Session

/new session

Create Session

/history

History

/session rename

Rename

/session delete

Delete

/session fork

Fork

/session archive

Archive

/session export

Export

---

# Workspace Commands

/workspace

Workspace Info

/workspace open

Open Folder

/workspace reload

Reload

/workspace scan

Analyze Project

/workspace files

Project Files

/workspace git

Git Status

---

# File Commands

/file

Open File

/read

Read File

/write

Write File

/edit

Edit File

/delete

Delete File

/search

Search Files

/replace

Replace Text

/diff

View Diff

---

# Git Commands

/git status

Status

/git diff

Diff

/git commit

Commit

/git branch

Branches

/git checkout

Checkout

/git log

History

---

# Tool Commands

/tools

Tool Manager

/tool permissions

Permissions

/tool logs

Logs

/tool reload

Reload

/tool disable

Disable

/tool enable

Enable

---

# MCP Commands

/mcp

MCP Manager

/mcp connect

Connect

/mcp disconnect

Disconnect

/mcp servers

Servers

/mcp reload

Reload

---

# Settings Commands

/settings

Open Settings

/theme

Themes

/config

Configuration

/language

Language

/reset

Reset Settings

---

# Help Commands

/help

Documentation

/about

About

/version

Version

/changelog

Updates

/license

License

/report

Bug Report

---

# Developer Commands

/debug

Debug Mode

/logs

System Logs

/cache clear

Clear Cache

/database

Database

/plugins

Plugin Manager

/reload

Reload CLI

---

# Search

The command palette must support

Fuzzy Search

Partial Search

Category Search

Alias Search

Keyword Search

Examples

pro

↓

/provider

/provider api

/provider switch

mod

↓

/model

/all models

add

↓

/add model

/provider api

---

# Command Aliases

/models

↓

/all models

/provider

↓

/providers

/agent

↓

/agents

/history

↓

/session history

Aliases improve usability.

---

# Autocomplete

Typing

/prov

↓

/provider

/provider api

/provider switch

/provider edit

/provider delete

---

# Keyboard Navigation

↑

Previous

↓

Next

Enter

Run

ESC

Close

Tab

Autocomplete

Ctrl+P

Recent Commands

---

# Recent Commands

The palette remembers

Last Used

Most Used

Pinned

Favorites

Example

Recent

/provider api

/all models

/agents

Pinned

/settings

/help

---

# Favorites

Users can favorite commands.

★

/provider api

★

/all models

★

/history

Favorites appear first.

---

# Command Metadata

Each command stores

Name

Description

Category

Aliases

Permissions

Shortcut

Version

Built-in

Hidden

---

# Permissions

Some commands require confirmation.

Examples

/provider delete

/model delete

/reset

/cache clear

Confirmation Dialog

↓

Execute

---

# Error Handling

Unknown Command

↓

Did you mean

/provider api

/help

/settings

---

# Command Lifecycle

Type "/"

↓

Open Palette

↓

Search

↓

Select

↓

Validate

↓

Execute

↓

Show Result

↓

Return Focus

---

# Performance Goals

Palette Open

<20ms

Search

Instant

Autocomplete

<10ms

Execution

Immediate

Memory

Minimal

---

# Future Features

Natural Language Commands

Voice Commands

Command Macros

Command Chaining

Custom Commands

Plugin Commands

AI Suggested Commands

Workspace Commands

Cloud Command Sync

---

# Design Principles

The Slash Command System is the heart of OpenChat CLI.

A user should never wonder where a feature is located.

If they remember only one thing, it should be:

Type "/"

From there, every capability of OpenChat CLI becomes searchable, discoverable, and executable within seconds.

The command palette should feel as powerful and intuitive as the command palette in VS Code while remaining optimized for a terminal environment and comfortable to use on Android Termux.