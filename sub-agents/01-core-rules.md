# 01-core-rules.md

# TermCode Core Rules

Version: 1.0.0

---

# Purpose

This document defines the global rules that every AI agent inside the TermCode ecosystem must follow.

These rules have the highest priority after the Master Architect.

Every specialized agent, planner, reviewer, reasoning engine, memory engine, workflow engine, and tool executor must obey these rules without exception.

---

# Rule Priority

Priority Order

```
1. Master Architect
2. Core Rules
3. Agent Protocol
4. Task Planner
5. Specialized Agents
6. User Task
```

Lower-priority rules must never override higher-priority rules.

---

# Rule 01

Architecture First

Never sacrifice architecture for speed.

Always preserve long-term maintainability.

---

# Rule 02

Understand Before Acting

Never generate code immediately.

Always:

- Read context
- Understand the request
- Identify affected modules
- Analyze dependencies
- Determine execution strategy

Only then begin implementation.

---

# Rule 03

No Assumptions

Never guess:

- APIs
- Folder names
- File names
- Data structures
- Database schema
- Existing logic

If required information is unavailable:

Stop.

Request clarification.

---

# Rule 04

Production Only

Never generate:

- Demo code
- Example code
- Placeholder logic
- Fake implementation
- Mock services
- Temporary fixes

Every implementation must be production-ready.

---

# Rule 05

Complete Solutions

Never return incomplete work.

Always provide:

- Complete implementation
- Required imports
- Error handling
- Validation
- Documentation updates when necessary

---

# Rule 06

Consistency

Maintain consistency across:

- Naming
- Folder structure
- Coding style
- UI
- Architecture
- Error handling
- Logging
- State management

---

# Rule 07

Mobile First

Every feature must be designed primarily for mobile devices.

Always consider:

- Small screen
- Portrait mode
- Touch interaction
- Virtual keyboard
- Limited resources

Desktop is secondary.

---

# Rule 08

Termux First

Every solution must run correctly inside Android Termux whenever possible.

Avoid dependencies that require:

- Desktop GUI
- Docker
- Platform-specific desktop APIs

Prefer lightweight tools.

---

# Rule 09

Offline First

Do not assume internet access.

The application must function locally whenever possible.

Remote services must remain optional.

---

# Rule 10

Go Best Practices

Always follow idiomatic Go.

Prefer:

- Small packages
- Small functions
- Explicit error handling
- Composition
- Interfaces where appropriate

Avoid unnecessary abstraction.

---

# Rule 11

Bubble Tea Standards

Every Bubble Tea component must:

- Be reusable
- Keep state isolated
- Minimize update complexity
- Render efficiently

Never place business logic inside rendering code.

---

# Rule 12

Single Responsibility

Every:

Function

Struct

Package

Agent

Module

Must have one clear responsibility.

---

# Rule 13

No Duplication

Before writing code:

Search existing implementation.

Reuse whenever possible.

Avoid duplicated:

Logic

Components

Utilities

Commands

---

# Rule 14

Standard Library First

Always prefer:

Go Standard Library

Before introducing external packages.

---

# Rule 15

Minimal Dependencies

Only install a dependency if:

It solves a real problem.

It is actively maintained.

It has stable releases.

It improves the project.

---

# Rule 16

Readable Code

Code must be understandable.

Prefer clarity over cleverness.

---

# Rule 17

Explicit Error Handling

Every possible error must be handled.

Never silently ignore errors.

Never use empty error blocks.

---

# Rule 18

Validation Required

Always validate:

Input

Configuration

Environment

File paths

User commands

Database queries

External responses

---

# Rule 19

Secure by Default

Never expose:

Secrets

API keys

Passwords

Private tokens

Environment variables

Database credentials

---

# Rule 20

Performance Awareness

Minimize:

Memory usage

CPU usage

Disk I/O

Network calls

Rendering cost

Context size

---

# Rule 21

State Management

State must be:

Predictable

Minimal

Explicit

Recoverable

Never create hidden state.

---

# Rule 22

Documentation

Major architectural changes require documentation updates.

Documentation is part of implementation.

---

# Rule 23

Logging

Logs must be:

Useful

Structured

Readable

Actionable

Never log secrets.

---

# Rule 24

Testing Mindset

Every implementation must be testable.

Avoid tightly coupled code.

---

# Rule 25

Clean Imports

Remove:

Unused imports

Unused variables

Dead code

Unused constants

Unused functions

---

# Rule 26

Deterministic Behavior

The same input should produce the same output whenever possible.

Avoid unpredictable behavior.

---

# Rule 27

Naming

Names must clearly describe purpose.

Avoid abbreviations unless universally accepted.

Good:

ProjectManager

Bad:

PM

---

# Rule 28

Small Functions

Prefer:

20–40 line functions.

Split large functions into logical units.

---

# Rule 29

File Organization

Keep related functionality together.

Avoid oversized files.

Prefer modular packages.

---

# Rule 30

No Circular Dependencies

Packages must remain independent.

Dependency direction must always remain clear.

---

# Rule 31

Review Before Finish

Every implementation must pass:

Architecture review

Logic review

Performance review

Security review

Consistency review

---

# Rule 32

Graceful Failure

If an operation fails:

Return meaningful errors.

Preserve application stability.

---

# Rule 33

Backward Compatibility

Avoid breaking existing behavior unless explicitly required.

---

# Rule 34

User Confirmation

Require confirmation before:

Deleting files

Overwriting files

Database destruction

Reset operations

Mass modifications

---

# Rule 35

Context Awareness

Always understand:

Current workspace

Current module

Current file

Current task

Current session

Before making changes.

---

# Rule 36

Tool Usage

Use tools only when necessary.

Avoid unnecessary tool execution.

---

# Rule 37

MCP Usage

Use the smallest number of MCP servers required to solve the task.

Avoid unnecessary remote operations.

---

# Rule 38

Self Review

Before completing work, verify:

Architecture

Code quality

Naming

Formatting

Performance

Security

Consistency

---

# Rule 39

Long-Term Thinking

Every decision should improve the project five years from now, not just today.

---

# Rule 40

Mission

Every action performed inside TermCode must contribute to building a reliable, maintainable, scalable, secure, mobile-first, Termux-compatible, production-ready local AI Coding CLI.