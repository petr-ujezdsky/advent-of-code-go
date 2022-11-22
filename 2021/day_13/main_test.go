package day_12

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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
	assert.Equal(t, utils.Vector2i{6, 10}, *world.points[0])
	assert.Equal(t, utils.Vector2i{9, 0}, *world.points[17])

	assert.Equal(t, 2, len(world.folds))
	assert.Equal(t, Fold{7, true}, world.folds[0])
	assert.Equal(t, Fold{5, false}, world.folds[1])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	dotsCount := FoldPaper(world, 1)
	assert.Equal(t, 17, dotsCount)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world, err := ParseInput(reader)
	assert.Nil(t, err)

	dotsCount := FoldPaper(world, 1)
	assert.Equal(t, 655, dotsCount)
}
