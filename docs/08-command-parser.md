# 08-command-parser.md

# Command Parser Architecture
## Complete Guide for AI CLI Coding Agents (OpenCode / Antigravity Style)

Version: 1.0

---

# What is a Command Parser?

A **Command Parser** is the subsystem responsible for understanding everything the user types into the CLI.

It converts raw text into structured actions that the Agent Engine can execute.

Instead of treating every input as plain chat, the parser determines whether the input is:

- A built-in command
- A slash command
- An AI prompt
- A shell command
- A plugin command
- A workflow
- A macro
- A shortcut

The parser is the first intelligent layer after user input.

---

# Why Command Parser?

Without Parser

```
User Input

↓

Agent

↓

Everything Sent To AI
```

Problems

- Slow
- Expensive
- Cannot execute local commands
- Difficult extension
- No autocomplete

---

With Parser

```
User Input

↓

Tokenizer

↓

Parser

↓

Command Router

↓

Target Handler
```

---

# Goals

A production Command Parser should provide

- Fast parsing
- Command routing
- Slash commands
- Flags
- Options
- Arguments
- Variables
- Aliases
- Autocomplete
- Plugin commands
- Validation
- Error recovery

---

# High-Level Architecture

```
               Keyboard
                   │
                   ▼
             Input Buffer
                   │
                   ▼
              Tokenizer
                   │
                   ▼
               Parser
                   │
         ┌─────────┼──────────┐
         ▼         ▼          ▼
    Validator   Resolver   Router
         │         │          │
         └────┬────┴────┬─────┘
              ▼         ▼
       Command Engine Agent Engine
```

---

# Folder Structure

```
src/

parser/

    CommandParser.ts

    Tokenizer.ts

    Lexer.ts

    AST.ts

    CommandRouter.ts

    CommandRegistry.ts

    AliasManager.ts

    ArgumentParser.ts

    FlagParser.ts

    OptionParser.ts

    Validator.ts

    AutoComplete.ts

    VariableResolver.ts

    History.ts

    ParserEvents.ts

    ParserMetrics.ts
```

---

# Core Components

## Command Parser

Central controller.

Responsibilities

- Parse input
- Build AST
- Validate commands
- Route execution

---

## Tokenizer

Splits raw text into tokens.

Example

Input

```
/open src/index.ts --line 25
```

Tokens

```
/open

src/index.ts

--line

25
```

---

## Lexer

Determines token types.

Examples

```
Command

Argument

Flag

Option

String

Number

Variable
```

---

## AST (Abstract Syntax Tree)

Represents parsed input.

Example

```
Command

├── Name

├── Flags

├── Options

└── Arguments
```

---

## Command Registry

Stores

```
Built-in Commands

Plugin Commands

Aliases

Descriptions

Permissions
```

---

## Command Router

Routes parsed command.

Example

```
/theme

↓

Theme Manager

-------------

/plugin

↓

Plugin Manager

-------------

/help

↓

Help System
```

---

## Validator

Checks

- Syntax
- Unknown commands
- Required arguments
- Invalid flags
- Permissions

---

## AutoComplete

Provides

```
Command Suggestions

Flags

Arguments

Files

History

Plugins
```

---

## Alias Manager

Maps

```
/ls

↓

/files

--------------

/quit

↓

/exit
```

---

## Variable Resolver

Expands variables.

Example

```
$HOME

$PWD

$PROJECT

$FILE
```

---

# Parsing Lifecycle

```
Receive Input

↓

Tokenize

↓

Lex

↓

Parse

↓

Validate

↓

Resolve

↓

Route

↓

Execute
```

---

# Input Types

The parser should recognize

### AI Prompt

```
Explain dependency injection
```

↓

Send to Agent

---

### Slash Command

```
/theme dark
```

↓

Theme Manager

---

### Shell Command

```
!git status
```

↓

Terminal

---

### Plugin Command

```
/deploy production
```

↓

Plugin

---

### Workflow

```
/workflow build
```

↓

Workflow Engine

---

### Macro

```
@review
```

↓

Macro Engine

---

# Command Grammar

General structure

```
command

argument

option

flag
```

Example

```
/open

src/main.ts

--line 40

--readonly
```

---

# Flags

Boolean values.

Examples

```
--force

--silent

--watch

--debug
```

---

# Options

Require values.

Examples

```
--model gpt-5

--theme dracula

--port 3000
```

---

# Arguments

