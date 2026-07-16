---
name: phase-boundary-hardening
description: Use before moving from spec, fixture, prototype, or local proof into live integrations, production publishing, CI, automation, or broad rollout; checks readiness across receipts, degradation, versioning, permissions, drift, rollback, and ownership.
---

# Phase Boundary Hardening

Use this skill when a project is about to cross from one truth mode to another, especially fixture-to-live or local-to-published.

## Trigger Signals

- "Start live integration", "wire Slack/Figma/GitHub", "publish", "launch", "production", "CI", "cron", "automation", "Blockcell", or "real data".
- A fixture dashboard, mock API, generated artifact, or PRP is about to become a live system.
- Multiple agents are about to work in parallel on integration code.

## Readiness Checklist

Before crossing the boundary, verify or create explicit follow-ups for:

- Source inventory and owners.
- Permission and requester-access model.
- Collector/input contracts.
- Graceful degradation and source health.
- Retry/failure behavior.
- Snapshot or artifact versioning.
- Receipt schema and claim validation.
- External drift check for issue trackers, docs, CI, or published URLs.
- Rollback or disable path.
- Manual correction or curator path.

## Output

Produce a short phase-boundary decision:

- `ready`: local proof and gates exist; live work may start.
- `ready_with_conditions`: live work may start only under named constraints.
- `not_ready`: blocking gap and owner/action are named.

Include exact validation commands and non-claims.

## Guardrail

Do not move live until the boundary is explicit. If the next action needs an owner decision, finish all local scaffolding and leave the owner action as the only blocker.
