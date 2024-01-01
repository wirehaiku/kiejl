// Package note implements the Note type and methods.
package note

import (
	"os"

	"github.com/wirehaiku/kiejl/kiejl/tools/file"
	"github.com/wirehaiku/kiejl/kiejl/tools/neat"
	"github.com/wirehaiku/kiejl/kiejl/tools/path"
)

// Note is a single plaintext note file.
type Note struct {
	Path string
	Mode os.FileMode
}

// New returns a new Note.
func New(path string, mode os.FileMode) *Note {
	path = neat.Path(path)
	return &Note{path, mode}
}

// Delete renames the Note to a ".deleted" extension in the same directory.
func (n *Note) Delete() error {
	return file.Delete(n.Path)
}

// Exists returns true if the Note's file exists.
func (n *Note) Exists() bool {
	return file.Exists(n.Path)
}

// Match returns true if the Note's name contains a substring.
func (n *Note) Match(text string) bool {
	return path.Match(n.Path, text)
}

// Name returns the Note's name.
func (n *Note) Name() string {
	name := path.Name(n.Path)
	return neat.Name(name)
}

// Read returns the Note's body as a string.
func (n *Note) Read() (string, error) {
	return file.Read(n.Path)
}

// Rename moves the Note's file to a new name in the same directory.
func (n *Note) Rename(name string) error {
	name = neat.Name(name)
	return file.Rename(n.Path, name)
}

// Search returns true if the Note's body contains a substring.
func (n *Note) Search(text string) (bool, error) {
	return file.Search(n.Path, text)
}

// Update overwrites the Note's body with a string.
func (n *Note) Update(body string) error {
	body = neat.Body(body)
	return file.Update(n.Path, body, n.Mode)
}
