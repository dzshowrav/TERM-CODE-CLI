# File System, Search & Code Libraries — Complete A-to-Z Reference

---

## 1. Fast-Glob (mrmlnc/fast-glob)
**GitHub**: https://github.com/mrmlnc/fast-glob | **Stars**: 9K+
**npm**: `fast-glob` | **Weekly**: ~67M | **License**: MIT

### 1.1 Installation
```bash
npm install fast-glob
```

### 1.2 Complete API

#### Main Export: `fg()`
```typescript
import fg from 'fast-glob';

// Sync
const entries: string[] = fg.sync(['**/*.ts', '!node_modules'], options?);
fg.sync(patterns, options);

// Async
const entries: string[] = await fg(['**/*.ts', '!node_modules'], options?);
fg.async(patterns, options);

// Stream
const stream: NodeJS.ReadableStream = fg.stream(patterns, options?);
stream.on('data', (entry: string) => {});
stream.on('end', () => {});
stream.on('error', (err) => {});
```

#### Options (all types)

```typescript
interface FastGlobOptions {
  // === Patterns ===
  cwd?: string;                          // Current working dir (default: process.cwd())
  deep?: number | Infinity;              // Max recursion depth
  ignore?: string[];                      // Additional ignore patterns

  // === File System ===
  dot?: boolean;                          // Include dotfiles (default: false)
  onlyFiles?: boolean;                    // Only files, no dirs (default: true)
  onlyDirectories?: boolean;              // Only directories
  followSymbolicLinks?: boolean;          // Follow symlinks (default: true)
  stats?: boolean;                        // Return fs.Stats (default: false)
  markDirectories?: boolean;              // Append / to directories
  objectMode?: boolean;                   // Return Entry objects
  unique?: boolean;                       // Dedup entries (default: true)
  braceExpansion?: boolean;               // Enable brace expansion (default: true)
  caseSensitiveMatch?: boolean;           // Case sensitive (default: true)

  // === Performance ===
  concurrency?: number;                   // Worker threads (default: os.cpus().length)
  baseNameMatch?: boolean;                // Match basename, not path (default: false)
  transform?: (entry: string) => any;     // Transform each entry
  suppressErrors?: boolean;               // Suppress EACCES/EPERM (default: false)

  // === Git & Ignore ===
  absolute?: boolean;                     // Absolute paths (default: false)
  errorOnBrokenSymbolicLinks?: boolean;   // Error on broken symlinks
  throwErrorOnBrokenSymbolicLinks?: boolean;
}

// When stats: true, returns:
interface Entry {
  path: string;
  depth: number;
  stats?: fs.Stats;
}

// When objectMode: true, returns Entry objects with stats
```

#### Pattern Syntax
```typescript
// Wildcards
'*'                 // any file in cwd
'*.ts'              // .ts files in cwd
'**/*.ts'           // .ts files recursive
'src/*.ts'          // .ts files in src/

// Brace expansion
'*.{ts,js}'         // .ts or .js files
'src/{a,b}/**'      // src/a/** or src/b/**

// Negation
'!node_modules'     // exclude node_modules
'!**/*.test.ts'     // exclude test files

// Extglob
'@(foo|bar)'        // foo or bar
'!(foo|bar)'        // not foo or bar
'?(foo|bar)'        // zero or one
'*(foo|bar)'        // zero or more
'+(foo|bar)'        // one or more

// Globstars
'src/**'            // src and all subdirs
'src/**/index.ts'   // any index.ts under src
```

#### Error Handling
```typescript
try {
  const files = await fg('**/*');
} catch (error) {
  // Error is FastGlobError
}

// Suppress permission denied
const files = await fg('**/*', { suppressErrors: true });

// Broken symlinks
const files = await fg('**/*', {
  followSymbolicLinks: true,
  errorOnBrokenSymbolicLinks: false,
});
```

#### Helper: `fg.isDynamicPattern()`
```typescript
fg.isDynamicPattern('*.ts');        // true
fg.isDynamicPattern('index.ts');    // false
fg.isDynamicPattern('**/*.ts');     // true
fg.isDynamicPattern('node_modules', { caseSensitiveMatch: false }); // false
```

#### Helper: `fg.escapePath()`
```typescript
fg.escapePath('path/to/[file].ts');  // path/to/[[]file[]].ts (escapes glob chars)
```

