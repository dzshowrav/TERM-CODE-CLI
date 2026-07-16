# OpenChat CLI

# Agent System Specification

Version: 1.0

---

# Overview

The Agent System is the brain of OpenChat CLI.

An Agent is an AI personality that controls how the model thinks, behaves, plans, uses tools, communicates, and solves problems.

The selected model provides intelligence.

The selected agent provides behavior.

Changing the model does not change the agent.

Changing the agent does not change the model.

Both are independent.

---

# Design Goals

✓ Unlimited Agents

✓ User Creatable

✓ Provider Independent

✓ Model Independent

✓ Reusable

✓ Shareable

✓ Versioned

✓ Import / Export

✓ Plugin Ready

---

# Architecture

Provider

↓

Model

↓

Agent

↓

Skills

↓

Context

↓

Conversation

↓

Tool Execution

↓

Response

---

# Database Schema

agents

id

name

description

icon

system_prompt

temperature

top_p

reasoning_level

preferred_model

preferred_provider

allowed_tools

default_skills

max_context

color

enabled

built_in

created_at

updated_at

---

# Agent Components

Every agent contains

• Name

• Description

• Icon

• System Prompt

• Default Skills

• Allowed Tools

• Preferred Model

• Preferred Provider

• Reasoning Level

• Creativity

• Context Size

---

# Built-in Agents

General Assistant

Laravel Specialist

PHP Expert

JavaScript Expert

TypeScript Expert

React Engineer

Vue Engineer

Next.js Engineer

Node.js Engineer

Python Engineer

Flutter Engineer

Android Engineer

Rust Engineer

Go Engineer

Docker Expert

DevOps Engineer

Database Architect

System Architect

Security Auditor

Performance Optimizer

Bug Hunter

Code Reviewer

Documentation Writer

Technical Writer

UI/UX Designer

Testing Engineer

AI Engineer

Prompt Engineer

Git Expert

Linux Expert

---

# General Agent

Purpose

General programming

Capabilities

Coding

Explaining

Debugging

Documentation

Tool Usage

All

---

# Laravel Agent

Purpose

Enterprise Laravel development

Automatically prioritizes

Laravel

PHP

Blade

Eloquent

Artisan

Composer

Routes

Middleware

Validation

Policies

Queues

Events

API Resources

---

# React Agent

Optimized for

React

Next.js

Hooks

TypeScript

Performance

Accessibility

Component Design

---

# Security Agent

Focus

Security Audit

OWASP

Authentication

Authorization

Encryption

SQL Injection

XSS

CSRF

Dependency Analysis

---

# Reviewer Agent

Responsibilities

Review Code

Suggest Improvements

Detect Bugs

Performance Issues

Security Issues

Architecture Problems

Never modifies files automatically.

---

# Agent Selection

Command

/agents

↓

Search

↓

Select

↓

Agent Activated

No restart required.

---

# Agent List

General

Laravel Specialist

React Engineer

Flutter Engineer

Security Auditor

Performance Optimizer

Code Reviewer

Database Architect

---

# Agent Details

Selecting an agent displays

Name

Description

Capabilities

Preferred Skills

Preferred Models

Allowed Tools

Reasoning Level

Version

Author

---

# Create Agent

Command

/new agent

Fields

Agent Name

Description

System Prompt

Preferred Provider

Preferred Model

Allowed Tools

Default Skills

Reasoning Level

Temperature

Color

Icon

Save

---

# Edit Agent

Editable

Name

Prompt

Skills

Tools

Provider

Model

Reasoning

Color

Icon

---

# Delete Agent

Only allowed for

Custom Agents

Built-in agents cannot be deleted.

---

# Clone Agent

Creates

Laravel Specialist Copy

Useful for customization.

---

# Agent Prompt

Every request begins with

System Prompt

↓

Agent Prompt

↓

Skill Prompts

↓

Workspace Context

↓

Conversation

↓

User Prompt

---

# Agent Memory

Stores

Preferred Tools

Preferred Skills

Conversation Style

Planning Style

Reasoning Strategy

---

# Reasoning Levels

Fast

Balanced

Deep

Maximum

Higher reasoning may increase response time.

---

# Tool Permissions

Agents may restrict tools.

Example

Reviewer

Read Files

Git Diff

Search

No Delete

No Bash

---

# Preferred Skills

Example

Laravel Agent

Automatically loads

Laravel

PHP

Composer

MySQL

Blade

Security

Performance

---

# Preferred Models

Optional

Example

GPT-5.5

Claude Opus

DeepSeek V4

Gemini Pro

If unavailable

Current model is used.

---

# Agent Categories

General

Frontend

Backend

Mobile

AI

Database

Cloud

Security

Testing

DevOps

Documentation

Architecture

Custom

---

# Icons

Every agent has an icon.

Examples

💻 Developer

🛡 Security

⚛ React

🚀 Laravel

🐍 Python

📱 Flutter

⚙ DevOps

🧠 AI

---

# Agent Color

Each agent may define

Accent Color

Used for

Headers

Status

Selection

---

# Agent Export

Export

JSON

Includes

Prompt

Settings

Skills

Permissions

No API Keys

---

# Agent Import

Import

↓

Validate

↓

Install

↓

Available

---

# Agent Marketplace

Future

Community Agents

Official Agents

Verified Authors

Version Updates

Ratings

---

# Agent Versioning

Every agent stores

Version

Author

License

Last Updated

---

# Agent Lifecycle

Create

↓

Validate

↓

Save

↓

Activate

↓

Use

↓

Update

↓

Export

↓

Delete

---

# Performance Goals

Agent Switch

<100ms

Prompt Build

<20ms

Skill Loading

Lazy Loaded

Memory Usage

Minimal

---

# Future Features

Multi-Agent Collaboration

Sub Agents

Agent Teams

Voice Agents

Vision Agents

Learning Agents

Scheduled Agents

Workflow Agents

Autonomous Agents

Cloud Agent Sync

---

# Design Principles

Agents define *how* the AI behaves, not *which* AI is used.

A user should be able to pair any provider, any model, and any agent.

Example

OpenCode Zen
+
DeepSeek V4
+
Laravel Specialist

or

OpenAI
+
GPT-5.5
+
Security Auditor

or

Local Ollama
+
Qwen Coder
+
Code Reviewer

This separation of Provider → Model → Agent makes OpenChat CLI flexible, scalable, and completely vendor-independent.