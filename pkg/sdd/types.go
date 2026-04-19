// Package sdd provides types and utilities for SDD (Spec-Driven Development) definitions.
package sdd

import (
	multiagentspec "github.com/plexusone/multi-agent-spec/sdk/go"
)

// Status is an alias for the multi-agent-spec Status type.
// This provides GO/WARN/NO-GO/SKIP evaluation status.
type Status = multiagentspec.Status

// Status constants are re-exported from multi-agent-spec.
const (
	StatusGo   = multiagentspec.StatusGo
	StatusWarn = multiagentspec.StatusWarn
	StatusNoGo = multiagentspec.StatusNoGo
	StatusSkip = multiagentspec.StatusSkip
)

// SDDType represents a spec-driven development methodology definition.
type SDDType struct {
	Name          string     `yaml:"name"`
	DisplayName   string     `yaml:"display_name"`
	Description   string     `yaml:"description"`
	Extends       string     `yaml:"extends,omitempty"`
	SpecDirectory string     `yaml:"spec_directory"`
	Files         []FileSpec `yaml:"files"`

	// Markdown body content (after frontmatter)
	Body string `yaml:"-"`

	// FileDefinitions holds the parsed file-specific definitions
	FileDefinitions map[string]*FileDefinition `yaml:"-"`
}

// FileSpec defines a spec file pattern and its requirements.
type FileSpec struct {
	Name             string   `yaml:"name"`
	Patterns         []string `yaml:"patterns"`
	Required         bool     `yaml:"required"`
	RequiresCodebase bool     `yaml:"requires_codebase,omitempty"`
}

// FileDefinition represents a file-specific definition with sections and criteria.
type FileDefinition struct {
	File        string          `yaml:"file"`
	DisplayName string          `yaml:"display_name"`
	Sections    []Section       `yaml:"sections"`
	Criteria    []CriterionSpec `yaml:"criteria"`

	// Parsed rubrics from markdown body
	Rubrics map[string]*Rubric `yaml:"-"`

	// Markdown body content
	Body string `yaml:"-"`
}

// Section defines a required or optional section in a spec file.
type Section struct {
	Name        string   `yaml:"name"`
	Required    bool     `yaml:"required"`
	Subsections []string `yaml:"subsections,omitempty"`
	Fields      []string `yaml:"fields,omitempty"`
	MinCount    int      `yaml:"min_count,omitempty"`
	Format      string   `yaml:"format,omitempty"`
}

// CriterionSpec defines evaluation criteria for a file.
type CriterionSpec struct {
	ID     string  `yaml:"id"`
	Weight float64 `yaml:"weight"`
}

// Rubric contains the evaluation guidance for a criterion.
type Rubric struct {
	ID          string
	Title       string
	Description string
	Levels      map[RubricLevel]string
}

// RubricLevel represents evaluation result levels.
type RubricLevel string

const (
	// RubricLevelGo indicates the criterion fully passes.
	RubricLevelGo RubricLevel = "GO"
	// RubricLevelWarn indicates partial compliance with issues.
	RubricLevelWarn RubricLevel = "WARN"
	// RubricLevelNoGo indicates the criterion fails.
	RubricLevelNoGo RubricLevel = "NO-GO"
	// RubricLevelSkip indicates the criterion was skipped.
	RubricLevelSkip RubricLevel = "SKIP"
)


// DetectionResult holds the result of SDD type detection.
type DetectionResult struct {
	Detected    bool
	SDDType     *SDDType
	MatchedFile string
	Confidence  float64
}

// ValidationResult holds the result of structure validation.
type ValidationResult struct {
	Valid          bool
	SDDType        string
	MissingSections []SectionValidation
	PresentSections []SectionValidation
	Errors         []string
	Warnings       []string
}

// SectionValidation holds validation info for a specific section.
type SectionValidation struct {
	File     string
	Section  string
	Required bool
	Found    bool
	Line     int
}

// EvaluationResult holds the result of LLM evaluation.
type EvaluationResult struct {
	SDDType   string
	Files     []FileEvaluation
	Summary   Summary
	Status    Status
}

// FileEvaluation holds evaluation results for a single file.
type FileEvaluation struct {
	File        string
	DisplayName string
	Criteria    []CriterionResult
	Score       float64
	Status      Status
}

// CriterionResult holds the evaluation result for a single criterion.
type CriterionResult struct {
	ID          string
	Title       string
	Weight      float64
	Status      Status
	Score       float64
	Reasoning   string
	Suggestions []string
}

// Summary provides an overall evaluation summary.
type Summary struct {
	TotalScore      float64
	Status          Status
	FilesEvaluated  int
	CriteriaPassed  int
	CriteriaWarned  int
	CriteriaFailed  int
	CriteriaSkipped int
}

// ToTeamReport converts an EvaluationResult to a multi-agent-spec TeamReport.
// This allows devspec output to be consumed by multi-agent-spec tooling.
func (r *EvaluationResult) ToTeamReport(project, version string) *multiagentspec.TeamReport {
	var teams []multiagentspec.TeamSection

	for _, fileEval := range r.Files {
		var tasks []multiagentspec.TaskResult
		for _, cr := range fileEval.Criteria {
			tasks = append(tasks, multiagentspec.TaskResult{
				ID:     cr.ID,
				Status: cr.Status,
				Detail: cr.Reasoning,
				Metadata: map[string]interface{}{
					"title":       cr.Title,
					"weight":      cr.Weight,
					"score":       cr.Score,
					"suggestions": cr.Suggestions,
				},
			})
		}

		teams = append(teams, multiagentspec.TeamSection{
			ID:     fileEval.File,
			Name:   fileEval.DisplayName,
			Tasks:  tasks,
			Status: fileEval.Status,
		})
	}

	return &multiagentspec.TeamReport{
		Schema:  "https://raw.githubusercontent.com/plexusone/multi-agent-spec/main/schema/report/team-report.schema.json",
		Title:   "SDD EVALUATION REPORT",
		Project: project,
		Version: version,
		Phase:   "EVALUATION",
		Teams:   teams,
		Status:  r.Status,
		Tags: map[string]string{
			"sdd_type": r.SDDType,
		},
	}
}
