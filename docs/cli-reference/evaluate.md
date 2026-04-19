# devspec evaluate

Evaluate specs using LLM-as-a-Judge.

## Synopsis

```bash
devspec evaluate [path] [flags]
```

## Description

Evaluates spec files against the detected or specified SDD type using LLM-as-a-Judge methodology.

This command requires LLM configuration.

## Arguments

| Argument | Description | Default |
|----------|-------------|---------|
| `path` | Directory to evaluate | Current directory |

## Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--type` | `-t` | SDD type (auto-detected if not specified) |
| `--llm` | | LLM provider (anthropic, openai, etc.) |
| `--model` | | LLM model to use |
| `--dry-run` | | Show prompts without calling LLM |
| `--definition` | | Path to custom SDD definition |

## Environment Variables

| Variable | Description |
|----------|-------------|
| `ANTHROPIC_API_KEY` | API key for Anthropic |
| `OPENAI_API_KEY` | API key for OpenAI |

## Examples

```bash
# Evaluate with Anthropic Claude
export ANTHROPIC_API_KEY=your-key
devspec evaluate ./my-project --llm anthropic

# Evaluate with OpenAI GPT-4
export OPENAI_API_KEY=your-key
devspec evaluate --llm openai --model gpt-4

# Dry run to see prompts
devspec evaluate --dry-run

# Force specific SDD type
devspec evaluate --type kiro

# Use custom definition
devspec evaluate --definition .devspec/definitions/my-sdd/
```

## Output

```
SDD Type: kiro
Status: [WARN]
Score: 75.0%

## requirements.md [GO]
Score: 85.0%

  [GO] EARS Format Compliance (25%)
    All acceptance criteria use proper EARS format with clear triggers
  [WARN] User Story Quality (20%)
    Some user stories missing benefit clause
    Suggestions:
      - Add "so that" clause to user story in line 45
      - Specify concrete role instead of generic "user"
  [GO] Testability (20%)
    Requirements have measurable acceptance criteria

---
Summary: 3 files, 12 passed, 3 warned, 0 failed, 0 skipped
```

## Scoring

- Each criterion is evaluated as GO (100%), WARN (50%), or NO-GO (0%)
- File score = weighted average of criterion scores
- Overall score = average of file scores
- Overall status = worst status across all files
