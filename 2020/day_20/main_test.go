package main

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	assert.Equal(t, 9, len(world.Tiles))
	assert.Equal(t, 3079, world.Tiles[3079].Id)
	assert.Equal(t, false, world.Tiles[3079].Data.Columns[1][0])
	assert.Equal(t, true, world.Tiles[3079].Data.Columns[1][1])
	assert.Equal(t, false, world.Tiles[3079].Data.Columns[1][2])
	assert.Equal(t, 10, world.Tiles[3079].Data.Width)
	assert.Equal(t, 10, world.Tiles[3079].Data.Height)
}

func Test_01_example_all(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	for _, tile := range world.Tiles {
		picture := ConnectTilesUsing(tile, world.Tiles)
		cornersProduct := multiplyCorners(picture)
		assert.Equal(t, 20899048083289, cornersProduct)
	}
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 20899048083289, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart01(world)
	assert.Equal(t, 11788777383197, result)
}

func Test_02_example_picture(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	tile := world.Tiles[3079]
	connectedTiles := ConnectTilesUsing(tile, world.Tiles)
	picture := TilesToPicture(connectedTiles)

	expected := utils.Msg(`
.#.#..#.##...#.##..#####
###....#.#....#..#......
##.##.###.#.#..######...
###.#####...#.#####.#..#
##.#....#.##.####...#.##
...########.#....#####.#
....#..#...##..#.#.###..
.####...#..#.....#......
#..#.##..#..###.#.##....
#.####..#.####.#.#.###..
###.#.#...#.######.#..##
#.####....##..########.#
##..##.#...#...#.#.#.#..
...#..#..#.#.##..###.###
.#.#....#.##.#...###.##.
###.#...#..#.##.######..
.#.#.###.##.##.#..#.##..
.####.###.#...###.#..#.#
..#.#..#..#.#.#.####.###
#..####...#.#.#.###.###.
#####..#####...###....##
#.##..#..#...#..####...#
.#.###..##..##..####.##.
...###...##...#...#..###`)

	assert.Equal(t, expected, picture.StringFmtSeparator("", matrix.FmtBoolean[bool]))
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 273, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	world := ParseInput(reader)

	result := DoWithInputPart02(world)
	assert.Equal(t, 2242, result)
}
