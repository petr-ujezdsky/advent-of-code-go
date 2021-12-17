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

func TestAbs(t *testing.T) {
	assert.Equal(t, 5, utils.Abs(5))
	assert.Equal(t, 5, utils.Abs(-5))
	assert.Equal(t, 0, utils.Abs(0))
	// does not work - overflows
	// assert.Equal(t, -math.MinInt, utils.Abs(math.MinInt))
}

func TestSumNtoM(t *testing.T) {
	assert.Equal(t, 1, utils.SumNtoM(0, 1))
	assert.Equal(t, 0, utils.SumNtoM(-50, 50))
	assert.Equal(t, 5050, utils.SumNtoM(0, 100))
	assert.Equal(t, 5050, utils.SumNtoM(1, 100))
	assert.Equal(t, -5050, utils.SumNtoM(-100, 0))
}

func TestClamp(t *testing.T) {
	assert.Equal(t, 0, utils.Clamp(-1, 0, 10))
	assert.Equal(t, 0, utils.Clamp(0, 0, 10))
	assert.Equal(t, 5, utils.Clamp(5, 0, 10))
	assert.Equal(t, 10, utils.Clamp(10, 0, 10))
	assert.Equal(t, 10, utils.Clamp(30, 0, 10))
}
