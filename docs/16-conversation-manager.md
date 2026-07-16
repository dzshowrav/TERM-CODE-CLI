# 16-conversation-manager.md

# Conversation Manager Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Conversation Manager?

A **Conversation Manager** is the subsystem responsible for managing the complete lifecycle of interactions between the user and the AI Agent.

It organizes messages, maintains conversation history, tracks context across turns, manages branching conversations, coordinates tool calls, and ensures every request is processed in the correct conversational state.

Unlike the **Session Manager**, which manages the entire application state, the Conversation Manager focuses only on the dialogue and reasoning flow.

---

# Why Conversation Manager?

Without Conversation Manager

```
User

↓

Message

↓

LLM

↓

Response
```

Problems

- No conversation history
- No context continuity
- No branching
- No message organization
- Poor multi-turn reasoning

---

With Conversation Manager

```
User

↓

Conversation Manager

↓

Conversation State

↓

Agent

↓

LLM

↓

Assistant
```

---

# Goals

A production Conversation Manager should provide

- Multi-turn conversations
- Conversation history
- Branching conversations
- Message organization
- Tool call tracking
- Streaming message support
- Conversation summaries
- Context synchronization
- Token management
- Conversation recovery

---

# High-Level Architecture

```
                 User

                  │

                  ▼

        Conversation Manager

                  │

      ┌───────────┼─────────────┐

      ▼           ▼             ▼

 Message      History       Branches

      ▼           ▼             ▼

 State      Summaries      Tokens

      └───────────┼─────────────┘

                  ▼

           Context Builder

                  ▼

            Prompt Builder
```

---

# Folder Structure

```
src/

conversation/

    ConversationManager.ts

    Conversation.ts

    ConversationStore.ts

    ConversationLoader.ts

    ConversationSaver.ts

    Message.ts

    MessageQueue.ts

    MessageValidator.ts

    HistoryManager.ts

    SummaryManager.ts

    BranchManager.ts

    TokenTracker.ts

    ConversationEvents.ts

    ConversationMetrics.ts
```

---

# Core Components

## Conversation Manager

Central controller.

Responsibilities

- Create conversations
- Receive messages
- Save history
- Manage branches
- Restore conversations

---

## Conversation Store

Persists

```
Messages

Branches

Summaries

Metadata
```

---

## History Manager

Maintains

- Ordered messages
- Conversation timeline
- Message lookup

---

## Message Queue

Buffers

```
Incoming Messages

↓

Processing

↓

Completed
```

Useful for streaming and concurrency.

---

## Branch Manager

Supports

```
Conversation

↓

Branch A

↓

Branch B

↓

Branch C
```

Users can explore multiple reasoning paths.

---

## Summary Manager

Compresses

```
Long Conversation

↓

Summary

↓

Reduced Tokens
```

---

## Token Tracker

Tracks

- Prompt tokens
- Response tokens
- Total conversation size

---

## Message Validator

Checks

- Message format
- Roles
- Metadata
- Attachments

---

# Conversation Lifecycle

```
Create

↓

Receive Message

↓

Validate

↓

Append History

↓

Update Summary

↓

Save

↓

Continue
```

---

# Message Object

Contains

```
Message ID

Role

Content

Timestamp

Attachments

Tool Calls

Metadata
```

---

# Message Roles

Supported roles

```
System

Developer

User

Assistant

Tool
```

---

# Conversation States

```
New

↓

Active

↓

Waiting

↓

Streaming

↓

Completed

↓

Archived
```

---

# Conversation Timeline

```
System

↓

User

↓

Assistant

↓

Tool

↓

Assistant

↓

User
```

Every event is recorded in chronological order.

---

# Branching Conversations

Example

```
Conversation

      │

 ┌────┴────┐

 ▼         ▼

Branch A  Branch B

           │

           ▼

       Branch C
```

Each branch has independent history.

---

# Streaming Support

```
LLM

↓

Token Stream

↓

Message Buffer

↓

Final Message
```

The Conversation Manager assembles streamed tokens into complete responses.

