package sdd

import (
	"embed"
)

//go:embed definitions
var EmbeddedDefinitions embed.FS

// BuiltinSDDTypes returns the names of all built-in SDD types.
func BuiltinSDDTypes() []string {
	return []string{"kiro", "speckit", "plexusone"}
}
