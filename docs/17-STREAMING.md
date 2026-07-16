# OpenChat CLI

# Streaming & Response Rendering Engine Specification

Version: 1.0

---

# Overview

The Streaming Engine is responsible for delivering AI responses in real time.

Instead of waiting for the entire response to complete, OpenChat CLI renders content token-by-token while simultaneously handling tool calls, markdown rendering, syntax highlighting, and status updates.

The goal is to make the AI feel responsive, transparent, and alive.

---

# Design Goals

✓ Real-time Streaming

✓ Low Latency

✓ Smooth Rendering

✓ Markdown Aware

✓ Tool Aware

✓ Interruptible

✓ Mobile Optimized

✓ Zero Flicker

✓ Memory Efficient

✓ Provider Independent

---

# Architecture

User Prompt

↓

API Request

↓

Provider Stream

↓

Streaming Parser

↓

Markdown Parser

↓

Syntax Highlighter

↓

Terminal Renderer

↓

Screen Refresh

↓

User

---

# Streaming Pipeline

User

↓

Provider

↓

Receive Tokens

↓

Buffer Tokens

↓

Render Markdown

↓

Highlight Code

↓

Update Screen

↓

Complete Response

---

# Supported Stream Types

OpenAI Stream

Server-Sent Events (SSE)

Chunked HTTP

Future

WebSocket

gRPC

---

# Rendering States

Idle

Connecting

Streaming

Thinking

Executing Tools

Waiting

Retrying

Completed

Interrupted

Cancelled

Error

---

# Connection State

Display

● Connecting...

↓

✓ Connected

or

✖ Failed

Users should always know the current state.

---

# Thinking State

When supported

Display

Thinking...

Planning...

Analyzing Repository...

Reading Files...

Never leave the screen blank.

---

# Token Streaming

Every received token is immediately rendered.

Example

User

Create Laravel Login.

Assistant

Creating▌

Creating authentication▌

Creating authentication controller...▌

Done

---

# Streaming Cursor

Cursor

▌

Visible only while streaming.

Disappears automatically when complete.

---

# Streaming Speed

Modes

Real Time

Fast

Instant

Developer Mode

Character-by-character

Default

Real Time

---

# Buffer System

Small responses

↓

Immediate Rendering

Large responses

↓

Buffered Rendering

↓

Smooth Output

Avoid excessive terminal redraws.

---

# Markdown Rendering

Supported

Headers

Lists

Tables

Blockquotes

Code Blocks

Task Lists

Links

Bold

Italic

Inline Code

Horizontal Rules

Nested Lists

---

# Code Block Rendering

Requirements

Language Detection

Syntax Highlighting

Copy Support

Collapsible

Scrollable

Optional Line Numbers

Example

```typescript

function hello() {
    console.log("OpenChat");
}

```

---

# Syntax Highlighting

Supported

PHP

JS

TS

Python

Go

Rust

Java

Kotlin

Swift

C

C++

C#

Dart

HTML

CSS

JSON

YAML

SQL

Markdown

Shell

Auto detect if language missing.

---

# Progressive Markdown

Instead of waiting

Render

Heading

↓

Paragraph

↓

List

↓

Code Block

↓

Table

As content arrives.

---

# Tool Event Rendering

Example

✓ Reading composer.json

✓ Searching routes

✓ Updating AuthController.php

✓ Running composer install

✓ Completed

Tool events are visually separated from chat text.

---

# Tool Progress

Display

[1/5]

Reading Files

[2/5]

Searching Workspace

[3/5]

Writing Controller

[4/5]

Running Tests

[5/5]

Done

---

# Inline Status Messages

Examples

Loading...

Retrying...

Waiting...

Parsing...

Formatting...

Completed

Keep messages short.

---

# Interrupt

Shortcut

Ctrl+C

or

ESC

Result

Stop Streaming

Cancel Pending Tool

Keep Conversation

Allow Retry

---

# Retry

Command

/retry

Uses

Same Prompt

Same Context

Same Model

Same Agent

---

# Resume Streaming

If provider supports

Resume

↓

Continue Stream

Otherwise

Retry

---

# Multi-Tool Streaming

Support

Tool 1

↓

Response

↓

Tool 2

↓

Response

↓

Tool 3

↓

Final Answer

Streaming never pauses unexpectedly.

---

# Error Handling

Connection Lost

↓

Retry

↓

Reconnect

↓

Resume

↓

Fail Gracefully

---

# Error Display

Examples

Connection Timeout

Authentication Failed

Model Busy

Network Lost

Rate Limited

Every error includes

Reason

Suggested Fix

Retry Button

---

# Progress Indicators

Spinner

●

○

◐

◑

◒

◓

Progress Bar

████████░░

Percentage

68%

Choose automatically based on task.

---

# Long Response Handling

Large outputs

↓

Chunked Rendering

↓

Minimal Memory

↓

Stable FPS

Never freeze the terminal.

---

# Response Completion

When stream finishes

Remove Cursor

↓

Finalize Markdown

↓

Save Session

↓

Ready for Input

---

# Performance Optimizations

Batch small updates

Minimize redraws

Reuse buffers

Cache syntax colors

Lazy markdown parsing

Only repaint changed lines

---

# Mobile Optimizations

Optimized for

Termux

80 Columns

Portrait

Landscape

Low Memory Devices

Slow CPUs

---

# Accessibility

High Contrast

Screen Reader Friendly

Configurable Colors

Disable Animations

Reduced Motion

---

# Configuration

/settings

Streaming

Options

Streaming Enabled

Thinking Indicator

Typing Cursor

Syntax Highlighting

Markdown Rendering

Render Speed

Maximum FPS

Buffer Size

---

# Metrics

Track

Time to First Token

Streaming Speed

Average Tokens/sec

Render FPS

Dropped Frames

Response Time

---

# Future Features

Live Diff Rendering

Streaming Images

Voice Streaming

Vision Streaming

Multi-Agent Streaming

Split Response View

Streaming Charts

Collaborative Streaming

---

# Performance Goals

First Token

<500ms

Render Delay

<10ms

Screen Refresh

60 FPS Equivalent

Markdown Parsing

Incremental

Memory Usage

Minimal

CPU Usage

Optimized

---

# Design Principles

The Streaming Engine defines how OpenChat CLI feels.

A fast model with poor rendering still feels slow.

A smooth streaming experience makes the AI feel responsive, intelligent, and interactive.

Every token, tool event, and status update should appear naturally, giving developers confidence that the system is actively working while remaining lightweight enough to perform exceptionally well on Android Termux, Linux, macOS, and Windows.