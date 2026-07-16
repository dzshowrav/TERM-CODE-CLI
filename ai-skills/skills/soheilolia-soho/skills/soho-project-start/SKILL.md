---
name: soho-project-start
description: Use as the Soho project-start/context-engineering skill when starting a new project, major spec, PRP, architecture brief, agent handoff, or workflow that must keep durable project markdown surfaces aligned, including AGENTS.md, CLAUDE.md, CODEX.md, README.md, ARCHITECTURE.md, PLANNING.md, TASK.md, lessons.md, docs/specs, docs/plans, docs/receipts, Linear, and issue-tracker or handoff notes.
---

# Soho Project Start

Use this with `using-soho` before substantial project/spec work. The goal is a context package that future agents can execute from, not a chat-only plan.

## Mode

- Use `recommend` when the project direction or ownership is still undecided.
- Use `solo` when one agent can inspect, write the context package, and verify it.
- Use `swarm` only when independent architecture, product, security, design, or implementation tracks genuinely need separate passes.

Record the mode and runtime in the receipt.

## 1. Orient

Before writing, inspect the existing project truth:

```bash
pwd
git status --short --branch
find . -maxdepth 3 -type f \( -name 'AGENTS.md' -o -name 'CLAUDE.md' -o -name 'CODEX.md' -o -name 'README.md' -o -name 'ARCHITECTURE.md' -o -name 'PLANNING.md' -o -name 'TASK.md' -o -name 'lessons.md' -o -path './docs/*' \) | sort
```

Also inspect local equivalents when present:

- `BUILDERBOT.md`, `GEMINI.md`, `GOOSE.md`, `.cursor/commands/`
- `.agentops/context/`, `.agentops/tasks/`, `.agentops/receipts/`
- `docs/specs/`, `docs/plans/`, `docs/receipts/`
- open PR, Linear, Slack, or handoff notes referenced by the user or repo

Do not create every possible file. Use existing conventions first. For a new empty project, the safe default is `README.md`, `AGENTS.md`, `docs/specs/`, `docs/plans/`, and `docs/receipts/`.

## 2. Build The Context Package

Borrow the useful Context Engineering pattern:

1. Initial request: what is being built and why.
2. Project rules: durable agent instructions and constraints.
3. Examples and patterns: files or prior artifacts to mirror.
4. PRP/spec: implementation blueprint with enough context for one-pass execution.
5. Validation gates: commands the agent can run and iterate on.
6. Receipt: what changed, what is proven, and what is not claimed.

The package must answer:

- Problem statement and user-visible outcome.
- Current truth layer: working tree, `HEAD`, generated artifacts, external systems.
- Non-goals and privacy/safety constraints.
- Source inventory and ownership.
- Existing patterns/examples to follow.
- Architecture decision and tradeoffs.
- Interfaces or artifact contracts.
- Manual curation or override path when generated output is involved.
- Tests, validation commands, and success criteria.
- Open questions with owner or next action.

## 2A. Orchestration Roadmap Gate

Before creating tickets or writing implementation code, turn the project into an orchestrated roadmap. The agent must be able to explain where the work is going, when the pieces converge, and what visible interface or workflow proves the project exists as a product.

Meta-level prompt to apply:

> We are not here to bang out disconnected tickets. Before implementation starts, define the end state, the visible product moment, and the orchestration path that turns separate tools into one usable workflow. Every milestone must have a tangible outcome a person can see, run, or decide from. Every ticket must map to a milestone, an interface/workflow checkpoint, and a truth layer: frontend/interface, backend/service, data/model, integration, docs/ops, or external decision. Maintain explicit Done, Todo, and Backlog states. If a task is only a fixture, prototype, generated artifact, or local proof, label it that way and do not let it masquerade as live readiness. If a backend or code dependency is required before the interface can come together, name it directly. Do not create a pile of tickets without a milestone map that shows how they converge into the actual experience.

Required roadmap surfaces:

- Milestone map with tangible outcomes.
- Done, Todo, and Backlog sections.
- Ticket-to-milestone mapping.
- Code/backend/interface/integration/doc/external-decision classification for each ticket.
- First visible product checkpoint.
- Orchestration checkpoint where separate tools become one workflow.
- Explicit non-claims for fixture-only or docs-only proof.

## 3. Update Markdown Surfaces In Lockstep

Update only the files that are present, expected by the repo, or necessary for the new project.

- `AGENTS.md`: repo-specific agent rules and required read order. Preserve managed blocks exactly.
- `CLAUDE.md` / `CODEX.md`: runtime-specific guidance only when the repo uses those files. Keep them aligned with `AGENTS.md` instead of creating conflicting rules.
- `README.md`: human entrypoint, run commands, and current status.
- `ARCHITECTURE.md` / `PLANNING.md`: system design, boundaries, source-of-truth decisions, and sequencing.
- `TASK.md`: active tasks, discovered work, and dated next actions.
- `lessons.md`: durable lessons and repeated failure modes only. Do not use it as a scratchpad.
- `docs/specs/`: product and technical spec.
- `docs/plans/`: executable implementation plan with exact paths and validation.
- `docs/receipts/`: proof, commands run, outputs, blockers, non-claims, and external update state.
- Issue tracker / Linear / PR notes: post only when explicitly authorized; otherwise draft exact text and label it not posted.

If two files would say the same thing, prefer one canonical file and add a pointer from the other. Avoid duplicate narratives that can drift.

## 4. PRP / Spec Minimum

A build-ready PRP or spec needs:

- Goal, why, and user-visible behavior.
- Current codebase/project tree and desired tree.
- Files to create or modify.
- Data models, interfaces, or artifact schemas.
- Implementation steps in execution order.
- Validation gates that are actually runnable.
- Risks, failure modes, and mitigations.
- Success metrics or acceptance criteria.
- Explicit fixture/local proof versus live integration proof.

For agent-facing PRPs, include exact files and examples to read. Do not rely on a model guessing the project conventions.

## 5. Validation

Before claiming the context package is ready:

- Re-read the changed markdown files for contradictions.
- Verify JSON/YAML receipts or schemas parse when applicable.
- Verify referenced files exist or are explicitly marked future/nonexistent.
- Verify validation commands are concrete, not placeholders.
- Re-run `git status --short --branch` and classify unrelated dirty work.
- Record any live-system action as `posted`, `drafted`, `blocked`, or `not attempted`.

## 6. Receipt

End with a receipt containing:

- mode and runtime
- files created or updated
- commands run and results
- context-engineering inputs used
- live systems contacted or explicitly not contacted
- what is proven now
- what is not claimed
- next recommended action

## Guardrails

- Preserve managed blocks and user-owned edits.
- Do not overwrite existing project doctrine with generic templates.
- Do not claim fixture proof as live integration proof.
- Do not create `CLAUDE.md`, `CODEX.md`, or `lessons.md` just to satisfy a checklist if the repo has another canonical context surface.
- Do not post Slack, Linear, email, or PR comments without explicit authorization.
- If sending on the user's behalf after authorization, prefix the message with `🤖`.
