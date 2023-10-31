package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 678, len(world.Program))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(1, world)
	assert.Equal(t, 16209841, result)
}

func Test_02_example(t *testing.T) {
	// equal to (position mode)
	world := ParseInput(strings.NewReader("3,9,8,9,10,9,4,9,99,-1,8"))
	assert.Equal(t, 0, DoWithInputPart02(0, world))
	assert.Equal(t, 0, DoWithInputPart02(1, world))
	assert.Equal(t, 0, DoWithInputPart02(2, world))
	assert.Equal(t, 0, DoWithInputPart02(3, world))
	assert.Equal(t, 0, DoWithInputPart02(4, world))
	assert.Equal(t, 0, DoWithInputPart02(5, world))
	assert.Equal(t, 0, DoWithInputPart02(6, world))
	assert.Equal(t, 0, DoWithInputPart02(7, world))
	assert.Equal(t, 1, DoWithInputPart02(8, world))
	assert.Equal(t, 0, DoWithInputPart02(9, world))

	// less than (position mode)
	world = ParseInput(strings.NewReader("3,9,7,9,10,9,4,9,99,-1,8"))
	assert.Equal(t, 1, DoWithInputPart02(0, world))
	assert.Equal(t, 1, DoWithInputPart02(1, world))
	assert.Equal(t, 1, DoWithInputPart02(2, world))
	assert.Equal(t, 1, DoWithInputPart02(3, world))
	assert.Equal(t, 1, DoWithInputPart02(4, world))
	assert.Equal(t, 1, DoWithInputPart02(5, world))
	assert.Equal(t, 1, DoWithInputPart02(6, world))
	assert.Equal(t, 1, DoWithInputPart02(7, world))
	assert.Equal(t, 0, DoWithInputPart02(8, world))
	assert.Equal(t, 0, DoWithInputPart02(9, world))

	// equal to (immediate mode)
	world = ParseInput(strings.NewReader("3,3,1108,-1,8,3,4,3,99"))
	assert.Equal(t, 0, DoWithInputPart02(0, world))
	assert.Equal(t, 0, DoWithInputPart02(1, world))
	assert.Equal(t, 0, DoWithInputPart02(2, world))
	assert.Equal(t, 0, DoWithInputPart02(3, world))
	assert.Equal(t, 0, DoWithInputPart02(4, world))
	assert.Equal(t, 0, DoWithInputPart02(5, world))
	assert.Equal(t, 0, DoWithInputPart02(6, world))
	assert.Equal(t, 0, DoWithInputPart02(7, world))
	assert.Equal(t, 1, DoWithInputPart02(8, world))
	assert.Equal(t, 0, DoWithInputPart02(9, world))

	// less than (immediate mode)
	world = ParseInput(strings.NewReader("3,3,1107,-1,8,3,4,3,99"))
	assert.Equal(t, 1, DoWithInputPart02(0, world))
	assert.Equal(t, 1, DoWithInputPart02(1, world))
	assert.Equal(t, 1, DoWithInputPart02(2, world))
	assert.Equal(t, 1, DoWithInputPart02(3, world))
	assert.Equal(t, 1, DoWithInputPart02(4, world))
	assert.Equal(t, 1, DoWithInputPart02(5, world))
	assert.Equal(t, 1, DoWithInputPart02(6, world))
	assert.Equal(t, 1, DoWithInputPart02(7, world))
	assert.Equal(t, 0, DoWithInputPart02(8, world))
	assert.Equal(t, 0, DoWithInputPart02(9, world))

	// jump (position mode)
	world = ParseInput(strings.NewReader("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"))
	assert.Equal(t, 0, DoWithInputPart02(0, world))
	assert.Equal(t, 1, DoWithInputPart02(1, world))
	assert.Equal(t, 1, DoWithInputPart02(2, world))
	assert.Equal(t, 1, DoWithInputPart02(3, world))
	assert.Equal(t, 1, DoWithInputPart02(4, world))
	assert.Equal(t, 1, DoWithInputPart02(5, world))
	assert.Equal(t, 1, DoWithInputPart02(6, world))
	assert.Equal(t, 1, DoWithInputPart02(7, world))
	assert.Equal(t, 1, DoWithInputPart02(8, world))
	assert.Equal(t, 1, DoWithInputPart02(9, world))

	// jump tests (immediate mode)
	world = ParseInput(strings.NewReader("3,3,1105,-1,9,1101,0,0,12,4,12,99,1"))
	assert.Equal(t, 0, DoWithInputPart02(0, world))
	assert.Equal(t, 1, DoWithInputPart02(1, world))
	assert.Equal(t, 1, DoWithInputPart02(2, world))
	assert.Equal(t, 1, DoWithInputPart02(3, world))
	assert.Equal(t, 1, DoWithInputPart02(4, world))
	assert.Equal(t, 1, DoWithInputPart02(5, world))
	assert.Equal(t, 1, DoWithInputPart02(6, world))
	assert.Equal(t, 1, DoWithInputPart02(7, world))
	assert.Equal(t, 1, DoWithInputPart02(8, world))
	assert.Equal(t, 1, DoWithInputPart02(9, world))

	// more complex
	world = ParseInput(strings.NewReader("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"))
	assert.Equal(t, 999, DoWithInputPart02(0, world))
	assert.Equal(t, 999, DoWithInputPart02(1, world))
	assert.Equal(t, 999, DoWithInputPart02(2, world))
	assert.Equal(t, 999, DoWithInputPart02(3, world))
	assert.Equal(t, 999, DoWithInputPart02(4, world))
	assert.Equal(t, 999, DoWithInputPart02(5, world))
	assert.Equal(t, 999, DoWithInputPart02(6, world))
	assert.Equal(t, 999, DoWithInputPart02(7, world))
	assert.Equal(t, 1000, DoWithInputPart02(8, world))
	assert.Equal(t, 1001, DoWithInputPart02(9, world))
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(5, world)
	assert.Equal(t, 8834787, result)
}
