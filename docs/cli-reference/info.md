# devspec info

Get SDD type information for coding assistants.

## Synopsis

```bash
devspec info <type> [flags]
```

## Description

Returns structured information about an SDD type including required and optional files, section structure, and template content.

Designed for coding assistants (Claude Code, Cursor, etc.) to guide users in creating spec files.

## Arguments

| Argument | Description |
|----------|-------------|
| `type` | SDD type (kiro, speckit, plexusone) |

## Examples

```bash
# Get Kiro info (TOON output)
devspec info kiro

# JSON for programmatic use
devspec info kiro --format json

# Markdown documentation
devspec info speckit --format markdown
```

## Output Fields

| Field | Description |
|-------|-------------|
| `name` | SDD type identifier |
| `display_name` | Human-readable name |
| `description` | Brief description |
| `spec_directory` | Where spec files are stored |
| `files` | List of file definitions |

Each file includes:

| Field | Description |
|-------|-------------|
| `name` | File identifier |
| `display_name` | Human-readable name |
| `patterns` | File path patterns |
| `required` | Whether file is required |
| `requires_codebase` | Whether codebase access is needed |
| `sections` | Required/optional sections |
| `template` | Template content |

## Use with Coding Assistants

Coding assistants can use this command to:

1. Understand what files are needed
2. Know which sections are required
3. Determine when codebase access is needed
4. Generate initial content from templates

```bash
# Get info and pipe to assistant context
devspec info kiro --format json | pbcopy
```
