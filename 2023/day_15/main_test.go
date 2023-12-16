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

	assert.Equal(t, 11, len(world.Steps))

	step := world.Steps[0]
	assert.Equal(t, "rn=1", step.Raw)
	assert.Equal(t, "rn", step.Label)
	assert.Equal(t, '=', step.Operation)
	assert.Equal(t, 1, step.FocalLength)

	step = world.Steps[1]
	assert.Equal(t, "cm-", step.Raw)
	assert.Equal(t, "cm", step.Label)
	assert.Equal(t, '-', step.Operation)
	assert.Equal(t, 0, step.FocalLength)

}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 1320, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 504036, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 145, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 57849, result)
}
