# Rubric Format

Rubrics define how spec files are evaluated.

## File Structure

Each file definition uses frontmatter + markdown:

```markdown
---
file: requirements
display_name: Requirements Document
sections:
  - name: Section Name
    required: true
    subsections: [Sub1, Sub2]
    fields: [Field1, Field2]
criteria:
  - id: criterion_id
    weight: 0.25
---

# Document Evaluation

## Criterion: criterion_id (25%)

**Title** - Brief description

### GO
What full compliance looks like...

### WARN
What partial compliance looks like...

### NO-GO
What failure looks like...
```

## Frontmatter Fields

### File Metadata

| Field | Type | Description |
|-------|------|-------------|
| `file` | string | File identifier (matches files[].name) |
| `display_name` | string | Human-readable name |

### Sections

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Section name (header text) |
| `required` | bool | Whether section is required |
| `subsections` | []string | Expected subsection names |
| `fields` | []string | Expected field names |
| `min_count` | int | Minimum occurrences |
| `format` | string | Expected format hint |

### Criteria

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique identifier |
| `weight` | float | Weight in score (0.0-1.0) |

Weights should sum to 1.0 for each file.

## Markdown Body

The markdown body contains rubric definitions:

```markdown
## Criterion: <id> (<weight>%)

**<Title>** - <Description>

### GO
<What full compliance looks like>
- Bullet points help
- Be specific

### WARN
<What partial compliance looks like>
- Common issues
- What's acceptable but not ideal

### NO-GO
<What failure looks like>
- Critical failures
- Blocking issues
```

## Parsing Rules

1. Criterion headers must match: `## Criterion: <id>`
2. Level headings (GO/WARN/NO-GO) must be `### <level>`
3. Content between level headings is captured as rubric text
4. Title is extracted from the line after `## Criterion:`

## Best Practices

### Weight Distribution

- Most important criteria: 25-30%
- Standard criteria: 15-20%
- Supporting criteria: 10-15%

### Writing Rubrics

**GO:**

- Be specific about what "good" looks like
- Include examples or patterns
- Make it objectively verifiable

**WARN:**

- Describe common partial compliance
- Explain what's missing vs GO
- Note what's acceptable

**NO-GO:**

- Describe clear failures
- Focus on blocking issues
- Be actionable (what to fix)

### Example: Well-Written Criterion

```markdown
## Criterion: testability (20%)

**Testability** - Requirements are verifiable

### GO
All requirements have measurable acceptance criteria:
- Specific numbers, thresholds, or conditions stated
- Clear pass/fail determination possible
- Can be automated or manually verified
- Example: "Response time < 200ms" not "fast response"

### WARN
Some requirements lack measurable criteria:
- Relative terms used ("faster", "better")
- Qualitative without definition
- Some criteria are testable, others vague

### NO-GO
Requirements cannot be tested:
- No acceptance criteria provided
- Purely subjective ("user-friendly", "intuitive")
- Cannot determine when requirement is met
```
