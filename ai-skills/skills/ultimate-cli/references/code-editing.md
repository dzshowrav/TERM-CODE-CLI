# Code Manipulation Libraries — Complete A-to-Z Reference

---

## 1. Babel (babel/babel)
**GitHub**: https://github.com/babel/babel | **Stars**: 43K+
**npm**: `@babel/core` | **Weekly**: ~45M | **License**: MIT
**Website**: https://babeljs.io/docs/

### 1.1 Installation
```bash
npm install @babel/core @babel/parser @babel/traverse @babel/generator @babel/types
```

### 1.2 Complete API

#### @babel/parser — Parse Code to AST
```typescript
import { parse, parseAsync, parseExpression } from '@babel/parser';

// Parse full file
const ast = parse('const x = 1;', {
  sourceType: 'module',          // 'module' | 'script' | 'unambiguous'
  sourceFilename?: string,
  startLine?: number,            // default: 1
  startColumn?: number,          // default: 0
  strictMode?: boolean,
  allowImportExportEverywhere?: boolean,
  allowReturnOutsideFunction?: boolean,
  allowSuperOutsideMethod?: boolean,
  allowUndeclaredExports?: boolean,
  errorRecovery?: boolean,        // continue on error (default: false, v7+)
  createParenthesizedExpressions?: boolean, // include parens nodes
  plugins: [
    // Syntax plugins:
    'jsx',
    'typescript',
    'flow',
    'decorators',
    'decorators-legacy',
    'classProperties',
    'classPrivateProperties',
    'classPrivateMethods',
    'optionalChaining',
    'nullishCoalescingOperator',
    'dynamicImport',
    'importMeta',
    'topLevelAwait',
    'asyncGenerators',
    'objectRestSpread',
    'optionalCatchBinding',
    'exportNamespaceFrom',
    'exportDefaultFrom',
    'pipelineOperator',
    'throwExpressions',
    'logicalAssignment',
    'numericSeparator',
    'importAssertions',
    'moduleStringNames',
    'destructuringPrivate',
    // Proposed:
    ['decorators', { decoratorsBeforeExport: true }],
    ['pipelineOperator', { proposal: 'minimal' }],
    ['recordAndTuple', { syntaxType: 'hash' }],
  ],
  tokens?: boolean,               // include token array
});

// Parse single expression
const expr = parseExpression('a + b', { plugins: ['typescript'] });

// Async parsing
const ast = await parseAsync(code, options);
```

