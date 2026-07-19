# TermCode Project Rules

## Build & Test Commands
- Build: `go build ./...` (CGO_ENABLED=0 for static builds)
- Test: `go test ./...` or `go test -v -count=1 ./...` (no cache)
- Lint: `gofumpt -l -w . && goimports -local -w . && go vet ./...`
- Type check: `go build ./...` (Go compiler enforces types)
- Format: `gofumpt -l -w .` and `goimports -local -w .`
- Git Push: **Always run `go build ./... && go vet ./... && go test ./...` before any `git push`. Never push without a clean build first.**
- **Build + Vet + Deploy after every change**: Always run `go build ./... && go vet ./...` then `go build -o tc ./cmd/tc/ && kill $(pidof tc) 2>/dev/null; cp -f ./tc /data/data/com.termux/files/usr/bin/tc` after every code modification. Never leave the binary stale.

## Project Structure
- Go modules with `cmd/`, `internal/`, `pkg/` layout
- Clean Architecture layers: domain, application, adapters, infrastructure
- Bubble Tea apps follow Model-Update-View pattern

## Core Conventions
- Mobile First: design for 80x24 terminal, 256 colors, aarch64/arm64
- Termux Only: CGO_ENABLED=0, static binaries, no platform-specific deps
- No Mock Data: use real implementations, fakes, or test containers
- No Emoji: text-only output
- Bangla in English letters only: When writing Bengali (Bangla), always use English alphabet transliteration (e.g. "bangla", "kemon acho"), never Bengali script characters. This is permanent and applies to ALL communication.
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

### HideThinking must reset thinkingDur — never let "Working..." leak past response
- **Root cause**: `HideThinking()` only cleared `showThinking` but not `thinkingDur`. The `else if` chain in `renderMessages()` hit `thinkingDur > 0` and rendered a stale "Working..." line that persisted after the response was already displayed.
- **Fix**: Always reset `thinkingDur = 0` inside `HideThinking()` — it represents the same "thinking is done" state as `showThinking = false`.
- **Check every state-clearing method**: `HideThinking()`, `ClearToolCards()`, and any "done" handler must zero out ALL related state fields, not just the primary flag.

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

### Cursor visibility — never hide, never override with placeholder
- **Cursor must always be visible on the focused field**. Never append `"█"` *then* check if value is empty — that throws the cursor away. Check `val == "" && i != focusField` *first* to show hint only for unfocused empty fields, then append cursor for the focused field unconditionally.
- **Pattern**: `if val == "" && i != s.focusField { val = hint }` → `if i == s.focusField { val += "█" }`. This guarantees the cursor shows even when the focused field is empty.
- **Check every input dialog** for this bug: `provider_add_screen.go`, `model_add_screen.go` both had it. The tell is `if val == "" || val == "█"` which catches the cursor itself and replaces it.
- **NEW INPUTS MUST FOLLOW THIS PATTERN** — every future input component/dialog must render the cursor unconditionally on the focused field. The "hint for empty unfocused" check must happen *before* the cursor append, never after.

### ChatScreen anomalies — state consistency, message routing, focus rendering
- **`ClearToolCards` must reset `thinkingDur`** — stale `thinkingDur > 0` leaks through the `else if` chain in `renderMessages()` after clearing, causing a phantom "▶ Thinking X.Xs" line to persist with no tool cards.
- **Every state-mutating handler must call `renderMessages()` or delegate through `UpdateToolCard`** — `handleCardFocusKey` up/down navigation changed `focusedCard` but omitted `renderMessages()`, so focus indicators didn't update on screen until the next interactive key (Enter/Space).
- **`ToolOutputMsg` and `ToolCompletedMsg` must use `msg.Index`** — unlike `ToolStartedMsg`, these handlers always targeted `len(s.toolCards)-1`, ignoring the explicit index from the callback. When `UpdateToolCard` already guards bounds, use the `msg.Index` pattern consistently across all tool lifecycle handlers.
- **`persistToolResults()` must be followed by `renderMessages()`** — tool metrics persisted to `s.messages` never appeared until the next unrelated re-render.
- **Alert pop-ups: "Press any key to search by name..."** appears because typing `/` to start a command (like `/help`) is intercepted by the chat screen's search-mode toggle *before* reaching the command input. `/` must be guarded with `len(s.toolCards) > 0` so it only activates when there are actually tool cards to search. Additionally, when `/` enters search mode, `focusedCard` must be set to `0` (not `-1`) so subsequent key presses route to `handleCardFocusKey` (which implements search input), not `handleScrollKey` (which ignores them).
- **Split msg.Content on `\n` before adding to viewport lines** — `renderMessages()` added `msg.Content` as a single viewport element. When the AI response was large with embedded newlines (paragraphs, code blocks), the viewport's wrapping algorithm split across `\n` characters, corrupting the display and breaking scroll offset calculations. Fix: split each `msg.Content` on `\n` so the viewport properly counts every line.

