---
file: requirements
display_name: Requirements Document
sections:
  - name: Document Information
    required: true
    fields: [Feature Name, Version, Author]
  - name: Introduction
    required: true
    subsections: [Feature Summary, Business Value, Scope]
  - name: Requirements
    required: true
    min_count: 1
    format: user_story_with_ears
  - name: Non-Functional Requirements
    required: false
  - name: Success Criteria
    required: true
criteria:
  - id: ears_format
    weight: 0.25
  - id: user_stories
    weight: 0.20
  - id: testability
    weight: 0.20
  - id: completeness
    weight: 0.20
  - id: no_implementation
    weight: 0.15
---

# Requirements Document Evaluation

## Criterion: ears_format (25%)

**EARS Format Compliance** - Acceptance criteria use EARS notation (WHEN/IF/WHILE/WHERE...SHALL)

### GO
All acceptance criteria use proper EARS format with clear triggers and system responses:
- WHEN [event/trigger] THEN [system] SHALL [response]
- IF [condition] THEN [system] SHALL [behavior]
- WHILE [ongoing condition] [system] SHALL [continuous behavior]
- WHERE [context] [system] SHALL [contextual behavior]

### WARN
Most criteria use EARS but some are informal or incomplete:
- Missing SHALL keyword in some criteria
- Vague triggers or conditions
- Mixing EARS with informal language

### NO-GO
No EARS format used or criteria are vague/untestable:
- Free-form text without structure
- Subjective criteria ("should be fast", "user-friendly")
- No clear system behavior defined

---

## Criterion: user_stories (20%)

**User Story Quality** - User stories follow standard format

### GO
All user stories have clear role, feature, and benefit:
- "As a [specific role], I want [concrete feature], so that [measurable benefit]"
- Roles are specific (not "user" but "authenticated customer")
- Benefits explain business value

### WARN
Some user stories missing role or benefit:
- Generic roles ("As a user...")
- Missing "so that" clause
- Benefits are implementation-focused

### NO-GO
No user stories or format not followed:
- Requirements stated without user context
- Technical specifications instead of user needs
- No clear value proposition

---

## Criterion: testability (20%)

**Testability** - Each requirement is testable and verifiable

### GO
All requirements have measurable acceptance criteria:
- Specific numbers, thresholds, or conditions
- Clear pass/fail criteria
- Can be automated or manually verified

### WARN
Some requirements are subjective or hard to test:
- Relative terms ("faster than before")
- Qualitative without definition ("high quality")

### NO-GO
Requirements are vague and untestable:
- No acceptance criteria
- Purely aspirational ("best in class")
- Cannot determine when requirement is met

---

## Criterion: completeness (20%)

**Completeness** - Covers functional, non-functional, and edge cases

### GO
Covers happy path, error cases, and NFRs:
- Normal user flows documented
- Error handling specified
- Performance, security, usability addressed
- Edge cases considered

### WARN
Missing some edge cases or NFRs:
- Only happy path detailed
- NFRs mentioned but not specified
- Some error cases missing

### NO-GO
Only covers happy path, missing critical scenarios:
- No error handling
- No NFRs
- No edge cases

---

## Criterion: no_implementation (15%)

**Implementation Independence** - Focuses on WHAT not HOW

### GO
Pure requirements without implementation details:
- Describes desired outcomes
- Technology-agnostic where possible
- Leaves design decisions to design phase

### WARN
Some implementation leakage:
- Specific technology mentioned unnecessarily
- UI details in requirements
- Database schema hints

### NO-GO
Heavily prescribes implementation:
- Specific frameworks required
- Code structure dictated
- No room for design alternatives
