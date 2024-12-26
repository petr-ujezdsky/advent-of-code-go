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

	assert.Equal(t, 19, len(world.DiskMap))
}

func Test_01_parseBlocks(t *testing.T) {
	blocks := parseBlocks([]int{1, 2, 3, 4, 5})

	assert.Equal(t, "0..111....22222", blocks.String())
}

func Test_01_parseBlocks2(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)
	blocks := parseBlocks(world.DiskMap)

	assert.Equal(t, "00...111...2...333.44.5555.6666.777.888899", blocks.String())
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 1928, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 6344673854800, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 2858, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 6360363199987, result)
}
