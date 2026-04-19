# SDD Types Overview

devspec includes three built-in SDD types, each with different workflows and focus areas.

## Comparison

| Aspect | Kiro | SpecKit | PlexusOne |
|--------|------|---------|-----------|
| **Origin** | AWS Kiro | GitHub | PlexusOne |
| **Focus** | EARS format, traceability | Outcome-focused, executable | Comprehensive enterprise |
| **Required Files** | 3 | 3 | 4 |
| **Optional Files** | 0 | 1 | 1 |
| **Spec Directory** | `.kiro/specs/` | `.speckit/` | `specs/` |

## File Comparison

| Purpose | Kiro | SpecKit | PlexusOne |
|---------|------|---------|-----------|
| Principles | - | constitution.md (opt) | - |
| Market/Business | - | - | MRD.md (opt) |
| Requirements | requirements.md | spec.md | PRD.md |
| Design | design.md | - | TRD.md |
| Planning | - | plan.md | PLAN.md |
| Tasks | tasks.md | tasks.md | TASKS.md |

## Codebase Access Requirements

Files are marked whether they require codebase access for proper creation/evaluation:

| File Type | Needs Codebase | Reason |
|-----------|----------------|--------|
| Requirements/Spec/PRD | No | Define WHAT, not HOW |
| Design/TRD | Yes | Must integrate with existing architecture |
| Plan | Yes | Requires understanding current state |
| Tasks | Yes | Reference specific code locations |

## Choosing an SDD Type

**Choose Kiro if:**

- You want strict EARS format for acceptance criteria
- Traceability from requirements to tasks is critical
- You're using AWS Kiro CLI

**Choose SpecKit if:**

- You prefer outcome-focused specifications
- You want checkbox-based task tracking
- You're using GitHub's specify CLI

**Choose PlexusOne if:**

- You need comprehensive documentation
- Market analysis (MRD) is part of your process
- You're in an enterprise environment
