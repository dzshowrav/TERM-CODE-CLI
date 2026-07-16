# Git Libraries — Complete A-to-Z Reference

---

## 1. Isomorphic-Git (isomorphic-git/isomorphic-git)
**GitHub**: https://github.com/isomorphic-git/isomorphic-git | **Stars**: 7.5K+
**npm**: `isomorphic-git` | **Weekly**: ~75K | **License**: MIT
**Website**: https://isomorphic-git.org/

### 1.1 Installation
```bash
npm install isomorphic-git
```

### 1.2 Complete API

#### Setup & Configuration
```typescript
import git from 'isomorphic-git';
import fs from 'fs';

const dir = '/path/to/repo';

// All functions require a `fs` and `dir` parameter:
// git.func({ fs, dir, ... })
```

#### Core Operations

##### init — Initialize Repository
```typescript
await git.init({ fs, dir });
await git.init({ fs, dir, bare: true });   // bare repo
await git.init({ fs, dir, defaultBranch: 'main' }); // custom default branch
```

##### clone — Clone Repository
```typescript
await git.clone({
  fs,
  dir,                    // destination
  url: 'https://github.com/user/repo.git',
  ref: 'main',            // branch/tag
  singleBranch: true,     // only clone one branch
  depth: 1,               // shallow clone
  noTags: false,
  remote: 'origin',       // remote name (default: 'origin')
  onProgress: (progress) => { console.log(progress); },
  onMessage: (message) => { console.log(message); },
  onAuth: () => ({ username: 'token', password: 'ghp_xxx' }),
  onAuthFailure: (url, auth) => { console.log('Auth failed'); },
  headers: { 'User-Agent': 'my-app' },
  noCheckout: false,       // bare-like (no checkout)
});
```

##### fetch — Fetch Without Merge
```typescript
await git.fetch({
  fs,
  dir,
  remote: 'origin',
  ref: 'main',
  depth: 1,
  singleBranch: true,
  prune: false,           // remove deleted remote refs
  pruneTags: false,
  onProgress: (p) => {},
  onAuth: () => ({ username: 'token', password: 'xxx' }),
});
```

##### push — Push Commits
```typescript
await git.push({
  fs,
  dir,
  remote: 'origin',
  ref: 'main',             // local branch to push
  remoteRef: 'refs/heads/main', // remote ref (or remote branch name)
  force: false,            // force push
  delete: false,           // delete remote branch
  onProgress: (p) => {},
  onAuth: () => ({ username: 'token' }),
  onAuthFailure: (url, auth) => {},
  remoteRef: 'main',
});
```

##### pull — Fetch + Merge
```typescript
await git.pull({
  fs,
  dir,
  ref: 'main',
  singleBranch: true,
  fastForwardOnly: true,
  onAuth: () => ({ username: 'token' }),
  // onProgress, onMessage, etc.
});
```

#### Status & Diff

##### status
```typescript
const status = await git.status({ fs, dir, filepath: 'src/index.ts' });
// Returns one of:
// 'ignored'     | .gitignore'd
// 'unmodified'  | no changes
// '*modified'   | has unstaged changes
// '*deleted'    | unstaged deletion
// '*added'      | unstaged new file
// 'added'       | staged new file
// 'deleted'     | staged deletion
// 'modified'    | staged modification
// 'absent'      | not tracked

// Check multiple files
const statusMatrix = await git.statusMatrix({
  fs, dir,
  filter: (f) => f.endsWith('.ts'),
});
// Returns: [[filepath, HEAD, WORKDIR, STAGE], ...]
// HEAD: 0 (absent) or 1 (exists)
// WORKDIR: 0 (absent) or 1 (exists) or 2 (modified)
// STAGE: similar
```

##### log — Commit History
```typescript
const commits = await git.log({
  fs, dir,
  ref: 'main',           // branch or tag
  depth: 10,             // number of commits
  since: new Date('2025-01-01'), // commits after date
  // or: since: timestamp_ms
});

// commit = {
//   oid: 'abc123...',
//   message: 'commit message',
//   tree: 'tree_oid',
//   parent: ['parent_oid'],
//   author: {
//     name: 'Author',
//     email: 'author@email.com',
//     timestamp: 1234567890,
//     timezoneOffset: 60,
//   },
//   committer: { ... },
// }
```

