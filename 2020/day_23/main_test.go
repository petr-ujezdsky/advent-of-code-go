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
	assert.Equal(t, "125467389", cup.String())
}

func Test_01_example(t *testing.T) {
	world := ParseInput("389125467")

	result := DoWithInputPart01(world)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	world := ParseInput("643719258")

	result := DoWithInputPart01(world)
	assert.Equal(t, 0, result)
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
