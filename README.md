# Dev Spec

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Docs][docs-mkdoc-svg]][docs-mkdoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/plexusone/dev-spec/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/plexusone/dev-spec/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/plexusone/dev-spec/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/plexusone/dev-spec/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/plexusone/dev-spec/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/plexusone/dev-spec/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/plexusone/dev-spec
 [goreport-url]: https://goreportcard.com/report/github.com/plexusone/dev-spec
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/plexusone/dev-spec
 [docs-godoc-url]: https://pkg.go.dev/github.com/plexusone/dev-spec
 [docs-mkdoc-svg]: https://img.shields.io/badge/docs-MkDocs-blue.svg
 [docs-mkdoc-url]: https://plexusone.dev/dev-spec/
 [viz-svg]: https://img.shields.io/badge/visualizaton-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=plexusone%2Fdev-spec
 [loc-svg]: https://tokei.rs/b1/github/plexusone/dev-spec
 [repo-url]: https://github.com/plexusone/dev-spec
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/plexusone/dev-spec/blob/main/LICENSE

**Spec-Driven Development Evaluation CLI**

devspec evaluates project specifications against multiple SDD (Spec-Driven Development) methodologies using LLM-as-a-Judge.

## Features

- **Multiple SDD Types**: Built-in support for Kiro, SpecKit, and PlexusOne methodologies
- **Deterministic Validation**: Structure validation without LLM calls
- **LLM-as-a-Judge**: AI-powered evaluation using customizable rubrics
- **Extensible**: Create custom SDD types or extend built-ins
- **Token-Efficient**: TOON output format saves ~40% tokens vs JSON
- **Multi-Provider LLM**: Supports Anthropic, OpenAI, AWS Bedrock, Google Gemini

## Installation

```bash
go install github.com/plexusone/dev-spec/cmd/devspec@latest
```

Or build from source:

```bash
git clone https://github.com/plexusone/dev-spec.git
cd dev-spec
go build -o devspec ./cmd/devspec
```

## Quick Start

```bash
# Initialize specs for a project using Kiro methodology
devspec init kiro ./my-project

# Check what SDD type is detected in a directory
devspec check ./my-project

# Validate spec structure (deterministic, no LLM)
devspec validate ./my-project

# Get SDD type info for coding assistants
devspec info kiro --format json

# Get evaluation rubrics
devspec rubrics kiro

# Evaluate specs with LLM (requires API key)
export ANTHROPIC_API_KEY=your-key
devspec evaluate ./my-project --llm anthropic
```

## Supported SDD Types

| Type | Description | Required Files |
|------|-------------|----------------|
| **Kiro** | AWS Kiro's three-phase workflow with EARS format | requirements.md, design.md, tasks.md |
| **SpecKit** | GitHub's spec-driven development | spec.md, plan.md, tasks.md |
| **PlexusOne** | Comprehensive enterprise SDD | PRD.md, TRD.md, PLAN.md, TASKS.md |

## Commands

| Command | Description |
|---------|-------------|
| `check` | Detect SDD type in a directory |
| `validate` | Validate spec structure against SDD type |
| `init` | Scaffold spec files for a chosen SDD type |
| `info` | Get detailed SDD type information |
| `rubrics` | Get evaluation rubrics (GO/WARN/NO-GO criteria) |
| `evaluate` | LLM-powered evaluation using LLM-as-a-Judge |
| `version` | Display version information |

## Output Formats

All commands support multiple output formats via `--format`:

| Format | Description |
|--------|-------------|
| `toon` | Token-efficient format (default), ~40% fewer tokens than JSON |
| `json` | Standard JSON with indentation |
| `json-compact` | Minified JSON |
| `text` | Human-readable plain text |
| `markdown` | Formatted markdown documentation |

## Deployment Models

### Coding Assistants

Use `devspec info` and `devspec rubrics` to provide context to coding assistants like Claude Code, Cursor, or Kiro CLI. The assistant can self-evaluate specs using the rubrics.

### CI/CD Pipeline

Run `devspec validate` for deterministic checks and `devspec evaluate` for LLM-powered quality gates in GitHub Actions, GitLab CI, or Jenkins.

```yaml
# GitHub Actions example
- name: Validate specs
  run: devspec validate . --format json

- name: Evaluate specs
  run: devspec evaluate . --llm anthropic --format json
  env:
    ANTHROPIC_API_KEY: ${{ secrets.ANTHROPIC_API_KEY }}
```

### Custom SDD Types

Create custom SDD types by placing definition files in:

- Project-local: `.devspec/types/`
- User-global: `~/.config/devspec/types/`

See the [documentation](https://plexusone.github.io/dev-spec/) for the definition format.

## Documentation

Full documentation is available at [plexusone.github.io/dev-spec](https://plexusone.github.io/dev-spec/).

## License

MIT
