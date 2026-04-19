# Extending Built-in Types

Extend existing SDD types with additional files or criteria.

## Using `extends`

Create a new type that inherits from a built-in:

`.devspec/definitions/my-kiro/my-kiro.md`:

```markdown
---
name: my-kiro
extends: kiro
display_name: My Company Kiro
description: Kiro with compliance requirements
files:
  # Add new file type
  - name: compliance
    patterns: ["compliance.md", ".kiro/specs/compliance.md"]
    required: true
    requires_codebase: false
---

# My Company Kiro

Extended Kiro with compliance documentation.
```

Then create the rubrics for the new file:

`.devspec/definitions/my-kiro/compliance.md`:

```markdown
---
file: compliance
display_name: Compliance Document
sections:
  - name: Security Review
    required: true
  - name: Privacy Impact
    required: true
criteria:
  - id: security_compliance
    weight: 0.50
  - id: privacy_compliance
    weight: 0.50
---

# Compliance Document Evaluation

## Criterion: security_compliance (50%)

**Security Compliance** - Security requirements met

### GO
Full security compliance...

### WARN
Partial compliance...

### NO-GO
Not compliant...
```

## Inheritance Behavior

When extending:

| Aspect | Behavior |
|--------|----------|
| Files | Child files are added to parent files |
| File Definitions | Child overrides parent for same file name |
| Spec Directory | Child's directory used if specified |
| Description | Child's description used |

## Overriding Rubrics

To change criteria for an existing file, create a file definition with the same name:

`.devspec/definitions/kiro/requirements.md`:

```markdown
---
file: requirements
display_name: Requirements Document
sections:
  - name: Document Information
    required: true
  # ... your sections
criteria:
  - id: ears_format
    weight: 0.30  # Changed from 0.25
  - id: my_custom_criterion
    weight: 0.20  # Added new criterion
  # ... other criteria
---

# Requirements Document Evaluation

## Criterion: ears_format (30%)

Your custom GO/WARN/NO-GO definitions...

## Criterion: my_custom_criterion (20%)

Your new criterion rubrics...
```

## Examples

### Add Security Review to SpecKit

```markdown
---
name: secure-speckit
extends: speckit
display_name: Secure SpecKit
files:
  - name: security
    patterns: ["security.md", ".speckit/security.md"]
    required: true
    requires_codebase: true
---
```

### Add ADRs to PlexusOne

```markdown
---
name: plexusone-adr
extends: plexusone
display_name: PlexusOne with ADRs
files:
  - name: adr
    patterns: ["ADR-*.md", "specs/adr/*.md"]
    required: false
    requires_codebase: true
---
```
