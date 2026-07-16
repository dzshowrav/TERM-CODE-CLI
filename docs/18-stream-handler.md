# 18-stream-handler.md

# Stream Handler Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Stream Handler?

A **Stream Handler** is the subsystem responsible for receiving, processing, buffering, validating, and distributing streaming data in real time.

In an AI Coding Agent, streaming typically means receiving partial responses (tokens or chunks) from an LLM provider and displaying them immediately instead of waiting for the complete response.

The Stream Handler separates streaming logic from the LLM Provider, Conversation Manager, and Renderer.

---

# Why Stream Handler?

Without Stream Handler

```
LLM

↓

Wait

↓

Entire Response

↓

User
```

Problems

- Slow feedback
- Poor user experience
- No live rendering
- Difficult interruption
- High perceived latency

---

With Stream Handler

```
LLM

↓

Stream Handler

↓

Token Buffer

↓

Conversation

↓

Renderer

↓

User
```

---

# Goals

A production Stream Handler should provide

- Real-time streaming
- Token buffering
- Chunk processing
- Partial rendering
- Stream interruption
- Stream validation
- Multi-provider compatibility
- Error recovery
- Event-driven architecture
- Backpressure handling

---

# High-Level Architecture

```
            LLM Provider

                 │

                 ▼

          Stream Handler

                 │

      ┌──────────┼──────────┐

      ▼          ▼          ▼

 Buffer     Validator   Dispatcher

      ▼          ▼          ▼

 Parser     Aggregator  Metrics

      └──────────┼──────────┘

                 ▼

 Conversation Manager

                 ▼

 Renderer

                 ▼

 User
```

---

# Folder Structure

```
src/

stream/

    StreamHandler.ts

    StreamBuffer.ts

    StreamParser.ts

    StreamAggregator.ts

    StreamDispatcher.ts

    StreamValidator.ts

    StreamController.ts

    StreamInterrupt.ts

    StreamMetrics.ts

    StreamEvents.ts

    StreamCache.ts

    StreamState.ts
```

---

# Core Components

## Stream Handler

Central controller.

Responsibilities

- Receive stream
- Process chunks
- Coordinate streaming
- Finalize output

---

## Stream Buffer

Temporarily stores

```
Incoming Tokens

↓

Buffered Chunks
```

Prevents rendering issues.

---

## Stream Parser

Parses

```
Provider Chunk

↓

Structured Token
```

Supports multiple provider formats.

---

## Stream Aggregator

Combines

```
Token

↓

Sentence

↓

Paragraph

↓

Complete Response
```

---

## Stream Dispatcher

Distributes updates to

- Conversation Manager
- Renderer
- Plugins
- Metrics

---

## Stream Validator

Checks

- Valid chunks
- Order
- Encoding
- Completion state

---

## Stream Controller

Controls

- Pause
- Resume
- Stop
- Restart

---

## Stream Interrupt

Supports

```
User Stops Response

↓

Cancel Stream

↓

Cleanup
```

---

# Stream Lifecycle

```
Open Stream

↓

Receive Chunks

↓

Parse

↓

Buffer

↓

Dispatch

↓

Render

↓

Complete

↓

Close
```

---

# Stream States

```
Idle

↓

Connecting

↓

Streaming

↓

Paused

↓

Completed

↓

Cancelled

↓

Error
```

---

# Chunk Object

Contains

```
Chunk ID

Token

Sequence

Timestamp

Provider Metadata

Completion Flag
```

---

# Complete Response

Built from

```
Chunk 1

↓

Chunk 2

↓

Chunk 3

↓

Final Response
```

---

# Streaming Flow

```
Provider

↓

Chunk

↓

Parser

↓

Buffer

↓

Aggregator

↓

Conversation

↓

Renderer
```

---

# Token Buffering

Purpose

```
Incoming Tokens

↓

Temporary Buffer

↓

Smooth Rendering
```

Reduces UI flickering.

---

# Incremental Rendering

Instead of

```
Entire Response
```

Use

```
Hello

↓

Hello World

↓

Hello World!
```

Users receive immediate feedback.

---

# Multi-Provider Support

Handle formats from

```
OpenAI

Anthropic

Gemini

OpenRouter

Ollama

LM Studio
```

