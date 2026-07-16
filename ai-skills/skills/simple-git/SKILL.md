# simple-git

Lightweight Node.js interface for running `git` commands programmatically. Promise, callback, and chaining APIs.

## Install

```bash
npm install simple-git
```

Requires `git` binary on PATH. Node.js >= 14.17 (ESM/CJS/TypeScript with bundled types).

## Setup

```typescript
import { simpleGit, SimpleGit, SimpleGitOptions } from 'simple-git';

// default: cwd = process.cwd()
const git: SimpleGit = simpleGit();

// with options
const options: Partial<SimpleGitOptions> = {
  baseDir: process.cwd(),
  binary: 'git',
  maxConcurrentProcesses: 6,
  trimmed: false,           // auto-trim whitespace from raw responses
  config: ['http.proxy=someproxy'],  // -c flags prefixed to every command
};
const git: SimpleGit = simpleGit(options);

// split baseDir (backward compatible)
const git: SimpleGit = simpleGit('/path/to/repo', { binary: 'git' });
```

## Chaining & Promises

```typescript
const git = simpleGit();

// chain — runs in series
await git.init().addRemote('origin', '...remote.git');

// or await individually
await git.init();
await git.addRemote('origin', '...remote.git');

// catch-all
try {
  await git.pull();
} catch (e) { /* handle */ }

// per-step catch (chain continues)
await git.init().catch(() => {});

// callback style
git.init((err, result) => {});
```

## API Reference

### init

```typescript
git.init(bare?: boolean, options?: Options): Promise<InitResult>
git.init(options?: Options): Promise<InitResult>
```

### clone

```typescript
git.clone(repoPath: string, localPath?: string, options?: Options): Promise<string>
git.mirror(repoPath: string, localPath?: string, options?: Options): Promise<string>
```

### add

```typescript
git.add(files: string | string[]): Promise<string>
git.addTag(name: string): Promise<{ name: string }>
git.addAnnotatedTag(name: string, message: string): Promise<{ name: string }>
```

### commit

```typescript
git.commit(message: string | string[], files?: string | string[], options?: Options): Promise<CommitResult>
```

### branch

```typescript
git.branch(options?: Options): Promise<BranchSummary>
git.branchLocal(): Promise<BranchSummary>
git.deleteLocalBranch(branchName: string, forceDelete?: boolean): Promise<BranchSingleDeleteResult>
git.deleteLocalBranches(branchNames: string[], forceDelete?: boolean): Promise<BranchSingleDeleteResult>
```

### checkout

```typescript
git.checkout(checkoutWhat: string, options?: Options): Promise<string>
git.checkoutBranch(branchName: string, startPoint: string): Promise<string>
git.checkoutLocalBranch(branchName: string): Promise<string>
```

### fetch

```typescript
git.fetch(options?: Options): Promise<FetchResult>
git.fetch(remote: string, branch: string, options?: Options): Promise<FetchResult>
```

### pull

```typescript
git.pull(options?: Options): Promise<PullResult>
git.pull(remote: string, branch: string, options?: Options): Promise<PullResult>
```

### push

```typescript
git.push(options?: Options): Promise<PushResult>
git.push(remote: string, branch: string, options?: Options): Promise<PushResult>
git.pushTags(remote: string, options?: Options): Promise<PushResult>
```

### merge

```typescript
git.merge(options?: Options): Promise<MergeSummary>
git.mergeFromTo(remote: string, branch: string, options?: Options): Promise<MergeSummary>
```

Conflicts produce a `GitResponseError` with `.git` containing the `MergeSummary`.

### rebase

```typescript
git.rebase(options?: string[] | Options): Promise<string>
```

### reset

```typescript
git.reset(mode: ResetMode | 'mixed' | 'soft' | 'hard' | 'merge' | 'keep', options?: Options): Promise<string>
git.reset(options?: Options): Promise<string>
git.reset(): Promise<string>
```

### revert

```typescript
git.revert(commit: string, options?: Options): Promise<string>
```

### status

```typescript
git.status(options?: Options): Promise<StatusResult>
```

### log

```typescript
git.log(options?: LogOptions): Promise<LogResult>
```

LogOptions:
| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `from` | string | — | Oldest commit |
| `to` | string | `HEAD` | Newest commit |
| `file` | string | — | Restrict to file path |
| `format` | object | — | Custom format `{ key: formatStr }` |
| `maxCount` | number | — | `--max-count` |
| `multiLine` | boolean | false | Multiline body |
| `strictDate` | boolean | false | Strict ISO 8601 dates |
| `symmetric` | boolean | true | Symmetric revision range |
| `splitter` | string | `ò` | Field delimiter |
| `mailMap` | boolean | true | Use `.mailmap` |

### diff

```typescript
git.diff(options?: Options): Promise<string>
git.diffSummary(options?: Options): Promise<DiffResult>
```

### stash

```typescript
git.stash(options?: Options | string[]): Promise<string>
git.stashList(options?: Options): Promise<LogResult>
```

### show

```typescript
git.show(options?: Options): Promise<string>
git.showBuffer(options?: Options): Promise<Buffer>
```

### tag

```typescript
git.tag(args?: string[]): Promise<string>
git.tags(options?: Options): Promise<TagResult>
```

### config

```typescript
git.addConfig(key: string, value: string, append?: boolean, scope?: GitConfigScope): Promise<string>
git.getConfig(key: string, scope?: GitConfigScope): Promise<ConfigGetResult>
git.listConfig(scope?: GitConfigScope): Promise<ConfigListSummary>
```

`GitConfigScope`: `'worktree' | 'local' | 'global' | 'system'`

### remote