### Session management — startup, dialog, navigation
- **App startup: fresh session, home screen, no old session load** — Never call `loadLastSession()` in `SetProviderService`. Use only `ensureSession()` + `m.screen = screenHome`. Old sessions remain accessible via `/sessions` dialog — users select one manually to resume.
- **Session dialog Enter must return session ID** — `session_screen.go` Enter handler left `result` empty, so selecting a session just closed the dialog. Fixed by setting `s.result = s.sessions[s.cursor].ID` and adding `loadSessionByID()` to load session messages into a fresh chat screen.
- **New session from `/sessions new` must go to home screen** — `cmdSessions()` should set `screen = screenHome` and return `__home__`, not `__chat__`. New session is created but home screen shows with fresh session ready. User goes to chat only when they type a message.

### Tool execution pipeline — validate, timeout, truncate, hooks
- **Execution pipeline**: Every tool goes through `Validate → OnBefore → OnStart → Execute → OnEnd/OnError → Truncate → Return`. Validation checks required fields and types against the JSON schema before execution starts.
- **Timeout enforcement**: Use `context.WithTimeout` with a goroutine + select pattern. Default 30s per tool. When timeout fires, set `StatusTimeout` and call `OnAbort` hook. The deferred cancel prevents resource leaks.
- **Result truncation**: Large outputs (>100KB) are truncated with `Result.Truncate(maxBytes)`. Always store `RawSize` and set `Truncated` flag. The model sees truncated output but knows how much was cut.
- **Goroutine safety**: Tool executors run in a goroutine with a recover-deferred `done` channel. Panics in tool code are caught and returned as errors, not crashes.
- **Executors return errors, not set status**: Each `exec*` method returns `error` (nil on success). Set `result.Error` in the executor, let the pipeline set `result.Status`. This keeps status logic centralized.
- **Registry aliases in index**: When registering a tool, add all aliases to the registry's name→index map so `Lookup` works for both primary names and aliases. Remove alias entries when a tool is removed.
- **Hook execution is optional but ordered**: Check each hook for nil before calling. Pipeline order: OnBefore → OnStart → (execute) → OnEnd (success) / OnError (failure) / OnAbort (timeout). Hooks receive the full context.

### Tool model — always add Category/Aliases/Capabilities to new Tool entries
- **Every tool must declare Category, Aliases, Capabilities, and Dangerous flag** alongside Name/Description/InputSchema. Define tools using `tool.New(name, desc, schema, category, capabilities...)` then set additional fields (Aliases, Dangerous, Version, Author) as needed.
- **`tool.MatchName(s string) bool`** resolves both primary name and aliases. Tool service `Execute()` should only switch on the primary name; alias resolution is the caller's responsibility.
- **`AvailableTools()` returns the full set** — the LLM sees these. `Execute()` must handle every tool returned, or return `"not implemented"` for stubs. A switch on tool name is the simplest dispatch.
- **Build function literals for tool definitions** to set fields after `tool.New()` — Go's struct composite literal syntax can't call methods; use an IIFE: `func() tool.Tool { t := tool.New(...); t.Aliases = ...; return t }()`.
- **Don't import gogit in the tool service** — the git service already wraps all types. Use `s.git.Open(path)` directly (with `:=` type inference), don't create a helper with an explicit return type annotation.

### View must set AltScreen on first render — or Bubble Tea v2 never sends WindowSizeMsg
- **Root cause**: The `View()` method returned `tea.NewView("Initializing...")` without setting
  `v.AltScreen = true`. Bubble Tea v2 only queries terminal size and sends `WindowSizeMsg` when
  the alternate screen is active. Without it, the program runs in inline mode, never receives
  `WindowSizeMsg`, and `ready` stays `false` forever — stuck on "Initializing...".
- **Fix**: Always set `v.AltScreen = true` and `v.MouseMode = tea.MouseModeAllMotion` on every
  `tea.View` returned from `View()`, including the initial "not ready" view.
