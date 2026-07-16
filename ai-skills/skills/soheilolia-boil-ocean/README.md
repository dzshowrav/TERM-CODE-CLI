# boil-ocean

Boil Ocean is Soheil's complete-implementation standard as an installable agent skill.

Use it when the work should ship as a finished product, not a plan: search before
building, define done, implement the real fix, test it, document receipts, and make
only claims that are backed by verified evidence.

Invoking it is an execution signal: agents should proceed end-to-end without waiting
for a separate Apples approval unless the request says read-only, review-only, no
edits, no commit, or no push.

## Install

```bash
curl -sL https://raw.githubusercontent.com/SoheilOlia/skills/main/install.sh | bash -s boil-ocean
```

## Skill name

Invoke as:

```text
$boil-ocean
```

## Local files

- `SKILL.md`: canonical skill instructions
- `agents/openai.yaml`: OpenAI/Codex metadata
- `claude-code/boil-ocean.md`: direct Claude Code-compatible copy
