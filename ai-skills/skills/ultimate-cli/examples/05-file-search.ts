#!/usr/bin/env node
/**
 * ============================================================
 *  FILE SEARCH & CODE ANALYSIS — Fast-Glob + Ignore + Tree-Sitter
 * ============================================================
 * Libraries: fast-glob, ignore, tree-sitter, tree-sitter-javascript,
 *            tree-sitter-typescript, chalk, commander
 *
 * A code search and analysis tool:
 *   - Search files by glob pattern
 *   - Search code by Tree-Sitter AST queries
 *   - Count functions, classes, imports in a codebase
 *   - Find dead code / unused exports
 *   - Dependency graph
 *
 * Run: npx ts-node examples/05-file-search.ts analyze src/
 *      npx ts-node examples/05-file-search.ts find "src/**/*.ts" --pattern "function"
 *      npx ts-node examples/05-file-search.ts stats src/
 * ============================================================
 */

import { Command } from 'commander';
import Parser from 'tree-sitter';
import JavaScript from 'tree-sitter-javascript';
import TypeScript from 'tree-sitter-typescript';
import fg from 'fast-glob';
import ignore, { type Ignore } from 'ignore';
import chalk from 'chalk';
import fs from 'fs/promises';
import path from 'path';
import os from 'os';

// ─── Tree-Sitter Setup ──────────────────────────────────
const parser = new Parser();

function getLanguage(filePath: string) {
  const ext = path.extname(filePath);
  if (['.js', '.jsx', '.mjs', '.cjs'].includes(ext)) return JavaScript;
  if (['.ts', '.tsx', '.mts', '.cts'].includes(ext)) return TypeScript;
  return null;
}

// ─── Types ───────────────────────────────────────────────
interface CodeStats {
  files: number;
  lines: number;
  code_lines: number;
  comments: number;
  blanks: number;
  functions: number;
  classes: number;
  interfaces: number;
  imports: number;
  exports: number;
  total_loc: number;
}

interface SearchMatch {
  file: string;
  line: number;
  column: number;
  text: string;
  type: string;
}

interface Dependency {
  file: string;
  imports: string[];
  exports: string[];
}

// ─── Ignore Setup ───────────────────────────────────────
async function loadGitignore(baseDir: string): Promise<Ignore> {
  const ig = ignore();
  ig.add(['node_modules', '.git', 'dist', 'build', 'coverage', '.next']);

  try {
    const gitignore = await fs.readFile(path.join(baseDir, '.gitignore'), 'utf-8');
    ig.add(gitignore);
  } catch {
    // No .gitignore
  }

  return ig;
}

