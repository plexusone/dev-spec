package main

import (
	"fmt"

	"github.com/plexusone/dev-spec/pkg/detect"
	"github.com/plexusone/dev-spec/pkg/format"
	"github.com/plexusone/dev-spec/pkg/output"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check [path]",
	Short: "Detect SDD type in a directory",
	Long: `Detect which SDD (Spec-Driven Development) type is present in a directory
based on file patterns and structure.

This is a deterministic command that does not use LLM calls.

Examples:
  devspec check                       # Check current directory (TOON output)
  devspec check ./my-project          # Check specific directory
  devspec check --format json         # Output as JSON
  devspec check --format text         # Human-readable output`,
	Args: cobra.MaximumNArgs(1),
	RunE: runCheck,
}

//nolint:errcheck // CLI output writes are intentionally fire-and-forget
func runCheck(cmd *cobra.Command, args []string) error {
	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	result, err := detect.DetectSDDType(path)
	if err != nil {
		return fmt.Errorf("detection failed: %w", err)
	}

	// Build output
	out := output.CheckOutput{
		Detected: result.Detected,
	}
	if result.Detected && result.SDDType != nil {
		out.SDDType = result.SDDType.Name
		out.DisplayName = result.SDDType.DisplayName
		out.Description = result.SDDType.Description
		out.SpecDirectory = result.SDDType.SpecDirectory
		out.MatchedFile = result.MatchedFile
		out.Confidence = result.Confidence
	}

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
