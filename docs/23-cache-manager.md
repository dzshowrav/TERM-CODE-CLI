# 23-cache-manager.md

# Cache Manager Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Cache Manager?

A **Cache Manager** is the subsystem responsible for temporarily storing frequently accessed data so it can be retrieved much faster than reading it again from the original source.

In an AI Coding Agent, the Cache Manager reduces latency, minimizes token usage, avoids repeated computations, and improves overall system performance.

Instead of recalculating or reloading the same information, the system retrieves it directly from the cache.

---

# Why Cache Manager?

Without Cache Manager

```
User Request

↓

Read Workspace

↓

Search

↓

LLM

↓

Response
```

Every request repeats expensive operations.

Problems

- Slow responses
- High token cost
- Duplicate computation
- Excessive disk access
- Poor scalability

---

With Cache Manager

```
User Request

↓

Cache Manager

↓

Cache Hit?

↓

Yes

↓

Return Cached Data

No

↓

Compute

↓

Store Cache

↓

Return
```

---

# Goals

A production Cache Manager should provide

- Fast data retrieval
- Automatic cache invalidation
- Multi-layer caching
- Memory optimization
- Token reduction
- Persistent cache
- Distributed cache support
- Cache metrics
- Expiration policies
- Thread-safe access

---

# High-Level Architecture

```
              Application

                    │

                    ▼

             Cache Manager

                    │

      ┌─────────────┼─────────────┐

      ▼             ▼             ▼

 Memory Cache   Disk Cache   Remote Cache

      ▼             ▼             ▼

 Policies      Expiration     Metrics

      └─────────────┼─────────────┘

                    ▼

              Data Source
```

---

# Folder Structure

```
src/

cache/

    CacheManager.ts

    CacheStore.ts

    MemoryCache.ts

    DiskCache.ts

    CacheKey.ts

    CachePolicy.ts

    CacheInvalidator.ts

    CacheSerializer.ts

    CacheMetrics.ts

    CacheEvents.ts

    CacheCleaner.ts

    CacheValidator.ts
```

---

# Core Components

## Cache Manager

Central controller.

Responsibilities

- Read cache
- Write cache
- Delete cache
- Manage expiration

---

## Cache Store

Stores

```
Key

Value

Metadata

Expiration
```

---

## Memory Cache

Fastest cache.

Suitable for

- Active session data
- Recent searches
- Recent prompts
- Recent responses

---

## Disk Cache

Persistent storage.

Suitable for

- Workspace indexes
- Embeddings
- Documentation
- Metadata

---

## Cache Policy

Defines

- Expiration
- Size limits
- Cleanup rules
- Priority

---

## Cache Invalidator

Removes outdated cache when

- Files change
- Configuration changes
- Session ends
- User clears cache

---

## Cache Serializer

Converts

```
Object

↓

Serialized Data

↓

Cache
```

and restores objects during retrieval.

---

## Cache Cleaner

Removes

- Expired entries
- Unused entries
- Oversized cache

---

## Cache Validator

Checks

- Integrity
- Expiration
- Data consistency

---

# Cache Lifecycle

```
Request

↓

Check Cache

↓

Hit?

↓

Return Cached Data

Miss?

↓

Load Source

↓

Store Cache

↓

Return
```

---

# Cache Object

Contains

```
Key

Value

Created Time

Expiration Time

Size

Metadata

Version
```

---

# Cache Levels

```
L1

Memory Cache

↓

L2

Disk Cache

↓

L3

Remote Cache

↓

Original Source
```

---

# Cache Types

Examples

```
Workspace Cache

Search Cache

Prompt Cache

Conversation Cache

Provider Cache

Tool Cache

Permission Cache

Embedding Cache
```

---

# Cache Keys

Examples

```
workspace:index

search:login

provider:gpt-5.5

conversation:123

tool:filesystem
```

Keys should be deterministic and unique.

---

# Cache Expiration

Strategies

```
Time To Live (TTL)

Sliding Expiration

Manual Invalidation

Version-Based Expiration
```

---

# Cache Policies

Examples

