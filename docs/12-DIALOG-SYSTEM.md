# Dialog System

## Overview

The dialog system provides modal and non-modal interfaces for user interaction beyond simple text input. Dialogs handle selections, confirmations, forms, and multi-step workflows.

## Dialog Types

### Confirmation Dialog
- Presents a question with Yes/No/Cancel options
- Used for destructive actions (delete, overwrite, uninstall)
- Shows relevant context (what will be deleted, what will change)
- Default action is highlighted
- Keyboard shortcuts for quick response (Y/N/C)

### Selection Dialog
- Presents a list of options with descriptions
- Single-select or multi-select modes
- Search/filter support for large lists
- Keyboard navigation (arrow keys, vim keys)
- Grouped options with section headers

### Form Dialog
- Multi-field input form
- Each field has a label, type, and validation
- Field types: text, password (masked), number, boolean, select
- Tab/Shift+Tab to navigate fields
- Form validation on submit
- Supports field dependencies (show/hide based on other fields)

### Confirmation Prompt
- Lightweight yes/no prompt inline in the conversation
- No modal overlay — just a question with options
- Used for permission approval (Allow/Deny)
- Auto-dismiss if no response within timeout

### Progress Dialog
- Shows progress for long-running operations
- Indeterminate (spinner) or determinate (percentage) progress
- Operation description and elapsed time
- Cancel button to abort

### Multi-Select Switcher
- Purpose-built dialog for switching between items
- Used for: session switching, agent selection, model selection, tool list, MCP servers, history
- Shows a list of items with metadata
- Search/filter at the top
- Quick-jump by typing
- Shows current selection and recent items at top
- Contextual actions per item (rename, delete, inspect)

## Dialog Stack

Multiple dialogs can be stacked:
- Only the topmost dialog is interactive
- Closing a dialog reveals the one beneath it
- Escape closes the current dialog (or cancels if at bottom)
- Dialog stack is visual — user sees they're in a dialog hierarchy

## Dialog States

Each dialog has:
- **Open** — visible and interactive
- **Submitting** — processing user's response
- **Closing** — running close animation/cleanup
- **Closed** — removed from display

## Keyboard Navigation (Within Dialogs)

Consistent across all dialogs:
- **Tab** — next interactive element
- **Shift+Tab** — previous interactive element
- **Enter/Space** — activate selected
- **Escape** — cancel/close
- **Arrow keys** — navigate list items
- **Ctrl+F** — focus search (if available)
- **Home/End** — first/last item in list

## Dialog Triggers

Dialogs can be opened by:
- Slash commands (e.g., `/agents` opens the agent switcher)
- Tool permission requests (approval prompt)
- User shortcuts (Ctrl+P for command palette)
- System events (update available, error, warning)
- Programmatic requests (AI requests user confirmation)

## Accessibility

- All dialogs work without a mouse
- Screen reader friendly (label associations)
- High contrast mode for all dialog elements
- No flashing or rapid animation
- Keyboard-only navigation always available

## Key Design Decisions

- Dialogs are modal but not blocking — the session continues in the background
- Escape is universal "go back / cancel" — the user can always back out
- Search in selection dialogs makes large lists navigable instantly
- Dialog stack prevents deep nesting confusion (only one dialog active at a time)
- Forms support full validation before submission
- Every dialog has a keyboard shortcut for quick action
- Confirmation dialogs show context, not just "are you sure?"
