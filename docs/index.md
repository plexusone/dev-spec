# devspec

**Spec-Driven Development Evaluation CLI**

devspec evaluates project specifications against multiple SDD (Spec-Driven Development) methodologies using LLM-as-a-Judge.

## Features

- **Multiple SDD Types**: Built-in support for Kiro, SpecKit, and PlexusOne methodologies
- **Deterministic Validation**: Structure validation without LLM calls
- **LLM-as-a-Judge**: AI-powered evaluation using customizable rubrics
- **Extensible**: Create custom SDD types or extend built-ins
- **Multiple Deployment Models**: Use with coding assistants, CI/CD, or as a service

## Supported SDD Types

| Type | Description | Files |
|------|-------------|-------|
| **Kiro** | AWS Kiro's three-phase workflow with EARS format | requirements, design, tasks |
| **SpecKit** | GitHub's spec-driven development | constitution, spec, plan, tasks |
| **PlexusOne** | Comprehensive enterprise SDD | MRD, PRD, TRD, PLAN, TASKS |

## Quick Example

```bash
# Initialize specs for a project
devspec init kiro ./my-project

# Check what SDD type is detected
devspec check ./my-project

# Validate structure
devspec validate ./my-project

# Get info for coding assistants
devspec info kiro --format json

# Evaluate with LLM
devspec evaluate ./my-project --llm anthropic
```

## Deployment Models

### Interactive (Coding Assistants)

Use `devspec info` and `devspec rubrics` to provide context to coding assistants like Claude Code, Cursor, or Kiro CLI. The assistant can then guide users in creating specs and self-evaluate using the rubrics.

### CI/CD Pipeline

Run `devspec evaluate` in GitHub Actions, Jenkins, or other CI systems for automated spec quality gates.

### Service (Future)

Deploy as an MCP server or HTTP API for integration with other tools.

## Output Formats

- **TOON** (default): Token-efficient format, ~40% fewer tokens than JSON
- **JSON**: Standard JSON with indentation
- **JSON-compact**: Minified JSON
- **Text**: Human-readable plain text
- **Markdown**: Formatted markdown documentation
