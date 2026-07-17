# TermCode Project Rules

## Build & Test Commands
- Build: `go build ./...` (CGO_ENABLED=0 for static builds)
- Test: `go test ./...` or `go test -v -count=1 ./...` (no cache)
- Lint: `gofumpt -l -w . && goimports -local -w . && go vet ./...`
- Type check: `go build ./...` (Go compiler enforces types)
- Format: `gofumpt -l -w .` and `goimports -local -w .`
- Git Push: **Always run `go build ./... && go vet ./... && go test ./...` before any `git push`. Never push without a clean build first.**

## Project Structure
- Go modules with `cmd/`, `internal/`, `pkg/` layout
- Clean Architecture layers: domain, application, adapters, infrastructure
- Bubble Tea apps follow Model-Update-View pattern

## Core Conventions
- Mobile First: design for 80x24 terminal, 256 colors, aarch64/arm64
- Termux Only: CGO_ENABLED=0, static binaries, no platform-specific deps
- No Mock Data: use real implementations, fakes, or test containers
- No Emoji: text-only output
- Full File Output: write complete files, never partial diffs
- Modular Code: small focused packages, clear interfaces
- Production Ready: graceful shutdown, structured logging, signal handling

## Code Standards
- Go best practices: idiomatic patterns, error wrapping, table-driven tests
- Bubble Tea Architecture: Elm-style model/update/view, cmd/message patterns
- Clean Architecture: dependency inversion, domain isolation

## Custom Agents (Legacy)
- @go-expert - General Go development
- @bubbletea-expert - General Bubble Tea TUI
- @lipgloss-expert - Terminal styling
- @mcp-expert - General MCP server development
- @treesitter-expert - Code parsing and analysis
- @postgres-expert - Database schema and queries
- @redis-expert - Caching and data structures
- @architect - System design and architecture review
- @security-auditor - General security analysis
- @perf-engineer - General performance optimization

## TermCode Multi-Agent System
Orchestration layer (planning/reasoning):
- @master-architect - Architecture decisions, agent coordination
- @task-planner - Request decomposition and execution planning
- @context-engine - Context collection and distribution
- @memory-engine - Long-term project knowledge
- @reasoning-engine - Structured decision making

Engineering agents (implementation):
- @go-engineer - Go business logic, services, domain models
- @bubbletea-engineer - Bubble Tea TUI architecture and screens
- @uiux-engineer - Mobile-first terminal UX and interaction design
- @terminal-engineer - ANSI, Unicode, keyboard, viewport
- @mcp-engineer - MCP integrations and tool execution
- @database-engineer - Schema, migrations, queries, persistence
- @git-engineer - Version control and repository workflow
- @testing-engineer - Quality assurance and automated testing
- @security-engineer - Vulnerability prevention and data protection
- @performance-engineer - Startup, memory, CPU, rendering optimization

Quality assurance layer:
- @documentation-engineer - Documentation and knowledge management
- @review-engineer - Final quality gate and code review
- @release-engineer - Release management and versioning

## Deployment Lessons
- **System binary vs local build**: After building with `go build -o tc ./cmd/tc/`, always replace the system binary at `/data/data/com.termux/files/usr/bin/tc` by running `cp ./tc /usr/bin/tc` (kill the running process first with `kill $(pidof tc)` if "Text file busy").
- Always verify which binary is running via `which tc` or `type tc` when changes don't appear to take effect.

## Automatic Lesson Extraction
Every bug fix or issue resolved MUST produce a "Learned Lessons" entry below. When you fix a bug,
identify the root cause pattern and write a specific, actionable rule that prevents recurrence.
Never fix a bug without recording the lesson.

## Learned Lessons

### BaseURL normalization — never store `/v1` suffix
- **Always strip `/v1` suffix from provider base URLs** at the service layer when creating/updating
  (via `normalizeBaseURL()`), since the code appends `/v1/...` paths at request time. Stored URLs
  should be the bare origin (e.g. `https://api.openai.com`), not `https://api.openai.com/v1`.
- **Placeholder text must match the expected stored format** — the "URL" field hint should show
  `https://api.openai.com` without `/v1`.
