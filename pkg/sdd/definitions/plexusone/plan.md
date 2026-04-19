---
file: plan
display_name: Implementation Plan
sections:
  - name: Overview
    required: true
  - name: Milestones
    required: true
  - name: Timeline
    required: true
  - name: Resources
    required: false
  - name: Risks
    required: true
  - name: Dependencies
    required: true
criteria:
  - id: milestone_clarity
    weight: 0.30
  - id: trd_alignment
    weight: 0.25
  - id: risk_management
    weight: 0.25
  - id: feasibility
    weight: 0.20
---

# Implementation Plan Evaluation

## Criterion: milestone_clarity (30%)

**Milestone Clarity** - Milestones are well-defined

### GO
Clear milestones:
- Specific deliverables per milestone
- Measurable completion criteria
- Logical progression
- Reasonable scope

### WARN
Milestones need work:
- Some vague milestones
- Criteria unclear

### NO-GO
Poor milestones:
- No clear milestones
- Cannot track progress

---

## Criterion: trd_alignment (25%)

**TRD Alignment** - Plan implements TRD

### GO
Full alignment:
- All TRD components planned
- Architecture reflected
- Technical decisions honored

### WARN
Partial alignment:
- Most items covered
- Some gaps

### NO-GO
Poor alignment:
- TRD not reflected
- Major components missing

---

## Criterion: risk_management (25%)

**Risk Management** - Risks are managed

### GO
Comprehensive risk management:
- Risks identified
- Impact assessed
- Mitigations planned
- Contingencies defined

### WARN
Basic risk awareness:
- Some risks noted
- Mitigations incomplete

### NO-GO
Risks ignored:
- No risk consideration
- No mitigations

---

## Criterion: feasibility (20%)

**Feasibility** - Plan is achievable

### GO
Feasible plan:
- Resources identified
- Timeline realistic
- Dependencies manageable
- Team capacity considered

### WARN
Questionable feasibility:
- Aggressive timeline
- Resource gaps

### NO-GO
Not feasible:
- Unrealistic expectations
- Cannot be achieved
