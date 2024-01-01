// Package main implements the main program function.
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/wirehaiku/kiejl/kiejl/items/book"
	"github.com/wirehaiku/kiejl/kiejl/tasks"
	"github.com/wirehaiku/kiejl/kiejl/tools/clui"
	"github.com/wirehaiku/kiejl/kiejl/tools/path"
)

// try prints a non-nil error and exits.
func try(err error) {
	if err != nil {
		fmt.Printf("Error: %s.\n", err.Error())
		os.Exit(1)
	}
}

// main runs the main Kiejl program.
func main() {
	// Collect configuration variables.
	dire, err := path.Evar("KIEJL_DIR")
	try(err)

	extn, err := path.Evar("KIEJL_EXT")
	try(err)

	// Collect and parse arguments and Task.
	name, elems := clui.Parse(os.Args[1:])
	task, err := tasks.Get(name)
	try(err)

	// Initialise Book object and Buffer.
	book := book.New(dire, extn, 0666)
	buff := bytes.NewBuffer(nil)

	// Run the collected Task against the Book and Buffer.
	err = task(book, buff, elems)
	try(err)

	// Print the buffered output.
	buff.WriteTo(os.Stdout)
}