#### @babel/types — AST Node Builders
```typescript
import * as t from '@babel/types';

// === Literals ===
t.stringLiteral('hello');              // StringLiteral
t.numericLiteral(42);                  // NumericLiteral
t.booleanLiteral(true);                // BooleanLiteral
t.nullLiteral();                       // NullLiteral
t.regExpLiteral('\\w+', 'g');         // RegExpLiteral
t.templateLiteral(quasis, expressions);
t.bigIntLiteral('9007199254740991');   // BigIntLiteral

// === Identifiers & References ===
t.identifier('name');                  // Identifier
t.identifier('myVar');
t.identifier('$');                     // special props

// === Expressions ===
t.assignmentExpression('=', left, right);   // AssignmentExpression
t.binaryExpression('+', left, right);       // BinaryExpression
t.unaryExpression('!', argument);           // UnaryExpression
t.callExpression(callee, args);             // CallExpression
t.newExpression(callee, args);              // NewExpression
t.memberExpression(object, property, computed?);
t.optionalMemberExpression(object, property, computed?, optional);
t.optionalCallExpression(callee, args, optional);
t.conditionalExpression(test, consequent, alternate);
t.logicalExpression('&&', left, right);
t.sequenceExpression(expressions);
t.arrowFunctionExpression(params, body, async?);
t.functionExpression(id, params, body, generator?, async?);
t.objectExpression(properties);
t.arrayExpression(elements);
t.yieldExpression(argument, delegate?);
t.awaitExpression(argument);
t.spreadElement(argument);
t.taggedTemplateExpression(tag, quasi);

// === Statements ===
t.expressionStatement(expression);
t.blockStatement(body);
t.returnStatement(argument?);
t.ifStatement(test, consequent, alternate?);
t.switchStatement(discriminant, cases);
t.throwStatement(argument);
t.tryStatement(block, handler?, finalizer?);
t.forStatement(init, test, update, body);
t.forInStatement(left, right, body);
t.forOfStatement(left, right, body);
t.whileStatement(test, body);
t.doWhileStatement(body, test);
t.breakStatement(label?);
t.continueStatement(label?);
t.debuggerStatement();
t.emptyStatement();

// === Declarations ===
t.variableDeclaration(kind, declarations);  // 'const' | 'let' | 'var'
t.variableDeclarator(id, init?);
t.functionDeclaration(id, params, body, generator?, async?);
t.classDeclaration(id, superClass?, body, decorators?);
t.exportNamedDeclaration(declaration?, specifiers?, source?);
t.exportDefaultDeclaration(declaration);
t.importDeclaration(specifiers, source);
t.importSpecifier(local, imported);
t.importDefaultSpecifier(local);
t.importNamespaceSpecifier(local);

// === JSX ===
t.jsxElement(openingElement, closingElement?, children?, selfClosing?);
t.jsxIdentifier(name);
t.jsxAttribute(name, value?);
t.jsxExpressionContainer(expression);
t.jsxText(value);
t.jsxFragment(openingFragment, closingFragment, children);
t.jsxSpreadAttribute(argument);

// === TypeScript ===
t.tsTypeAnnotation(typeAnnotation);         // wrap with :type
t.tsStringKeyword();                         // string
t.tsNumberKeyword();                         // number
t.tsBooleanKeyword();                        // boolean
t.tsAnyKeyword();                            // any
t.tsUnknownKeyword();                        // unknown
t.tsNeverKeyword();                          // never
t.tsVoidKeyword();                           // void
t.tsNullKeyword();                           // null
t.tsUndefinedKeyword();                      // undefined
t.tsObjectKeyword();                         // object
t.tsArrayType(elementType);                  // Type[]
t.tsUnionType(types);                        // A | B
t.tsIntersectionType(types);                 // A & B
t.tsFunctionType(params, typeAnnotation);    // (p: T) => R
t.tsTypeReference(typeName, typeParameters?);// Type<A, B>
t.tsGenericType(typeParameters);
t.tsInterfaceBody(body);                     // interface body
t.tsInterfaceDeclaration(id, body);          // interface X {}
t.tsTypeAliasDeclaration(id, typeAnnotation); // type X = Y
t.tsEnumDeclaration(id, members);            // enum X {}
t.tsEnumMember(id, initializer?);           // X = 1
t.tsAsExpression(expression, typeAnnotation);// expr as Type
t.tsNonNullExpression(expression);           // expr!
t.tsOptionalCallExpression(callee, args);   // expr?.()
t.tsOptionalProperty(name);                 // property?
t.tsReadonlyKeyword();                      // readonly

// === Utilities ===
t.isIdentifier(node);            // type check
t.isFunctionDeclaration(node);   // type check
t.assertFunctionDeclaration(node); // assert type
t.isNodesEquivalent(a, b);       // deep equality check
t.cloneNode(node);               // deep clone
t.cloneWithoutLoc(node);         // clone without location info
t.removeProperties(node);       // remove location/comment metadata
t.toIdentifier(name);            // convert to valid identifier
t.isValidIdentifier(name);       // check if valid identifier
t.react.isReactComponent(node); // check if React component (class)
t.react.isCompatTag(tag);       // check if HTML tag (vs component)
```

