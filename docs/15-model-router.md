# 15-model-router.md

# Model Router Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Model Router?

A **Model Router** is the subsystem responsible for selecting the most appropriate Large Language Model (LLM) for each request.

Instead of sending every request to the same model, the Model Router analyzes the task, user preferences, available providers, model capabilities, pricing, latency, context limits, and tool requirements before routing the request.

The router separates **decision-making** from **model execution**.

---

# Why Model Router?

Without Model Router

```
User Request

↓

Single Model

↓

Response
```

Problems

- Higher cost
- Poor model selection
- No failover
- Slow responses
- No provider flexibility

---

With Model Router

```
User Request

↓

Model Router

↓

Best Model

↓

LLM Provider

↓

Response
```

---

# Goals

A production Model Router should provide

- Intelligent model selection
- Multi-provider support
- Automatic failover
- Cost optimization
- Latency optimization
- Capability matching
- Context limit awareness
- Tool compatibility
- Streaming support
- Load balancing

---

# High-Level Architecture

```
            Prompt Builder

                  │

                  ▼

             Model Router

                  │

      ┌───────────┼────────────┐

      ▼           ▼            ▼

 Capability   Cost Engine   Latency

      ▼           ▼            ▼

 Provider   Availability   Policies

      └───────────┼────────────┘

                  ▼

        Selected Model

                  ▼

      LLM Provider Manager
```

---

# Folder Structure

```
src/

router/

    ModelRouter.ts

    ModelRegistry.ts

    ModelSelector.ts

    ProviderResolver.ts

    CapabilityMatcher.ts

    CostEngine.ts

    LatencyTracker.ts

    AvailabilityMonitor.ts

    RoutingPolicies.ts

    LoadBalancer.ts

    FailoverManager.ts

    RouterCache.ts

    RouterMetrics.ts

    RouterEvents.ts
```

---

# Core Components

## Model Router

Central controller.

Responsibilities

- Analyze request
- Select model
- Apply routing rules
- Return routing decision

---

## Model Registry

Stores

```
Available Models

Capabilities

Providers

Pricing

Context Limits

Metadata
```

---

## Capability Matcher

Matches request requirements with model abilities.

Examples

```
Reasoning

Coding

Vision

Function Calling

Long Context

Streaming
```

---

## Cost Engine

Calculates

```
Input Tokens

Output Tokens

Estimated Price
```

Supports budget-aware routing.

---

## Latency Tracker

Measures

- Response time
- Queue time
- Throughput
- Historical averages

---

## Availability Monitor

Tracks

```
Healthy

Busy

Offline

Rate Limited
```

---

## Provider Resolver

Maps

```
Model

↓

Provider
```

Example

```
GPT

↓

OpenAI

-------------

Gemini

↓

Google

-------------

Claude

↓

Anthropic
```

---

## Load Balancer

Distributes traffic across

- Multiple providers
- Multiple regions
- Multiple endpoints

---

## Failover Manager

Automatically switches provider if

- Timeout
- Rate limit
- Network error
- Service unavailable

---

# Routing Lifecycle

```
Receive Prompt

↓

Analyze Requirements

↓

Match Capabilities

↓

Apply Policies

↓

Select Provider

↓

Return Model
```

---

# Routing Factors

The router evaluates

- Task type
- Required tools
- Context length
- User preference
- Model health
- Cost
- Latency
- Permissions
- Streaming support

---

# Request Object

Contains

```
Prompt

Task Type

Workspace

Tools

Session

Preferences

Metadata
```

---

# Routing Decision

Returns

```
Provider

Model

Endpoint

Streaming

Fallback

Metadata
```

---

# Routing Policies

Examples

```
Coding

↓

Best Coding Model

-------------

Large Context

↓

Long Context Model

-------------

Low Budget

↓

Cheapest Model

-------------

Fast Response

↓

Lowest Latency Model
```

---

# Cost Optimization

Example

```
Simple Question

↓

Small Model

-------------

Complex Refactoring

↓

Advanced Model
```

Avoid using expensive models unnecessarily.

---

# Context Limit Awareness

Before routing

Check

