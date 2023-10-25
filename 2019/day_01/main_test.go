package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 100, len(world.Weights))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 3325156, result)
}

func Test_02_example(t *testing.T) {
	assert.Equal(t, 50346, calculateFuel(100756))
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 4984866, result)
}
