package book

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kiejl/kiejl/items/note"
	"github.com/wirehaiku/kiejl/kiejl/tools/test"
)

func p(book *Book, base string) string {
	return filepath.Join(book.Dire, base)
}

func tBook(t *testing.T) *Book {
	dire := test.Dire(t)
	return New(dire, ".extn", 0666)
}

func TestNew(t *testing.T) {
	// success
	book := tBook(t)
	assert.NotEmpty(t, book.Dire)
	assert.Equal(t, ".extn", book.Extn)
	assert.Equal(t, os.FileMode(0666), book.Mode)
}

func TestCreate(t *testing.T) {
	// setup
	book := tBook(t)

	// success
	note, err := book.Create("test")
	assert.Equal(t, p(book, "test.extn"), note.Path)
	assert.FileExists(t, note.Path)
	assert.NoError(t, err)
}

func TestFilter(t *testing.T) {
	// setup
	book := tBook(t)

	// success
	notes, err := book.Filter(func(note *note.Note) (bool, error) {
		return note.Name() == "alpha", nil
	})
	assert.Len(t, notes, 1)
	assert.Equal(t, p(book, "alpha.extn"), notes[0].Path)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	book := tBook(t)

	// success - note exists
	note := book.Get("alpha")
	assert.Equal(t, p(book, "alpha.extn"), note.Path)

	// success - note does not exist
	note = book.Get("nope")
	assert.Nil(t, note)
}

func TestGetOrCreate(t *testing.T) {
	// setup
	book := tBook(t)

	// success - note exists
	note, err := book.GetOrCreate("alpha")
	assert.Equal(t, p(book, "alpha.extn"), note.Path)
	assert.NoError(t, err)

	// success - note does not exist
	note, err = book.GetOrCreate("test")
	assert.Equal(t, p(book, "test.extn"), note.Path)
	assert.FileExists(t, note.Path)
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	// setup
	book := tBook(t)

	// success
	notes := book.List()
	assert.Len(t, notes, 2)
	assert.Equal(t, p(book, "alpha.extn"), notes[0].Path)
	assert.Equal(t, p(book, "bravo.extn"), notes[1].Path)
}

func TestMatch(t *testing.T) {
	// setup
	book := tBook(t)

	// success
	notes := book.Match("alph")
	assert.Len(t, notes, 1)
	assert.Equal(t, p(book, "alpha.extn"), notes[0].Path)
}

func TestSearch(t *testing.T) {
	// setup
	book := tBook(t)

	// success
	notes, err := book.Search("alph")
	assert.Len(t, notes, 1)
	assert.Equal(t, p(book, "alpha.extn"), notes[0].Path)
	assert.NoError(t, err)
}
