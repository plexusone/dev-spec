# Core Concepts

## Spec-Driven Development (SDD)

Spec-Driven Development is a methodology where detailed specifications are written before implementation. This ensures:

- Clear requirements before coding
- Alignment between stakeholders
- Traceable implementation
- Quality gates at each phase

## SDD Types

An SDD Type defines a specific methodology with:

- **Files**: Which spec documents are required/optional
- **Sections**: Required sections within each file
- **Criteria**: Evaluation criteria with weights
- **Rubrics**: GO/WARN/NO-GO definitions for each criterion

## Evaluation Levels

| Level | Meaning |
|-------|---------|
| **GO** | Criterion fully met |
| **WARN** | Partially met, needs improvement |
| **NO-GO** | Not met, blocking issue |
| **SKIP** | Not applicable or not evaluated |

## Codebase Access

Some spec files require access to the existing codebase:

| Phase | Needs Codebase | Reason |
|-------|----------------|--------|
| Requirements/PRD | No | Defines WHAT to build |
| Design/TRD | Yes | Must integrate with existing architecture |
| Tasks | Yes | References specific code locations |

## Deterministic vs LLM Commands

**Deterministic** (no LLM):

- `check` - Detect SDD type
- `validate` - Structure validation
- `init` - Scaffold templates
- `info` - Get SDD metadata
- `rubrics` - Get evaluation rubrics

**LLM-powered**:

- `evaluate` - Quality evaluation using LLM-as-a-Judge

## Output Formats

| Format | Use Case |
|--------|----------|
| TOON | Token-efficient for LLM context (~40% smaller than JSON) |
| JSON | Programmatic processing |
| Markdown | Documentation, reports |
| Text | Human reading in terminal |