- **Belt-and-suspenders**: Set `ready = true` in `NewApp()` since `layout()` already handles
  zero dimensions gracefully (`if m.width == 0 || m.height == 0 { return }`). This removes the
  dependency on `WindowSizeMsg` for basic rendering.
- **First render guard is dead code**: With `ready = true` from the start, the `if !m.ready`
  branch in View() never executes — but keep it as a safety net with proper AltScreen/MouseMode
  flags so it works if ever needed.

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

### Thinking animation — ShowThinking must start its own tick chain
- **Self-contained tick**: `ShowThinking()` must return a `tea.Cmd` that starts a
  `ThinkingTickMsg` tick chain at 100ms. Do NOT rely on forwarding global ticks
  (like `AnimTickMsg`) — the chat screen may not be active when the tick fires.
- **Tick re-fires while active**: The `ThinkingTickMsg` handler re-fires itself
  as long as `showThinking || thinkingDur > 0 || hasActiveToolCards()`. This
  single tick chain covers all animation: thinking spinner, live timer, and tool
  card frame updates.
- **Tool events don't start their own ticks**: `ToolQueuedMsg` and
  `ToolStartedMsg` should NOT return `tea.Tick()` — the thinking tick chain
  keeps running via `hasActiveToolCards()` and `IsRunning()` checks.
- **Braille spinner at 100ms**: Use `[]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}` with `s.animTick%10` and a 100ms tick interval for smooth animation.
- **Live timer**: Calculate elapsed time at render time with `time.Since(s.thinkingStart)`.
  The periodic `renderMessages()` call from the tick keeps the counter live.
- **responseText flushes before next user message**: Use `FlushResponse()` (adds
  to `s.messages` as `{Role: "assistant"}`) before `AddMessage("user", ...)` in
  the submit handler. Do NOT flush in `ClearToolCards()` — that causes
  duplicate rendering of old and new responses.
- **Animation tick after tools**: Never use a separate `ToolAnimTick` type;
  `ThinkingTickMsg` is the single animation tick message. `ToolAnimTick` type
  should be removed once all references are migrated.

### Double vertical padding in dialogs — trust AppModel to center
- **Root cause**: `MCPAuthDialog`, `ToolConfirmDialog`, and `WorkspaceTrustDialog` calculated vertical spacing and padded their output with newlines internally, while the outer `AppModel.View()` layout also vertically padded active dialogs. This resulted in double padding, pushing the dialog too far down or off the bottom of the screen.
- **Fix**: Remove all vertical padding calculations and raw newline additions from individual dialog `View()` implementations, leaving them to return only their unpadded box. Allow `AppModel`'s layout to serve as the single authority for vertical dialog centering.

### Variable-height list scrolling — check lines instead of items
- **Root cause**: The `ToolListScreen` calculated maximum visible items based on a fixed item height of 1 line. When a tool card was expanded to show details (adding ~10 lines), it caused the dialog content to exceed the terminal boundary and overflow.
- **Fix**: Instead of assuming 1 line per item, dynamically calculate the total rendered height inside `ensureVisible()` and `View()` by calling `strings.Count(renderedLine, "\n") + 1` for each visible item. Only append items to the viewport list if they fit within the dynamic available vertical height.

### Out-of-bounds slicing panic on custom scroll views — guard upper bound
- **Root cause**: Custom scroll views like `DiffScreen` incremented their scroll offset `s.scroll` upon pressing key down without checking the total line bounds. This caused `rendered[s.scroll:]` to slice past the slice's actual length, leading to a Go runtime bounds panic.
- **Fix**: Always calculate `maxScroll = len(contentLines) - viewportHeight` and clamp the scroll index at `maxScroll` when handling key down and rendering contents.

### Unregistered slash commands — implement DialogScreen interface
- **Root cause**: Commands like `/settings` were listed in `/help` but not registered in `commands.go`. The underlying screen component (`SettingsScreen`) did not implement the `DialogScreen` interface (`Update`, `Done`, `Result` methods), making it unusable as an interactive dialog.
- **Fix**: Concurrently implement the full `DialogScreen` interface on any dialog screen, register its slash command handler, and display it with rounded borders (`styles.DialogBox`).

