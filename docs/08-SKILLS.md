# OpenChat CLI

# Skill System Specification

Version: 1.0

---

# Overview

The Skill System is one of the most powerful features of OpenChat CLI.

Skills are reusable knowledge packages that teach an agent how to perform specific tasks.

Unlike Agents, which define personality and behavior, Skills provide domain-specific knowledge, best practices, templates, workflows, and coding standards.

Skills should be automatically loaded only when required.

This minimizes context usage while maximizing response quality.

---

# Philosophy

Provider

↓

Model

↓

Agent

↓

Skills

↓

Context

↓

AI Response

Agent decides HOW to think.

Skills provide WHAT to know.

---

# Goals

✓ Modular

✓ Reusable

✓ Lightweight

✓ Lazy Loaded

✓ Project Aware

✓ Language Aware

✓ Framework Aware

✓ Shareable

✓ Versioned

✓ Plugin Ready

---

# Skill Structure

Each skill contains

Name

Description

Version

Author

Category

Instructions

Examples

Templates

Best Practices

Dependencies

Keywords

Supported Languages

Supported Frameworks

---

# Folder Structure

skills/

laravel/

SKILL.md

examples/

templates/

rules/

metadata.json

---

react/

SKILL.md

examples/

metadata.json

---

docker/

SKILL.md

templates/

metadata.json

---

# Skill Metadata

Every skill contains

ID

Name

Version

Description

Author

License

Priority

Category

Tags

Dependencies

Languages

Frameworks

Keywords

---

# Database Schema

skills

id

name

version

description

category

priority

enabled

built_in

author

created_at

updated_at

---

# Skill Categories

Backend

Frontend

Mobile

AI

Database

DevOps

Testing

Security

Performance

Architecture

Documentation

Cloud

API

CLI

Automation

Custom

---

# Built-in Skills

Laravel

PHP

Composer

Blade

Eloquent

MySQL

PostgreSQL

SQLite

Redis

JavaScript

TypeScript

React

Vue

Next.js

Node.js

Express

Tailwind CSS

Bootstrap

Docker

Git

GitHub

Linux

REST API

GraphQL

JSON

Authentication

Authorization

Caching

Queue

Performance

Testing

Debugging

Logging

Refactoring

Documentation

Clean Code

Design Patterns

SOLID

DDD

Microservices

CI/CD

OpenAPI

MCP

OpenAI API

---

# Skill Loading

Skills are never loaded all at once.

Instead

Project Analysis

↓

Detect Technologies

↓

Select Relevant Skills

↓

Load Skill Instructions

↓

Inject into Context

Example

Laravel Project

↓

PHP

Laravel

Composer

Blade

Security

Performance

Only these skills are loaded.

---

# Automatic Detection

The system automatically detects

composer.json

↓

Laravel

package.json

↓

React

Next.js

Vue

Cargo.toml

↓

Rust

go.mod

↓

Go

requirements.txt

↓

Python

pubspec.yaml

↓

Flutter

---

# Manual Skill Selection

Command

/skills

↓

Search

↓

Select

↓

Enable

↓

Use

---

# Skill Manager

Command

/skills

Displays

Laravel

PHP

Docker

Git

React

Security

Performance

Testing

---

# Skill Details

Selecting a skill shows

Name

Description

Category

Version

Dependencies

Examples

Templates

Supported Languages

Supported Frameworks

---

# Skill Dependencies

Example

Laravel

Depends On

PHP

Composer

Authentication

Validation

These dependencies load automatically.

---

# Skill Priority

Highest

Security

Performance

Architecture

Medium

Framework

Language

Low

Documentation

Examples

Priority controls load order.

---

# Skill Context

Each skill contributes

Rules

Patterns

Templates

Examples

Coding Standards

Best Practices

Common Mistakes

Performance Tips

Security Rules

---

# Skill Examples

Laravel Skill

Includes

Controller Example

Migration Example

Validation Example

Policy Example

API Resource Example

Queue Example

---

# Skill Templates

Example

Laravel Controller

Laravel Middleware

Dockerfile

GitHub Workflow

Docker Compose

React Component

Node API

---

# Skill Rules

Example

Laravel

Always use Form Request Validation.

Use Route Model Binding.

Prefer Eloquent Relationships.

Avoid Raw SQL.

Use Queues for long-running jobs.

---

# Skill Search

Search by

Name

Category

Language

Framework

Keyword

Tag

Instant filtering.

---

# Enable / Disable

Users may disable

Unused skills

Custom skills

Experimental skills

Disabled skills are ignored.

---

# Custom Skills

Command

/new skill

Fields

Name

Description

Category

Instructions

Examples

Templates

Keywords

Dependencies

Save

---

# Edit Skill

Editable

Description

Instructions

Examples

Templates

Dependencies

Priority

---

# Delete Skill

Only

Custom Skills

Built-in skills cannot be deleted.

---

# Export Skill

Supports

JSON

Markdown

ZIP

---

# Import Skill

Import

↓

Validate

↓

Install

↓

Available

---

# Skill Marketplace

Future

Official Skills

Community Skills

Verified Authors

Ratings

Downloads

Updates

---

# Skill Versioning

Every skill stores

Version

Author

License

Compatibility

Minimum CLI Version

---

# Context Optimization

The Context Manager should only load

Required Sections

Never the entire skill.

Example

Laravel Skill

Contains

20 sections

Only

Routing

Validation

Controllers

are loaded

if needed.

This greatly reduces token usage.

---

# AI Workflow

User Prompt

↓

Analyze Intent

↓

Detect Project

↓

Load Agent

↓

Load Required Skills

↓

Build Context

↓

Call Model

↓

Generate Response

---

# Performance Goals

Skill Detection

<20ms

Skill Load

<50ms

Search

Instant

Memory Usage

Minimal

---

# Future Features

Skill Packs

Premium Skills

Cloud Sync

Automatic Updates

AI Generated Skills

Company Skill Libraries

Workspace Skills

Team Skills

Skill Analytics

Skill Recommendations

---

# Design Principles

Skills are knowledge modules.

They should never replace Agents.

Instead, they enhance the selected Agent with focused expertise.

By loading only the relevant skills at the right time, OpenChat CLI delivers higher-quality responses while keeping token usage low and maintaining excellent performance on Android Termux.

A developer should be able to install hundreds of skills without slowing down the application because only the skills needed for the current task are loaded into the AI context.