Each adapter feeds normalized chunks to the Stream Handler.

---

# Stream Completion

Conditions

- Finish reason received
- Provider closes connection
- User cancels
- Timeout reached

---

# Event Bus Integration

Common events

```
stream:start

stream:chunk

stream:update

stream:complete

stream:cancel

stream:error
```

---

# Conversation Manager Integration

```
Chunk

↓

Conversation Buffer

↓

Final Message
```

The conversation updates continuously during streaming.

---

# Renderer Integration

```
New Chunk

↓

Partial Render

↓

Terminal Update
```

The Renderer only displays processed content.

---

# LLM Provider Integration

```
Provider Adapter

↓

Stream Handler
```

The provider manager never renders tokens directly.

---

# Session Integration

Store

```
Streaming State

Interrupted Responses

Partial Output
```

for crash recovery if needed.

---

# Plugin Integration

Plugins may

- Observe streams
- Modify chunks
- Add analytics
- Export streams

---

# Skills Integration

Skills may react to

```
Streaming Started

↓

Prepare Workspace

Streaming Completed

↓

Trigger Workflow
```

---

# Backpressure Handling

If rendering is slower than streaming

```
Incoming Chunks

↓

Buffer

↓

Controlled Dispatch

↓

Renderer
```

Prevents data loss.

---

# Error Handling

```
Invalid Chunk

↓

Discard

↓

Continue
```

or

```
Connection Lost

↓

Retry

↓

Resume

↓

Abort
```

---

# Timeout Strategy

```
No Data

↓

Timeout

↓

Cancel Stream

↓

Notify User
```

---

# Cache Strategy

Cache

```
Partial Responses

Recent Chunks

Provider Metadata
```

until completion.

---

# Security

Always

- Validate chunk encoding
- Verify provider identity
- Sanitize streamed content
- Limit buffer growth

Never

- Trust malformed chunks
- Expose provider internals
- Allow unlimited memory usage

---

# Performance Optimizations

Use

- Incremental parsing
- Buffered rendering
- Chunk batching
- Memory reuse
- Asynchronous processing
- Event-driven dispatch

Avoid

- Rendering every byte
- Blocking the UI thread
- Duplicating chunk processing

---

# Best Practices

Always

- Normalize provider streams
- Buffer before rendering
- Support interruption
- Validate chunks
- Emit lifecycle events
- Separate parsing from rendering

Never

- Couple rendering to providers
- Ignore ordering
- Drop valid chunks
- Leak buffers after completion

---

# Common Mistakes

Bad

```
Provider

↓

Renderer
```

Renderer becomes provider-dependent.

---

Good

```
Provider

↓

Stream Handler

↓

Conversation

↓

Renderer
```

Clean separation of responsibilities.

---

# Testing Checklist

- Stream start
- Chunk ordering
- Incremental rendering
- Buffering
- Completion
- Cancellation
- Retry
- Timeout
- Multi-provider compatibility
- Error recovery

---

# Advantages

- Faster perceived responses
- Better user experience
- Real-time rendering
- Provider abstraction
- Reliable interruption
- Modular architecture

---

# Disadvantages

- Buffer management complexity
- Provider-specific stream formats
- Backpressure handling
- Synchronization overhead

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

# Complete Streaming Flow

```
User Prompt

↓

Model Router

↓

LLM Provider Manager

↓

Provider Adapter

↓

Stream Handler

↓

Chunk Parser

↓

Token Buffer

↓

Aggregator

↓

Conversation Manager

↓

Renderer

↓

User

↓

Final Response Stored
```

---

# Summary

The **Stream Handler** is the real-time processing layer responsible for receiving, buffering, validating, aggregating, and distributing streamed AI responses.

A production-grade Stream Handler should include:

- Stream Handler
- Stream Buffer
- Stream Parser
- Stream Aggregator
- Stream Dispatcher
- Stream Validator
- Stream Controller
- Stream Interrupt
- Stream Cache
- Metrics
- Event Bus Integration

By separating streaming logic from provider communication, conversation management, and rendering, the Stream Handler enables responsive, reliable, and scalable real-time AI interactions across multiple LLM providers, making it an essential architectural component of modern AI Coding Agents such as OpenCode, Antigravity CLI, and enterprise AI platforms.