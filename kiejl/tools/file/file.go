// Package file implements filesystem handling functions.
package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Create creates a new empty file.
func Create(path string) error {
	if Exists(path) {
		return fmt.Errorf("cannot create %q: file exists", path)
	}

	if _, err := os.Create(path); err != nil {
		return fmt.Errorf("cannot create %q: %w", path, err)
	}

	return nil
}

// Delete renames a file to a ".deleted" extension in the same directory.
func Delete(path string) error {
	dire := filepath.Dir(path)
	base := filepath.Base(path)
	extn := filepath.Ext(base)
	name := strings.TrimSuffix(base, extn)
	dest := filepath.Join(dire, name+".deleted")

	if !Exists(path) {
		return fmt.Errorf("cannot delete %q: file does not exist", path)
	}

	if err := os.Rename(path, dest); err != nil {
		return fmt.Errorf("cannot delete %q: %w", path, err)
	}

	return nil
}

// Exists returns true if a file or directory exists.
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

// Read returns a file's body as a string.
func Read(path string) (string, error) {
	if !Exists(path) {
		return "", fmt.Errorf("cannot read %q: file does not exist", path)
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("cannot read %q: %w", path, err)
	}

	return string(bytes), nil
}

// Rename moves a file to a new name in the same directory.
func Rename(path, name string) error {
	dire := filepath.Dir(path)
	extn := filepath.Ext(path)
	dest := filepath.Join(dire, name+extn)

	if !Exists(path) {
		return fmt.Errorf("cannot rename %q: file does not exist", path)
	}

	if err := os.Rename(path, dest); err != nil {
		return fmt.Errorf("cannot rename %q: %w", path, err)
	}

	return nil
}

// Search returns true if a file's body contains a substring.
func Search(path, text string) (bool, error) {
	if !Exists(path) {
		return false, fmt.Errorf("cannot search %q: file does not exist", path)
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("cannot search %q: %w", path, err)
	}

	body := strings.ToLower(string(bytes))
	text = strings.ToLower(text)
	return strings.Contains(body, text), nil
}

// Update overwrites an existing file's body with a string.
func Update(path, body string, mode os.FileMode) error {
	if !Exists(path) {
		return fmt.Errorf("cannot update %q: file does not exist", path)
	}

	if err := os.WriteFile(path, []byte(body), mode); err != nil {
		return fmt.Errorf("cannot update %q: %w", path, err)
	}

	return nil
}
