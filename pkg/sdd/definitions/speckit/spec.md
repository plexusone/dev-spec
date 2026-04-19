---
file: spec
display_name: Specification Document
sections:
  - name: Overview
    required: true
  - name: Goals
    required: true
  - name: Non-Goals
    required: false
  - name: User Stories
    required: true
  - name: Success Metrics
    required: false
criteria:
  - id: outcome_focus
    weight: 0.30
  - id: user_stories
    weight: 0.25
  - id: scope_clarity
    weight: 0.25
  - id: measurability
    weight: 0.20
---

# Specification Document Evaluation

## Criterion: outcome_focus (30%)

**Outcome Focus** - Spec describes WHAT not HOW

### GO
Specification is outcome-focused:
- Describes desired end state
- Technology-agnostic where possible
- Focuses on user value
- No implementation details

### WARN
Mostly outcome-focused:
- Some implementation leakage
- Generally describes what, not how

### NO-GO
Implementation-focused:
- Prescribes specific technologies
- Describes how to build, not what to build
- Design decisions in spec

---

## Criterion: user_stories (25%)

**User Stories** - User needs are clearly defined

### GO
User stories are well-defined:
- Clear user roles identified
- Specific needs documented
- Value proposition clear
- Acceptance criteria implied

### WARN
User stories present but incomplete:
- Generic user roles
- Some needs unclear

### NO-GO
User stories missing or poor:
- No user perspective
- Technical requirements only

---

## Criterion: scope_clarity (25%)

**Scope Clarity** - Boundaries are well-defined

### GO
Scope is crystal clear:
- Goals explicitly stated
- Non-goals documented
- Boundaries well-defined
- Out of scope items listed

### WARN
Scope mostly clear:
- Goals stated but boundaries fuzzy
- Some ambiguity remains

### NO-GO
Scope unclear:
- No clear boundaries
- Goals vague
- Scope creep likely

---

## Criterion: measurability (20%)

**Measurability** - Success can be measured

### GO
Success metrics defined:
- Quantifiable metrics
- Clear success criteria
- Measurable outcomes

### WARN
Some metrics present:
- Partial metrics
- Some criteria vague

### NO-GO
No measurable criteria:
- Cannot determine success
- No metrics defined
