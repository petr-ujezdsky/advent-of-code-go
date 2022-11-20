package day_12

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, 6, len(world.nodes))
	assert.Equal(t, 2, len(world.startNode.neighbours))

	neighbour := world.startNode.neighbours[0]
	assert.Equal(t, "A", neighbour.id)
	assert.Equal(t, big, neighbour.nodeType)

	neighbour = world.startNode.neighbours[1]
	assert.Equal(t, "b", neighbour.id)
	assert.Equal(t, small, neighbour.nodeType)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	paths := FindAllPaths(world)
	fmt.Println(paths)
	assert.Equal(t, 10, len(paths))
}

func Test_01_example_02(t *testing.T) {
	reader, err := os.Open("data-00-example-02.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	paths := FindAllPaths(world)
	fmt.Println(paths)
	assert.Equal(t, 19, len(paths))
}

func Test_01_example_03(t *testing.T) {
	reader, err := os.Open("data-00-example-03.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	paths := FindAllPaths(world)
	assert.Equal(t, 226, len(paths))
}

func Test_011(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	paths := FindAllPaths(world)
	assert.Equal(t, 5254, len(paths))
}
