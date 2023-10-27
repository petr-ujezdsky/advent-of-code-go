package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 9, len(world.Wire1.Steps))
	assert.Equal(t, 10, len(world.Wire1.Points))
	assert.Equal(t, 8, len(world.Wire2.Steps))
	assert.Equal(t, 9, len(world.Wire2.Points))
}

func Test_01_example_01(t *testing.T) {
	reader, err := os.Open("data-00-example-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 159, result)
}

func Test_01_example_02(t *testing.T) {
	reader, err := os.Open("data-00-example-02.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 135, result)
}

func Test_01_example_03(t *testing.T) {
	reader, err := os.Open("data-00-example-03.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 6, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 217, result)
}

func Test_02_example_01(t *testing.T) {
	reader, err := os.Open("data-00-example-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 610, result)
}

func Test_02_example_02(t *testing.T) {
	reader, err := os.Open("data-00-example-02.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 410, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 2450, result)
}
