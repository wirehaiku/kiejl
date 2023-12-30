// Package test implements unit-testing helper functions.
package test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockFiles is a map of mock files for unit-testing.
var MockFiles = map[string]string{
	"alpha.extn": "Alpha.\n",
	"bravo.extn": "Bravo.\n",
}

// AssertErr asserts an error's message matches a regular expression.
func AssertErr(t *testing.T, err error, regx string) {
	regx = fmt.Sprintf("^%s$", regx)
	assert.Regexp(t, regx, err.Error())
}

// AssertFile asserts the contents of a file from a string.
func AssertFile(t *testing.T, path, body string) {
	bytes, err := os.ReadFile(path)
	assert.Equal(t, body, string(bytes))
	assert.NoError(t, err)
}

// Dire returns a temporary directory populated from MockFiles.
func Dire(t *testing.T) string {
	dire := t.TempDir()
	for name, body := range MockFiles {
		path := filepath.Join(dire, name)
		os.WriteFile(path, []byte(body), 0666)
	}

	return dire
}

// File returns a temporary file populated from an entry in MockFiles.
func File(t *testing.T, name string) string {
	dire := t.TempDir()
	path := filepath.Join(dire, name)
	os.WriteFile(path, []byte(MockFiles[name]), 0666)
	return path
}
