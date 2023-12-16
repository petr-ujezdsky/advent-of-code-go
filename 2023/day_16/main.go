package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Tile struct {
	Char    rune
	NextFor [][]utils.Direction4
	Visited []bool
}

type World struct {
	Tiles matrix.Matrix[Tile]
}

func DoWithInputPart01(world World) int {
	return Walk(world.Tiles)
}

func Walk(tiles matrix.Matrix[Tile]) int {
	energized := make(map[utils.Vector2i]struct{})

	walkRecursive(utils.Vector2i{X: -1, Y: tiles.Height - 1}, utils.Right, energized, tiles)

	return len(energized)
}

func walkRecursive(position utils.Vector2i, direction utils.Direction4, energized map[utils.Vector2i]struct{}, tiles matrix.Matrix[Tile]) {
	nextPosition := position.Add(direction.ToStep())
	nextTile, ok := tiles.GetVSafe(nextPosition)

	// outside the world
	if !ok {
		return
	}

	// already visited from this direction
	if nextTile.Visited[direction] {
		return
	}

	// mark as visited from this direction
	nextTile.Visited[direction] = true

	// energize
	energized[nextPosition] = struct{}{}

	// compute where to go next
	nextDirections := nextTile.NextFor[direction]

	for _, nextDirection := range nextDirections {
		walkRecursive(nextPosition, nextDirection, energized, tiles)
	}
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(char rune) Tile {
		nextFor := make([][]utils.Direction4, 4)

		switch char {
		case '.':
			{
				nextFor[utils.Up] = []utils.Direction4{utils.Up}
				nextFor[utils.Right] = []utils.Direction4{utils.Right}
				nextFor[utils.Down] = []utils.Direction4{utils.Down}
				nextFor[utils.Left] = []utils.Direction4{utils.Left}
			}
		case '\\':
			{
				nextFor[utils.Up] = []utils.Direction4{utils.Left}
				nextFor[utils.Right] = []utils.Direction4{utils.Down}
				nextFor[utils.Down] = []utils.Direction4{utils.Right}
				nextFor[utils.Left] = []utils.Direction4{utils.Up}
			}
		case '/':
			{
				nextFor[utils.Up] = []utils.Direction4{utils.Right}
				nextFor[utils.Right] = []utils.Direction4{utils.Up}
				nextFor[utils.Down] = []utils.Direction4{utils.Left}
				nextFor[utils.Left] = []utils.Direction4{utils.Down}
			}
		case '-':
			{
				nextFor[utils.Up] = []utils.Direction4{utils.Left, utils.Right}
				nextFor[utils.Right] = []utils.Direction4{utils.Right}
				nextFor[utils.Down] = []utils.Direction4{utils.Left, utils.Right}
				nextFor[utils.Left] = []utils.Direction4{utils.Left}
			}
		case '|':
			{
				nextFor[utils.Up] = []utils.Direction4{utils.Up}
				nextFor[utils.Right] = []utils.Direction4{utils.Up, utils.Down}
				nextFor[utils.Down] = []utils.Direction4{utils.Down}
				nextFor[utils.Left] = []utils.Direction4{utils.Up, utils.Down}
			}
		}

		return Tile{
			Char:    char,
			NextFor: nextFor,
			Visited: make([]bool, 4),
		}
	}

	tiles := parsers.ParseToMatrix(r, parseItem).FlipVertical()

	return World{Tiles: tiles}
}