#### @babel/traverse — Walk & Transform AST
```typescript
import traverse, { NodePath, Visitor } from '@babel/traverse';

const visitor: Visitor = {
  // Enter: called when entering node
  enter(path) {
    // `this` = traverse state
  },
  // Exit: called when leaving node
  exit(path) {},

  // Type-specific visitors
  Identifier(path) {
    // Only for Identifier nodes
  },
  FunctionDeclaration(path) {
    // Only for function declarations
  },

  // With type annotation
  'FunctionDeclaration|ClassDeclaration'(path) {},

  // Conditional
  'BinaryExpression'(path) {
    if (path.node.operator === '+') {}
  },
};

traverse(ast, visitor, scope?, state?, path?);

// Alternative: scoped traversal
traverse(ast, {
  enter(path) {
    if (path.isIdentifier({ name: 'test' })) {
      // ...
    }
  },
});
```

#### NodePath API
```typescript
// === Path Information ===
path.node;                    // current AST node
path.parent;                  // parent node
path.parentPath;              // parent NodePath
path.scope;                   // current scope
path.type;                    // node type string
path.key;                     // key in parent (e.g., 'body', 'arguments')
path.listKey;                 // list key if in array
path.container;               // parent container (array or object)
path.inList;                  // boolean: is part of array?
path.depth;                   // depth in tree
path.isStatement;             // boolean: is statement?
path.isExpression;            // boolean: is expression?

// === Type Checking ===
path.isIdentifier(opts?);                 // path.node.type === 'Identifier'
path.isStringLiteral();
path.isFunction();
path.isDeclaration();
path.isScope();                           // introduces new scope?
path.isReferencedIdentifier();            // identifier in reference context?
path.isReferenced();                      // is node referenced?
path.isPure();                            // has no side effects?

// With object filter:
path.isIdentifier({ name: 'require' });

// === Navigation ===
path.get('key');                          // get child path by key
path.get('body.0');                       // nested
path.get('arguments.0');
path.get('callee');
path.getAllNextSiblings();                // NodePath[]
path.getAllPrevSiblings();                // NodePath[]
path.getSibling(key);                     // get sibling by index
path.getStatementParent();                // get enclosing statement
path.getDeepestCommonAncestorFrom(paths); // common ancestor

// === Replacement ===
path.replaceWith(replacementNode);        // replace node
path.replaceWithMultiple(nodes);          // replace with array
path.replaceInline(nodes);                // replace in parent's array
path.replaceExpressionWithStatements(nodes); // expression to statements

// === Insertion ===
path.insertBefore(nodes);                 // insert before
path.insertAfter(nodes);                  // insert after
path.pushContainer(key, nodes);           // push to container array
path.unshiftContainer(key, nodes);        // prepend to container

// === Removal ===
path.remove();                            // remove node
path.prune();                             // remove + avoid dangling comma

// === Scope & Binding ===
path.scope;                               // Scope object
path.bindings;                            // bindings in scope
path.getBinding(name);                    // get binding by name
path.getOwnBinding(name);                 // get binding declared in this scope

// === Comments ===
path.addComment(type, content, line?);    // add comment
path.getComments(type);                   // get comments
path.node.leadingComments;                // comments before node
path.node.trailingComments;               // comments after node
path.node.innerComments;                  // comments inside node

// === Other ===
path.skip();                              // don't traverse children
path.stop();                              // stop traversal entirely
path.removed;                             // boolean: was removed?
path.evaluate();                          // evaluate constant expressions
path.hub;                                 // hub reference
path.requeue();                           // re-queue for re-traversal
path.setData(key, value);                 // store on path
path.getData(key);                        // retrieve stored data
path.visit();                             // visit children
path.matchesPattern(pattern);             // match pattern like "obj.prop"
path.nodeToString();                      // convert to source string
```

