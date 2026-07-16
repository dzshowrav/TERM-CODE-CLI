# Context Management

## Overview

Context management controls what information is sent to the AI model with each request. The goal is to fit the most relevant information within the model's context window while minimizing token usage.

## Context Composition

Every request to the AI model includes:

1. **System prompt** — the agent's identity, capabilities, tool definitions, and behavioral instructions
2. **Workspace context** — project type, detected frameworks, git state, file structure
3. **Skill instructions** — specialized knowledge injected based on the current task
4. **Conversation history** — previous messages (user + assistant)
5. **Tool results** — outputs from recently executed tools
6. **Current message** — the user's latest input

## Context Window Awareness

The system tracks:
- Model's maximum context window (from model definition)
- Current token usage of the assembled context
- Remaining budget for the response

When the context is approaching the limit, the system takes action before sending.

## Compaction Strategies

When the assembled context exceeds the model's limit (or a configurable threshold), the system applies compaction:

### Strategy 1: Drop Oldest Messages
- Remove the oldest user/assistant message pairs first
- Preserve the system prompt, workspace context, and recent messages
- Simple, fast, but loses information

### Strategy 2: Summarize Old History
- Compress old conversation segments into short summaries
- Preserves key decisions and outcomes without full verbatim text
- More expensive (requires an AI call to summarize), but preserves more meaning

### Strategy 3: Selective Retention
- Keep messages that reference files the user is currently working on
- Keep messages containing tool results the model might reuse
- Drop messages about completed, unrelated tasks
- Priority scoring based on recency, relevance, and content type

### Strategy 4: Tool Result Truncation
- Long tool outputs are summarized or truncated (head/tail)
- Full results are stored in the session but only the summary is in context
- The model can request the full output if needed

## Compaction Triggers

Compaction is triggered automatically when:
- Context exceeds 80% of the model's limit (threshold is configurable)
- Context would exceed the limit when adding the current message
- Explicit user request (via command)

## Compaction Visibility

The user is always informed when compaction occurs:
- Which strategy was applied
- How many tokens were saved
- What information was dropped or summarized

The user can also:
- Manually trigger compaction
- Undo compaction (restore full history from session storage)
- Configure which strategy to prefer

## System Prompt Design

The system prompt is composed of layered sections:
1. **Core identity** — who the agent is, its purpose, its tone
2. **Capability declaration** — what tools the agent has, how to use them
3. **Behavioral rules** — how the agent should act (ask before destructive actions, prefer safe approaches, etc.)
4. **Workspace context** — dynamically injected project information
5. **Skill injections** — specialized knowledge for the current task
6. **User preferences** — tone, verbosity, safety settings

Each section is independently maintainable and contributeable.

## Reference Files

The user can pin files or folders as "always in context." These are:
- Read and included in every request
- Updated when the file changes on disk
- Excluded from compaction

## Key Design Decisions

- Context compaction should be lossless where possible — dropped info is stored in the session for retrieval
- The user should never feel "the AI forgot something" — retrieval mechanisms compensate for compaction
- Compaction strategy is tunable per-session
- System prompt composition is modular — skills, workspace, and user preferences are independently injected
- Context window pressure is visible to the user via a context gauge or token count display
