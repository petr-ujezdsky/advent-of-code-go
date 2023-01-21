package day_11

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	energyLevels := ParseInput(reader)

	// first column
	assert.Equal(t, []int{5, 2, 5, 6, 6, 4, 2, 6, 4, 5}, energyLevels.Columns[0])

	// last column
	assert.Equal(t, []int{3, 1, 3, 6, 8, 5, 1, 4, 4, 6}, energyLevels.Columns[9])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	energyLevels := ParseInput(reader)

	flashesCount, allFlashedStepNumber := CountFlashes(energyLevels, 100)

	assert.Equal(t, 1656, flashesCount)
	assert.Equal(t, 195, allFlashedStepNumber)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	energyLevels := ParseInput(reader)

	flashesCount, allFlashedStepNumber := CountFlashes(energyLevels, 100)

	assert.Equal(t, 1773, flashesCount)
	assert.Equal(t, 494, allFlashedStepNumber)
}