#### @babel/generator — AST to Code
```typescript
import generate from '@babel/generator';

const result = generate(ast, options?, code?);
// {
//   code: string,
//   map: object | null    // source map if options.sourceMaps
// }

// Options:
interface GeneratorOptions {
  retainLines?: boolean;
  retainFunctionParens?: boolean;
  comments?: boolean;                    // include comments (default: true)
  compact?: boolean | 'auto';
  concise?: boolean;
  minified?: boolean;
  sourceMaps?: boolean | 'inline' | 'both';
  sourceFileName?: string;
  sourceRoot?: string;
  jsescOption?: object;                  // JSON.stringify options
  auxiliaryCommentBefore?: string;
  auxiliaryCommentAfter?: string;
  shouldPrintComment?: (comment: string) => boolean;
  keepImportAttributes?: boolean;
  retainImportAttributes?: boolean;
  filename?: string;
}
```

#### @babel/code-frame — Pretty Error Messages
```typescript
import { codeFrameColumns } from '@babel/code-frame';

const result = codeFrameColumns(rawLines, location, options?);

// Raw source + location:
const result = codeFrameColumns(`const x = 1\nconst y = 2\n`, {
  start: { line: 1, column: 1 },
  end: { line: 1, column: 12 },
});

// Returns:
// > 1 | const x = 1
//     | ^^^^^^^^^^^

interface CodeFrameOptions {
  linesAbove?: number;      // context lines above (default: 2)
  linesBelow?: number;      // context lines below (default: 2)
  forceColor?: boolean;     // force ANSI colors
  highlightCode?: boolean;  // highlight with color (default: true)
  message?: string;         // error message
}
```

---

## 2. TS-Morph (dsherret/ts-morph)
**GitHub**: https://github.com/dsherret/ts-morph | **Stars**: 5.5K+
**npm**: `ts-morph` | **Weekly**: ~1.5M | **License**: MIT

### 2.1 Installation
```bash
npm install ts-morph
```

### 2.2 Complete API

#### Project
```typescript
import { Project, SyntaxKind, Node, SourceFile } from 'ts-morph';

// === Create Project ===
const project = new Project();

// Project options:
const project = new Project({
  tsConfigFilePath: 'tsconfig.json',
  // or manual compiler options:
  compilerOptions: {
    target: ScriptTarget.ES2020,
    module: ModuleKind.ESNext,
    strict: true,
    jsx: JsxEmit.React,
    declaration: true,
    outDir: './dist',
    rootDir: './src',
    paths: {
      '@/*': ['./src/*'],
    },
    baseUrl: '.',
  },
  // File system (for tests):
  fileSystem: new VirtualFileSystem(),
  // Resolution:
  resolutionHost: (moduleResolutionHost) => resolvedHost,
  // Skip loading:
  skipLoadingLibFiles: false,
  // Manipulation settings:
  manipulationSettings: {
    indentationText: IndentationText.TwoSpaces, // or FourSpaces, EightSpaces, Tab
    newLineKind: NewLineKind.LineFeed,  // or CarriageReturnLineFeed
    insertSpaceAfterOpeningAndBeforeClosingNonemptyBraces: true,
    // ... many more formatting options
  },
});

// === Add Source Files ===
const sourceFile = project.addSourceFileAtPath('src/index.ts');
const files = project.addSourceFileAtPaths('src/**/*.ts');
const newFile = project.createSourceFile('src/new.ts', 'export const x = 1;', { overwrite: true });
const [src] = project.addSourceFilesFromTsConfig('tsconfig.json');

// === Get Source Files ===
const files = project.getSourceFiles();
const file = project.getSourceFile('src/index.ts');
const fileOrUndefined = project.getSourceFileOrThrow('src/index.ts');

// === Compile & Emit ===
project.emit();                   // emit all files
project.emitToMemory();           // emit to in-memory
project.emitSync();               // sync emit
const diagnostics = project.getPreEmitDiagnostics();

// === Type Checker ===
const typeChecker = project.getTypeChecker();
const type = typeChecker.getTypeAtLocation(node);
const symbol = typeChecker.getSymbolAtLocation(node);
```

