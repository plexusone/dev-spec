---
file: tasks
display_name: Tasks Document
sections:
  - name: Implementation Overview
    required: true
  - name: Implementation Plan
    required: true
    format: phased_tasks
criteria:
  - id: task_granularity
    weight: 0.25
  - id: requirements_traceability
    weight: 0.25
  - id: dependency_order
    weight: 0.20
  - id: testing_included
    weight: 0.15
  - id: design_coverage
    weight: 0.15
---

# Tasks Document Evaluation

## Criterion: task_granularity (25%)

**Task Granularity** - Tasks are appropriately sized

### GO
Tasks are well-sized for implementation:
- Each task is completable in a reasonable timeframe
- Tasks are specific and actionable
- Clear definition of done for each task
- Neither too large nor too granular

### WARN
Task sizing inconsistent:
- Some tasks too large or vague
- Some tasks too granular
- Mixed levels of detail

### NO-GO
Poor task granularity:
- Tasks are too vague to implement
- No clear task boundaries
- Monolithic tasks without breakdown

---

## Criterion: requirements_traceability (25%)

**Requirements Traceability** - Tasks link to requirements

### GO
Full traceability to requirements:
- Each task references requirement(s) it fulfills
- All requirements have implementing tasks
- Clear mapping between tasks and user stories

### WARN
Partial traceability:
- Some tasks lack requirement references
- Some requirements not clearly covered

### NO-GO
No traceability:
- Tasks don't reference requirements
- Cannot verify requirement coverage

---

## Criterion: dependency_order (20%)

**Dependency Order** - Tasks are properly ordered

### GO
Clear dependency management:
- Tasks ordered for logical implementation
- Dependencies explicitly stated
- Parallel work opportunities identified
- Critical path understood

### WARN
Order mostly clear:
- Some dependencies unclear
- Order generally logical

### NO-GO
No clear ordering:
- Tasks in random order
- Dependencies not considered
- Implementation sequence unclear

---

## Criterion: testing_included (15%)

**Testing Included** - Testing tasks are present

### GO
Comprehensive testing coverage:
- Unit test tasks included
- Integration test tasks present
- Test tasks aligned with features
- Testing strategy from design reflected

### WARN
Basic testing tasks:
- Some testing mentioned
- Not all features have test tasks

### NO-GO
Testing missing:
- No test tasks
- Testing not considered

---

## Criterion: design_coverage (15%)

**Design Coverage** - Tasks cover design components

### GO
All design components have tasks:
- Each component has implementation tasks
- Interfaces have implementation tasks
- Data models have tasks

### WARN
Most design covered:
- Some components lack explicit tasks
- Minor gaps in coverage

### NO-GO
Poor design coverage:
- Major design components missing tasks
- Design not reflected in tasks
