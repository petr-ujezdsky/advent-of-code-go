package day_09

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	heightMap, err := ParseInput(reader)
	assert.Nil(t, err)

	// first column
	assert.Equal(t, []int{2, 3, 9, 8, 9}, heightMap.Columns[0])

	// last column
	assert.Equal(t, []int{0, 1, 2, 9, 8}, heightMap.Columns[9])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	heightMap, err := ParseInput(reader)
	assert.Nil(t, err)

	sum := FindLowPointsAndSum(heightMap)
	assert.Equal(t, 15, sum)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	heightMap, err := ParseInput(reader)
	assert.Nil(t, err)

	sum := FindLowPointsAndSum(heightMap)
	assert.Equal(t, 486, sum)
}
