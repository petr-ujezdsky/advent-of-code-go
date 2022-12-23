package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader, 4)

	assert.Equal(t, 7, len(world.Steps))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader, 4)

	result := Walk(world)
	assert.Equal(t, 6032, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader, 50)

	result := Walk(world)
	assert.Equal(t, 186128, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader, 4)

	result := Walk3D(world, patchEdgesExample)
	assert.Equal(t, 5031, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader, 50)

	result := Walk3D(world, patchEdgesMain)
	assert.Equal(t, 0, result)
}
