# OpenChat CLI

# Provider Management System

Version: 1.0

---

# Overview

The Provider Management System is one of the core innovations of OpenChat CLI.

Unlike traditional AI coding assistants that ship with fixed providers, OpenChat CLI allows users to register, manage, test, and switch between unlimited OpenAI-compatible providers directly from the CLI.

The application does not know or care whether the provider is:

• OpenAI
• OpenRouter
• Groq
• DeepSeek
• Together AI
• Azure OpenAI
• LiteLLM
• Ollama
• LM Studio
• vLLM
• OpenCode Zen
• Self-hosted API
• Any future OpenAI-compatible service

If the provider implements the OpenAI API specification, OpenChat CLI can use it.

---

# Design Goals

✓ Unlimited providers

✓ No configuration file editing

✓ Interactive UI

✓ Secure API key storage

✓ Instant provider switching

✓ Connection testing

✓ Health monitoring

✓ Import / Export

✓ Mobile-friendly workflow

---

# Provider Structure

Every provider contains:

Provider Name

Base URL

API Key

Description (optional)

Status

Latency

Created At

Updated At

Example

Name

OpenCode Zen

Base URL

https://example.com/v1

API Key

sk-xxxxxxxxxxxx

Description

Personal AI Proxy

---

# Database Schema

providers

id

name

base_url

api_key

description

status

latency

created_at

updated_at

---

# UI Workflow

Command

/provider api

↓

Add Provider Dialog

↓

Validate Input

↓

Test Connection

↓

Save

↓

Provider Available

---

# Add Provider Screen

Provider Name

[________________]

Base URL

[________________]

API Key

[****************]

Description

[________________]

────────────────────────

Save

Cancel

---

# Validation Rules

Provider Name

• Required

• Unique

• Maximum 50 characters

Base URL

• Required

• Valid URL

• Must use HTTP or HTTPS

• Remove trailing slash automatically

API Key

• Required

• Hidden while typing

Description

• Optional

---

# Connection Test

Before saving, OpenChat CLI automatically verifies the provider.

Request

GET

/models

or

POST

/chat/completions

Expected Result

200 OK

↓

Provider Saved

If verification fails

Show error

Do not save invalid providers.

---

# Provider List

Command

/providers

Displays

● OpenCode Zen

Connected

Latency: 312ms

○ OpenAI

Disconnected

○ Local Ollama

Connected

Latency: 18ms

---

# Provider Details

Selecting a provider opens

Name

Base URL

Description

Status

Latency

Number of Models

Last Connected

Created Date

Updated Date

---

# Provider Actions

Every provider supports

Edit

Delete

Duplicate

Export

Import

Test Connection

Set Default

View Models

---

# Edit Provider

Allows editing

Provider Name

Base URL

API Key

Description

After saving

Automatically re-test connection.

---

# Delete Provider

Before deletion

Check

Does this provider own models?

If yes

Warn user

"This provider has associated models."

Options

Delete Provider Only

Delete Provider + Models

Cancel

---

# Duplicate Provider

Creates a copy.

Useful for

Multiple API keys

Different environments

Testing

---

# Export Provider

Exports

Provider Name

Base URL

Description

Does NOT export API Key unless user explicitly enables it.

Default

API Key excluded.

---

# Import Provider

Supports

JSON

Import

↓

Validate

↓

Test

↓

Save

---

# Health Check

Every provider stores

Status

Latency

Last Success

Last Failure

Failure Count

Health is updated

Application Startup

Manual Test

Model Switch

Periodic Background Check

---

# Status Values

Connected

Connecting

Disconnected

Offline

Authentication Failed

Timeout

Unknown

---

# Connection Indicator

Green

Connected

Yellow

Slow

Red

Offline

Gray

Unknown

---

# Latency Display

Display

18ms

135ms

502ms

1.2s

Users can quickly identify slow providers.

---

# Authentication

Default

Authorization

Bearer API_KEY

Future Support

Custom Headers

OAuth

Token Refresh

Session Tokens

---

# OpenAI Compatibility

The provider must support

/models

/chat/completions

Streaming

Tool Calling (optional)

Reasoning Models (optional)

Vision Models (optional)

---

# Provider Priority

Each provider can have

Priority

1

↓

Highest

Useful for

Automatic Failover

---

# Default Provider

Only one provider may be default.

Changing default

Does not automatically change

Current Model

---

# Provider Switching

Command

/provider switch

↓

List Providers

↓

Select

↓

Provider Active

↓

Keep Current Session

No restart required.

---

# Provider Search

Search Provider...

Supports

Name

Base URL

Status

Description

Instant filtering.

---

# Provider Cache

Cache

Connection

Latency

Capabilities

Supported Models

Cache expires after configurable duration.

---

# Error Messages

Invalid URL

Unable to connect

Authentication failed

Timeout

Unsupported API

Malformed response

SSL error

Every error should explain

What happened

Why

How to fix it

---

# Security

API Keys are

Hidden in UI

Masked in logs

Never printed

Never exported by default

Never included in crash reports

Future

Encrypted local storage

---

# Future Features

Automatic Model Discovery

Provider Tags

Provider Groups

Cloud Sync

Usage Statistics

Cost Tracking

Rate Limit Monitoring

Quota Monitoring

OAuth Providers

Multi-Key Rotation

Provider Templates

---

# Provider Lifecycle

Create

↓

Validate

↓

Test

↓

Save

↓

Load

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

Provider management should be so simple that users never need to manually edit JSON files.

Everything should be discoverable through interactive dialogs, searchable lists, and slash commands, making OpenChat CLI a universal frontend for any OpenAI-compatible AI service while remaining secure, transparent, and easy to use.