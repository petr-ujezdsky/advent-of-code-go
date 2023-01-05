package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	source := make(map[string]int)

	source["aaa"] = 5
	source["bbb"] = -10

	copied := Copy(source)

	delete(source, "aaa")

	_, ok := source["aaa"]
	assert.False(t, ok)

	assert.Equal(t, 5, copied["aaa"])
}
