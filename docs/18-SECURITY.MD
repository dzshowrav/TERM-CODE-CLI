# OpenChat CLI

# Security System Specification

Version: 1.0

---

# Overview

The Security System protects

• User Data

• API Keys

• Source Code

• Workspaces

• AI Context

• Tool Execution

• MCP Servers

• Plugins

• Local Database

Security is a core feature, not an optional add-on.

Every action performed by OpenChat CLI must be transparent, permission-aware, and reversible whenever possible.

---

# Security Goals

✓ Secure by Default

✓ Offline First

✓ Least Privilege

✓ User Consent

✓ Transparent Actions

✓ Secret Protection

✓ Audit Logging

✓ Plugin Isolation

✓ MCP Security

✓ Enterprise Ready

---

# Security Architecture

User

↓

Permission Manager

↓

Security Manager

↓

Tool Manager

↓

Operating System

↓

Workspace

No tool can bypass the Security Manager.

---

# Security Layers

Layer 1

Authentication

↓

Layer 2

Authorization

↓

Layer 3

Permission Manager

↓

Layer 4

Sandbox

↓

Layer 5

Audit Logs

↓

Layer 6

Recovery

---

# Permission Model

Every action belongs to one of four levels.

Read

Safe

Write

Sensitive

Execute

Dangerous

Network

External

The UI should always communicate which level an action belongs to.

---

# Permission Options

Always Allow

Allow Once

Ask Every Time

Always Deny

Permissions are stored locally and can be reset.

---

# Read Permissions

Examples

Read File

Search Project

List Directory

Git Status

Workspace Scan

Default

Always Allow

---

# Write Permissions

Examples

Write File

Edit File

Rename File

Create Folder

Default

Ask Every Time

---

# Dangerous Permissions

Examples

Delete File

Delete Folder

Force Git Reset

Shell Commands

Database Writes

Plugin Installation

MCP Registration

Default

Always Ask

---

# Network Permissions

Examples

HTTP Request

Download

Upload

Remote MCP

Cloud Sync

Default

Ask

---

# Secret Management

Secrets include

API Keys

Access Tokens

Bearer Tokens

Passwords

Private Keys

Cookies

Session Tokens

Secrets must never be displayed in plain text after saving.

---

# API Key Storage

Provider Keys

Stored Encrypted (future)

Masked in UI

Example

sk-********************************

Copy requires confirmation.

---

# Secret Redaction

Logs automatically replace

API Keys

Tokens

Passwords

Cookies

Authorization Headers

With

********

Secrets never appear in logs or exported diagnostics.

---

# Environment Variables

Supported

OPENAI_API_KEY

ANTHROPIC_API_KEY

GEMINI_API_KEY

CUSTOM_PROVIDER_KEY

Environment values override saved settings only for the current runtime.

---

# Workspace Protection

The AI may only access

Current Workspace

unless the user explicitly grants access elsewhere.

Attempting to read outside the workspace triggers a confirmation.

---

# Filesystem Boundaries

Allowed

Workspace

User-approved folders

Denied by default

System directories

Sensitive home folders

OS configuration

unless explicitly approved.

---

# Shell Execution

Every shell command passes through

Permission Manager

↓

Preview

↓

User Confirmation (if required)

↓

Execution

↓

Output

↓

Audit Log

---

# Command Preview

Before executing

Show

Command

Working Directory

Environment

Potential Risk

Example

rm -rf build/

Risk

High

Requires confirmation.

---

# Safe Shell Policy

Block obviously destructive commands unless explicitly approved.

Examples

rm -rf /

mkfs

shutdown

reboot

Privilege escalation attempts

System-wide package removal

The user may still choose to execute if the policy allows it and the risk is acknowledged.

---

# Plugin Security

Plugins cannot access

Database

Secrets

Filesystem

Network

unless permissions are granted.

Every plugin declares required capabilities.

---

# Plugin Permissions

Filesystem

Network

Clipboard

Database

MCP

Shell

Git

Workspace

Users approve capabilities during installation.

---

# MCP Security

Every MCP Server must declare

Tools

Resources

Prompts

Permissions

Authentication

OpenChat validates these before enabling the server.

---

# Remote MCP

Display

Server URL

Authentication Method

Certificate Status (HTTPS)

Permissions Requested

Connection Status

Require confirmation before first connection.

---

# Prompt Injection Defense

The Chat Engine should detect attempts to

Ignore system instructions

Reveal secrets

Disable permissions

Access unauthorized files

Manipulate tool execution

When detected

Warn the user

Ignore malicious instructions

Continue safely

---

# Workspace Trust

When opening a project

Workspace Trusted?

Yes

No

If untrusted

Disable

Shell

Plugins

Remote MCP

Automatic Tool Execution

Until trusted.

---

# Audit Log

Every sensitive action records

Timestamp

User Action

Tool

Target

Result

Duration

Permission Used

Secrets are never logged.

---

# Security Dashboard

Command

/security

Displays

Workspace Trust

Plugin Status

MCP Status

Recent Permissions

Audit Summary

Risk Alerts

---

# Recovery

If corruption or suspicious activity is detected

Create Backup

↓

Disable Risky Features

↓

Notify User

↓

Allow Recovery

---

# Backup Security

Backups may optionally be

Password Protected

Encrypted (future)

Verified before restore.

---

# Import Validation

Imported

Settings

Plugins

MCP Configurations

Agents

Skills

Are validated before activation.

Invalid or malformed files are rejected.

---

# Rate Limiting

Protect against

Accidental loops

Runaway tool execution

Repeated network failures

Excessive retries

Limits are configurable.

---

# Data Privacy

By default

No telemetry

No analytics

No cloud upload

No conversation sharing

Everything stays on the local device unless the user explicitly enables additional features.

---

# Security Notifications

Examples

New Plugin Installed

New MCP Connected

Permission Changed

Workspace Trust Changed

API Key Updated

Notifications are concise and actionable.

---

# Enterprise Features (Future)

Encrypted Database

Hardware-backed Key Storage

Multi-user Profiles

Organization Policies

Role-based Permissions

Signed Plugins

Signed MCP Servers

Security Compliance Reports

SSO Integration

---

# Performance Goals

Permission Check

<5ms

Audit Write

<10ms

Workspace Trust Check

Instant

Secret Redaction

Automatic

Minimal runtime overhead.

---

# Design Principles

Security should protect the user without interrupting productive work.

OpenChat CLI must clearly explain why permissions are requested, make sensitive actions visible before execution, and keep all secrets under the user's control.

The system should default to the safest behavior while remaining flexible enough for advanced developers who intentionally need powerful capabilities.