package utils_test

import (
	"os"
	"testing"

	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseToInts(t *testing.T) {
	reader, err := os.Open("test-data/parse-to-ints.txt")
	assert.Nil(t, err)

	parsed, err := utils.ParseToInts(reader)
	assert.Nil(t, err)

	expected := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263}

	assert.Equal(t, expected, parsed)
}
