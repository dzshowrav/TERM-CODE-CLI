#!/usr/bin/env node
/**
 * ============================================================
 *  CODE TRANSFORMER — Babel + TS-Morph + JSCodeshift
 * ============================================================
 * Libraries: @babel/core, @babel/traverse, @babel/types,
 *            ts-morph, jscodeshift, fast-glob, ignore
 *
 * A CLI tool that performs code transformations:
 *   - Convert CommonJS to ESM
 *   - Rename React class components to hooks
 *   - Transform import paths
 *   - Add/remove TypeScript types
 *
 * Run: npx ts-node examples/02-code-transformer.ts convert src/ --cjs-to-esm
 * ============================================================
 */

import { Command } from 'commander';
import * as t from '@babel/types';
import { parse, type ParseResult } from '@babel/parser';
import traverse from '@babel/traverse';
import generate from '@babel/generator';
import chalk from 'chalk';
import fg from 'fast-glob';
import fs from 'fs/promises';
import path from 'path';

// ─── Types ───────────────────────────────────────────────
interface TransformOptions {
  cjsToEsm?: boolean;
  classToHooks?: boolean;
  addReactImport?: boolean;
  renameImports?: [string, string][];
  dryRun?: boolean;
  verbose?: boolean;
}

interface TransformResult {
  file: string;
  transformed: boolean;
  changes: string[];
}

// ─── Babel: CommonJS → ESM Transformer ──────────────────
function transformCJStoESM(code: string): { code: string; changes: string[] } {
  const changes: string[] = [];
  const ast = parse(code, {
    sourceType: 'module',
    plugins: ['jsx', 'typescript', 'decorators-legacy'],
  });

  traverse(ast, {
    // require('x') → import 'x'
    CallExpression(path) {
      const node = path.node;
      if (
        t.isIdentifier(node.callee, { name: 'require' }) &&
        node.arguments.length === 1 &&
        t.isStringLiteral(node.arguments[0])
      ) {
        const source = node.arguments[0].value;
        const parent = path.parent;

        // const x = require('x') → import x from 'x'
        if (t.isVariableDeclarator(parent) && t.isIdentifier(parent.id)) {
          const importDecl = t.importDeclaration(
            [t.importDefaultSpecifier(t.identifier(parent.id.name))],
            t.stringLiteral(source)
          );
          path.parentPath.parentPath?.replaceWith(importDecl);
          changes.push(`require('${source}') → import {default} from '${source}'`);
        }
        // const { a, b } = require('x') → import { a, b } from 'x'
        else if (t.isVariableDeclarator(parent) && t.isObjectPattern(parent.id)) {
          const specifiers = parent.id.properties.map((prop: any) =>
            t.importSpecifier(
              t.identifier(prop.value.name),
              t.identifier(prop.key.name)
            )
          );
          const importDecl = t.importDeclaration(specifiers, t.stringLiteral(source));
          path.parentPath.parentPath?.replaceWith(importDecl);
          changes.push(`require('${source}') → import { destructured } from '${source}'`);
        }
      }
    },

    // module.exports = ... → export default ...
    AssignmentExpression(path) {
      const node = path.node;
      if (
        t.isMemberExpression(node.left) &&
        t.isIdentifier(node.left.object, { name: 'module' }) &&
        t.isIdentifier(node.left.property, { name: 'exports' })
      ) {
        const exportDecl = t.exportDefaultDeclaration(node.right);
        path.replaceWith(exportDecl);
        changes.push('module.exports → export default');
      }
    },

    // exports.x = ... → export const x = ...
    ExpressionStatement(path) {
      const node = path.node;
      if (
        t.isAssignmentExpression(node.expression) &&
        t.isMemberExpression(node.expression.left) &&
        t.isIdentifier(node.expression.left.object, { name: 'exports' })
      ) {
        const prop = node.expression.left.property;
        if (t.isIdentifier(prop)) {
          const varDecl = t.variableDeclaration('const', [
            t.variableDeclarator(
              t.identifier(prop.name),
              node.expression.right
            ),
          ]);
          const exportDecl = t.exportNamedDeclaration(varDecl);
          path.replaceWith(exportDecl);
          changes.push(`exports.${prop.name} → export const ${prop.name}`);
        }
      }
    },
  });

  const result = generate(ast, { retainLines: true, retainFunctionParens: true }, code);
  return { code: result.code, changes };
}

