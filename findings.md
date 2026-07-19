# Full Project Sync Research Report

## CRITICAL BUGS (5)

| # | Bug | Location | Impact |
|---|-----|----------|--------|
| C1 | **`m.state` mutated from goroutine without mutex** | app.go:441-463,884-912 | Data race — `m.state` written by `onToolEvent` goroutine AND `AppModel.Update` simultaneously. Bubble Tea reads `m.state` in `View()` while goroutine writes it. |
| C2 | **`persistToolResults()` dead code — tool results lost** | chat_screen.go:279 (defined, never called) | Tool cards are ephemeral. After next user message, all tool output disappears from display. Only survives in viewport scrollback until cleared. |
| C3 | **`cost.Engine` completely disconnected** | chat_service.go:57,78; converseStream:377 | `SetCostEngine` never called. Budget warnings, cost tracking, token accounting all dead code. Full implementation, zero integration. |
| C4 | **Encryption key derived from hostname — not portable** | crypto.go:27-77 | Moving DB between devices (or hostname change on mobile) makes all API keys undecryptable. No master key, no fallback. |
| C5 | **`m.streamBuf` / `s.responseText` double buffer — can diverge** | app.go:947, chat_screen.go:206 | Both hold same content independently. If `/continue` or pause/cancel happens mid-stream, partialContent and responseText can diverge showing different text. |

## HIGH SEVERITY (7)

| # | Issue | Location |
|---|-------|----------|
| H1 | **Tick chain stops when reasoning arrives — no re-renders during streaming** | chat_screen.go:599-612 | `ThinkingTickMsg` checks `showThinking||thinkingDur>0||hasActiveToolCards()` but NOT `streamActive`. When reasoning comes before content, tick stops, viewport doesn't update until streamDoneMsg. |
| H2 | **Branches entirely in-memory — lost on restart** | app.go:142-150 | `ConversationBranch` stored in `[]ConversationBranch`. Never persisted to DB. Auto-save overwrites current branch but DB has no branch concept. |
| H3 | **Undo/redo not synced to DB — orphaned message rows** | app.go:1199-1221 | `editMessage()` creates NEW DB row instead of UPDATE. Old row stays as orphan. `deleteMessage()` removes from memory only — DB row persists. |
| H4 | **Session MessageCnt desync** | app.go:381-385 | `saveMessage()` updates count, but `editMessage()`/`deleteMessage()` don't. DB count drifts from actual messages over time. |
| H5 | **`SetOutputCallback` unprotected — data race risk** | tool_service.go:72-74,299-301 | `outputCallback` field is bare `func(string)`. Currently safe (serial tool calls), but unsafe by design for concurrent use. |
| H6 | **`m.ctx` cancel function leak** | app.go:869,988 | Old `m.cancel` not called before reassignment. Context leaks if rapid submit. Only set to nil after stream completion — not on pause/cancel. |
| H7 | **`ToolTrustManager` dead code — trust decisions never persisted** | trust_manager.go (entire file) | `NewTrustManager` never called. Permission decisions stored in-memory only. Restart loses all "always allow" decisions. |

## MEDIUM SEVERITY (10)

| # | Issue | Location |
|---|-------|----------|
| M1 | **Git branch detected once at startup — stale** | app.go:294 | `detectGitBranch()` only called in `SetProviderService()`. `/git checkout` doesn't update `m.gitBranch` or `statusBar.SetBranch()`. |
| M2 | **Duplicate goroutine logic in SubmitMsg and startChat** | app.go:862-942 vs 408-495 | ~50 lines of identical code duplicated. `startChat()` used by `cmdRetry`, but `SubmitMsg` handler has inline copy. |
| M3 | **ChatScreen rebuilt on every undo/redo/edit — expensive + loses state** | app.go:1177,1216,1233,1260,1285 | Creates fresh `NewChatScreen()`, re-adds all messages. Resets viewport scroll, markdown cache, tool card state. O(n) per operation. |
| M4 | **`SetStreamActive(true)` never called** | chat_screen.go:254 | `streamActive` always `false`. `cachedBase` (line 500-501) allocated but never read. Wasteful allocation on every streaming render. |
| M5 | **`retry.Do` return value discarded** | tool_service.go:286 | Retry error info lost when all attempts exhausted. Only last `result.Error` survives. |
| M6 | **`OnStart` hook fires on every retry attempt** | tool_service.go:293-295 | Duplicate UI events if OnStart sends "tool started" message. Output callback also fires per attempt with potentially partial output. |
| M7 | **Dialog screens hold stale data (SessionScreen, ModelListScreen)** | session_screen.go, model_list_screen.go | Local slices not refreshed after mutations inside dialog. Deletes don't update displayed list until dialog is closed and reopened. |
| M8 | **Provider selection from dialog doesn't call layout()** | app.go:718-720 (implied) | `screen = screenChat` set but `layout()` not called after provider select. Dimensions may not recalculate. |
| M9 | **`m.ctx` not canceled on stream completion** | app.go:954-994 | `m.cancel` never called when stream finishes naturally (no error). Cancel function only called on explicit `/stop` or `CancelStreamMsg`. |
| M10 | **Session `provider_id`/`model_id` have no FK constraints** | migrations.go:53-54 | Dangling references if provider/model deleted. No validation on session write. |

## LOW SEVERITY (12)

