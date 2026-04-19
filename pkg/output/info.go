package output

import (
	"fmt"
	"strings"

	"github.com/plexusone/dev-spec/pkg/sdd"
)

// SDDInfo is the structured output for coding assistants.
type SDDInfo struct {
	Name          string     `json:"name"`
	DisplayName   string     `json:"display_name"`
	Description   string     `json:"description"`
	SpecDirectory string     `json:"spec_directory"`
	Files         []FileInfo `json:"files"`
}

// FileInfo describes a file in the SDD.
type FileInfo struct {
	Name             string        `json:"name"`
	DisplayName      string        `json:"display_name"`
	Patterns         []string      `json:"patterns"`
	Required         bool          `json:"required"`
	RequiresCodebase bool          `json:"requires_codebase"`
	Sections         []SectionInfo `json:"sections,omitempty"`
	Template         string        `json:"template,omitempty"`
}

// SectionInfo describes a section in a file.
type SectionInfo struct {
	Name        string   `json:"name"`
	Required    bool     `json:"required"`
	Subsections []string `json:"subsections,omitempty"`
	Fields      []string `json:"fields,omitempty"`
}

// NewSDDInfo creates an SDDInfo from an SDDType.
func NewSDDInfo(sddType *sdd.SDDType) SDDInfo {
	info := SDDInfo{
		Name:          sddType.Name,
		DisplayName:   sddType.DisplayName,
		Description:   sddType.Description,
		SpecDirectory: sddType.SpecDirectory,
	}

	for _, fileSpec := range sddType.Files {
		fileInfo := FileInfo{
			Name:             fileSpec.Name,
			Patterns:         fileSpec.Patterns,
			Required:         fileSpec.Required,
			RequiresCodebase: fileSpec.RequiresCodebase,
		}

		// Add file definition details if available
		if fileDef, ok := sddType.FileDefinitions[fileSpec.Name]; ok {
			fileInfo.DisplayName = fileDef.DisplayName

			for _, section := range fileDef.Sections {
				fileInfo.Sections = append(fileInfo.Sections, SectionInfo{
					Name:        section.Name,
					Required:    section.Required,
					Subsections: section.Subsections,
					Fields:      section.Fields,
				})
			}

			// Generate template
			fileInfo.Template = GenerateTemplateContent(fileDef)
		}

		info.Files = append(info.Files, fileInfo)
	}

	return info
}

// GenerateTemplateContent generates template content from a file definition.
func GenerateTemplateContent(fileDef *sdd.FileDefinition) string {
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

// FormatText implements format.TextFormatter.
func (s SDDInfo) FormatText() ([]byte, error) {
	var sb strings.Builder

	fmt.Fprintf(&sb, "# %s\n\n", s.DisplayName)
	fmt.Fprintf(&sb, "%s\n\n", s.Description)
	fmt.Fprintf(&sb, "Spec Directory: %s\n\n", s.SpecDirectory)

	sb.WriteString("## Files\n\n")
	for _, f := range s.Files {
		req := "optional"
		if f.Required {
			req = "required"
		}
		codebase := ""
		if f.RequiresCodebase {
			codebase = ", needs codebase"
		}
		fmt.Fprintf(&sb, "### %s (%s%s)\n", f.Name, req, codebase)
		fmt.Fprintf(&sb, "Patterns: %v\n", f.Patterns)

		if len(f.Sections) > 0 {
			sb.WriteString("Sections:\n")
			for _, sec := range f.Sections {
				sreq := ""
				if sec.Required {
					sreq = " (required)"
				}
				fmt.Fprintf(&sb, "  - %s%s\n", sec.Name, sreq)
			}
		}
		sb.WriteString("\n")
	}

	return []byte(sb.String()), nil
}

// FormatMarkdown implements format.MarkdownFormatter.
func (s SDDInfo) FormatMarkdown() ([]byte, error) {
	var sb strings.Builder

	fmt.Fprintf(&sb, "# %s\n\n", s.DisplayName)
	fmt.Fprintf(&sb, "%s\n\n", s.Description)
	fmt.Fprintf(&sb, "**Spec Directory**: `%s`\n\n", s.SpecDirectory)

	sb.WriteString("## Files\n\n")
	sb.WriteString("| File | Required | Needs Codebase | Patterns |\n")
	sb.WriteString("|------|----------|----------------|----------|\n")
	for _, f := range s.Files {
		patterns := strings.Join(f.Patterns, ", ")
		fmt.Fprintf(&sb, "| %s | %v | %v | `%s` |\n", f.Name, f.Required, f.RequiresCodebase, patterns)
	}
	sb.WriteString("\n")

	for _, f := range s.Files {
		if len(f.Sections) > 0 {
			fmt.Fprintf(&sb, "### %s Sections\n\n", f.Name)
			sb.WriteString("| Section | Required |\n")
			sb.WriteString("|---------|----------|\n")
			for _, sec := range f.Sections {
				fmt.Fprintf(&sb, "| %s | %v |\n", sec.Name, sec.Required)
			}
			sb.WriteString("\n")
		}
	}

	return []byte(sb.String()), nil
}
