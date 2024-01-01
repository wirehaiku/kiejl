// Package tasks implements the Task type and functions.
package tasks

import (
	"bytes"
	"fmt"

	"github.com/wirehaiku/kiejl/kiejl/items/book"
	"github.com/wirehaiku/kiejl/kiejl/tools/neat"
)

// Task is a high-level user function.
type Task func(*book.Book, *bytes.Buffer, []string) error

// Tasks is a map of all existing Tasks.
var Tasks = map[string]Task{}

// Get returns an existing Task or an error.
func Get(name string) (Task, error) {
	name = neat.Name(name)
	task, ok := Tasks[name]
	if !ok {
		return nil, fmt.Errorf("task %q does not exist", name)
	}

	return task, nil
}
