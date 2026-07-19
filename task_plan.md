# Task Plan: Full Project UI/UX ↔ Backend Sync & Bug Fix

## Goal
Fix every synchronization gap between UI/UX and backend — eliminate data races, persist ephemeral state, sync in-memory state with DB, wire disconnected components.

## Current Phase
Phase 1 (Research complete — report delivered)

## Phases

### Phase 1: Deep Research & Report
- [x] Map every UI component to its backend service
- [x] Trace every message type from emission to consumption
- [x] Find all synchronization gaps
- [x] Identify stale state / orphaned references
- [x] Detect callback chains that can break
- [x] Compile findings report
- **Status:** complete

### Phase 2: Fix Critical Bugs (C1-C5)
- [ ] Add mutex protection on `m.state` — data race fix
- [ ] Wire `persistToolResults()` — tool results survive next message
- [ ] Wire `costEngine` — connect budget optimizer
- [ ] Fix `m.ctx` lifecycle — cancel on completion, clean up on reassignment
- **Status:** pending

### Phase 3: Fix High Issues (H1-H7)
- [ ] Fix ThinkingTickMsg condition — keep alive during streaming
- [ ] Fix undo/redo DB sync — UPDATE not CREATE, delete from DB
- [ ] Fix message count desync — recalculate on edit/delete
- [ ] Wire `SetCheckpointPath` — pass real path
- [ ] Wire `ToolTrustManager` — instantiate at startup
- **Status:** pending

### Phase 4: Medium Issues (M1-M10)
- [ ] Refresh git branch on checkout operations
- [ ] Refresh dialog local lists after in-dialog mutations
- [ ] Add converseStream timeout
- [ ] Remove dead code
- **Status:** pending

### Phase 5: Polish (L1-L12, A1-A6)
- [ ] Consolidate duplicate logic, clean dead code
- [ ] Wire encryption key file fallback
- **Status:** pending

## Key Questions
1. Should branches be persisted to DB? → Yes, needs new `branches` table
2. Should we remove dead code or keep for future? → Remove dead, keep documented

## Decisions Made
| Decision | Rationale |
|----------|-----------|
| Report first, fix later | User explicitly asked for deep research before changes |
| Priority: Critical → High → Medium → Low | Safety and correctness before polish |
