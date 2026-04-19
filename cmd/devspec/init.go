package main

import (
	"fmt"

	"github.com/plexusone/dev-spec/pkg/scaffold"
	"github.com/plexusone/dev-spec/pkg/sdd"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <type> [path]",
	Short: "Initialize spec scaffolding for an SDD type",
	Long: `Initialize spec file scaffolding for a specific SDD type.
Creates the required directory structure and template files.

Available types:
  - kiro       AWS Kiro's three-phase workflow
  - speckit    GitHub SpecKit
  - plexusone  PlexusOne comprehensive SDD

Examples:
  devspec init kiro              # Create .kiro/specs/ in current directory
  devspec init kiro ./my-project # Create specs in specified directory
  devspec init speckit           # Create SpecKit templates
  devspec init plexusone         # Create specs/ with PlexusOne templates`,
	Args: cobra.RangeArgs(1, 2),
	RunE: runInit,
}

var initForce bool

func init() {
	initCmd.Flags().BoolVar(&initForce, "force", false, "Overwrite existing files")
}

//nolint:errcheck // CLI output writes are intentionally fire-and-forget
func runInit(cmd *cobra.Command, args []string) error {
	typeName := args[0]
	basePath := "."
	if len(args) > 1 {
		basePath = args[1]
	}

	resolver := sdd.NewResolver()
	sddType, err := resolver.ResolveType(typeName)
	if err != nil {
		return fmt.Errorf("unknown SDD type %q: %w", typeName, err)
	}

	opts := scaffold.Options{
		Force:   initForce,
		Verbose: verbose,
	}

	result, err := scaffold.Init(basePath, sddType, opts)
	if err != nil {
		return err
	}

	// Print results
	for _, f := range result.CreatedFiles {
		fmt.Fprintf(cmd.OutOrStdout(), "Created %s\n", f)
	}
	for _, f := range result.SkippedFiles {
		fmt.Fprintf(cmd.OutOrStdout(), "Skipping %s (exists, use --force to overwrite)\n", f)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "\nInitialized %s spec structure in %s/\n", sddType.DisplayName, result.SpecDirectory)
	fmt.Fprintln(cmd.OutOrStdout(), "\nNext steps:")
	fmt.Fprintln(cmd.OutOrStdout(), "  1. Fill in the template files with your specifications")
	fmt.Fprintln(cmd.OutOrStdout(), "  2. Run 'devspec validate' to check structure")
	fmt.Fprintln(cmd.OutOrStdout(), "  3. Run 'devspec evaluate' for LLM-powered evaluation")

	return nil
}
