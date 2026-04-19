// Package format provides output format handling for the devspec CLI.
package format

import (
	"encoding/json"
	"fmt"

	"github.com/toon-format/toon-go"
)

// Format represents an output format.
type Format string

// Supported output formats.
const (
	TOON        Format = "toon"
	JSON        Format = "json"
	JSONCompact Format = "json-compact"
	Text        Format = "text"
	Markdown    Format = "markdown"
)

// Default is the default output format.
const Default = TOON

// All returns all supported formats.
func All() []Format {
	return []Format{TOON, JSON, JSONCompact, Text, Markdown}
}

// Parse parses a format string into a Format.
// Empty string defaults to TOON.
func Parse(s string) (Format, error) {
	switch s {
	case "toon", "":
		return TOON, nil
	case "json":
		return JSON, nil
	case "json-compact":
		return JSONCompact, nil
	case "text":
		return Text, nil
	case "markdown", "md":
		return Markdown, nil
	default:
		return "", fmt.Errorf("unknown format %q: use toon, json, json-compact, text, or markdown", s)
	}
}

// IsStructured returns true if the format is a structured data format (TOON, JSON).
func (f Format) IsStructured() bool {
	switch f {
	case TOON, JSON, JSONCompact:
		return true
	default:
		return false
	}
}

// Marshal serializes a value to the specified format.
// For Text and Markdown formats, the value must implement TextFormatter or MarkdownFormatter.
func Marshal(v any, f Format) ([]byte, error) {
	switch f {
	case TOON:
		return toon.Marshal(v)
	case JSON:
		return json.MarshalIndent(v, "", "  ")
	case JSONCompact:
		return json.Marshal(v)
	case Text:
		if tm, ok := v.(TextFormatter); ok {
			return tm.FormatText()
		}
		return nil, fmt.Errorf("value does not implement TextFormatter")
	case Markdown:
		if mm, ok := v.(MarkdownFormatter); ok {
			return mm.FormatMarkdown()
		}
		// Fall back to text if markdown not implemented
		if tm, ok := v.(TextFormatter); ok {
			return tm.FormatText()
		}
		return nil, fmt.Errorf("value does not implement MarkdownFormatter or TextFormatter")
	default:
		return toon.Marshal(v)
	}
}

// TextFormatter is implemented by types that can format themselves as human-readable text.
// Note: This uses FormatText() instead of MarshalText() to avoid conflicting with
// encoding.TextMarshaler, which json.Marshal would call.
type TextFormatter interface {
	FormatText() ([]byte, error)
}

// MarkdownFormatter is implemented by types that can format themselves as markdown.
type MarkdownFormatter interface {
	FormatMarkdown() ([]byte, error)
}
