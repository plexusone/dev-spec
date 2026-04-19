package main

import (
	"fmt"

	"github.com/plexusone/dev-spec/pkg/format"
	"github.com/plexusone/dev-spec/pkg/output"
	"github.com/plexusone/dev-spec/pkg/sdd"
	"github.com/spf13/cobra"
)

var rubricsCmd = &cobra.Command{
	Use:   "rubrics <type>",
	Short: "Get evaluation rubrics for coding assistants",
	Long: `Returns evaluation rubrics and criteria for an SDD type.

This is designed for coding assistants (Claude Code, Cursor, etc.) to evaluate
spec files themselves using the rubrics, rather than having the CLI make LLM calls.

The output includes:
- Criteria IDs and weights for each file type
- GO/WARN/NO-GO descriptions for each criterion
- Scoring guidance

Examples:
  devspec rubrics kiro                        # Get Kiro rubrics (TOON output)
  devspec rubrics kiro --format json          # JSON output
  devspec rubrics kiro --format markdown      # Markdown documentation
  devspec rubrics kiro --file requirements    # Rubrics for specific file only`,
	Args: cobra.ExactArgs(1),
	RunE: runRubrics,
}

var rubricsFile string

func init() {
	rubricsCmd.Flags().StringVar(&rubricsFile, "file", "", "Get rubrics for specific file only")
}

//nolint:errcheck // CLI output writes are intentionally fire-and-forget
func runRubrics(cmd *cobra.Command, args []string) error {
	typeName := args[0]

	resolver := sdd.NewResolver()
	sddType, err := resolver.ResolveType(typeName)
	if err != nil {
		return fmt.Errorf("unknown SDD type %q: %w", typeName, err)
	}

	out := output.NewRubricsOutput(sddType, rubricsFile)

	// Get format and marshal
	f, err := getFormat()
	if err != nil {
		return err
	}

	data, err := format.Marshal(out, f)
	if err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}
