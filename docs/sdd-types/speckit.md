# SpecKit SDD

GitHub's spec-driven development with specify CLI.

## Overview

SpecKit follows an outcome-focused workflow:

```
Constitution (opt) → Spec → Plan → Tasks
```

## Files

| File | Required | Needs Codebase | Purpose |
|------|----------|----------------|---------|
| constitution.md | No | No | Project principles and constraints |
| spec.md | Yes | No | What the system should do |
| plan.md | Yes | Yes | Implementation phases |
| tasks.md | Yes | Yes | Detailed task checklist |

## Directory Structure

```
project/
└── .speckit/
    ├── constitution.md  (optional)
    ├── spec.md
    ├── plan.md
    └── tasks.md
```

## Constitution Document (Optional)

Defines guiding principles and constraints for the project.

### Sections

- **Purpose**: Why this project exists
- **Principles**: Guiding technical principles
- **Constraints**: Technical or business constraints
- **Quality Standards**: Quality expectations

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Clarity | 40% | Principles are clear and understandable |
| Actionability | 35% | Principles guide decisions |
| Completeness | 25% | Covers key areas |

## Specification Document

### Required Sections

- **Overview**: What the system does
- **Goals**: What success looks like
- **Non-Goals**: What's explicitly out of scope
- **User Stories**: User needs
- **Success Metrics**: How to measure success

### Key Principle: What Not How

The spec focuses on **outcomes**, not implementation details:

```markdown
# Good (What)
Users can search for products by name and filter by category.

# Bad (How)
Use Elasticsearch with a React frontend and debounced input.
```

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Outcome Focus | 30% | Describes WHAT not HOW |
| User Stories | 25% | User needs are defined |
| Scope Clarity | 25% | Boundaries are clear |
| Measurability | 20% | Success can be measured |

## Plan Document

### Required Sections

- **Approach**: High-level implementation approach
- **Phases**: Implementation phases with milestones
- **Risks**: Identified risks
- **Dependencies**: External/internal dependencies

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Phase Breakdown | 35% | Phases are well-defined |
| Spec Alignment | 30% | Plan covers specification |
| Risk Awareness | 20% | Risks are identified |
| Dependency Clarity | 15% | Dependencies documented |

## Tasks Document

### Key Feature: Checkbox Tracking

Tasks use markdown checkboxes for progress tracking:

```markdown
## Phase 1: Foundation

- [x] Set up project structure
- [x] Configure build system
- [ ] Implement core data models
- [ ] Add basic API endpoints
```

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Task Clarity | 30% | Tasks are clear and actionable |
| Plan Coverage | 30% | Tasks cover the plan |
| Checkbox Format | 20% | Uses checkbox tracking |
| Task Ordering | 20% | Tasks are logically ordered |

## Example

```bash
# Initialize SpecKit specs
devspec init speckit ./my-project

# Validate structure
devspec validate ./my-project

# Get rubrics
devspec rubrics speckit --format markdown
```
