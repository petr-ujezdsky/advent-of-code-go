package main

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
	"math"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example1.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 40, len(world.Asteroids))
}

func Test_01_example0(t *testing.T) {
	reader, err := os.Open("data-00-example0.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 8, result)
}

func Test_01_example1(t *testing.T) {
	reader, err := os.Open("data-00-example1.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 33, result)
}

func Test_01_example2(t *testing.T) {
	reader, err := os.Open("data-00-example2.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 35, result)
}

func Test_01_example3(t *testing.T) {
	reader, err := os.Open("data-00-example3.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 41, result)
}

func Test_01_example4(t *testing.T) {
	reader, err := os.Open("data-00-example4.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 210, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 0, result)
}

func Test_angle(t *testing.T) {
	//	assert.Equal(t, 0.0, angle(utils.Vector2i{0, 1}))
	//	assert.Equal(t, 1*math.Pi/4, angle(utils.Vector2i{1, 1}))
	//	assert.Equal(t, 2*math.Pi/4, angle(utils.Vector2i{1, 0}))
	//	assert.Equal(t, 3*math.Pi/4, angle(utils.Vector2i{1, -1}))
	//	assert.Equal(t, 4*math.Pi/4, angle(utils.Vector2i{0, -1}))
	//	assert.Equal(t, 5*math.Pi/4, angle(utils.Vector2i{-1, -1}))
	//	assert.Equal(t, 6*math.Pi/4, angle(utils.Vector2i{-1, 0}))
	//	assert.Equal(t, 7*math.Pi/4, angle(utils.Vector2i{-1, 1}))

	assert.Equal(t, 0.0, angle(utils.Vector2i{0, -1}))
	assert.Equal(t, 1*math.Pi/4, angle(utils.Vector2i{1, -1}))
	assert.Equal(t, 2*math.Pi/4, angle(utils.Vector2i{1, 0}))
	assert.Equal(t, 3*math.Pi/4, angle(utils.Vector2i{1, 1}))
	assert.Equal(t, 4*math.Pi/4, angle(utils.Vector2i{0, 1}))
	assert.Equal(t, 5*math.Pi/4, angle(utils.Vector2i{-1, 1}))
	assert.Equal(t, 6*math.Pi/4, angle(utils.Vector2i{-1, 0}))
	assert.Equal(t, 7*math.Pi/4, angle(utils.Vector2i{-1, -1}))
}

func Test_02_angles(t *testing.T) {
	assert.Equal(t, math.Pi/2, math.Atan2(1, 0))

	assert.Equal(t, math.Pi/4, math.Atan2(1, 1))
	assert.Equal(t, 3*math.Pi/4, math.Atan2(1, -1))
	assert.Equal(t, -3*math.Pi/4, math.Atan2(-1, -1))
	assert.Equal(t, -math.Pi/4, math.Atan2(-1, 1))
}

func Test_02_example4(t *testing.T) {
	reader, err := os.Open("data-00-example4.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 802, result)
}

func Test_02_example5(t *testing.T) {
	reader, err := os.Open("data-00-example5.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 802, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 608, result)
}