```
Prompt Size

↓

Fits Model?

↓

Yes

↓

Route

-------------

No

↓

Choose Larger Context Model
```

---

# Streaming Support

Select models supporting

```
Real-time Token Streaming
```

when requested.

---

# Tool Compatibility

Some requests require

```
Function Calling

MCP

Tool Use

JSON Output
```

Only compatible models are selected.

---

# Multi-Provider Support

Example

```
OpenAI

Anthropic

Google

OpenRouter

Local Models
```

The router treats them through a unified interface.

---

# Failover Flow

```
Primary Provider

↓

Failure

↓

Failover Manager

↓

Secondary Provider

↓

Continue
```

---

# Load Balancing

Strategies

- Round Robin
- Least Latency
- Least Cost
- Weighted Routing
- Health-based Routing

---

# Event Bus Integration

Common events

```
router:start

router:select

router:fallback

router:error

router:complete
```

---

# Prompt Builder Integration

```
Prompt Builder

↓

Model Router
```

The router never builds prompts.

---

# Provider Manager Integration

```
Model Router

↓

LLM Provider Manager
```

The router selects, the provider manager executes.

---

# Session Integration

Session stores

```
Preferred Model

Preferred Provider

Recent Usage
```

---

# Plugin Integration

Plugins may register

- New providers
- Custom routing rules
- Capability extensions

---

# Skills Integration

Skills may request

```
Vision Model

Reasoning Model

Coding Model
```

The router resolves the appropriate choice.

---

# Cache Strategy

Cache routing decisions using

```
Task Type

+

Model Requirements

+

Workspace
```

to reduce repeated computation.

---

# Error Handling

```
No Matching Model

↓

Fallback Policy

↓

Default Model

↓

Continue
```

---

# Security

Always

- Validate provider endpoints
- Verify model availability
- Protect API credentials
- Respect routing policies

Never

- Route to unauthorized providers
- Ignore health checks
- Expose internal endpoints

---

# Performance Optimizations

Use

- Routing cache
- Health monitoring
- Latency tracking
- Capability indexing
- Parallel provider checks

Avoid

- Querying every provider every request
- Recalculating identical routes
- Blocking routing decisions

---

# Best Practices

Always

- Separate routing from execution
- Track provider health
- Optimize for task requirements
- Support failover
- Log routing decisions

Never

- Hardcode provider choices
- Ignore context limits
- Skip capability validation

---

# Common Mistakes

Bad

```
Every Request

↓

Same Model
```

Expensive and inefficient.

---

Good

```
Analyze Request

↓

Capability Match

↓

Cost Check

↓

Health Check

↓

Best Model
```

Intelligent routing.

---

# Testing Checklist

- Model selection
- Capability matching
- Cost optimization
- Context limits
- Provider failover
- Load balancing
- Streaming support
- Health monitoring
- Cache
- Error recovery

---

# Advantages

- Lower costs
- Faster responses
- Better reliability
- Multi-provider flexibility
- Intelligent routing
- Improved scalability

---

# Disadvantages

- Routing complexity
- Health monitoring overhead
- Provider compatibility maintenance
- Policy management

---

# Used In

- OpenCode
- Antigravity CLI
- Cursor
- Claude Code
- Continue.dev
- GitHub Copilot Agents
- Enterprise AI Platforms

---

# Complete Routing Flow

```
Prompt Builder

↓

Model Router

↓

Capability Analysis

↓

Policy Evaluation

↓

Cost Engine

↓

Latency Check

↓

Health Check

↓

Provider Selection

↓

LLM Provider Manager

↓

Model

↓

Response
```

---

# Summary

The **Model Router** is the intelligent decision-making layer that determines which language model should handle a given request.

A production-grade Model Router should include:

- Model Router
- Model Registry
- Capability Matcher
- Cost Engine
- Latency Tracker
- Availability Monitor
- Provider Resolver
- Load Balancer
- Failover Manager
- Routing Policies
- Cache
- Event Bus Integration

By separating routing decisions from model execution, the Model Router enables an AI Coding Agent to intelligently balance capability, cost, latency, reliability, and scalability across multiple providers, making it a core architectural component of modern systems such as OpenCode, Antigravity CLI, and other enterprise AI development platforms.