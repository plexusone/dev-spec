# PlexusOne SDD

Comprehensive enterprise spec-driven development.

## Overview

PlexusOne uses a full document hierarchy:

```
MRD (opt) → PRD → TRD → PLAN → TASKS
```

## Files

| File | Required | Needs Codebase | Purpose |
|------|----------|----------------|---------|
| MRD.md | No | No | Market requirements and business case |
| PRD.md | Yes | No | Product requirements and user needs |
| TRD.md | Yes | Yes | Technical architecture and design |
| PLAN.md | Yes | Yes | Implementation roadmap |
| TASKS.md | Yes | Yes | Detailed task breakdown |

## Directory Structure

```
project/
└── specs/
    ├── MRD.md   (optional)
    ├── PRD.md
    ├── TRD.md
    ├── PLAN.md
    └── TASKS.md
```

## Market Requirements Document (Optional)

Business justification and market analysis.

### Sections

- **Executive Summary**: Brief overview
- **Market Analysis**: Market size, trends
- **Competitive Landscape**: Competitors analysis
- **Business Opportunity**: Why now, differentiation
- **Target Audience**: Who we're building for
- **Success Metrics**: Business KPIs

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Market Understanding | 30% | Market is well-researched |
| Opportunity Clarity | 30% | Business opportunity is clear |
| Audience Definition | 25% | Target audience defined |
| Business Case | 15% | Financial viability shown |

## Product Requirements Document

### Required Sections

- **Executive Summary**: Brief overview
- **Problem Statement**: What problem we're solving
- **Goals and Objectives**: What we want to achieve
- **User Personas**: Who the users are
- **Requirements**: Detailed requirements
- **Success Metrics**: How to measure success

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Problem Clarity | 25% | Problem is well-defined |
| Requirements Quality | 30% | Requirements are well-written |
| User Focus | 25% | User needs are central |
| Success Criteria | 20% | Success is measurable |

## Technical Requirements Document

### Required Sections

- **Technical Overview**: Summary of technical approach
- **Architecture**: System architecture
- **Data Model**: Data structures
- **API Design**: API contracts
- **Security**: Security design
- **Performance**: Performance considerations
- **Testing Strategy**: Test approach

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Architecture Design | 30% | Architecture is well-designed |
| PRD Alignment | 25% | Design addresses PRD |
| Security Design | 20% | Security is addressed |
| Implementation Guidance | 25% | Enables implementation |

## Implementation Plan

Can be derived from Claude Code plan files or similar planning artifacts.

### Required Sections

- **Overview**: Implementation approach
- **Milestones**: Key milestones
- **Timeline**: Schedule
- **Resources**: Team/resource needs
- **Risks**: Risk identification
- **Dependencies**: What we depend on

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Milestone Clarity | 30% | Milestones are well-defined |
| TRD Alignment | 25% | Plan implements TRD |
| Risk Management | 25% | Risks are managed |
| Feasibility | 20% | Plan is achievable |

## Tasks Document

### Required Sections

- **Overview**: Task organization approach
- **Tasks**: Detailed task list

### Evaluation Criteria

| Criterion | Weight | Description |
|-----------|--------|-------------|
| Task Breakdown | 30% | Tasks properly decomposed |
| Traceability | 30% | Tasks trace to requirements |
| Estimation | 20% | Tasks have estimates |
| Dependencies | 20% | Dependencies documented |

## Example

```bash
# Initialize PlexusOne specs
devspec init plexusone ./my-project

# Validate structure
devspec validate ./my-project

# Get rubrics
devspec rubrics plexusone --format markdown
```
