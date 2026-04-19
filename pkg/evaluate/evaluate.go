// Package evaluate provides LLM-as-a-Judge evaluation for spec files.
package evaluate

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/plexusone/dev-spec/pkg/llm"
	"github.com/plexusone/dev-spec/pkg/sdd"
	"github.com/plexusone/dev-spec/pkg/tools"
)

// Options configures the evaluator.
type Options struct {
	LLMProvider string
	LLMModel    string
	DryRun      bool
	Verbose     bool
	UseTools    bool // Enable tool-based evaluation
}

// Evaluator performs LLM-as-a-Judge evaluation.
type Evaluator struct {
	opts   Options
	client *llm.Client
}

// NewEvaluator creates a new evaluator with the given options.
func NewEvaluator(opts Options) (*Evaluator, error) {
	var client *llm.Client
	var err error

	if !opts.DryRun {
		client, err = llm.NewClient(llm.Config{
			Provider: opts.LLMProvider,
			Model:    opts.LLMModel,
		})
		if err != nil {
			return nil, fmt.Errorf("create LLM client: %w", err)
		}
	}

	return &Evaluator{
		opts:   opts,
		client: client,
	}, nil
}

// Close releases resources.
func (e *Evaluator) Close() error {
	if e.client != nil {
		return e.client.Close()
	}
	return nil
}

// Evaluate runs LLM-as-a-Judge evaluation on spec files.
func (e *Evaluator) Evaluate(path string, sddType *sdd.SDDType) (*sdd.EvaluationResult, error) {
	return e.EvaluateWithContext(context.Background(), path, sddType)
}

// EvaluateWithContext runs evaluation with context support.
func (e *Evaluator) EvaluateWithContext(ctx context.Context, path string, sddType *sdd.SDDType) (*sdd.EvaluationResult, error) {
	if e.opts.UseTools {
		return e.evaluateWithTools(ctx, path, sddType)
	}
	return e.evaluateStandard(ctx, path, sddType)
}

// evaluateWithTools uses the tool-enabled LLM to evaluate specs.
// The LLM can call tools to read files, validate structure, and get rubrics.
func (e *Evaluator) evaluateWithTools(ctx context.Context, path string, sddType *sdd.SDDType) (*sdd.EvaluationResult, error) {
	if e.opts.DryRun {
		return e.dryRunResult(sddType), nil
	}

	// Register tools with the client
	tools.RegisterAllTools(e.client, path)

	// Build system prompt for evaluation
	systemPrompt := buildEvaluationSystemPrompt(sddType)

	// Build user prompt
	userPrompt := buildEvaluationUserPrompt(path, sddType)

	// Run evaluation with tools
	response, err := e.client.CompleteWithTools(ctx, systemPrompt, userPrompt)
	if err != nil {
		return nil, fmt.Errorf("LLM evaluation: %w", err)
	}

	// Parse the evaluation response
	return ParseEvaluationResponse(sddType, response)
}

