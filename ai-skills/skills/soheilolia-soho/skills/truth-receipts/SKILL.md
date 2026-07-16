---
name: truth-receipts
description: Use when a task relies on receipts, claim ledgers, generated artifacts, status reports, proof boundaries, or external-system claims; enforce schema, truth layers, evidence, non-claims, and receipt validation before completion.
---

# Truth Receipts

Use this skill when work could be misreported as more complete than it is. The goal is to make receipts machine-checkable and human-trustworthy.

## Trigger Signals

- The task creates or updates `docs/receipts/`, a claims ledger, status report, roadmap, launch note, PR summary, or weekly update.
- Generated artifacts could be confused with live or published systems.
- External systems are mentioned: Linear, Slack, GitHub, Figma, Blockcell, CI, production, dashboards, docs, or calendars.
- The user asks for proof, receipts, "what is real", "show receipts", or "be honest".

## Required Truth Layers

Separate these explicitly when relevant:

- `HEAD`: committed code/docs at the current git commit.
- `working_tree`: uncommitted local changes.
- `generated_artifact`: files produced by commands, especially ignored outputs.
- `receipt`: what the receipt claims and when it was written.
- `external_system`: Linear/Slack/GitHub/Figma/CI/published state.
- `not_claimed`: what was not done, not verified, or not live.

## Receipt Minimum

A receipt must include:

- generated date/time
- host, mode, and runtime when Soho was used
- task or artifact summary
- files created/updated or outputs produced
- commands/evidence and results
- truth layers
- live systems contacted or explicitly not contacted
- blockers/risks
- not-claimed list
- confidence/status

## Workflow

1. Inspect existing receipt conventions before adding a new shape.
2. Prefer a schema or validator if the repo already has one.
3. If no schema exists, add a follow-up or implement the smallest validator in scope.
4. Run JSON/YAML parsing checks for structured receipts.
5. Verify generated artifacts from source commands, not manual edits.
6. Before final status, compare the receipt claims to git status, test output, and external-system state.

## Guardrail

Never let a local fixture, generated file, draft, or unposted note masquerade as live readiness.
