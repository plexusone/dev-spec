# LLM Evaluation Guide

How to get the most out of LLM-as-a-Judge evaluation.

## How It Works

1. **Load Rubrics**: devspec loads criteria and GO/WARN/NO-GO definitions
2. **Build Prompt**: Spec content + rubrics combined into evaluation prompt
3. **LLM Evaluation**: LLM evaluates each criterion
4. **Parse Results**: Response parsed into structured results
5. **Calculate Scores**: Weighted scores computed

## Choosing an LLM Provider

### Anthropic Claude

```bash
export ANTHROPIC_API_KEY=your-key
devspec evaluate --llm anthropic --model claude-sonnet-4-20250514
```

**Pros**: Strong reasoning, follows rubrics well
**Best for**: Complex evaluations, nuanced criteria

### OpenAI GPT-4

```bash
export OPENAI_API_KEY=your-key
devspec evaluate --llm openai --model gpt-4
```

**Pros**: Fast, good at structured output
**Best for**: High-volume evaluations

### Model Selection

| Use Case | Recommended Model |
|----------|-------------------|
| Development/testing | GPT-3.5, Claude Haiku |
| Production evaluation | GPT-4, Claude Sonnet |
| Critical decisions | Claude Opus |

## Interpreting Results

### Status Meanings

| Status | Score | Meaning |
|--------|-------|---------|
| GO | 100% | Criterion fully met |
| WARN | 50% | Partially met, needs improvement |
| NO-GO | 0% | Not met, blocking issue |
| SKIP | N/A | Not evaluated (missing file, etc.) |

### Score Calculation

```
File Score = Σ(criterion_score × criterion_weight)

Where:
  GO = 1.0
  WARN = 0.5
  NO-GO = 0.0
```

Example:
```
EARS Format (25%): GO    → 0.25 × 1.0 = 0.25
User Stories (20%): WARN → 0.20 × 0.5 = 0.10
Testability (20%): GO   → 0.20 × 1.0 = 0.20
Completeness (20%): WARN → 0.20 × 0.5 = 0.10
No Impl (15%): GO       → 0.15 × 1.0 = 0.15
                                 Total = 0.80 (80%)
```

### Understanding Reasoning

Each criterion includes reasoning:

```json
{
  "id": "ears_format",
  "status": "WARN",
  "reasoning": "Most acceptance criteria use EARS format, but REQ-003 and REQ-007 use informal language without clear triggers.",
  "suggestions": [
    "Rewrite REQ-003 using WHEN...THEN...SHALL format",
    "Add specific trigger condition to REQ-007"
  ]
}
```

## Improving Scores

### Common Issues by Criterion

**EARS Format (Kiro)**

| Issue | Solution |
|-------|----------|
| Missing SHALL | Add "SHALL" to system behavior |
| Vague triggers | Be specific: "clicks button" not "interacts" |
| Mixed formats | Standardize all criteria to EARS |

**User Stories**

| Issue | Solution |
|-------|----------|
| Generic roles | "authenticated admin" not "user" |
| Missing benefit | Always include "so that" clause |
| Technical focus | Describe user value, not implementation |

**Testability**

| Issue | Solution |
|-------|----------|
| Relative terms | Use specific numbers |
| Subjective criteria | Define measurable outcomes |
| No acceptance criteria | Add verifiable conditions |

## Dry Run Mode

Preview prompts without LLM calls:

```bash
devspec evaluate --dry-run
```

This shows:

- What files will be evaluated
- Prompts that would be sent
- Expected response format

Useful for:

- Debugging custom rubrics
- Understanding evaluation process
- Estimating token costs

## Cost Optimization

### Reduce Token Usage

1. **Use TOON format** for context (40% smaller than JSON)
2. **Evaluate changed files only** in CI
3. **Cache results** for unchanged specs

### Batch Evaluations

```bash
# Evaluate specific file only
devspec evaluate --file requirements

# Skip files that haven't changed
devspec evaluate --changed-only
```

### Model Tiering

```bash
# Quick validation with smaller model
devspec evaluate --llm openai --model gpt-3.5-turbo

# Full evaluation with capable model
devspec evaluate --llm anthropic --model claude-sonnet-4-20250514
```

## Troubleshooting

### Low Scores Despite Good Specs

1. Check rubric alignment - your style may differ from rubric expectations
2. Review reasoning - LLM explains its evaluation
3. Consider custom rubrics - extend or override built-ins

### Inconsistent Results

1. Use deterministic settings if available
2. Run multiple times and average
3. Use more capable models

### Missing Context

1. Ensure all referenced files exist
2. Check that spec is self-contained
3. Provide necessary background in spec itself
