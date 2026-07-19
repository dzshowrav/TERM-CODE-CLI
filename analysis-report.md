# TermCode Full App Analysis Report

Date: 2026-07-19
Scope: Complete codebase audit by 5 specialized agents (init flow, commands, chat/stream, dialog/layout, persistence/DB)
Total issues found: **51**

---

## CRITICAL (crashes, data loss, unrecoverable)

| # | Bug | File:Line |
|---|-----|-----------|
| C1 | **`deleteMessage()` never deletes from DB** — message row stays in `messages` table forever. On reload, "deleted" messages reappear. | `app.go:1242-1258` |
| C2 | **`compressContext()` only truncates in-memory** — DB has all original rows. On reload, full history returns and the summary message is gone. | `app.go:1161-1179` |
| C3 | **Undo/redo never writes to DB** — all undo operations lost on restart. `delete_message` undo reinserts into `m.history` but DB row was never deleted. | `app.go:1260-1283` |
| C4 | **`switchToBranch()` writes new messages to original session ID** — branch history diverges from DB permanently. Messages written from a branch go to the wrong session. | `app.go:1181-1195` |
| C5 | **StatePaused has no escape** — every keypress is swallowed, `CommandInput.Update` never called. User cannot type `/continue`, `/retry`, or any command. Only Ctrl+C (app quit) works. | `app.go:844-854` |
| C6 | **`sessionRepo.Create()` error silently ignored** in `ensureSession()` — `m.currentSess` points to a session not in DB. All subsequent `messageRepo.Create()` calls fail with FK violation. Messages silently disappear forever. | `app.go:400` |
| C7 | **19 DB errors silently ignored** — `Create`, `Update`, `Delete` return values discarded across all repos. User never knows when persistence fails. | `app.go:400,415,421,1235,1237,1254` + many more |
| C8 | **Stale tool events from cancelled goroutine** — after `/cancel` + new SubmitMsg, old goroutine's `prog.Send(ToolQueuedMsg{})` arrives and `append`s phantom tool cards into the new conversation. | `app.go:920-935` + `chat_screen.go:559-573` |
| C9 | **No signal handler** — `tea.WithoutSignalHandler()` means SIGTERM/SIGHUP from Termux process management kills instantly without restoring alt screen. Terminal left broken (cursor invisible). | `main.go:128` |
| C10 | **Byte/rune mismatch in cursor position** — `cursor` is byte offset but `utf8.RuneCountInString` slices by byte index. Bangla/CJK/emoji input corrupts cursor display. | `command_input.go:161-162` |

---

## HIGH (wrong behavior, UX broken, data inconsistency)

| # | Bug | File:Line |
|---|-----|-----------|
| H1 | **Pause state overwritten by `streamDoneMsg`** — Esc sets `StatePaused`, but the goroutine returns `context.Canceled` error, `streamDoneMsg` handler overwrites to `StateCancelled`. State indicator shows "cancelled" (red) instead of "paused" (yellow). | `app.go:951-963` |
| H2 | **Visual flicker during streaming** — `AppendResponseText` adds raw text to viewport, `renderMessages()` (called 10x/sec by tick) rebuilds with markdown. User sees raw text snap to formatted text on every chunk. | `chat_screen.go:237-247` + `chat_screen.go:375-544` |
| H3 | **`SetViewportSize` ignores `scrolledAway`** — unconditionally calls `GotoBottom()` on resize. When Termux keyboard opens/closes, user loses scroll position. | `chat_screen.go:142` |
| H4 | **`cmdContinue` creates a NEW session** — sets `currentSess = nil` then `ensureSession()`. Partial response goes to a different session than the original conversation. | `commands.go:767-769` |
| H5 | **7 command handlers will panic if services not initialized** — `cmdTools` (chatSvc), `cmdStop` (eventBus), `cmdContinue/Retry/Cancel` (chatScreen), `cmdImport` (messageRepo), `cmdCollab` (chatScreen). | `commands.go:401,759,771,651,666,892,955` |
| H6 | **Dialog stack scope asymmetry** — nested dialogs misalign keybinding scope stack. After enough nesting, `PopScope` no-ops and `CurrentScope()` returns wrong scope. | `app.go:551-564` |
| H7 | **`layout()` doesn't account for toast/gauge/state lines** — rendered output exceeds terminal height by 2-5 lines when toast is visible or state is non-idle. Bottom content clips off-screen. | `app.go:606-607` |
| H8 | **`ToolConfirmDialog` navigation broken** — `shift+tab` and `left` both increment focus instead of decrementing. | `tool_confirm_dialog.go:50-51` |
| H9 | **`WorkspaceTrustDialog` defaults to Trust** — `focusField: 1` means pressing Enter trusts by default. Security anti-pattern. | `workspace_trust_dialog.go:26,53` |
| H10 | **`MessageCnt` permanently wrong after compression** — `len(m.history)` written to DB after in-memory truncation. On reload, `MessageCnt` shows compressed count but all messages are restored. | `app.go:419` |
| H11 | **7+ dialog screens missing `PasteMsg`** — pasted text silently dropped in session_screen, batch_edit_screen, git_add_screen, message_screen, settings_screen, agent_select_screen, model_list_screen. | Multiple `screens/*.go` |
| H12 | **No FK on `conversations.session_id`** — orphan conversation rows survive session deletion. CASCADE never fires. | `migrations.go:99-109` |
| H13 | **Missing indexes on 6 key query columns** — `sessions.updated_at`, `sessions.status`, `models.provider_id`, `workspaces.path`, `agents.is_default`, `providers.is_default`. Full table scans on every list query. | `migrations.go` |
| H14 | **`/plugin` command not listed in `/help`** — command exists but is invisible to users. | `help_screen.go` |
| H15 | **`/collab start` goroutine leak** — calling `/collab start` twice creates a second goroutine while old server keeps running (no stop). | `commands.go:922-927` |
| H16 | **`sessionRepo.Update` can nil-panic** in `editMessage` — `m.sessionRepo.Update(m.ctx, m.currentSess)` without nil guard on `sessionRepo`. | `app.go:1237` |

