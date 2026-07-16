---
name: external-drift-check
description: Use when repo docs, receipts, roadmaps, PRs, or status updates reference external systems such as Linear, GitHub, Slack, Figma, CI, Blockcell, deployed URLs, or calendars; verify live state or mark it posted, drafted, blocked, not attempted, or drifted.
---

# External Drift Check

Use this skill when local repo claims depend on external systems. The goal is to stop stale links and status mismatches from becoming project truth.

## Trigger Signals

- Linear issue IDs, PR numbers, Slack posts, published URLs, CI checks, Figma files, launch trackers, or docs are referenced.
- The user asks where things stand, what shipped, what is posted, or whether something is synced.
- A receipt says an external action happened or did not happen.

## Status Vocabulary

Use one of:

- `posted`: verified live in the external system.
- `drafted`: local draft exists but was not sent.
- `blocked`: attempted or required but blocked; include reason.
- `not_attempted`: no live action was tried.
- `drift`: repo/local claim disagrees with external state.
- `unknown`: cannot verify with available tools; say why.

## Workflow

1. Extract external references from docs, receipts, commits, or task text.
2. Use the available connector/CLI/tool when present.
3. If no tool exists, say the state is `unknown` or `not_attempted`; do not infer.
4. Update or draft a receipt/status note with exact evidence.
5. If sending on the user's behalf, follow the repo's message-prefix rule.

## Guardrail

Never convert a local draft into a posted claim unless the external system was actually contacted and verified.
