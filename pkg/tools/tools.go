// Package tools provides LLM-callable tools for spec evaluation.
package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/plexusone/dev-spec/pkg/detect"
	"github.com/plexusone/dev-spec/pkg/format"
	"github.com/plexusone/dev-spec/pkg/llm"
	"github.com/plexusone/dev-spec/pkg/output"
	"github.com/plexusone/dev-spec/pkg/sdd"
	"github.com/plexusone/dev-spec/pkg/validate"
)

// BaseTool provides a base implementation for tools.
type BaseTool struct {
	name        string
	description string
	parameters  map[string]interface{}
	handler     func(ctx context.Context, args json.RawMessage) (string, error)
}

// NewBaseTool creates a new base tool.
func NewBaseTool(name, description string, parameters map[string]interface{}, handler func(ctx context.Context, args json.RawMessage) (string, error)) *BaseTool {
	return &BaseTool{
		name:        name,
		description: description,
		parameters:  parameters,
		handler:     handler,
	}
}

func (t *BaseTool) Name() string                       { return t.name }
func (t *BaseTool) Description() string                { return t.description }
func (t *BaseTool) Parameters() map[string]interface{} { return t.parameters }

func (t *BaseTool) Execute(ctx context.Context, args json.RawMessage) (string, error) {
	return t.handler(ctx, args)
}

// Ensure BaseTool implements llm.Tool interface.
var _ llm.Tool = (*BaseTool)(nil)

// NewValidateTool creates the devspec_validate tool.
func NewValidateTool(resolver *sdd.Resolver) llm.Tool {
	return NewBaseTool(
		"devspec_validate",
		"Validate spec file structure against an SDD type. Returns validation results including missing sections and errors.",
		map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path": map[string]interface{}{
					"type":        "string",
					"description": "Path to the directory containing spec files",
				},
				"type": map[string]interface{}{
					"type":        "string",
					"description": "SDD type to validate against (kiro, speckit, plexusone). If not specified, auto-detects.",
				},
			},
			"required": []string{"path"},
		},
		func(ctx context.Context, args json.RawMessage) (string, error) {
			var params struct {
				Path string `json:"path"`
				Type string `json:"type"`
			}
			if err := json.Unmarshal(args, &params); err != nil {
				return "", fmt.Errorf("parse arguments: %w", err)
			}

			// Detect or use specified type
			var sddTypeName string
			if params.Type != "" {
				sddTypeName = params.Type
			} else {
				result, err := detect.DetectSDDType(params.Path)
				if err != nil {
					return "", fmt.Errorf("detection failed: %w", err)
				}
				if !result.Detected {
					return "No SDD type detected in the specified path", nil
				}
				sddTypeName = result.SDDType.Name
			}

			// Run validation
			validationResult, err := validate.ValidateStructure(params.Path, sddTypeName)
			if err != nil {
				return "", fmt.Errorf("validation failed: %w", err)
			}

			// Format as TOON for token efficiency
			out := output.ValidateOutput{ValidationResult: validationResult}
			data, err := format.Marshal(out, format.TOON)
			if err != nil {
				return "", err
			}

			return string(data), nil
		},
	)
}

// NewInfoTool creates the devspec_info tool.
func NewInfoTool(resolver *sdd.Resolver) llm.Tool {
	return NewBaseTool(
		"devspec_info",
		"Get detailed information about an SDD type including required files, sections, and evaluation criteria.",
		map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"type": map[string]interface{}{
					"type":        "string",
					"description": "SDD type to get info for (kiro, speckit, plexusone)",
				},
			},
			"required": []string{"type"},
		},
		func(ctx context.Context, args json.RawMessage) (string, error) {
			var params struct {
				Type string `json:"type"`
			}
			if err := json.Unmarshal(args, &params); err != nil {
				return "", fmt.Errorf("parse arguments: %w", err)
			}

			sddType, err := resolver.ResolveType(params.Type)
			if err != nil {
				return "", fmt.Errorf("unknown SDD type %q: %w", params.Type, err)
			}

			info := output.NewSDDInfo(sddType)

			// Format as TOON for token efficiency
			data, err := format.Marshal(info, format.TOON)
			if err != nil {
				return "", err
			}

			return string(data), nil
		},
	)
}

