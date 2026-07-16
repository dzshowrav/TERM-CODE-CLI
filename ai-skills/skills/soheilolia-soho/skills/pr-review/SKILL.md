---
name: pr-review
description: "Use when reading, reviewing, triaging, fixing, or responding to pull requests, PR bot review comments, review threads, CI failures, merge readiness, branch updates, or PR handoff decisions. Enforces Soho PR Review v2: evidence-first diff inspection, GPT-5.5 Extra High where available, scoped fixes only, tested proof, no merge without explicit human instruction."
---

# PR Review v2

Use this skill for any PR review, PR comment remediation, bot review response, CI triage, or merge-readiness decision. Load `using-soho` first unless the host has already selected Soho mode for the task.

## Required Defaults

- Preferred model for model-backed PR review: `gpt-5.5` with `xhigh` reasoning, described to users as GPT-5.5 Extra High.
- If the host cannot select that model or reasoning level, continue with the best available model and record the actual model or host limitation in the final output or receipt.
- Do not add a robot icon to PR comments, review replies, commits, or handoff text by default.
- If AI disclosure is required, use plain text such as `AI-assisted review:` or `Updated by <agent name> (AI-assisted)`.
- Repo-local or user-level instructions can still require a specific signature. If they conflict, name the conflict instead of silently violating either policy.
- Never merge a PR unless the user explicitly asks for a merge in the current session.
- Never force-push, resolve review threads, approve a PR, mark a draft PR ready, dismiss reviews, or post comments on the user's behalf unless permission is explicit.

## Phase 1: Establish PR Truth

Collect current facts before editing or recommending action.

- Confirm GitHub access with `gh auth status`.
- Fetch PR metadata: URL, number, title, author, base branch, head branch, head SHA, draft state, mergeability, review decision, labels, requested reviewers, and linked issues if available.
- Fetch review comments, bot comments, unresolved threads, and latest CI/check status.
- Fetch the full diff against the PR base, not just local working-tree changes.
- Confirm whether the local checkout matches the PR head SHA. If it does not, fetch or check out the real PR branch before making claims.
- Read repo instructions such as `AGENTS.md`, `CLAUDE.md`, `.cursor/rules`, `CONTRIBUTING.md`, `README.md`, test docs, and relevant handoff/receipt files.
- Record truth mode explicitly when it matters: PR head, local HEAD, working tree, remote branch, CI status, receipt/docs.

## Phase 2: Protect Scope and Worktree

- Inspect `git status --short --branch` before editing.
- Treat existing unrelated changes as user-owned. Do not stage, revert, format, or "clean up" files outside the PR-review fix scope.
- Prefer checking out the PR branch or an isolated worktree over editing `main`.
- If the PR branch has drifted from the fetched PR head, stop and reconcile before editing.
- If the base branch moved, inspect conflicts or use a non-destructive merge/rebase preview before declaring the PR ready.
- Keep fixes surgical. Do not bundle opportunistic refactors, broad formatting, dependency upgrades, or unrelated docs.

## Phase 3: Review the Diff With Teeth

Classify changed files and look for failure modes beyond the comment text.

- Runtime behavior: control flow, data contracts, defaults, error paths, retries, idempotency, state transitions.
- Tests: missing coverage, weak assertions, fixture drift, flaky timing, untested negative paths.
- Generated artifacts: source-of-truth versus generated output, stale build products, accidental committed snapshots.
- Config and CI: workflows, secrets, permissions, environment variables, branch protections, package locks.
- Security and privacy: credentials, tokens, PII, authz/authn, logging, request signing, injection, path traversal.
- Data and migrations: irreversible writes, migration ordering, backfills, compatibility, rollback behavior.
- API and schema compatibility: versioning, optional fields, unknown field preservation, serialization drift.
- Observability: silent failure, misleading logs, missing receipts, overbroad success messages.
- Product/docs claims: docs saying something is wired, tested, or landed when code or CI does not prove it.

## Phase 4: Adjudicate Review Comments

Do not blindly obey bot or human comments. For each finding, choose one outcome:

- `fix`: the finding is valid and the PR should change.
- `prove false`: the concern does not apply; cite code, tests, or docs.
- `defer`: the concern is valid but out of scope; capture the exact follow-up and owner.
- `ask`: the comment is ambiguous or high-risk; ask before changing behavior.

For bot comments specifically:

- Verify the cited line still exists at the current PR head.
- Check whether the bot missed project context, generated-file boundaries, local conventions, or a newer commit.
- If a bot suggestion creates a worse contract, do not apply it. Explain the tradeoff and propose the narrower fix.

## Phase 5: Implement Required Fixes Only

- Edit the smallest set of files required to address valid findings.
- Preserve existing style, public APIs, serialization shape, and generated boundaries unless the finding requires changing them.
- Add or update tests only where they prove the reviewed behavior.
- If multiple findings are independent, keep commits separable when committing is requested or appropriate.
- Do not resolve, reply to, or close review threads until the fix is pushed and the user has allowed that action.

## Phase 6: Prove the Result

Run the tightest meaningful verification for the changed surface.

- Run the tests/checks suggested by repo docs, CI config, or the PR itself.
- If a command is unavailable, run the nearest equivalent and record the limitation.
- For docs-only PRs, verify links, file references, generated snippets, and truth claims.
- For generated outputs, rerun the generator if that is the canonical path, or state clearly that regeneration was not performed.
- Inspect final `git diff` and `git status --short --branch` before summarizing.
- If pushing, confirm the pushed SHA matches the intended PR branch after `git push`.

## Phase 7: Output Contract

Final or handoff output must include:

- PR URL and PR head SHA reviewed.
- Review comments addressed, rejected, deferred, or left for human decision.
- Files changed and why.
- Commands run and results.
- Whether updates were pushed.
- Whether comments were posted or threads resolved.
- Residual risks, blocked checks, and exact next step for the PR.

## Hard Stops

Stop and ask for direction when any of these are true:

- The PR branch cannot be checked out or fetched.
- The checkout has unrelated changes in files that must be edited.
- The requested fix requires changing public behavior outside the review finding.
- CI or tests require credentials, paid services, production access, or destructive data.
- A force-push, rebase of published history, branch deletion, merge, approval, or review dismissal would be needed.
- The model/runtime cannot provide enough evidence to support the requested decision.