// ─── Babel: Class → Hooks Transformer ───────────────────
function transformClassToHooks(code: string): { code: string; changes: string[] } {
  const changes: string[] = [];
  const ast = parse(code, {
    sourceType: 'module',
    plugins: ['jsx', 'typescript'],
  });

  traverse(ast, {
    ClassDeclaration(path) {
      const node = path.node;
      const className = node.id?.name;
      if (!className) return;

      // Check if it extends React.Component or Component
      const superClass = node.superClass;
      if (!superClass) return;

      // Convert to function declaration
      const bodyStatements: t.Statement[] = [];
      let hasState = false;

      // Collect state
      node.body.body.forEach((member) => {
        if (t.isClassProperty(member)) {
          const key = (member.key as t.Identifier).name;
          if (key === 'state' && t.isObjectExpression(member.value)) {
            hasState = true;
            // Add useState hooks
            member.value.properties.forEach((prop) => {
              if (t.isObjectProperty(prop) && t.isIdentifier(prop.key)) {
                const stateName = prop.key.name;
                const setterName = `set${stateName.charAt(0).toUpperCase() + stateName.slice(1)}`;
                bodyStatements.push(
                  t.variableDeclaration('const', [
                    t.variableDeclarator(
                      t.arrayPattern([
                        t.identifier(stateName),
                        t.identifier(setterName),
                      ]),
                      t.callExpression(t.identifier('useState'), [prop.value])
                    ),
                  ])
                );
                changes.push(`this.state.${stateName} → useState(${stateName})`);
              }
            });
          }
        }
      });

      // Add hooks import if needed
      if (hasState) {
        bodyStatements.unshift(
          t.importDeclaration(
            [t.importSpecifier(t.identifier('useState'), t.identifier('useState'))],
            t.stringLiteral('react')
          )
        );
      }

      // Convert render method
      node.body.body.forEach((member) => {
        if (
          t.isClassMethod(member) &&
          t.isIdentifier(member.key, { name: 'render' })
        ) {
          // Replace return statement
          member.body.body.forEach((stmt) => {
            if (t.isReturnStatement(stmt)) {
              bodyStatements.push(stmt);
            }
          });
          changes.push(`class ${className} → function ${className}`);
        }
      });

      const funcDecl = t.functionDeclaration(
        node.id,
        node.body.body
          .filter((m) => t.isClassMethod(m) && m.key && (m.key as t.Identifier).name === 'constructor')
          .flatMap((ctor: any) => ctor.params || []),
        t.blockStatement(bodyStatements)
      );

      path.replaceWith(
        t.exportNamedDeclaration(
          t.variableDeclaration('const', [
            t.variableDeclarator(t.identifier(className), t.arrowFunctionExpression(
              funcDecl.params as any,
              funcDecl.body as any
            )),
          ])
        )
      );
    },
  });

  const result = generate(ast, { retainLines: true }, code);
  return { code: result.code, changes };
}

