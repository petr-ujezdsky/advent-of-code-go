package main

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 5, len(world.Lines))
	assert.Equal(t, utils.Vector3i{X: 20, Y: 19, Z: 15}, world.Lines[4].Position)
	assert.Equal(t, utils.Vector3i{X: 1, Y: -5, Z: -3}, world.Lines[4].Direction)
}

func Test_findCrossing2D(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	line1 := world.Lines[1]
	line2 := world.Lines[3]

	cross, ok := findCrossing2D(line1, line2)

	assert.True(t, ok)
	assert.Equal(t, utils.Vector2f{X: -6, Y: -5}, cross)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world, 7, 27)
	assert.Equal(t, 2, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world, 200000000000000, 400000000000000)
	assert.Equal(t, 18651, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}