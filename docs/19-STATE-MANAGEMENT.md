# State Management

## Overview

The state management system defines how the agent maintains, updates, and propagates its internal state. It follows an event-driven, reactive model where state changes trigger UI updates automatically.

## State Model

The system has a single source of truth for UI state, distinct from persisted data:

### State Categories
- **Session state** — current session ID, messages, active tools, scroll position
- **UI state** — current layout mode, dialog stack, focused element, theme
- **Connection state** — provider connection status, streaming state, model info
- **Input state** — current input text, cursor position, paste detection state
- **Configuration state** — active configuration (merged from all layers)
- **Permission state** — pending permission requests, policy evaluation results

### State Shape
- State is organized as a structured tree
- Each branch is independently updatable
- State is immutable — updates produce a new state tree
- Derived state is computed from source state (never stored separately)

## Event-Driven Updates

State changes are triggered by events:

### Event Types
- **User events** — keystrokes, mouse clicks, resize signals
- **AI events** — token received, tool call started, response complete
- **System events** — file changed, process exited, timer fired
- **Network events** — connected, disconnected, error, reconnected

### Event Flow
1. An event is emitted
2. The event is dispatched to registered handlers
3. Handlers compute new state based on the event
4. The new state is published
5. UI components subscribed to the changed state re-render

### Event Batching
- Multiple rapid events are batched
- State is updated once per batch
- UI re-renders once per batch
- Batching prevents unnecessary intermediate renders

## Reactive Updates

The UI reacts to state changes automatically:

### Subscription Model
- UI components declare which state branches they depend on
- When a dependency changes, the component is marked for update
- Only changed components re-render
- Unchanged components are skipped

### Derived State
- Computed values (e.g., formatted timestamps, filtered lists) are derived
- Derived state is recomputed lazily when its source changes
- Derived state is cached until a source dependency changes
- Examples: visible messages (filtered by search), context usage percentage

### Side Effects
Side effects are triggered by state changes:
- **Auto-save** — session state changes trigger deferred save
- **File watching** — file changes trigger content refresh
- **Reconnection** — network state change triggers reconnection logic
- **Attention** — certain state changes trigger attention indicators

## Performance

- State updates are batched and throttled
- Deep state trees use structural sharing (unchanged branches are reused)
- Subscriptions are granular (component subscribes to specific paths, not entire state)
- State diffs are computed for debugging and dev tools
- State is serializable for debugging and replay

## State Debugging

- State can be inspected at runtime
- State change history is recorded (circular buffer)
- Replay mode allows stepping through state changes
- State snapshots can be exported for bug reports

## Key Design Decisions

- Single state tree makes the system predictable and debuggable
- Immutable state enables time-travel debugging and undo
- Event-driven updates decouple event sources from UI rendering
- Granular subscriptions prevent unnecessary re-renders
- Derived state is computed, not stored — no stale computed values
- State is serializable — the entire UI state can be saved and restored
- Batching prevents visual flicker during rapid updates
