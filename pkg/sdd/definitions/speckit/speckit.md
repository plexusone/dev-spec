---
name: speckit
display_name: GitHub SpecKit
description: GitHub's spec-driven development with specify CLI
spec_directory: .speckit
files:
  - name: constitution
    patterns: ["constitution.md", ".speckit/constitution.md"]
    required: false
    requires_codebase: false
  - name: spec
    patterns: ["spec.md", ".speckit/spec.md"]
    required: true
    requires_codebase: false
  - name: plan
    patterns: ["plan.md", ".speckit/plan.md"]
    required: true
    requires_codebase: true
  - name: tasks
    patterns: ["tasks.md", ".speckit/tasks.md"]
    required: true
    requires_codebase: true
---

# GitHub SpecKit

SpecKit follows: Constitution (optional) -> Specification -> Plan -> Tasks -> Implementation

## Key Characteristics

- **What Not How**: spec.md focuses on outcomes, not tech stack
- **Executable Specs**: Specifications directly generate implementations
- **Checkbox Tracking**: tasks.md uses [x]/[ ] for progress
- **Constitution**: Optional guiding principles for the project

## Workflow

1. **Constitution** (optional): Define project principles and constraints
2. **Specification**: Describe what the system should do
3. **Plan**: Break down into implementation phases (codebase access recommended)
4. **Tasks**: Detailed task list with checkboxes for tracking (codebase access required)

## Codebase Access

| File | Codebase Required | Reason |
|------|-------------------|--------|
| constitution.md | No | High-level principles independent of implementation |
| spec.md | No | Defines WHAT to build, outcome-focused |
| plan.md | Yes | Planning requires understanding existing architecture |
| tasks.md | Yes | Tasks must reference specific code locations and patterns |
