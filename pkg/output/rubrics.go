package output

import (
	"fmt"
	"strings"

	"github.com/plexusone/dev-spec/pkg/sdd"
)

// RubricsOutput is the structured rubrics output for coding assistants.
type RubricsOutput struct {
	SDDType     string            `json:"sdd_type"`
	DisplayName string            `json:"display_name"`
	Files       []FileRubricsInfo `json:"files"`
}

// FileRubricsInfo contains rubrics for a single file type.
type FileRubricsInfo struct {
	File        string         `json:"file"`
	DisplayName string         `json:"display_name"`
	Criteria    []CriteriaInfo `json:"criteria"`
}

// CriteriaInfo contains a single criterion with its rubric.
type CriteriaInfo struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Weight      float64 `json:"weight"`
	Description string  `json:"description,omitempty"`
	Go          string  `json:"go"`
	Warn        string  `json:"warn"`
	NoGo        string  `json:"no_go"`
}

// NewRubricsOutput creates a RubricsOutput from an SDDType.
// If fileFilter is non-empty, only that file's rubrics are included.
func NewRubricsOutput(sddType *sdd.SDDType, fileFilter string) RubricsOutput {
	output := RubricsOutput{
		SDDType:     sddType.Name,
		DisplayName: sddType.DisplayName,
	}

	for _, fileSpec := range sddType.Files {
		// Skip if filtering to specific file
		if fileFilter != "" && fileSpec.Name != fileFilter {
			continue
		}

		fileDef, ok := sddType.FileDefinitions[fileSpec.Name]
		if !ok {
			continue
		}

		fileRubrics := FileRubricsInfo{
			File:        fileSpec.Name,
			DisplayName: fileDef.DisplayName,
		}

		for _, criterion := range fileDef.Criteria {
			rubric, ok := fileDef.Rubrics[criterion.ID]
			if !ok {
				continue
			}

			criteriaInfo := CriteriaInfo{
				ID:     criterion.ID,
				Title:  rubric.Title,
				Weight: criterion.Weight,
				Go:     rubric.Levels[sdd.RubricLevelGo],
				Warn:   rubric.Levels[sdd.RubricLevelWarn],
				NoGo:   rubric.Levels[sdd.RubricLevelNoGo],
			}

			fileRubrics.Criteria = append(fileRubrics.Criteria, criteriaInfo)
		}

		output.Files = append(output.Files, fileRubrics)
	}

	return output
}

// FormatText implements format.TextFormatter.
func (r RubricsOutput) FormatText() ([]byte, error) {
	var sb strings.Builder

	fmt.Fprintf(&sb, "# %s Evaluation Rubrics\n\n", r.DisplayName)

	for _, file := range r.Files {
		fmt.Fprintf(&sb, "## %s\n\n", file.DisplayName)

		for _, c := range file.Criteria {
			fmt.Fprintf(&sb, "### %s (%.0f%%)\n\n", c.Title, c.Weight*100)

			sb.WriteString("**GO**:\n")
			fmt.Fprintf(&sb, "%s\n\n", c.Go)

			sb.WriteString("**WARN**:\n")
			fmt.Fprintf(&sb, "%s\n\n", c.Warn)

			sb.WriteString("**NO-GO**:\n")
			fmt.Fprintf(&sb, "%s\n\n", c.NoGo)

			sb.WriteString("---\n\n")
		}
	}

	return []byte(sb.String()), nil
}

// FormatMarkdown implements format.MarkdownFormatter.
func (r RubricsOutput) FormatMarkdown() ([]byte, error) {
	var sb strings.Builder

	fmt.Fprintf(&sb, "# %s Evaluation Rubrics\n\n", r.DisplayName)

	for _, file := range r.Files {
		fmt.Fprintf(&sb, "## %s\n\n", file.DisplayName)

		for _, c := range file.Criteria {
			fmt.Fprintf(&sb, "### %s\n\n", c.Title)
			fmt.Fprintf(&sb, "**Weight**: %.0f%%\n\n", c.Weight*100)

			sb.WriteString("#### GO (Pass)\n\n")
			fmt.Fprintf(&sb, "%s\n\n", c.Go)

			sb.WriteString("#### WARN (Partial)\n\n")
			fmt.Fprintf(&sb, "%s\n\n", c.Warn)

			sb.WriteString("#### NO-GO (Fail)\n\n")
			fmt.Fprintf(&sb, "%s\n\n", c.NoGo)

			sb.WriteString("---\n\n")
		}
	}

	return []byte(sb.String()), nil
}
