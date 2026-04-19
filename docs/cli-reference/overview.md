# CLI Reference

## Global Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--format` | `-f` | Output format (toon, json, json-compact, text, markdown) |
| `--verbose` | `-v` | Enable verbose output |
| `--help` | `-h` | Show help |

## Output Formats

| Format | Description | Use Case |
|--------|-------------|----------|
| `toon` | Token-Oriented Object Notation | Default; LLM context (40% smaller than JSON) |
| `json` | Indented JSON | Programmatic processing |
| `json-compact` | Minified JSON | API responses, storage |
| `text` | Human-readable text | Terminal output |
| `markdown` | Markdown tables | Documentation, reports |

## Commands

### Deterministic (No LLM)

| Command | Description |
|---------|-------------|
| `check` | Detect SDD type in a directory |
| `validate` | Validate spec structure |
| `init` | Initialize spec scaffolding |
| `info` | Get SDD type information |
| `rubrics` | Get evaluation rubrics |

### LLM-Powered

| Command | Description |
|---------|-------------|
| `evaluate` | Full LLM-as-a-Judge evaluation |

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Invalid arguments |
| 3 | SDD type not found |
| 4 | Validation failed |
