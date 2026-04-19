package main

import (
	"fmt"

	"github.com/plexusone/dev-spec/pkg/format"
	"github.com/plexusone/dev-spec/pkg/output"
	"github.com/plexusone/dev-spec/pkg/sdd"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <type>",
	Short: "Get SDD type information for coding assistants",
	Long: `Returns structured information about an SDD type including:
- Required and optional files with patterns
- Section structure for each file
- Template content for scaffolding

This is designed for coding assistants (Claude Code, Cursor, etc.) to guide
users in creating spec files.

Examples:
  devspec info kiro                   # Get Kiro SDD info (TOON output)
  devspec info speckit --format json  # JSON output for programmatic use
  devspec info plexusone --format text  # Human-readable output`,
	Args: cobra.ExactArgs(1),
	RunE: runInfo,
}

//nolint:errcheck // CLI output writes are intentionally fire-and-forget
func runInfo(cmd *cobra.Command, args []string) error {
	typeName := args[0]

	resolver := sdd.NewResolver()
	sddType, err := resolver.ResolveType(typeName)
	if err != nil {
		return fmt.Errorf("unknown SDD type %q: %w", typeName, err)
	}

	info := output.NewSDDInfo(sddType)

	// Get format and marshal
	f, err := getFormat()
	if err != nil {
		return err
	}

	data, err := format.Marshal(info, f)
	if err != nil {
		return err
	}

	fmt.Fprintln(cmd.OutOrStdout(), string(data))
	return nil
}
