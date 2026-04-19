# Writing Good Specs

Best practices for writing high-quality specifications.

## General Principles

### 1. Be Specific, Not Vague

```markdown
# Bad
The system should be fast and user-friendly.

# Good
The search API SHALL return results within 200ms for 95% of queries.
Users SHALL complete the checkout flow in 3 steps or fewer.
```

### 2. Focus on WHAT, Not HOW

```markdown
# Bad (Implementation details)
Use Redis for caching with a 5-minute TTL.
Implement using React with Redux state management.

# Good (Requirements)
Frequently accessed data SHALL be cached to reduce latency.
The UI SHALL update in real-time when data changes.
```

### 3. Make Everything Testable

```markdown
# Bad (Untestable)
The application should have good performance.
Users should find the interface intuitive.

# Good (Testable)
Page load time SHALL be under 2 seconds on 3G connections.
New users SHALL complete onboarding without documentation.
```

## EARS Format (Kiro)

Use EARS (Easy Approach to Requirements Syntax) for acceptance criteria:

### Templates

| Type | Template | Use When |
|------|----------|----------|
| Event-driven | WHEN [trigger] THEN [system] SHALL [response] | User action triggers response |
| Conditional | IF [condition] THEN [system] SHALL [behavior] | Behavior depends on state |
| State-driven | WHILE [state] [system] SHALL [behavior] | Ongoing behavior during state |
| Optional | WHERE [feature enabled] [system] SHALL [behavior] | Feature flags |

### Examples

```markdown
# Event-driven
WHEN a user submits the login form
THEN the system SHALL validate credentials within 500ms
AND redirect to the dashboard on success
AND display an error message on failure

# Conditional
IF the user's cart total exceeds $100
THEN the system SHALL apply free shipping
AND display the savings to the user

# State-driven
WHILE a file upload is in progress
the system SHALL display a progress indicator
AND allow the user to cancel the upload

# Optional
WHERE two-factor authentication is enabled
the system SHALL require a verification code after password entry
```

## User Stories

### Format

```markdown
As a [specific role],
I want [concrete feature],
so that [measurable benefit].
```

### Good Examples

```markdown
As an authenticated customer,
I want to save items to a wishlist,
so that I can purchase them later without searching again.

As a team administrator,
I want to bulk-invite users via CSV upload,
so that I can onboard large teams in minutes instead of hours.
```

### Common Mistakes

| Mistake | Example | Fix |
|---------|---------|-----|
| Generic role | "As a user" | "As an authenticated customer" |
| Vague feature | "I want better search" | "I want to filter search by date range" |
| No benefit | Missing "so that" | Always include business value |

## Document Structure

### Requirements Document

```markdown
# Requirements Document

## Document Information
- **Feature Name**: User Authentication
- **Version**: 1.0
- **Author**: Jane Smith

## Introduction
### Feature Summary
Brief description of the feature.

### Business Value
Why this feature matters to the business.

### Scope
What's included and excluded.

## Requirements

### REQ-001: User Login
As an existing user...
**Acceptance Criteria:**
- WHEN user enters valid credentials...
- IF credentials are invalid...

## Non-Functional Requirements
- Performance: Response time < 500ms
- Security: Passwords hashed with bcrypt

## Success Criteria
- 95% of users complete login on first attempt
- Zero security vulnerabilities in penetration testing
```

## Common Pitfalls

### 1. Implementation Leakage

```markdown
# Bad
Use PostgreSQL for the database.
Implement with microservices architecture.

# Good
Data SHALL persist across server restarts.
Services SHALL be independently deployable.
```

### 2. Ambiguous Language

Avoid: should, might, could, nice to have, etc.

Use: SHALL (required), SHOULD (recommended), MAY (optional)

### 3. Missing Edge Cases

Always consider:

- Error states
- Empty states
- Boundary conditions
- Concurrent access
- Network failures

### 4. Scope Creep in Specs

Keep specs focused. If scope grows, create new specs rather than expanding existing ones.