// NewRubricsTool creates the devspec_rubrics tool.
func NewRubricsTool(resolver *sdd.Resolver) llm.Tool {
	return NewBaseTool(
		"devspec_rubrics",
		"Get evaluation rubrics for an SDD type. Returns GO/WARN/NO-GO criteria for evaluating spec files.",
		map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"type": map[string]interface{}{
					"type":        "string",
					"description": "SDD type to get rubrics for (kiro, speckit, plexusone)",
				},
				"file": map[string]interface{}{
					"type":        "string",
					"description": "Specific file to get rubrics for (e.g., requirements, design, tasks). If not specified, returns all rubrics.",
				},
			},
			"required": []string{"type"},
		},
		func(ctx context.Context, args json.RawMessage) (string, error) {
			var params struct {
				Type string `json:"type"`
				File string `json:"file"`
			}
			if err := json.Unmarshal(args, &params); err != nil {
				return "", fmt.Errorf("parse arguments: %w", err)
			}

			sddType, err := resolver.ResolveType(params.Type)
			if err != nil {
				return "", fmt.Errorf("unknown SDD type %q: %w", params.Type, err)
			}

			out := output.NewRubricsOutput(sddType, params.File)

			// Format as TOON for token efficiency
			data, err := format.Marshal(out, format.TOON)
			if err != nil {
				return "", err
			}

			return string(data), nil
		},
	)
}

// NewReadFileTool creates the read_file tool for reading spec and code files.
func NewReadFileTool(basePath string) llm.Tool {
	return NewBaseTool(
		"read_file",
		"Read the contents of a file. Use this to read spec files or source code files for evaluation.",
		map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path": map[string]interface{}{
					"type":        "string",
					"description": "Path to the file to read (relative to the project root or absolute)",
				},
			},
			"required": []string{"path"},
		},
		func(ctx context.Context, args json.RawMessage) (string, error) {
			var params struct {
				Path string `json:"path"`
			}
			if err := json.Unmarshal(args, &params); err != nil {
				return "", fmt.Errorf("parse arguments: %w", err)
			}

			// Resolve path relative to basePath if not absolute
			filePath := params.Path
			if !filepath.IsAbs(filePath) {
				filePath = filepath.Join(basePath, filePath)
			}

			// Security: ensure path is within basePath
			absPath, err := filepath.Abs(filePath)
			if err != nil {
				return "", fmt.Errorf("resolve path: %w", err)
			}
			absBase, err := filepath.Abs(basePath)
			if err != nil {
				return "", fmt.Errorf("resolve base path: %w", err)
			}

			rel, err := filepath.Rel(absBase, absPath)
			if err != nil || len(rel) > 2 && rel[:2] == ".." {
				return "", fmt.Errorf("path outside project directory: %s", params.Path)
			}

			content, err := os.ReadFile(absPath)
			if err != nil {
				if os.IsNotExist(err) {
					return fmt.Sprintf("File not found: %s", params.Path), nil
				}
				return "", fmt.Errorf("read file: %w", err)
			}

			// Limit file size for token efficiency
			const maxSize = 50000
			if len(content) > maxSize {
				return string(content[:maxSize]) + "\n\n[File truncated - exceeded 50KB limit]", nil
			}

			return string(content), nil
		},
	)
}

