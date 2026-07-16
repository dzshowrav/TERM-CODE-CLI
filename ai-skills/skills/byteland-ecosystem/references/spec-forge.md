# spec-forge — Reference

> **GitHub**: https://github.com/ByteLandTechnology/spec-forge
> **Stars**: 0 (new) | **License**: MIT | **Status**: v1.0.0 (3 commits)
> **Languages**: Rust 83.3%, JavaScript 15.6%, Shell 1.1%

A **YAML-first workflow** for turning a rough request into an approved implementation spec and an auditable implementation report. Combines agent-guided stage progression with a Rust CLI that persists durable state under `.spec-forge/`.

---

## When to Use

- Turn a vague feature request into a structured, reviewable spec
- Persist planning state in versionable files (YAML) instead of chat memory
- Pause and resume multi-stage spec work safely
- Require explicit approvals before stage advancement

---

## Architecture: 6 Stages (P0-P5)

| Stage | Purpose |
|-------|---------|
| **P0 — Intake** | Frame the request, stakeholders, scope, and constraints |
| **P1 — Architecture** | Lock solution outline + journey and component indexes |
| **P2 — Journeys** | Refine in-scope journeys in reviewable batches |
| **P3 — Components** | Refine in-scope components into implementation-facing contracts |
| **P4 — Readiness** | Consolidate approved work into final implementation spec |
| **P5 — Implement** | Record delivery status, validations, blockers, and closure |

---

## Workspace Model

Persistent state lives under `.spec-forge/`:

```
.spec-forge/
├── registry.yaml
└── specs/
    └── <spec-id>/
        ├── pipeline-state.yaml        ← stage, focus, UX handoff
        ├── handoff.yaml
        ├── framing/
        ├── architecture/
        ├── journeys/
        ├── components/
        ├── synthesis/
        │   ├── implementation-spec.yaml    ← implementation-ready contract
        │   └── implementation-report.yaml  ← delivery status & validations
        └── gates/
            └── <stage>.yaml               ← gate evaluation per stage
```

---

## CLI Commands (`spec-forge-cli`)

```bash
# Initialize a workspace
spec-forge-cli init --target . --spec-id demo --request-title "Demo Spec"

# Resolve a stage
spec-forge-cli resolve --target . --spec-id demo --skill spec-forge --stage router --write

# Validate UX
spec-forge-cli ux validate --target .
```

---

## Agent Entry Points

```markdown
Start: "Use $spec-forge to turn this request into an approved spec."
Resume: "Use $spec-forge in this repo and continue checkout-redesign."
Find next: "Use $spec-forge to find the next incomplete stage and keep going."
```

---

## Skill Family Structure

| Path | Purpose |
|------|---------|
| `spec-forge/` | Router skill + shared UX contracts |
| `spec-forge-intake/` | P0 stage skill |
| `spec-forge-architecture/` | P1 stage skill |
| `spec-forge-journeys/` | P2 stage skill |
| `spec-forge-components/` | P3 stage skill |
| `spec-forge-readiness/` | P4 stage skill |
| `spec-forge-implement/` | P5 stage skill |
| `spec-forge-cli/` | Rust CLI, npm wrapper, release scripts |
