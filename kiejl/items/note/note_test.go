package note

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kiejl/kiejl/tools/test"
)

func tNote(t *testing.T) *Note {
	path := test.File(t, "alpha.extn")
	return New(path, 0666)
}

func TestNew(t *testing.T) {
	// setup
	note := tNote(t)

	// success
	assert.Contains(t, note.Path, "alpha.extn")
	assert.Equal(t, os.FileMode(0666), note.Mode)
}

func TestDelete(t *testing.T) {
	// setup
	note := tNote(t)
	dest := strings.Replace(note.Path, ".extn", ".deleted", 1)

	// success
	err := note.Delete()
	assert.NoFileExists(t, note.Path)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	note := tNote(t)

	// success
	ok := note.Exists()
	assert.True(t, ok)
}

func TestMatch(t *testing.T) {
	// setup
	note := tNote(t)

	// success
	ok := note.Match("ALPH")
	assert.True(t, ok)
}

func TestName(t *testing.T) {
	// setup
	note := tNote(t)

	// success
	name := note.Name()
	assert.Equal(t, "alpha", name)
}

func TestRead(t *testing.T) {
	// setup
	note := tNote(t)

	// success
	body, err := note.Read()
	assert.Equal(t, "Alpha.\n", body)
	assert.NoError(t, err)
}

func TestRename(t *testing.T) {
	// setup
	note := tNote(t)
	dest := strings.Replace(note.Path, "alpha", "test", 1)

	// success
	err := note.Rename("test")
	assert.NoFileExists(t, note.Path)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	note := tNote(t)

	// success
	ok, err := note.Search("ALPH")
	assert.True(t, ok)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// setup
	note := tNote(t)

	// success
	err := note.Update("test\n")
	test.AssertFile(t, note.Path, "test\n")
	assert.NoError(t, err)
}
