---
name: plexusone
display_name: PlexusOne SDD
description: PlexusOne's comprehensive spec-driven development
spec_directory: specs
files:
  - name: mrd
    patterns: ["MRD.md", "specs/MRD.md"]
    required: false
    requires_codebase: false
  - name: prd
    patterns: ["PRD.md", "specs/PRD.md"]
    required: true
    requires_codebase: false
  - name: trd
    patterns: ["TRD.md", "specs/TRD.md"]
    required: true
    requires_codebase: true
  - name: plan
    patterns: ["PLAN.md", "specs/PLAN.md"]
    required: true
    requires_codebase: true
  - name: tasks
    patterns: ["TASKS.md", "specs/TASKS.md"]
    required: true
    requires_codebase: true
---

# PlexusOne SDD

PlexusOne uses a comprehensive document hierarchy:
MRD (optional) -> PRD -> TRD -> PLAN (optional) -> TASKS

## Key Characteristics

- **MRD**: Market Requirements Document - business opportunity and market analysis
- **PRD**: Product Requirements Document - user needs, features, acceptance criteria
- **TRD**: Technical Requirements Document - architecture and design decisions
- **PLAN**: Implementation plan with milestones
- **TASKS**: Detailed, traceable task breakdown

## Workflow

1. **MRD** (optional): Define market opportunity and business case
2. **PRD**: Specify product requirements and user needs
3. **TRD**: Design technical architecture and solutions (codebase access required)
4. **PLAN**: Create implementation roadmap (codebase access required) - can be derived from Claude Code plan files
5. **TASKS**: Break down into implementable tasks (codebase access required)

## Codebase Access

| File | Required | Codebase Required | Reason |
|------|----------|-------------------|--------|
| MRD.md | No | No | Market analysis independent of implementation |
| PRD.md | Yes | No | Product requirements define WHAT, not HOW |
| TRD.md | Yes | Yes | Technical design must integrate with existing architecture |
| PLAN.md | Yes | Yes | Implementation roadmap; can be copied from Claude Code plan file |
| TASKS.md | Yes | Yes | Tasks must reference specific files, modules, and APIs |
