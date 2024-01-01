// Package book implements the Book type and methods.
package book

import (
	"os"

	"github.com/wirehaiku/kiejl/kiejl/items/note"
	"github.com/wirehaiku/kiejl/kiejl/tools/file"
	"github.com/wirehaiku/kiejl/kiejl/tools/neat"
	"github.com/wirehaiku/kiejl/kiejl/tools/path"
)

// Book is a single directory full of Notes.
type Book struct {
	Dire string
	Extn string
	Mode os.FileMode
}

// New returns a new Book.
func New(dire, extn string, mode os.FileMode) *Book {
	dire = neat.Path(dire)
	extn = neat.Extn(extn)
	return &Book{dire, extn, mode}
}

// Create creates and returns a new Note in the Book.
func (b *Book) Create(name string) (*note.Note, error) {
	name = neat.Name(name)
	path := path.Join(b.Dire, name, b.Extn)
	if err := file.Create(path); err != nil {
		return nil, err
	}

	return note.New(path, b.Mode), nil
}

// Filter returns all Notes in the Book passing a filter function.
func (b *Book) Filter(filt func(*note.Note) (bool, error)) ([]*note.Note, error) {
	var notes []*note.Note
	for _, note := range b.List() {
		ok, err := filt(note)
		switch {
		case err != nil:
			return nil, err
		case ok:
			notes = append(notes, note)
		}
	}

	return notes, nil
}

// Get returns an existing Note from the Book, or nil.
func (b *Book) Get(name string) *note.Note {
	name = neat.Name(name)
	for _, note := range b.List() {
		if note.Name() == name {
			return note
		}
	}

	return nil
}

// GetOrCreate returns an existing or newly created Note from the Book.
func (b *Book) GetOrCreate(name string) (*note.Note, error) {
	if note := b.Get(name); note != nil {
		return note, nil
	}

	return b.Create(name)
}

// List returns all Notes in the Book in alphabetical order.
func (b *Book) List() []*note.Note {
	var notes []*note.Note
	for _, path := range path.Glob(b.Dire, b.Extn) {
		notes = append(notes, note.New(path, b.Mode))
	}

	return notes
}

// Match returns all Notes with names containing a substring.
func (b *Book) Match(text string) []*note.Note {
	notes, _ := b.Filter(func(note *note.Note) (bool, error) {
		return note.Match(text), nil
	})

	return notes
}

// Search returns all Notes with bodies containing a substring.
func (b *Book) Search(text string) ([]*note.Note, error) {
	return b.Filter(func(note *note.Note) (bool, error) {
		return note.Search(text)
	})
}
