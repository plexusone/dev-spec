# Coding Assistants Integration

Use devspec with Claude Code, Cursor, Kiro CLI, and other coding assistants.

## Overview

In this model:

- **devspec** provides deterministic tools (no LLM calls)
- **Coding assistant** provides LLM capabilities
- **User** interacts with the assistant

## Workflow

### 1. Initialize Project

```bash
# Assistant runs:
devspec init kiro ./my-project
```

### 2. Get SDD Information

Assistant retrieves structure information:

```bash
devspec info kiro --format json
```

This tells the assistant:

- What files are needed
- Which sections are required
- When codebase access is needed
- Template content for scaffolding

### 3. Guide Spec Creation

The assistant uses `info` output to:

- Prompt user for each required section
- Ensure EARS format for acceptance criteria (Kiro)
- Check for completeness before moving on

### 4. Self-Evaluate Using Rubrics

```bash
devspec rubrics kiro --format json
```

The assistant can then:

- Read the spec file
- Evaluate each criterion using GO/WARN/NO-GO definitions
- Calculate weighted scores
- Provide improvement suggestions

## Example: Claude Code Integration

### CLAUDE.md Instructions

Add to your project's `CLAUDE.md`:

```markdown
## Spec-Driven Development

This project uses PlexusOne SDD methodology.

### Creating Specs

1. Run `devspec info plexusone --format json` to understand requirements
2. Create specs in order: PRD → TRD → PLAN → TASKS
3. Files requiring codebase access: TRD, PLAN, TASKS

### Evaluating Specs

1. Run `devspec rubrics plexusone --file prd --format json`
2. Evaluate each criterion against the rubric
3. Provide GO/WARN/NO-GO ratings with reasoning
```

### Assistant Prompt Pattern

```
When asked to create or evaluate specs:

1. Detect existing specs: `devspec check`
2. Get structure info: `devspec info <type> --format json`
3. For creation: Use sections and templates from info
4. For evaluation: Use `devspec rubrics <type> --format json`
5. Validate structure: `devspec validate`
```

## Codebase Access Awareness

The `requires_codebase` field tells the assistant when to analyze code:

```json
{
  "files": [
    {"name": "requirements", "requires_codebase": false},
    {"name": "design", "requires_codebase": true},
    {"name": "tasks", "requires_codebase": true}
  ]
}
```

**Before writing design.md or tasks.md:**

1. Explore existing codebase structure
2. Identify integration points
3. Reference specific files and functions

## Benefits

| Benefit | Description |
|---------|-------------|
| **No API costs** | LLM calls use user's existing quota |
| **Full context** | Assistant has full codebase access |
| **Interactive** | User can clarify and iterate |
| **Consistent** | Same rubrics ensure consistent evaluation |
