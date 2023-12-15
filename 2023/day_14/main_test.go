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

	assert.Equal(t, 10, world.Tiles.Width)
	assert.Equal(t, 10, world.Tiles.Height)
}

func TestMoveRocks(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	MoveRocks(world.Tiles, utils.Up)

	str := world.Tiles.StringFmtSeparator("", func(tile Tile) string { return string(tile.Char) })
	fmt.Println(str)

	expected := utils.Msg(`
OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#....`)

	assert.Equal(t, expected, str)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 136, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 110128, result)
}

func TestSpinCycleRocks(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	SpinCycleRocks(world)

	str := world.Tiles.StringFmtSeparator("", func(tile Tile) string { return string(tile.Char) })
	fmt.Println(str)

	expected := utils.Msg(`
.....#....
....#...O#
...OO##...
.OO#......
.....OOO#.
.O#...O#.#
....O#....
......OOOO
#...O###..
#..OO#....`)

	assert.Equal(t, expected, str)

	SpinCycleRocks(world)

	str = world.Tiles.StringFmtSeparator("", func(tile Tile) string { return string(tile.Char) })
	fmt.Println(str)

	expected = utils.Msg(`
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#..OO###..
#.OOO#...O`)

	assert.Equal(t, expected, str)

	SpinCycleRocks(world)

	str = world.Tiles.StringFmtSeparator("", func(tile Tile) string { return string(tile.Char) })
	fmt.Println(str)

	expected = utils.Msg(`
.....#....
....#...O#
.....##...
..O#......
.....OOO#.
.O#...O#.#
....O#...O
.......OOO
#...O###.O
#.OOO#...O`)

	assert.Equal(t, expected, str)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 64, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 103861, result)
}