// ─── File Processing ────────────────────────────────────
async function processFile(
  filePath: string,
  options: TransformOptions
): Promise<TransformResult> {
  const result: TransformResult = {
    file: filePath,
    transformed: false,
    changes: [],
  };

  const code = await fs.readFile(filePath, 'utf-8');

  // Skip if file is small or binary-like
  if (code.length < 10 || code.includes('\x00')) return result;

  let currentCode = code;

  // CommonJS → ESM
  if (options.cjsToEsm && /require\(|module\.exports|exports\./.test(currentCode)) {
    const transformed = transformCJStoESM(currentCode);
    currentCode = transformed.code;
    result.changes.push(...transformed.changes);
  }

  // Class → Hooks
  if (options.classToHooks && /extends\s+(React\.)?Component|PureComponent/.test(currentCode)) {
    const transformed = transformClassToHooks(currentCode);
    currentCode = transformed.code;
    result.changes.push(...transformed.changes);
  }

  // Add React import
  if (options.addReactImport && !/import\s+React/.test(currentCode) && /jsx|JSX/.test(currentCode)) {
    // Simple — add at top
    const lines = currentCode.split('\n');
    lines.unshift(`import React from 'react';`);
    currentCode = lines.join('\n');
    result.changes.push('Added React import');
  }

  // Rename imports
  if (options.renameImports?.length) {
    const ast = parse(currentCode, {
      sourceType: 'module',
      plugins: ['jsx', 'typescript'],
    });
    traverse(ast, {
      ImportDeclaration(path) {
        const source = path.node.source.value;
        for (const [oldName, newName] of options.renameImports!) {
          if (source === oldName) {
            path.node.source = t.stringLiteral(newName);
            result.changes.push(`import from '${oldName}' → '${newName}'`);
          }
        }
      },
    });
    currentCode = generate(ast, {}, currentCode).code;
  }

  // Write changes
  result.transformed = result.changes.length > 0;
  if (result.transformed && !options.dryRun) {
    await fs.writeFile(filePath, currentCode, 'utf-8');
  }

  return result;
}

// ─── CLI ────────────────────────────────────────────────
const program = new Command()
  .name('code-transformer')
  .description('Transform JavaScript/TypeScript code')
  .version('1.0.0');

program
  .command('convert')
  .description('Apply transformations to files')
  .argument('<pattern>', 'Glob pattern (e.g., "src/**/*.{ts,tsx,js,jsx}")')
  .option('--cjs-to-esm', 'Convert CommonJS to ES modules')
  .option('--class-to-hooks', 'Convert React class components to hooks')
  .option('--add-react-import', 'Add missing React imports')
  .option('--rename-import <mapping...>', 'Rename import paths (old=new old2=new2)')
  .option('--dry-run', 'Preview changes without writing')
  .option('-v, --verbose', 'Show all changes')
  .option('--ignore <pattern>', 'Ignore pattern', 'node_modules')
  .action(async (pattern: string, opts) => {
    const options: TransformOptions = {
      cjsToEsm: opts.cjsToEsm,
      classToHooks: opts.classToHooks,
      addReactImport: opts.addReactImport,
      dryRun: opts.dryRun,
      verbose: opts.verbose,
      renameImports: opts.renameImport?.map((m: string) => {
        const [oldPath, newPath] = m.split('=');
        return [oldPath, newPath || oldPath] as [string, string];
      }),
    };

    if (!options.cjsToEsm && !options.classToHooks && !options.addReactImport && !options.renameImports?.length) {
      console.log(chalk.red('✖ No transformations specified. Use --help to see options.'));
      process.exit(1);
    }

    console.log(chalk.bold.cyan('\n  🔧 Code Transformer\n'));
    if (options.dryRun) console.log(chalk.yellow('  ⚠ DRY RUN — no files will be modified\n'));

    // Find files
    const files = await fg(pattern, {
      ignore: [opts.ignore, '**/node_modules/**', '**/dist/**'],
      onlyFiles: true,
    });

    if (!files.length) {
      console.log(chalk.yellow('  No files matched'));
      process.exit(0);
    }

    console.log(chalk.dim(`  Found ${files.length} files\n`));

    // Process files
    let totalChanged = 0;
    let totalChanges = 0;

    for (const file of files) {
      const result = await processFile(file, options);
      if (result.transformed) {
        totalChanged++;
        totalChanges += result.changes.length;
        console.log(chalk.green('  ✓'), chalk.bold(path.relative(process.cwd(), file)));
        if (options.verbose) {
          result.changes.forEach((c) => console.log(chalk.dim('      →'), c));
        }
      }
    }

    // Summary
    const verb = options.dryRun ? 'would be' : 'were';
    console.log(chalk.bold(`\n  📊 Summary:`));
    console.log(chalk.dim(`     Files scanned: ${files.length}`));
    console.log(chalk.green(`     Files modified: ${totalChanged}`));
    console.log(chalk.dim(`     Changes made: ${totalChanges}`));
    console.log();
  });

program.parse(process.argv);
