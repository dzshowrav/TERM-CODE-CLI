# Tool Architecture

## Overview

Tools are the agent's primary means of interacting with the user's system. Each tool encapsulates a capability: read a file, write a file, run a command, search code, etc. The tool system manages the lifecycle of tool definition, registration, execution, and result reporting.

## Tool Definition

Every tool has a structured definition:

### Core Fields
- **Name** — unique identifier (snake_case, e.g., `read_file`, `execute_command`)
- **Display name** — human-readable label for UI display
- **Description** — what the tool does, when to use it (used by the AI model to choose tools)
- **Parameters** — JSON Schema definition of required and optional inputs
- **Category** — functional grouping (file, shell, search, network, etc.)

### Lifecycle Hooks (Optional)
- **onBefore** — runs before execution (validation, logging)
- **onStart** — runs when execution begins
- **onDelta** — receives intermediate output during execution (for streaming tools)
- **onEnd** — runs when execution completes successfully
- **onError** — runs on failure
- **onAbort** — runs when execution is cancelled

### Execution Constraints
- **Timeout** — maximum execution duration (default: 30 seconds)
- **Allowed contexts** — what the tool can access (filesystem, network, environment)
- **Permission level** — what level of user approval is required
- **Rate limit** — how often the tool can be called

## Tool Registration

Tools are registered from multiple sources:
1. **Built-in tools** — shipped with the agent (file ops, shell, search, etc.)
2. **Skill tools** — tools defined by loaded skills
3. **MCP tools** — tools exposed by external MCP servers
4. **User-defined tools** — custom tools created by the user

Registration involves:
- Validating the tool definition
- Resolving conflicts (same-name tools are flagged)
- Indexing for AI model selection
- Categorizing for UI display

## Tool Execution Pipeline

When the AI model requests a tool call:

1. **Parse** — extract tool name and arguments from the model's response
2. **Validate** — validate arguments against the parameter schema
3. **Permission check** — check if user approval is needed
4. **Pre-execution hooks** — run onBefore hooks (logging, metrics, pre-validation)
5. **Execute** — run the tool's handler with validated arguments
6. **Stream** — if the tool supports streaming, emit delta events
7. **Complete** — capture the final result
8. **Post-execution hooks** — run onEnd hooks (result processing, context injection)
9. **Return** — format and return the result to the AI model

## Timeout Handling

Every tool execution has a timeout:
- Default: 30 seconds (configurable per tool, per session)
- When timeout fires:
  1. The tool receives an abort signal
  2. Any partial output is captured and returned
  3. The model is informed the tool timed out
  4. The model may retry the tool or proceed without it

Timeout is enforced via:
- An external timer that fires after the configured duration
- An abort signal passed to the tool handler

## Streaming Tool Output

Some tools produce output incrementally (e.g., shell commands with ongoing output). These tools:
- Emit delta events as output arrives
- The UI updates in real-time as deltas arrive
- The final result includes all accumulated output
- The model receives only the final result (not intermediate deltas)

## Tool Results

Every tool execution produces a result:
- **Status** — success, error, timeout, aborted
- **Output** — the tool's return value (text, structured data, or both)
- **Duration** — how long execution took
- **Metadata** — tool-specific metadata (file path, command, etc.)
- **Truncation info** — if output was truncated, how much was kept

Results are:
- Returned to the AI model for the next response
- Stored in the session for later reference
- Displayed in the UI (formatted by result type)
- Available for the user to inspect

## Tool Discovery

The user can:
- List all available tools with descriptions
- Inspect a specific tool's definition and parameters
- Enable/disable tools per session
- Test a tool interactively

## Key Design Decisions

- Tools are pure functions of their inputs — no hidden state between calls
- Tool execution is observable at every stage (start, delta, end, error)
- The AI model never directly executes tools — it only requests execution
- Tool timeout prevents runaway operations
- Tool outputs can be large — truncation is applied, but full output is storable
- Tools are immutable (defined once, used many times) — state lives in the session
