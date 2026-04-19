# devspec init

Initialize spec scaffolding for an SDD type.

## Synopsis

```bash
devspec init <type> [path] [flags]
```

## Description

Creates the required directory structure and template files for a specific SDD type.

## Arguments

| Argument | Description | Default |
|----------|-------------|---------|
| `type` | SDD type (kiro, speckit, plexusone) | Required |
| `path` | Directory to initialize | Current directory |

## Flags

| Flag | Description |
|------|-------------|
| `--force` | Overwrite existing files |

## Examples

```bash
# Initialize Kiro in current directory
devspec init kiro

# Initialize in specific directory
devspec init kiro ./my-project

# Initialize PlexusOne
devspec init plexusone ./my-project

# Overwrite existing files
devspec init kiro --force
```

## Output

```
Created ./my-project/.kiro/specs/requirements.md
Created ./my-project/.kiro/specs/design.md
Created ./my-project/.kiro/specs/tasks.md

Initialized Kiro Spec-Driven Development spec structure in ./my-project/.kiro/specs/

Next steps:
  1. Fill in the template files with your specifications
  2. Run 'devspec validate' to check structure
  3. Run 'devspec evaluate' for LLM-powered evaluation
```

## Generated Templates

Templates include:

- Required section headers
- `[TODO: Add content]` placeholders
- Field prompts where applicable

Example generated `requirements.md`:

```markdown
# Requirements Document

## Document Information (Required)

- **Feature Name**: 
- **Version**: 
- **Author**: 

## Introduction (Required)

### Feature Summary

[TODO: Add content]

### Business Value

[TODO: Add content]
```
