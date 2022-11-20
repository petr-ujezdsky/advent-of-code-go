package day_12

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

	assert.Equal(t, 6, len(world.nodes))
	assert.Equal(t, 2, len(world.startNode.neighbours))

	neighbour := world.startNode.neighbours[0]
	assert.Equal(t, "A", neighbour.id)
	assert.Equal(t, big, neighbour.nodeType)

	neighbour = world.startNode.neighbours[1]
	assert.Equal(t, "b", neighbour.id)
	assert.Equal(t, small, neighbour.nodeType)
}