#### Helper: `fg.convertPathToPattern()`
```typescript
fg.convertPathToPattern('/path/to/special{file}.ts');
// /path/to/special{file}.ts → escapes special chars
```

#### Task System (internal)
```typescript
import fg from 'fast-glob';
const tasks = fg.generateTasks('**/*.ts', { cwd: '/src' });
// tasks: [{ dynamic: true, patterns: [...], base: '/src/dir' }]
```

---

## 2. Globby (sindresorhus/globby)
**GitHub**: https://github.com/sindresorhus/globby | **Stars**: 2.5K+
**npm**: `globby` | **Weekly**: ~62M | **License**: MIT

### 2.1 Installation
```bash
npm install globby
```

### 2.2 Complete API
```typescript
import globby from 'globby';
// Powered by fast-glob with gitignore integration
```

#### globby()
```typescript
// Async
const paths: string[] = await globby('*.ts');
const paths: string[] = await globby(['**/*.ts', '!node_modules']);

// With gitignore support
const paths: string[] = await globby('**/*', { gitignore: true });
```

#### globby.sync()
```typescript
const paths: string[] = globby.sync('**/*.ts', options?);
```

#### globby.stream()
```typescript
const stream = globby.stream('**/*.ts');
stream.on('data', (path) => {});
```

#### globby.generateGlobTasks()
```typescript
const tasks = await globby.generateGlobTasks('**/*.ts', options?);
// [{ pattern: '**/*.ts', options: { ... } }]
```

#### globby.hasMagic()
```typescript
globby.hasMagic('**/*.ts');    // true
globby.hasMagic('index.ts');   // false
```

#### globby.gitignore()
```typescript
const isIgnored = await globby.gitignore({ cwd: '/project' });
// isIgnored({ path: 'node_modules/test' }) → true
```

#### glomatch — Pattern Matching
```typescript
import { isMatch } from 'globby';
// or:
import { isGitIgnored } from 'globby';

const isIgnored = await isGitIgnored({ cwd: '/project' });
```

### 2.3 Options
Supports all FastGlob options plus:

```typescript
interface GlobbyOptions extends FastGlobOptions {
  gitignore?: boolean;        // respect .gitignore (default: false)
  expandDirectories?: boolean | ExpandDirectoriesOption;
  // ExpandDirectoriesOption:
  //   true               - expand all
  //   false              - no expansion
  //   { files: string[], extensions: string[] }
}
```

### 2.4 Convenience: isGitIgnored / isMatch
```typescript
import { isGitIgnored } from 'globby';

// Check if a path is gitignored
const check = isGitIgnored();
check('/absolute/path/to/file');       // Promise<boolean>

// With custom cwd
const check = isGitIgnored({ cwd: '/project' });
```

---

## 3. Ignore (kaelzhang/node-ignore)
**GitHub**: https://github.com/kaelzhang/node-ignore | **Stars**: 6K+
**npm**: `ignore` | **Weekly**: ~84M | **License**: MIT

### 3.1 Installation
```bash
npm install ignore
```

### 3.2 Complete API

#### Create Instance
```typescript
import ignore from 'ignore';

const ig = ignore();
const ig2 = ignore({ ignorecase: true });  // case insensitive (default: false)
const ig3 = ignore({ allowRelativePaths: true }); // allow relative (default: true)
```

#### .add()
```typescript
// Add from string
ig.add('.gitignore');
ig.add('node_modules');
ig.add('*.log');

// Add from array
ig.add(['.env', '.DS_Store', '*.tmp']);

// Add from string with \n
ig.add('node_modules\ndist\n*.log');

// Add from .gitignore content
ig.add(fs.readFileSync('.gitignore', 'utf8'));

// Returns the instance (chainable)
ig.add('pattern').add('another');
```

#### .ignores()
```typescript
// Check if path is ignored
const result: boolean = ig.ignores('node_modules/foo.js');      // true
const result: boolean = ig.ignores('src/index.ts');              // false
const result: boolean = ig.ignores('.gitignore');                // false (unless pattern)
```

#### .filter()
```typescript
// Filter array of paths
const filtered: string[] = ig.filter([
  'src/index.ts',
  'node_modules/foo.js',
  'dist/bundle.js',
]);
// ['src/index.ts']  (non-ignored paths)
```