#### SourceFile
```typescript
const sourceFile = project.getSourceFileOrThrow('src/index.ts');

// === Read ===
sourceFile.getText();                      // full text
sourceFile.getFullText();                  // with leading trivia
sourceFile.getFilePath();                  // file path
sourceFile.getBaseName();                  // file name
sourceFile.getExtension();                 // .ts
sourceFile.isDeclarationFile();            // .d.ts
sourceFile.getLanguageVariant();           // Standard, JSX
sourceFile.getReferencingSourceFiles();    // files that import this
sourceFile.getReferencedSourceFiles();     // files that this imports

// === Statements ===
sourceFile.getStatements();
sourceFile.getStatementByKind(SyntaxKind.VariableStatement);
sourceFile.getExportDeclarations();
sourceFile.getImportDeclarations();
sourceFile.getClass('MyClass');
sourceFile.getClassOrThrow('MyClass');
sourceFile.getClasses();
sourceFile.getFunction('myFunc');
sourceFile.getFunctions();
sourceFile.getInterface('MyInterface');
sourceFile.getInterfaces();
sourceFile.getTypeAlias('MyType');
sourceFile.getTypeAliases();
sourceFile.getEnum('MyEnum');
sourceFile.getEnums();
sourceFile.getVariableDeclaration('myVar');
sourceFile.getVariableDeclarations();

// === Write ===
sourceFile.addImportDeclaration({
  moduleSpecifier: './utils',
  namedImports: ['helper'],
  defaultImport: 'utils',
});

sourceFile.addExportDeclaration({
  moduleSpecifier: './types',
  namedExports: ['User'],
});

sourceFile.addStatements('console.log("hello");');
sourceFile.insertStatements(0, '// top comment');
sourceFile.replaceWithText('const y = 2;');
sourceFile.remove();

// === Manipulation ===
sourceFile.formatText();                   // apply formatting
sourceFile.indent();                       // indent
sourceFile.unindent();                     // unindent
sourceFile.save();                         // save to disk
sourceFile.saveSync();
sourceFile.copy('path/to/copy.ts');        // copy
sourceFile.move('path/to/new.ts');        // move/rename
```

#### Classes
```typescript
const classDec = sourceFile.getClassOrThrow('MyClass');

// === Read ===
classDec.getName();                    // MyClass
classDec.getExtends();                 // base class
classDec.getImplements();              // interfaces
classDec.getProperties();
classDec.getProperty('name');
classDec.getMethods();
classDec.getMethod('getName');
classDec.getConstructors();
classDec.getDecorators();
classDec.getTypeParameters();
classDec.isAbstract();
classDec.isExported();
classDec.isDefaultExport();

// === Add Members ===
classDec.addProperty({
  name: 'age',
  type: 'number',
  initializer: '0',
  scope: Scope.Private,    // Public, Protected, Private
  isReadonly: true,
  hasQuestionToken: true,
  decorators: [{ name: 'inject' }],
  leadingTrivia: '// comment\n',
});

classDec.addMethod({
  name: 'getName',
  returnType: 'string',
  parameters: [{ name: 'prefix', type: 'string' }],
  statements: 'return this.prefix + this.name;',
  scope: Scope.Public,
  isAbstract: false,
  isStatic: false,
  isAsync: true,
  isGenerator: false,
});

classDec.addConstructor({
  parameters: [{ name: 'name', type: 'string', scope: Scope.Private, isReadonly: true }],
});
```

#### Functions
```typescript
const func = sourceFile.getFunctionOrThrow('myFunc');

func.getName();
func.getParameters();
func.getReturnType();
func.getBody();
func.isAsync();
func.isGenerator();
func.isExported();
func.isDefaultExport();

func.addParameter({ name: 'options', type: 'Options' });
func.setReturnType('Promise<void>');
func.remove();
```

