# Custom SDD Types

Create your own SDD type from scratch.

## Directory Structure

```
.devspec/definitions/
└── my-sdd/
    ├── my-sdd.md      # Main SDD type definition
    ├── spec.md        # File definition with rubrics
    └── design.md      # File definition with rubrics
```

## Main Definition File

`my-sdd.md`:

```markdown
---
name: my-sdd
display_name: My Custom SDD
description: Custom spec-driven development workflow
spec_directory: specs
files:
  - name: spec
    patterns: ["spec.md", "specs/spec.md"]
    required: true
    requires_codebase: false
  - name: design
    patterns: ["design.md", "specs/design.md"]
    required: true
    requires_codebase: true
  - name: tasks
    patterns: ["tasks.md", "specs/tasks.md"]
    required: true
    requires_codebase: true
---

# My Custom SDD

Description of your SDD methodology.

## Workflow

1. **Spec**: Define what to build
2. **Design**: Design how to build it
3. **Tasks**: Break down implementation
```

## File Definition with Rubrics

`spec.md`:

```markdown
---
file: spec
display_name: Specification Document
sections:
  - name: Overview
    required: true
  - name: Requirements
    required: true
  - name: Success Metrics
    required: false
criteria:
  - id: clarity
    weight: 0.40
  - id: completeness
    weight: 0.35
  - id: testability
    weight: 0.25
---

# Specification Document Evaluation

## Criterion: clarity (40%)

**Clarity** - Requirements are clear and unambiguous

### GO
All requirements are crystal clear:
- No ambiguous language
- Specific and measurable
- Easy to understand

### WARN
Most requirements clear:
- Some ambiguity exists
- Generally understandable

### NO-GO
Requirements are unclear:
- Vague language throughout
- Cannot determine intent
```

## Using Your Custom Type

```bash
# Use the custom type
devspec init my-sdd ./my-project

# Validate against it
devspec validate ./my-project --type my-sdd

# Get info
devspec info my-sdd
```

## Search Paths

Definitions are searched in order:

1. `.devspec/definitions/` - Project-local
2. `~/.config/devspec/definitions/` - User-global
3. Built-in definitions

Project-local definitions override user-global and built-in.
