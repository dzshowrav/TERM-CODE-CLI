# 01-event-bus.md

# Event Bus Architecture
### Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is an Event Bus?

An **Event Bus** is the communication backbone of an application.

Instead of every module talking directly to every other module, each module communicates through a central hub.

Think of it like this:

```
Agent
   │
   ▼
Event Bus
 ▲  ▲  ▲
 │  │  │
UI FileWatcher MCP
```

Nobody knows who receives the event.

Everyone simply emits events.

---

# Why Event Bus?

Without Event Bus:

```
Agent
 ├── UI
 ├── MCP
 ├── Theme
 ├── Plugin
 ├── Logger
 ├── History
 ├── Session
 ├── File Watcher
 └── Skills
```

Everything becomes tightly coupled.

Every module depends on every other module.

Maintenance becomes difficult.

---

With Event Bus

```
Agent
   │
Event Bus
 │ │ │ │ │ │ │
 ▼ ▼ ▼ ▼ ▼ ▼ ▼
UI
Logger
Theme
MCP
Plugin
Watcher
History
```

Every component stays independent.

---

# Goals

An Event Bus should provide:

- Loose coupling
- High scalability
- High performance
- Async communication
- Easy extension
- Easy testing
- Plugin compatibility
- Skill compatibility
- Multiple listeners

---

# Real World Usage

OpenCode

```
User types message

↓

Command Parser

↓

Agent

↓

EventBus.emit("agent:start")

↓

UI updates

↓

Logger saves

↓

Spinner starts

↓

Plugin notified

↓

History updated
```

Nobody directly calls each module.

Everything is event driven.

---

# Core Principles

## 1. Publisher

Produces event.

Example

```
Agent
```

publishes

```
agent:start
```

---

## 2. Subscriber

Listens to event.

Example

```
UI

Logger

Plugin
```

---

## 3. Event

Contains information.

Example

```
{
    id,
    timestamp,
    type,
    payload
}
```

---

## 4. Bus

Routes events.

```
Publisher

↓

Bus

↓

Subscribers
```

---

# Architecture

```
            User

             │

             ▼

      Command Parser

             │

             ▼

        Agent Engine

             │

             ▼

         Event Bus

 ┌────────┼─────────┐

 ▼        ▼         ▼

UI     Logger     Plugin

 ▼        ▼         ▼

Theme  History   Skills

             ▼

            MCP

             ▼

        File Watcher
```

---

# Folder Structure

```
src/

 events/

     EventBus.ts

     Event.ts

     EventTypes.ts

     Subscriber.ts

     Publisher.ts

     EventQueue.ts

     Middleware.ts

     Dispatcher.ts

     EventStore.ts

     EventContext.ts

     EventPriority.ts

     EventHistory.ts

     EventFilter.ts

     EventScheduler.ts

     EventMetrics.ts
```

---

# Core Components

## EventBus

Central dispatcher.

Responsibilities

- Register listeners
- Remove listeners
- Emit events
- Queue events
- Dispatch events
- Middleware
- Logging

---

## Event

Represents one action.

Contains

```
id

name

payload

timestamp

source

priority

metadata
```

---

## Subscriber

Listens.

Example

```
Theme

UI

Plugin

Logger
```

---

## Publisher

Produces event.

Example

```
Agent

Command Parser

MCP

Watcher
```

---

# Event Lifecycle

```
Create

↓

Validate

↓

Middleware

↓

Queue

↓

Dispatch

↓

Subscribers

↓

Complete

↓

Store History
```

---

# Event Flow

```
User

↓

Agent

↓

EventBus.emit()

↓

Queue

↓

Dispatcher

↓

Subscriber

↓

UI Update
```

---

# Types of Events

## UI Events

```
ui:render

ui:update

ui:refresh

ui:clear

ui:error
```

---

## Agent Events

```
agent:start

agent:stop

agent:thinking

agent:response

agent:error
```

---

## MCP Events

```
mcp:connect

mcp:disconnect

mcp:tool

mcp:error
```

---

## Theme Events

