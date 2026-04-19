---
file: prd
display_name: Product Requirements Document
sections:
  - name: Executive Summary
    required: true
  - name: Problem Statement
    required: true
  - name: Goals and Objectives
    required: true
  - name: User Personas
    required: false
  - name: Requirements
    required: true
  - name: Success Metrics
    required: true
criteria:
  - id: problem_clarity
    weight: 0.25
  - id: requirements_quality
    weight: 0.30
  - id: user_focus
    weight: 0.25
  - id: success_criteria
    weight: 0.20
---

# Product Requirements Document Evaluation

## Criterion: problem_clarity (25%)

**Problem Clarity** - Problem is well-defined

### GO
Problem statement is clear:
- Specific problem identified
- Impact quantified where possible
- Root cause understood
- Scope of problem defined

### WARN
Problem partially defined:
- Problem stated but vague
- Impact unclear

### NO-GO
Problem unclear:
- No clear problem statement
- Solution looking for a problem

---

## Criterion: requirements_quality (30%)

**Requirements Quality** - Requirements are well-written

### GO
High-quality requirements:
- Specific and measurable
- Testable acceptance criteria
- Priority indicated
- Dependencies noted

### WARN
Requirements need work:
- Some vague requirements
- Missing acceptance criteria

### NO-GO
Poor requirements:
- Vague or ambiguous
- Not testable
- No acceptance criteria

---

## Criterion: user_focus (25%)

**User Focus** - User needs are central

### GO
User-centric requirements:
- User personas defined
- User journeys documented
- User value clear
- User feedback incorporated

### WARN
Some user focus:
- Basic user consideration
- Missing personas or journeys

### NO-GO
Not user-focused:
- Technical requirements only
- No user perspective

---

## Criterion: success_criteria (20%)

**Success Criteria** - Success is measurable

### GO
Clear success criteria:
- Quantifiable metrics
- Baseline defined
- Target values set
- Measurement method clear

### WARN
Partial criteria:
- Some metrics defined
- Targets unclear

### NO-GO
No success criteria:
- Cannot measure success
- No metrics defined
