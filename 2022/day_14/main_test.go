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

	world := ParseInput(reader, false)

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

	assert.Equal(t, expected, world.Cave.StringFmtSeparator("", utils.FmtFmt[rune]("%c")))

	fmt.Println(expected)
	fmt.Println()
	fmt.Println(world.Cave.StringFmtSeparator("", utils.FmtFmt[rune]("%c")))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader, false)

	result := PourSand(world, false)
	assert.Equal(t, 24, result)
	fmt.Println(world.Cave.StringFmtSeparator("", utils.FmtFmt[rune]("%c")))
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader, false)

	fmt.Println(world.Cave.StringFmtSeparator("", utils.FmtFmt[rune]("%c")))

	result := PourSand(world, false)
	assert.Equal(t, 888, result)

	fmt.Println()
	fmt.Println(world.Cave.StringFmtSeparator("", utils.FmtFmt[rune]("%c")))
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader, true)

	result := PourSand(world, true)
	assert.Equal(t, 93, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader, true)

	result := PourSand(world, true)
	assert.Equal(t, 26461, result)
}
