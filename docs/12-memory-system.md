# 12-memory-system.md

# Memory System Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is Memory?

A **Memory System** enables an AI agent to remember information beyond the current prompt.

Unlike a Session, which stores the current working state, Memory stores knowledge that helps the Agent make better decisions over time.

Memory may last for

- One conversation
- One project
- Multiple sessions
- Permanently

---

# Why Memory?

Without Memory

```
User

↓

"Use TypeScript"

↓

Agent forgets

↓

Repeats mistakes
```

---

With Memory

```
User

↓

Preference Stored

↓

Future Requests

↓

Automatically Applied
```

---

# Goals

A production Memory System should provide

- Long-term memory
- Short-term memory
- Project memory
- Semantic retrieval
- Memory ranking
- Automatic summarization
- Memory expiration
- Fast lookup
- Secure storage
- Context injection

---

# High-Level Architecture

```
              Agent Engine

                   │

                   ▼

            Memory Manager

                   │

      ┌────────────┼─────────────┐

      ▼            ▼             ▼

 Short Memory  Long Memory  Project Memory

      │            │             │

      └──────┬─────┴──────┬──────┘

             ▼            ▼

        Memory Index   Embeddings

             │

             ▼

      Context Builder
```

---

# Folder Structure

```
src/

memory/

    MemoryManager.ts

    MemoryStore.ts

    MemoryIndex.ts

    MemoryRetriever.ts

    MemoryRanker.ts

    MemoryCompressor.ts

    MemorySummarizer.ts

    MemoryValidator.ts

    MemoryEvents.ts

    MemoryMetrics.ts

    EmbeddingStore.ts

    ProjectMemory.ts

    ConversationMemory.ts
```

---

# Types of Memory

## Working Memory

Stores

- Current prompt
- Current tool results
- Active reasoning

Lifetime

```
One request
```

---

## Short-Term Memory

Stores

- Current conversation
- Recent actions
- Recent files

Lifetime

```
Current session
```

---

## Long-Term Memory

Stores

- User preferences
- Coding style
- Learned facts
- Frequently used workflows

Lifetime

```
Multiple sessions
```

---

## Project Memory

Stores

- Folder structure
- Framework
- Coding conventions
- Architecture
- APIs
- Important files

---

# Memory Lifecycle

```
Capture

↓

Validate

↓

Index

↓

Store

↓

Retrieve

↓

Update

↓

Expire
```

---

# Memory Object

Contains

```
ID

Type

Content

Source

Timestamp

Importance

Tags

Workspace

Embedding

Metadata
```

---

# Retrieval Flow

```
User Prompt

↓

Retriever

↓

Similarity Search

↓

Ranking

↓

Top Results

↓

Prompt Builder
```

---

# Memory Ranking

Factors

- Similarity
- Recency
- Importance
- Frequency
- Workspace match

---

# Compression

Old conversations become

```
100 Messages

↓

Summary

↓

Memory Entry
```

This saves context tokens.

---

# Event Bus Integration

Events

```
memory:create

memory:update

memory:retrieve

memory:delete

memory:expire
```

---

# Agent Integration

```
User Prompt

↓

Retrieve Memories

↓

Inject Context

↓

LLM
```

---

# Session Integration

```
Session Ends

↓

Important Facts

↓

Long-Term Memory
```

---

# Skills Integration

Skills may save

```
Templates

Rules

Examples
```

for future reuse.

---

# Plugin Integration

Plugins may create

- Custom memories
- Domain-specific indexes
- External storage providers

---

# Performance Optimizations

Use

- Vector index
- Embedding cache
- Memory compression
- Incremental indexing
- Lazy loading

Avoid

- Loading every memory
- Storing duplicate facts
- Keeping irrelevant history

---

# Security

Always

- Encrypt sensitive memories
- Respect workspace isolation
- Validate stored data
- Support memory deletion

Never

- Store secrets in plain text
- Mix memories between projects

---

# Best Practices

Always

- Rank memories
- Compress old conversations
- Separate project memory
- Keep metadata
- Remove obsolete memories

Never

- Inject every memory
- Ignore expiration
- Duplicate entries

---

# Testing Checklist

- Memory creation
- Retrieval
- Ranking
- Compression
- Expiration
- Multi-project isolation
- Search accuracy
- Recovery
- Performance

---

# Advantages

- Personalized AI
- Better reasoning
- Less repeated context
- Faster responses
- Lower token usage
- Consistent coding style

---

# Used In

- OpenCode
- Claude Code
- Cursor
- Continue.dev
- GitHub Copilot Agents
- Enterprise AI Assistants

---

# Complete Memory Flow

```
User Prompt

↓

Memory Retriever

↓

Similarity Search

↓

Ranking

↓

Context Builder

↓

Agent Engine

↓

LLM

↓

Response

↓

New Knowledge

↓

Memory Store
```

---

# Summary

The **Memory System** provides persistent intelligence for an AI Coding Agent by storing, retrieving, ranking, and summarizing information across requests, sessions, and projects.

A production-grade implementation should include:

- Memory Manager
- Short-Term Memory
- Long-Term Memory
- Project Memory
- Memory Index
- Retriever
- Ranker
- Compressor
- Summarizer
- Embedding Store
- Event Bus Integration

A robust Memory System enables the Agent to evolve from a stateless chatbot into a context-aware development assistant capable of maintaining consistency, reducing token usage, and improving productivity over time.