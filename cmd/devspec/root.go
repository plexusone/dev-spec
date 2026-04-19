package main

import (
	"github.com/plexusone/dev-spec/pkg/format"
	"github.com/spf13/cobra"
)

var (
	// Version is set during build
	Version = "dev"

	// Global flags
	outputFormat string
	verbose      bool
)

var rootCmd = &cobra.Command{
	Use:   "devspec",
	Short: "Spec-Driven Development Evaluation CLI",
	Long: `devspec evaluates project specifications against multiple SDD
(Spec-Driven Development) methodologies using LLM-as-a-Judge.

Supported SDD types:
  - kiro       AWS Kiro's three-phase workflow (Requirements → Design → Tasks)
  - speckit    GitHub SpecKit (Spec → Plan → Tasks)
  - plexusone  PlexusOne comprehensive SDD (MRD → PRD → TRD → PLAN → TASKS)

Output formats:
  - toon (default): Token-Oriented Object Notation, ~40% fewer tokens than JSON
  - json: Standard JSON with indentation
  - json-compact: Minified JSON
  - text: Human-readable text
  - markdown: Markdown documentation

Commands:
  check       Detect SDD type in current directory (deterministic)
  validate    Structure validation against SDD type (deterministic)
  init        Scaffold specs for a specific SDD type
  info        Get SDD type info (files, sections, templates)
  rubrics     Get evaluation rubrics for coding assistant evaluation
  evaluate    Full LLM-as-a-Judge evaluation (requires LLM config)
`,
	SilenceUsage: true,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "format", "f", "", "Output format (toon, json, json-compact, text, markdown)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	rootCmd.AddCommand(checkCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(rubricsCmd)
	rootCmd.AddCommand(evaluateCmd)
	rootCmd.AddCommand(versionCmd)
}

// getFormat parses the output format flag.
func getFormat() (format.Format, error) {
	return format.Parse(outputFormat)
}
