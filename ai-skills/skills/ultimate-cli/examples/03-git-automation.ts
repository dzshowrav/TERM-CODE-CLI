#!/usr/bin/env node
/**
 * ============================================================
 *  GIT AUTOMATION — Simple-Git + Enquirer + Chalk
 * ============================================================
 * Libraries: simple-git, enquirer, chalk, fast-glob
 *
 * An interactive Git workflow assistant with:
 *   - Smart commit with conventional commit format
 *   - Branch management (create, switch, merge, delete)
 *   - Interactive staging with file selection
 *   - Auto-detect changed files with fast-glob
 *   - Changelog generation
 *
 * Run: npx ts-node examples/03-git-automation.ts commit
 * ============================================================
 */

import simpleGit, { type StatusResult, type DefaultLogFields } from 'simple-git';
import Enquirer from 'enquirer';
import chalk from 'chalk';
import fg from 'fast-glob';
import ignore from 'ignore';
import fs from 'fs/promises';
import path from 'path';
import { Command } from 'commander';

// ─── Types ───────────────────────────────────────────────
interface CommitInfo {
  type: string;
  scope?: string;
  description: string;
  breaking?: string;
  body?: string;
  issues?: string;
}

// ─── Git Wrapper ─────────────────────────────────────────
const git = simpleGit(process.cwd(), { binary: 'git' });

// ─── Helpers ─────────────────────────────────────────────
const CONVENTIONAL_TYPES = [
  { name: '✨ feat     — A new feature', value: 'feat' },
  { name: '🐛 fix      — A bug fix', value: 'fix' },
  { name: '📚 docs     — Documentation changes', value: 'docs' },
  { name: '💄 style    — Code style (formatting, etc.)', value: 'style' },
  { name: '♻️  refactor — Code refactoring', value: 'refactor' },
  { name: '⚡ perf     — Performance improvements', value: 'perf' },
  { name: '🧪 test     — Adding/updating tests', value: 'test' },
  { name: '🔧 build    — Build system or dependencies', value: 'build' },
  { name: '👷 ci       — CI/CD changes', value: 'ci' },
  { name: '🔨 chore    — Other changes', value: 'chore' },
];

function formatCommitMessage(info: CommitInfo): string {
  let msg = `${info.type}`;
  if (info.scope) msg += `(${info.scope})`;
  if (info.breaking) msg += '!';
  msg += `: ${info.description}`;
  if (info.body) msg += `\n\n${info.body}`;
  if (info.issues) msg += `\n\nCloses: ${info.issues}`;
  return msg;
}

function colorStatus(status: string): string {
  switch (status) {
    case 'modified': return chalk.yellow('M');
    case 'added': return chalk.green('A');
    case 'deleted': return chalk.red('D');
    case 'renamed': return chalk.blue('R');
    case 'conflicted': return chalk.bgRed.white('!');
    default: return chalk.dim('?');
  }
}

// ─── Commands ────────────────────────────────────────────

// === status ===
async function showStatus(): Promise<void> {
  const status: StatusResult = await git.status();

  console.log(chalk.bold('\n  📊 Git Status\n'));
  console.log(chalk.dim(`  Branch: `), chalk.bold(status.current));
  console.log(chalk.dim(`  Ahead:  `), status.ahead);
  console.log(chalk.dim(`  Behind: `), status.behind);
  console.log();

  const sections = [
    { title: 'Staged', files: status.staged, color: chalk.green },
    { title: 'Modified', files: status.modified, color: chalk.yellow },
    { title: 'Created', files: status.created, color: chalk.green },
    { title: 'Deleted', files: status.deleted, color: chalk.red },
    { title: 'Renamed', files: status.renamed, color: chalk.blue },
    { title: 'Conflicted', files: status.conflicted, color: chalk.bgRed.white },
    { title: 'Not Staged', files: status.not_added, color: chalk.dim },
  ];

  for (const section of sections) {
    if (section.files.length > 0) {
      console.log(chalk.bold(`  ${section.title}:`));
      section.files.forEach((f) => console.log(`    ${section.color('▸')} ${f}`));
      console.log();
    }
  }

  if (status.isClean()) {
    console.log(chalk.green('  ✅ Working tree clean\n'));
  }
}

