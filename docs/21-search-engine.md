# 21-search-engine.md

# Search Engine Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Search Engine?

A **Search Engine** is the subsystem responsible for locating relevant information from the workspace, memory, documentation, conversation history, symbols, plugins, MCP servers, and external sources.

Instead of forcing the AI Agent to scan every available resource, the Search Engine retrieves only the most relevant information needed for the current task.

The Search Engine is one of the primary retrieval layers in an AI Coding Agent.

---

# Why Search Engine?

Without Search Engine

```
User Request

↓

Entire Workspace

↓

LLM
```

Problems

- Slow
- High token usage
- Poor scalability
- Duplicate processing
- Unnecessary file loading

---

With Search Engine

```
User Request

↓

Search Engine

↓

Relevant Results

↓

Context Builder

↓

LLM
```

---

# Goals

A production Search Engine should provide

- Fast retrieval
- Full-text search
- Symbol search
- Semantic search
- Hybrid search
- Ranking
- Filtering
- Incremental indexing
- Multi-source search
- Low latency

---

# High-Level Architecture

```
             User Query

                  │

                  ▼

            Search Engine

                  │

      ┌───────────┼────────────┐

      ▼           ▼            ▼

 Query      Search Index    Ranking

      ▼           ▼            ▼

 Filters   Semantic DB   Cache

      └───────────┼────────────┘

                  ▼

        Context Builder
```

---

# Folder Structure

```
src/

search/

    SearchEngine.ts

    QueryParser.ts

    SearchIndex.ts

    SearchRanker.ts

    SearchFilter.ts

    SemanticSearch.ts

    FullTextSearch.ts

    SymbolSearch.ts

    SearchCache.ts

    SearchMetrics.ts

    SearchEvents.ts

    SearchValidator.ts
```

---

# Core Components

## Search Engine

Central controller.

Responsibilities

- Receive query
- Search multiple sources
- Rank results
- Return relevant matches

---

## Query Parser

Converts

```
Natural Language

↓

Search Query
```

Example

```
Find login controller

↓

login controller
```

---

## Search Index

Stores

```
Files

Symbols

Text

Metadata

Documentation
```

Provides fast lookups.

---

## Search Ranker

Ranks by

- Relevance
- Similarity
- Recency
- Popularity
- Workspace proximity

---

## Search Filter

Supports

```
Language

Folder

Extension

Workspace

Tags

Metadata
```

---

## Semantic Search

Uses embeddings to find

```
Meaning

↓

Relevant Code
```

instead of exact text matches.

---

## Full-Text Search

Matches

```
Exact Words

↓

Relevant Files
```

---

## Symbol Search

Searches

```
Functions

Classes

Interfaces

Methods

Enums

Variables
```

---

## Search Cache

Caches

```
Recent Queries

Results

Rankings
```

Improves performance.

---

## Search Validator

Checks

- Query format
- Search scope
- Permissions

---

# Search Lifecycle

```
Receive Query

↓

Parse

↓

Search

↓

Filter

↓

Rank

↓

Return Results
```

---

# Search Sources

The engine may search

```
Workspace

Memory

Conversation

Documentation

Git History

Plugins

Skills

MCP

External APIs
```

---

# Query Object

Contains

```
Query

Filters

Workspace

Language

Search Type

Metadata
```

---

# Result Object

Contains

```
File

Symbol

Score

Snippet

Path

Metadata
```

---

# Search Types

Supported

```
Full-Text Search

Symbol Search

Semantic Search

Hybrid Search

Metadata Search
```

---

# Hybrid Search

Combines

```
Full-Text

+

Semantic

+

Symbol

↓

Ranked Results
```

Provides higher accuracy.

---

# Ranking Strategy

Factors

- Similarity
- Keyword match
- File importance
- Dependency distance
- User activity
- Recent edits

---

# Filtering

Example

```
Language

↓

TypeScript

Folder

↓

src/

Extension

↓

.ts
```

---

# Semantic Search Flow