##### readBlob — Read File at Commit
```typescript
const { blob, oid } = await git.readBlob({
  fs, dir,
  oid: 'commit_sha',     // OR
  filepath: 'src/index.ts',
});

// blob is a Uint8Array
const text = new TextDecoder().decode(blob);
```

##### readTree — Read Directory Tree
```typescript
const tree = await git.readTree({
  fs, dir,
  oid: 'tree_sha',       // OR
  ref: 'main',
  filepath: 'src',       // specific directory
});

// tree = {
//   oid: 'tree_sha',
//   tree: [
//     { mode: '100644', path: 'index.ts', oid: 'blob_sha', type: 'blob' },
//     { mode: '040000', path: 'utils', oid: 'tree_sha', type: 'tree' },
//   ],
// }
```

##### diff — Raw Diff
```typescript
// There is no built-in formatted diff, but you can compare:
await git.walk({
  fs, dir,
  trees: [git.HEAD({ fs, dir }), git.WORKDIR({ fs, dir })],
  // Hand-rolled diff via map/reduce
});
```

#### Branch Management

##### branch — List/Create/Delete Branches
```typescript
// Create branch
await git.branch({ fs, dir, ref: 'feature/new', checkout: false });
await git.branch({ fs, dir, ref: 'feature/new', object: 'commit_sha' });

// Delete branch
await git.branch({ fs, dir, ref: 'feature/old', delete: true });

// Rename
await git.branch({ fs, dir, ref: 'old-name', rename: true });
// ^ use checkout + branch for renaming

// List branches
const branches = await git.listBranches({ fs, dir });
// ['main', 'develop', 'feature/a']

// List remote branches
const remoteBranches = await git.listBranches({ fs, dir, remote: 'origin' });
```

##### checkout — Switch Branch
```typescript
await git.checkout({
  fs, dir,
  ref: 'main',
  noCheckout: false,         // no checkout files (bare-like)
  noUpdateHead: false,        // don't update HEAD
  force: false,
  remote: 'origin',           // remote to track
});

// Create + checkout
await git.checkout({
  fs, dir,
  ref: 'feature/new',
  noCheckout: false,
});
```

#### Staging & Committing

##### add — Stage Files
```typescript
await git.add({ fs, dir, filepath: 'src/index.ts' });
await git.add({ fs, dir, filepath: '.' });       // stage all
await git.add({ fs, dir, filepath: ['src/a.ts', 'src/b.ts'] });
```

##### remove — Remove Files
```typescript
await git.remove({ fs, dir, filepath: 'src/old.ts' });
```

##### commit — Create Commit
```typescript
const sha = await git.commit({
  fs, dir,
  message: 'feat: add new feature',
  author: {
    name: 'Author Name',
    email: 'author@email.com',
  },
  committer: {
    name: 'Committer',
    email: 'committer@email.com',
  },
  date: new Date(),             // author date (or timestamp)
  signingKey: '...',            // GPG signing
  amend: false,                 // amend last commit
  noVerify: false,              // skip hooks
// ^ returns commit SHA (oid)
});
```

##### stash — Save WIP
```typescript
await git.stash({
  fs, dir,
  message: 'WIP: in progress',
});

// There is no unstash directly; use log on refs/stash
```

#### Reset & Revert

##### reset — Reset HEAD
```typescript
await git.reset({
  fs, dir,
  ref: 'HEAD~1',         // OR 'main', 'commit_sha'
  hard: false,           // reset index (staging)
  // hard: true also resets working directory
});
```

#### Tags

##### tag — Add/List/Delete Tags
```typescript
// Add tag
const oid = await git.tag({
  fs, dir,
  ref: 'v1.0.0',
  object: 'commit_sha',         // (optional, HEAD if not specified)
  type: 'annotated',            // or 'lightweight'
  tagger: { name: 'Author', email: '...', timestamp: ... },
  message: 'Release v1.0.0',
  force: false,
});

// Delete tag
await git.tag({ fs, dir, ref: 'v1.0.0', delete: true });

// List tags
const tags = await git.listTags({ fs, dir });
// ['v1.0.0', 'v1.1.0']

// Get tag details
const tag = await git.readTag({ fs, dir, oid: 'tag_sha' });
```

#### Reference Management

