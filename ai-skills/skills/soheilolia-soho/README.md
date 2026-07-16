# soho

Soho is a disciplined execution skill pack. It is not Goose and does not require
Goose. These are plain skills that can be installed globally and used by Codex,
Claude Code, or any agent runtime that reads local skill folders.

## Install the full Soho pack

```bash
curl -sL https://raw.githubusercontent.com/SoheilOlia/skills/main/install.sh | bash -s soho
```

This installs every skill under `soho/skills/` into the global skill directory.
The `using-soho` skill also includes Cursor command shims for `/soho`,
`/soho-plan`, and `/soho-swarm` where the host scans `~/.cursor/commands`.
For Cursor projects that require project-local commands, install the source
plugin's commands into that project with:

```bash
~/agent-plugins/soho/scripts/install-cursor-project.sh /path/to/project
```

## Included skills

- `using-soho`: choose solo, swarm, or recommendation mode before work starts
- `soho-project-start`: start projects/specs with aligned markdown context, PRP/spec inputs, validation gates, and receipts
- `pr-review`: inspect PR truth, bot review comments, CI, diff risk, merge readiness, and required fixes with GPT-5.5 Extra High where available
- `brainstorming`: clarify creative implementation work before editing
- `writing-plans`: create exact plans for multi-step work
- `test-driven-development`: drive behavior changes through tests
- `systematic-debugging`: reproduce and root-cause failures before fixing
- `verification-before-completion`: verify outputs before claiming done
- `orchestrating-swarms`: decompose work across roles and boundaries
- `selecting-topology`: choose the right Soho swarm topology
- `subagent-driven-development`: execute a plan through role-bounded agents
- `synthesizing-results`: reconcile role outputs into one verified result

## Source

These skills were exported from `/Users/soheil/agent-plugins/soho/skills`.