| # | Issue | Location |
|---|-------|----------|
| L1 | `StateThinking` defined but never used | app.go:40 | Dead enum value in state machine |
| L2 | Two parallel animation tick chains | app.go:599-601,734-735 | `AnimTickMsg` from both `Init()` and `Update()` handler — redundant |
| L3 | `No tea.PasteMsg` handler in AppModel | app.go:738-806 | Works only because commandInput catches it via default fallthrough |
| L4 | Checkpoint path never set | chat_service.go:94-131 | `SaveCheckpoint()` is no-op — can't load/save without path |
| L5 | 28/35 event bus events dead | eventbus/events.go | Only 7 emitted, 5 subscribed. Rest are enum constants unused |
| L6 | `stream` package dead code | internal/application/stream/ | `NewHandler`/`NewTracker` never called. Duplicates `converseStream` |
| L7 | `session.Service` application layer unused | session_service.go | All session ops go directly through SQLite repos, bypassing service layer |
| L8 | `ErrCannotDelete`, `ErrNoMessages`, `ErrEmptyAPIKey` defined never used | Various errors.go | Dead error constants |
| L9 | Settings screen silently ignores List errors | commands.go:584 | `if err == nil { ... }` — no error displayed |
| L10 | `"allow_once"` falls through switch case silently | tool_service.go:129-139 | Works correctly but confusing — no explicit case |
| L11 | `SessionScreen` delete modifies local slice but closure captures stale `providersCopy` | commands.go:230-245 | Index may shift after deletion |
| L12 | `converseStream()` has no timeout — stream can hang forever | chat_service.go:335-412 | No `context.WithTimeout` wrapper |

## ARCHITECTURAL CONCERNS (6)

| # | Concern | Details |
|---|---------|---------|
| A1 | **Direct repo access from UI layer** | `AppModel` calls `sessionRepo`, `messageRepo`, `settingsRepo`, `modelRepo` directly instead of going through service layer. Bypasses all application logic, transactions, validation. |
| A2 | **No dependency injection container** | Services instantiated manually in `SetProviderService()` (app.go:234-300). No interface-based DI — concrete types everywhere. Hard to test, hard to swap implementations. |
| A3 | **In-memory state is source of truth, not DB** | `m.history`, `m.branches`, `m.currentSess` are primary. DB is written to asynchronously. On crash, in-memory state after last DB write is lost. |
| A4 | **`m.history` shared between AppModel and goroutine** | Copy made at line 426 but original modified by `saveMessage()` while goroutine runs. Data race on the slice elements. |
| A5 | **Dialog stack can grow unbounded** | `dialogStack` has no max size. Nested dialogs push but if close path is missed, stack accumulates. |
| A6 | **`keyMgr.Resolve()` result discarded** | app.go:805 | Key binding resolution executed but never acted upon. Decorative only. |

## POSITIVE FINDINGS (what works correctly)

Good news — these systems are solid:

1. **Tool execution pipeline** — Permission check → validate → execute → retry → hooks → truncate → return. All wired end-to-end.
2. **Streaming message flow** — `LLM adapter → converseStream → onChunk → prog.Send → AppModel → ChatScreen.AppendResponseText`. Clean unidirectional data flow.
3. **Tool event UI wiring** — `onToolEvent → 6 message types → ChatScreen.Update handles all 6`. No gaps.
4. **Crash recovery checkpoint** — Atomic tmp+rename, called after tool rounds AND final completion. (Just needs `SetCheckpointPath` call.)
5. **Session DB cascade** — `ON DELETE CASCADE` from messages to sessions. No orphaned messages at DB level.
6. **Provider delete cascades to models** — Both DB-level CASCADE and app-level double-delete. Robust.
7. **Dialog layout** — `SetSize` called on all dialogs through `layout()`. Responsive to resize.
8. **Event bus thread safety** — Bus is structurally thread-safe with RWMutex. Suitable for toast notifications.
9. **Thinking animation** — Self-contained tick chain, properly starts/stops, handles tool card animations.
10. **Help screen** — Full 34 commands listed with search/filter/scroll.

## RECOMMENDED FIX ORDER

Priority-based fix plan:

### Phase 1 (Critical — crashing/data loss)
1. Add mutex on `m.state` — prevent data race between goroutine and Update
2. Call `persistToolResults()` in `streamDoneMsg` handler — tool results survive next message
3. Wire `costEngine` — call `chatSvc.SetCostEngine(cost.New())` in `SetProviderService`
4. Fix `m.ctx` lifecycle — cancel old ctx before creating new, cancel on stream complete

### Phase 2 (High — correctness)
5. Fix ThinkingTickMsg condition — add `streamActive` check to keep tick alive during streaming
6. Fix undo/redo DB sync — use UPDATE not CREATE for message edits, delete from DB
7. Fix message count desync — recalculate from actual messages on edit/delete
8. Wire `SetCheckpointPath` — pass a real path from AppModel to chat service
9. Wire `ToolTrustManager` — instantiate in `SetProviderService`, pass settings repo

### Phase 3 (Medium — UX)
10. Refresh git branch on `cmdGit` checkout operations
11. Refresh dialog local lists after in-dialog mutations
12. Add `converseStream` timeout wrapping with context.WithTimeout
13. Remove dead code: `stream` package, dead event bus types, `StateThinking`, `cachedBase`

### Phase 4 (Low — polish)
14. Consolidate duplicate goroutine logic (SubmitMsg ↔ startChat)
15. Clean up event bus — remove dead events, wire useful ones
16. Add encryption key file as fallback for hostname-based derivation
