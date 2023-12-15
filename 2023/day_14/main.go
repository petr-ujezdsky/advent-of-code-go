package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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
	totalLoad := MoveRocks(world.Tiles, utils.Up)

	return totalLoad
}

func MoveRocks(tiles matrix.Matrix[Tile], direction utils.Direction4) int {
	totalLoad := 0

	// direction number is equal to rotation steps count
	view := matrix.NewMatrixViewRotated90CounterClockwise[Tile](tiles, int(direction))
	for x := 0; x < view.GetWidth(); x++ {
		emptyTileY := -1

		for y := 0; y < view.GetHeight(); y++ {
			tile := view.Get(x, y)
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
					empty := view.Get(x, emptyTileY)
					view.Set(x, emptyTileY, tile)
					view.Set(x, y, empty)

					// update total load
					totalLoad += view.GetHeight() - emptyTileY

					// empty tile is right after the moved rock
					emptyTileY++
				} else {
					totalLoad += view.GetHeight() - y
				}
			}
		}
	}

	return totalLoad
}

func DoWithInputPart02(world World) int {
	lastTotalLoad := 0

	for i := 0; i < 1_000_000_000; i++ {
		lastTotalLoad = SpinCycleRocks(world)
	}

	return lastTotalLoad
}

func SpinCycleRocks(world World) int {
	MoveRocks(world.Tiles, utils.Up)
	MoveRocks(world.Tiles, utils.Left)
	MoveRocks(world.Tiles, utils.Down)
	return MoveRocks(world.Tiles, utils.Right)
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
