# 13-context-builder.md

# Context Builder Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Context Builder?

A **Context Builder** is responsible for collecting, filtering, organizing, and preparing all relevant information before sending a request to the AI model.

The AI model should **never** receive only the user's message.

Instead, it receives a carefully constructed context containing everything needed to produce the best possible response.

---

# Why Context Builder?

Without Context Builder

```
User Prompt

↓

LLM
```

Problems

- No project awareness
- No memory
- No workspace knowledge
- Poor answers
- Repeated questions

---

With Context Builder

```
User Prompt

↓

Context Builder

↓

LLM
```

The model receives everything it needs.

---

# Goals

A production Context Builder should provide

- Intelligent context collection
- Token optimization
- Workspace awareness
- Conversation awareness
- Memory integration
- Tool awareness
- Skill awareness
- Plugin awareness
- Fast context assembly
- Automatic context compression

---

# High-Level Architecture

```
                User Prompt

                     │

                     ▼

             Context Builder

                     │

     ┌───────────────┼────────────────┐

     ▼               ▼                ▼

 Conversation     Workspace       Memory

     ▼               ▼                ▼

 Skills          Plugins         MCP

     ▼               ▼                ▼

 Search          File Cache      Session

     └───────────────┼────────────────┘

                     ▼

            Context Optimizer

                     ▼

             Prompt Builder
```

---

# Folder Structure

```
src/

context/

    ContextBuilder.ts

    ContextCollector.ts

    ContextResolver.ts

    ContextOptimizer.ts

    ContextFilter.ts

    ContextRanker.ts

    ContextCompressor.ts

    ContextValidator.ts

    ContextCache.ts

    ContextMetrics.ts

    ContextEvents.ts

    ContextSources.ts
```

---

# Core Components

## Context Builder

Central controller.

Responsibilities

- Build context
- Merge sources
- Optimize tokens
- Return final context

---

## Context Collector

Collects information from

- Conversation
- Files
- Memory
- Plugins
- Skills
- MCP
- Workspace

---

## Context Resolver

Determines

```
Relevant

↓

Irrelevant
```

Only useful information continues.

---

## Context Filter

Removes

- Duplicate content
- Empty entries
- Outdated data
- Unused files

---

## Context Ranker

Ranks context using

- Similarity
- Importance
- Recency
- Workspace relevance

---

## Context Compressor

Compresses

```
Large History

↓

Summary

↓

Smaller Context
```

---

## Context Cache

Stores recently built contexts.

Benefits

- Faster requests
- Reduced recomputation

---

## Context Validator

Checks

- Size
- Completeness
- Required sections
- Token limits

---

# Context Sources

The builder may collect information from

```
Conversation

Workspace

Memory

Session

Open Files

Git Status

Configuration

Environment

Skills

Plugins

MCP Tools

Search Results

Recent Commands
```

---

# Context Lifecycle

```
Collect

↓

Filter

↓

Rank

↓

Compress

↓

Merge

↓

Validate

↓

Return
```

---

# Context Object

Contains

```
System Prompt

Conversation

Workspace

Memory

Files

Tools

Skills

Plugins

Metadata
```

---

# Context Assembly Flow

```
User Prompt

↓

Collect Sources

↓

Resolve Relevance

↓

Remove Duplicates

↓

Rank

↓

Compress

↓

Merge

↓

Return
```

---

# Workspace Context

May include

```
Project Name

Framework

Folder Structure

Important Files

Dependencies

Configuration
```

---

# Conversation Context

Includes

```
Recent Messages

Tool Calls

AI Responses

Errors

Decisions
```

---

# Memory Context

Includes

```
User Preferences

Coding Style

Previous Solutions

Persistent Facts
```

---

# Session Context

Includes

```
Open Files

Current Directory

Selected Model

Loaded Skills

Loaded Plugins
```

---

# File Context

Collect