##### readRef — Read References
```typescript
const ref = await git.readRef({
  fs, dir,
  ref: 'HEAD',
});

// Returns:
// { oid: 'abc...', target: 'refs/heads/main', type: 'symbolic' }
// { oid: 'abc...', type: 'direct' }  (detached HEAD)
```

##### log / reflog — Reflog
```typescript
const reflog = await git.log({
  fs, dir,
  ref: 'HEAD',
  depth: 100,
});
// Returns commits with reflog entries
```

#### Remote Management

##### listRemotes — List Remotes
```typescript
const remotes = await git.listRemotes({ fs, dir });
// [{ remote: 'origin', url: 'https://...' }]
```

##### addRemote / deleteRemote
```typescript
await git.addRemote({
  fs, dir,
  remote: 'upstream',
  url: 'https://...',
  force: false,
});

await git.deleteRemote({
  fs, dir,
  remote: 'origin',
});
```

##### getRemoteInfo — Remote Info Without Clone
```typescript
const info = await git.getRemoteInfo({
  url: 'https://github.com/user/repo.git',
  forPush: false,
  protocolVersion: 2,    // 0 = auto, 1 = v1, 2 = v2
  onAuth: () => {},
  // onProgress, onMessage
});

// Returns:
// {
//   capabilities: ['...'],
//   refs: {
//     HEAD: { oid: '...' },
//     'refs/heads/main': { oid: '...' },
//     'refs/tags/v1.0.0': { oid: '...' },
//   },
// }
```

#### Merge

##### merge — Merge Branches
```typescript
const mergeResult = await git.merge({
  fs, dir,
  ours: 'main',           // current branch
  theirs: 'feature',      // branch to merge
  fastForwardOnly: false,
  abortOnConflict: false,
  author: { name: '...', email: '...' },
  committer: { ... },
  message: 'Merge feature into main',
  signingKey: '...',
  dryRun: false,          // check for conflicts without merging
  noVerify: false,
});

// On success:
// { oid: 'merge_commit_sha', tree: 'tree_sha' }

// On conflict (without abortOnConflict):
// Emits conflicts that must be resolved manually
```

#### findRoot — Find Repo Root
```typescript
const root = await git.findRoot({
  fs,
  filepath: '/path/to/repo/src/index.ts',
});
// Returns '/path/to/repo' (or throws if not in a repo)
```

#### ExpandOid — Expand SHA
```typescript
const oid = await git.expandOid({
  fs, dir,
  oid: 'abc123',          // partial SHA
});
// Returns full SHA 'abc123def456...'
```

#### HashBlob — Hash Object
```typescript
const { oid, type, object, format } = await git.hashBlob({
  object: 'file content',
  format: 'utf8',         // 'utf8' | 'wrapped' | 'content' | 'parsed'
});
```

#### ListFiles — List All Files
```typescript
const files = await git.listFiles({ fs, dir });
// ['src/index.ts', 'package.json', ...]

// At specific ref:
const refFiles = await git.listFiles({ fs, dir, ref: 'main' });
```

#### Pack Objects
```typescript
// Low-level packfile operations
const pack = await git.packObjects({
  fs, dir,
  oids: ['commit_sha1', 'commit_sha2'],
  writePack: false,
});

// Unpack
await git.unpackObjects({
  fs, dir,
  objects: pack,
});
```

#### Walk — Custom Tree Walk
```typescript
import { TREE } from 'isomorphic-git';

const results = await git.walk({
  fs, dir,
  trees: [git.HEAD({ fs, dir }), git.WORKDIR({ fs, dir })],
  map: async (filepath, [head, workdir]) => {
    if (head === null && workdir !== null) {
      return 'ADDED: ' + filepath;
    }
    if (head !== null && workdir === null) {
      return 'DELETED: ' + filepath;
    }
    if (head.oid !== workdir.oid) {
      return 'MODIFIED: ' + filepath;
    }
  },
});

// Available tree sources:
git.HEAD({ fs, dir });        // HEAD tree
git.WORKDIR({ fs, dir });     // working directory
git.STAGE({ fs, dir });       // staging area
git.TREE({ fs, dir, ref: 'main', filepath: 'src' }); // specific tree
```

#### Network

##### walkRemote — List Remote Objects
```typescript
await git.walkRemote({
  fs, dir,
  url: 'https://...',
  oids: ['...', '...'],
});
```

