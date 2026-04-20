// Package scaffold provides spec file scaffolding functionality.
package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/plexusone/dev-spec/pkg/sdd"
)

// Options configures scaffold behavior.
type Options struct {
	Force   bool // Overwrite existing files
	Verbose bool // Print progress
}

// Result contains the result of scaffolding.
type Result struct {
	SpecDirectory string
	CreatedFiles  []string
	SkippedFiles  []string
}

// Init initializes spec scaffolding for an SDD type in the given path.
func Init(basePath string, sddType *sdd.SDDType, opts Options) (*Result, error) {
	result := &Result{}

	originalSpecDir := sddType.SpecDirectory
	if originalSpecDir == "" {
		originalSpecDir = "specs"
	}

	// Prepend base path to spec directory
	specDir := filepath.Join(basePath, originalSpecDir)
	result.SpecDirectory = specDir

	// Create spec directory
	if err := os.MkdirAll(specDir, 0755); err != nil {
		return nil, fmt.Errorf("create spec directory: %w", err)
	}

	// Create template files for each file spec
	for _, fileSpec := range sddType.Files {
		if len(fileSpec.Patterns) == 0 {
			continue
		}

		// Use first pattern to determine filename
		filename := fileSpec.Patterns[0]
		// If pattern has directory prefix matching original spec dir, prepend only basePath
		// Otherwise, prepend full spec directory (basePath + specDir)
		var fullPath string
		if strings.HasPrefix(filename, originalSpecDir) {
			fullPath = filepath.Join(basePath, filename)
		} else {
			fullPath = filepath.Join(specDir, filename)
		}

		// Check if file exists
		if _, err := os.Stat(fullPath); err == nil && !opts.Force {
			result.SkippedFiles = append(result.SkippedFiles, fullPath)
			continue
		}

		// Ensure parent directory exists
		if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return nil, fmt.Errorf("create directory for %s: %w", fullPath, err)
		}

		// Generate template content
		content := GenerateTemplate(sddType, fileSpec)

		if err := os.WriteFile(fullPath, []byte(content), 0600); err != nil {
			return nil, fmt.Errorf("write %s: %w", fullPath, err)
		}

		result.CreatedFiles = append(result.CreatedFiles, fullPath)
	}

	return result, nil
}

// GenerateTemplate generates template content for a file spec.
func GenerateTemplate(sddType *sdd.SDDType, fileSpec sdd.FileSpec) string {
	fileDef, ok := sddType.FileDefinitions[fileSpec.Name]
	if !ok {
		return generateBasicTemplate(fileSpec)
	}

	var content string
	content += fmt.Sprintf("# %s\n\n", fileDef.DisplayName)

	for _, section := range fileDef.Sections {
		required := ""
		if section.Required {
			required = " (Required)"
		}
		content += fmt.Sprintf("## %s%s\n\n", section.Name, required)

		if len(section.Fields) > 0 {
			for _, field := range section.Fields {
				content += fmt.Sprintf("- **%s**: \n", field)
			}
			content += "\n"
		}

		if len(section.Subsections) > 0 {
			for _, sub := range section.Subsections {
				content += fmt.Sprintf("### %s\n\n[TODO: Add content]\n\n", sub)
			}
		} else {
			content += "[TODO: Add content]\n\n"
		}
	}

	return content
}

func generateBasicTemplate(fileSpec sdd.FileSpec) string {
	return fmt.Sprintf(`# %s

## Overview

[TODO: Add overview]

## Details

[TODO: Add details]
`, fileSpec.Name)
}
