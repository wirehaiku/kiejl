package neat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	// success
	body := Body("\tbody\n")
	assert.Equal(t, "body\n", body)
}

func TestExtn(t *testing.T) {
	// success
	extn := Extn("\tEXTN\n")
	assert.Equal(t, ".extn", extn)
}

func TestName(t *testing.T) {
	// success
	name := Name("\tNAME\n")
	assert.Equal(t, "name", name)
}

func TestPath(t *testing.T) {
	// success
	dire := Path("\t/dire\n")
	assert.Equal(t, "/dire", dire)
}
