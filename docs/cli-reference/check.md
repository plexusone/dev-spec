# devspec check

Detect SDD type in a directory.

## Synopsis

```bash
devspec check [path] [flags]
```

## Description

Scans the specified directory for spec files and identifies which SDD type is present based on file patterns and structure.

This is a deterministic command that does not use LLM calls.

## Arguments

| Argument | Description | Default |
|----------|-------------|---------|
| `path` | Directory to check | Current directory |

## Examples

```bash
# Check current directory
devspec check

# Check specific directory
devspec check ./my-project

# Output as JSON
devspec check ./my-project --format json
```

## Output

```
Detected: kiro
  Display Name: Kiro Spec-Driven Development
  Description: AWS Kiro's three-phase spec workflow with EARS format
  Spec Directory: .kiro/specs
  Matched File: ./my-project/.kiro/specs/requirements.md
  Confidence: 100%
```

## JSON Output

```json
{
  "detected": true,
  "sdd_type": "kiro",
  "display_name": "Kiro Spec-Driven Development",
  "description": "AWS Kiro's three-phase spec workflow with EARS format",
  "spec_directory": ".kiro/specs",
  "matched_file": "./my-project/.kiro/specs/requirements.md",
  "confidence": 1.0
}
```

## See Also

- [validate](validate.md) - Validate spec structure
- [info](info.md) - Get SDD type information
