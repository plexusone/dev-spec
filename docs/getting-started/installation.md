# Installation

## From Source

```bash
# Clone the repository
git clone https://github.com/plexusone/dev-spec.git
cd dev-spec

# Install
go install ./cmd/devspec

# Verify
devspec --version
```

## Prerequisites

- Go 1.21 or later
- For LLM evaluation: API key for your chosen provider (Anthropic, OpenAI, etc.)

## Shell Completion

Generate shell completions for your shell:

```bash
# Bash
devspec completion bash > /etc/bash_completion.d/devspec

# Zsh
devspec completion zsh > "${fpath[1]}/_devspec"

# Fish
devspec completion fish > ~/.config/fish/completions/devspec.fish
```