#### Import/Export
```typescript
const importDecl = sourceFile.getImportDeclarations()[0];

importDecl.getModuleSpecifier();      // './utils'
importDecl.getDefaultImport();        // default import name
importDecl.getNamespaceImport();      // * as name
importDecl.getNamedImports();         // { a, b, c }

importDecl.addNamedImport('newImport');
importDecl.insertNamedImport(1, 'anotherImport');
importDecl.removeNamedImport('oldImport');
importDecl.setModuleSpecifier('./new-path');
importDecl.remove();
```

#### Type System
```typescript
const type = classDec.getProperty('name')!.getType();

type.getText();                       // 'string'
type.isString();                      // true
type.isNumber();
type.isBoolean();
type.isArray();
type.isClass();
type.isInterface();
type.isUnion();
type.isIntersection();
type.isEnum();
type.isLiteral();
type.isNullable();
type.isUndefined();
type.isAny();
type.isUnknown();
type.isVoid();
type.isNever();
type.isObject();

const unionTypes = type.getUnionTypes();    // for union types
const intersectionTypes = type.getIntersectionTypes();
const arrayElementType = type.getArrayElementType();
const typeArguments = type.getTypeArguments();
```

#### Node (AST)
```typescript
const node: Node = someDeclaration;

// === General ===
node.getKind();              // SyntaxKind enum
node.getKindName();          // 'VariableStatement'
node.getText();              // source text
node.getStart();             // start position
node.getEnd();               // end position
node.getFullStart();         // start with leading trivia

// === Navigation ===
node.getParent();
node.getParentIfOrThrow(predicate);
node.getChildCount();
node.getChildren();
node.getChildAt(index);
node.getFirstChild();
node.getLastChild();
node.getNextSibling();
node.getPreviousSibling();
node.getDescendants();
node.getDescendantAtPos(pos);
node.forEachDescendant((descendant, traversal) => void);

// === Traversal ===
node.getFirstDescendantByKind(SyntaxKind.Identifier);
node.getFirstDescendantByKindOrThrow(kind);
node.getDescendantsOfKind(SyntaxKind.CallExpression);
node.getFirstAncestorByKind(SyntaxKind.FunctionDeclaration);
node.getFirstAncestor(predicate);

// === Type ===
node.getType();
node.getSymbol();
node.getAliasSymbol();

// === Manipulation ===
node.replaceWithText('new text');
node.remove();
node.removeText();
node.removeChildren();

// === Location ===
node.getSourceFile();
node.getFilePath();
node.getLineAndColumnAtPos(pos);
node.getLineNumber();
node.getColumn();

// === Formatting ===
node.formatText();
node.indent();
node.unindent();
```

#### Symbols & References
```typescript
const symbol = sourceFile.getClassOrThrow('MyClass')!.getSymbol()!;

symbol.getName();
symbol.getFullyQualifiedName();
symbol.getDeclarations();
symbol.getDeclaredType();
symbol.getValueDeclaration();

// Find references
const references = symbol.findReferences();
// { sourceFile, references: [{ textSpan, ... }] }

// Rename
symbol.rename('NewName');

// Go to definition
const definitions = symbol.getDefinitions();
```

### 2.3 Code Fixes & Transforms
```typescript
// Language service integration
const diagnostics = project.getPreEmitDiagnostics();
for (const diag of diagnostics) {
  const fixes = diag.getCodeFixes();
  // Apply fix
  fixes[0]?.apply();
}

// Get quick fixes
const sourceFile = project.getSourceFileOrThrow('file.ts');
const quickFixes = sourceFile.getLanguageService().getCodeFixesAtPosition(
  sourceFile.getFilePath(),
  position,
  position,
  []
);
```

---

## 3. JSCodeshift (facebook/jscodeshift)
**GitHub**: https://github.com/facebook/jscodeshift | **Stars**: 9K+
**npm**: `jscodeshift` | **Weekly**: ~1.7M | **License**: MIT

### 3.1 Installation
```bash
npm install jscodeshift
```

### 3.2 Complete API

