package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	Matrix matrix.Matrix[rune]
	Start  utils.Vector2i
}

func walk(matrix matrix.Matrix[rune], start utils.Vector2i) map[utils.Vector2i]struct{} {
	dir := utils.Down
	pos := start

	visited := make(map[utils.Vector2i]struct{})

	for {
		visited[pos] = struct{}{}

		for i := 0; i < 5; i++ {
			if i == 2 {
				// do not return back
				continue
			}

			if i == 4 {
				panic("Unable to find rotation")
			}

			newDir := dir.Rotate(-i)
			newPos := pos.Add(newDir.ToStep())

			char, ok := matrix.GetVSafe(newPos)
			if !ok {
				return visited
			}

			if char == '.' {
				dir = newDir
				pos = newPos
				break
			}

			if char != '#' {
				panic("Unknown char " + string(char))
			}
		}

	}

	return visited
}

func DoWithInputPart01(world World) int {
	visited := walk(world.Matrix, world.Start)

	return len(visited)
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	start := utils.Vector2i{}

	mapper := func(char rune, x, y int) rune {
		if char == '^' {
			start = utils.Vector2i{X: x, Y: y}
			return '.'
		}

		return char
	}

	return World{
		Matrix: parsers.ParseToMatrixIndexed(r, mapper),
		Start:  start,
	}
}