- **Always add HTML detection to error responses** when calling external APIs. A malformed URL
  often returns a redirect to an HTML page instead of JSON; detect `<!DOCTYPE` or `<html` and
  replace with a human-readable message instead of dumping raw HTML.
- **Normalize at every usage point, not just at write time.** Existing stored data may already have
  the `/v1` suffix. The `normalizeBaseURL()` helper must be called at every URL construction site
  (adapter constructors, discovery, test connections), not only in `Create`/`Update`.

### Keyboard-aware layout — handle terminal resize for virtual keyboard
- **Viewport height must account for all fixed elements** (model header, separator, command input,
  status bar). Formula: `vpHeight = totalHeight - numFixedElements`, not `totalHeight - 1`.
  On Termux, opening the virtual keyboard shrinks the terminal; the layout must fit exactly
  without overflow.
- **Always call `GotoBottom()` on viewport resize** — when the terminal height changes (keyboard
  opens/closes), the viewport's offset can leave latest messages scrolled out of view.
  `SetSize` must scroll to bottom so the most recent content stays visible.

### Termux scroll flicker — use MouseMode + AltScreen on `tea.View`
- **Root cause**: On Termux, touch-scroll makes the emulator scroll its native scrollback buffer
  while Bubble Tea simultaneously re-renders, creating a collision that causes flicker.
- **Fix**: Set `v.MouseMode = tea.MouseModeAllMotion` on the `tea.View` returned by `View()`. This
  tells the terminal to send touch/scroll interactions as mouse wheel messages to the program
  instead of handling them natively.
- **AltScreen**: Set `v.AltScreen = true` on the view for an isolated rendering buffer.
- **MouseWheelMsg**: Handle `case tea.MouseWheelMsg:` with `msg.Mouse().Button == tea.MouseWheelUp/Down`
  to process scroll events from touch gestures.
- **Program options**: Add `tea.WithoutSignalHandler()` and `tea.WithFPS(30)` to `tea.NewProgram()`.
- **Fallback**: If flicker persists, user can tap the SCROLL lock button (⇳) in Termux extra keys row.
- In Bubble Tea v2, mouse mode and alt screen are fields on `tea.View` struct, not program options.

### Tool execution UI — wire tool events through callbacks + prog.Send
- Tool execution happens in a goroutine inside `conversation.Service.converse()`. To show it in the UI,
  add a `ToolEventCallback` parameter that fires events at each lifecycle stage (queued/started/output/
  completed). The callback sends typed messages to the Bubble Tea program via `prog.Send()`.
- ChatScreen maintains a `[]*components.ToolCard` slice. Tool event messages (ToolQueuedMsg,
  ToolStartedMsg, ToolCompletedMsg) are handled in `ChatScreen.Update` to add/update cards.
- The ToolCard component follows the spec: status icons (○/●/✓/✗/…), color coding (green/blue/
  yellow/red), collapsed/expanded states, progress bar, arguments panel, truncated output, duration.
- Auto-collapse completed cards would ideally use `tea.Tick` with a 2s delay, but for now
  expand/collapse is toggled via Enter key.

### "clean cache" / "clear cache" means `go clean -cache`
- When a user says "clean cache" or "clear cache", run `go clean -cache` to clear the Go build
  cache. This is a Go toolchain command, not related to any application-level caching feature.
- **NEVER decompose PasteMsg into KeyPressMsgs** — this breaks non-ASCII characters
  (multi-byte UTF-8 produces `len() > 1`, filtered out by `len(msg.String()) == 1` guard).
  Instead, forward PasteMsg directly to the input component's `Update`.
- **Every input component must handle PasteMsg** — it's not optional. When creating a new dialog
  with form fields, check that `case tea.PasteMsg:` is present alongside `case tea.KeyMsg:`.
- **When adding a new dialog, audit message types** — grep existing dialogs for the full set of
  `case` patterns (`tea.KeyMsg`, `tea.PasteMsg`, custom messages) and match them all.

### Centralized layout — single authority for all component sizes
- **NEVER distribute layout height math across components.** Having `ChatScreen.SetSize`
  subtract overhead that it doesn't own (input, status bar, palette) creates a fragile implicit
  contract that breaks when any component changes height. Instead, use a single `layout()`
  method on the top-level model that computes every component's size from total terminal height
  minus every fixed element.