---

# Conversation Summary

Summarize

```
Old Messages

↓

Conversation Summary

↓

Recent Messages
```

Keeps prompts within token limits.

---

# Tool Call Tracking

Each tool interaction records

```
Tool Name

Arguments

Execution Time

Result

Status
```

---

# Event Bus Integration

Common events

```
conversation:create

conversation:update

message:add

message:stream

conversation:summary

conversation:error
```

---

# Agent Integration

```
User Message

↓

Conversation Manager

↓

Context Builder

↓

Agent Engine
```

---

# Context Builder Integration

Provides

```
Recent Messages

Conversation Summary

Tool History
```

for context assembly.

---

# Session Integration

Stores

```
Active Conversation

Current Branch

Last Message
```

for restoration.

---

# Memory Integration

Important conversation facts may be promoted to

```
Long-Term Memory
```

---

# Plugin Integration

Plugins may

- Add message types
- Extend metadata
- Analyze conversations
- Generate summaries

---

# Skills Integration

Skills may inject

- Conversation templates
- Domain-specific prompts
- Conversation rules

---

# Search Integration

Search

```
Conversation History

↓

Relevant Messages

↓

Context
```

---

# Error Handling

```
Invalid Message

↓

Validation

↓

Reject

↓

Notify User
```

---

# Token Management

Monitor

```
Conversation Size

↓

Threshold

↓

Summarize Older Messages

↓

Continue
```

---

# Cache Strategy

Cache

```
Recent Messages

Conversation Summary

Branch Metadata
```

to improve performance.

---

# Security

Always

- Validate messages
- Sanitize attachments
- Protect conversation privacy
- Encrypt stored history if required

Never

- Mix conversations between sessions
- Store sensitive information without protection
- Ignore corrupted history

---

# Performance Optimizations

Use

- Incremental saving
- Conversation summaries
- Lazy loading
- Message caching
- Branch indexing
- Streaming buffers

Avoid

- Reloading entire history
- Re-summarizing unchanged conversations
- Blocking message processing

---

# Best Practices

Always

- Separate conversation from session
- Track message metadata
- Support branching
- Summarize long histories
- Validate messages
- Preserve chronological order

Never

- Lose conversation state
- Duplicate messages
- Ignore token growth
- Hardcode message roles

---

# Common Mistakes

Bad

```
Messages

↓

LLM

↓

Forget History
```

No continuity.

---

Good

```
Messages

↓

Conversation Manager

↓

History

↓

Summary

↓

LLM
```

Structured and scalable.

---

# Testing Checklist

- Conversation creation
- Message ordering
- Branching
- Streaming
- Tool tracking
- Summarization
- Recovery
- Cache
- Token tracking
- Error handling

---

# Advantages

- Multi-turn continuity
- Better AI reasoning
- Branch support
- Reduced token usage
- Organized history
- Improved user experience

---

# Disadvantages

- Storage overhead
- Summary complexity
- Branch management
- Token tracking maintenance

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

# Complete Conversation Flow

```
User Message

↓

Conversation Manager

↓

Validate Message

↓

Append History

↓

Track Tool Calls

↓

Update Summary

↓

Context Builder

↓

Prompt Builder

↓

Model Router

↓

LLM

↓

Assistant Response

↓

Conversation Store

↓

Ready For Next Turn
```

---

# Summary

The **Conversation Manager** is responsible for managing the complete lifecycle of AI conversations, ensuring continuity, organization, and efficient context management across multiple interactions.

A production-grade Conversation Manager should include:

- Conversation Manager
- Conversation Store
- History Manager
- Message Queue
- Branch Manager
- Summary Manager
- Token Tracker
- Message Validator
- Cache
- Event Bus Integration

By separating conversation management from session management and integrating with the Context Builder, Prompt Builder, Agent Engine, Memory System, and Session Manager, the Conversation Manager enables scalable, context-aware, multi-turn AI interactions similar to those found in OpenCode, Antigravity CLI, and other modern AI development platforms.