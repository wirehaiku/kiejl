// Package clui implements command-line user interface functions.
package clui

import "fmt"

// Check returns an error if an argument slice is less than an integer.
func Check(elems []string, size int) error {
	if len(elems) < size {
		return fmt.Errorf("not enough arguments")
	}

	return nil
}

// Parse returns a Task name and argument slice from an argument slice.
func Parse(elems []string) (string, []string) {
	switch len(elems) {
	case 0:
		return "", nil
	case 1:
		return elems[0], nil
	default:
		return elems[0], elems[1:]
	}
}
