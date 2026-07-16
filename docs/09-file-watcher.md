# 09-file-watcher.md

# File Watcher Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a File Watcher?

A **File Watcher** is a background service that continuously monitors the filesystem for changes and notifies the application whenever files or directories are created, modified, renamed, or deleted.

In an AI Coding Agent, the File Watcher keeps the Agent synchronized with the user's workspace without requiring manual refreshes.

---

# Why File Watcher?

Without File Watcher

```
User edits file

↓

Agent doesn't know

↓

Stale context

↓

Wrong response
```

Problems

- Outdated workspace context
- Manual refresh required
- Slow feedback
- Broken cache
- Poor developer experience

---

With File Watcher

```
Filesystem

↓

File Watcher

↓

Event Bus

↓

Agent

↓

Renderer

↓

Workspace Updated
```

---

# Goals

A production File Watcher should provide

- Real-time monitoring
- Recursive directory watching
- Cross-platform support
- Debouncing
- Event filtering
- Ignore rules
- Symbolic link handling
- Low CPU usage
- Cache synchronization
- Event-driven architecture

---

# High-Level Architecture

```
               Workspace

                   │

                   ▼

          Operating System

                   │

                   ▼

           File Watcher

                   │

        ┌──────────┼──────────┐

        ▼          ▼          ▼

   Event Queue  Ignore Filter Cache

        │          │          │

        └─────┬────┴─────┬────┘

              ▼          ▼

          Event Bus   Metrics

              │

              ▼

      Agent / TUI / Plugins
```

---

# Folder Structure

```
src/

watcher/

    FileWatcher.ts

    WatchManager.ts

    WatchSession.ts

    FileEvent.ts

    EventQueue.ts

    IgnoreManager.ts

    PathResolver.ts

    RecursiveWatcher.ts

    Debouncer.ts

    EventDispatcher.ts

    CacheSynchronizer.ts

    Snapshot.ts

    WatchMetrics.ts

    WatchEvents.ts

    WatchValidator.ts
```

---

# Core Components

## File Watcher

Central monitoring service.

Responsibilities

- Watch filesystem
- Receive OS notifications
- Forward events
- Handle lifecycle

---

## Watch Manager

Controls

- Start watching
- Stop watching
- Restart watching
- Multiple workspaces

---

## Recursive Watcher

Automatically watches

```
project/

src/

components/

pages/

config/

...
```

Including newly created folders.

---

## Ignore Manager

Ignores unnecessary files.

Examples

```
node_modules/

.git/

.cache/

dist/

build/

vendor/

coverage/
```

---

## Event Queue

Buffers filesystem events.

Prevents

```
1000 File Events

↓

1000 UI Updates
```

Instead

```
Queue

↓

Batch

↓

Dispatch
```

---

## Debouncer

Merges repeated events.

Example

```
save()

↓

change

↓

change

↓

change

↓

Emit Once
```

---

## Path Resolver

Converts

```
Relative

↓

Absolute

↓

Normalized Path
```

Supports

- Windows
- Linux
- macOS

---

## Cache Synchronizer

Updates

```
Workspace Cache

↓

Agent Context

↓

Search Index

↓

File Tree
```

---

## Snapshot Manager

Stores

```
Previous State

↓

Compare

↓

Detect Changes
```

Useful when native watching is unavailable.

---

# Watch Lifecycle

```
Application Starts

↓

Load Workspace

↓

Start Watcher

↓

Receive Events

↓

Filter

↓

Queue

↓

Dispatch

↓

Update Cache

↓

Notify Event Bus
```

---

# File Event Types

Supported events

```
Create

Modify

Delete

Rename

Move

Permission Change

Metadata Change
```

---

# Event Object

Contains

```
Event Type

Path

Timestamp

Workspace

Old Path

New Path

Metadata
```

---

# Event Flow

```
OS Notification

↓

Watcher

↓

Ignore Filter

↓

Debouncer

↓

Queue

↓

Dispatcher

↓

Event Bus
```

---

# OS Integration

Typical native APIs

```
Linux

inotify

-------------

macOS

FSEvents

-------------

Windows

ReadDirectoryChangesW
```

The Watcher abstracts platform differences.

---

# Recursive Watching

Example

```
workspace/

    src/

        app/

        ui/

    docs/

    config/

```

Every folder is monitored automatically.

---

# Ignore Rules

Support

