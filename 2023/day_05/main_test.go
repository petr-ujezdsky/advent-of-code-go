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

	assert.Equal(t, []int{79, 14, 55, 13}, world.Seeds)
	assert.Equal(t, 7, len(world.Mappings))
	dest, ok := world.Mappings[0].Mappers[0](98)
	assert.True(t, ok)
	assert.Equal(t, dest, 50)

	assert.Equal(t, 52, world.Mappings[0].Map(50))

	// examples
	assert.Equal(t, 81, world.Mappings[0].Map(79))
	assert.Equal(t, 14, world.Mappings[0].Map(14))
	assert.Equal(t, 57, world.Mappings[0].Map(55))
	assert.Equal(t, 13, world.Mappings[0].Map(13))

}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 35, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 26273516, result)
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
