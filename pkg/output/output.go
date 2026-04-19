// Package output provides output formatting for evaluation results.
package output

import (
	"fmt"
	"strings"

	"github.com/plexusone/dev-spec/pkg/sdd"
)

// EvaluationOutput wraps EvaluationResult with text/markdown marshaling.
type EvaluationOutput struct {
	*sdd.EvaluationResult
}

// FormatText implements format.TextFormatter.
func (e EvaluationOutput) FormatText() ([]byte, error) {
	var sb strings.Builder
	r := e.EvaluationResult

	fmt.Fprintf(&sb, "SDD Type: %s\n", r.SDDType)
	fmt.Fprintf(&sb, "Status: %s\n", statusEmoji(r.Status))
	fmt.Fprintf(&sb, "Score: %.1f%%\n\n", r.Summary.TotalScore*100)

	for _, file := range r.Files {
		fmt.Fprintf(&sb, "## %s %s\n", file.DisplayName, statusEmoji(file.Status))
		fmt.Fprintf(&sb, "Score: %.1f%%\n\n", file.Score*100)

		for _, criterion := range file.Criteria {
			fmt.Fprintf(&sb, "  %s %s (%.0f%%)\n", statusEmoji(criterion.Status), criterion.Title, criterion.Weight*100)
			if criterion.Reasoning != "" {
				fmt.Fprintf(&sb, "    %s\n", criterion.Reasoning)
			}
			if len(criterion.Suggestions) > 0 {
				sb.WriteString("    Suggestions:\n")
				for _, s := range criterion.Suggestions {
					fmt.Fprintf(&sb, "      - %s\n", s)
				}
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString("---\n")
	fmt.Fprintf(&sb, "Summary: %d files, %d passed, %d warned, %d failed, %d skipped\n",
		r.Summary.FilesEvaluated,
		r.Summary.CriteriaPassed,
		r.Summary.CriteriaWarned,
		r.Summary.CriteriaFailed,
		r.Summary.CriteriaSkipped,
	)

	return []byte(sb.String()), nil
}

// FormatMarkdown implements format.MarkdownFormatter.
func (e EvaluationOutput) FormatMarkdown() ([]byte, error) {
	var sb strings.Builder
	r := e.EvaluationResult

	sb.WriteString("# Evaluation Report\n\n")
	fmt.Fprintf(&sb, "**SDD Type**: %s\n\n", r.SDDType)
	fmt.Fprintf(&sb, "**Status**: %s\n\n", statusEmoji(r.Status))
	fmt.Fprintf(&sb, "**Score**: %.1f%%\n\n", r.Summary.TotalScore*100)

	for _, file := range r.Files {
		fmt.Fprintf(&sb, "## %s\n\n", file.DisplayName)
		fmt.Fprintf(&sb, "**Status**: %s | **Score**: %.1f%%\n\n", statusEmoji(file.Status), file.Score*100)

		sb.WriteString("| Criterion | Status | Weight | Score |\n")
		sb.WriteString("|-----------|--------|--------|-------|\n")
		for _, criterion := range file.Criteria {
			fmt.Fprintf(&sb, "| %s | %s | %.0f%% | %.0f%% |\n",
				criterion.Title, statusEmoji(criterion.Status), criterion.Weight*100, criterion.Score*100)
		}
		sb.WriteString("\n")

		for _, criterion := range file.Criteria {
			if criterion.Reasoning != "" {
				fmt.Fprintf(&sb, "### %s\n\n", criterion.Title)
				fmt.Fprintf(&sb, "%s\n\n", criterion.Reasoning)
				if len(criterion.Suggestions) > 0 {
					sb.WriteString("**Suggestions**:\n\n")
					for _, s := range criterion.Suggestions {
						fmt.Fprintf(&sb, "- %s\n", s)
					}
					sb.WriteString("\n")
				}
			}
		}
	}

	sb.WriteString("## Summary\n\n")
	sb.WriteString("| Metric | Value |\n")
	sb.WriteString("|--------|-------|\n")
	fmt.Fprintf(&sb, "| Files Evaluated | %d |\n", r.Summary.FilesEvaluated)
	fmt.Fprintf(&sb, "| Criteria Passed | %d |\n", r.Summary.CriteriaPassed)
	fmt.Fprintf(&sb, "| Criteria Warned | %d |\n", r.Summary.CriteriaWarned)
	fmt.Fprintf(&sb, "| Criteria Failed | %d |\n", r.Summary.CriteriaFailed)
	fmt.Fprintf(&sb, "| Criteria Skipped | %d |\n", r.Summary.CriteriaSkipped)

	return []byte(sb.String()), nil
}

func statusEmoji(status sdd.Status) string {
	switch status {
	case sdd.StatusGo:
		return "[GO]"
	case sdd.StatusWarn:
		return "[WARN]"
	case sdd.StatusNoGo:
		return "[NO-GO]"
	case sdd.StatusSkip:
		return "[SKIP]"
	default:
		return "[?]"
	}
}

// FormatStatus formats a status with optional color codes.
func FormatStatus(status sdd.Status, useColor bool) string {
	emoji := statusEmoji(status)
	if !useColor {
		return emoji
	}

	// ANSI color codes
	var color string
	switch status {
	case sdd.StatusGo:
		color = "\033[32m" // Green
	case sdd.StatusWarn:
		color = "\033[33m" // Yellow
	case sdd.StatusNoGo:
		color = "\033[31m" // Red
	case sdd.StatusSkip:
		color = "\033[90m" // Gray
	default:
		color = ""
	}

	reset := "\033[0m"
	return color + emoji + reset
}

// FormatScore formats a score as a percentage.
func FormatScore(score float64) string {
	return fmt.Sprintf("%.1f%%", score*100)
}
