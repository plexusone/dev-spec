---
file: tasks
display_name: Tasks Document
sections:
  - name: Overview
    required: true
  - name: Tasks
    required: true
criteria:
  - id: task_breakdown
    weight: 0.30
  - id: traceability
    weight: 0.30
  - id: estimation
    weight: 0.20
  - id: dependencies
    weight: 0.20
---

# Tasks Document Evaluation

## Criterion: task_breakdown (30%)

**Task Breakdown** - Tasks are properly decomposed

### GO
Well-decomposed tasks:
- Appropriate granularity
- Each task is actionable
- Clear deliverables
- Reasonable scope

### WARN
Partial breakdown:
- Some tasks too large
- Mixed granularity

### NO-GO
Poor breakdown:
- Tasks too vague
- Monolithic items
- Not actionable

---

## Criterion: traceability (30%)

**Traceability** - Tasks trace to requirements

### GO
Full traceability:
- Tasks reference PRD/TRD items
- All requirements covered
- Clear mapping maintained

### WARN
Partial traceability:
- Some references present
- Gaps in coverage

### NO-GO
No traceability:
- No requirement references
- Cannot verify coverage

---

## Criterion: estimation (20%)

**Estimation** - Tasks have estimates

### GO
Tasks are estimated:
- Effort estimates present
- Reasonable estimates
- Complexity noted

### WARN
Partial estimates:
- Some tasks estimated
- Estimates inconsistent

### NO-GO
No estimates:
- Cannot plan timeline
- No effort indication

---

## Criterion: dependencies (20%)

**Dependencies** - Dependencies documented

### GO
Dependencies clear:
- Task dependencies noted
- External dependencies listed
- Blocking items identified
- Order logical

### WARN
Some dependencies:
- Partial coverage
- Some implicit

### NO-GO
Dependencies missing:
- No dependency tracking
- Order unclear
