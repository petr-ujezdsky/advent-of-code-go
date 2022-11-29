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

func TestSignum(t *testing.T) {
	assert.Equal(t, -1, utils.Signum(-600))
	assert.Equal(t, -1, utils.Signum(-1))
	assert.Equal(t, 0, utils.Signum(0))
	assert.Equal(t, 1, utils.Signum(1))
	assert.Equal(t, 1, utils.Signum(20))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 1, utils.Max(0, 1))
	assert.Equal(t, 5, utils.Max(5, -1))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 0, utils.Min(0, 1))
	assert.Equal(t, -1, utils.Min(5, -1))
}

func TestArgMin(t *testing.T) {
	index, min := utils.ArgMin(5, 4, 3)
	assert.Equal(t, 2, index)
	assert.Equal(t, 3, min)

	index, min = utils.ArgMin([]int{5, -300, 80, 500}...)
	assert.Equal(t, 1, index)
	assert.Equal(t, -300, min)
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

func TestShallowCopy(t *testing.T) {
	data := []int{1, 2, 3}
	clone := utils.ShallowCopy(data)

	// modify original data
	data[0] = 9

	assert.Equal(t, 9, data[0])
	assert.Equal(t, 1, clone[0])
}

func TestReverse(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, utils.Reverse([]int{1, 2, 3}))
	assert.Equal(t, []int{4, 3, 2, 1}, utils.Reverse([]int{1, 2, 3, 4}))
}

func TestRemoveUnordered(t *testing.T) {
	assert.Equal(t, []int{3, 2}, utils.RemoveUnordered([]int{1, 2, 3}, 0))
	assert.Equal(t, []int{1, 3}, utils.RemoveUnordered([]int{1, 2, 3}, 1))
	assert.Equal(t, []int{1, 2}, utils.RemoveUnordered([]int{1, 2, 3}, 2))
}
