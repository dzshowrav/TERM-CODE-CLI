# 20-workspace-indexer.md

# Workspace Indexer Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Workspace Indexer?

A **Workspace Indexer** is the subsystem responsible for scanning, analyzing, indexing, and maintaining a searchable representation of an entire project workspace.

Instead of reading every file on every request, the AI Agent queries the Workspace Indexer to quickly locate relevant files, symbols, dependencies, documentation, and project metadata.

The Workspace Indexer acts as the knowledge base of the current workspace.

---

# Why Workspace Indexer?

Without Workspace Indexer

```
User Request

↓

Read Entire Project

↓

LLM
```

Problems

- Slow
- Expensive
- High token usage
- Duplicate file reads
- Poor scalability

---

With Workspace Indexer

```
Workspace

↓

Workspace Indexer

↓

Fast Search

↓

Relevant Files

↓

Context Builder
```

---

# Goals

A production Workspace Indexer should provide

- Fast project search
- File indexing
- Symbol indexing
- Dependency graph
- Language awareness
- Incremental updates
- Workspace metadata
- Semantic search
- Cache support
- Real-time synchronization

---

# High-Level Architecture

```
              Workspace

                   │

                   ▼

         Workspace Indexer

                   │

     ┌─────────────┼─────────────┐

     ▼             ▼             ▼

 File Index   Symbol Index   Metadata

     ▼             ▼             ▼

 Search      Dependency     Cache

     └─────────────┼─────────────┘

                   ▼

          Context Builder
```

---

# Folder Structure

```
src/

workspace/

    WorkspaceIndexer.ts

    WorkspaceScanner.ts

    FileIndexer.ts

    SymbolIndexer.ts

    DependencyIndexer.ts

    MetadataIndexer.ts

    SearchIndex.ts

    WorkspaceCache.ts

    WorkspaceWatcher.ts

    WorkspaceValidator.ts

    WorkspaceEvents.ts

    WorkspaceMetrics.ts
```

---

# Core Components

## Workspace Indexer

Central controller.

Responsibilities

- Scan workspace
- Build indexes
- Update indexes
- Provide search APIs

---

## Workspace Scanner

Scans

```
Workspace

↓

Folders

↓

Files
```

Discovers project structure.

---

## File Indexer

Indexes

- File names
- Paths
- Extensions
- Sizes
- Timestamps

---

## Symbol Indexer

Indexes

```
Classes

Functions

Variables

Interfaces

Enums

Methods
```

Supports language-aware navigation.

---

## Dependency Indexer

Tracks

```
Imports

Exports

Packages

Modules

Relationships
```

Builds dependency graphs.

---

## Metadata Indexer

Stores

```
Framework

Language

Configuration

Build Tools

Project Name

Package Information
```

---

## Search Index

Supports

- File search
- Symbol search
- Text search
- Semantic search

---

## Workspace Cache

Caches

```
Indexes

Metadata

Search Results
```

Improves lookup performance.

---

## Workspace Watcher

Monitors

```
File Created

File Modified

File Deleted

Folder Renamed
```

Triggers incremental reindexing.

---

## Workspace Validator

Checks

- Workspace integrity
- Index consistency
- File accessibility

---

# Indexing Lifecycle

```
Open Workspace

↓

Scan Files

↓

Extract Symbols

↓

Build Indexes

↓

Cache

↓

Ready
```

---

# Indexed Objects

Each indexed entry may contain

```
ID

Path

Name

Type

Language

Symbols

Dependencies

Metadata

Timestamp
```

---

# Supported Indexes

```
File Index

Folder Index

Symbol Index

Dependency Index

Package Index

Configuration Index

Documentation Index
```

---

# File Discovery

Scans

```
Source Files

Configuration Files

Documentation

Assets

Scripts

Templates
```

---

# Symbol Extraction

Extract

```
Functions

Classes

Methods

Interfaces

Enums

Constants
```

Language-specific parsers may be used.

---

# Dependency Graph

Example

```
main.ts

↓

App.ts

↓

UserService.ts

↓

Database.ts
```

Helps identify related files.

---

# Search Flow

```
Search Query

↓

Search Index

↓

Rank Results

↓

Relevant Files
```

