package tasks

import (
	"bytes"
	"fmt"

	"github.com/wirehaiku/kiejl/kiejl/items/book"
	"github.com/wirehaiku/kiejl/kiejl/tools/clui"
)

// List is a Task that lists existing Notes.
func List(book *book.Book, buff *bytes.Buffer, elems []string) error {
	text := clui.Default(elems, 0, "")

	for _, note := range book.Match(text) {
		fmt.Fprintf(buff, "%s\n", note.Name())
	}

	return nil
}
