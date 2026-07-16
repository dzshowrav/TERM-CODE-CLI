# UI Shell

## Overview

The UI shell is the structural layout of the terminal interface. It defines how information is organized across the available screen space, what regions exist, and how they behave under different terminal sizes.

## Screen Regions

The terminal screen is divided into conceptual regions:

### Region 1: Status Bar (Top)
A single-line bar at the top showing:
- **Session name** — current session identifier or title
- **Model/provider** — currently active model with provider indicator
- **Mode indicators** — auto mode, compact mode, silent mode
- **Network status** — connected/connecting/disconnected/reconnecting
- **Context gauge** — approximate context usage (visual bar or percentage)
- **MCP count** — number of active MCP connections
- **Git branch** — current git branch (if in a git repo)
- **Attention indicator** — visual cue when the agent needs attention
- **Dynamic status words** — AI-generated one-word status messages that cycle during tool execution (e.g., "thinking", "analyzing", "writing", "searching", "compiling")

The status bar is always visible (except in ultra-compact mode where it may be hidden).

### Region 2: Main Content (Center)
The primary conversation area:
- **Message history** — scrollable list of user and AI messages
- **Tool call cards** — real-time tool execution display (interleaved with messages)
- **Streaming response** — current AI response being generated
- **System messages** — status updates, errors, notifications

### Region 3: Input Area (Bottom)
The text input region:
- **Prompt prefix** — cursor indicator (> or $ or custom)
- **Input buffer** — the user's current input text
- **Character count** — optional, shown when approaching limits
- **Mode indicator** — insert/normal mode, input mode
- **Suggestion dropdown** — command completion or @-mention suggestions

## Layout Modes

### Normal Mode
Full layout with all three regions visible. Status bar at top, content area (taking remaining space), input at bottom.

### Compact Mode
Optimized for smaller terminals (< 30 rows):
- Status bar is minimized (only essential indicators)
- Content area is the full terminal minus input line
- Tool call cards are compact (one-line summary instead of full cards)
- Input area is single-line (no multi-line editing visible until expanded)

### Ultra-Compact Mode
For very small terminals (< 15 rows):
- Status bar is hidden entirely
- Content area shares space with input (input overlays the last line of content)
- Messages are shown without headers (just content)
- Tool calls are shown as icons only
- Only the most recent messages are visible

### Mobile Mode
For mobile-terminal users (< 18 rows, touch keyboard):
- All compact mode features
- Input area has larger touch targets
- Long-press detection is more sensitive
- Additional padding for touch keyboard overlap
- Simplified dialog layouts (full-screen instead of modal)
- Status bar minimized to a single line with only critical info

## Status Bar (Detailed)

The status bar is composed of segments:
- **Left-aligned** — session, mode indicators, status words
- **Center** — dynamic status (current operation, if any)
- **Right-aligned** — model, context, git, MCP, attention

Segments auto-hide when the terminal is too narrow to show them all. Priority determines which segments are dropped first (lowest priority hidden first).

## Attention System

The agent needs attention when:
- A tool requires permission
- A command completes with an error
- A background task finishes
- The AI generates a question or asks for input

Attention is indicated by:
- **Visual cue** — status bar indicator (colored dot or icon)
- **Status bar flash** — brief color change on the status bar
- **Terminal bell** — audible beep (configurable)
- **System notification** — desktop notification (configurable)
- **Title bar update** — terminal title blinks or changes

The user can configure which attention signals are active.

## Layout Transitions

When the terminal is resized:
1. The layout recalculates within one render cycle
2. Content reflows to fit the new dimensions
3. If the terminal shrinks below a threshold, the layout mode auto-switches
4. If the terminal grows, the layout mode switches back
5. Input content is preserved across resize (no lost characters)
6. Scroll position adjusts to keep the latest content visible

## Key Design Decisions

- Status bar is informational, not interactive — it's purely a display
- Layout modes auto-switch based on terminal size, not user preference (user can override)
- Ultra-compact mode is intentionally minimal — the agent is still fully functional
- Attention system is non-intrusive by default — the user isn't bombarded
- Every region is independently scrollable — the user can review history while input is active
- The input area is always visible — the user is always ready to type
- Dynamic status words make tool execution feel alive and responsive
