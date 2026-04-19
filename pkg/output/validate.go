package output

import (
	"fmt"
	"strings"

	"github.com/plexusone/dev-spec/pkg/sdd"
)

// ValidateOutput wraps ValidationResult with text/markdown marshaling.
type ValidateOutput struct {
	*sdd.ValidationResult
}

// FormatText implements format.TextFormatter.
func (v ValidateOutput) FormatText() ([]byte, error) {
	var sb strings.Builder
	r := v.ValidationResult

	fmt.Fprintf(&sb, "SDD Type: %s\n", r.SDDType)
	fmt.Fprintf(&sb, "Valid: %v\n\n", r.Valid)

	if len(r.PresentSections) > 0 {
		sb.WriteString("Present Sections:\n")
		for _, s := range r.PresentSections {
			status := "optional"
			if s.Required {
				status = "required"
			}
			fmt.Fprintf(&sb, "  [+] %s/%s (%s)\n", s.File, s.Section, status)
		}
	}

	if len(r.MissingSections) > 0 {
		sb.WriteString("\nMissing Sections:\n")
		for _, s := range r.MissingSections {
			status := "optional"
			if s.Required {
				status = "required"
			}
			fmt.Fprintf(&sb, "  [-] %s/%s (%s)\n", s.File, s.Section, status)
		}
	}

	if len(r.Errors) > 0 {
		sb.WriteString("\nErrors:\n")
		for _, e := range r.Errors {
			fmt.Fprintf(&sb, "  - %s\n", e)
		}
	}

	if len(r.Warnings) > 0 {
		sb.WriteString("\nWarnings:\n")
		for _, w := range r.Warnings {
			fmt.Fprintf(&sb, "  - %s\n", w)
		}
	}

	return []byte(sb.String()), nil
}

// FormatMarkdown implements format.MarkdownFormatter.
func (v ValidateOutput) FormatMarkdown() ([]byte, error) {
	var sb strings.Builder
	r := v.ValidationResult

	sb.WriteString("# Validation Report\n\n")
	fmt.Fprintf(&sb, "**SDD Type**: %s\n\n", r.SDDType)
	fmt.Fprintf(&sb, "**Valid**: %v\n\n", r.Valid)

	if len(r.PresentSections) > 0 {
		sb.WriteString("## Present Sections\n\n")
		sb.WriteString("| File | Section | Status |\n")
		sb.WriteString("|------|---------|--------|\n")
		for _, s := range r.PresentSections {
			status := "optional"
			if s.Required {
				status = "required"
			}
			fmt.Fprintf(&sb, "| %s | %s | %s |\n", s.File, s.Section, status)
		}
		sb.WriteString("\n")
	}

	if len(r.MissingSections) > 0 {
		sb.WriteString("## Missing Sections\n\n")
		sb.WriteString("| File | Section | Status |\n")
		sb.WriteString("|------|---------|--------|\n")
		for _, s := range r.MissingSections {
			status := "optional"
			if s.Required {
				status = "required"
			}
			fmt.Fprintf(&sb, "| %s | %s | %s |\n", s.File, s.Section, status)
		}
		sb.WriteString("\n")
	}

	if len(r.Errors) > 0 {
		sb.WriteString("## Errors\n\n")
		for _, e := range r.Errors {
			fmt.Fprintf(&sb, "- %s\n", e)
		}
		sb.WriteString("\n")
	}

	if len(r.Warnings) > 0 {
		sb.WriteString("## Warnings\n\n")
		for _, w := range r.Warnings {
			fmt.Fprintf(&sb, "- %s\n", w)
		}
	}

	return []byte(sb.String()), nil
}
