package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_01_parse(t *testing.T) {
	world := ParseInput("389125467")

	assert.Equal(t, 9, len(world.CupsByLabel))
	cup := world.CupsByLabel[1]
	assert.Equal(t, 1, cup.Label)
	assert.Equal(t, "1 2 5 4 6 7 3 8 9", cup.String())
	assert.Equal(t, "(1) 2 5 4 6 7 3 8 9", cup.StringHighlighted(1))
}

func Test_01_example_small(t *testing.T) {
	world := ParseInput("389125467")

	result := DoWithInputPart01(world, 10)
	assert.Equal(t, "92658374", result)
}

func Test_01_example_big(t *testing.T) {
	world := ParseInput("389125467")

	result := DoWithInputPart01(world, 100)
	assert.Equal(t, "67384529", result)
}

func Test_01(t *testing.T) {
	world := ParseInput("643719258")

	result := DoWithInputPart01(world, 100)
	assert.Equal(t, "54896723", result)
}

func Test_02_example(t *testing.T) {
	world := ParseInput("")

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	world := ParseInput("")

	result := DoWithInputPart02(world)
	assert.Equal(t, 0, result)
}
