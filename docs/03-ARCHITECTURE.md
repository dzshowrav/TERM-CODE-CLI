# OpenChat CLI

# Software Architecture Specification

Version: 1.0

---

# Purpose

This document defines the internal software architecture of OpenChat CLI.

The architecture follows a modular, event-driven, plugin-first design where every component is isolated, reusable, and replaceable.

The application must never depend on a specific AI provider. Every request flows through a unified OpenAI-compatible API layer.

---

# Architectural Principles

The entire application follows these principles.

• Modular
• Event Driven
• Plugin Ready
• OpenAI Compatible
• Mobile First
• Keyboard First
• Offline Friendly
• Cross Platform
• Strongly Typed
• Async by Default

---

# Layered Architecture

                    UI Layer
                        │
                        ▼
                Command Layer
                        │
                        ▼
                Service Layer
                        │
                        ▼
                 Engine Layer
                        │
                        ▼
              Provider API Layer
                        │
                        ▼
                 Storage Layer

Every layer only communicates with adjacent layers.

---

# Complete Folder Structure

openchat-cli/

├── src/
│
├── core/
│   ├── engine/
│   ├── events/
│   ├── state/
│   ├── services/
│   ├── logger/
│   ├── config/
│   └── constants/
│
├── ui/
│   ├── screens/
│   ├── components/
│   ├── dialogs/
│   ├── markdown/
│   ├── syntax/
│   └── theme/
│
├── providers/
│
├── models/
│
├── agents/
│
├── skills/
│
├── commands/
│
├── tools/
│
├── mcp/
│
├── plugins/
│
├── database/
│
├── sessions/
│
├── api/
│
├── utils/
│
└── index.ts

---

# Engine Hierarchy

Application Engine

↓

Workspace Engine

↓

Session Engine

↓

Context Engine

↓

Tool Engine

↓

Provider Engine

↓

Streaming Engine

↓

Renderer

Each engine has one responsibility.

---

# Event Bus

Every module communicates using events.

Example Events

app.started

provider.changed

model.changed

session.created

session.deleted

message.sent

message.received

tool.started

tool.finished

tool.failed

stream.started

stream.finished

workspace.changed

plugin.loaded

plugin.unloaded

---

# State Manager

The application keeps only one global state.

Example

Application

↓

Workspace

↓

Provider

↓

Model

↓

Agent

↓

Session

↓

Conversation

↓

Tool Status

↓

Streaming Status

Every UI component reads from this state.

---

# Dependency Flow

UI

↓

Commands

↓

Services

↓

Engine

↓

Storage

Never allow

Storage

↓

UI

or

API

↓

UI

Everything passes through the Engine.

---

# Workspace Architecture

Workspace

↓

Files

↓

Git Repository

↓

Language Detection

↓

Package Manager

↓

Project Metadata

↓

Index

↓

Context Builder

---

# Conversation Architecture

User Message

↓

Conversation Manager

↓

Agent Prompt

↓

Skill Prompt

↓

Workspace Context

↓

File Context

↓

History

↓

API Request

↓

Streaming Response

↓

Tool Execution

↓

Conversation Saved

---

# Request Lifecycle

User enters prompt

↓

Prompt validated

↓

Conversation updated

↓

Context collected

↓

Workspace analyzed

↓

Agent loaded

↓

Skills loaded

↓

API request created

↓

Provider selected

↓

Model selected

↓

Streaming starts

↓

Tool calls executed

↓

Response rendered

↓

Conversation stored

---

# Provider Lifecycle

Provider Created

↓

API Validation

↓

Health Check

↓

Saved

↓

Activated

↓

Used

↓

Updated

↓

Deleted

Deleting a provider automatically disables dependent models until reassigned.

---

# Model Lifecycle

Create

↓

Validate

↓

Assign Provider

↓

Save

↓

Activate

↓

Use

↓

Update

↓

Delete

---

# Session Lifecycle

New Session

↓

Title Generated

↓

Messages Stored

↓

Files Indexed

↓

Resume

↓

Archive

↓

Export

↓

Delete

---

# Tool Lifecycle

AI Requests Tool

↓

Permission Check

↓

Execute

↓

Capture Output

↓

Return Result

↓

Continue AI Response

---

# Permission Flow

Tool Requested

↓

Policy Check

↓

Always Allow

↓

Execute

OR

Ask User

↓

Allow Once

↓

Execute

OR

Denied

↓

Cancel Tool

---

# API Lifecycle

Create Request

↓

Inject System Prompt

↓

Inject Agent Prompt

↓

Inject Skills

↓

Inject Conversation

↓

Inject Tool Schema

↓

Send Request

↓

Receive Stream

↓

Render Output

↓

Handle Tool Calls

↓

Finish

---

# Storage Architecture

SQLite

↓

Repositories

↓

Services

↓

Application

The UI never communicates directly with SQLite.

---

# Rendering Pipeline

Markdown

↓

Syntax Highlight

↓

ANSI Colors

↓

Terminal Renderer

↓

Screen Refresh

---

# Logging Pipeline

Action

↓

Logger

↓

Log Formatter

↓

SQLite

↓

Log Viewer

Sensitive information such as API keys must never be written to logs.

---

# Plugin Architecture

Plugin

↓

Manifest

↓

Register

↓

Commands

↓

Skills

↓

Agents

↓

Themes

↓

Tools

↓

Unload

Plugins should not modify the core source code.

---

# Theme Architecture

Theme

↓

Palette

↓

Borders

↓

Icons

↓

Markdown Colors

↓

Syntax Colors

↓

Prompt Style

↓

Status Line

Themes only affect presentation.

---

# Configuration Hierarchy

Default Settings

↓

Global Configuration

↓

Workspace Configuration

↓

Session Overrides

↓

Runtime Changes

Lower levels override higher levels.

---

# Error Recovery

If a request fails

↓

Retry

↓

Fallback Provider (optional)

↓

Show Error

↓

Keep Session Active

The application should recover gracefully whenever possible.

---

# Security Boundaries

User Input

↓

Validation

↓

Permission System

↓

Tool Execution

↓

Filesystem

↓

Operating System

Never execute shell commands without passing through the permission layer.

---

# Performance Objectives

Application Startup

< 1 second

Command Palette

< 50 ms

Provider Switch

< 100 ms

Model Switch

< 100 ms

Screen Refresh

60 FPS equivalent

Streaming Latency

Minimal

Memory Usage

Optimized for Android Termux

---

# Future Scalability

Multi-Agent Collaboration

↓

Distributed Workers

↓

Remote Workspaces

↓

Cloud Synchronization

↓

Plugin Marketplace

↓

Web Dashboard

↓

Voice Interaction

↓

Vision Models

↓

Image Generation

↓

Enterprise Collaboration

---

# Architecture Summary

OpenChat CLI is built around a clean separation of concerns.

The UI displays information.

The Engine coordinates behavior.

The Service Layer performs logic.

The Storage Layer persists data.

The Provider Layer communicates with AI.

Every feature is modular, replaceable, and extensible, allowing the project to grow from a lightweight Termux coding assistant into a full-featured universal AI development platform without redesigning its foundation.