### 1.3 Plugin System

```typescript
import { plugins } from 'isomorphic-git';

// Custom HTTP plugin
plugins.set('http', {
  request: async ({ url, method, headers, body }) => {
    const response = await fetch(url, { method, headers, body });
    return {
      url: response.url,
      method: response.method,
      headers: Object.fromEntries(response.headers),
      body: [new Uint8Array(await response.arrayBuffer())],
      statusCode: response.status,
      statusMessage: response.statusText,
    };
  },
});

// Custom FS plugin (you passed fs directly, but for default)
// Note: In v1.x, fs is always passed explicitly. Plugins are for http, emitter, etc.

// Emitter (for events)
plugins.set('emitter', {
  emit: (event, data) => {
    console.log(`[git] ${event}:`, data);
  },
});
```

### 1.4 Events
```typescript
// git emits events via optional `emitter` in fs:
// event: 'message', 'progress', 'push', 'fetch', 'clone', 'checkout'
// Use onProgress callback to receive progress updates
```

---

## 2. Simple-Git (steveukx/git-js)
**GitHub**: https://github.com/steveukx/git-js | **Stars**: 6K+
**npm**: `simple-git` | **Weekly**: ~13M | **License**: MIT

### 2.1 Installation
```bash
npm install simple-git
```

### 2.2 Complete API

#### Setup
```typescript
import simpleGit from 'simple-git';

// Basic
const git = simpleGit();

// With options
const git = simpleGit({
  baseDir: '/path/to/repo',
  binary: 'git',
  maxConcurrentProcesses: 6,
  trimmed: false,            // trim trailing whitespace
});

// With working directory
const git = simpleGit('/path/to/repo');

// Create from existing
const git = simpleGit().cwd('/path/to/sub/dir');

// Spawn options
const git = simpleGit({
  spawnOptions: {
    stdio: ['ignore', 'pipe', 'pipe'],
    env: { ...process.env, GIT_SSH: '/usr/bin/ssh' },
  },
});
```

#### Status
```typescript
const status = await git.status();
// {
//   not_added: string[],
//   conflicted: string[],
//   created: string[],
//   deleted: string[],
//   modified: string[],
//   renamed: string[],
//   files: FileStatusResult[],
//   staged: string[],
//   ahead: number,
//   behind: number,
//   current: string | null,   // current branch
//   tracking: string | null,  // tracking branch
//   detached: boolean,
//   isClean: () => boolean,
// }

// Quick check
const clean = await git.status().then(s => s.isClean());
```

#### Log
```typescript
const log = await git.log();
// {
//   all: DefaultLogFields[],
//   latest: DefaultLogFields | null,
//   total: number,
// }

// With options
const log = await git.log({
  from: 'main',
  to: 'feature',
  maxCount: 10,
  multiLine: true,
  symmetric: false,
  file: 'src/index.ts',
  format: { date: 'YYYY-MM-DD', hash: '%H', message: '%s' },
  number: 20,
  splitter: '----',
  // OR custom format:
  '--diff-filter': 'M',      // only modified
  since: '2025-01-01',
  until: '2025-06-01',
});

// log.all[0]:
// {
//   hash: string,
//   date: string,
//   message: string,
//   refs: string,
//   body: string,
//   author_name: string,
//   author_email: string,
// }

// Pagination
const commits = await git.log(['--max-count=50']);
const nextPage = await git.log(['--skip=50', '--max-count=50']);
```

#### Branches
```typescript
// List branches
const branches = await git.branch();
// {
//   all: string[],
//   branches: { [name]: BranchSummary },
//   current: string,
//   detached: boolean,
// }

// With options
const branches = await git.branch(['-a']);           // all branches
const branches = await git.branch(['-r']);           // remote branches
const merged = await git.branch(['--merged']);

// Local branches only
const local = await git.branchLocal();

// Branch details
const branchSummary = await git.branch();
branchSummary.branches['main'];
// { current: true, name: 'main', commit: 'abc...', label: 'main' }

// Create branch
await git.branch(['feature/x']);

// Delete branch
await git.branch(['-D', 'feature/old']);

// Rename branch
await git.branch(['-m', 'old-name', 'new-name']);
```

