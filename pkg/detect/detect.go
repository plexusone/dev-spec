// Package detect provides SDD type detection based on file patterns.
package detect

import (
	"os"
	"path/filepath"

	"github.com/plexusone/dev-spec/pkg/sdd"
)

// DetectSDDType attempts to detect which SDD type is present in a directory.
func DetectSDDType(path string) (*sdd.DetectionResult, error) {
	resolver := sdd.NewResolver()
	types, err := resolver.ListTypes()
	if err != nil {
		return nil, err
	}

	var bestMatch *sdd.DetectionResult
	var bestScore float64

	for _, typeName := range types {
		sddType, err := resolver.ResolveType(typeName)
		if err != nil {
			continue
		}

		result := matchSDDType(path, sddType)
		if result.Detected && result.Confidence > bestScore {
			bestMatch = result
			bestScore = result.Confidence
		}
	}

	if bestMatch != nil {
		return bestMatch, nil
	}

	return &sdd.DetectionResult{Detected: false}, nil
}

// matchSDDType checks how well a directory matches an SDD type.
func matchSDDType(path string, sddType *sdd.SDDType) *sdd.DetectionResult {
	result := &sdd.DetectionResult{
		Detected: false,
		SDDType:  sddType,
	}

	// Check if spec directory exists
	specDir := filepath.Join(path, sddType.SpecDirectory)
	if _, err := os.Stat(specDir); os.IsNotExist(err) {
		// Try matching patterns directly in path
		specDir = path
	}

	matchedFiles := 0
	requiredFiles := 0
	requiredMatched := 0
	var firstMatch string

	for _, fileSpec := range sddType.Files {
		if fileSpec.Required {
			requiredFiles++
		}

		for _, pattern := range fileSpec.Patterns {
			matches, err := filepath.Glob(filepath.Join(specDir, pattern))
			if err != nil {
				continue
			}

			// Also try matching in the base path
			if len(matches) == 0 {
				matches, _ = filepath.Glob(filepath.Join(path, pattern))
			}

			if len(matches) > 0 {
				matchedFiles++
				if fileSpec.Required {
					requiredMatched++
				}
				if firstMatch == "" {
					firstMatch = matches[0]
				}
				break
			}
		}
	}

	// Detection criteria:
	// - At least one required file must be present
	// - More matches = higher confidence
	if requiredMatched == 0 && requiredFiles > 0 {
		return result
	}

	if matchedFiles == 0 {
		return result
	}

	result.Detected = true
	result.MatchedFile = firstMatch

	// Calculate confidence based on how many files match
	totalFiles := len(sddType.Files)
	if totalFiles > 0 {
		result.Confidence = float64(matchedFiles) / float64(totalFiles)
	} else {
		result.Confidence = 0.5
	}

	// Boost confidence if all required files are present
	if requiredFiles > 0 && requiredMatched == requiredFiles {
		result.Confidence = (result.Confidence + 1.0) / 2.0
	}

	return result
}

// DetectWithType checks if a specific SDD type is present in a directory.
func DetectWithType(path string, typeName string) (*sdd.DetectionResult, error) {
	resolver := sdd.NewResolver()
	sddType, err := resolver.ResolveType(typeName)
	if err != nil {
		return nil, err
	}

	return matchSDDType(path, sddType), nil
}
