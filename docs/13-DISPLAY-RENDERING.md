# Display Rendering

## Overview

The display rendering system converts AI responses, tool results, and system messages into formatted terminal output. It handles markdown, code blocks, diffs, streaming text, and syntax highlighting.

## Markdown Rendering

Markdown output from the AI is rendered with full formatting:

### Elements
- **Headings** (# through ######) — rendered with bold or inverse styling
- **Bold/Italic** — standard terminal bold/italic
- **Lists** (ordered and unordered) — indented with bullets or numbers
- **Links** — displayed with descriptive text, URL shown on hover or in status bar
- **Blockquotes** — indented with a vertical bar
- **Horizontal rules** — terminal-width separator line
- **Tables** — rendered with column alignment and borders
- **Task lists** — checkboxes (checked/unchecked)

### Code Blocks
- Fenced code blocks (```) are rendered with:
  - A header showing the language name
  - Syntax highlighting (see below)
  - Line numbers (optional, configurable)
  - A copy button (or shortcut) to copy the content
  - Horizontal scroll for long lines
  - Collapsible for very large blocks
- Inline code (`code`) — rendered with a distinct background or color

### Image Placeholders
- Markdown images show the alt text and URL
- If the terminal supports images (Kitty protocol, iTerm2), display inline
- Otherwise, show a placeholder with dimensions

## Syntax Highlighting

Code blocks are highlighted based on language:

- Support for all common languages (JS, TS, Python, Rust, Go, Java, Ruby, C++, SQL, YAML, JSON, HTML, CSS, Shell, etc.)
- Highlighting is performed off the main rendering thread (background computation)
- Colors come from the active theme
- Highlighting handles:
  - Keywords, strings, numbers, comments, operators, types, functions, variables
  - Multi-line constructs
  - Embedded languages (JS in HTML, SQL in Python)
- On-theme-change, highlighting is recomputed
- Fallback to plain display if the language is unknown

## Streaming Text Display

During AI response generation:
- Tokens are displayed as they arrive (character-by-character or word-by-word)
- The display updates incrementally — previously rendered content is not re-rendered
- Streaming text is rendered through the markdown pipeline incrementally
- The cursor/insertion point is shown at the end of the stream
- The user can scroll back to read previous content while streaming continues

### Streaming Challenges
- Incomplete markdown is rendered progressively (e.g., opening ``` without closing still shows code-like styling)
- Lists and tables are re-flowed as new tokens arrive
- The display never "jumps" — content is appended, not replaced
- Streaming smoothes out via token throttling (configurable display rate)

## Diff Visualization

File diffs are displayed with:

### Diff Elements
- **File header** — filename and change summary (+N / -M)
- **Hunk header** — line range context (@@ -L,S +L,S @@)
- **Added lines** — green/highlighted with leading +
- **Removed lines** — red/highlighted with leading -
- **Context lines** — normal display
- **Line numbers** — old and new line numbers side by side

### Diff Enhancements
- **Word-level highlighting** — within-line changes are highlighted
- **Collapsible hunks** — large diffs can be collapsed by hunk
- **Scrollable** — long diffs scroll within a fixed-height viewport
- **Animation** — lines transition in (optional, configurable)
- **Copy** — shortcut to copy the diff content
- **Apply** — shortcut to apply the diff (with confirmation)

### Diff Metadata
- Number of files changed
- Insertions/deletions count per file
- Total diff size
- Language detection per file

## Tool Call Cards

When the AI executes a tool, a visual card shows the progress:

### Card Structure
- **Tool name** — icon + name (e.g., "Read File", "Execute Command")
- **Status indicator** — spinning (in progress), check (done), X (failed), clock (timeout)
- **Arguments** — formatted tool arguments (truncated if long)
- **Status message** — dynamic, AI-generated status text (e.g., "Reading 15 files…", "Compiling project…")
- **Duration** — elapsed time during execution, total time after completion
- **Result preview** — first few lines of output, with expand option

### States
- **Queued** — dimmed, grey, waiting
- **Running** — animated spinner, status message updates in real-time
- **Complete** — green check, result shown, duration displayed
- **Failed** — red X, error message, retry option
- **Timed Out** — clock icon, partial output shown
- **Aborted** — cancelled icon, partial output shown

### Card Transitions
- Cards smoothly transition between states (no abrupt UI jumps)
- Completed/failed cards collapse to a compact summary after a configurable timeout
- Multiple parallel tool calls show stacked cards

## Performance

- Rendering is incremental — only changed parts of the display are updated
- Syntax highlighting runs on a background thread/process
- Long output is virtualized (only visible portion is rendered)
- Markdown parsing is streaming — partial results can be rendered
- All rendering is stateless — given the same state, the same output is produced

## Key Design Decisions

- Markdown rendering is live — the user sees it take shape as the AI generates
- Syntax highlighting is off-thread to keep input responsive
- Diffs are the primary way file changes are communicated — the user always sees what changed
- Tool call cards provide real-time visibility into the agent's actions
- The display never blocks on rendering — new input is always accepted
- All rendered output is scannable — the user can quickly find what matters