// === commit ===
async function doCommit(options?: { message?: string }): Promise<void> {
  const status = await git.status();

  if (status.isClean() && status.not_added.length === 0 && status.created.length === 0) {
    console.log(chalk.yellow('\n  ⚠ Nothing to commit\n'));
    return;
  }

  console.log(chalk.bold('\n  📝 New Commit\n'));

  // Stage files interactively
  const allChanged = [
    ...status.not_added.map(f => ({ name: `[NEW]  ${f}`, value: f, checked: true })),
    ...status.modified.map(f => ({ name: `[MOD]  ${f}`, value: f, checked: true })),
    ...status.deleted.map(f => ({ name: `[DEL]  ${f}`, value: f, checked: true })),
    ...status.created.map(f => ({ name: `[NEW]  ${f}`, value: f, checked: true })),
    ...status.renamed.map(f => ({ name: `[REN]  ${f}`, value: f, checked: true })),
  ];

  const selectedFiles: string[] = options?.message
    ? allChanged.map(f => f.value)  // non-interactive: stage all
    : await Enquirer.prompt({
        type: 'multiselect',
        name: 'files',
        message: 'Select files to stage:',
        choices: allChanged,
        validate: (selected) => selected.length > 0 || 'Select at least one file',
      }).then((r: any) => r.files) as string[];

  // Stage selected
  await git.add(selectedFiles);

  // Commit message
  let commitInfo: CommitInfo;
  if (options?.message) {
    // Parse conventional commit from message
    const match = options.message.match(/^(\w+)(?:\((.+)\))?(!)?:\s*(.+)$/);
    if (match) {
      commitInfo = {
        type: match[1],
        scope: match[2],
        breaking: match[3] || undefined,
        description: match[4],
      };
    } else {
      commitInfo = { type: 'chore', description: options.message };
    }
  } else {
    const type: string = await Enquirer.prompt({
      type: 'select',
      name: 'type',
      message: 'Commit type:',
      choices: CONVENTIONAL_TYPES,
      pointer: chalk.green('❯'),
    }).then((r: any) => r.type);

    const scope: string | undefined = await Enquirer.prompt({
      type: 'input',
      name: 'scope',
      message: 'Scope (optional, e.g., "api", "ui"):',
      initial: '',
    }).then((r: any) => r.scope || undefined);

    const description: string = await Enquirer.prompt({
      type: 'input',
      name: 'description',
      message: 'Short description:',
      validate: (v: string) => v.length > 0 ? true : 'Required',
    }).then((r: any) => r.description);

    const breaking: string | undefined = await Enquirer.prompt({
      type: 'confirm',
      name: 'breaking',
      message: 'Breaking change?',
      initial: false,
    }).then((r: any) => r.breaking ? 'BREAKING CHANGE' : undefined);

    const body: string | undefined = await Enquirer.prompt({
      type: 'input',
      name: 'body',
      message: 'Body (optional, multi-line with \\n):',
    }).then((r: any) => r.body || undefined);

    commitInfo = { type, scope, description, breaking, body };
  }

  const commitMessage = formatCommitMessage(commitInfo);

  // Confirm
  console.log(chalk.dim(`\n  Commit message:\n`));
  console.log(chalk.cyan(`  ${commitMessage}`));
  console.log();

  const confirm = options?.message || await Enquirer.prompt({
    type: 'confirm',
    name: 'confirmed',
    message: 'Commit?',
    initial: true,
  }).then((r: any) => r.confirmed);

  if (confirm) {
    const result = await git.commit(commitMessage);
    console.log(chalk.green(`\n  ✅ ${result.summary?.changes || 0} files changed, ${result.summary?.insertions || 0} insertions, ${result.summary?.deletions || 0} deletions\n`));
  } else {
    console.log(chalk.yellow('\n  ⏹ Commit cancelled\n'));
  }
}

// === branch ===
async function branchManagement(): Promise<void> {
  const branches = await git.branchLocal();

  const action: string = await Enquirer.prompt({
    type: 'select',
    name: 'action',
    message: 'Branch action:',
    choices: [
      { name: 'Create branch', value: 'create' },
      { name: 'Switch branch', value: 'switch' },
      { name: 'Merge branch', value: 'merge' },
      { name: 'Delete branch', value: 'delete' },
      { name: 'Create from current', value: 'checkout-new' },
    ],
  }).then((r: any) => r.action);

  if (action === 'create' || action === 'checkout-new') {
    const name: string = await Enquirer.prompt({
      type: 'input',
      name: 'name',
      message: 'Branch name:',
      validate: (v: string) => /^[a-zA-Z0-9_/.-]+$/.test(v) || 'Invalid branch name',
    }).then((r: any) => r.name);

    if (action === 'checkout-new') {
      await git.checkoutLocalBranch(name);
    } else {
      await git.branch([name]);
    }
    console.log(chalk.green(`\n  ✅ Branch '${name}' created\n`));
  }

  if (action === 'switch') {
    const branchNames = Object.keys(branches.branches);
    const name: string = await Enquirer.prompt({
      type: 'autocomplete',
      name: 'name',
      message: 'Switch to branch:',
      choices: branchNames.map(b => ({ name: b === branches.current ? `${b} (current)` : b, value: b })),
    }).then((r: any) => r.name);

    if (name !== branches.current) {
      await git.checkout(name);
      console.log(chalk.green(`\n  ✅ Switched to '${name}'\n`));
    }
  }

  if (action === 'merge') {
    const targetBranch: string = await Enquirer.prompt({
      type: 'autocomplete',
      name: 'name',
      message: 'Merge which branch into current?',
      choices: Object.keys(branches.branches)
        .filter(b => b !== branches.current)
        .map(b => ({ name: b, value: b })),
    }).then((r: any) => r.name);

    try {
      const result = await git.merge([targetBranch]);
      console.log(chalk.green(`\n  ✅ Merged '${targetBranch}' into '${branches.current}'\n`));
    } catch (error: any) {
      console.log(chalk.red(`\n  ✖ Merge conflict: ${error.message}\n`));
    }
  }

  if (action === 'delete') {
    const toDelete: string = await Enquirer.prompt({
      type: 'select',
      name: 'name',
      message: 'Delete branch:',
      choices: Object.keys(branches.branches)
        .filter(b => b !== branches.current)
        .map(b => ({ name: b, value: b })),
    }).then((r: any) => r.name);

    const confirm = await Enquirer.prompt({
      type: 'confirm',
      name: 'confirmed',
      message: `Delete '${toDelete}'?`,
      initial: false,
    }).then((r: any) => r.confirmed);

    if (confirm) {
      await git.branch(['-D', toDelete]);
      console.log(chalk.green(`\n  ✅ Deleted '${toDelete}'\n`));
    }
  }
}

