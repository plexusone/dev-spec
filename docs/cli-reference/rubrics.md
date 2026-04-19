# devspec rubrics

Get evaluation rubrics for coding assistants.

## Synopsis

```bash
devspec rubrics <type> [flags]
```

## Description

Returns evaluation rubrics and criteria for an SDD type. Designed for coding assistants to evaluate spec files themselves using the rubrics, rather than having the CLI make LLM calls.

## Arguments

| Argument | Description |
|----------|-------------|
| `type` | SDD type (kiro, speckit, plexusone) |

## Flags

| Flag | Description |
|------|-------------|
| `--file` | Get rubrics for specific file only |

## Examples

```bash
# Get all rubrics for Kiro
devspec rubrics kiro

# Get rubrics for specific file
devspec rubrics kiro --file requirements

# JSON output
devspec rubrics kiro --format json

# Markdown for documentation
devspec rubrics plexusone --format markdown
```

## Output Structure

Each criterion includes:

| Field | Description |
|-------|-------------|
| `id` | Criterion identifier |
| `title` | Human-readable title |
| `weight` | Weight in overall score (0.0-1.0) |
| `go` | Description of GO (pass) condition |
| `warn` | Description of WARN (partial) condition |
| `no_go` | Description of NO-GO (fail) condition |

## Example Output

```json
{
  "sdd_type": "kiro",
  "files": [
    {
      "file": "requirements",
      "criteria": [
        {
          "id": "ears_format",
          "title": "EARS Format Compliance",
          "weight": 0.25,
          "go": "All acceptance criteria use proper EARS format...",
          "warn": "Most criteria use EARS but some are informal...",
          "no_go": "No EARS format used or criteria are vague..."
        }
      ]
    }
  ]
}
```

## Use with Coding Assistants

Coding assistants can use rubrics to:

1. Self-evaluate spec quality
2. Provide specific improvement suggestions
3. Calculate scores using weights
4. Generate evaluation reports

```bash
# Coding assistant workflow
devspec rubrics kiro --format json > rubrics.json
# Assistant reads spec file and rubrics
# Assistant evaluates each criterion
# Assistant provides GO/WARN/NO-GO ratings
```
