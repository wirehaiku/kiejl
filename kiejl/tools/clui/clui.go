// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"os"
	"strings"
)

// Check returns an error if an argument slice is less than an integer.
func Check(elems []string, size int) error {
	if len(elems) < size {
		return fmt.Errorf("not enough arguments")
	}

	return nil
}

// Default returns an element from an argument slice, or a default string.
func Default(elems []string, indx int, dflt string) string {
	if len(elems) > indx {
		return elems[indx]
	}

	return dflt
}

// Evar returns an environment variable by name, or an error.
func Evar(name string) (string, error) {
	name = strings.ToUpper(name)
	data, ok := os.LookupEnv(name)
	data = strings.TrimSpace(data)

	switch {
	case !ok:
		return "", fmt.Errorf("environment variable %q does not exist", name)
	case data == "":
		return "", fmt.Errorf("environment variable %q is blank", name)
	default:
		return data, nil
	}
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