```
Glob Patterns

Regular Expressions

Exact Paths

Extensions
```

Examples

```
*.log

*.tmp

node_modules/**

.git/**

dist/**
```

---

# Debounce Strategy

Without debounce

```
Save File

↓

20 Events

↓

20 Renders
```

With debounce

```
Save File

↓

20 Events

↓

1 Update
```

---

# Event Bus Integration

Common events

```
watch:start

watch:stop

file:create

file:update

file:delete

watch:error
```

---

# Agent Integration

```
File Changed

↓

Invalidate Context

↓

Reload File

↓

Continue Reasoning
```

---

# TUI Integration

```
File Created

↓

Explorer Refresh

↓

Render

↓

Done
```

---

# Plugin Integration

Plugins may

- Watch additional folders
- Filter events
- Register custom handlers
- Trigger workflows

---

# Skills Integration

Skills may react to

```
File Created

↓

Generate Template

-------------

Config Updated

↓

Reload Rules
```

---

# MCP Integration

Watcher may notify

```
Remote Sync

↓

Git Status

↓

Cloud Workspace

↓

External Tools
```

---

# Search Index Integration

```
File Updated

↓

Re-index

↓

Semantic Search Updated
```

---

# Session Integration

Remember

```
Open Files

Watched Paths

Ignored Paths
```

Restore automatically.

---

# Error Handling

```
Watcher Failure

↓

Retry

↓

Fallback Polling

↓

Log

↓

Continue
```

---

# Polling Fallback

If native watching fails

```
Timer

↓

Snapshot

↓

Compare

↓

Detect Changes
```

Slower but reliable.

---

# Performance Optimizations

Use

- Native OS watchers
- Debouncing
- Event batching
- Recursive watching
- Ignore filters
- Incremental cache updates
- Lazy indexing

Avoid

- Polling continuously
- Watching temporary folders
- Reloading the entire workspace

---

# Security

Always

- Validate paths
- Restrict workspace boundaries
- Prevent symbolic link loops
- Ignore restricted directories
- Sanitize file metadata

Never

- Watch arbitrary system folders
- Trust external paths
- Follow infinite symlink chains

---

# Best Practices

Always

- Watch recursively
- Debounce updates
- Ignore unnecessary folders
- Cache file metadata
- Emit structured events
- Handle rename correctly
- Support multiple workspaces

Never

- Reload everything
- Block the UI thread
- Ignore filesystem errors
- Assume event ordering

---

# Common Mistakes

Bad

```
File Changed

↓

Reload Entire Project
```

Very slow.

---

Good

```
File Changed

↓

Update Cache

↓

Refresh Only Affected Components
```

Efficient and scalable.

---

# Testing Checklist

- File creation
- File modification
- File deletion
- Rename
- Move
- Recursive watching
- Ignore rules
- Debounce
- Multiple workspaces
- Symbolic links
- Native watcher failure
- Polling fallback

---

# Advantages

- Real-time updates
- Better AI context
- Automatic synchronization
- Faster UI refresh
- Improved search indexing
- Lower manual effort
- Better developer experience

---

# Disadvantages

- Platform-specific APIs
- Native watcher limits
- Large workspaces require optimization
- Event storms without debouncing

---

# Used In

- OpenCode
- Antigravity CLI
- VS Code
- Cursor
- Zed Editor
- IntelliJ IDEA
- WebStorm
- Neovim
- Helix
- Sublime Text

---

# Complete Integration Flow

```
User Saves File

↓

Operating System

↓

Native File Watcher

↓

Ignore Filter

↓

Debouncer

↓

Event Queue

↓

Dispatcher

↓

Event Bus

↓

Agent Context Updated

↓

Search Index Updated

↓

Plugin Hooks Triggered

↓

TUI Refresh

↓

User Sees Latest State
```

---

# Summary

The **File Watcher** is the synchronization layer between the filesystem and the AI Coding Agent.

A production-grade File Watcher should include:

- Native OS integration
- Recursive directory monitoring
- Ignore Manager
- Event Queue
- Debouncer
- Path Resolver
- Cache Synchronizer
- Snapshot/Polling Fallback
- Event Bus integration
- Plugin and Skill support

By monitoring the workspace in real time and efficiently propagating file changes throughout the application, the File Watcher ensures that the Agent, UI, plugins, search index, and all other subsystems always operate on the latest project state with minimal latency and maximum performance.