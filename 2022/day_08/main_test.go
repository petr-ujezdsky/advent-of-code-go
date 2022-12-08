package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	heights := ParseInput(reader)

	result := CountVisibleTrees(heights)
	assert.Equal(t, 21, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	heights := ParseInput(reader)

	result := CountVisibleTrees(heights)
	assert.Equal(t, 1796, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	heights := ParseInput(reader)

	result := FindBestTreeHouseLocationScore(heights)
	assert.Equal(t, 8, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	heights := ParseInput(reader)

	result := FindBestTreeHouseLocationScore(heights)
	assert.Equal(t, 288120, result)
}
