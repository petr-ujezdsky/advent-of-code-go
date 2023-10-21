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

	assert.Equal(t, 4, len(world.Foods))
	assert.Equal(t, 7, len(world.AllIngredients))
	assert.Equal(t, 3, len(world.AllAllergens))

	food := world.Foods[0]
	assert.Equal(t, 4, len(food.Ingredients))
	assert.Equal(t, 2, len(food.Allergens))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	part1, _ := DoWithInput(world)
	assert.Equal(t, 5, part1)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	part1, _ := DoWithInput(world)
	assert.Equal(t, 2374, part1)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	_, part2 := DoWithInput(world)
	assert.Equal(t, "mxmxvkd,sqjhc,fvjkl", part2)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	_, part2 := DoWithInput(world)
	assert.Equal(t, "fbtqkzc,jbbsjh,cpttmnv,ccrbr,tdmqcl,vnjxjg,nlph,mzqjxq", part2)
}
