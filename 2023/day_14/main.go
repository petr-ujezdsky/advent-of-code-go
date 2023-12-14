package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type TileType int

const (
	Empty TileType = iota
	Movable
	Solid
)

type Tile struct {
	Char rune
	Type TileType
}

func (t Tile) String() string {
	return string(t.Char)
}

type World struct {
	Tiles matrix.Matrix[Tile]
}

func DoWithInputPart01(world World) int {
	totalLoad := MoveRocksUp(world.Tiles)

	return totalLoad
}

func MoveRocksUp(tiles matrix.Matrix[Tile]) int {
	totalLoad := 0

	for x, column := range tiles.Columns {
		emptyTileY := -1
		for y, tile := range column {
			char := tile.Char

			switch char {
			case '.':
				if emptyTileY == -1 {
					emptyTileY = y
				}
			case '#':
				emptyTileY = -1
			case 'O':
				if emptyTileY != -1 {
					// switch empty tile with rock
					tiles.Columns[x][y], tiles.Columns[x][emptyTileY] = tiles.Columns[x][emptyTileY], tiles.Columns[x][y]

					// update total load
					totalLoad += len(column) - emptyTileY

					// empty tile is right after the moved rock
					emptyTileY++
				} else {
					totalLoad += len(column) - y
				}
			}
		}
	}

	return totalLoad
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(char rune) Tile {
		var tileType TileType
		switch char {
		case '.':
			tileType = Empty
		case 'O':
			tileType = Movable
		case '#':
			tileType = Solid
		}

		return Tile{
			Char: char,
			Type: tileType,
		}
	}

	return World{Tiles: parsers.ParseToMatrix(r, parseItem)}
}
