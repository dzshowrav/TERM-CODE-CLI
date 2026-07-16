# AI Model & Provider System

## Overview

The AI model & provider system is the only component translated directly from the current project. It handles the registration, configuration, authentication, and routing to various AI model providers.

## Provider

A provider is a service that hosts AI models and exposes them via an API endpoint. Each provider has:

### Definition
- **Name** — human-readable label (e.g., "OpenAI", "Anthropic", "Google AI", "local Ollama")
- **Base URL** — the API endpoint base path
- **Authentication** — API key, token, or other credential storage
- **Models** — list of models available through this provider
- **Capabilities** — what this provider supports (streaming, tool calls, vision, thinking, etc.)
- **Rate limits** — requests-per-minute, tokens-per-minute limits
- **Custom headers** — additional headers to include in every request
- **Priority** — when multiple providers support the same model, priority determines selection

### Provider CRUD
- **Add** — register a new provider with URL, auth, and model list
- **Edit** — modify any field of an existing provider
- **Delete** — remove a provider entirely (warns if sessions still reference it)
- **Test** — verify connectivity and authentication with a lightweight request
- **List** — show all registered providers with their status

### Authentication Storage
- Credentials are stored securely (obfuscated or encrypted in storage)
- The user can set credentials via environment variables or interactive prompt
- Credentials are never displayed in logs or session exports

## Model

A model is a specific AI model hosted by a provider. Each model has:

### Definition
- **Name** — model identifier (e.g., "gpt-4o", "claude-sonnet-4-20250514", "qwen3-235b-a22b")
- **Provider reference** — which provider hosts this model
- **Context window** — maximum input context size in tokens
- **Max output tokens** — maximum generation length
- **Pricing** — per-input-token and per-output-token cost (for usage tracking)
- **Capabilities** — what this model can do:
  - Streaming (token-by-token output)
  - Tool/function calling
  - Vision (image input)
  - Thinking/reasoning (extended chain-of-thought)
  - Structured output (JSON mode)
  - System prompt support
- **Version** — model version string

### Model Organization
Models are organized by provider. A provider can offer multiple models. Models can be:
- **Featured** — marked as recommended
- **Deprecated** — still available but marked for removal
- **Custom** — user-defined model pointing to a custom endpoint

### Default Provider & Model
- A default provider and model are configured for new sessions
- The default can be overridden per-session or per-message
- If the default model is unavailable, the system falls back to the next available model

## Model Selection

When sending a request, the model is selected by:
1. Message-level override (if specified)
2. Session-level default (if configured)
3. Global default

If the selected model is unavailable (rate-limited, down, removed), the system:
1. Tries the same model on a different provider (if available)
2. Tries the next available model from the same provider
3. Reports the failure and asks the user what to do

## Usage Tracking

The system tracks per-session and cumulative:
- Token counts (input + output)
- Cost (based on model pricing)
- Request count
- Response time

Usage data is visible in session details and can be reset.

## Key Design Decisions

- Provider configuration is entirely user-controlled — no hardcoded providers
- Multiple providers for the same model enable automatic failover
- All provider communication goes through a unified request interface
- Authentication failures are reported immediately with clear guidance
- Usage tracking is optional and opt-in