Examples

```
filename

directory

URL

project
```

---

# Quoted Strings

Example

```
/chat "Explain async await"
```

Tokenizer keeps it together as one argument.

---

# Escape Characters

Support

```
\"

\\

\n

\t
```

---

# Variable Expansion

Example

Input

```
/open $PROJECT/src/index.ts
```

Resolved

```
/workspace/myapp/src/index.ts
```

---

# Environment Variables

Supported

```
$HOME

$USER

$PWD

$LANG

$SHELL
```

---

# Command Routing

```
Parser

↓

Router

↓

Registry

↓

Handler

↓

Execution
```

---

# Event Bus Integration

Common events

```
parser:start

parser:complete

command:execute

command:error

autocomplete:update
```

---

# Agent Integration

```
Unknown Command

↓

Treat As Prompt

↓

Agent
```

or

```
Known Command

↓

Execute Handler
```

---

# Plugin Integration

Plugins may register

```
Commands

Aliases

Autocomplete

Validators
```

No parser changes required.

---

# Skills Integration

Skills may contribute

```
Command Templates

Prompt Templates

Workflow Commands
```

---

# MCP Integration

Command may invoke MCP tools.

Example

```
/github search
```

↓

GitHub MCP Server

---

# Error Handling

```
Invalid Input

↓

Suggestion

↓

Retry

↓

Help
```

Example

```
/plgin

↓

Did you mean

/plugin ?
```

---

# Help System

Each command provides

```
Description

Arguments

Flags

Examples

Permissions
```

---

# Autocomplete Flow

```
User Types

↓

Tokenizer

↓

Registry Search

↓

Suggestions

↓

Display
```

---

# History Integration

Stores

```
Recent Commands

Frequency

Pinned Commands

Favorites
```

Supports navigation with arrow keys.

---

# Permission Checks

Before execution

Verify

```
Command Exists

↓

Permission Granted

↓

Arguments Valid

↓

Execute
```

---

# Performance Optimizations

Use

- Trie for commands
- Cached parsing
- Incremental tokenization
- Lazy autocomplete
- Immutable AST
- Fast lookup tables

Avoid

- Linear search
- Re-tokenizing unchanged input
- Blocking autocomplete

---

# Security

Always

- Validate commands
- Sanitize arguments
- Restrict dangerous operations
- Verify permissions
- Escape shell input

Never

- Execute raw input directly
- Trust plugin arguments
- Ignore malformed syntax

---

# Best Practices

Always

- Separate parsing from execution
- Build an AST
- Support autocomplete
- Provide helpful errors
- Keep grammar consistent
- Register commands dynamically

Never

- Mix parser with business logic
- Hardcode plugin commands
- Ignore invalid syntax
- Execute before validation

---

# Common Mistakes

Bad

```
Input

↓

String Compare

↓

Execute
```

Hard to maintain.

---

Good

```
Input

↓

Tokenizer

↓

Parser

↓

AST

↓

Router

↓

Handler
```

Modular and extensible.

---

# Testing Checklist

- Tokenization
- Parsing
- Flags
- Options
- Arguments
- Variables
- Aliases
- Autocomplete
- Invalid syntax
- Unknown commands
- Plugin commands
- Permission checks

---

# Example Built-in Commands

Examples

```
/help

/exit

/clear

/theme

/model

/plugin

/skill

/session

/config

/history

/open

/save

/reload

/version
```

---

# Advantages

- Fast command execution
- Modular architecture
- Extensible commands
- Better UX
- Autocomplete
- Structured validation
- Plugin support
- Workflow support

---

# Disadvantages

- Grammar maintenance
- Parser complexity
- Alias conflicts
- Version compatibility

---

# Used In

- OpenCode
- Antigravity CLI
- Claude Code
- Git CLI
- Docker CLI
- Kubernetes kubectl
- GitHub CLI
- AWS CLI
- Azure CLI
- VS Code Command Palette

---

# Summary

The **Command Parser** is the gateway between user input and application execution.

A production-grade parser should include:

- Tokenizer
- Lexer
- AST Builder
- Validator
- Command Registry
- Router
- Alias Manager
- Variable Resolver
- Autocomplete
- Permission System
- Event Bus Integration

By separating parsing, validation, routing, and execution, the CLI becomes fast, extensible, secure, and capable of supporting built-in commands, plugins, workflows, AI prompts, and future features without changing the core parsing architecture.