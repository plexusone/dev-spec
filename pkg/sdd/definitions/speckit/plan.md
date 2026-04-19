---
file: plan
display_name: Plan Document
sections:
  - name: Approach
    required: true
  - name: Phases
    required: true
  - name: Risks
    required: false
  - name: Dependencies
    required: false
criteria:
  - id: phase_breakdown
    weight: 0.35
  - id: spec_alignment
    weight: 0.30
  - id: risk_awareness
    weight: 0.20
  - id: dependency_clarity
    weight: 0.15
---

# Plan Document Evaluation

## Criterion: phase_breakdown (35%)

**Phase Breakdown** - Implementation phases are well-defined

### GO
Phases are well-structured:
- Logical progression of work
- Clear milestones
- Deliverables per phase
- Reasonable scope per phase

### WARN
Phases present but incomplete:
- Some phases too large
- Milestones unclear

### NO-GO
Poor phase breakdown:
- No clear phases
- Monolithic plan
- No milestones

---

## Criterion: spec_alignment (30%)

**Spec Alignment** - Plan addresses specification

### GO
Plan covers entire specification:
- All spec items have plan coverage
- Goals addressed
- User stories mapped to phases

### WARN
Partial alignment:
- Most spec items covered
- Some gaps exist

### NO-GO
Poor alignment:
- Major spec items missing
- Plan doesn't match spec

---

## Criterion: risk_awareness (20%)

**Risk Awareness** - Risks are identified

### GO
Risks well-documented:
- Key risks identified
- Mitigation strategies noted
- Dependencies as risks considered

### WARN
Some risks noted:
- Partial risk coverage
- Mitigations incomplete

### NO-GO
Risks not addressed:
- No risk consideration
- Blind to potential issues

---

## Criterion: dependency_clarity (15%)

**Dependency Clarity** - Dependencies documented

### GO
Dependencies clear:
- External dependencies listed
- Internal dependencies noted
- Blocking items identified

### WARN
Some dependencies noted:
- Partial coverage
- Some unclear

### NO-GO
Dependencies missing:
- No dependency consideration
- Unknown blockers likely