### Status bar animation — propagate tick commands and use rune-by-rune styling
- **Root cause**: The status bar's loading/thinking animation `[⬝⬝⬝⬝⬝⬝]` did not animate properly. This was due to two issues:
  1. The animation was rendered using raw byte slicing (`frame[:4]` and `frame[4:]`) on UTF-8 strings. Since `[` is 1 byte, and `■` and `⬝` are 3-byte UTF-8 runes, the slice `frame[:4]` always only styled the left bracket and the very first block, leaving the remaining blocks uncolored/stale.
  2. The `tea.Cmd` returned by `m.statusBar.Update(msg)` and `m.statusBar.SetWorking(true)` was discarded in `AppModel.Update()`, preventing the tick loop message (`workingTickMsg`) from scheduling future frames.
- **Fix**:
  1. Refactor `progressBar()` to convert the frame string to `[]rune` and style runes individually: render `■` with `barFilled` (yellow) and everything else (`⬝`, `[`, `]`) with `barEmpty` (dark gray).
  2. Modify `SetWorking(true)` to return a `tea.Cmd` from `b.tick()`, and ensure `AppModel.Update()` batches and returns the `tea.Cmd` returned by both `m.statusBar.SetWorking()` and `m.statusBar.Update(msg)`.

### Dialog resizing and event swallowing — intercept inputs only and dock status bar
- **Root cause**: Opening a dialog screen (implementing `DialogScreen`) caused layout and update anomalies:
  1. The `if m.activeDialog != nil` block in `AppModel.Update` intercepted all bubbletea messages indiscriminately. This swallowed critical global messages (like `WindowSizeMsg`, background `AnimTickMsg`, and streaming LLM token messages) preventing global resizing, background animation tick scheduling, and update handlers from working when dialogs were open.
  2. Dialog box rendering calculations used a fixed overhead offset that did not fill up to `m.height` lines, causing the status bar to float in the middle of the terminal rather than stay docked at the bottom.
- **Fix**:
  1. Handle `WindowSizeMsg` globally at the top of `AppModel.Update()` to update sizes and propagate resize events to both the main view and the active dialog.
  2. Restructure the active dialog interception in `AppModel.Update()` to only intercept user inputs (`KeyMsg` and `PasteMsg`), letting system events fall through.
  3. Ensure `layout()` calls `activeDialog.SetSize()` to keep dialog dimensions in sync with window bounds.
  4. Dynamically compute the top and bottom padding when a dialog is active inside `AppModel.View()`, centering the dialog vertically and appending `bottomPad` newlines so the status bar stays docked at the bottom.

### Dialog stack — ShowDialog from within a dialog's Update must not be overwritten
- **Root cause**: When `s.onEdit(p.Name)` inside `ProviderListScreen.Update` called `r.app.ShowDialog(editScreen)`, it set `m.activeDialog = editScreen`. But immediately after, the Update handler's `m.activeDialog = d` (where `d` is the return value of `m.activeDialog.Update(msg)`, i.e., the provider list) overwrote it. The edit screen was assigned but never rendered.
- **Fix**: Added `dialogUpdated` flag. `ShowDialog` sets it. The Update handler checks `if m.dialogUpdated { return m, nil }` before `m.activeDialog = d`, preserving the new dialog.
- **Dialog stack**: Added `dialogStack []screens.DialogScreen`. `ShowDialog` pushes the current dialog before replacing. `closeDialog` pops from the stack instead of setting `activeDialog = nil`. This lets nested dialogs (provider list → edit form) work naturally.
- **Stale data in stacked dialogs**: When a dialog is restored from the stack, its local data may be stale. Add an `OnRefresh` callback to the dialog screen that re-fetches data from the DB. Call it in `SetSize` (which is called when the dialog is restored and layout runs). The callback pattern: re-`List()` from service, rebuild items, call `SetProviders()`.

### Dialog text must be truncated to prevent layout overflow
- **Root cause**: `ProviderListScreen.View()` rendered long provider URLs (e.g., `https://generativelanguage.googleapis.com/v1beta/openai`) and names without length checks. Lines wider than the dialog box content area (`s.width - 2`) overflowed past the border, producing double-border artifacts (`││`).
- **Fix**: Truncate every rendered string field to `innerW - offset` using `[]rune` slicing, appending `"..."` when truncated. Always account for indentation, cursor prefixes, icons, and badges in the offset calculation.
- **Check every View() for untruncated raw strings**: Provider names, URLs, descriptions, and any dynamic text must be truncated before rendering. A single long line can break the entire dialog layout.


