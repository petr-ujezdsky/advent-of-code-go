package main

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	expected := utils.Msg(`
......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
........#.
#########.`)

	assert.Equal(t, expected, world.StringFmtSeparator("", utils.FmtFmt[rune]("%c")))

	fmt.Println(expected)
	fmt.Println()
	fmt.Println(world.StringFmtSeparator("", utils.FmtFmt[rune]("%c")))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInput(world)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInput(world)
	assert.Equal(t, 0, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInput(world)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInput(world)
	assert.Equal(t, 0, result)
}
