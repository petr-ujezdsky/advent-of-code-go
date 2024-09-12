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

	assert.Equal(t, utils.Vector2i{X: 5, Y: 5}, world.Start)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world, 6)
	assert.Equal(t, 16, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world, 64)
	assert.Equal(t, 3699, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world, 6)
	assert.Equal(t, 16, result)

	result = DoWithInputPart02(world, 10)
	assert.Equal(t, 50, result)

	result = DoWithInputPart02(world, 50)
	assert.Equal(t, 1594, result)

	result = DoWithInputPart02(world, 100)
	assert.Equal(t, 6536, result)

	result = DoWithInputPart02(world, 500)
	assert.Equal(t, 167_004, result)

	result = DoWithInputPart02(world, 1000)
	assert.Equal(t, 668_697, result)

	result = DoWithInputPart02(world, 5000)
	assert.Equal(t, 16_733_044, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world, 26_501_365)
	assert.Equal(t, 0, result)
}

func Benchmark_Part02(b *testing.B) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(b, err)

	world := ParseInput(reader)

	for i := 0; i < b.N; i++ {
		result := DoWithInputPart02(world, 50)
		assert.Equal(b, 1594, result)
	}
}
