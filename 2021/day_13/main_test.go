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

	assert.Equal(t, 18, len(world.points))
	assert.Equal(t, 6, world.points[0].X)
	assert.Equal(t, 10, world.points[0].Y)

	assert.Equal(t, 2, len(world.folds))
	assert.Equal(t, 7, world.folds[0].index)
	assert.Equal(t, false, world.folds[0].dirX)
}