#### Core API

```typescript
const j = require('jscodeshift');
// or: import j from 'jscodeshift';

// === Parse ===
const root = j(sourceCode);
// equivalent: j(sourceCode, { parser: 'babylon' })

// Parser options:
const root = j(sourceCode, {
  parser: 'babylon',     // default
  // or:
  parser: 'flow',
  parser: 'ts',           // TypeScript
  parser: 'tsx',
  parser: 'babel',
});

// === Output ===
root.toSource();                          // generate code
root.toSource({ quote: 'single' });       // options
root.toSource({ trailingComma: true });
root.toSource({ wrapColumn: 80 });

// === Collection ===
root.find(j.CallExpression);             // find all call expressions
root.find(j.Identifier);                  // find all identifiers
root.find(j.VariableDeclaration);         // find all variable declarations
root.find(j.FunctionDeclaration);
root.find(j.ImportDeclaration);
root.find(j.ClassDeclaration);
root.find(j.ArrowFunctionExpression);
root.find(j.ExportDeclaration);
root.find(j.ConditionalExpression);

// With filter:
root.find(j.CallExpression, {
  callee: { name: 'require' },           // CallExpression where callee.name === 'require'
  arguments: { length: 1 },
});

// Nested filter:
root.find(j.CallExpression, {
  callee: {
    type: 'MemberExpression',
    object: { name: 'console' },
    property: { name: 'log' },
  },
});
```

#### Collection API
```typescript
const collection = root.find(j.CallExpression);

// === Transformation ===
collection.forEach((path) => {
  // path: NodePath
  // path.node: actual AST node
  // path.parentPath: parent
  // path.scope: scope
});

collection.replaceWith((path) => {
  // Each node → one replacement
  return j.identifier('replacement');
});

collection.replaceWith((path) => {
  // Return null/undefined to remove
  // Return array to replace with multiple
  return null;
});

collection.remove();                      // remove all matched nodes

// === Insertion ===
collection.at(0).insertAfter(j.identifier('newNode'));
collection.at(0).insertBefore(j.identifier('preNode'));
path.insertAfter(collection);            // at path level

// === Filtering ===
collection.filter((path) => {
  return path.node.callee?.name === 'test';
});

// === Cloning ===
const clone = collection.clone();

// === Nodes ===
collection.nodes();                       // get all AST nodes
collection.paths();                       // get all NodePaths
collection.size();                        // number of matches
collection.get();                         // get first NodePath
collection.at(index);                     // get at index
```

