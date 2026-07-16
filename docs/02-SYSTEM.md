# OpenChat CLI

# System Design Specification

Version: 1.0

---

# Purpose

This document defines the complete internal system architecture of OpenChat CLI.

The system is designed to be modular, scalable, provider-independent, and mobile-first.

Every module should work independently while communicating through the Core Engine.

---

# High Level Architecture

                    User
                      │
                      ▼
             OpenChat CLI
                      │
     ┌────────────────────────────────┐
     │          Core Engine           │
     └────────────────────────────────┘
          │      │      │      │
          ▼      ▼      ▼      ▼
     Providers Models Sessions Settings
          │
          ▼
      Agent Engine
          │
     ┌────┴────────────────────────┐
     │                             │
     ▼                             ▼
 Context Manager             Tool Manager
     │                             │
     └──────────────┬──────────────┘
                    ▼
             OpenAI API Engine
                    │
                    ▼
          Selected Provider
                    │
                    ▼
              Selected Model
                    │
                    ▼
             Streaming Response
                    │
                    ▼
                 Terminal

---

# Core Modules

The application consists of the following modules.

Core Engine

Provider Manager

Model Manager

Agent Manager

Session Manager

Settings Manager

Storage Engine

Database Layer

API Engine

Streaming Engine

Context Manager

Tool Manager

Permission Manager

Plugin Manager

MCP Manager

Logger

Theme Manager

Command Palette

---

# Core Engine

The Core Engine controls the entire application.

Responsibilities

• initialize application

• load settings

• load providers

• load models

• load active session

• initialize database

• initialize plugins

• initialize MCP

• initialize tools

• initialize command palette

• initialize AI engine

Every other module communicates through the Core Engine.

---

# Provider Manager

Responsible for

Adding providers

Editing providers

Deleting providers

Testing providers

Importing providers

Exporting providers

Selecting default provider

Each provider stores

Provider Name

Base URL

API Key

Status

Latency

Created Date

Updated Date

Example

OpenCode Zen

https://example.com/v1

sk-xxxxxxxx

---

# Model Manager

Responsible for

Adding models

Editing models

Deleting models

Selecting current model

Filtering models

Searching models

Each model belongs to exactly one provider.

Example

DeepSeek V4

↓

OpenCode Zen

GPT-5.5

↓

OpenAI

Gemini 3

↓

OpenRouter

---

# Session Manager

Stores every conversation.

Features

Resume

Fork

Export

Rename

Delete

Archive

Favorite

Each session stores

Title

Messages

Files

Tools Used

Model

Provider

Created Date

Updated Date

---

# Agent Manager

Agents define AI behavior.

Example

General

Laravel

React

Node

Python

Flutter

Security

Reviewer

Architect

Debugger

DevOps

Each agent contains

Name

Description

System Prompt

Allowed Tools

Temperature

Reasoning Level

Preferred Skills

---

# Skill Manager

Skills are reusable instruction libraries.

Example

Laravel

React

Bootstrap

Tailwind

Docker

Git

Testing

Security

Performance

A skill contains

Instructions

Examples

Templates

Best Practices

---

# Tool Manager

Responsible for tool execution.

Built-in tools

Read File

Write File

Replace File

Delete File

Search

Glob

Grep

Bash

Git

Clipboard

HTTP

Download

Todo

Diff

Patch

MCP

Browser

Each tool returns structured results.

---

# Permission Manager

Every tool requires permission.

Permission Levels

Allow Always

Allow Once

Ask

Deny

Example

Bash

Ask

Delete File

Ask

Read File

Always Allow

---

# API Engine

Responsible for all AI communication.

Supports

Chat Completions

Streaming

Tool Calling

Embeddings (future)

Reasoning Models

Vision Models

The API Engine only understands the OpenAI API format.

It never depends on a specific provider.

---

# Context Manager

Responsible for building prompts.

Collects

Conversation

Files

Git Changes

Workspace

Skills

Agent Prompt

System Prompt

User Prompt

Everything is combined before sending to the AI.

---

# Streaming Engine

Responsibilities

Receive tokens

Render tokens

Markdown rendering

Code block rendering

Tool execution events

Interrupt handling

Retry handling

---

# Storage Engine

Responsible for

Reading

Writing

Caching

Indexing

Auto Save

History

Bookmarks

Recent Files

---

# Database Layer

SQLite

Tables

providers

models

sessions

messages

agents

skills

settings

plugins

history

permissions

mcp_servers

logs

---

# Logger

Stores

Errors

Warnings

Requests

Responses

Latency

Tool Calls

Performance

Logs should never contain API keys.

---

# Plugin Manager

Future extension system.

Plugins may add

Commands

Agents

Skills

Themes

Tools

Panels

MCP Servers

---

# MCP Manager

Responsible for

Connecting MCP

Disconnecting MCP

Tool Discovery

Server Health

Authentication

Permission

Multiple MCP servers should work simultaneously.

---

# Theme Manager

Controls

Colors

Borders

Icons

Syntax Highlighting

Markdown

Prompt Style

Status Line

Cursor

---

# Command Palette

Opened by typing

/

Responsible for

Searching commands

Running commands

Filtering

Keyboard navigation

No mouse support required.

---

# Application Startup

Launch

↓

Load Settings

↓

Initialize Database

↓

Load Providers

↓

Load Models

↓

Load Plugins

↓

Load MCP

↓

Initialize Tools

↓

Load Theme

↓

Load Previous Session

↓

Display Home Screen

---

# Chat Request Flow

User Prompt

↓

Context Builder

↓

Current Session

↓

Current Agent

↓

Current Skills

↓

Workspace Analysis

↓

Tool Planning

↓

API Request

↓

Streaming

↓

Tool Execution

↓

Continue Conversation

↓

Save Session

---

# Error Handling

If Provider Offline

↓

Retry

↓

Fallback

↓

Display Error

↓

Keep Session Alive

The application should never crash because of an API failure.

---

# Design Principles

Every module must be

Independent

Reusable

Replaceable

Testable

Well Documented

Strongly Typed

Asynchronous

Plugin Friendly

---

# Performance Goals

Startup

< 1 second

Command Palette

< 50 ms

Provider Switch

< 100 ms

Model Switch

< 100 ms

Streaming Delay

< 200 ms

Tool Execution

Real Time

Database

SQLite optimized

Memory Usage

Minimal

---