// evaluateStandard performs standard prompt-based evaluation without tools.
func (e *Evaluator) evaluateStandard(ctx context.Context, path string, sddType *sdd.SDDType) (*sdd.EvaluationResult, error) {
	result := &sdd.EvaluationResult{
		SDDType: sddType.Name,
		Status:  sdd.StatusGo,
	}

	// Find spec directory
	specDir := filepath.Join(path, sddType.SpecDirectory)
	if _, err := os.Stat(specDir); os.IsNotExist(err) {
		specDir = path
	}

	// Evaluate each file
	for _, fileSpec := range sddType.Files {
		// Find matching file
		var matchedFile string
		for _, pattern := range fileSpec.Patterns {
			matches, err := filepath.Glob(filepath.Join(specDir, pattern))
			if err != nil {
				continue
			}
			if len(matches) == 0 {
				matches, _ = filepath.Glob(filepath.Join(path, pattern))
			}
			if len(matches) > 0 {
				matchedFile = matches[0]
				break
			}
		}

		if matchedFile == "" {
			if fileSpec.Required {
				result.Files = append(result.Files, sdd.FileEvaluation{
					File:        fileSpec.Name,
					DisplayName: fileSpec.Name,
					Status:      sdd.StatusNoGo,
					Criteria: []sdd.CriterionResult{{
						ID:        "file_present",
						Title:     "File Present",
						Status:    sdd.StatusNoGo,
						Reasoning: "Required file not found",
					}},
				})
				result.Summary.CriteriaFailed++
			} else {
				result.Files = append(result.Files, sdd.FileEvaluation{
					File:        fileSpec.Name,
					DisplayName: fileSpec.Name,
					Status:      sdd.StatusSkip,
					Criteria: []sdd.CriterionResult{{
						ID:        "file_present",
						Title:     "File Present",
						Status:    sdd.StatusSkip,
						Reasoning: "Optional file not found",
					}},
				})
				result.Summary.CriteriaSkipped++
			}
			continue
		}

		// Get file definition
		fileDef, ok := sddType.FileDefinitions[fileSpec.Name]
		if !ok {
			continue
		}

		// Read file content
		content, err := os.ReadFile(matchedFile)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", matchedFile, err)
		}

		// Evaluate file
		fileEval, err := e.evaluateFile(ctx, fileSpec.Name, fileDef, string(content))
		if err != nil {
			return nil, fmt.Errorf("evaluate %s: %w", matchedFile, err)
		}

		result.Files = append(result.Files, *fileEval)

		// Update summary
		for _, cr := range fileEval.Criteria {
			switch cr.Status {
			case sdd.StatusGo:
				result.Summary.CriteriaPassed++
			case sdd.StatusWarn:
				result.Summary.CriteriaWarned++
			case sdd.StatusNoGo:
				result.Summary.CriteriaFailed++
			case sdd.StatusSkip:
				result.Summary.CriteriaSkipped++
			}
		}
	}

	// Calculate summary
	result.Summary.FilesEvaluated = len(result.Files)
	result.Summary.TotalScore = calculateTotalScore(result)
	result.Summary.Status = calculateOverallStatus(result)
	result.Status = result.Summary.Status

	return result, nil
}

// evaluateFile evaluates a single spec file against its criteria.
func (e *Evaluator) evaluateFile(ctx context.Context, name string, fileDef *sdd.FileDefinition, content string) (*sdd.FileEvaluation, error) {
	eval := &sdd.FileEvaluation{
		File:        name,
		DisplayName: fileDef.DisplayName,
		Status:      sdd.StatusGo,
	}

	for _, criterion := range fileDef.Criteria {
		rubric, ok := fileDef.Rubrics[criterion.ID]
		if !ok {
			continue
		}

		var cr sdd.CriterionResult
		var err error

		if e.opts.DryRun {
			cr = e.dryRunCriterion(criterion, rubric, content)
		} else {
			cr, err = e.evaluateCriterion(ctx, criterion, rubric, content)
			if err != nil {
				return nil, fmt.Errorf("evaluate criterion %s: %w", criterion.ID, err)
			}
		}

		eval.Criteria = append(eval.Criteria, cr)
	}

	// Calculate file score and status
	eval.Score = calculateFileScore(eval)
	eval.Status = calculateFileStatus(eval)

	return eval, nil
}

// evaluateCriterion evaluates a single criterion using LLM.
func (e *Evaluator) evaluateCriterion(ctx context.Context, criterion sdd.CriterionSpec, rubric *sdd.Rubric, content string) (sdd.CriterionResult, error) {
	prompt := BuildPrompt(rubric, content)

	response, err := e.client.CompleteWithContext(ctx, prompt)
	if err != nil {
		return sdd.CriterionResult{}, err
	}

	return ParseResponse(criterion, rubric, response)
}

// dryRunCriterion returns a placeholder result for dry-run mode.
func (e *Evaluator) dryRunCriterion(criterion sdd.CriterionSpec, rubric *sdd.Rubric, content string) sdd.CriterionResult {
	prompt := BuildPrompt(rubric, content)

	if e.opts.Verbose {
		fmt.Printf("\n=== Criterion: %s ===\n", criterion.ID)
		fmt.Printf("Prompt:\n%s\n", prompt)
		fmt.Println("---")
	}

	return sdd.CriterionResult{
		ID:        criterion.ID,
		Title:     rubric.Title,
		Weight:    criterion.Weight,
		Status:    sdd.StatusSkip,
		Reasoning: "[Dry run - no LLM call made]",
	}
}

