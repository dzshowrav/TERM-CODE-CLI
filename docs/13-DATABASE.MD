# OpenChat CLI

# Database System Specification

Version: 1.0

---

# Overview

OpenChat CLI uses SQLite as its primary database.

SQLite was selected because it is

• Fast

• Lightweight

• Offline

• Cross-platform

• Zero Configuration

• Perfect for Android (Termux)

The database stores only application metadata.

Project source code is never stored inside the database.

---

# Design Principles

✓ Normalized

✓ Fast

✓ ACID

✓ Indexed

✓ Extensible

✓ Versioned

✓ Migration Ready

✓ Backup Friendly

---

# Database Engine

SQLite 3

Database File

~/.openchat/openchat.db

WAL Mode

Enabled

Foreign Keys

Enabled

Automatic Vacuum

Enabled

---

# Database Architecture

SQLite

│

├── Providers

├── Models

├── Agents

├── Skills

├── Sessions

├── Messages

├── Workspaces

├── Settings

├── MCP

├── Plugins

├── Permissions

├── Logs

└── Cache

---

# Tables

providers

models

agents

skills

sessions

messages

workspace

settings

permissions

tools

tool_logs

mcp_servers

plugins

themes

favorites

history

recent

cache

indexes

bookmarks

attachments

statistics

---

# providers

Stores AI providers.

Columns

id

uuid

name

description

base_url

api_key

status

latency

priority

is_default

created_at

updated_at

Indexes

name

base_url

---

# models

Stores registered AI models.

Columns

id

provider_id

model_id

display_name

description

category

supports_streaming

supports_tools

supports_reasoning

supports_vision

supports_json

supports_audio

supports_embedding

max_context

max_output

favorite

enabled

created_at

updated_at

Relationship

Provider

↓

Many Models

---

# agents

Stores AI agents.

Columns

id

name

description

icon

prompt

reasoning_level

temperature

default_skills

allowed_tools

enabled

built_in

created_at

updated_at

---

# skills

Stores installed skills.

Columns

id

name

version

category

description

priority

path

enabled

built_in

created_at

updated_at

---

# sessions

Stores conversations.

Columns

id

title

provider_id

model_id

agent_id

workspace_id

summary

favorite

archived

created_at

updated_at

Relationship

Session

↓

Many Messages

---

# messages

Columns

id

session_id

role

content

tool_calls

token_input

token_output

created_at

---

# workspace

Stores project information.

Columns

id

name

path

language

framework

package_manager

git_branch

last_scan

created_at

---

# settings

Stores global configuration.

Columns

key

value

updated_at

Examples

theme

language

default_provider

default_model

default_agent

streaming

permissions

---

# permissions

Stores tool permissions.

Columns

tool_name

permission

updated_at

Values

always_allow

allow_once

ask

deny

---

# tools

Stores available tools.

Columns

id

name

category

description

plugin

enabled

version

---

# tool_logs

Stores execution history.

Columns

id

tool

arguments

status

duration

created_at

Logs never store API keys.

---

# mcp_servers

Stores MCP configuration.

Columns

id

name

transport

command

url

status

enabled

created_at

---

# plugins

Stores installed plugins.

Columns

id

name

version

author

description

enabled

installed_at

---

# themes

Stores themes.

Columns

id

name

author

version

primary_color

accent_color

dark_mode

---

# favorites

Stores user favorites.

Types

Model

Command

Skill

Agent

Session

---

# history

Stores recently used items.

Examples

Commands

Models

Providers

Files

Sessions

---

# cache

Stores temporary data.

Examples

Provider Health

Workspace Scan

Token Count

Capabilities

Autocomplete

Cache can be cleared safely.

---

# bookmarks

Stores

Favorite Messages

Favorite Files

Favorite Sessions

Pinned Conversations

---

# attachments

Future

Images

PDF

Logs

Audio

Screenshots

---

# statistics

Stores usage statistics.

Examples

Messages

Sessions

Providers Used

Models Used

Tool Calls

Average Latency

Total Tokens

---

# Relationships

Provider

↓

Models

↓

Sessions

↓

Messages

Agent

↓

Sessions

Workspace

↓

Sessions

Plugin

↓

Tools

---

# Foreign Keys

Provider

↓

Models

Model

↓

Sessions

Session

↓

Messages

Workspace

↓

Sessions

---

# Indexes

Optimize

Provider Name

Model ID

Session Date

Message Date

Workspace Path

Tool Name

Command History

Search Keywords

---

# Transactions

All write operations use transactions.

Example

Create Session

↓

Insert Session

↓

Insert Message

↓

Commit

Failure

↓

Rollback

---

# Auto Save

Every message

↓

Immediately saved

No manual save required.

---

# Backup

Command

/database backup

Creates

SQLite Backup

Compressed Backup

Timestamped Backup

---

# Restore

/database restore

↓

Validate

↓

Replace Database

↓

Restart Required

---

# Migrations

Each database version contains

Migration Number

Description

SQL Script

Rollback Script

Example

001_initial

002_models

003_agents

004_skills

---

# Performance Goals

Startup

<50ms

Provider Lookup

<5ms

Model Lookup

<5ms

Session Load

<30ms

Search

Instant

---

# Future Tables

cloud_sync

organizations

teams

shared_sessions

workflow

voice

vision

billing

marketplace

analytics

---

# Security

API Keys are encrypted (future).

Sensitive values are masked in logs.

Foreign key constraints prevent orphaned records.

Database corruption is automatically detected during startup.

Automatic backups can be scheduled.

---

# Design Principles

The database is the foundation of OpenChat CLI.

It should remain lightweight, reliable, and fully offline while supporting millions of messages, thousands of sessions, hundreds of providers and models, and future enterprise features without requiring architectural changes.

SQLite provides the perfect balance of simplicity, speed, and portability for a mobile-first AI coding assistant.