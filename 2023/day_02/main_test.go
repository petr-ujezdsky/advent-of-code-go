package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 5, len(world.Games))

	game := world.Games[0]
	assert.Equal(t, 1, game.Id)
	assert.Equal(t, CubeSet{Red: 4, Blue: 3}, game.Examples[0])
	assert.Equal(t, CubeSet{Red: 1, Green: 2, Blue: 6}, game.Examples[1])
	assert.Equal(t, CubeSet{Green: 2}, game.Examples[2])

	game = world.Games[2]
	assert.Equal(t, 3, game.Id)
	assert.Equal(t, CubeSet{Red: 20, Green: 8, Blue: 6}, game.Examples[0])
	assert.Equal(t, CubeSet{Red: 4, Green: 13, Blue: 5}, game.Examples[1])
	assert.Equal(t, CubeSet{Red: 1, Green: 5, Blue: 0}, game.Examples[2])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 8, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 2265, result)
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
