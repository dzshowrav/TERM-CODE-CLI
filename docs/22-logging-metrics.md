# 22-logging-metrics.md

# Logging & Metrics Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What are Logging & Metrics?

**Logging** records everything that happens inside the AI Coding Agent.

**Metrics** measure how well the system performs.

Together they provide observability, debugging, monitoring, auditing, analytics, and performance optimization.

A production AI Agent should never operate without Logging and Metrics.

---

# Why Logging & Metrics?

Without Logging

```
Agent

↓

Error

↓

Unknown Cause
```

Problems

- Difficult debugging
- No audit trail
- Hidden failures
- Poor diagnostics

---

Without Metrics

```
Agent

↓

Runs

↓

No Performance Data
```

Problems

- Unknown latency
- Unknown costs
- Unknown bottlenecks
- No optimization

---

With Logging & Metrics

```
System Event

↓

Logger

↓

Metrics

↓

Storage

↓

Dashboard

↓

Developer
```

---

# Goals

A production Logging & Metrics system should provide

- Structured logging
- Performance metrics
- Usage analytics
- Error tracking
- Audit logs
- Token monitoring
- Cost tracking
- Latency measurement
- Health monitoring
- Event correlation

---

# High-Level Architecture

```
             Application

                  │

                  ▼

        Logging & Metrics

                  │

      ┌───────────┼────────────┐

      ▼           ▼            ▼

 Logger      Metrics      Events

      ▼           ▼            ▼

 Storage     Dashboard    Alerts

      └───────────┼────────────┘

                  ▼

            Monitoring
```

---

# Folder Structure

```
src/

logging/

    Logger.ts

    LogFormatter.ts

    LogWriter.ts

    LogStorage.ts

    Metrics.ts

    MetricsCollector.ts

    MetricsAggregator.ts

    MetricsExporter.ts

    AuditLogger.ts

    HealthMonitor.ts

    Alerts.ts

    LoggingEvents.ts
```

---

# Core Components

## Logger

Central logging controller.

Responsibilities

- Record events
- Format logs
- Write logs
- Export logs

---

## Log Formatter

Formats logs as

```
JSON

Text

Structured Output
```

---

## Log Writer

Writes logs to

- File
- Console
- Database
- Remote service

---

## Metrics Collector

Collects

- Performance
- Usage
- Resource consumption
- Errors

---

## Metrics Aggregator

Aggregates

```
Raw Metrics

↓

Statistics
```

---

## Metrics Exporter

Exports metrics to

- Dashboard
- Monitoring systems
- APIs

---

## Audit Logger

Records

```
Security Events

Permissions

Approvals

Administrative Actions
```

---

## Health Monitor

Tracks

```
CPU

Memory

Providers

Tools

Plugins

Workspace
```

---

## Alert Manager

Generates alerts when

- Errors exceed threshold
- Latency increases
- Provider fails
- Memory usage spikes

---

# Logging Lifecycle

```
Application Event

↓

Logger

↓

Formatter

↓

Writer

↓

Storage
```

---

# Metrics Lifecycle

```
Operation

↓

Collector

↓

Aggregator

↓

Storage

↓

Dashboard
```

---

# Log Entry

Contains

```
Timestamp

Level

Component

Message

Metadata

Session

Workspace
```

---

# Log Levels

```
TRACE

DEBUG

INFO

WARNING

ERROR

FATAL
```

---

# Metrics Types

Supported

```
Performance

Usage

Resource

Business

Security
```

---

# Performance Metrics

Track

```
Latency

Execution Time

Response Time

Queue Time

Streaming Speed
```

---

# Token Metrics

Track

```
Input Tokens

Output Tokens

Cached Tokens

Total Tokens
```

---

# Cost Metrics

Calculate

```
Provider

Model

Token Cost

Total Cost
```

Useful for optimization.

---

# Tool Metrics

Measure

```
Execution Time

Success Rate

Failure Rate

Retries
```

---

# Search Metrics

