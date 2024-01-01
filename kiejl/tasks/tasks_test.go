package tasks

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kiejl/kiejl/items/book"
	"github.com/wirehaiku/kiejl/kiejl/tools/test"
)

func run(t *testing.T, task Task, elems []string) (string, error) {
	dire := test.Dire(t)
	book := book.New(dire, ".extn", 0666)
	buff := bytes.NewBuffer(nil)
	err := task(book, buff, elems)
	return buff.String(), err
}

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