```
Current File

Referenced Files

Related Files

Imports

Dependencies
```

Avoid loading the entire project unnecessarily.

---

# Search Context

May include

```
Workspace Search

Symbol Search

Recent Files

Relevant Snippets
```

---

# Skill Context

Provides

```
Templates

Rules

Examples

Coding Standards
```

---

# Plugin Context

Plugins may contribute

```
Additional Instructions

Metadata

Project Information
```

---

# MCP Context

MCP servers may provide

```
Git Status

Database Schema

Browser State

External Documents
```

---

# Token Budget

Example

```
Available

100000 Tokens

↓

Conversation

20%

Workspace

25%

Memory

15%

Files

25%

Skills

10%

Reserved

5%
```

The Context Builder allocates the available budget intelligently.

---

# Ranking Strategy

Prioritize

1. Current file
2. User request
3. Recent conversation
4. Relevant memory
5. Related files
6. Skills
7. Plugins

---

# Compression Strategy

Example

```
200 Messages

↓

Conversation Summary

↓

Latest Messages
```

Old information becomes summarized.

---

# Event Bus Integration

Common events

```
context:start

context:collect

context:compress

context:complete

context:error
```

---

# Agent Integration

```
User Prompt

↓

Context Builder

↓

Prompt Builder

↓

LLM
```

---

# Search Engine Integration

```
Missing Context

↓

Search

↓

Relevant Results

↓

Context
```

---

# File Watcher Integration

```
File Changed

↓

Invalidate Context Cache

↓

Rebuild
```

---

# Session Integration

Restore

```
Workspace

Conversation

Open Files
```

before building context.

---

# Performance Optimizations

Use

- Incremental context updates
- Context cache
- Similarity ranking
- Lazy loading
- Compression
- Parallel collection

Avoid

- Reading every file
- Loading unnecessary history
- Rebuilding unchanged context

---

# Security

Always

- Respect workspace boundaries
- Filter sensitive files
- Remove secrets
- Validate context sources

Never

- Include private keys
- Leak unrelated projects
- Exceed model token limits

---

# Best Practices

Always

- Keep context relevant
- Compress older history
- Rank by importance
- Cache results
- Validate token usage

Never

- Send entire repositories
- Include duplicate information
- Ignore workspace relevance

---

# Common Mistakes

Bad

```
Entire Repository

↓

LLM
```

Huge cost and poor performance.

---

Good

```
Relevant Files

+

Recent Conversation

+

Memory

↓

LLM
```

Efficient and accurate.

---

# Testing Checklist

- Context collection
- Ranking
- Compression
- Token limits
- Cache
- File relevance
- Memory integration
- Skill integration
- Plugin integration
- Error handling

---

# Advantages

- Better AI responses
- Lower token usage
- Faster execution
- Project awareness
- Consistent reasoning
- Modular architecture

---

# Disadvantages

- Ranking complexity
- Compression overhead
- Token budgeting challenges
- Cache invalidation

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Cursor
- Continue.dev
- GitHub Copilot Agents
- Enterprise AI platforms

---

# Complete Context Flow

```
User Prompt

↓

Conversation

↓

Workspace

↓

Memory

↓

Files

↓

Skills

↓

Plugins

↓

MCP

↓

Ranking

↓

Compression

↓

Context Validation

↓

Prompt Builder

↓

LLM
```

---

# Summary

The **Context Builder** is the intelligence layer responsible for preparing the optimal input for an AI model.

A production-grade Context Builder should include:

- Context Collector
- Resolver
- Filter
- Ranker
- Compressor
- Cache
- Validator
- Token Budget Manager
- Event Bus Integration
- Workspace, Memory, Session, Skill, Plugin, and MCP integration

By collecting only the most relevant information and optimizing it within the available token budget, the Context Builder enables an AI Coding Agent to produce accurate, context-aware, and efficient responses while scaling to large codebases like OpenCode and Antigravity CLI.