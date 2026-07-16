# OpenChat CLI

# Action Engine Specification

Version: 1.0

---

# Overview

The Action Engine is the autonomous execution layer of OpenChat CLI.

Unlike a traditional chatbot that simply responds to prompts, the Action Engine understands goals, creates execution plans, coordinates tools, tracks progress, requests approval when necessary, and verifies results before completing a task.

The Action Engine is responsible for turning natural language into structured, executable workflows.

---

# Philosophy

User Goal

↓

Analyze Intent

↓

Create Plan

↓

Request Approval (if needed)

↓

Execute Actions

↓

Verify Results

↓

Complete Task

The AI should behave like a senior software engineer solving a problem methodically.

---

# Design Goals

✓ Goal Oriented

✓ Intelligent Planning

✓ Transparent

✓ Safe

✓ Reversible

✓ Tool Aware

✓ Context Aware

✓ Workspace Aware

✓ Provider Independent

✓ Mobile Friendly

---

# Architecture

User

↓

Chat Engine

↓

Action Engine

↓

Planner

↓

Permission Manager

↓

Tool Manager

↓

Workspace

↓

Result

↓

AI Response

---

# Action Lifecycle

Receive Goal

↓

Analyze Intent

↓

Create Execution Plan

↓

Estimate Complexity

↓

Ask Permission

↓

Execute Steps

↓

Verify Output

↓

Generate Summary

↓

Save Session

---

# Action Types

Read

Analyze

Search

Modify

Generate

Review

Debug

Refactor

Test

Build

Deploy (Future)

---

# Planning Levels

Simple

Single Step

Example

Read README.md

Medium

Multiple Steps

Example

Fix authentication bug

Complex

Long-running Workflow

Example

Build complete Laravel authentication system

---

# Task Planner

Before executing anything

The planner creates

Goal

Steps

Dependencies

Required Tools

Risk Level

Estimated Duration

Expected Output

---

# Execution Plan Example

Goal

Add Laravel Authentication

Plan

1.

Read composer.json

2.

Detect Laravel Version

3.

Analyze Routes

4.

Create Controller

5.

Create Middleware

6.

Update Routes

7.

Run Tests

8.

Generate Summary

---

# Action Queue

Pending

Running

Waiting Approval

Completed

Failed

Cancelled

---

# Parallel Execution

Allowed

Read Files

Search

Workspace Scan

Dependency Analysis

Git Status

Parallel tasks improve performance.

---

# Sequential Execution

Required

Edit File

Delete File

Shell Commands

Git Commit

Database Writes

Deployment

---

# Approval System

Before dangerous actions

Display

Execution Plan

↓

Requested Changes

↓

Risk Level

↓

Approve

Cancel

Modify Plan

---

# Risk Levels

Low

Read Only

Medium

File Modifications

High

Delete

Shell

Git Push

Database Write

Critical

System Operations

Require explicit approval.

---

# TODO Manager

Every plan automatically generates

Task List

Example

☐ Read Project

☐ Analyze Dependencies

☐ Update Routes

☐ Generate Controller

☐ Run Tests

☐ Verify

☐ Complete

---

# Progress Tracking

Display

0%

Planning

↓

25%

Analyzing

↓

50%

Editing

↓

75%

Testing

↓

100%

Completed

---

# Checkpoints

Long-running tasks automatically create

Restore Points

Checkpoint 1

Project Analysis

Checkpoint 2

File Changes

Checkpoint 3

Testing

Allows recovery after interruption.

---

# Rollback

If execution fails

Restore Previous State

Undo File Changes

Restore Session

Report Errors

Rollback should be available whenever possible.

---

# Verification

Every completed action should verify

Files Created

Files Modified

Tests Passed

Errors Fixed

Dependencies Installed

Verification happens before success is reported.

---

# Action Summary

At completion display

Goal

Completed Steps

Modified Files

Tool Usage

Execution Time

Warnings

Next Recommendations

---

# Tool Coordination

The Action Engine selects

Filesystem

Git

Shell

Workspace

Search

Database

MCP

Without exposing implementation details to the user.

---

# Context Awareness

Planner considers

Conversation

Workspace

Agent

Skills

Model Capabilities

Previous Actions

---

# Interruption

Ctrl+C

↓

Pause Action

↓

Save State

↓

Resume Later

---

# Resume

Command

/action resume

Restores

Plan

Progress

Context

Tool State

---

# Cancel

Command

/action cancel

Stops

Current Execution

Keeps

Conversation

Workspace

History

---

# Multiple Actions

Queue

Action A

↓

Action B

↓

Action C

Users can reorder pending actions.

---

# Action History

Stores

Goal

Plan

Execution Time

Success

Failure

Warnings

Tool Calls

---

# Action Search

Search

Goal

Workspace

Date

Status

Files

Instant filtering.

---

# AI Decision Rules

The planner should

Prefer Read before Write

Verify before Modify

Ask before Dangerous Actions

Reuse Existing Context

Avoid Duplicate Work

---

# Failure Recovery

If a step fails

Retry

↓

Alternative Strategy

↓

Request User Input

↓

Cancel Gracefully

Never silently ignore failures.

---

# Performance Goals

Plan Generation

<100ms

Action Queue Update

Instant

Progress Refresh

Real Time

Verification

Automatic

Minimal overhead.

---

# Future Features

Autonomous Project Mode

Scheduled Actions

Background Tasks

Multi-Agent Planning

Collaborative Execution

Workflow Templates

Custom Action Pipelines

Cloud Execution

CI/CD Integration

---

# Design Principles

The Action Engine is the decision-making brain of OpenChat CLI.

Instead of reacting to prompts one message at a time, it transforms user goals into structured execution plans, coordinates tools intelligently, verifies results, and keeps the user informed throughout the process.

By combining planning, transparency, permissions, and verification, the Action Engine enables OpenChat CLI to function as a true AI coding agent rather than a simple conversational assistant.