package evaluate

import (
	"fmt"
	"strings"

	"github.com/plexusone/dev-spec/pkg/sdd"
)

// BuildPrompt constructs an LLM prompt for evaluating a criterion.
func BuildPrompt(rubric *sdd.Rubric, content string) string {
	var sb strings.Builder

	sb.WriteString("You are evaluating a software specification document against specific quality criteria.\n\n")

	sb.WriteString("## Criterion\n\n")
	fmt.Fprintf(&sb, "**%s**: %s\n\n", rubric.ID, rubric.Title)

	sb.WriteString("## Evaluation Rubric\n\n")

	if goLevel, ok := rubric.Levels[sdd.RubricLevelGo]; ok && goLevel != "" {
		sb.WriteString("### GO (Pass)\n")
		sb.WriteString(goLevel)
		sb.WriteString("\n\n")
	}

	if warnLevel, ok := rubric.Levels[sdd.RubricLevelWarn]; ok && warnLevel != "" {
		sb.WriteString("### WARN (Partial)\n")
		sb.WriteString(warnLevel)
		sb.WriteString("\n\n")
	}

	if nogoLevel, ok := rubric.Levels[sdd.RubricLevelNoGo]; ok && nogoLevel != "" {
		sb.WriteString("### NO-GO (Fail)\n")
		sb.WriteString(nogoLevel)
		sb.WriteString("\n\n")
	}

	sb.WriteString("## Document to Evaluate\n\n")
	sb.WriteString("```markdown\n")
	sb.WriteString(content)
	sb.WriteString("\n```\n\n")

	sb.WriteString("## Instructions\n\n")
	sb.WriteString("Evaluate the document against the criterion above. Respond in the following JSON format:\n\n")
	sb.WriteString("```json\n")
	sb.WriteString(`{
  "status": "GO" | "WARN" | "NO-GO",
  "score": 0.0 to 1.0,
  "reasoning": "Brief explanation of your evaluation",
  "suggestions": ["Optional improvement suggestions"]
}`)
	sb.WriteString("\n```\n\n")
	sb.WriteString("Be objective and thorough. Reference specific examples from the document to support your evaluation.")

	return sb.String()
}

// BuildFilePrompt constructs an LLM prompt for evaluating an entire file.
func BuildFilePrompt(fileDef *sdd.FileDefinition, content string) string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "You are evaluating a %s document.\n\n", fileDef.DisplayName)

	sb.WriteString("## Criteria to Evaluate\n\n")

	for _, criterion := range fileDef.Criteria {
		rubric, ok := fileDef.Rubrics[criterion.ID]
		if !ok {
			continue
		}

		fmt.Fprintf(&sb, "### %s (Weight: %.0f%%)\n\n", rubric.Title, criterion.Weight*100)

		if goLevel, ok := rubric.Levels[sdd.RubricLevelGo]; ok && goLevel != "" {
			sb.WriteString("**GO**: ")
			sb.WriteString(summarizeLevel(goLevel))
			sb.WriteString("\n")
		}

		if warnLevel, ok := rubric.Levels[sdd.RubricLevelWarn]; ok && warnLevel != "" {
			sb.WriteString("**WARN**: ")
			sb.WriteString(summarizeLevel(warnLevel))
			sb.WriteString("\n")
		}

		if nogoLevel, ok := rubric.Levels[sdd.RubricLevelNoGo]; ok && nogoLevel != "" {
			sb.WriteString("**NO-GO**: ")
			sb.WriteString(summarizeLevel(nogoLevel))
			sb.WriteString("\n")
		}

		sb.WriteString("\n")
	}

	sb.WriteString("## Document to Evaluate\n\n")
	sb.WriteString("```markdown\n")
	sb.WriteString(content)
	sb.WriteString("\n```\n\n")

	sb.WriteString("## Instructions\n\n")
	sb.WriteString("Evaluate the document against each criterion. Respond with a JSON array:\n\n")
	sb.WriteString("```json\n")
	sb.WriteString(`[
  {
    "criterion_id": "criterion_name",
    "status": "GO" | "WARN" | "NO-GO",
    "score": 0.0 to 1.0,
    "reasoning": "Brief explanation",
    "suggestions": ["Improvements"]
  }
]`)
	sb.WriteString("\n```\n")

	return sb.String()
}

// summarizeLevel extracts the first line or sentence from a rubric level description.
func summarizeLevel(level string) string {
	lines := strings.Split(level, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "-") {
			// Limit to first sentence or 100 chars
			if idx := strings.Index(line, ". "); idx > 0 && idx < 100 {
				return line[:idx+1]
			}
			if len(line) > 100 {
				return line[:97] + "..."
			}
			return line
		}
	}
	return ""
}
