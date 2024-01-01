// Package path implements filepath manipulation functions.
package path

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Dire returns a path's parent directory.
func Dire(path string) string {
	return filepath.Dir(path)
}

// Evar returns an environment variable by name, or an error.
func Evar(name string) (string, error) {
	name = strings.ToUpper(name)
	data, ok := os.LookupEnv(name)
	data = strings.TrimSpace(data)

	switch {
	case !ok:
		return "", fmt.Errorf("environment variable %q does not exist", name)
	case data == "":
		return "", fmt.Errorf("environment variable %q is blank", name)
	default:
		return data, nil
	}
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

// Match returns true if a path's base name contains a substring.
func Match(path, text string) bool {
	base := filepath.Base(path)
	extn := filepath.Ext(base)
	name := strings.TrimSuffix(base, extn)
	name = strings.ToLower(name)
	text = strings.ToLower(text)
	return strings.Contains(name, text)
}

// Name returns a path's base name without the extension.
func Name(path string) string {
	base := filepath.Base(path)
	extn := filepath.Ext(base)
	return strings.TrimSuffix(base, extn)
}
