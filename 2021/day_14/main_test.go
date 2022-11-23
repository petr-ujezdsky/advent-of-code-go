package day_14

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, "NNCB", world.template)
	assert.Equal(t, 16, len(world.rules))
	assert.Equal(t, "B", world.rules["CH"])
	assert.Equal(t, "C", world.rules["CN"])
}

func Test_01_example_single(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	polymer := world.template
	polymer = GrowPolymerStepIter(polymer, world.rules)
	assert.Equal(t, "NCNBCHB", polymer)

	polymer = GrowPolymerStepIter(polymer, world.rules)
	assert.Equal(t, "NBCCNBBBCBHCB", polymer)

	polymer = GrowPolymerStepIter(polymer, world.rules)
	assert.Equal(t, "NBBBCNCCNBBNBNBBCHBHHBCHB", polymer)

	polymer = GrowPolymerStepIter(polymer, world.rules)
	assert.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", polymer)
}

func Test_01_example_multi(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	polymer := GrowPolymerIter(world.template, world.rules, 4)
	assert.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", polymer)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	polymer := GrowPolymerIter(world.template, world.rules, 10)
	score := PolymerScore(polymer)
	assert.Equal(t, 1588, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	polymer := GrowPolymerIter(world.template, world.rules, 10)
	score := PolymerScore(polymer)
	assert.Equal(t, 3555, score)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	// never finishes
	polymer := GrowPolymerIter(world.template, world.rules, 40)
	score := PolymerScore(polymer)
	assert.Equal(t, -1, score)
}
