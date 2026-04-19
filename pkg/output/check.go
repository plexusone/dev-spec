package output

import (
	"fmt"
	"strings"
)

// CheckOutput is the structured output for the check command.
type CheckOutput struct {
	Detected      bool    `json:"detected"`
	SDDType       string  `json:"sdd_type,omitempty"`
	DisplayName   string  `json:"display_name,omitempty"`
	Description   string  `json:"description,omitempty"`
	SpecDirectory string  `json:"spec_directory,omitempty"`
	MatchedFile   string  `json:"matched_file,omitempty"`
	Confidence    float64 `json:"confidence,omitempty"`
}

// FormatText implements format.TextFormatter.
func (c CheckOutput) FormatText() ([]byte, error) {
	var sb strings.Builder
	if !c.Detected {
		sb.WriteString("No SDD type detected\n")
		sb.WriteString("\nHint: Initialize specs with 'devspec init <type>'\n")
		sb.WriteString("Available types: kiro, speckit, plexusone\n")
		return []byte(sb.String()), nil
	}

	fmt.Fprintf(&sb, "Detected: %s\n", c.SDDType)
	fmt.Fprintf(&sb, "  Display Name: %s\n", c.DisplayName)
	fmt.Fprintf(&sb, "  Description: %s\n", c.Description)
	fmt.Fprintf(&sb, "  Spec Directory: %s\n", c.SpecDirectory)
	fmt.Fprintf(&sb, "  Matched File: %s\n", c.MatchedFile)
	fmt.Fprintf(&sb, "  Confidence: %.0f%%\n", c.Confidence*100)
	return []byte(sb.String()), nil
}

// FormatMarkdown implements format.MarkdownFormatter.
func (c CheckOutput) FormatMarkdown() ([]byte, error) {
	var sb strings.Builder
	if !c.Detected {
		sb.WriteString("# SDD Type Detection\n\n")
		sb.WriteString("**No SDD type detected**\n\n")
		sb.WriteString("Initialize specs with `devspec init <type>`\n\n")
		sb.WriteString("Available types: `kiro`, `speckit`, `plexusone`\n")
		return []byte(sb.String()), nil
	}

	sb.WriteString("# SDD Type Detection\n\n")
	fmt.Fprintf(&sb, "**Detected**: %s\n\n", c.SDDType)
	sb.WriteString("| Property | Value |\n")
	sb.WriteString("|----------|-------|\n")
	fmt.Fprintf(&sb, "| Display Name | %s |\n", c.DisplayName)
	fmt.Fprintf(&sb, "| Description | %s |\n", c.Description)
	fmt.Fprintf(&sb, "| Spec Directory | `%s` |\n", c.SpecDirectory)
	fmt.Fprintf(&sb, "| Matched File | `%s` |\n", c.MatchedFile)
	fmt.Fprintf(&sb, "| Confidence | %.0f%% |\n", c.Confidence*100)
	return []byte(sb.String()), nil
}