---

## MEDIUM (degradation, waste, inconsistency)

| # | Bug | File:Line |
|---|-----|-----------|
| M1 | **`streamBuf` unbounded** — grows to multi-MB for long responses. Only freed on next submit. | `app.go:940` |
| M2 | **`ThinkingTickMsg` runs 10x/sec between conversations** — `responseText != ""` keeps it alive even when user is reading (no animation visible). | `chat_screen.go:637-649` |
| M3 | **`renderMessages()` called from 18+ sites** — each call rebuilds entire viewport O(n*m) where n=messages, m=content. 10 calls/sec during streaming. | `chat_screen.go:159-543` |
| M4 | **`history` grows unbounded** — `compressContext()` only fires after ~384K chars. Until then, every message retained in memory. | `app.go:412` |
| M5 | **`cmdProvider` OnSelect not persisted** — selecting a default provider is lost on restart. | `commands.go:283-285` |
| M6 | **Dialog stack unbounded** — can grow to arbitrary size by nesting ProviderList→Edit→ProviderList→Edit... | `app.go:539-541` |
| M7 | **3 screens use `Content()` instead of `DialogBox()`** — renders with plain padding instead of rounded borders. Inconsistent with all other dialogs. | `model_list_screen.go:114`, `model_selector.go:166`, `tool_execution_screen.go:118` |
| M8 | **`ProviderAddScreen` Enter on input submits immediately** — pressing Enter on Name field submits form with only name filled in. | `provider_add_screen.go:72-86` |
| M9 | **`ModelSelector` no scroll bounds** — renders all models regardless of terminal height, overflows screen. | `model_selector.go:55-92` |
| M10 | **7 missing transactions** — `saveMessage`, `editMessage`, session deletion, `SyncFromProvider`, etc. Non-atomic multi-write operations can leave DB inconsistent on crash. | Multiple locations |
| M11 | **Content/reasoning ordering flipped per chunk** — `streamContentMsg` sent before `streamReasoningMsg` for same LLM chunk. Content displayed before its reasoning. | `chat_service.go:370-379` |
| M12 | **`cmdStop` vs `cmdCancel` inconsistency** — one sets `cancel = nil`, other doesn't. | `commands.go:752 vs 663` |
| M13 | **Anim tick never restarts** — leaves home screen, tick stops, never restarts when returning home. Particles frozen. | `app.go:773-779` |
| M14 | **`onToolEvent` closure captures `m` — `eventBus.Emit` from goroutine** — concurrent `Emit` from goroutine + main loop unsynchronized. | `app.go:478-507` |
| M15 | **`dialogs/` package is 1650+ lines of dead code** — all 12 files have zero imports anywhere. Old dialog pattern replaced by `screens/`. | `internal/adapters/tui/dialogs/*` |

---

## LOW (cosmetic, minor edge cases)

| # | Bug | File:Line |
|---|-----|-----------|
| L1 | `var _ = fmt.Sprintf` dead keep-alive — `fmt` is actually used, so this is unnecessary | `app.go:1308` |
| L2 | `ready = true` hardcoded + dead "Initializing..." branch in View() | `app.go:207,1043-1048` |
| L3 | `/tmp/termcode.log` unbounded, no rotation, may not exist on Termux | `main.go:139` |
| L4 | `cmdClear` doesn't call `layout()` — home screen might misrender | `commands.go:611-618` |
| L5 | No quoted argument support — `/import "/path/with spaces/file"` breaks | `commands.go:100` |
| L6 | Migration v2 is irreversible (ALTER TABLE ADD COLUMN) | `migrations.go:250` |
| L7 | `ModelSelector` dead code — no scroll bounds, may overflow | `model_selector.go` |
| L8 | 2 screens pass `s.width-4` to DialogBox instead of `s.width` — box 4 cols too narrow | `network_screen.go:67`, `about_screen.go:44` |
| L9 | `ProviderEditScreen` title shows stale `origName` — never updated after edit | `provider_edit_screen.go:21,148` |
| L10 | `conversations`, `bookmarks`, `attachments`, `themes` tables created but never wired to any UI or service | `migrations.go:99-231` |

---

## Summary

| Severity | Count |
|----------|-------|
| CRITICAL | 10 |
| HIGH | 16 |
| MEDIUM | 15 |
| LOW | 10 |
| **Total** | **51** |

### Most impactful

**C1-C4** — In-memory/DB divergence. Deletions, compression, undo, and branches all silently lose or corrupt data on restart. Every `deleteMessage()` call is a no-op that resets on reload. Every undo is ephemeral. Every branch write goes to the wrong session. Every compression is invisible on reload.

**C5** — StatePaused dead-end. Once paused, the user cannot type any command. Only Ctrl+C (which quits the app) works.

**C7** — 19 ignored DB errors. Persistence failures are completely invisible to the user. Messages, sessions, and settings can silently fail to save.

**H7** — Layout overflow. Toast notifications, context gauge, and state line push content past the terminal bottom. On mobile (Termux keyboard), this is especially severe.
