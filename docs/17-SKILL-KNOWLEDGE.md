# Skill & Knowledge System

## Overview

The skill and knowledge system provides the agent with specialized capabilities and domain-specific knowledge. Skills are additive — loading a skill gives the agent new tools, knowledge, and behavioral instructions without modifying the core system.

## Skill Definition

A skill is a package of:

### Skill Components
- **Name** — unique identifier
- **Description** — what the skill provides and when to use it
- **System prompt injection** — additional instructions added to the system prompt when this skill is active
- **Tools** — additional tools provided by this skill (e.g., a Docker skill provides docker-related tools)
- **Knowledge** — reference information injected into context (e.g., framework documentation, API references)
- **Dependencies** — other skills this skill depends on
- **Triggers** — conditions that auto-activate this skill
- **Templates** — prompt templates for common workflows

### Trigger-Based Auto-Load

Skills can auto-detect when they're needed:
- **Pattern-based** — if the user mentions specific keywords (e.g., "docker", "container", "image")
- **File-based** — if the workspace contains certain files (e.g., `Dockerfile` triggers Docker skill, `package.json` triggers Node.js skill)
- **Task-based** — if the AI model selects a tool from the skill
- **Context-based** — if the current workspace matches the skill's project type

When a trigger fires:
1. The skill is loaded into the current session
2. Its system prompt is injected
3. Its tools become available
4. The user is notified (briefly) that the skill is active
5. Skills are unloaded when no longer relevant (with user notification)

## Skill Library

The system ships with a comprehensive skill library:
- **Languages** — JavaScript, TypeScript, Python, Rust, Go, Java, Ruby, C++, SQL, Shell, HTML, CSS
- **Frameworks** — React, Vue, Angular, Astro, Next.js, Express, Nest, FastAPI, Django, Rails, Spring
- **Tools** — Docker, Git, AWS, GCP, Azure, Kubernetes, Terraform
- **Testing** — Jest, Vitest, Playwright, Cypress, pytest, JUnit, RSpec
- **Quality** — Accessibility, Performance, Security, SEO, Code Review
- **Process** — Architecture, Planning, Debugging, Refactoring, Documentation

Total: 145+ skills.

## Knowledge Injection

Knowledge consists of reference information the agent can use:

### Knowledge Sources
- **Documentation** — framework guides, API references, best practice docs
- **Code patterns** — common patterns, idioms, conventions
- **Configuration** — recommended configs for tools and frameworks
- **Templates** — boilerplate code, starter templates
- **Cheatsheets** — quick references for commands, APIs, syntax

### Knowledge Lifecycle
1. **Stored** — knowledge is stored in a structured format
2. **Injected** — when relevant, knowledge is injected into the AI's context
3. **Retrieved** — knowledge can be explicitly retrieved by the AI model
4. **Updated** — knowledge can be updated or added by the user

Knowledge is injected via:
- **Context injection** — added to the system prompt
- **Tool-accessible** — available via a search/retrieve tool
- **On-demand** — fetched when the AI requests it

## Prompt Templates

Skills can provide prompt templates for common workflows:
- "Write unit tests for this component"
- "Review this PR for security issues"
- "Optimize this database query"
- "Refactor this class following best practices"

Each template includes:
- The prompt text to send to the AI
- Required parameters (file paths, names, etc.)
- Expected output format
- Skills and tools that should be active

## Skill Management

Users can:
- **List** — show all available skills with descriptions
- **Load** — manually activate a skill
- **Unload** — deactivate a skill
- **Search** — find skills by name, description, or trigger
- **Configure** — set skill-specific options
- **Create** — define custom skills
- **Disable** — prevent a skill from auto-loading

## Key Design Decisions

- Skills are purely additive — they never override core behavior
- Auto-loading is non-intrusive — a brief notification is sufficient
- Skills are scoped to sessions — loading a skill doesn't affect other sessions
- Knowledge is injectable but also retrievable — the AI can look things up
- Templates provide consistency for common workflows
- The skill library is extensive but lazy-loaded — only active skills consume context space
- Custom skills are as powerful as built-in skills — no artificial limitations
- Dependencies ensure skills can build on each other (e.g., React skill depends on JavaScript skill)
