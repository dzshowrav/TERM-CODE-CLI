# OpenChat CLI

# Chat Engine Specification

Version: 1.0

---

# Overview

The Chat Engine is the heart of OpenChat CLI.

Every interaction between the user and the AI passes through the Chat Engine.

Its responsibilities include

• Conversation management

• Context building

• Prompt construction

• Streaming

• Tool execution

• Reasoning

• Session memory

• Interrupt handling

• Response rendering

The Chat Engine should be provider-independent and work with any OpenAI-compatible API.

---

# Philosophy

The Chat Engine should feel like talking to an intelligent software engineer instead of sending isolated prompts.

Every response should consider

Current Request

Conversation History

Workspace

Agent

Skills

Tools

Permissions

Model Capabilities

---

# Request Pipeline

User Prompt

↓

Input Validation

↓

Conversation Update

↓

Workspace Analysis

↓

Agent Loading

↓

Skill Loading

↓

Context Builder

↓

Tool Planning

↓

Provider Request

↓

Streaming Response

↓

Tool Calls

↓

Continue Response

↓

Save Session

↓

Ready

---

# Conversation Structure

Every message contains

Role

Content

Timestamp

Model

Provider

Agent

Tool Calls

Attachments

Token Count

Example

User

↓

Build Laravel Authentication

Assistant

↓

Planning...

↓

Reading routes

↓

Editing controller

↓

Done

---

# Message Roles

System

Developer

Agent

Skill

User

Assistant

Tool

Reasoning

Only User and Assistant appear in chat.

The remaining roles are internal.

---

# Conversation Memory

Every conversation stores

Messages

Files

Workspace

Git Branch

Selected Model

Selected Agent

Skills

Tool History

Context Usage

Token Usage

---

# Context Builder

Before every request

Collect

↓

System Prompt

↓

Agent Prompt

↓

Skill Instructions

↓

Conversation History

↓

Workspace Context

↓

File Context

↓

Git Context

↓

User Prompt

↓

Final Request

---

# Context Priority

Highest

System Prompt

↓

Agent Prompt

↓

Skill Rules

↓

Workspace

↓

Conversation

↓

User Prompt

---

# Context Optimization

Never send unnecessary information.

Only include

Relevant Files

Recent Messages

Required Skills

Required Git Changes

Current Task

Goal

Reduce token usage while preserving quality.

---

# Streaming Engine

Responses must stream immediately.

Never wait for complete responses.

Stages

Connecting...

↓

Thinking...

↓

Generating...

↓

Executing Tools...

↓

Finalizing...

↓

Done

---

# Streaming Cursor

During generation

Display

▌

Example

Creating authentication▌

The cursor disappears after completion.

---

# Thinking Mode

If supported by the model

Display

Thinking...

Planning...

Analyzing Project...

Thinking content should never be editable.

---

# Tool Calling

The AI may request tools.

Example

Read File

↓

Permission Check

↓

Execute

↓

Return Result

↓

Continue Response

The chat must continue automatically after the tool returns.

---

# Multiple Tool Calls

Example

Read File

↓

Search Project

↓

Run Bash

↓

Write File

↓

Git Diff

↓

Final Response

Support sequential and parallel execution where safe.

---

# Tool Status

Display inline

✓ Reading routes/web.php

✓ Reading composer.json

✓ Updating AuthController.php

✓ Running composer install

✓ Tests Passed

---

# Markdown Rendering

Support

Headers

Lists

Tables

Task Lists

Code Blocks

Quotes

Links

Bold

Italic

Inline Code

Horizontal Rules

---

# Code Blocks

Requirements

Syntax Highlighting

Copy Button

Language Detection

Line Numbers (optional)

Scrollable

Example

```php
Route::get('/login');
```

---

# Interrupt Handling

User presses

Ctrl+C

or

ESC

↓

Stop Streaming

↓

Keep Conversation

↓

Allow Retry

No data loss.

---

# Retry

Command

/retry

Resends

Last Prompt

Current Context

Current Agent

Current Model

---

# Edit Previous Prompt

Users may edit

Their previous prompt

↓

Resend

↓

Continue Conversation

---

# Response Actions

Every assistant message supports

Copy

Retry

Export

Explain

Continue

Summarize

Regenerate

---

# Auto Title Generation

After the first exchange

Generate

Session Title

Example

"Laravel Authentication"

"Fix Docker Build"

"React Dashboard"

---

# Attachments

Future Support

Images

PDF

Text Files

Logs

Screenshots

Audio

---

# Workspace Awareness

Before responding

Detect

Language

Framework

Package Manager

Git

Configuration

Project Structure

The AI should automatically adapt.

---

# Conversation Modes

Chat

Coding

Review

Debug

Explain

Planning

Architecture

Documentation

Future

Voice

Vision

Autonomous

---

# Token Management

Track

Input Tokens

Output Tokens

Context Tokens

Remaining Context

Display

Context

52%

128K

---

# Long Conversation Handling

When context becomes large

Summarize Older Messages

↓

Compress Context

↓

Keep Important Information

↓

Continue Chat

The user should never lose important context.

---

# Error Handling

Connection Lost

↓

Retry

↓

Continue

Provider Error

↓

Show Error

↓

Keep Session

Tool Error

↓

Display

↓

Continue Conversation

---

# Conversation Search

Future

Search

Messages

Code

Files

Tool Calls

Dates

Keywords

---

# Export

Supported Formats

Markdown

JSON

HTML

PDF

Conversation only

or

Conversation + Tool History

---

# Import

Future

Markdown

JSON

Conversation Backup

---

# Performance Goals

First Token

<500ms

Streaming

Real Time

Markdown Rendering

Instant

Interrupt

Immediate

Retry

<100ms

---

# Future Features

Conversation Branching

Pinned Messages

Bookmarks

Shared Sessions

Cloud Sync

Live Collaboration

Voice Chat

Vision Chat

Multi-Agent Conversations

Memory Profiles

---

# Design Principles

The Chat Engine should make conversations feel natural, continuous, and workspace-aware.

The user should never feel like they are repeatedly starting from scratch.

Every interaction should build on previous knowledge while remaining transparent about tool usage, context, and AI actions.

The Chat Engine is not just a messaging system.

It is the intelligent orchestration layer that transforms OpenChat CLI into a true AI coding partner.