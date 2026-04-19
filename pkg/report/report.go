// Package report provides report generation for evaluation results.
package report

import (
	"github.com/plexusone/dev-spec/pkg/sdd"
)

// TeamReport represents a complete evaluation report in multi-agent-spec format.
type TeamReport struct {
	TeamName string        `json:"team_name"`
	Status   string        `json:"status"`
	Agents   []AgentResult `json:"agents"`
	Summary  ReportSummary `json:"summary"`
}

// AgentResult represents an agent's evaluation result.
type AgentResult struct {
	AgentName string       `json:"agent_name"`
	Status    string       `json:"status"`
	Tasks     []TaskResult `json:"tasks"`
}

// TaskResult represents a task's evaluation result.
type TaskResult struct {
	TaskName    string   `json:"task_name"`
	Status      string   `json:"status"`
	Score       float64  `json:"score,omitempty"`
	Reasoning   string   `json:"reasoning,omitempty"`
	Suggestions []string `json:"suggestions,omitempty"`
}

// ReportSummary provides overall statistics.
type ReportSummary struct {
	TotalScore      float64 `json:"total_score"`
	FilesEvaluated  int     `json:"files_evaluated"`
	CriteriaPassed  int     `json:"criteria_passed"`
	CriteriaWarned  int     `json:"criteria_warned"`
	CriteriaFailed  int     `json:"criteria_failed"`
	CriteriaSkipped int     `json:"criteria_skipped"`
}

// FromEvaluationResult converts an EvaluationResult to a TeamReport.
func FromEvaluationResult(result *sdd.EvaluationResult) *TeamReport {
	report := &TeamReport{
		TeamName: "spec-evaluation",
		Status:   string(result.Status),
		Summary: ReportSummary{
			TotalScore:      result.Summary.TotalScore,
			FilesEvaluated:  result.Summary.FilesEvaluated,
			CriteriaPassed:  result.Summary.CriteriaPassed,
			CriteriaWarned:  result.Summary.CriteriaWarned,
			CriteriaFailed:  result.Summary.CriteriaFailed,
			CriteriaSkipped: result.Summary.CriteriaSkipped,
		},
	}

	for _, file := range result.Files {
		agent := AgentResult{
			AgentName: file.DisplayName + " Judge",
			Status:    string(file.Status),
		}

		for _, criterion := range file.Criteria {
			task := TaskResult{
				TaskName:    criterion.Title,
				Status:      string(criterion.Status),
				Score:       criterion.Score,
				Reasoning:   criterion.Reasoning,
				Suggestions: criterion.Suggestions,
			}
			agent.Tasks = append(agent.Tasks, task)
		}

		report.Agents = append(report.Agents, agent)
	}

	return report
}