```typescript
git.addRemote(name: string, repo: string, options?: Options): Promise<string>
git.removeRemote(name: string): Promise<string>
git.getRemotes(verbose?: boolean): Promise<RemoteWithRefs[]>
git.listRemote(options?: Options): Promise<string>
git.remote(options?: Options): Promise<string>
```

### clean

```typescript
import { CleanOptions } from 'simple-git';
git.clean(mode: CleanOptions | 'n' | 'f'): Promise<string>
git.clean(switches: string, options?: Options): Promise<string>
```

### rm

```typescript
git.rm(files: string | string[]): Promise<string>
git.rmKeepLocal(files: string | string[]): Promise<string>
```

### mv

```typescript
git.mv(from: string | string[], to: string): Promise<string>
```

### submodule

```typescript
git.subModule(options?: Options): Promise<string>
git.submoduleAdd(repo: string, path: string): Promise<string>
git.submoduleInit(options?: Options): Promise<string>
git.submoduleUpdate(subModuleName?: string, options?: Options): Promise<string>
```

### applyPatch

```typescript
git.applyPatch(patch: string, options?: Options): Promise<string>
git.applyPatch(patches: string[], options?: Options): Promise<string>
```

### grep

```typescript
import { grepQueryBuilder } from 'simple-git';
git.grep(searchTerm: string, options?: Options): Promise<GrepResult>
git.grep(query: ReturnType<typeof grepQueryBuilder>, options?: Options): Promise<GrepResult>
```

### Check / Info

```typescript
git.checkIsRepo(): Promise<boolean>
git.checkIsRepo('bare' | 'root'): Promise<boolean>
git.firstCommit(): Promise<string>
git.revparse(options: string | Options): Promise<string>
git.catFile(options: string[]): Promise<string>
git.countObjects(): Promise<CountObjectsResult>
git.hashObject(filePath: string, write?: boolean): Promise<string>
git.version(): Promise<GitVersionResult>
```

### checkIgnore

```typescript
git.checkIgnore(filepaths: string | string[]): Promise<string[]>
```

### cwd (change working directory)

```typescript
git.cwd(workingDirectory: string): Promise<string>
git.cwd({ path: string, root?: boolean }): Promise<string>
```

### env (environment variables)

```typescript
git.env(name: string, value: string): this
git.env(env: Record<string, string>): this
```

### raw (arbitrary git commands)

```typescript
git.raw(args: string | string[]): Promise<string>
git.raw(...args: string[]): Promise<string>
```

### exec (run function in chain)

```typescript
git.exec(handler: () => void): this
```

### customBinary

```typescript
git.customBinary(gitPath: string): this
```

### outputHandler

```typescript
git.outputHandler(fn: (command: string, stdout: Readable, stderr: Readable) => void): this
```

## Options

Tasks accept options as an **object** or **array of strings**.

Object form — string value → `name=value`, null/other value → just the key:
```typescript
git.pull('origin', 'master', { '--no-rebase': null });
// => git pull origin master --no-rebase

git.pull('origin', 'master', { '--rebase': 'true' });
// => git pull origin master --rebase=true

git.log({ '--grep': ['bug', 'fix'] });
// => git log --grep=bug --grep=fix
```

Array form:
```typescript
git.pull('origin', 'master', ['--no-rebase']);
```

### pathspec helper

```typescript
import { pathspec } from 'simple-git';

git.status([pathspec('sub-dir')]);
git.status({ 'sub-dir': pathspec('sub-dir') });
```

## Plugins

| Plugin | Purpose |
|--------|---------|
| AbortController | Terminate pending/future tasks (node >= 16) |
| Custom Binary | Override `git` binary path |
| Completion Detection | Customise end-of-process detection |
| Error Detection | Customise error detection |
| Progress Events | Receive `{ method, stage, progress }` callbacks |
| Spawn Options | Configure `uid`/`gid` for spawned processes |
| Timeout | Auto-kill git process after rolling timeout |
| Unsafe | Opt out of safety precautions |

Progress events via options:
```typescript
const git = simpleGit({
  progress({ method, stage, progress }) {
    console.log(`git.${method} ${stage} stage ${progress}% complete`);
  },
});
```

## Error Handling

```typescript
import { GitResponseError } from 'simple-git';

try {
  const summary = await git.merge();
} catch (err) {
  const mergeSummary = (err as GitResponseError<MergeSummary>).git;
  console.error(`${mergeSummary.conflicts.length} conflicts`);
}
```

`GitError` — no parser available. `GitResponseError` — parser ran, `.git` has parsed content.

## Debug Logging

```bash
DEBUG=simple-git node app.js
DEBUG=simple-git:task:*,simple-git:output:* node app.js
```

## Common Patterns

```typescript
// init new repo
await git.init().add('./*').commit('first commit!')
  .addRemote('origin', 'https://github.com/user/repo.git')
  .push('origin', 'master');

// check if repo, init if not
if (!await git.checkIsRepo()) {
  await git.init().addRemote('origin', 'https://some.git.repo');
}

// parallel independent tasks
const [cdup, prefix] = await Promise.all([
  git.raw('rev-parse', '--show-cdup').catch(() => null),
  git.raw('rev-parse', '--show-prefix').catch(() => null),
]);

// set author per-commit
await git.addConfig('user.name', 'Some One').addConfig('user.email', 'some@one.com');
await git.commit('message', 'file', { '--author': '"Another Person <a@b.com>"' });

// restart app on changes
const { summary } = await git.pull();
if (summary.changes) require('child_process').exec('npm restart');

// auth via URL
const remote = `https://${USER}:${PASS}@github.com/username/private-repo`;
await git.clone(remote);
```
