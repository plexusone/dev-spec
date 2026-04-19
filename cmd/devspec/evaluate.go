package main

import (
	"fmt"

	"github.com/plexusone/dev-spec/pkg/detect"
	"github.com/plexusone/dev-spec/pkg/evaluate"
	"github.com/plexusone/dev-spec/pkg/format"
	"github.com/plexusone/dev-spec/pkg/output"
	"github.com/plexusone/dev-spec/pkg/sdd"
	"github.com/spf13/cobra"
)

var evaluateCmd = &cobra.Command{
	Use:   "evaluate [path]",
	Short: "Evaluate specs using LLM-as-a-Judge",
	Long: `Evaluate spec files against the detected or specified SDD type
using LLM-as-a-Judge methodology.

This command requires LLM configuration (via environment or flags).

Examples:
  devspec evaluate                              # Evaluate current directory (TOON output)
  devspec evaluate --type kiro                  # Force Kiro evaluation
  devspec evaluate --llm anthropic              # Use Anthropic API
  devspec evaluate --llm openai --model gpt-4   # Use OpenAI GPT-4
  devspec evaluate --tools                      # Enable tool-based evaluation (LLM can read files)
  devspec evaluate --format json                # Output as JSON
  devspec evaluate --format markdown            # Output as Markdown report`,
	Args: cobra.MaximumNArgs(1),
	RunE: runEvaluate,
}

var (
	evaluateType       string
	evaluateLLM        string
	evaluateModel      string
	evaluateDryRun     bool
	evaluateDefinition string
	evaluateTools      bool
)

func init() {
	evaluateCmd.Flags().StringVarP(&evaluateType, "type", "t", "", "SDD type to evaluate against (auto-detected if not specified)")
	evaluateCmd.Flags().StringVar(&evaluateLLM, "llm", "", "LLM provider (anthropic, openai, etc.)")
	evaluateCmd.Flags().StringVar(&evaluateModel, "model", "", "LLM model to use")
	evaluateCmd.Flags().BoolVar(&evaluateDryRun, "dry-run", false, "Show prompts without calling LLM")
	evaluateCmd.Flags().StringVar(&evaluateDefinition, "definition", "", "Path to custom SDD definition")
	evaluateCmd.Flags().BoolVar(&evaluateTools, "tools", false, "Enable tool-based evaluation (LLM can call tools to read files, validate, etc.)")
}

//nolint:errcheck // CLI output writes are intentionally fire-and-forget
func runEvaluate(cmd *cobra.Command, args []string) error {
	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	// Load SDD type
	var sddType *sdd.SDDType
	var err error

	resolver := sdd.NewResolver()

	if evaluateDefinition != "" {
		// Load from custom path
		sddType, err = resolver.LoadFromPath(evaluateDefinition)
		if err != nil {
			return fmt.Errorf("load custom definition: %w", err)
		}
	} else if evaluateType != "" {
		// Use specified type
		sddType, err = resolver.ResolveType(evaluateType)
		if err != nil {
			return fmt.Errorf("resolve SDD type %q: %w", evaluateType, err)
		}
	} else {
		// Auto-detect
		result, err := detect.DetectSDDType(path)
		if err != nil {
			return fmt.Errorf("detection failed: %w", err)
		}
		if !result.Detected {
			return fmt.Errorf("no SDD type detected; use --type to specify")
		}
		sddType = result.SDDType
	}

	if verbose {
		fmt.Fprintf(cmd.OutOrStdout(), "Evaluating against: %s\n", sddType.DisplayName)
	}

	// Create evaluator
	opts := evaluate.Options{
		LLMProvider: evaluateLLM,
		LLMModel:    evaluateModel,
		DryRun:      evaluateDryRun,
		Verbose:     verbose,
		UseTools:    evaluateTools,
	}

	evaluator, err := evaluate.NewEvaluator(opts)
	if err != nil {
		return fmt.Errorf("create evaluator: %w", err)
	}
	defer evaluator.Close() //nolint:errcheck // Best effort close

	// Run evaluation
	result, err := evaluator.Evaluate(path, sddType)
	if err != nil {
		return fmt.Errorf("evaluation failed: %w", err)
	}

	// Get format and marshal
	f, err := getFormat()
	if err != nil {
		return err
	}

	evalOutput := output.EvaluationOutput{EvaluationResult: result}
	data, err := format.Marshal(evalOutput, f)
	if err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}
