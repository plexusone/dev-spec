# Quick Start

This guide walks you through using devspec to create and evaluate project specifications.

## 1. Initialize Specs

Choose an SDD type and initialize your project:

```bash
# For Kiro (AWS's spec workflow)
devspec init kiro ./my-project

# For SpecKit (GitHub's approach)
devspec init speckit ./my-project

# For PlexusOne (comprehensive enterprise SDD)
devspec init plexusone ./my-project
```

This creates template files in the appropriate directory structure.

## 2. Fill in Your Specs

Edit the generated template files. Each file has sections marked with `[TODO: Add content]`.

For example, with Kiro:

```
my-project/
└── .kiro/
    └── specs/
        ├── requirements.md  # User stories with EARS acceptance criteria
        ├── design.md        # Architecture and components
        └── tasks.md         # Implementation tasks
```

## 3. Validate Structure

Check that your specs have the required sections:

```bash
devspec validate ./my-project
```

This is a deterministic check (no LLM calls) that verifies:

- Required files are present
- Required sections exist
- Basic structure compliance

## 4. Evaluate Quality (Optional)

For AI-powered quality evaluation:

```bash
# Using Anthropic
export ANTHROPIC_API_KEY=your-key
devspec evaluate ./my-project --llm anthropic

# Using OpenAI
export OPENAI_API_KEY=your-key
devspec evaluate ./my-project --llm openai --model gpt-4
```

## 5. View Results

Results show GO/WARN/NO-GO status for each criterion:

```
SDD Type: kiro
Status: [WARN]
Score: 75.0%

## requirements.md [GO]
Score: 85.0%

  [GO] EARS Format Compliance (25%)
    All acceptance criteria use proper EARS format
  [WARN] User Story Quality (20%)
    Some user stories missing "so that" clause
    Suggestions:
      - Add benefit clause to user story #3
```

## Output Formats

Choose your preferred format:

```bash
devspec validate ./my-project --format json      # JSON
devspec validate ./my-project --format markdown  # Markdown
devspec validate ./my-project --format text      # Plain text
devspec validate ./my-project                    # TOON (default)
```
