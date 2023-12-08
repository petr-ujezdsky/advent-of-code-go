package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_regex(t *testing.T) {
	parts := regexMapDef.FindStringSubmatch("AAA = (BBB, CCC)")

	assert.Equal(t, 4, len(parts))
	assert.Equal(t, "AAA", parts[1])
	assert.Equal(t, "BBB", parts[2])
	assert.Equal(t, "CCC", parts[3])
}

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 2, len(world.Directions))
	assert.Equal(t, 7, len(world.Maps))

	m := world.Maps["AAA"]
	assert.Equal(t, "AAA", m.Name)
	assert.Equal(t, "BBB", m.Next[Left].Name)
	assert.Equal(t, "CCC", m.Next[Right].Name)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 2, result)
}

func Test_01_example_02(t *testing.T) {
	reader, err := os.Open("data-00-example-02.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 6, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 19951, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-01-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 2, len(world.StartingMaps))

	result := DoWithInputPart02(world)
	assert.Equal(t, 6, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 16342438708751, result)
}