// NewGlobFilesTool creates the glob_files tool for finding files.
func NewGlobFilesTool(basePath string) llm.Tool {
	return NewBaseTool(
		"glob_files",
		"Find files matching a glob pattern. Use this to discover spec files or source code files.",
		map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"pattern": map[string]interface{}{
					"type":        "string",
					"description": "Glob pattern to match files (e.g., '**/*.md', 'src/**/*.go')",
				},
				"path": map[string]interface{}{
					"type":        "string",
					"description": "Base path to search from (default: project root)",
				},
			},
			"required": []string{"pattern"},
		},
		func(ctx context.Context, args json.RawMessage) (string, error) {
			var params struct {
				Pattern string `json:"pattern"`
				Path    string `json:"path"`
			}
			if err := json.Unmarshal(args, &params); err != nil {
				return "", fmt.Errorf("parse arguments: %w", err)
			}

			searchPath := basePath
			if params.Path != "" {
				if filepath.IsAbs(params.Path) {
					searchPath = params.Path
				} else {
					searchPath = filepath.Join(basePath, params.Path)
				}
			}

			// Security: ensure path is within basePath
			absPath, err := filepath.Abs(searchPath)
			if err != nil {
				return "", fmt.Errorf("resolve path: %w", err)
			}
			absBase, err := filepath.Abs(basePath)
			if err != nil {
				return "", fmt.Errorf("resolve base path: %w", err)
			}

			rel, err := filepath.Rel(absBase, absPath)
			if err != nil || len(rel) > 2 && rel[:2] == ".." {
				return "", fmt.Errorf("path outside project directory")
			}

			// Use filepath.Glob for simple patterns
			pattern := filepath.Join(absPath, params.Pattern)
			matches, err := filepath.Glob(pattern)
			if err != nil {
				return "", fmt.Errorf("glob pattern error: %w", err)
			}

			if len(matches) == 0 {
				return "No files found matching pattern: " + params.Pattern, nil
			}

			// Convert to relative paths for readability
			var result []string
			for _, match := range matches {
				relPath, err := filepath.Rel(absBase, match)
				if err != nil {
					relPath = match
				}
				result = append(result, relPath)
			}

			// Format as newline-separated list
			output := fmt.Sprintf("Found %d files:\n", len(result))
			for _, path := range result {
				output += "  " + path + "\n"
			}

			return output, nil
		},
	)
}

// NewCheckTool creates the devspec_check tool for detecting SDD types.
func NewCheckTool() llm.Tool {
	return NewBaseTool(
		"devspec_check",
		"Detect the SDD type in a directory based on file patterns.",
		map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path": map[string]interface{}{
					"type":        "string",
					"description": "Path to check for SDD type (default: current directory)",
				},
			},
			"required": []string{},
		},
		func(ctx context.Context, args json.RawMessage) (string, error) {
			var params struct {
				Path string `json:"path"`
			}
			if err := json.Unmarshal(args, &params); err != nil {
				return "", fmt.Errorf("parse arguments: %w", err)
			}

			path := params.Path
			if path == "" {
				path = "."
			}

			result, err := detect.DetectSDDType(path)
			if err != nil {
				return "", fmt.Errorf("detection failed: %w", err)
			}

			out := output.CheckOutput{
				Detected: result.Detected,
			}
			if result.SDDType != nil {
				out.SDDType = result.SDDType.Name
				out.DisplayName = result.SDDType.DisplayName
				out.Description = result.SDDType.Description
				out.SpecDirectory = result.SDDType.SpecDirectory
			}
			out.MatchedFile = result.MatchedFile
			out.Confidence = result.Confidence

			// Format as TOON for token efficiency
			data, err := format.Marshal(out, format.TOON)
			if err != nil {
				return "", err
			}

			return string(data), nil
		},
	)
}

// RegisterAllTools registers all devspec tools with the client.
func RegisterAllTools(client *llm.Client, basePath string) {
	resolver := sdd.NewResolver()

	client.RegisterTools(
		NewCheckTool(),
		NewValidateTool(resolver),
		NewInfoTool(resolver),
		NewRubricsTool(resolver),
		NewReadFileTool(basePath),
		NewGlobFilesTool(basePath),
	)
}