```
theme:change

theme:reload
```

---

## Plugin Events

```
plugin:load

plugin:enable

plugin:disable

plugin:error
```

---

## Watcher Events

```
watch:file

watch:create

watch:update

watch:delete
```

---

## Session Events

```
session:start

session:end

session:save

session:restore
```

---

# Event Payload

Example

```
{
    type:"agent:start",

    timestamp:123456,

    source:"Agent",

    payload:{
        prompt:"Explain React"
    }
}
```

---

# Priority Levels

```
CRITICAL

HIGH

NORMAL

LOW

BACKGROUND
```

Example

```
User Input

↓

HIGH

Theme Refresh

↓

LOW

Analytics

↓

BACKGROUND
```

---

# Queue

Incoming events

↓

FIFO Queue

↓

Dispatcher

↓

Subscribers

Queue prevents blocking.

---

# Dispatcher

Responsibilities

- Read queue
- Route events
- Execute listeners
- Retry failures
- Remove completed

---

# Middleware

Middleware executes before subscribers.

Example

```
Validation

↓

Authentication

↓

Logging

↓

Metrics

↓

Dispatch
```

---

# Event History

Stores

```
Timestamp

Type

Payload

Duration

Status
```

Useful for debugging.

---

# Metrics

Track

```
Events/sec

Queue length

Subscriber time

Dispatch latency

Dropped events

Retries
```

---

# Error Handling

```
Emit

↓

Subscriber Error

↓

Catch

↓

Retry

↓

Fallback

↓

Logger
```

One broken listener should never stop the bus.

---

# Best Practices

Always

- Use immutable payloads
- Keep events small
- Use meaningful names
- Separate event types
- Log failures
- Add timestamps
- Use priorities
- Keep subscribers independent

Never

- Directly call another module
- Share mutable objects
- Create circular events
- Block event loop
- Perform heavy work in subscribers

---

# Common Mistakes

Bad

```
Agent

↓

UI.update()

↓

Logger.save()

↓

Plugin.notify()

↓

Theme.refresh()
```

Everything is coupled.

Good

```
Agent

↓

emit()

↓

EventBus

↓

Everyone listens
```

---

# Performance Optimizations

Use

- Event queue
- Async dispatch
- Event batching
- Priority scheduling
- Debounce
- Throttle
- Worker threads (if needed)

---

# Event Naming Convention

Use

```
module:action
```

Examples

```
agent:start

agent:complete

plugin:load

plugin:error

theme:update

ui:refresh
```

Avoid

```
start

event1

abc

run
```

---

# Debugging Strategy

Log

```
Time

Event

Source

Target

Duration

Error
```

Example

```
12:11:04

agent:start

↓

UI

4ms

SUCCESS
```

---

# Testing Checklist

- Single listener
- Multiple listeners
- No listener
- Priority events
- Queue overflow
- Retry logic
- Error handling
- Middleware execution
- History storage
- Concurrent events

---

# Advantages

- Decoupled architecture
- Easy maintenance
- Easier testing
- Plugin friendly
- Skill friendly
- Highly scalable
- Better debugging
- Better performance
- Cleaner architecture

---

# Disadvantages

- Harder tracing without logs
- Too many events can become noisy
- Needs proper naming convention
- Infinite event loops must be prevented

---

# Where It Is Used

- OpenCode
- Antigravity CLI
- VS Code
- IntelliJ Platform
- Chrome DevTools
- Electron Applications
- Large AI Agents
- Game Engines
- Enterprise Systems

---

# Summary

The Event Bus is the **central nervous system** of a modern AI CLI application.

Every major subsystem, including the Agent Engine, TUI, MCP, Plugins, Skills, Themes, Logger, File Watcher, and Session Manager, should communicate through the Event Bus instead of calling each other directly.

This architecture provides:

- Loose coupling
- High scalability
- Better maintainability
- Easier debugging
- Plugin extensibility
- Production-grade reliability

For an OpenCode or Antigravity-style coding agent, the Event Bus is one of the most critical foundational components and should be designed before implementing higher-level features.