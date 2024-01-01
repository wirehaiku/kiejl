package tasks

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kiejl/kiejl/items/book"
	"github.com/wirehaiku/kiejl/kiejl/tools/test"
)

func TestGet(t *testing.T) {
	// setup
	Tasks["test"] = func(*book.Book, *bytes.Buffer, []string) error {
		return nil
	}

	// success
	task, err := Get("test")
	assert.NotNil(t, task)
	assert.NoError(t, err)

	// failure - task does not exist
	task, err = Get("nope")
	assert.Nil(t, task)
	test.AssertErr(t, err, "task .* does not exist")
}