#### .createFilter()
```typescript
const filter = ig.createFilter();
const filtered: string[] = paths.filter(filter);
// Like .filter() but returns a reusable function
```

#### .test()
```typescript
const result = ig.test('src/index.ts').ignored;     // boolean
const detail = ig.test('node_modules/foo.js');
// {
//   ignored: true,
//   unignored: false,
//   rule: {
//     pattern: 'node_modules',
//     negative: false
//   }
// }
```

#### Negation Patterns
```typescript
ig.add(['*.log', '!important.log']);  // ignore all logs except important.log

ig.ignores('debug.log');          // true
ig.ignores('important.log');      // false (unignored)
```

#### Instance Options
```typescript
const ig = ignore({
  ignorecase: false,          // case-sensitive matching
  allowRelativePaths: true,   // allow non-absolute paths
  dotfiles: true,             // include dotfiles (default: platform ✓)
});

// Platform-specific behavior
// Windows: path.sep handling
// POSIX: forward slashes only
```

---

## 4. Tree-Sitter (tree-sitter/tree-sitter)
**GitHub**: https://github.com/tree-sitter/tree-sitter | **Stars**: 19K+
**npm**: `tree-sitter` (also `web-tree-sitter` for browser)
**License**: MIT | **Website**: https://tree-sitter.github.io/tree-sitter/

### 4.1 Installation
```bash
npm install tree-sitter
# Install language parser for each language:
npm install tree-sitter-javascript
npm install tree-sitter-typescript
npm install tree-sitter-python
# ... etc
```

### 4.2 Complete API

#### Parser
```typescript
import Parser from 'tree-sitter';
import JavaScript from 'tree-sitter-javascript';

// Initialize
const parser = new Parser();

// Set language
parser.setLanguage(JavaScript);

// Parse source code
const tree: Tree = parser.parse('const x = 1;');
// or with callback
const tree = parser.parse((index, offset) => sourceCode.slice(offset));

// Parse with old tree for incremental parsing
const newTree = parser.parse('const x = 2;', tree);

// Get language
const lang = parser.getLanguage();

// Get logger
const logger = parser.getLogger();

// Set logger
parser.setLogger((message, params) => {
  console.log(message, params);
});

// Debugging
parser.printDotGraph();  // prints DOT graph of parser state
```

#### Tree
```typescript
// Get root node
const root: SyntaxNode = tree.rootNode;

// Incremental edits
tree.edit({
  startIndex: 0,
  oldEndIndex: 3,
  newEndIndex: 5,
  startPosition: { row: 0, column: 0 },
  oldEndPosition: { row: 0, column: 3 },
  newEndPosition: { row: 0, column: 5 },
});

// Get changed ranges (compared to another tree)
const ranges: Range[] = tree.getChangedRanges(oldTree);

// Walk tree
const cursor: TreeCursor = tree.walk();

// Clone tree
const clone: Tree = tree.copy();

// Language
const lang = tree.getLanguage();
```

#### SyntaxNode
```typescript
const node: SyntaxNode = rootNode;

// Type info
node.type;                   // string: 'function_declaration', 'variable_declaration', etc.
node.isNamed;                // boolean: true for named nodes
node.isExtra;                // boolean: is extra (e.g. comments)
node.hasChanges;             // boolean: changed since last edit

// Position
node.startIndex;             // number: byte offset start
node.endIndex;               // number: byte offset end
node.startPosition;          // { row: number, column: number }
node.endPosition;            // { row: number, column: number }

// Tree navigation
node.parent;                 // SyntaxNode | null
node.child(index);           // SyntaxNode | null
node.childCount;             // number
node.namedChild(index);      // SyntaxNode | null
node.namedChildCount;        // number
node.children;               // SyntaxNode[] (all children)
node.namedChildren;          // SyntaxNode[] (named only)
node.nextSibling;            // SyntaxNode | null
node.prevSibling;            // SyntaxNode | null
node.nextNamedSibling;       // SyntaxNode | null
node.prevNamedSibling;       // SyntaxNode | null

// Content
node.text;                   // string: source text
node.grammarSymbol;          // number: symbol id in grammar

// Querying
node.firstChildForIndex(index);       // SyntaxNode | null
node.firstNamedChildForIndex(index);  // SyntaxNode | null

// Descendant search
node.descendantForIndex(start, end);       // SyntaxNode
node.namedDescendantForIndex(start, end);  // SyntaxNode
node.descendantForPosition(start, end);    // SyntaxNode
node.namedDescendantForPosition(start, end); // SyntaxNode

// Descendant count
node.descendantCount;        // number (including self)

// Equality
node.equals(other);          // boolean: same position in syntax tree

// toString
node.toString();             // S-expression representation
```