```
Keep 30 Minutes

Keep Until File Changes

Keep Current Session

Never Cache
```

---

# Cache Hit Flow

```
Request

↓

Cache Found

↓

Return Immediately
```

Very low latency.

---

# Cache Miss Flow

```
Request

↓

Cache Missing

↓

Read Source

↓

Save Cache

↓

Return
```

---

# Cache Invalidation

Triggers

```
Workspace Modified

↓

Invalidate Index

↓

Rebuild Cache
```

Ensures consistency.

---

# Event Bus Integration

Common events

```
cache:hit

cache:miss

cache:write

cache:invalidate

cache:clear

cache:error
```

---

# Workspace Indexer Integration

Caches

```
Workspace Metadata

Symbol Index

Dependency Graph
```

---

# Search Engine Integration

Caches

```
Search Queries

Search Results

Rankings
```

---

# Conversation Manager Integration

Caches

```
Conversation Summary

Recent Messages

Branch State
```

---

# Prompt Builder Integration

Caches

```
Generated Prompts

Prompt Templates
```

---

# Model Router Integration

Caches

```
Routing Decisions

Provider Health

Capability Maps
```

---

# LLM Provider Integration

Caches

```
Model Metadata

Provider Status

Supported Models
```

---

# Session Integration

Stores

```
Session Cache

Temporary Data

User Preferences
```

---

# Plugin Integration

Plugins may

- Register cache providers
- Define cache policies
- Store plugin metadata
- Clear plugin cache

---

# Skills Integration

Skills may cache

- Templates
- Knowledge
- Analysis results
- Search expansions

---

# Error Handling

```
Cache Corrupted

↓

Delete Entry

↓

Reload Source

↓

Continue
```

Never return invalid cache.

---

# Security

Always

- Encrypt sensitive cache
- Validate cached objects
- Respect workspace permissions
- Clear confidential data when required

Never

- Cache API keys
- Cache passwords
- Cache unauthorized data
- Return expired sensitive data

---

# Performance Optimizations

Use

- Memory cache
- Lazy loading
- Incremental invalidation
- Background cleanup
- Compression
- Object pooling

Avoid

- Unlimited cache growth
- Duplicate cache entries
- Full cache rebuilds
- Frequent disk writes

---

# Best Practices

Always

- Cache expensive operations
- Use meaningful keys
- Invalidate stale data
- Track cache metrics
- Separate memory and disk cache

Never

- Cache everything
- Ignore expiration
- Trust corrupted cache
- Mix unrelated cache entries

---

# Common Mistakes

Bad

```
Request

↓

Always Compute

↓

Return
```

No caching.

---

Good

```
Request

↓

Cache

↓

Hit

↓

Return

Miss

↓

Compute

↓

Store

↓

Return
```

Fast and efficient.

---

# Testing Checklist

- Cache hit
- Cache miss
- Expiration
- Invalidation
- Memory cache
- Disk cache
- Serialization
- Cleanup
- Metrics
- Error recovery

---

# Advantages

- Faster responses
- Lower token usage
- Reduced CPU usage
- Reduced disk access
- Better scalability
- Improved user experience

---

# Disadvantages

- Memory consumption
- Cache invalidation complexity
- Synchronization overhead
- Storage management

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

# Complete Cache Flow

```
Application Request

↓

Cache Manager

↓

Memory Cache

↓

Disk Cache

↓

Original Source

↓

Store Cache

↓

Return Result
```

---

# Summary

The **Cache Manager** is the performance optimization layer responsible for storing and serving frequently accessed data with minimal latency.

A production-grade Cache Manager should include:

- Cache Manager
- Cache Store
- Memory Cache
- Disk Cache
- Cache Policy
- Cache Invalidator
- Cache Serializer
- Cache Cleaner
- Cache Validator
- Cache Metrics
- Event Bus Integration

By implementing multi-layer caching, intelligent invalidation, and efficient storage strategies, the Cache Manager enables AI Coding Agents such as OpenCode, Antigravity CLI, Cursor, Claude Code, and enterprise AI platforms to deliver faster responses, reduce token consumption, minimize redundant computation, and scale efficiently across large projects.