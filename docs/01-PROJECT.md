# OpenChat CLI

> The Ultimate Universal AI Coding Agent for the Terminal

Version: 1.0 (Project Specification)

Author: Your Name

---

# Vision

OpenChat CLI is a next-generation terminal-based AI coding assistant designed for developers who want complete freedom over their AI providers.

Unlike traditional AI coding tools that lock users into specific providers, OpenChat CLI supports **any OpenAI-compatible API** through a simple Provider Management system.

Users can connect:

- OpenAI
- OpenRouter
- DeepSeek
- Groq
- Together AI
- LiteLLM
- OpenCode Zen
- Ollama
- LM Studio
- vLLM
- Azure OpenAI
- Self-hosted APIs
- Any OpenAI-compatible endpoint

without modifying configuration files manually.

Everything is managed inside the CLI.

---

# Core Philosophy

OpenChat CLI follows five principles.

## 1. Provider Agnostic

The application never assumes a specific AI provider.

If the endpoint supports the OpenAI API specification, it works.

---

## 2. Chat First

Everything revolves around conversation.

The terminal becomes an intelligent workspace.

No complicated commands.

Simply chat.

---

## 3. Keyboard Driven

Every feature must be usable without touching a mouse.

Navigation should be extremely fast inside Termux.

---

## 4. Mobile Friendly

The UI is optimized primarily for Android + Termux.

Desktop support is secondary.

---

## 5. Transparent AI

Users should always know

- what the AI is doing
- what tool is executing
- what files are changing
- which model is responding
- which provider is connected

No hidden actions.

---

# Goals

Build an AI coding assistant that combines the best parts of

- OpenCode
- Claude Code
- Codex CLI
- Gemini CLI

while introducing a universal provider architecture.

---

# Main Features

## Universal Provider Management

Users can register unlimited providers.

Each provider contains

- Name
- Base URL
- API Key

Example

OpenCode Zen

https://example.com/v1

sk-xxxxxxxx

---

## Model Library

Users create their own model list.

Each model belongs to a provider.

Example

GPT-5.5

↓

OpenAI

DeepSeek V4

↓

OpenCode Zen

Gemini 3

↓

OpenRouter

---

## Interactive Chat

Streaming responses

Markdown rendering

Syntax highlighting

Tool execution

Code generation

Refactoring

Debugging

Documentation

---

## Coding Agent

The assistant can

Read files

Create files

Edit files

Delete files

Search code

Run commands

Review code

Generate documentation

Fix bugs

Execute tools

Use MCP servers

---

## Session Management

Unlimited chat sessions.

Resume conversations.

Export conversations.

Fork sessions.

Rename sessions.

Delete sessions.

---

## Tool System

Built-in tools include

Filesystem

Bash

Git

Search

Replace

HTTP

Clipboard

Downloads

MCP

Task Runner

Todo Manager

---

# User Interface Philosophy

No sidebar.

No floating windows.

No complex layouts.

The application uses only three major interfaces.

## Home Screen

ASCII logo

Prompt input

Provider

Current model

Agent

Workspace

Tips

Footer

---

## Chat Screen

Conversation

Streaming

Tool execution

Prompt input

Status line

---

## Command Palette

Everything is accessible from "/"

Examples

/provider api

/add model

/all models

/settings

/history

/help

---

# Target Platforms

Primary

Android (Termux)

Secondary

Linux

macOS

Windows

---

# Technology Stack

Language

TypeScript

Runtime

Node.js

UI

Ink React

Storage

SQLite

HTTP

Fetch API

Markdown

Marked

Syntax Highlighting

Shiki

Database

better-sqlite3

Package Manager

npm

---

# Project Structure

OpenChat CLI

↓

Core Engine

↓

Provider Manager

↓

Model Manager

↓

Agent Engine

↓

Tool System

↓

OpenAI Compatible API

↓

AI Model

↓

Streaming Response

↓

Terminal

---

# User Workflow

Launch CLI

↓

Select Provider

↓

Select Model

↓

Open Project

↓

Start Chat

↓

AI Reads Context

↓

Uses Tools

↓

Edits Files

↓

Shows Diff

↓

Finish

---

# Competitive Advantages

✔ Unlimited Providers

✔ Unlimited Models

✔ Interactive Provider Manager

✔ Interactive Model Manager

✔ OpenAI-Compatible Architecture

✔ Mobile First UX

✔ Keyboard Driven

✔ Tool Execution Transparency

✔ Streaming Responses

✔ SQLite Storage

✔ Plugin Ready

✔ MCP Support

✔ Session History

✔ Multiple Agents

✔ Skills

✔ Slash Commands

✔ Enterprise Architecture

---

# Project Mission

Build the most flexible and developer-friendly terminal AI coding assistant.

The application should never restrict users to a specific AI company.

Instead, it should provide a universal platform where developers can connect any OpenAI-compatible provider, organize models, switch between them instantly, and work inside a clean, keyboard-first, mobile-optimized terminal experience.