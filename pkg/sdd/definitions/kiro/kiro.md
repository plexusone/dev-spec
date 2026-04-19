---
name: kiro
display_name: Kiro Spec-Driven Development
description: AWS Kiro's three-phase spec workflow with EARS format
spec_directory: .kiro/specs
files:
  - name: requirements
    patterns: ["requirements.md", "**/requirements.md"]
    required: true
    requires_codebase: false
  - name: design
    patterns: ["design.md", "**/design.md"]
    required: true
    requires_codebase: true
  - name: tasks
    patterns: ["tasks.md", "**/tasks.md"]
    required: true
    requires_codebase: true
---

# Kiro Spec-Driven Development

Kiro uses a three-phase workflow: Requirements -> Design -> Tasks.

## Key Characteristics

- **EARS Format**: All acceptance criteria use Easy Approach to Requirements Syntax
- **User Stories**: Requirements structured as "As a [role], I want [feature], so that [benefit]"
- **Traceability**: Tasks link back to requirements and design components

## Workflow

1. **Requirements Phase**: Define what the system should do using user stories and EARS acceptance criteria
2. **Design Phase**: Architect the solution with components, interfaces, and data models (codebase access recommended for integration context)
3. **Tasks Phase**: Break down implementation into traceable, ordered tasks (codebase access required)

## Codebase Access

| File | Codebase Required | Reason |
|------|-------------------|--------|
| requirements.md | No | Defines WHAT to build, independent of existing code |
| design.md | Yes | Architecture must integrate with existing codebase |
| tasks.md | Yes | Tasks reference specific files, functions, and code patterns |
