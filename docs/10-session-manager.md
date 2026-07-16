# 10-session-manager.md

# Session Manager Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Session Manager?

A **Session Manager** is responsible for creating, maintaining, restoring, and ending an AI conversation session.

A session represents the complete working state of the application.

Unlike chat history, a session stores everything required to continue work exactly where the user stopped.

---

# Why Session Manager?

Without Session Manager

```
User

↓

Close CLI

↓

Everything Lost
```

Problems

- Conversation lost
- Context lost
- Open files lost
- Running tasks lost
- User preferences forgotten

---

With Session Manager

```
Start Session

↓

Work

↓

Save

↓

Exit

↓

Restore

↓

Continue
```

---

# Goals

A production Session Manager should provide

- Session persistence
- Auto save
- Crash recovery
- Workspace restoration
- Conversation history
- Context preservation
- Multi-session support
- Fast loading
- Compression
- Secure storage

---

# High-Level Architecture

```
                 User
                   │
                   ▼
            Session Manager
                   │
      ┌────────────┼─────────────┐
      ▼            ▼             ▼
 Session Store  State Manager  History
      │            │             │
      └──────┬─────┴──────┬──────┘
             ▼            ▼
        Cache Layer   Event Bus
             │
             ▼
       Agent / Renderer
```

---

# Folder Structure

```
src/

session/

    SessionManager.ts

    Session.ts

    SessionStore.ts

    SessionLoader.ts

    SessionSaver.ts

    SessionCache.ts

    SessionHistory.ts

    SessionSnapshot.ts

    SessionRecovery.ts

    SessionValidator.ts

    SessionEvents.ts

    SessionMetrics.ts

    SessionConfig.ts
```

---

# Core Components

## Session Manager

Central controller.

Responsibilities

- Create session
- Save session
- Restore session
- Delete session
- Switch session

---

## Session Store

Responsible for persistent storage.

Stores

```
JSON

SQLite

Binary

Encrypted Files
```

---

## Session Loader

Loads

- Conversation
- Context
- Open files
- Settings
- Active workspace

---

## Session Saver

Writes session safely.

Supports

- Incremental save
- Atomic save
- Auto save

---

## Session Cache

Stores current session in memory.

Benefits

- Fast access
- Reduced disk reads
- Better performance

---

## Session Snapshot

Creates restore points.

Example

```
Snapshot 1

↓

Snapshot 2

↓

Snapshot 3
```

Useful for recovery.

---

## Session Recovery

Handles

```
Crash

↓

Recover Latest Session
```

---

## Session Validator

Checks

- Version
- Corruption
- Missing fields
- Compatibility

---

# Session Lifecycle

```
Application Starts

↓

Create/Open Session

↓

Restore State

↓

Work

↓

Auto Save

↓

Exit

↓

Persist Session
```

---

# Session Object

Contains

```
Session ID

Workspace

Conversation

Agent State

Model

Settings

Timestamp

Version

Metadata
```

---

# Session States

```
New

↓

Active

↓

Saving

↓

Idle

↓

Closed

↓

Archived
```

---

# Auto Save Flow

```
State Changed

↓

Debounce

↓

Save Queue

↓

Write To Disk

↓

Done
```

---

# Restore Flow

```
Application Starts

↓

Read Session

↓

Validate

↓

Restore Workspace

↓

Restore Conversation

↓

Ready
```

---

# Multi-Session Support

Example

```
Session A

Laravel Project

--------------

Session B

React Project

--------------

Session C

Research
```

Each session remains isolated.

---

# Conversation Storage

Stores

```
System Messages

User Messages

Assistant Messages

Tool Calls

Tool Results
```

---

# Workspace State

Stores

```
Current Directory

Open Files

Selected File

Cursor Position

Scroll Position
```

---

# Agent State

Stores

```
Current Model

Loaded Skills

Loaded Plugins

Pending Tasks

Memory
```

---

# Configuration

Stores

```
Theme

Language

Permissions

Window Layout

Preferences
```

---

# Event Bus Integration

Common events

```
session:create

session:load

session:save

session:restore

session:close

session:error
```

---

# Agent Integration

```
Session Loaded

↓

Restore Context

↓

Continue Conversation
```

---

# Renderer Integration

```
Session Restored

↓

Restore UI Layout

↓

Render
```

---

# Plugin Integration

Plugins may save

```
Plugin Settings

Plugin Cache

Plugin State
```

---

# Skills Integration

Skills may restore

```
Loaded Skills

Skill Configuration

Cached Templates
```

---

# File Watcher Integration

```
Workspace Restored

↓

Restart Watching
```

---

# MCP Integration

Reconnect

```
Configured Servers

↓

Authentication

↓

Restore Ready State
```

---

# Crash Recovery

```
Unexpected Exit

↓

Recovery Manager

↓

Latest Snapshot

↓

Restore
```

---

# Version Migration

If session version differs

```
Old Version

↓

Migration

↓

Current Version
```

---

# Storage Layout

Example

```
sessions/

    active/

    archived/

    backups/

    cache/

    snapshots/
```

---

# Backup Strategy

Automatically create

```
Daily Backup

Manual Snapshot

Before Upgrade Backup
```

---

# Security

Always

- Encrypt sensitive data
- Validate session files
- Restrict file access
- Backup before overwrite

Never

- Store API keys in plain text
- Trust corrupted sessions
- Ignore validation failures

---

# Performance Optimizations

Use

- Incremental saving
- Compression
- Lazy loading
- Memory cache
- Background saving
- Snapshot reuse

Avoid

- Saving entire session every keystroke
- Blocking UI while saving
- Duplicate serialization

---

# Best Practices

Always

- Auto save
- Validate sessions
- Keep backups
- Separate cache from storage
- Version session format
- Support migration

Never

- Assume sessions are valid
- Overwrite without backup
- Block startup unnecessarily

---

# Common Mistakes

Bad

```
Exit

↓

Lose Everything
```

---

Good

```
State Changed

↓

Auto Save

↓

Crash Recovery

↓

Continue Later
```

---

# Testing Checklist

- Session creation
- Auto save
- Manual save
- Restore
- Multi-session
- Crash recovery
- Backup
- Migration
- Corrupted session
- Version compatibility

---

# Advantages

- Persistent workflow
- Faster startup
- Better user experience
- Crash recovery
- Multi-project support
- Reduced context rebuilding
- Seamless continuation

---

# Disadvantages

- Storage management
- Migration complexity
- Backup overhead
- Corruption handling

---

# Used In

- OpenCode
- Antigravity CLI
- VS Code
- Cursor
- Zed
- JetBrains IDEs
- Claude Code
- Continue.dev

---

# Complete Session Flow

```
User Opens CLI

↓

Session Manager

↓

Load Previous Session

↓

Restore Workspace

↓

Restore Conversation

↓

Reconnect MCP

↓

Load Skills

↓

Load Plugins

↓

Restart File Watcher

↓

Restore UI

↓

Ready
```

---

# Summary

The **Session Manager** preserves the complete working state of an AI Coding Agent so users can resume work without losing context.

A production-grade Session Manager should include:

- Session Manager
- Session Store
- Loader
- Saver
- Cache
- Snapshot System
- Recovery Manager
- Validator
- Version Migration
- Auto Save
- Event Bus Integration

By combining persistent storage, automatic recovery, version-aware migration, and integration with the Agent Engine, Renderer, File Watcher, MCP, Plugins, and Skills, the Session Manager provides a reliable and uninterrupted development experience comparable to modern tools like OpenCode and Antigravity CLI.