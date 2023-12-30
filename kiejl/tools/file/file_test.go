package file

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kiejl/kiejl/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	dire := test.Dire(t)
	path := filepath.Join(dire, "test.extn")

	// success
	err := Create(path)
	assert.FileExists(t, path)
	assert.NoError(t, err)

	// failure - file exists
	err = Create(path)
	test.AssertErr(t, err, "cannot create .*: file exists")

	// failure - other error
	err = Create("/root")
	test.AssertErr(t, err, "cannot create .*: .*")
}

func TestDelete(t *testing.T) {
	// setup
	path := test.File(t, "alpha.extn")
	dest := strings.Replace(path, ".extn", ".deleted", 1)

	// success
	err := Delete(path)
	assert.NoFileExists(t, path)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// failure - file does not exist
	err = Delete(path)
	test.AssertErr(t, err, "cannot delete .*: file does not exist")

	// failure - other error
	err = Delete("/")
	test.AssertErr(t, err, "cannot delete .*: .*")
}

func TestExists(t *testing.T) {
	// setup
	path := test.File(t, "alpha.extn")

	// success - true
	ok := Exists(path)
	assert.True(t, ok)

	// success - false
	ok = Exists("/nope")
	assert.False(t, ok)
}

func TestRead(t *testing.T) {
	// setup
	path := test.File(t, "alpha.extn")

	// success
	body, err := Read(path)
	assert.Equal(t, "Alpha.\n", body)
	assert.NoError(t, err)

	// failure - file does not exist
	body, err = Read("/nope")
	assert.Empty(t, body)
	test.AssertErr(t, err, "cannot read .*: file does not exist")

	// failure - other error
	body, err = Read("/")
	assert.Empty(t, body)
	test.AssertErr(t, err, "cannot read .*: .*")
}

func TestRename(t *testing.T) {
	// setup
	path := test.File(t, "alpha.extn")
	dire := filepath.Dir(path)
	dest := filepath.Join(dire, "test.extn")

	// success
	err := Rename(path, "test")
	assert.NoFileExists(t, path)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// failure - file does not exist
	err = Rename(path, "test")
	test.AssertErr(t, err, "cannot rename .*: file does not exist")

	// failure - other error
	err = Rename("/", "test")
	test.AssertErr(t, err, "cannot rename .*: .*")
}

func TestSearch(t *testing.T) {
	// setup
	path := test.File(t, "alpha.extn")

	// success - true
	ok, err := Search(path, "alph")
	assert.True(t, ok)
	assert.NoError(t, err)

	// success - false
	ok, err = Search(path, "nope")
	assert.False(t, ok)
	assert.NoError(t, err)

	// failure - file does not exist
	ok, err = Search("/nope", "nope")
	assert.False(t, ok)
	test.AssertErr(t, err, "cannot search .*: file does not exist")

	// failure - other error
	ok, err = Search("/", "nope")
	assert.False(t, ok)
	test.AssertErr(t, err, "cannot search .*: .*")
}

func TestUpdate(t *testing.T) {
	// setup
	path := test.File(t, "alpha.extn")

	// success
	err := Update(path, "test\n", 0666)
	test.AssertFile(t, path, "test\n")
	assert.NoError(t, err)

	// failure - file does not exist
	err = Update("/nope", "test\n", 0666)
	test.AssertErr(t, err, "cannot update .*: file does not exist")

	// failure - other error
	err = Update("/", "test\n", 0666)
	test.AssertErr(t, err, "cannot update .*: .*")
}
