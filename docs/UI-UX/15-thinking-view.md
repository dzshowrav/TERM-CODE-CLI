# 15-thinking-view.md

# Thinking View
## Mobile AI CLI UI/UX Specification
### Version 1.0

---

# Purpose

The Thinking View defines how the Mobile AI CLI communicates that the AI is processing a request before producing a response.

Its purpose is to reassure users that work is in progress without exposing internal reasoning or interrupting the conversation.

The Thinking View must remain lightweight, temporary, and terminal-native.

---

# Design Goals

The Thinking View must be

- Mobile First
- Terminal Native
- Transparent
- Minimal
- Non-Intrusive
- Real-Time
- Performance Friendly
- Keyboard Safe

---

# Design Philosophy

Users should always know that the AI is actively working.

The interface should communicate progress without revealing hidden reasoning or internal chain-of-thought.

The Thinking View is a status indicator, not a reasoning transcript.

---

# Visibility

The Thinking View appears only while the AI is preparing a response.

It automatically disappears when streaming begins.

---

# Lifecycle

```
User Prompt

↓

Queued

↓

Thinking

↓

Streaming

↓

Completed
```

---

# Display Location

The Thinking View appears inside the Conversation Area.

It is displayed as a temporary conversation item.

It never replaces the Chat Screen.

---

# Layout

```text
Assistant

Thinking...
```

---

# States

Supported

```
Queued

Thinking

Preparing Tools

Waiting

Streaming

Completed

Cancelled

Failed
```

Only one state is active.

---

# Queued

Displayed immediately after prompt submission.

Example

```text
Queued...
```

---

# Thinking

Displayed while the AI prepares its response.

Example

```text
Thinking...
```

---

# Preparing Tools

If tool selection is occurring

Display

```text
Preparing tools...
```

---

# Waiting

If the AI is waiting for

- Permission
- Tool completion
- Network response

Display

```text
Waiting...
```

---

# Streaming

The Thinking View disappears automatically.

The assistant response replaces it.

---

# Completed

No Thinking View remains after the response has completed.

---

# Cancelled

Display

```text
Cancelled
```

The placeholder disappears shortly afterward.

---

# Failed

Display

```text
Unable to generate response.
```

A retry action may be available.

---

# Animation

Allowed

Minimal terminal-friendly animation.

Examples

```text
Thinking.

Thinking..

Thinking...
```

or

```text
Thinking ⠁

Thinking ⠂

Thinking ⠄
```

Animation must remain subtle.

---

# Prohibited Animation

Never use

- Flashing
- Bouncing
- Scaling
- Rotation
- Decorative effects

---

# Duration

The Thinking View remains visible only while required.

Remove immediately when streaming begins.

---

# Replacement Behavior

```
Thinking...

↓

Assistant Response
```

No abrupt layout shifts.

---

# Conversation Behavior

The Thinking View behaves like a normal conversation item.

Users may

- Scroll
- Read previous messages
- Type the next prompt

---

# Input Behavior

The Command Input remains fully interactive.

Typing is never blocked.

---

# Status Bar Integration

Possible status values

```
Thinking

Preparing

Waiting

Ready
```

Updates immediately.

---

# Tool Integration

If tools are required

Display

```text
Preparing tools...
```

Actual tool execution appears separately in Tool Execution blocks.

---

# Streaming Integration

When the first response token arrives

```
Thinking View

↓

Removed

↓

Streaming Response Begins
```

---

# Keyboard Behavior

Keyboard opening

Reduces Conversation height.

Thinking View remains visible if necessary.

---

# Safe Area

The Thinking View always remains inside the Conversation Area.

Never overlap

- Command Input
- Status Bar
- Keyboard

---

# Accessibility

Support

- Large Fonts
- High Contrast
- Screen Readers

The displayed state must always be readable.

---

# Performance

Only update changed text.

Avoid unnecessary redraws.

Animation frequency should remain low.

---

# Error Handling

If preparation fails

Display

```text
Unable to prepare response.
```

Do not expose internal debugging information.

---

# Retry

If supported

Display

```text
Retry Available
```

Retry remains user initiated.

---

# Security

The Thinking View must never reveal

- Internal prompts
- Hidden instructions
- Chain-of-thought
- Model reasoning
- Internal planning
- Confidential tool decisions

Only high-level status information may be displayed.

---

# Responsive Behavior

On terminal resize

```
Pause Rendering

↓

Recalculate Layout

↓

Continue
```

No visible flicker.

---

# Restrictions

Never

- Display hidden reasoning
- Display chain-of-thought
- Freeze the interface
- Replace the Chat Screen
- Block typing
- Block scrolling
- Remain visible after streaming starts

---

# Example Layout

```text
Assistant

Thinking...
```

---

# Tool Preparation Example

```text
Assistant

Preparing tools...
```

---

# Waiting Example

```text
Assistant

Waiting for permission...
```

---

# Transition Example

```text
Assistant

Thinking...

↓

Assistant

Creating a React project...
```

---

# Thinking View Checklist

Every Thinking View must

- Be temporary
- Show only high-level status
- Never expose reasoning
- Stay inside the Conversation Area
- Disappear when streaming starts
- Support scrolling
- Support typing
- Respect safe areas
- Update efficiently
- Remain terminal-native

---

# Core Rules

1. The Thinking View communicates status, not reasoning.
2. Never expose chain-of-thought.
3. Display only high-level progress.
4. Remove the Thinking View when streaming begins.
5. Never block user interaction.
6. Keep animations minimal.
7. Respect the Safe Area.
8. Update the Status Bar simultaneously.
9. Optimize for incremental rendering.
10. Maintain Android Termux compatibility.

---

# Summary

The Thinking View provides a simple, transparent indication that the AI is actively processing a request before generating a response. By limiting itself to high-level status messages and avoiding disclosure of internal reasoning, it maintains user confidence while preserving privacy, performance, and a clean terminal-native experience for Android Termux.