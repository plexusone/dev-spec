package evaluate

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/plexusone/dev-spec/pkg/sdd"
)

// llmResponse represents the expected JSON response from the LLM.
type llmResponse struct {
	Status      string   `json:"status"`
	Score       float64  `json:"score"`
	Reasoning   string   `json:"reasoning"`
	Suggestions []string `json:"suggestions"`
}

// toolEvaluationResponse represents the response from tool-based evaluation.
type toolEvaluationResponse struct {
	Files []struct {
		File        string `json:"file"`
		DisplayName string `json:"display_name"`
		Criteria    []struct {
			CriterionID string   `json:"criterion_id"`
			Status      string   `json:"status"`
			Score       float64  `json:"score"`
			Reasoning   string   `json:"reasoning"`
			Suggestions []string `json:"suggestions"`
		} `json:"criteria"`
	} `json:"files"`
	OverallStatus string `json:"overall_status"`
	Summary       string `json:"summary"`
}

// ParseResponse parses an LLM response into a CriterionResult.
func ParseResponse(criterion sdd.CriterionSpec, rubric *sdd.Rubric, response string) (sdd.CriterionResult, error) {
	result := sdd.CriterionResult{
		ID:     criterion.ID,
		Title:  rubric.Title,
		Weight: criterion.Weight,
	}

	// Extract JSON from response (it might be wrapped in markdown code blocks)
	jsonStr := extractJSON(response)
	if jsonStr == "" {
		return result, fmt.Errorf("no JSON found in response")
	}

	var llmResp llmResponse
	if err := json.Unmarshal([]byte(jsonStr), &llmResp); err != nil {
		return result, fmt.Errorf("parse JSON: %w", err)
	}

	// Map status string to Status type
	switch strings.ToUpper(llmResp.Status) {
	case "GO":
		result.Status = sdd.StatusGo
	case "WARN":
		result.Status = sdd.StatusWarn
	case "NO-GO", "NOGO":
		result.Status = sdd.StatusNoGo
	case "SKIP":
		result.Status = sdd.StatusSkip
	default:
		return result, fmt.Errorf("unknown status: %s", llmResp.Status)
	}

	result.Score = llmResp.Score
	result.Reasoning = llmResp.Reasoning
	result.Suggestions = llmResp.Suggestions

	// Validate score is in range
	if result.Score < 0 {
		result.Score = 0
	}
	if result.Score > 1 {
		result.Score = 1
	}

	return result, nil
}

// ParseEvaluationResponse parses a tool-based evaluation response into EvaluationResult.
func ParseEvaluationResponse(sddType *sdd.SDDType, response string) (*sdd.EvaluationResult, error) {
	// Extract JSON from response
	jsonStr := extractJSON(response)
	if jsonStr == "" {
		return nil, fmt.Errorf("no JSON found in evaluation response")
	}

	var evalResp toolEvaluationResponse
	if err := json.Unmarshal([]byte(jsonStr), &evalResp); err != nil {
		return nil, fmt.Errorf("parse evaluation JSON: %w", err)
	}

	result := &sdd.EvaluationResult{
		SDDType: sddType.Name,
	}

	// Convert files
	for _, fileResp := range evalResp.Files {
		fileEval := sdd.FileEvaluation{
			File:        fileResp.File,
			DisplayName: fileResp.DisplayName,
		}

		for _, cr := range fileResp.Criteria {
			criterionResult := sdd.CriterionResult{
				ID:          cr.CriterionID,
				Score:       cr.Score,
				Reasoning:   cr.Reasoning,
				Suggestions: cr.Suggestions,
			}

			// Map status
			switch strings.ToUpper(cr.Status) {
			case "GO":
				criterionResult.Status = sdd.StatusGo
				result.Summary.CriteriaPassed++
			case "WARN":
				criterionResult.Status = sdd.StatusWarn
				result.Summary.CriteriaWarned++
			case "NO-GO", "NOGO":
				criterionResult.Status = sdd.StatusNoGo
				result.Summary.CriteriaFailed++
			case "SKIP":
				criterionResult.Status = sdd.StatusSkip
				result.Summary.CriteriaSkipped++
			default:
				criterionResult.Status = sdd.StatusSkip
				result.Summary.CriteriaSkipped++
			}

			// Validate score
			if criterionResult.Score < 0 {
				criterionResult.Score = 0
			}
			if criterionResult.Score > 1 {
				criterionResult.Score = 1
			}

			fileEval.Criteria = append(fileEval.Criteria, criterionResult)
		}

		// Calculate file score and status
		fileEval.Score = calculateFileScore(&fileEval)
		fileEval.Status = calculateFileStatus(&fileEval)

		result.Files = append(result.Files, fileEval)
	}

	// Set overall status
	switch strings.ToUpper(evalResp.OverallStatus) {
	case "GO":
		result.Status = sdd.StatusGo
	case "WARN":
		result.Status = sdd.StatusWarn
	case "NO-GO", "NOGO":
		result.Status = sdd.StatusNoGo
	default:
		result.Status = calculateOverallStatus(result)
	}

	// Calculate summary
	result.Summary.FilesEvaluated = len(result.Files)
	result.Summary.TotalScore = calculateTotalScore(result)
	result.Summary.Status = result.Status

	return result, nil
}

