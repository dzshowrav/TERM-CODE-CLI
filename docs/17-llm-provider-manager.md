# 17-llm-provider-manager.md

# LLM Provider Manager Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is an LLM Provider Manager?

An **LLM Provider Manager** is the subsystem responsible for communicating with external and local Large Language Model (LLM) providers.

The **Model Router** decides **which model** should be used.

The **LLM Provider Manager** is responsible for **executing the request**, handling authentication, retries, streaming, provider-specific APIs, rate limits, and normalizing responses into a common format.

It provides a single interface regardless of which provider is being used.

---

# Why LLM Provider Manager?

Without Provider Manager

```
Agent

↓

OpenAI API

Anthropic API

Gemini API

OpenRouter API

Local Model

↓

Response
```

Problems

- Different APIs
- Duplicate code
- Hard maintenance
- Provider lock-in

---

With Provider Manager

```
Agent

↓

Provider Manager

↓

Provider Adapter

↓

Provider API

↓

Normalized Response
```

---

# Goals

A production LLM Provider Manager should provide

- Unified provider interface
- Multi-provider support
- Authentication
- Request normalization
- Response normalization
- Streaming support
- Retry handling
- Rate limit handling
- Timeout handling
- Usage tracking

---

# High-Level Architecture

```
             Model Router

                  │

                  ▼

      LLM Provider Manager

                  │

      ┌───────────┼────────────┐

      ▼           ▼            ▼

 Provider    Authentication   Cache

      ▼           ▼            ▼

 Adapters    Retry Engine   Metrics

      └───────────┼────────────┘

                  ▼

        External Providers

      OpenAI

      Anthropic

      Gemini

      OpenRouter

      Ollama

      LM Studio

      Local Models
```

---

# Folder Structure

```
src/

providers/

    ProviderManager.ts

    ProviderRegistry.ts

    ProviderAdapter.ts

    ProviderFactory.ts

    ProviderClient.ts

    Authentication.ts

    RequestBuilder.ts

    ResponseParser.ts

    StreamingManager.ts

    RetryManager.ts

    TimeoutManager.ts

    RateLimiter.ts

    ProviderCache.ts

    ProviderMetrics.ts

    ProviderEvents.ts
```

---

# Core Components

## Provider Manager

Central controller.

Responsibilities

- Execute requests
- Select adapter
- Authenticate
- Normalize responses

---

## Provider Registry

Stores

```
Available Providers

Supported Models

Capabilities

Endpoints

Metadata
```

---

## Provider Adapter

Each provider implements the same interface.

Examples

```
OpenAI Adapter

Anthropic Adapter

Gemini Adapter

OpenRouter Adapter

Ollama Adapter
```

---

## Provider Factory

Creates

```
Provider

↓

Adapter Instance
```

---

## Authentication Manager

Handles

```
API Keys

OAuth

Bearer Tokens

Environment Variables
```

---

## Request Builder

Transforms

```
Unified Request

↓

Provider Request
```

Each provider has different request formats.

---

## Response Parser

Transforms

```
Provider Response

↓

Unified Response
```

The Agent always receives the same structure.

---

## Streaming Manager

Supports

```
Streaming Tokens

↓

Partial Responses

↓

Final Response
```

---

## Retry Manager

Handles

```
Failure

↓

Retry

↓

Success
```

or

```
Failure

↓

Fallback

↓

Another Provider
```

---

## Timeout Manager

Controls

```
Start

↓

Maximum Time

↓

Cancel

↓

Error
```

---

## Rate Limiter

Tracks

- Requests
- Tokens
- Quotas
- Cooldowns

---

## Provider Cache

Caches

```
Metadata

Capabilities

Model Lists

Health Status
```

---

# Request Lifecycle

```
Receive Request

↓

Authenticate

↓

Build Provider Request

↓

Execute

↓

Receive Response

↓

Normalize

↓

Return
```

---

# Unified Request Object

Contains

```
Model

Messages

Tools

Streaming

Temperature

Max Tokens

Metadata
```

---

# Unified Response Object

Contains

```
Content

Tool Calls

Usage

Finish Reason

Metadata

Provider Info
```

---

# Provider Adapters

Examples

```
OpenAI

Anthropic

Gemini

OpenRouter

Ollama

LM Studio

vLLM

Custom APIs
```

Each adapter hides provider-specific implementation details.

---

# Streaming Flow