Track

```
Queries

Search Time

Cache Hits

Ranking Time
```

---

# Workspace Metrics

Measure

```
Indexed Files

Search Requests

Workspace Size

Index Time
```

---

# Streaming Metrics

Track

```
Tokens Per Second

Chunks

Completion Time

Interruptions
```

---

# Provider Metrics

Monitor

```
Latency

Availability

Failures

Rate Limits
```

---

# Event Bus Integration

Common events

```
log:write

metrics:update

health:check

alert:trigger

audit:event
```

---

# Conversation Manager Integration

Log

```
Messages

Streaming

Summaries

Branches
```

---

# Tool Manager Integration

Record

```
Tool Calls

Arguments

Results

Errors
```

---

# Permission Engine Integration

Audit

```
Approvals

Denials

Policy Violations
```

---

# Search Engine Integration

Track

```
Search Queries

Ranking

Cache Usage
```

---

# Model Router Integration

Measure

```
Routing Decisions

Latency

Cost

Fallbacks
```

---

# LLM Provider Integration

Record

```
Provider

Model

Usage

Streaming

Failures
```

---

# Plugin Integration

Plugins may

- Export metrics
- Add log sources
- Register dashboards
- Generate reports

---

# Session Integration

Store

```
Session Metrics

Session Logs

Usage Summary
```

---

# Cache Strategy

Cache

```
Recent Logs

Recent Metrics

Health Status
```

Flush periodically.

---

# Error Handling

```
Logger Failure

↓

Fallback Logger

↓

Continue
```

Logging should never crash the application.

---

# Security

Always

- Mask API keys
- Mask passwords
- Encrypt sensitive logs
- Restrict audit access
- Protect log storage

Never

- Store secrets
- Log private credentials
- Expose internal tokens
- Ignore audit integrity

---

# Performance Optimizations

Use

- Asynchronous logging
- Buffered writes
- Metrics batching
- Incremental aggregation
- Compression
- Background exports

Avoid

- Blocking log writes
- Duplicate metrics
- Excessive disk writes

---

# Best Practices

Always

- Use structured logs
- Separate logs from metrics
- Monitor performance
- Track costs
- Audit security events
- Generate health reports

Never

- Print sensitive information
- Ignore warnings
- Mix audit logs with debug logs
- Lose failure information

---

# Common Mistakes

Bad

```
console.log()

↓

Done
```

No structure or analytics.

---

Good

```
Application

↓

Structured Logger

↓

Metrics

↓

Dashboard
```

Reliable and observable.

---

# Testing Checklist

- Structured logging
- Log levels
- Metrics collection
- Performance tracking
- Token tracking
- Cost tracking
- Health monitoring
- Alerts
- Audit logging
- Export
- Error recovery

---

# Advantages

- Easier debugging
- Better observability
- Performance optimization
- Security auditing
- Cost analysis
- Operational monitoring

---

# Disadvantages

- Storage requirements
- Metric aggregation complexity
- Logging overhead
- Dashboard maintenance

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

# Complete Logging Flow

```
Application Event

↓

Logger

↓

Formatter

↓

Storage

↓

Metrics Collector

↓

Aggregator

↓

Dashboard

↓

Alerts

↓

Developer
```

---

# Summary

The **Logging & Metrics** subsystem provides complete observability for an AI Coding Agent by recording events, measuring performance, tracking resource usage, and generating actionable insights.

A production-grade Logging & Metrics system should include:

- Logger
- Log Formatter
- Log Writer
- Log Storage
- Metrics Collector
- Metrics Aggregator
- Metrics Exporter
- Audit Logger
- Health Monitor
- Alert Manager
- Event Bus Integration

By combining structured logging, performance metrics, security auditing, health monitoring, and operational analytics, the Logging & Metrics subsystem enables AI Coding Agents such as OpenCode, Antigravity CLI, Cursor, Claude Code, and enterprise AI platforms to operate reliably, securely, efficiently, and at scale.