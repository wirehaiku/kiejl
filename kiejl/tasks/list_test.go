package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	// success - no arguments
	buff, err := run(t, List, nil)
	assert.Equal(t, "alpha\nbravo\n", buff)
	assert.NoError(t, err)

	// success - with arguments
	buff, err = run(t, List, []string{"alph"})
	assert.Equal(t, "alpha\n", buff)
	assert.NoError(t, err)
}