- **Account for border heights.** `lipgloss.BorderTop(true)` adds 1 line; `BorderBottom(true)`
  adds 1 line. A bordered component is 2–3 lines, not 1. Always count actual rendered lines.
- **Add 1 line of bottom safe-area padding.** Termux keyboards may overlay or clip the very last
  line. Reserve one unused line at the bottom so the status bar is never hidden behind the
  keyboard.
- **Call layout() after every state change** that affects any component's height (palette show/
  hide/filter, terminal resize). Use a single dedicated method rather than inline size calls.
- **ChatScreen exposes `SetViewportSize(w, h)`** — it takes viewport dimensions directly and
  does NO math beyond storing width. All layout arithmetic lives in the caller.
- **Call `layout()` after every palette navigation event** (Up/Down changes scroll indicators
  that change palette height). Palette arrow handlers returned early without relayout, causing
  the palette to grow at the bottom edge and overflow the screen.

### Dialog box design — rounded border, full interactivity, dynamic layout
- **Dialog boxes must always use `styles.DialogBox()`** which wraps content in a rounded-border
  box (`╭─╮││╰─╯`) with colorized borders (240). Never use `styles.Content()` (plain padding)
  for dialogs — that's for full-page screens.
- **Every dialog must be fully interactive**: Tab/Shift+Tab/Up/Down cycles through BOTH input
  fields AND action buttons. `focusField` must cover the entire interactive surface (fields +
  buttons), not just input fields.
- **Buttons must visually reflect focus state** — use `formBtnNormal` (dim 245) when unfocused
  and `formBtnActive` (blue 39 + bold) when focused. Enter on a focused button triggers its
  action.
- **Prevent text input when button is focused** — guard `addChar`/`deleteChar` with
  `if s.focusField < len(fields)` so keyboard events don't leak into fields when navigating.
- **Always handle `tea.PasteMsg`** in every input-capable dialog. Decomposed multi-byte UTF-8
  characters (len > 1) are silently dropped by the `len(msg.String()) == 1` guard in the
  KeyMsg default handler.
- **Separator lines inside dialogs must use `styles.DialogSep(width)`** (color 236) instead of
  bare `strings.Repeat("─", width)` for consistent colorization.

### Interactive selection dialogs — search, scroll, cursor
- **DialogScreen goes in `screens/`, not `dialogs/`**. The old dialog pattern (`dialogs/`) uses `tea.Model` directly; new code uses the `DialogScreen` interface (`SetSize`, `Update`, `View`, `Done`, `Result`).
- **Search filtering**: On each keystroke, rebuild `filtered` from `models`. Reset cursor/scroll to 0 after filtering. Clamp cursor to `max(0, len(filtered)-1)`.
- **Scrolling**: Maintain `cursor` (current selection) and `scroll` (viewport offset). `ensureVisible()` adjusts scroll when cursor goes out of view. `entryBounds()` computes max visible entries from terminal height minus fixed overhead.
- **Height accounting for terminal**: Each dialog has 7 fixed body lines (title, sep, search, sep, header, blank, bottomBar) + model entries + scroll indicators + 2 dialog borders + 3 AppModel overhead lines. `entryBounds()` returns `s.height - 12 - scrollIndicators`.
- **SetModels takes value types**: `modelSvc.List()` returns `[]*domainmodel.Model`. Dereference with `models[i] = *m` before passing to `SetModels([]domainmodel.Model)`.
- **`tea.PasteMsg` has a `String()` method**: Use `msg.String()` not `string(msg)`. PasteMsg is a struct with a `Content` string field.
- **Type-to-search**: Capture typing in Update's default KeyMsg case. Convert to `[]rune` and check `!unicode.IsControl(r[0])` to filter non-printing chars.
- **Free tag**: Check `m.PricingInput == 0 && m.PricingOut == 0` to show "Free" badge. Render with `styles.Active` (blue/bold).
- **Named returns vs `:=`**: In `func() (x int)`, use `x = 1` not `x := 1` — `:=` shadows the named return and the outer variable goes unused.

## Custom Commands
- /feature - Create a new feature
- /component - Create a Bubble Tea component
- /command - Create a custom command
- /mcp - Create an MCP server
- /screen - Create a Bubble Tea screen
- /test - Write tests
- /docs - Write documentation
