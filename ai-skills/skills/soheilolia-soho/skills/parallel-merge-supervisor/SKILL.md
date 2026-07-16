---
name: parallel-merge-supervisor
description: Use when multiple agents, sessions, branches, worktrees, or tools are working on one repo; define write ownership, hot files, merge order, supervisor checks, final test gate, and integration receipts.
---

# Parallel Merge Supervisor

Use this skill before or during multi-agent work. The goal is velocity without merge chaos or unverified integration.

## Trigger Signals

- The user mentions parallel agents, sidecars, Claude/Codex/Goose/Nexus, huddle, swarm, lanes, or multiple sessions.
- More than one branch/worktree/session may edit the same repo.
- A plan has "lanes" but no final integration owner.

## Required Plan

Define:

- Supervisor agent/session.
- Active write-agent cap.
- Lane owners and disjoint write sets.
- Hot files that require serialization.
- Merge order.
- Required receipt per lane.
- Final full-suite validation after all lanes land.
- External-system update policy: posted, drafted, blocked, or not attempted.

## Merge Gate

No lane is done until the supervisor verifies:

- The lane stayed within its write set or documented an approved exception.
- Tests for the lane and the full project pass after integration.
- `git diff --check` passes when applicable.
- Generated artifacts were created from commands, not hand-edited.
- Receipts name files, commands, evidence, non-claims, and external-system status.
- Roadmap/TASK/Linear status does not drift.

## Hot File Rule

If two lanes touch the same hot file, stop parallel writes and serialize the merge.

## Guardrail

Do not claim "swarm" or "parallel" unless the host actually ran independent agents. If work was serial role-play, record `runtime: prompt-backed`.