// ─── Stats Command ───────────────────────────────────────
async function analyzeStats(pattern: string, baseDir: string = '.'): Promise<void> {
  const ig = await loadGitignore(baseDir);

  console.log(chalk.bold.cyan('\n  📊 Codebase Analysis\n'));

  const files = await fg(pattern, {
    cwd: baseDir,
    absolute: true,
    ignore: ['**/node_modules/**', '**/.git/**', '**/dist/**'],
    onlyFiles: true,
  });

  // Filter by gitignore
  const filteredFiles = files.filter((f) => {
    const rel = path.relative(baseDir, f);
    return !ig.ignores(rel);
  });

  if (!filteredFiles.length) {
    console.log(chalk.yellow('  No files matched\n'));
    return;
  }

  console.log(chalk.dim(`  Scanning ${filteredFiles.length} files...\n`));

  const stats: CodeStats = {
    files: 0,
    lines: 0,
    code_lines: 0,
    comments: 0,
    blanks: 0,
    functions: 0,
    classes: 0,
    interfaces: 0,
    imports: 0,
    exports: 0,
    total_loc: 0,
  };

  // Per-file stats
  const fileStats: { file: string; lines: number; functions: number }[] = [];

  for (const file of filteredFiles) {
    const lang = getLanguage(file);
    if (!lang) continue;

    const content = await fs.readFile(file, 'utf-8');
    const lines = content.split('\n');
    const lineCount = lines.length;

    stats.files++;
    stats.lines += lineCount;

    // Count blanks & comments
    for (const line of lines) {
      const trimmed = line.trim();
      if (!trimmed) {
        stats.blanks++;
      } else if (trimmed.startsWith('//') || trimmed.startsWith('/*') || trimmed.startsWith('*')) {
        stats.comments++;
      } else {
        stats.code_lines++;
      }
    }

    // Tree-Sitter analysis
    try {
      parser.setLanguage(lang as any);
      const tree = parser.parse(content);
      const root = tree.rootNode;

      // Count functions
      const funcs = root.descendantsOfType?.('function_declaration') || [];
      const arrowFuncs = root.descendantsOfType?.('arrow_function') || [];
      const methods = root.descendantsOfType?.('method_definition') || [];
      stats.functions += funcs.length + arrowFuncs.length + methods.length;

      // Count classes
      const classes = root.descendantsOfType?.('class_declaration') || [];
      stats.classes += classes.length;

      // Count imports
      const imports = root.descendantsOfType?.('import_statement') || [];
      const imports2 = root.descendantsOfType?.('import_declaration') || [];
      stats.imports += imports.length + imports2.length;

      // Count exports
      const exports = root.descendantsOfType?.('export_statement') || [];
      const exports2 = root.descendantsOfType?.('export_declaration') || [];
      stats.exports += exports.length + exports2.length;

      // TypeScript-specific
      if (lang === TypeScript) {
        const ifaces = root.descendantsOfType?.('interface_declaration') || [];
        stats.interfaces += ifaces.length;
      }

      fileStats.push({
        file: path.relative(baseDir, file),
        lines: lineCount,
        functions: funcs.length + arrowFuncs.length + methods.length,
      });
    } catch (err) {
      // Skip parsing errors
    }
  }

  stats.total_loc = stats.code_lines + stats.comments;

  // Print summary
  console.log(chalk.bold('  Summary:\n'));
  console.log(`  ${chalk.dim('Files:')}     ${chalk.bold(String(stats.files))}`);
  console.log(`  ${chalk.dim('Lines:')}     ${chalk.bold(String(stats.lines))}`);
  console.log(`  ${chalk.dim('Code:')}      ${chalk.green(String(stats.code_lines))}`);
  console.log(`  ${chalk.dim('Comments:')}  ${chalk.blue(String(stats.comments))}`);
  console.log(`  ${chalk.dim('Blanks:')}    ${chalk.dim(String(stats.blanks))}`);
  console.log();
  console.log(`  ${chalk.dim('Functions:')} ${chalk.yellow(String(stats.functions))}`);
  console.log(`  ${chalk.dim('Classes:')}   ${chalk.magenta(String(stats.classes))}`);
  if (stats.interfaces > 0) {
    console.log(`  ${chalk.dim('Interfaces:')} ${chalk.cyan(String(stats.interfaces))}`);
  }
  console.log(`  ${chalk.dim('Imports:')}   ${chalk.white(String(stats.imports))}`);
  console.log(`  ${chalk.dim('Exports:')}   ${chalk.white(String(stats.exports))}`);
  console.log();

  // Top 10 largest files
  const topFiles = fileStats.sort((a, b) => b.lines - a.lines).slice(0, 10);
  if (topFiles.length > 0) {
    console.log(chalk.bold('  Top 10 largest files:\n'));
    for (const f of topFiles) {
      console.log(`  ${chalk.dim(String(f.lines).padStart(6))} ${f.file}`);
    }
    console.log();
  }
}

// ─── Find / Search Command ───────────────────────────────
async function searchFiles(
  pattern: string,
  query: string,
  options: { type?: string; context?: number; regex?: boolean; baseDir?: string }
): Promise<void> {
  const ig = await loadGitignore(options.baseDir || '.');
  const baseDir = options.baseDir || '.';

  console.log(chalk.bold.cyan('\n  🔍 Code Search\n'));
  console.log(chalk.dim(`  Pattern: ${pattern}`));
  console.log(chalk.dim(`  Query:   ${query}\n`));

  const files = await fg(pattern, {
    cwd: baseDir,
    absolute: true,
    ignore: ['**/node_modules/**', '**/.git/**', '**/dist/**'],
    onlyFiles: true,
  });

  const filtered = files.filter((f) => !ig.ignores(path.relative(baseDir, f)));

  if (!filtered.length) {
    console.log(chalk.yellow('  No files matched\n'));
    return;
  }

  const matches: SearchMatch[] = [];
  const regex = options.regex ? new RegExp(query, 'gi') : null;

  for (const file of filtered) {
    const content = await fs.readFile(file, 'utf-8');
    const lines = content.split('\n');

    for (let i = 0; i < lines.length; i++) {
      const line = lines[i];
      let match: RegExpExecArray | RegExpMatchArray | null = null;

      if (regex) {
        regex.lastIndex = 0;
        match = regex.exec(line);
      } else if (line.toLowerCase().includes(query.toLowerCase())) {
        match = [query] as unknown as RegExpMatchArray;
      }

      if (match) {
        const col = line.indexOf(match[0] || query) + 1;
        const type = options.type || 'text';

        // Print with context
        const relFile = path.relative(baseDir, file);
        if (options.context && options.context > 0) {
          const start = Math.max(0, i - options.context);
          const end = Math.min(lines.length, i + options.context + 1);

          console.log(chalk.underline(`${relFile}:${i + 1}:${col}`));
          for (let j = start; j < end; j++) {
            const prefix = j === i ? chalk.green('❯') : ' ';
            const lineNum = String(j + 1).padStart(4, ' ');
            let displayLine = lines[j];
            if (j === i) {
              displayLine = displayLine.replace(
                new RegExp(query.replace(/[.*+?^${}()|[\]\\]/g, '\\$&'), 'gi'),
                (m) => chalk.bgYellow.black(m)
              );
            }
            console.log(`  ${prefix} ${chalk.dim(lineNum)} ${displayLine}`);
          }
          console.log();
        } else {
          const truncated = line.length > 100 ? line.slice(0, 97) + '...' : line;
          console.log(`  ${chalk.green(relFile)}:${chalk.yellow(String(i + 1))}:${chalk.cyan(String(col))}`);
          console.log(`    ${truncated.trim()}`);
          console.log();
        }
      }
    }
  }

  console.log(chalk.dim(`  Found ${matches.length} matches in ${files.length} files\n`));
}