#### TreeCursor
```typescript
const cursor: TreeCursor = tree.walk();

// Navigation
cursor.gotoFirstChild();         // boolean
cursor.gotoLastChild();          // boolean
cursor.gotoNextSibling();        // boolean
cursor.gotoPreviousSibling();    // boolean
cursor.gotoParent();             // boolean
cursor.gotoDescendant(index);    // jump to descendant by index
cursor.gotoFirstChildForIndex(index); // boolean
cursor.gotoFirstChildForPosition(pos);// boolean

// Current node
cursor.nodeType;                 // string
cursor.nodeIsNamed;              // boolean
cursor.currentNode();            // SyntaxNode
cursor.currentFieldName();       // string | null

// Position
cursor.startIndex;               // number
cursor.endIndex;                 // number
cursor.startPosition;            // Point
cursor.endPosition;              // Point

// Cloning
cursor.copy();                  // TreeCursor

// Reset
cursor.reset(node);             // reposition
cursor.delete();                // cleanup
```

#### Language
```typescript
const JS = await Parser.Language.load('/path/to/tree-sitter-javascript.wasm');

// Query language
const query = JS.query('(function_declaration name: (identifier) @fnname)');

// Get node types
const nodeTypes = JS.nodeTypeInfoById;  // Map<number, NodeTypeInfo>

// Field names
const fieldNames = JS.fieldNamesForNodeType(nodeTypeId);
```

#### Query
```typescript
const query = JS.query(`
  (function_declaration
    name: (identifier) @function.name
    parameters: (formal_parameters) @function.params)
  (comment) @comment
`);

// Match captures
const matches = query.matches(rootNode);
// [{ pattern: 0, captures: [{ name: 'function.name', node: SyntaxNode }, ...] }]

// Or iterate
query.captures(rootNode, (captures) => {
  for (const capture of captures) {
    console.log(capture.name, capture.node.text);
  }
});

// Predicates
// #eq?, #match?, #any-of?, #set!, #is?, #not?, #is-not?
const query = JS.query(`
  (call_expression
    function: (identifier) @fn
    (#eq? @fn "require"))
`);

// Range queries
query.matches(rootNode, {
  startPosition: { row: 10, column: 0 },
  endPosition: { row: 20, column: 0 },
  startIndex: 500,
  endIndex: 1000,
});

query.captures(rootNode, (captures) => {
  // process captures
}, {
  startPosition: { row: 10, column: 0 },
  endPosition: { row: 20, column: 0 },
});
```

#### Range Type
```typescript
interface Range {
  startIndex: number;
  endIndex: number;
  startPosition: Point;
  endPosition: Point;
}

interface Point {
  row: number;
  column: number;
}
```

### 4.3 Available Language Parsers
```
tree-sitter-javascript         JS / JSX
tree-sitter-typescript         TypeScript (.ts, .tsx)
tree-sitter-python             Python
tree-sitter-rust               Rust
tree-sitter-go                 Go
tree-sitter-java               Java
tree-sitter-c                  C
tree-sitter-cpp                C++
tree-sitter-c-sharp            C#
tree-sitter-ruby               Ruby
tree-sitter-php                PHP
tree-sitter-swift              Swift
tree-sitter-json               JSON
tree-sitter-yaml               YAML
tree-sitter-toml               TOML
tree-sitter-html               HTML
tree-sitter-css                CSS
tree-sitter-bash               Bash
tree-sitter-markdown           Markdown
tree-sitter-sql                SQL
tree-sitter-regex              Regex
...and 100+ more
```

### 4.4 Language Methods
```typescript
// After setLanguage(JS), the language object provides:
const lang = parser.getLanguage();

// Query
const query = lang.query('(identifier) @id');

// Node type info
lang.nodeTypeInfoById;         // all node types
lang.nodeTypeInfoForNode(node); // get type info for node

// Field info
lang.fields;                   // field names map

// Next state
lang.nextState(node, stateId); // for state machine
```
