package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"math"
)

type Tile struct {
	Char rune
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

var metricGlobal = utils.NewMetric("Global")

func DoWithInputPart02(world World) int {
	metricGlobal.Enable()
	totalCycles := 1_000_000_000

	loopStartTiles := world.Tiles
	loopCheckAfter := 1000

	for i := 0; i < totalCycles; i++ {
		SpinCycleRocks(world)
		metricGlobal.TickTotal(100_000, totalCycles)

		fmt.Printf("#%d load: %d\n", i, ComputeLoad(world))

		if i == loopCheckAfter {
			loopStartTiles = world.Tiles.Clone()
		}

		if i > loopCheckAfter {
			if matrix.EqualFunc(world.Tiles, loopStartTiles, func(a, b Tile) bool { return a.Char == b.Char }) {
				loopLength := i - loopCheckAfter
				remainingCycles := totalCycles - i - 1

				i += (remainingCycles / loopLength) * loopLength
				loopCheckAfter = math.MaxInt
				fmt.Printf("Found loop of length %d, going to %d\n", loopLength, i)
			}
		}
	}

	return ComputeLoad(world)
}

func SpinCycleRocks(world World) {
	MoveRocks(world.Tiles, utils.Up)
	MoveRocks(world.Tiles, utils.Left)
	MoveRocks(world.Tiles, utils.Down)
	MoveRocks(world.Tiles, utils.Right)
}

func ComputeLoad(world World) int {
	totalLoad := 0

	for x := 0; x < world.Tiles.GetWidth(); x++ {
		for y := 0; y < world.Tiles.GetHeight(); y++ {
			tile := world.Tiles.Get(x, y)
			if tile.Char == 'O' {
				totalLoad += world.Tiles.GetHeight() - y
			}
		}
	}

	return totalLoad
}

func ParseInput(r io.Reader) World {
	parseItem := func(char rune) Tile {
		return Tile{
			Char: char,
		}
	}

	return World{Tiles: parsers.ParseToMatrix(r, parseItem)}
}
