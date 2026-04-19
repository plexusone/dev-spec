# Kiro SDD

AWS Kiro's three-phase spec workflow with EARS format.

## Overview

Kiro uses a structured three-phase workflow:

```
Requirements → Design → Tasks
```

## Files

| File | Required | Needs Codebase | Purpose |
|------|----------|----------------|---------|
| requirements.md | Yes | No | User stories with EARS acceptance criteria |
| design.md | Yes | Yes | Architecture and component design |
| tasks.md | Yes | Yes | Implementation task breakdown |

## Directory Structure

```
project/
└── .kiro/
    └── specs/
        ├── requirements.md
        ├── design.md
        └── tasks.md
```

## Requirements Document

### Required Sections

- **Document Information**: Feature name, version, author
- **Introduction**: Feature summary, business value, scope
- **Requirements**: User stories with EARS acceptance criteria
- **Success Criteria**: How to measure success

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| EARS Format | 25% | Acceptance criteria use EARS notation |
| User Stories | 20% | Stories follow "As a/I want/So that" format |
| Testability | 20% | Requirements are measurable and verifiable |
| Completeness | 20% | Covers happy path, errors, and NFRs |
| No Implementation | 15% | Focuses on WHAT, not HOW |

### EARS Format

EARS (Easy Approach to Requirements Syntax) provides templates:

```
WHEN [trigger] THEN [system] SHALL [response]
IF [condition] THEN [system] SHALL [behavior]
WHILE [state] [system] SHALL [behavior]
WHERE [context] [system] SHALL [behavior]
```

**Example:**
```markdown
WHEN a user clicks the "Submit" button 
THEN the system SHALL validate all form fields 
AND display validation errors within 200ms
```

## Design Document

### Required Sections

- **Overview**: Design goals, key decisions
- **Architecture**: System architecture
- **Components and Interfaces**: Component definitions and APIs
- **Data Models**: Data structures and schemas
- **Security Considerations**: Auth, data protection
- **Testing Strategy**: Test approach

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Architecture Clarity | 25% | Architecture is well-defined |
| Requirements Alignment | 20% | Design addresses all requirements |
| Component Interfaces | 20% | APIs and contracts defined |
| Security Coverage | 15% | Security considerations addressed |
| Implementation Ready | 20% | Enough detail to implement |

## Tasks Document

### Required Sections

- **Implementation Overview**: Summary of implementation approach
- **Implementation Plan**: Phased task breakdown

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Task Granularity | 25% | Tasks are appropriately sized |
| Requirements Traceability | 25% | Tasks link to requirements |
| Dependency Order | 20% | Tasks are properly ordered |
| Testing Included | 15% | Test tasks are present |
| Design Coverage | 15% | Tasks cover design components |

## Example

```bash
# Initialize Kiro specs
devspec init kiro ./my-project

# Validate structure
devspec validate ./my-project

# Get rubrics for self-evaluation
devspec rubrics kiro --format markdown
```
