# CI/CD Pipeline Integration

Automate spec validation and evaluation in your CI/CD pipeline.

## Overview

In this model:

- **devspec** makes LLM calls directly
- **Pipeline** runs devspec as a quality gate
- **Results** determine pass/fail

## GitHub Actions

### Basic Validation (No LLM)

```yaml
name: Spec Validation

on:
  pull_request:
    paths:
      - 'specs/**'
      - '.kiro/**'

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Install devspec
        run: go install github.com/plexusone/dev-spec/cmd/devspec@latest
      
      - name: Check SDD type
        run: devspec check
      
      - name: Validate structure
        run: devspec validate --format json
```

### Full Evaluation (With LLM)

```yaml
name: Spec Evaluation

on:
  pull_request:
    paths:
      - 'specs/**'

jobs:
  evaluate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Install devspec
        run: go install github.com/plexusone/dev-spec/cmd/devspec@latest
      
      - name: Evaluate specs
        env:
          ANTHROPIC_API_KEY: ${{ secrets.ANTHROPIC_API_KEY }}
        run: |
          devspec evaluate --llm anthropic --format json > evaluation.json
          
      - name: Check score threshold
        run: |
          SCORE=$(jq '.summary.total_score' evaluation.json)
          if (( $(echo "$SCORE < 0.7" | bc -l) )); then
            echo "Spec score $SCORE is below threshold 0.7"
            exit 1
          fi
          
      - name: Upload evaluation report
        uses: actions/upload-artifact@v4
        with:
          name: spec-evaluation
          path: evaluation.json
```

### PR Comment with Results

```yaml
      - name: Comment on PR
        uses: actions/github-script@v7
        with:
          script: |
            const fs = require('fs');
            const eval = JSON.parse(fs.readFileSync('evaluation.json'));
            
            let body = `## Spec Evaluation Results\n\n`;
            body += `**Score**: ${(eval.summary.total_score * 100).toFixed(1)}%\n`;
            body += `**Status**: ${eval.status}\n\n`;
            
            for (const file of eval.files) {
              body += `### ${file.display_name} - ${file.status}\n`;
              for (const c of file.criteria) {
                body += `- ${c.status} ${c.title}\n`;
              }
            }
            
            github.rest.issues.createComment({
              issue_number: context.issue.number,
              owner: context.repo.owner,
              repo: context.repo.repo,
              body: body
            });
```

## GitLab CI

```yaml
spec-validation:
  stage: test
  script:
    - go install github.com/plexusone/dev-spec/cmd/devspec@latest
    - devspec validate
  rules:
    - changes:
        - specs/**

spec-evaluation:
  stage: test
  script:
    - go install github.com/plexusone/dev-spec/cmd/devspec@latest
    - devspec evaluate --llm anthropic --format json > evaluation.json
    - |
      SCORE=$(jq '.summary.total_score' evaluation.json)
      if [ $(echo "$SCORE < 0.7" | bc) -eq 1 ]; then
        exit 1
      fi
  variables:
    ANTHROPIC_API_KEY: $ANTHROPIC_API_KEY
  rules:
    - changes:
        - specs/**
```

## Jenkins

```groovy
pipeline {
    agent any
    
    environment {
        ANTHROPIC_API_KEY = credentials('anthropic-api-key')
    }
    
    stages {
        stage('Validate Specs') {
            steps {
                sh 'devspec validate --format json > validation.json'
            }
        }
        
        stage('Evaluate Specs') {
            steps {
                sh 'devspec evaluate --llm anthropic --format json > evaluation.json'
                script {
                    def eval = readJSON file: 'evaluation.json'
                    if (eval.summary.total_score < 0.7) {
                        error "Spec score ${eval.summary.total_score} below threshold"
                    }
                }
            }
        }
    }
    
    post {
        always {
            archiveArtifacts artifacts: '*.json'
        }
    }
}
```

## Quality Gates

### Score Thresholds

| Level | Threshold | Action |
|-------|-----------|--------|
| Blocking | < 50% | Fail pipeline |
| Warning | 50-70% | Require approval |
| Passing | > 70% | Auto-merge eligible |

### Status-Based Gates

```bash
# Fail if any NO-GO
if devspec evaluate --format json | jq -e '.files[].criteria[] | select(.status == "NO-GO")' > /dev/null; then
  echo "Found NO-GO criteria"
  exit 1
fi
```

## Cost Management

LLM calls in CI can add up. Consider:

1. **Only evaluate on spec changes** - Use path filters
2. **Cache evaluations** - Skip if spec unchanged
3. **Use smaller models** - GPT-3.5 for validation, GPT-4 for final
4. **Rate limit** - Don't evaluate every commit