#### Checkout
```typescript
await git.checkout('main');
await git.checkout('feature/new');
await git.checkout('-b', 'feature/new');      // create + checkout
await git.checkoutLocalBranch('feature/new');  // create local branch
await git.checkoutBranch('feature/new', 'origin/feature/new'); // track remote
await git.checkout('commit_sha');               // detached HEAD
```

#### Add / Commit
```typescript
// Stage specific files
await git.add('src/index.ts');
await git.add(['src/a.ts', 'src/b.ts']);

// Stage all
await git.add('.');
await git.add('./*');

// Stage with options
await git.add('.', '-f');   // force

// Commit
const result = await git.commit('feat: message');
// { commit: 'sha', branch: 'main', summary: { changes: 10, insertions: 5, deletions: 3 } }

// Commit with options
await git.commit('feat: message', 'src/index.ts');   // specific files
await git.commit('feat: message', { '--no-verify': null });  // skip hooks
await git.commit('amend: message', { '--amend': null, '--no-edit': null });

// Amend
await git.commit('', { '--amend': null, '--no-edit': null });
```

#### Diff
```typescript
const diff = await git.diff();
// plain text diff

const diffSummary = await git.diffSummary();
// { files: [{ file: string, insertions: number, deletions: number, changes: number, binary: boolean }], insertions: number, deletions: number, changes: number }

// Options
await git.diff(['--cached']);              // staged diff
await git.diff(['main', 'feature']);      // between branches
await git.diff(['HEAD~1', 'HEAD']);       // last commit
await git.diff(['--name-only']);          // only filenames
await git.diffSummary(['main', 'feature']);
```

#### Remote Operations

##### clone
```typescript
const result = await git.clone('https://github.com/user/repo.git');
const result = await git.clone('https://...', '/local/path');
const result = await git.clone('https://...', './repo', {
  '--depth': 1,
  '--branch': 'main',
  '--single-branch': null,
});
```

##### pull
```typescript
const result = await git.pull();
// { remote: string[], branch: Summary, ... }

// With options
await git.pull('origin', 'main');
await git.pull('origin', 'main', {
  '--rebase': null,
  '--ff-only': null,
});
```

##### push
```typescript
await git.push();
await git.push('origin', 'main');
await git.push('origin', 'feature', {
  '--force': null,
  '--set-upstream': null,
  '--tags': null,
});
await git.push(['--delete', 'origin', 'feature/old']);
```

##### fetch
```typescript
await git.fetch();
await git.fetch('origin');
await git.fetch('origin', 'main');
await git.fetch({ '--depth': 1, '--prune': null });
await git.fetch(['--all']);             // all remotes
```

##### remote
```typescript
// List remotes
const remotes = await git.getRemotes(true);
// [{ name: 'origin', refs: { fetch: '...', push: '...' } }]

// Add remote
await git.addRemote('upstream', 'https://...');

// Remove remote
await git.removeRemote('origin');

// Update remote URL
await git.remote(['set-url', 'origin', 'https://...']);
```

#### Merge
```typescript
// Merge branch
await git.mergeFromTo('feature', 'main');

// Or simpler:
await git.merge(['feature']);

// Merge options
await git.merge(['--no-ff', 'feature']);
await git.merge(['--squash', 'feature']);
await git.merge(['--abort']);   // abort conflicted merge
await git.merge(['--continue']);

// Check merge
const mergeResult = await git.merge(['--no-commit', '--no-ff', 'feature']);
// mergeResult = { result: string?, updates: {}, commits: string[] }
```

#### Rebase
```typescript
await git.rebase(['main']);
await git.rebase(['--onto', 'main', 'feature']);
await git.rebase(['--abort']);
await git.rebase(['--continue']);
await git.rebase(['--skip']);
```

#### Reset
```typescript
await git.reset(['--soft', 'HEAD~1']);     // keep changes staged
await git.reset(['--mixed', 'HEAD~1']);    // keep changes unstaged
await git.reset(['--hard', 'HEAD~1']);     // discard changes
await git.reset(['HEAD']);                  // unstage all
await git.reset(['--hard', 'origin/main']); // reset to remote
```

#### Stash
```typescript
await git.stash();                          // push stash
await git.stash(['push', '-m', 'message']);
await git.stash(['pop']);                   // apply + drop
await git.stash(['apply']);                 // apply only
await git.stash(['drop']);                  // drop latest
await git.stash(['drop', 'stash@{0}']);     // drop specific
await git.stash(['list']);                  // list stashes
await git.stash(['clear']);                 // clear all
await git.stash(['branch', 'feature/x']);   // branch from stash
const stashList = await git.stash(['list']);
```

