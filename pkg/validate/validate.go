// Package validate provides structure validation for spec files.
package validate

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/plexusone/dev-spec/pkg/sdd"
)

// ValidateStructure validates the structure of spec files against an SDD type.
func ValidateStructure(path string, typeName string) (*sdd.ValidationResult, error) {
	resolver := sdd.NewResolver()
	sddType, err := resolver.ResolveType(typeName)
	if err != nil {
		return nil, fmt.Errorf("resolve SDD type: %w", err)
	}

	result := &sdd.ValidationResult{
		Valid:   true,
		SDDType: typeName,
	}

	// Find spec directory
	specDir := filepath.Join(path, sddType.SpecDirectory)
	if _, err := os.Stat(specDir); os.IsNotExist(err) {
		specDir = path
	}

	// Check each file spec
	for _, fileSpec := range sddType.Files {
		// Find matching file
		var matchedFile string
		for _, pattern := range fileSpec.Patterns {
			matches, err := filepath.Glob(filepath.Join(specDir, pattern))
			if err != nil {
				continue
			}
			if len(matches) == 0 {
				matches, _ = filepath.Glob(filepath.Join(path, pattern))
			}
			if len(matches) > 0 {
				matchedFile = matches[0]
				break
			}
		}

		if matchedFile == "" {
			if fileSpec.Required {
				result.Valid = false
				result.Errors = append(result.Errors, fmt.Sprintf("Required file missing: %s", fileSpec.Name))
			} else {
				result.Warnings = append(result.Warnings, fmt.Sprintf("Optional file missing: %s", fileSpec.Name))
			}
			continue
		}

		// Validate sections if file definition exists
		fileDef, hasDef := sddType.FileDefinitions[fileSpec.Name]
		if !hasDef {
			continue
		}

		// Parse sections from file
		sections, err := parseSections(matchedFile)
		if err != nil {
			result.Warnings = append(result.Warnings, fmt.Sprintf("Could not parse %s: %v", matchedFile, err))
			continue
		}

		// Check each expected section
		for _, sectionDef := range fileDef.Sections {
			found, line := findSection(sections, sectionDef.Name)
			sv := sdd.SectionValidation{
				File:     fileSpec.Name,
				Section:  sectionDef.Name,
				Required: sectionDef.Required,
				Found:    found,
				Line:     line,
			}

			if found {
				result.PresentSections = append(result.PresentSections, sv)
			} else {
				result.MissingSections = append(result.MissingSections, sv)
				if sectionDef.Required {
					result.Valid = false
					result.Errors = append(result.Errors, fmt.Sprintf("Required section missing in %s: %s", fileSpec.Name, sectionDef.Name))
				}
			}
		}
	}

	return result, nil
}

// parseSections extracts section headers from a markdown file.
func parseSections(filepath string) (sections map[string]int, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	sections = make(map[string]int)
	scanner := bufio.NewScanner(file)
	lineNum := 0

	// Match ## Section Name or # Section Name
	headerRe := regexp.MustCompile(`^(#{1,6})\s+(.+)$`)

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if match := headerRe.FindStringSubmatch(line); match != nil {
			sectionName := strings.TrimSpace(match[2])
			// Normalize section name (remove trailing punctuation, extra markup)
			sectionName = strings.TrimRight(sectionName, ":")
			sectionName = strings.TrimSuffix(sectionName, " (Required)")
			sectionName = strings.TrimSuffix(sectionName, " (Optional)")
			sections[sectionName] = lineNum
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sections, nil
}

// findSection checks if a section exists in the parsed sections.
func findSection(sections map[string]int, name string) (bool, int) {
	// Direct match
	if line, ok := sections[name]; ok {
		return true, line
	}

	// Case-insensitive match
	nameLower := strings.ToLower(name)
	for sectionName, line := range sections {
		if strings.ToLower(sectionName) == nameLower {
			return true, line
		}
	}

	// Partial match (section name contains the target)
	for sectionName, line := range sections {
		if strings.Contains(strings.ToLower(sectionName), nameLower) {
			return true, line
		}
	}

	return false, 0
}
