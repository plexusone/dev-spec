# devspec validate

Validate spec structure against SDD type.

## Synopsis

```bash
devspec validate [path] [flags]
```

## Description

Validates the structure of spec files against the detected or specified SDD type. Checks for required sections, file presence, and format compliance.

This is a deterministic command that does not use LLM calls.

## Arguments

| Argument | Description | Default |
|----------|-------------|---------|
| `path` | Directory to validate | Current directory |

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--type` | `-t` | SDD type to validate against (auto-detected if not specified) |

## Examples

```bash
# Validate current directory
devspec validate

# Validate specific directory
devspec validate ./my-project

# Force specific SDD type
devspec validate --type kiro

# Output as markdown
devspec validate --format markdown
```

## Output

```
SDD Type: kiro
Valid: true

Present Sections:
  [+] requirements/Document Information (required)
  [+] requirements/Introduction (required)
  [+] requirements/Requirements (required)
  [+] design/Architecture (required)

Missing Sections:
  [-] requirements/Non-Functional Requirements (optional)

Warnings:
  - design.md: Security Considerations section is brief
```
