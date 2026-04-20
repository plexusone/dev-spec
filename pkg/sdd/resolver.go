package sdd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var (
	ErrSDDTypeNotFound = errors.New("SDD type not found")
	ErrCircularExtends = errors.New("circular extends detected")
)

// Resolver discovers and loads SDD type definitions.
type Resolver struct {
	// Search paths for definitions (in priority order)
	searchPaths []string

	// Embedded filesystem for built-in definitions
	embeddedFS fs.FS

	// Cache of loaded definitions
	cache map[string]*SDDType
}

// NewResolver creates a new Resolver with default search paths.
func NewResolver() *Resolver {
	homeDir, _ := os.UserHomeDir()

	searchPaths := []string{
		".devspec/definitions",                                // Project-local
		filepath.Join(homeDir, ".config/devspec/definitions"), // User-global
	}

	return &Resolver{
		searchPaths: searchPaths,
		embeddedFS:  EmbeddedDefinitions,
		cache:       make(map[string]*SDDType),
	}
}

// ResolveType loads and resolves an SDD type by name.
// It searches in order: project-local, user-global, built-in.
func (r *Resolver) ResolveType(name string) (*SDDType, error) {
	// Check cache first
	if cached, ok := r.cache[name]; ok {
		return cached, nil
	}

	// Search in local and global paths
	for _, searchPath := range r.searchPaths {
		defPath := filepath.Join(searchPath, name, name+".md")
		if _, err := os.Stat(defPath); err == nil {
			sddType, err := LoadSDDType(os.DirFS(searchPath), filepath.Join(name, name+".md"))
			if err != nil {
				continue
			}
			resolved, err := r.resolveInheritance(sddType, nil)
			if err != nil {
				return nil, err
			}
			r.cache[name] = resolved
			return resolved, nil
		}
	}

	// Try embedded definitions
	defPath := filepath.Join("definitions", name, name+".md")
	sddType, err := LoadSDDType(r.embeddedFS, defPath)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrSDDTypeNotFound, name)
	}

	resolved, err := r.resolveInheritance(sddType, nil)
	if err != nil {
		return nil, err
	}
	r.cache[name] = resolved
	return resolved, nil
}

// ListTypes returns all available SDD types.
func (r *Resolver) ListTypes() ([]string, error) {
	types := make(map[string]struct{})

	// Add built-in types
	for _, name := range BuiltinSDDTypes() {
		types[name] = struct{}{}
	}

	// Add types from search paths
	for _, searchPath := range r.searchPaths {
		entries, err := os.ReadDir(searchPath)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			if entry.IsDir() {
				types[entry.Name()] = struct{}{}
			}
		}
	}

	result := make([]string, 0, len(types))
	for name := range types {
		result = append(result, name)
	}
	return result, nil
}

// resolveInheritance resolves the extends chain for an SDD type.
func (r *Resolver) resolveInheritance(sddType *SDDType, visited map[string]struct{}) (*SDDType, error) {
	if sddType.Extends == "" {
		return sddType, nil
	}

	if visited == nil {
		visited = make(map[string]struct{})
	}

	if _, ok := visited[sddType.Name]; ok {
		return nil, fmt.Errorf("%w: %s", ErrCircularExtends, sddType.Name)
	}
	visited[sddType.Name] = struct{}{}

	parent, err := r.ResolveType(sddType.Extends)
	if err != nil {
		return nil, fmt.Errorf("resolve parent %s: %w", sddType.Extends, err)
	}

	// Merge parent into child (child overrides parent)
	merged := r.mergeTypes(parent, sddType)
	return merged, nil
}

// mergeTypes merges a parent and child SDD type, with child taking precedence.
func (r *Resolver) mergeTypes(parent, child *SDDType) *SDDType {
	merged := &SDDType{
		Name:            child.Name,
		DisplayName:     child.DisplayName,
		Description:     child.Description,
		SpecDirectory:   child.SpecDirectory,
		Body:            child.Body,
		FileDefinitions: make(map[string]*FileDefinition),
	}

	// Use parent's spec directory if child doesn't specify
	if merged.SpecDirectory == "" {
		merged.SpecDirectory = parent.SpecDirectory
	}

	// Merge files (child additions/overrides)
	fileMap := make(map[string]FileSpec)
	for _, f := range parent.Files {
		fileMap[f.Name] = f
	}
	for _, f := range child.Files {
		fileMap[f.Name] = f
	}
	for _, f := range fileMap {
		merged.Files = append(merged.Files, f)
	}

	// Merge file definitions
	for name, def := range parent.FileDefinitions {
		merged.FileDefinitions[name] = def
	}
	for name, def := range child.FileDefinitions {
		merged.FileDefinitions[name] = def
	}

	return merged
}

// LoadFromPath loads an SDD type from a specific path.
func (r *Resolver) LoadFromPath(path string) (*SDDType, error) {
	dir := filepath.Dir(path)
	base := filepath.Base(path)

	sddType, err := LoadSDDType(os.DirFS(dir), base)
	if err != nil {
		return nil, err
	}

	return r.resolveInheritance(sddType, nil)
}
