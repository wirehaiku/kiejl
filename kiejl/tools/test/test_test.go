package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertFile(t *testing.T) {
	// setup
	dire := t.TempDir()
	path := filepath.Join(dire, "file.extn")
	os.WriteFile(path, []byte("test\n"), 0666)

	// success
	AssertFile(t, path, "test\n")
}

func TestDire(t *testing.T) {
	// success
	dire := Dire(t)
	for name, body := range MockFiles {
		path := filepath.Join(dire, name)
		bytes, err := os.ReadFile(path)
		assert.Equal(t, body, string(bytes))
		assert.NoError(t, err)
	}
}

func TestFile(t *testing.T) {
	// success
	path := File(t, "alpha.extn")
	bytes, err := os.ReadFile(path)
	assert.Equal(t, MockFiles["alpha.extn"], string(bytes))
	assert.NoError(t, err)
}
