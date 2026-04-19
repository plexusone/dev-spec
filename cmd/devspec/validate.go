package main

import (
	"fmt"

	"github.com/plexusone/dev-spec/pkg/detect"
	"github.com/plexusone/dev-spec/pkg/format"
	"github.com/plexusone/dev-spec/pkg/output"
	"github.com/plexusone/dev-spec/pkg/validate"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate [path]",
	Short: "Validate spec structure against SDD type",
	Long: `Validate the structure of spec files against the detected or specified
SDD type. This checks for required sections, file presence, and format compliance.

This is a deterministic command that does not use LLM calls.

Examples:
  devspec validate                    # Validate current directory (TOON output)
  devspec validate --type kiro        # Force Kiro validation
  devspec validate ./my-project       # Validate specific directory
  devspec validate --format json      # Output as JSON`,
	Args: cobra.MaximumNArgs(1),
	RunE: runValidate,
}

var validateType string

func init() {
	validateCmd.Flags().StringVarP(&validateType, "type", "t", "", "SDD type to validate against (auto-detected if not specified)")
}

//nolint:errcheck // CLI output writes are intentionally fire-and-forget
func runValidate(cmd *cobra.Command, args []string) error {
	path := "."
	if len(args) > 0 {
		path = args[0]
	}

	// Detect or use specified type
	var sddTypeName string
	if validateType != "" {
		sddTypeName = validateType
	} else {
		result, err := detect.DetectSDDType(path)
		if err != nil {
			return fmt.Errorf("detection failed: %w", err)
		}
		if !result.Detected {
			return fmt.Errorf("no SDD type detected; use --type to specify")
		}
		sddTypeName = result.SDDType.Name
	}

	// Run validation
	validationResult, err := validate.ValidateStructure(path, sddTypeName)
	if err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Get format and marshal
	f, err := getFormat()
	if err != nil {
		return err
	}

	out := output.ValidateOutput{ValidationResult: validationResult}
	data, err := format.Marshal(out, f)
	if err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}