// ─── Find Unused Exports ────────────────────────────────
async function findUnused(pattern: string, baseDir: string = '.'): Promise<void> {
  const ig = await loadGitignore(baseDir);

  console.log(chalk.bold.cyan('\n  🗑 Finding Unused Exports\n'));

  const files = await fg(pattern, {
    cwd: baseDir,
    absolute: true,
    ignore: ['**/node_modules/**', '**/.git/**', '**/dist/**'],
    onlyFiles: true,
  });

  // Collect all exports
  const exportMap = new Map<string, string[]>(); // file -> exports[]
  const importMap = new Map<string, Set<string>>(); // import name -> files that use it

  for (const file of files) {
    if (!getLanguage(file)) continue;
    const content = await fs.readFile(file, 'utf-8');

    // Simple regex-based export detection
    const exports: string[] = [];
    const exportRegex = /export\s+(?:default\s+)?(?:function|const|let|var|class|interface|type|enum|async\s+function)\s+(\w+)/g;
    let match;
    while ((match = exportRegex.exec(content)) !== null) {
      exports.push(match[1]);
    }

    // Named exports
    const namedRegex = /export\s*\{\s*([^}]+)\s*\}/g;
    while ((match = namedRegex.exec(content)) !== null) {
      match[1].split(',').forEach((name) => {
        exports.push(name.trim().split(/\s+as\s+/).pop()!);
      });
    }

    if (exports.length > 0) {
      exportMap.set(file, exports);
    }

    // Track imports
    const importRegex = /import\s+(?:\{\s*([^}]+)\s*\}|\*\s+as\s+(\w+)|(\w+))\s+from/g;
    while ((match = importRegex.exec(content)) !== null) {
      const names = [match[1], match[2], match[3]].filter(Boolean);
      for (const name of names) {
        name.split(',').forEach((n) => {
          const trimmed = n.trim().split(/\s+as\s+/).pop()!;
          if (!importMap.has(trimmed)) importMap.set(trimmed, new Set());
          importMap.get(trimmed)!.add(file);
        });
      }
    }
  }

  // Find unused
  let unusedCount = 0;
  for (const [file, exports] of exportMap) {
    for (const exp of exports) {
      const users = importMap.get(exp);
      if (!users || users.size === 0) {
        unusedCount++;
        console.log(`  ${chalk.red('✖')} ${chalk.bold(exp)} ${chalk.dim(`in ${path.relative(baseDir, file)}`)}`);
      }
    }
  }

  if (unusedCount === 0) {
    console.log(chalk.green('  ✅ No unused exports found\n'));
  } else {
    console.log(chalk.yellow(`\n  Found ${unusedCount} potentially unused exports\n`));
  }
}

// ─── CLI ────────────────────────────────────────────────
const program = new Command()
  .name('file-search')
  .description('Code search and analysis with Tree-Sitter')
  .version('1.0.0');

program
  .command('stats')
  .description('Analyze codebase statistics')
  .argument('[pattern]', 'Glob pattern', '**/*.{ts,tsx,js,jsx}')
  .option('--base-dir <dir>', 'Base directory', '.')
  .action((pattern, opts) => analyzeStats(pattern, opts.baseDir));

program
  .command('find')
  .description('Search code for a pattern')
  .argument('<pattern>', 'Glob pattern')
  .argument('<query>', 'Text to search for')
  .option('--regex', 'Use regex search')
  .option('-C, --context <lines>', 'Context lines', parseInt)
  .option('--base-dir <dir>', 'Base directory', '.')
  .action((pattern, query, opts) => searchFiles(pattern, query, opts));

program
  .command('unused')
  .description('Find unused exports')
  .argument('[pattern]', 'Glob pattern', 'src/**/*.{ts,tsx}')
  .option('--base-dir <dir>', 'Base directory', '.')
  .action((pattern, opts) => findUnused(pattern, opts.baseDir));

program.parse(process.argv);
