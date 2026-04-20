package sdd

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrNoFrontmatter      = errors.New("no frontmatter found")
	ErrInvalidFrontmatter = errors.New("invalid frontmatter")
)

// LoadSDDType loads an SDD type definition from a filesystem.
// It expects a main definition file (e.g., kiro.md) and file-specific definitions.
func LoadSDDType(fsys fs.FS, mainFile string) (*SDDType, error) {
	content, err := fs.ReadFile(fsys, mainFile)
	if err != nil {
		return nil, fmt.Errorf("read main file: %w", err)
	}

	frontmatter, body, err := ParseFrontmatter(string(content))
	if err != nil {
		return nil, fmt.Errorf("parse frontmatter: %w", err)
	}

	var sddType SDDType
	if err := yaml.Unmarshal([]byte(frontmatter), &sddType); err != nil {
		return nil, fmt.Errorf("unmarshal frontmatter: %w", err)
	}
	sddType.Body = body
	sddType.FileDefinitions = make(map[string]*FileDefinition)

	// Load file-specific definitions from the same directory
	dir := filepath.Dir(mainFile)
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, fmt.Errorf("read directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || entry.Name() == filepath.Base(mainFile) {
			continue
		}
		if !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		filePath := filepath.Join(dir, entry.Name())
		fileDef, err := LoadFileDefinition(fsys, filePath)
		if err != nil {
			// Skip files that don't have proper frontmatter (might be README, etc.)
			continue
		}

		if fileDef.File != "" {
			sddType.FileDefinitions[fileDef.File] = fileDef
		}
	}

	return &sddType, nil
}

// LoadFileDefinition loads a file-specific definition.
func LoadFileDefinition(fsys fs.FS, path string) (*FileDefinition, error) {
	content, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	frontmatter, body, err := ParseFrontmatter(string(content))
	if err != nil {
		return nil, err
	}

	var fileDef FileDefinition
	if err := yaml.Unmarshal([]byte(frontmatter), &fileDef); err != nil {
		return nil, fmt.Errorf("unmarshal frontmatter: %w", err)
	}
	fileDef.Body = body

	// Parse rubrics from markdown body
	fileDef.Rubrics = ParseRubrics(body)

	return &fileDef, nil
}

// ParseFrontmatter extracts YAML frontmatter and markdown body from content.
func ParseFrontmatter(content string) (frontmatter, body string, err error) {
	lines := strings.Split(content, "\n")
	if len(lines) < 3 {
		return "", "", ErrNoFrontmatter
	}

	// Check for opening delimiter
	if strings.TrimSpace(lines[0]) != "---" {
		return "", "", ErrNoFrontmatter
	}

	// Find closing delimiter
	endIndex := -1
	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "---" {
			endIndex = i
			break
		}
	}

	if endIndex == -1 {
		return "", "", ErrInvalidFrontmatter
	}

	frontmatter = strings.Join(lines[1:endIndex], "\n")
	if endIndex+1 < len(lines) {
		body = strings.TrimSpace(strings.Join(lines[endIndex+1:], "\n"))
	}

	return frontmatter, body, nil
}

// ParseRubrics extracts rubrics from markdown content.
// It looks for patterns like:
// ## Criterion: criterion_id (weight%)
// ### GO
// ### WARN
// ### NO-GO
func ParseRubrics(content string) map[string]*Rubric {
	rubrics := make(map[string]*Rubric)

	// Match criterion headers: ## Criterion: id (weight%)
	criterionRe := regexp.MustCompile(`(?m)^##\s+Criterion:\s+(\w+)\s*(?:\(\d+%\))?\s*$`)
	matches := criterionRe.FindAllStringSubmatchIndex(content, -1)

	for i, match := range matches {
		if len(match) < 4 {
			continue
		}

		id := content[match[2]:match[3]]

		// Find the section content (until next ## or end)
		startIdx := match[1]
		endIdx := len(content)
		if i+1 < len(matches) {
			endIdx = matches[i+1][0]
		}

		sectionContent := content[startIdx:endIdx]

		// Extract title (first line after criterion header that starts with **)
		titleRe := regexp.MustCompile(`(?m)^\*\*(.+?)\*\*`)
		titleMatch := titleRe.FindStringSubmatch(sectionContent)
		title := id
		if len(titleMatch) > 1 {
			title = strings.TrimSpace(titleMatch[1])
			// Remove trailing " -" if present
			title = strings.TrimSuffix(title, " -")
		}

		rubric := &Rubric{
			ID:     id,
			Title:  title,
			Levels: make(map[RubricLevel]string),
		}

		// Parse GO/WARN/NO-GO sections
		rubric.Levels[RubricLevelGo] = extractRubricLevel(sectionContent, "GO")
		rubric.Levels[RubricLevelWarn] = extractRubricLevel(sectionContent, "WARN")
		rubric.Levels[RubricLevelNoGo] = extractRubricLevel(sectionContent, "NO-GO")

		rubrics[id] = rubric
	}

	return rubrics
}

// extractRubricLevel extracts the content of a rubric level section.
func extractRubricLevel(content, level string) string {
	// Match ### GO, ### WARN, ### NO-GO
	pattern := fmt.Sprintf(`(?m)^###\s+%s\s*$`, regexp.QuoteMeta(level))
	re := regexp.MustCompile(pattern)
	match := re.FindStringIndex(content)
	if match == nil {
		return ""
	}

	// Find content until next ### or ---
	startIdx := match[1]

	scanner := bufio.NewScanner(strings.NewReader(content[startIdx:]))
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "###") || trimmed == "---" {
			break
		}
		lines = append(lines, line)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}
