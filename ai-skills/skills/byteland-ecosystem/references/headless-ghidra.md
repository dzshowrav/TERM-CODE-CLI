# headless-ghidra — Reference

> **GitHub**: https://github.com/ByteLandTechnology/headless-ghidra
> **Stars**: ⭐6 | **License**: MIT | **Status**: v1.8.0 (18 releases, 59 commits)
> **Languages**: Rust 55%, Java 28.8%, JavaScript 8.8%, Python 5.7%, Shell 1.7%

A Ghidra reverse-engineering **skill family** for AI agents. Provides reproducible, evidence-backed workflows with audit-ready Markdown outputs.

---

## Architecture (5 Phases + 1 Analysis)

| Phase | Skill Directory | Purpose |
|-------|----------------|---------|
| **P0 — Intake** | `headless-ghidra-intake` | Confirm target, initialize workspace, set scope |
| **P1 — Baseline** | `headless-ghidra-baseline` | Import into Ghidra, export baseline artifacts, record runtime |
| **P2 — Evidence** | `headless-ghidra-evidence` | Identify third-party code and evidence sources |
| **P3 — Discovery** | `headless-ghidra-discovery` | Enrich names, signatures, types, constants, strings |
| **P4 — Batch Decompile** | `headless-ghidra-batch-decompile` | Apply metadata and decompile selected functions |

Plus:
- **`headless-ghidra-analyze-function`** — Deep analysis of a single function (types → constants → vtables → identity → decompilation)
- **`headless-ghidra` (router)** — Orchestration layer that coordinates phase progression
- **`ghidra-agent-cli`** — Rust CLI invoked by skills (not user-facing)

---

## Installation

```bash
# Via Codex
$skill-installer install all skills from https://github.com/ByteLandTechnology/headless-ghidra

# Via npx (install all skills for all agents)
npx --yes skills add https://github.com/ByteLandTechnology/headless-ghidra --all

# For specific agent
npx --yes skills add https://github.com/ByteLandTechnology/headless-ghidra --agent codex --skill '*' --yes
```

## Usage

```markdown
Start a new analysis:
"Use the headless-ghidra skill to analyze ./sample-target. Start at P0 intake,
choose a stable target id, and stop after each phase gate so I can review."

Resume an existing target:
"Resume the same target and continue through P1 baseline."
```

## Prerequisites

- Ghidra installed locally
- Target binary in a workspace path the agent can read
- Frida (optional, for runtime observation)

## Output Directory Structure

```
targets/<target-id>/
artifacts/<target-id>/
```

Not in the skill directory — runtime output goes to the active workspace.
