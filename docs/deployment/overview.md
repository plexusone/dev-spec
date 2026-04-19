# Deployment Overview

devspec supports three deployment models for different use cases.

## Deployment Models

| Model | LLM Calls | Use Case |
|-------|-----------|----------|
| **Coding Assistants** | By assistant | Interactive development |
| **CI/CD Pipeline** | By devspec | Automated quality gates |
| **Service** | By service | API/MCP integration |

## Command Categories

### Deterministic (No LLM)

These commands work in all deployment models without LLM configuration:

- `check` - Detect SDD type
- `validate` - Structure validation
- `init` - Scaffold templates
- `info` - Get SDD metadata
- `rubrics` - Get evaluation rubrics

### LLM-Powered

These require LLM configuration:

- `evaluate` - Full quality evaluation

## Choosing a Model

```
┌─────────────────────────────────────────────────────────────┐
│                    Is this interactive?                      │
└─────────────────────────────────────────────────────────────┘
                            │
              ┌─────────────┴─────────────┐
              ▼                           ▼
            Yes                          No
              │                           │
              ▼                           ▼
    ┌─────────────────┐         ┌─────────────────┐
    │ Coding Assistant │         │   Automated?    │
    │ (info, rubrics)  │         └─────────────────┘
    └─────────────────┘                   │
                              ┌───────────┴───────────┐
                              ▼                       ▼
                            Yes                      No
                              │                       │
                              ▼                       ▼
                    ┌─────────────────┐     ┌─────────────────┐
                    │    CI/CD        │     │    Service      │
                    │   (evaluate)    │     │   (MCP/HTTP)    │
                    └─────────────────┘     └─────────────────┘
```

## Feature Comparison

| Feature | Coding Assistant | CI/CD | Service |
|---------|------------------|-------|---------|
| Interactive feedback | ✅ | ❌ | ✅ |
| Automated gates | ❌ | ✅ | ✅ |
| LLM cost control | User's quota | Your cost | Your cost |
| Codebase context | Full | Limited | API-based |
| Custom definitions | Local | Repo | Server-wide |
