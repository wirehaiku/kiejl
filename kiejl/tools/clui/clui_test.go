package clui

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kiejl/kiejl/tools/test"
)

func TestCheck(t *testing.T) {
	// success
	err := Check([]string{"one"}, 1)
	assert.NoError(t, err)

	// failure - not enough arguments
	err = Check(nil, 1)
	test.AssertErr(t, err, "not enough arguments")
}

func TestDefault(t *testing.T) {
	// success - given argument
	elem := Default([]string{"one"}, 0, "default")
	assert.Equal(t, "one", elem)

	// success - default argument
	elem = Default([]string{"one"}, 1, "default")
	assert.Equal(t, "default", elem)
}

func TestParse(t *testing.T) {
	// success - no arguments
	name, elems := Parse(nil)
	assert.Empty(t, name)
	assert.Empty(t, elems)

	// success - one argument
	name, elems = Parse([]string{"one"})
	assert.Equal(t, "one", name)
	assert.Empty(t, elems)

	// success - multiple arguments
	name, elems = Parse([]string{"one", "two"})
	assert.Equal(t, "one", name)
	assert.Equal(t, []string{"two"}, elems)
}
