# Agent Architecture

## Overview

The agent architecture defines how the AI assistant is organized into roles, how specialized agents are created, and how work is delegated between them.

## Multi-Agent System

The system supports multiple agent profiles. Each agent is a distinct persona with its own system prompt, tool set, and behavioral rules.

### Default Agents

Two primary agents are included:

**Architect Agent**
- Focus: system design, architecture decisions, high-level planning
- Strengths: understanding large codebases, making design trade-offs, creating technical specifications
- When to use: designing new features, refactoring architecture, code review
- Tool focus: read/search/glob (more reading, less writing)
- Behavior: explains rationale before acting, asks clarifying questions, considers alternatives

**Developer Agent**
- Focus: implementation, coding, debugging, testing
- Strengths: writing and editing code, fixing bugs, running tests
- When to use: implementing features, fixing issues, writing tests
- Tool focus: read/write/edit/shell (balanced reading and writing)
- Behavior: acts quickly, makes reasonable assumptions, asks when truly stuck

### Agent Profiles

Each agent profile contains:
- **Name** — identifier for the agent
- **Display name** — human-readable label
- **Description** — when and why to use this agent
- **System prompt** — core identity and behavioral instructions
- **Default tools** — which tools are available (subset of all tools)
- **Configuration** — model preference, temperature, verbosity
- **Skills** — preloaded skills relevant to this agent's focus

The user can:
- Switch between agents mid-session
- Create custom agent profiles
- Import/export agent definitions

## Sub-Agent Delegation

An agent can spawn sub-agents for specialized tasks:

### Delegation Model
1. The parent agent identifies a sub-task suitable for delegation
2. The parent creates a sub-agent with a specific goal and context
3. The sub-agent works autonomously on the task
4. The sub-agent returns its results (code, analysis, findings)
5. The parent incorporates the results into the main workflow

### Sub-Agent Lifecycle
- **Spawned** — created with a goal, context, and available tools
- **Working** — actively executing tools and generating responses
- **Reporting** — returning results to the parent
- **Complete** — task finished, sub-agent terminated
- **Failed** — sub-agent encountered an unrecoverable error

### Sub-Agent Visibility
- The user sees when a sub-agent is spawned and what its task is
- Sub-agent activity is visible in the conversation (with indentation or nesting)
- The user can inspect the sub-agent's full interaction
- The user can cancel a sub-agent mid-execution

### Deep Delegation
- Sub-agents can themselves create sub-agents (recursive delegation)
- Maximum depth is configurable (default: 3 levels)
- Each level is visually distinct in the UI
- Deep delegation is logged for audit

## Task Routing

When a new task arrives, it can be:
1. **Handled by the current agent** — the active agent processes it
2. **Routed to a specific agent** — user specifies which agent should handle it
3. **Auto-routed** — the system analyzes the task and routes it to the best-fit agent based on:
   - Task content keywords
   - Required tools
   - Complexity
   - User's past routing patterns

## Agent Communication

Agents communicate through the session context:
- The parent agent's context is available to the sub-agent (read-only)
- The sub-agent's results are returned to the parent's context
- Agents cannot communicate directly (no inter-agent messaging)
- All context flows through the session, making it auditable

## Key Design Decisions

- Two default agents cover the primary use cases (plan + execute)
- Custom agents make the system extensible without code changes
- Sub-agent delegation enables complex multi-step workflows
- Each agent operates in isolation — no shared mutable state
- The user always knows which agent is working and why
- Delegation is visible and auditable
- Auto-routing learns from user behavior but can be overridden
- Agent profiles are portable (import/export)