// ParseBatchResponse parses an LLM response containing multiple criteria results.
func ParseBatchResponse(criteria []sdd.CriterionSpec, rubrics map[string]*sdd.Rubric, response string) ([]sdd.CriterionResult, error) {
	jsonStr := extractJSON(response)
	if jsonStr == "" {
		return nil, fmt.Errorf("no JSON found in response")
	}

	var llmResps []struct {
		CriterionID string   `json:"criterion_id"`
		Status      string   `json:"status"`
		Score       float64  `json:"score"`
		Reasoning   string   `json:"reasoning"`
		Suggestions []string `json:"suggestions"`
	}

	if err := json.Unmarshal([]byte(jsonStr), &llmResps); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}

	// Map responses by criterion ID
	respMap := make(map[string]int)
	for i, resp := range llmResps {
		respMap[resp.CriterionID] = i
	}

	var results []sdd.CriterionResult

	for _, criterion := range criteria {
		rubric, ok := rubrics[criterion.ID]
		if !ok {
			continue
		}

		result := sdd.CriterionResult{
			ID:     criterion.ID,
			Title:  rubric.Title,
			Weight: criterion.Weight,
		}

		if idx, ok := respMap[criterion.ID]; ok {
			resp := llmResps[idx]

			switch strings.ToUpper(resp.Status) {
			case "GO":
				result.Status = sdd.StatusGo
			case "WARN":
				result.Status = sdd.StatusWarn
			case "NO-GO", "NOGO":
				result.Status = sdd.StatusNoGo
			case "SKIP":
				result.Status = sdd.StatusSkip
			default:
				result.Status = sdd.StatusSkip
				result.Reasoning = fmt.Sprintf("Unknown status: %s", resp.Status)
			}

			result.Score = resp.Score
			result.Reasoning = resp.Reasoning
			result.Suggestions = resp.Suggestions

			if result.Score < 0 {
				result.Score = 0
			}
			if result.Score > 1 {
				result.Score = 1
			}
		} else {
			result.Status = sdd.StatusSkip
			result.Reasoning = "No evaluation provided"
		}

		results = append(results, result)
	}

	return results, nil
}

// extractJSON extracts JSON from a response that might contain markdown code blocks.
func extractJSON(response string) string {
	// Try to find JSON in code blocks first
	codeBlockRe := regexp.MustCompile("(?s)```(?:json)?\\s*\n?(.*?)```")
	if matches := codeBlockRe.FindStringSubmatch(response); len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}

	// Try to find raw JSON (object or array)
	response = strings.TrimSpace(response)

	// Check for JSON object
	if strings.HasPrefix(response, "{") {
		end := strings.LastIndex(response, "}")
		if end > 0 {
			return response[:end+1]
		}
	}

	// Check for JSON array
	if strings.HasPrefix(response, "[") {
		end := strings.LastIndex(response, "]")
		if end > 0 {
			return response[:end+1]
		}
	}

	return ""
}