```
Query

↓

Embedding

↓

Vector Search

↓

Similarity Score

↓

Results
```

---

# Full-Text Search Flow

```
Query

↓

Index

↓

Matching Files
```

---

# Symbol Search Flow

```
Search Function

↓

Symbol Index

↓

Function Location
```

---

# Incremental Updates

```
File Changed

↓

Workspace Indexer

↓

Search Index Updated
```

No full rebuild required.

---

# Event Bus Integration

Common events

```
search:start

search:query

search:result

search:error

search:complete
```

---

# Workspace Indexer Integration

```
Workspace Index

↓

Search Engine
```

The Search Engine never scans the workspace directly.

---

# Context Builder Integration

Provides

```
Relevant Files

Relevant Symbols

Relevant Documentation
```

for prompt generation.

---

# Memory Integration

Search

```
Long-Term Memory

↓

Relevant Knowledge
```

before querying the workspace.

---

# Conversation Integration

Search

```
Conversation History

↓

Relevant Messages
```

to avoid repeated explanations.

---

# Plugin Integration

Plugins may add

- Search providers
- Ranking algorithms
- External sources
- Filters

---

# MCP Integration

Search external resources through

```
MCP Servers

↓

Results
```

Examples

- GitHub
- Databases
- Documentation
- APIs

---

# Skills Integration

Skills may

- Expand queries
- Add filters
- Improve ranking
- Provide search templates

---

# Cache Strategy

Cache

```
Search Query

↓

Results

↓

Ranking
```

Invalidate when indexes change.

---

# Error Handling

```
No Results

↓

Expand Search

↓

Retry

↓

Return Empty
```

Never crash because of missing matches.

---

# Security

Always

- Respect workspace permissions
- Validate search scope
- Filter sensitive files
- Protect private indexes

Never

- Search unauthorized directories
- Leak confidential data
- Ignore permission rules

---

# Performance Optimizations

Use

- Incremental indexes
- Query cache
- Parallel search
- Lazy loading
- Result pagination
- Background ranking

Avoid

- Full workspace scans
- Duplicate searches
- Blocking the UI

---

# Best Practices

Always

- Separate indexing from searching
- Rank results
- Support semantic search
- Cache frequent queries
- Filter before ranking

Never

- Return every match
- Ignore workspace context
- Rebuild indexes during search

---

# Common Mistakes

Bad

```
Query

↓

Entire Repository

↓

Results
```

Slow and expensive.

---

Good

```
Query

↓

Search Index

↓

Rank

↓

Relevant Results
```

Fast and scalable.

---

# Testing Checklist

- Full-text search
- Symbol search
- Semantic search
- Hybrid search
- Ranking
- Filtering
- Cache
- Incremental updates
- MCP search
- Plugin integration
- Error handling

---

# Advantages

- Fast retrieval
- Better AI context
- Reduced token usage
- Workspace awareness
- Scalable architecture
- Improved developer productivity

---

# Disadvantages

- Index maintenance
- Embedding storage
- Ranking complexity
- Cache synchronization

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

# Complete Search Flow

```
User Query

↓

Query Parser

↓

Search Engine

↓

Workspace Index

↓

Memory Search

↓

Conversation Search

↓

Semantic Search

↓

Ranking

↓

Filtering

↓

Search Cache

↓

Context Builder

↓

Prompt Builder

↓

LLM
```

---

# Summary

The **Search Engine** is the intelligent retrieval layer responsible for locating, ranking, and delivering the most relevant information from the workspace and other knowledge sources.

A production-grade Search Engine should include:

- Search Engine
- Query Parser
- Search Index
- Search Ranker
- Search Filter
- Full-Text Search
- Symbol Search
- Semantic Search
- Search Cache
- Validator
- Event Bus Integration

By combining full-text, symbol, semantic, and hybrid search with efficient indexing and ranking, the Search Engine enables AI Coding Agents such as OpenCode, Antigravity CLI, Cursor, Claude Code, and enterprise AI platforms to retrieve highly relevant context quickly while minimizing latency and token consumption.