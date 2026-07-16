---
name: bubbletea-architecture
description: Bubble Tea application architecture using The Elm Architecture pattern for terminal UIs
license: MIT
compatibility: opencode
metadata:
  audience: developers
  framework: bubbletea
---

## Model-Update-View Pattern
- Model: pure data struct implementing `tea.Model`
- Init: `func (m Model) Init() tea.Cmd` - startup side effects
- Update: `func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd)` - state transitions
- View: `func (m Model) View() string` - render based on current state

## Component Composition
- Parent model holds child models as fields
- Parent Update delegates to child Update when appropriate
- Parent View includes child View in rendered output
- Use `tea.Batch` to combine Cmds from multiple children
- Message delegation: parent wraps child messages for context

## State Management
- Finite state machine pattern: enum for view states
- Messages carry all context needed for the update
- No side effects in View or Model methods
- Use `tea.WindowSizeMsg` for responsive layouts

## Cmd Patterns
- `tea.Tick` for timers and animations
- `tea.Every` for periodic polling
- `tea.Sequence` for ordered async operations
- `tea.Batch` for parallel operations
- Custom Cmds via `func() tea.Msg` closures

## Production Quality
- Graceful quit handling with `tea.Quit`
- Signal handling: `tea.WithSignals(os.Interrupt, syscall.SIGTERM)`
- Error boundary: top-level error handling in main update loop
- Resource cleanup via `tea.Quit` cmd in cleanup scenarios
- Testing: simulate messages, check model state, snapshot view output
