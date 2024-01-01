package path

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wirehaiku/kiejl/kiejl/tools/test"
)

func TestDire(t *testing.T) {
	// success
	dire := Dire("/dire/name.extn")
	assert.Equal(t, "/dire", dire)
}

func TestEvar(t *testing.T) {
	// setup
	os.Setenv("TEST", "test\n")
	os.Setenv("BLANK", "\n")

	// success
	evar, err := Evar("test")
	assert.Equal(t, "test", evar)
	assert.NoError(t, err)

	// failure - variable does not exist
	evar, err = Evar("nope")
	assert.Empty(t, evar)
	test.AssertErr(t, err, "environment variable .* does not exist")

	// failure - variable is blank
	evar, err = Evar("blank")
	assert.Empty(t, evar)
	test.AssertErr(t, err, "environment variable .* is blank")
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn("/dire/name.extn")
	assert.Equal(t, ".extn", extn)
}

func TestGlob(t *testing.T) {
	// setup
	dire := test.Dire(t)

	// success
	paths := Glob(dire, ".extn")
	assert.Equal(t, []string{
		filepath.Join(dire, "alpha.extn"),
		filepath.Join(dire, "bravo.extn"),
	}, paths)
}

func TestJoin(t *testing.T) {
	// success
	path := Join("/dire", "name", ".extn")
	assert.Equal(t, "/dire/name.extn", path)
}

func TestMatch(t *testing.T) {
	// success - true
	ok := Match("/dire/name.extn", "NAME")
	assert.True(t, ok)

	// success - false
	ok = Match("/dire/name.extn", "NOPE")
	assert.False(t, ok)
}

func TestName(t *testing.T) {
	// success
	name := Name("/dire/name.extn")
	assert.Equal(t, "name", name)
}
