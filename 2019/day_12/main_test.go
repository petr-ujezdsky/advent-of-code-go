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

	assert.Equal(t, 4, len(world.Moons))
	assert.Equal(t, utils.Vector3i{X: 2, Y: -10, Z: -7}, world.Moons[1].Position)
}

func Test_01_example1(t *testing.T) {
	reader, err := os.Open("data-00-example1.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world, 10)
	assert.Equal(t, 179, result)
}

func Test_01_example2(t *testing.T) {
	reader, err := os.Open("data-00-example2.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world, 100)
	assert.Equal(t, 1940, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world, 1000)
	assert.Equal(t, 5517, result)
}

func Test_02_example1(t *testing.T) {
	reader, err := os.Open("data-00-example1.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 2772, result)
}

func Test_02_example2(t *testing.T) {
	reader, err := os.Open("data-00-example2.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 4686774924, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 303070460651184, result)
}