// dryRunResult returns a placeholder result for dry-run mode with tools.
func (e *Evaluator) dryRunResult(sddType *sdd.SDDType) *sdd.EvaluationResult {
	result := &sdd.EvaluationResult{
		SDDType: sddType.Name,
		Status:  sdd.StatusSkip,
	}

	for _, fileSpec := range sddType.Files {
		result.Files = append(result.Files, sdd.FileEvaluation{
			File:        fileSpec.Name,
			DisplayName: fileSpec.Name,
			Status:      sdd.StatusSkip,
			Criteria: []sdd.CriterionResult{{
				ID:        "dry_run",
				Title:     "Dry Run",
				Status:    sdd.StatusSkip,
				Reasoning: "[Dry run with tools - no LLM call made]",
			}},
		})
	}

	return result
}

func calculateFileScore(eval *sdd.FileEvaluation) float64 {
	if len(eval.Criteria) == 0 {
		return 0
	}

	var totalWeight float64
	var weightedScore float64

	for _, cr := range eval.Criteria {
		totalWeight += cr.Weight
		weightedScore += cr.Score * cr.Weight
	}

	if totalWeight == 0 {
		return 0
	}

	return weightedScore / totalWeight
}

func calculateFileStatus(eval *sdd.FileEvaluation) sdd.Status {
	hasNoGo := false
	hasWarn := false

	for _, cr := range eval.Criteria {
		switch cr.Status {
		case sdd.StatusNoGo:
			hasNoGo = true
		case sdd.StatusWarn:
			hasWarn = true
		}
	}

	if hasNoGo {
		return sdd.StatusNoGo
	}
	if hasWarn {
		return sdd.StatusWarn
	}
	return sdd.StatusGo
}

func calculateTotalScore(result *sdd.EvaluationResult) float64 {
	if len(result.Files) == 0 {
		return 0
	}

	var total float64
	for _, f := range result.Files {
		total += f.Score
	}

	return total / float64(len(result.Files))
}

func calculateOverallStatus(result *sdd.EvaluationResult) sdd.Status {
	hasNoGo := false
	hasWarn := false

	for _, f := range result.Files {
		switch f.Status {
		case sdd.StatusNoGo:
			hasNoGo = true
		case sdd.StatusWarn:
			hasWarn = true
		}
	}

	if hasNoGo {
		return sdd.StatusNoGo
	}
	if hasWarn {
		return sdd.StatusWarn
	}
	return sdd.StatusGo
}

// buildEvaluationSystemPrompt creates the system prompt for tool-based evaluation.
func buildEvaluationSystemPrompt(sddType *sdd.SDDType) string {
	return fmt.Sprintf(`You are an expert evaluator for %s spec-driven development methodology.

Your task is to evaluate spec files against the %s criteria using the tools available to you.

## Available Tools

- devspec_check: Detect SDD type in a directory
- devspec_validate: Validate spec file structure
- devspec_info: Get detailed SDD type information
- devspec_rubrics: Get evaluation rubrics (GO/WARN/NO-GO criteria)
- read_file: Read file contents
- glob_files: Find files matching a pattern

## Evaluation Process

1. First, use devspec_rubrics to get the evaluation criteria for each file type
2. Use glob_files or read_file to read the actual spec files
3. For files marked as requiring codebase access, also examine relevant source code
4. Evaluate each criterion using the rubrics as guidance
5. Provide your final evaluation in the specified JSON format

## Output Format

After completing your evaluation, respond with a JSON object:

{
  "files": [
    {
      "file": "requirements",
      "display_name": "Requirements Document",
      "criteria": [
        {
          "criterion_id": "ears_format",
          "status": "GO" | "WARN" | "NO-GO",
          "score": 0.0 to 1.0,
          "reasoning": "Brief explanation",
          "suggestions": ["Optional improvements"]
        }
      ]
    }
  ],
  "overall_status": "GO" | "WARN" | "NO-GO",
  "summary": "Brief overall assessment"
}

Be thorough and objective. Reference specific examples from the documents.`,
		sddType.DisplayName, sddType.DisplayName)
}

// buildEvaluationUserPrompt creates the user prompt for evaluation.
func buildEvaluationUserPrompt(path string, sddType *sdd.SDDType) string {
	return fmt.Sprintf(`Please evaluate the %s specs in the directory: %s

Start by using the available tools to:
1. Get the rubrics for %s
2. Read the spec files
3. Evaluate each file against its criteria

If any files are marked as requiring codebase access, also examine the relevant source code to verify alignment.

Provide your complete evaluation in the JSON format specified.`,
		sddType.DisplayName, path, sddType.Name)
}
