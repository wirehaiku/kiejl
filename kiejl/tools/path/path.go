// Package path implements filepath manipulation functions.
package path

import (
	"path/filepath"
	"sort"
	"strings"
)

// Dire returns a path's parent directory.
func Dire(path string) string {
	return filepath.Dir(path)
}

// Extn returns a path's file extension with a leading dot.
func Extn(path string) string {
	return filepath.Ext(path)
}

// Glob returns all paths in a directory matching an extension, sorted alphabetically.
func Glob(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	paths, _ := filepath.Glob(glob)
	sort.Strings(paths)
	return paths
}

// Join returns a path from directory, name and extension strings.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Name returns a path's base name without the extension.
func Name(path string) string {
	base := filepath.Base(path)
	extn := filepath.Ext(base)
	return strings.TrimSuffix(base, extn)
}
