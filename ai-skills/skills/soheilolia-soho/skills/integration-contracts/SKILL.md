---
name: integration-contracts
description: Use before wiring external systems such as Slack, Figma, GitHub, Builderbot, Google Drive, Linear, Blockcell, CI, or APIs; define collector contracts, fixture importers, source health, degradation, permissions, and snapshot interfaces first.
---

# Integration Contracts

Use this skill before live connectors. The goal is to wire external systems through a stable seam, not scatter API assumptions through the codebase.

## Trigger Signals

- "Wire", "integrate", "collector", "connector", "API", "Slack", "Figma", "GitHub", "Builderbot", "Google Drive", "Blockcell", "CI", "publish", or "sync".
- A task touches 3+ integration-adjacent files.
- A live system will feed a local model, dashboard, digest, or bot.

## Contract First

Define before implementation:

- Input source and access method.
- Allowlist and permission boundary.
- Output schema or normalized entity shape.
- Source metadata fields.
- Error and partial-failure shape.
- Source health reporting.
- Fixture importer matching the live contract.
- Degradation behavior when a source is unavailable.
- Test fixtures for success, empty, partial, and failed source states.

## Implementation Rule

Live API code must feed the same contract as fixtures. Do not let each collector invent a new data shape.

## Validation

At minimum, test:

- valid source payload
- missing/empty source
- malformed source
- restricted or disallowed source
- source failure with degraded output

## Guardrail

No network call should be required to prove the contract. Use fixtures first, then live wiring once the seam is stable.
