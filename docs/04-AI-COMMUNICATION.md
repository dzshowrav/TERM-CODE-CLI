# AI Communication

## Overview

The AI communication system handles every aspect of sending messages to and receiving responses from AI models. It wraps the provider's raw API into a consistent, observable interaction pattern.

## Message Submission

When the user submits a message:

1. **Context assembly** — system prompt + conversation history + current tool results are assembled into a message array
2. **Context compaction** — if the context exceeds the model's limit, old messages are summarized or dropped (see Context Management)
3. **Tool definitions** — available tools are serialized into the provider's tool format
4. **Request dispatch** — the assembled request is sent to the provider via the appropriate transport
5. **Response handling** — the response stream is parsed and dispatched to the appropriate handler

## Streaming

All supported providers stream their responses token-by-token. The streaming system:

- Receives raw tokens as they arrive
- Accumulates them into a growing response buffer
- Emits events on each token, each complete sentence, and each complete response
- Handles provider-specific streaming formats (SSE, WebSocket, chunked HTTP)
- Maintains a configurable token throttle for display smoothness

### Streaming States
- **Connecting** — establishing connection to provider
- **Streaming** — actively receiving tokens
- **Paused** — stream is paused (user interrupt or flow control)
- **Complete** — stream finished naturally
- **Error** — stream terminated with error
- **Aborted** — user or system aborted the stream

## Thinking / Reasoning Display

When a model supports extended reasoning (a "thinking" phase):

- The thinking tokens are displayed in a distinct visual style (dimmed, italic, or in a collapsible section)
- The user can expand/collapse the thinking block
- The thinking phase is separated from the final answer by a visual divider
- The model knows when to stop thinking and start answering
- Thinking tokens are counted separately for usage tracking

The thinking display serves two purposes:
1. **Transparency** — the user sees the model's reasoning process
2. **Debugging** — the user can verify the model is reasoning correctly

## Tool Call Lifecycle

When the model requests tool execution, the system manages a complete lifecycle:

### States
1. **Planned** — model has declared intent to call a tool (name + arguments known)
2. **Pending** — waiting for user approval (if permission requires it)
3. **Executing** — tool is running
4. **Completed** — tool returned successfully
5. **Failed** — tool returned an error
6. **Timed Out** — tool exceeded its execution limit
7. **Aborted** — user or system cancelled the tool

### Lifecycle Events
- **onToolStart(name, args)** — tool call begins
- **onToolDelta(chunk)** — streaming tool output (for real-time display)
- **onToolEnd(result)** — tool completes with result
- **onToolError(error)** — tool fails with error
- **onToolTimeout(name, duration)** — tool exceeds timeout

### Parallel Tool Calls
The model may request multiple tool calls simultaneously:
- Each tool call gets its own lifecycle instance
- Tool calls can be approved/denied individually or as a batch
- Results are collected and returned to the model in the order they complete
- Failed tool calls don't block successful ones (unless configured otherwise)

## Abort

The user can abort an ongoing AI response at any point:
- Abort stops token streaming immediately
- In-flight tool calls are aborted (tool receives abort signal)
- Partial response is preserved in the session (marked as aborted)
- The system returns to input mode, ready for the next message
- Abort can also be triggered by timeout or rate-limit detection

## Request Batching

Multiple user messages may be batched before sending:
- Rapid consecutive inputs are accumulated
- After a configurable silence period, the batch is submitted
- This reduces API calls and allows the model to see more context

## Key Design Decisions

- Streaming is the default, never opt-in — the user always sees output as it arrives
- Thinking display is not configurable by the model — the user controls visibility
- Tool calls are always observable — the user sees what the model is doing
- Abort is instant — no cooldown, no confirmation for cancellation
- The communication layer abstracts provider differences behind a unified interface
- Rate limiting and retry logic is handled at this layer, not by the provider config