// === merge ===
async function mergeBranch(branch: string) {
  const status = await git.status();
  if (!status.isClean()) {
    console.log(chalk.red('\n  ✖ Working tree is not clean. Commit or stash first.\n'));
    return;
  }

  try {
    const result = await git.merge([branch]);
    console.log(chalk.green(`\n  ✅ Merged '${branch}' into '${status.current}'\n`));
  } catch (error: any) {
    console.log(chalk.red(`\n  ✖ Merge conflict: ${error.message}`));
    console.log(chalk.yellow('  Resolve conflicts and run: git merge --continue\n'));
  }
}

// === log ===
async function showLog(limit: number = 10): Promise<void> {
  const log = await git.log({ maxCount: limit });

  console.log(chalk.bold(`\n  📜 Recent Commits (last ${limit})\n`));
  for (const commit of log.all) {
    const date = new Date(commit.date);
    const formatted = date.toLocaleDateString('en-US', {
      month: 'short', day: 'numeric', hour: '2-digit', minute: '2-digit',
    });
    const msg = commit.message.split('\n')[0];
    const hash = commit.hash.slice(0, 7);
    console.log(`  ${chalk.yellow(hash)} ${chalk.dim(formatted)} ${msg}`);
  }
  console.log();
}

// === changelog ===
async function generateChangelog(tag?: string): Promise<void> {
  const log = await git.log({
    maxCount: 100,
    ...(tag ? { from: tag, to: 'HEAD' } : {}),
  });

  const conventional = log.all.filter((c) =>
    /^(feat|fix|docs|style|refactor|perf|test|build|ci|chore)(\(.+\))?(!)?:\s/.test(c.message)
  );

  const changes = {
    feat: [] as DefaultLogFields[],
    fix: [] as DefaultLogFields[],
    docs: [] as DefaultLogFields[],
    refactor: [] as DefaultLogFields[],
    perf: [] as DefaultLogFields[],
    chore: [] as DefaultLogFields[],
  };

  for (const commit of conventional) {
    const type = commit.message.match(/^(\w+)/)?.[1] as keyof typeof changes;
    if (type in changes) changes[type].push(commit);
  }

  console.log(chalk.bold('\n  📋 Changelog\n'));
  console.log(chalk.dim(`  ${new Date().toISOString().split('T')[0]}\n`));

  const sections: [string, string, DefaultLogFields[]][] = [
    ['🚀 Features', 'feat', changes.feat],
    ['🐛 Bug Fixes', 'fix', changes.fix],
    ['📚 Documentation', 'docs', changes.docs],
    ['♻️ Refactoring', 'refactor', changes.refactor],
    ['⚡ Performance', 'perf', changes.perf],
    ['🔧 Maintenance', 'chore', changes.chore],
  ];

  for (const [title, _, items] of sections) {
    if (items.length > 0) {
      console.log(chalk.bold(`  ${title}:`));
      for (const item of items) {
        const msg = item.message.replace(/^(feat|fix|docs|refactor|perf|chore)(\(.+\))?(!)?:\s/, '');
        console.log(`    • ${msg} (${chalk.dim(item.hash.slice(0, 7))})`);
      }
      console.log();
    }
  }
}

// ─── CLI ────────────────────────────────────────────────
const program = new Command()
  .name('git-auto')
  .description('Interactive Git workflow automation')
  .version('2.0.0');

program
  .command('status')
  .description('Show detailed working tree status')
  .action(showStatus);

program
  .command('commit')
  .description('Interactive conventional commit')
  .option('-m, --message <msg>', 'Commit message (skips prompts)')
  .action((opts) => doCommit(opts));

program
  .command('branch')
  .description('Interactive branch management')
  .alias('br')
  .action(branchManagement);

program
  .command('merge <branch>')
  .description('Merge a branch into current')
  .action(mergeBranch);

program
  .command('log')
  .description('Show recent commits')
  .option('-n, --number <count>', 'Number of commits', parseInt, 10)
  .action((opts) => showLog(opts.number));

program
  .command('changelog')
  .description('Generate changelog from conventional commits')
  .option('--from <tag>', 'Starting tag')
  .action((opts) => generateChangelog(opts.from));

program.parse(process.argv);