#### Tag
```typescript
// Create tag
await git.addTag('v1.0.0');
await git.addTag('v1.0.0', 'Release message');  // annotated tag
await git.tag(['--delete', 'v1.0.0']);          // delete tag
await git.tag(['--list', '--sort=-v:refname']); // list sorted
const tags = await git.tags();
// { all: ['v1.0.0', 'v1.1.0'], latest: 'v1.1.0' }

// Push tags
await git.pushTags('origin');
```

#### Blame
```typescript
const blame = await git.blame({
  file: 'src/index.ts',
  // optional revision:
  revision: 'main',
  // optional starting line:
  startingLine: 1,
  // ending line:
  endingLine: 10,
});

// Returns:
// {
//   commits: {
//     'sha1': { author, email, date, message, summary, line, ... },
//   },
//   lines: { lineNumber: 'sha1', ... }
// }
```

#### Submodules
```typescript
await git.subModule(['init']);
await git.subModule(['update']);
await git.subModule(['add', 'https://...', 'lib/mylib']);
await git.subModule(['foreach', 'git checkout main']);
```

#### Configuration
```typescript
// Read config
const username = await git.raw(['config', 'user.name']);
const email = await git.raw(['config', 'user.email']);

// Set config
await git.addConfig('user.name', 'My Name');
await git.addConfig('user.email', 'me@email.com', true); // global

// List config
const configList = await git.listConfig();
// { all: { 'user.name': '...' }, values: { files: [...] } }
```

#### Raw Commands
```typescript
// Run any git command
const result: string = await git.raw(['log', '--oneline', '-5']);
const result: Buffer = await git.binaryCatFile(['commit', 'sha']);
const result: string = await git.raw([
  'diff', '--name-only', 'HEAD~1', 'HEAD',
]);
```

#### Revert
```typescript
await git.revert(['HEAD']);            // revert last commit
await git.revert(['--no-commit', 'sha']);
await git.revert(['main~3..main~2']);
```

#### Cherry-pick
```typescript
await git.cherryPick(['commit_sha']);
await git.cherryPick(['--no-commit', 'commit_sha']);
await git.cherryPick(['main~1', 'main~3']);
await git.cherryPick(['--abort']);
await git.cherryPick(['--continue']);
```

#### Clean
```typescript
await git.clean('f');            // -f: force
await git.clean('f', ['-d']);    // -d: directories
await git.clean('f', ['-d', '-x']); // -x: ignored too
await git.clean(CleanOptions.FORCE + CleanOptions.RECURSIVE);
```

#### Rewriting History
```typescript
// Not directly supported — use raw:
await git.raw(['filter-branch', '--force', '--index-filter', '...']);
await git.raw(['rebase', '-i', '--root']);
```

### 2.3 Error Handling
```typescript
import simpleGit, { GitError, GitResponseError, GitConstructError } from 'simple-git';

try {
  await git.clone('https://invalid-url');
} catch (error) {
  if (error instanceof GitResponseError) {
    console.error('Git error:', error.message);
    // error.git = { ... raw git output }
  }
  if (error.git?.exitCode === 128) {
    // Authentication error
  }
}
```

### 2.4 Events
```typescript
const git = simpleGit()
  .outputHandler((command, stdout, stderr) => {
    stdout.on('data', (data) => console.log('out:', data));
    stderr.on('data', (data) => console.warn('err:', data));
  });
```

### 2.5 TypeScript
```typescript
import simpleGit, {
  SimpleGit,
  StatusResult,
  BranchSummary,
  LogResult,
  DiffResult,
  MergeResult,
  ResetMode,
  CleanOptions,
} from 'simple-git';

const git: SimpleGit = simpleGit();
const status: StatusResult = await git.status();
const log: LogResult = await git.log();

// Reset modes
await git.reset(ResetMode.HARD, 'HEAD~1');
await git.reset(ResetMode.SOFT, 'HEAD~1');

// Clean options
const clean = CleanOptions.FORCE + CleanOptions.RECURSIVE;
await git.clean(clean);

// Type guards
if (status.isClean()) {
  console.log('Clean working tree');
}
```
