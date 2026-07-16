# OpenChat CLI

# Model Management System

Version: 1.0

---

# Overview

The Model Management System allows users to build and manage their own AI model library.

Unlike other CLI tools that automatically fetch models from providers every time, OpenChat CLI maintains a local model registry controlled by the user.

A model always belongs to exactly one provider.

This approach provides:

• Faster model switching

• Offline model library

• Custom display names

• Better organization

• Future capability metadata

---

# Design Goals

✓ Unlimited Models

✓ Unlimited Providers

✓ Instant Switching

✓ Search

✓ Favorites

✓ Categories

✓ Capability Detection

✓ No JSON Editing

✓ Mobile Friendly

---

# Architecture

Provider

↓

Model

↓

Capabilities

↓

Current Model

↓

Conversation

---

# Database Schema

models

id

provider_id

model_id

display_name

description

category

supports_streaming

supports_tools

supports_reasoning

supports_vision

supports_embeddings

supports_json_mode

supports_function_calling

max_context

max_output_tokens

favorite

enabled

created_at

updated_at

---

# Model Identity

Every model contains

Model ID

Display Name

Provider

Capabilities

Status

Description

---

# Example

Model ID

deepseek-v4

Display Name

DeepSeek V4

Provider

OpenCode Zen

Description

Fast reasoning model

---

# Model Categories

General

Coding

Reasoning

Vision

Embedding

Audio

Experimental

Custom

---

# Add Model

Command

/add model

---

# Add Model Dialog

Model ID

[__________________]

Display Name

[__________________]

Provider

▼ OpenCode Zen

Description

[__________________]

Category

▼ Coding

────────────────────

Capabilities

☑ Streaming

☑ Tool Calling

☑ Reasoning

☐ Vision

☐ Embeddings

☑ JSON Mode

────────────────────

Maximum Context

128000

Maximum Output

8192

────────────────────

Save

Cancel

---

# Validation

Model ID

Required

Unique within provider

Maximum 100 characters

Provider

Required

Must exist

Display Name

Optional

If empty

Display Name = Model ID

---

# Automatic Validation

Before saving

Send

GET

/models

or

Check Provider

If model exists

↓

Offer

Import Capabilities

or

Manual Configuration

---

# Model List

Command

/all models

---

Layout

Search...

────────────────────

★ DeepSeek V4

OpenCode Zen

Coding

128K

────────────────────

GPT-5.5

OpenAI

Reasoning

200K

────────────────────

Gemini Pro

Google

Vision

1M

────────────────────

Claude Opus

Anthropic Proxy

Coding

200K

────────────────────

---

# Active Model

The active model is highlighted.

Example

● DeepSeek V4

OpenCode Zen

---

# Model Details

Selecting a model displays

Display Name

Model ID

Provider

Description

Category

Capabilities

Context Size

Output Limit

Created

Updated

---

# Model Actions

Activate

Edit

Delete

Duplicate

Favorite

Disable

Export

Copy ID

View Provider

---

# Edit Model

Users may edit

Display Name

Description

Category

Capabilities

Limits

Provider

---

# Delete Model

Confirmation

Delete

Cancel

If the model is active

Automatically ask user

Select another model

---

# Duplicate Model

Creates

DeepSeek V4 Copy

Useful for

Testing

Different configurations

---

# Favorite Models

Users may favorite models.

Example

★

DeepSeek

★

GPT-5.5

Favorites appear first.

---

# Search

Supports

Display Name

Model ID

Provider

Category

Capabilities

Instant filtering.

---

# Sorting

By Name

By Provider

By Category

By Context

By Recently Used

By Favorites

---

# Capabilities

Streaming

Tool Calling

Reasoning

Vision

Embeddings

JSON Mode

Function Calling

Image Input

Audio Input

Audio Output

Code Interpreter

Computer Use

Future

Web Search

MCP Native

---

# Context Size

Store

4096

8192

32000

128000

200000

1000000

Display

128K

200K

1M

---

# Maximum Output

Store

1024

2048

4096

8192

16384

---

# Provider Relationship

One Provider

↓

Many Models

Deleting Provider

↓

Warning

↓

Delete Models?

↓

Yes

No

---

# Current Model

Only one model may be active.

Switching model

↓

Immediately affects

New requests

Existing conversations continue normally.

---

# Quick Switch

Command

/model

↓

Search

↓

Select

↓

Done

No restart.

---

# Recent Models

Recently used models appear first.

Example

Recent

DeepSeek V4

GPT-5.5

Claude Opus

---

# Empty State

No Models Found

Run

/add model

---

# Import Models

Supported

JSON

Future

YAML

CSV

Remote Registry

---

# Export Models

Export

Selected

Favorites

All

API Keys are never exported.

---

# Health

Each model stores

Available

Unavailable

Unknown

Disabled

---

# Indicators

● Active

★ Favorite

⚠ Disabled

✖ Missing Provider

---

# Automatic Discovery

Future

Provider

↓

GET /models

↓

Import Wizard

↓

User Selects

↓

Save

---

# Compatibility

Every model must inherit authentication from its provider.

Models never store API keys.

---

# Performance

Model Switch

<100ms

Search

Instant

Load

<50ms

---

# Future Features

Model Tags

Cost Per Token

Speed Rating

Quality Rating

Community Ratings

Benchmarks

Recommended Tasks

Automatic Capability Detection

Model Templates

Cloud Sync

Usage Statistics

Token Cost Calculator

---

# Lifecycle

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

Monitor

↓

Edit

↓

Delete

---

# Design Principles

The Model Library should behave like an app library on a smartphone.

Users should be able to search, organize, favorite, categorize, and instantly switch between models without touching configuration files.

Every model is independent, belongs to a provider, and exposes its capabilities so the rest of the system can automatically enable or disable features such as tool calling, vision, reasoning, or structured JSON output.

The result is a clean, scalable, and provider-independent model management experience that remains simple enough to use comfortably inside Termux on Android.