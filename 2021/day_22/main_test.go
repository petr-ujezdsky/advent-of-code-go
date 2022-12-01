package day_22

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_regexCube(t *testing.T) {
	matches := regexCube.FindStringSubmatch("on x=10..12,y=10..12,z=10..12")
	assert.Equal(t, []string{"on x=10..12,y=10..12,z=10..12", "on", "10", "12", "10", "12", "10", "12"}, matches)

	matches = regexCube.FindStringSubmatch("on x=-54112..-39298,y=-85059..-49293,z=-27449..7877")
	assert.Equal(t, []string{"on x=-54112..-39298,y=-85059..-49293,z=-27449..7877", "on", "-54112", "-39298", "-85059", "-49293", "-27449", "7877"}, matches)
}

func Test_01_example_1(t *testing.T) {
	reader, err := os.Open("data-00-example-1.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)
	assert.Equal(t, 4, len(cubes))

	world := NewCubeSymmetric(50, false)

	count := NaiveCount(world, cubes)
	assert.Equal(t, 39, count)
}

func Test_01_example_2(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)
	assert.Equal(t, 22, len(cubes))

	world := NewCubeSymmetric(50, false)

	count := NaiveCount(world, cubes[:20])
	assert.Equal(t, 590784, count)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	cubes := ParseInput(reader)
	assert.Equal(t, 420, len(cubes))

	world := NewCubeSymmetric(50, false)

	count := NaiveCount(world, cubes[:20])
	assert.Equal(t, 620241, count)
}