```
Provider

↓

Streaming Manager

↓

Conversation Manager

↓

Renderer
```

---

# Authentication Flow

```
API Key

↓

Authentication Manager

↓

Provider Request
```

Supports secure credential handling.

---

# Retry Strategy

Example

```
Timeout

↓

Retry

↓

Success
```

or

```
Rate Limit

↓

Wait

↓

Retry
```

---

# Provider Failover

```
Primary Provider

↓

Unavailable

↓

Secondary Provider

↓

Continue
```

Works together with the Model Router.

---

# Response Normalization

Different providers return different formats.

Normalize into

```
Unified Response

↓

Agent
```

The Agent never parses provider-specific data.

---

# Tool Calling Support

Support

```
Function Calling

Tool Use

Structured Output

JSON Mode
```

Only when supported by the provider.

---

# Event Bus Integration

Common events

```
provider:start

provider:request

provider:stream

provider:complete

provider:error

provider:retry
```

---

# Model Router Integration

```
Model Router

↓

Provider Manager

↓

Selected Provider
```

---

# Conversation Manager Integration

```
Streaming Tokens

↓

Conversation Update

↓

Renderer
```

---

# Session Integration

Store

```
Last Provider

Usage

Preferred Endpoint
```

---

# Plugin Integration

Plugins may register

- New providers
- Custom adapters
- Authentication methods
- Response parsers

---

# Skills Integration

Skills may require

```
Vision Support

Tool Calling

Large Context

Streaming
```

The Provider Manager executes accordingly.

---

# Cache Strategy

Cache

```
Provider Metadata

Supported Models

Health Status
```

Refresh periodically.

---

# Error Handling

```
Network Error

↓

Retry

↓

Fallback

↓

Report
```

Never expose raw provider errors directly to the user.

---

# Security

Always

- Encrypt API keys
- Validate endpoints
- Verify SSL/TLS
- Mask credentials in logs
- Limit retries

Never

- Store API keys in plain text
- Trust unknown endpoints
- Leak provider credentials

---

# Performance Optimizations

Use

- Connection pooling
- Streaming
- Provider cache
- HTTP keep-alive
- Parallel health checks
- Request reuse

Avoid

- Creating new connections for every request
- Duplicate authentication
- Blocking streaming responses

---

# Best Practices

Always

- Normalize requests
- Normalize responses
- Separate adapters
- Support retries
- Track usage
- Handle provider failures gracefully

Never

- Couple the Agent to provider APIs
- Hardcode provider logic
- Ignore rate limits
- Skip authentication validation

---

# Common Mistakes

Bad

```
Agent

↓

OpenAI API

↓

Agent

↓

Gemini API

↓

Agent
```

Multiple incompatible implementations.

---

Good

```
Agent

↓

Provider Manager

↓

Adapter

↓

Provider

↓

Normalized Response
```

Clean abstraction.

---

# Testing Checklist

- Authentication
- Request building
- Response parsing
- Streaming
- Retry
- Timeout
- Rate limiting
- Failover
- Cache
- Usage tracking
- Error handling

---

# Advantages

- Unified provider interface
- Easy provider switching
- Better maintainability
- Improved reliability
- Streaming support
- Multi-provider architecture

---

# Disadvantages

- Adapter maintenance
- Provider API changes
- Authentication complexity
- Additional abstraction layer

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

# Complete Provider Flow

```
Prompt Builder

↓

Model Router

↓

LLM Provider Manager

↓

Authentication

↓

Provider Adapter

↓

Request Builder

↓

Provider API

↓

Streaming Response

↓

Response Parser

↓

Unified Response

↓

Conversation Manager

↓

Renderer

↓

User
```

---

# Summary

The **LLM Provider Manager** is the execution layer responsible for communicating with language model providers through a unified interface.

A production-grade LLM Provider Manager should include:

- Provider Manager
- Provider Registry
- Provider Factory
- Provider Adapters
- Authentication Manager
- Request Builder
- Response Parser
- Streaming Manager
- Retry Manager
- Timeout Manager
- Rate Limiter
- Provider Cache
- Event Bus Integration

By abstracting provider-specific APIs behind standardized adapters, the LLM Provider Manager enables seamless integration with multiple cloud and local LLM providers while improving reliability, scalability, maintainability, and interoperability across modern AI Coding Agents such as OpenCode, Antigravity CLI, and enterprise AI platforms.