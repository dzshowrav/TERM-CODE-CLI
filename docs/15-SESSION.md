# OpenChat CLI

# Session Management System Specification

Version: 1.0

---

# Overview

The Session Management System is responsible for preserving every conversation, workspace state, AI context, tool history, and project metadata.

A Session is much more than chat history.

A Session is a complete development workspace.

It remembers

• Conversation

• Provider

• Model

• Agent

• Skills

• Workspace

• Open Files

• Tool History

• Git Context

• AI State

This allows users to stop working and continue later without losing context.

---

# Philosophy

One Project

↓

Multiple Sessions

↓

Multiple Conversations

↓

Persistent Memory

Sessions should feel like reopening VS Code after closing it yesterday.

Everything should continue naturally.

---

# Goals

✓ Unlimited Sessions

✓ Automatic Saving

✓ Resume Anytime

✓ Search

✓ Archive

✓ Favorites

✓ Export

✓ Import

✓ Fork

✓ Workspace Awareness

---

# Session Structure

Every Session contains

Title

Description

Workspace

Provider

Model

Agent

Skills

Conversation

Tool History

Git Branch

Bookmarks

Statistics

Created Date

Updated Date

---

# Session Lifecycle

Create

↓

Start Chat

↓

Auto Save

↓

Pause

↓

Resume

↓

Archive

↓

Export

↓

Delete

---

# Database

sessions

Columns

id

uuid

title

description

workspace_id

provider_id

model_id

agent_id

status

favorite

archived

created_at

updated_at

---

# Status

Active

Paused

Archived

Deleted

---

# Message Storage

Each session owns

Many Messages

Every message stores

Role

Content

Timestamp

Tokens

Tool Calls

Attachments

Metadata

---

# Auto Save

Every action triggers

Auto Save

Examples

Send Message

Receive Response

Switch Model

Tool Execution

Workspace Scan

Permission Change

No manual save required.

---

# Session Resume

Command

/session resume

Displays

Recent Sessions

↓

Select

↓

Conversation Restored

↓

Workspace Loaded

↓

Continue Chat

---

# New Session

Command

/new session

Options

Blank Session

Current Workspace

Duplicate Current

Template

---

# Session Title

Generated automatically.

Example

Laravel Authentication

Fix Docker Build

React Dashboard

Bug Fix

API Documentation

Users may rename anytime.

---

# Rename Session

/session rename

New Title

[________________]

Save

---

# Delete Session

Confirmation

Delete

Cancel

Soft Delete by default.

Permanent deletion available separately.

---

# Archive Session

Archived sessions

Disappear from Recent

Remain searchable

May be restored anytime.

---

# Favorite Session

★

Pinned at top.

Useful for

Long-term projects

---

# Session Search

Supports

Title

Workspace

Provider

Model

Agent

Date

Keywords

Tags (Future)

---

# Session History

/history

Displays

Today

Yesterday

Last Week

Last Month

Older

Search is always available.

---

# Session Details

Displays

Title

Workspace

Provider

Model

Agent

Skills

Messages

Tool Calls

Duration

Created

Updated

---

# Fork Session

Creates

Complete Copy

Including

Conversation

Workspace

Context

Useful for

Trying different solutions.

---

# Export

Supported Formats

Markdown

JSON

HTML

PDF

Future

ZIP Archive

---

# Import

Supported

Markdown

JSON

OpenChat Backup

Future

OpenCode Sessions

Claude Code Sessions

Gemini CLI Sessions

---

# Workspace Binding

Each session belongs to

One Workspace

Switching Workspace

Creates

New Session

or

Prompts user

---

# Git Integration

Store

Repository

Current Branch

Latest Commit

Modified Files

Never stores Git objects.

---

# Attachments

Future

Images

PDF

Logs

Code Files

Screenshots

Audio

Videos

---

# Bookmarks

Bookmark

Messages

Code Blocks

Tool Outputs

Bookmarks appear in

/session bookmarks

---

# Notes

Users may attach

Personal Notes

To any session.

Notes are not sent to the AI unless explicitly requested.

---

# Statistics

Per Session

Messages

Input Tokens

Output Tokens

Tool Calls

Files Read

Files Modified

Duration

Cost (Future)

---

# Session Timeline

Create

↓

Chat

↓

Read Files

↓

Write Files

↓

Git

↓

Continue

↓

Archive

Timeline helps users review work.

---

# Filters

Active

Archived

Favorites

Recent

Longest

Most Used

Current Workspace

---

# Recovery

Unexpected Shutdown

↓

Recover Session

↓

Restore Chat

↓

Restore Context

↓

Continue

No conversation loss.

---

# Auto Cleanup

Optional

Delete sessions older than

30 Days

90 Days

1 Year

Never

---

# Security

Sessions stored locally.

No cloud upload by default.

Encrypted backups planned.

Sensitive information masked where possible.

---

# Performance Goals

Create Session

<20ms

Resume Session

<100ms

Search

Instant

Export

<1 Second

Auto Save

Background

---

# Future Features

Cloud Sync

Shared Sessions

Team Sessions

Live Collaboration

Session Templates

Session Tags

Session Analytics

AI Session Summary

Timeline Visualization

Workspace Snapshots

---

# Design Principles

A session is the developer's memory.

It should preserve not only messages but the complete development environment, allowing developers to continue work exactly where they left off.

OpenChat CLI should make sessions feel reliable, intelligent, searchable, and effortless, providing a seamless experience across long-running software projects.