#### Builder Functions (AST Creators)
```typescript
// === Identifiers ===
j.identifier('name');                     // Identifier('name')
j.identifier.from({ name: 'test' });

// === Literals ===
j.literal('string');                      // StringLiteral
j.literal(42);                            // NumericLiteral
j.literal(true);                          // BooleanLiteral
j.stringLiteral('string');                // StringLiteral (exact)
j.numericLiteral(42);                     // NumericLiteral (exact)
j.nullLiteral();
j.booleanLiteral(true);

// === Expressions ===
j.callExpression(callee, args);           // func(args)
j.memberExpression(obj, prop);            // obj.prop
j.memberExpression(obj, prop, true);      // obj[prop]
j.arrowFunctionExpression(params, body);  // (params) => body
j.functionExpression(id, params, body);   // function() {}
j.assignmentExpression('=', left, right);  // left = right
j.binaryExpression('+', left, right);     // left + right
j.unaryExpression('!', arg);              // !arg
j.conditionalExpression(test, cons, alt); // test ? cons : alt
j.arrayExpression(elements);              // [elements]
j.objectExpression(properties);           // { properties }
j.spreadElement(arg);                     // ...arg
j.templateLiteral(quasis, expressions);   // `str${expr}`
j.yieldExpression(arg);                   // yield arg
j.awaitExpression(arg);                   // await arg
j.sequenceExpression(exprs);              // a, b, c
j.logicalExpression('&&', left, right);  // left && right
j.newExpression(callee, args);            // new Func()
j.thisExpression();                       // this
j.super();                                // super
j.optionalCallExpression(callee, args, optional); // func?.(args)
j.optionalMemberExpression(obj, prop, computed, optional); // obj?.prop

// === Statements ===
j.expressionStatement(expression);         // expr;
j.blockStatement(body);                   // { body }
j.returnStatement(arg?);                  // return arg;
j.ifStatement(test, cons, alt?);         // if (test) {} else {}
j.forStatement(init, test, update, body);
j.forInStatement(left, right, body);
j.forOfStatement(left, right, body);
j.whileStatement(test, body);
j.doWhileStatement(body, test);
j.tryStatement(block, handler, finalizer?);
j.throwStatement(arg);
j.switchStatement(discriminant, cases);
j.breakStatement();
j.continueStatement();
j.debuggerStatement();
j.emptyStatement();

// === Declarations ===
j.variableDeclaration('const', [j.variableDeclarator(id, init)]);
j.functionDeclaration(id, params, body);
j.classDeclaration(id, body);
j.importDeclaration(specifiers, source);
j.exportNamedDeclaration(declaration?, specifiers?, source?);
j.exportDefaultDeclaration(declaration);

// === JSX ===
j.jsxElement(opening, closing, children, selfClosing);
j.jsxIdentifier(name);
j.jsxAttribute(name, value?);
j.jsxExpressionContainer(expr);
j.jsxText(value);
j.jsxSpreadAttribute(arg);
j.jsxFragment(opening, closing, children);

// === TypeScript ===
j.tsTypeAnnotation(typeAnnotation);
j.tsStringKeyword();
j.tsNumberKeyword();
j.tsBooleanKeyword();
j.tsAnyKeyword();
j.tsUnionType(types);
j.tsIntersectionType(types);
j.tsArrayType(elementType);
j.tsFunctionType(params, typeAnnotation);
j.tsTypeReference(typeName, typeParameters?);
j.tsInterfaceBody(body);
j.tsInterfaceDeclaration(id, body);
j.tsTypeAliasDeclaration(id, typeAnnotation);
j.tsEnumDeclaration(id, members);
j.tsAsExpression(expr, typeAnnotation);

// === Patterns ===
j.objectPattern(properties);              // { a, b } destructuring
j.arrayPattern(elements);                 // [a, b] destructuring
j.restElement(arg);                       // ...rest
j.assignmentPattern(left, right);         // a = 1 in params
```

#### Transforms (Module Pattern)
```typescript
// transform.js — for use with jscodeshift CLI
/**
 * @param {import('jscodeshift').FileInfo} file
 * @param {import('jscodeshift').API} api
 * @param {object} options
 */
module.exports = function(file, api, options) {
  const j = api.jscodeshift;
  const root = j(file.source);

  // Transform
  root.find(j.Identifier)
    .filter(path => path.node.name === 'oldName')
    .replaceWith(j.identifier('newName'));

  return root.toSource();
};

// Run with:
// jscodeshift -t transform.js src/
// jscodeshift -t transform.js src/ --parser=ts
// jscodeshift -t transform.js src/ --dry-run
// jscodeshift -t transform.js src/ --verbose=2
```

#### CLI Options
```bash
jscodeshift [options] <path...>

# Common options:
-t, --transform FILE     # transform file
-c, --cpus N             # number of CPUs (default: max)
-d, --dry-run            # dry run (no changes)
-p, --print              # print transformed files
--parser PARSE           # babylon, flow, ts, tsx, babel
--extensions EXT         # file extensions (default: js)
--ignore-pattern PATTERN # ignore pattern (e.g., **/node_modules/**)
--ignore-config FILE     # ignore config file
--run-in-band           # run serially (no worker threads)
--silent                # less output
--verbose N              # verbosity level
--stdin                 # read from stdin
```