---

# Semantic Search

Supports

```
Natural Language

↓

Relevant Code
```

Example

```
"login system"

↓

AuthenticationController.ts

AuthService.ts

LoginForm.tsx
```

---

# Incremental Indexing

Instead of

```
Entire Workspace
```

Only update

```
Changed Files
```

This improves performance.

---

# Event Bus Integration

Common events

```
workspace:scan

workspace:index

workspace:update

workspace:search

workspace:error
```

---

# File Watcher Integration

```
File Changed

↓

Workspace Watcher

↓

Incremental Index

↓

Cache Update
```

---

# Context Builder Integration

Provides

```
Relevant Files

Symbols

Dependencies

Metadata
```

instead of scanning the project every time.

---

# Search Engine Integration

The Search Engine queries

```
Workspace Index

↓

Results
```

for fast retrieval.

---

# Session Integration

Session stores

```
Current Workspace

Last Scan

Open Files

Search History
```

---

# Memory Integration

Project-specific knowledge may be linked to indexed files.

---

# Plugin Integration

Plugins may add

- Custom indexers
- Additional metadata
- Language parsers
- Search providers

---

# Skills Integration

Skills may consume

```
Workspace Metadata

Dependency Graph

Symbol Index
```

to generate better responses.

---

# Cache Strategy

Cache

```
Indexes

Search Results

Metadata

Dependency Graph
```

Invalidate only when necessary.

---

# Error Handling

```
Unreadable File

↓

Skip

↓

Log

↓

Continue
```

Indexing should not stop because of one failure.

---

# Security

Always

- Respect workspace boundaries
- Ignore hidden secrets unless authorized
- Validate file access
- Protect cached metadata

Never

- Index unauthorized directories
- Expose private files
- Traverse outside allowed workspaces

---

# Performance Optimizations

Use

- Incremental indexing
- Parallel scanning
- Lazy symbol extraction
- Persistent cache
- Background indexing

Avoid

- Full rescans
- Duplicate parsing
- Blocking the UI

---

# Best Practices

Always

- Separate scanning from indexing
- Keep indexes updated
- Support multiple languages
- Cache search results
- Monitor file changes

Never

- Read the whole project repeatedly
- Ignore dependency relationships
- Rebuild indexes unnecessarily

---

# Common Mistakes

Bad

```
User Request

↓

Read Entire Repository
```

Slow and inefficient.

---

Good

```
Workspace

↓

Workspace Index

↓

Relevant Files

↓

Context Builder
```

Fast and scalable.

---

# Testing Checklist

- Workspace scanning
- File indexing
- Symbol extraction
- Dependency graph
- Metadata indexing
- Search accuracy
- Incremental updates
- Cache behavior
- File watcher integration
- Error recovery

---

# Advantages

- Faster project understanding
- Reduced token usage
- Efficient search
- Better context quality
- Scalable architecture
- Language-aware navigation

---

# Disadvantages

- Initial indexing cost
- Cache maintenance
- Language parser complexity
- Index synchronization

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Cursor
- Continue.dev
- GitHub Copilot Agents
- Enterprise AI Platforms

---

# Complete Workspace Indexing Flow

```
Workspace

↓

Workspace Scanner

↓

File Indexer

↓

Symbol Indexer

↓

Dependency Indexer

↓

Metadata Indexer

↓

Search Index

↓

Workspace Cache

↓

Context Builder

↓

Prompt Builder

↓

LLM
```

---

# Summary

The **Workspace Indexer** is the project intelligence layer responsible for scanning, indexing, organizing, and searching an entire development workspace.

A production-grade Workspace Indexer should include:

- Workspace Indexer
- Workspace Scanner
- File Indexer
- Symbol Indexer
- Dependency Indexer
- Metadata Indexer
- Search Index
- Workspace Cache
- Workspace Watcher
- Validator
- Event Bus Integration

By maintaining searchable indexes of files, symbols, dependencies, and project metadata, the Workspace Indexer enables AI Coding Agents such as OpenCode, Antigravity CLI, Cursor, and Claude Code to understand large codebases efficiently while minimizing latency, token usage, and unnecessary file processing.