package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Vector2i = utils.Vector2i
type MatrixInt = utils.MatrixInt

var dirs = []Vector2i{
	// left
	{-1, 0},
	// up
	{0, -1},
	// right
	{1, 0},
	// down
	{0, 1},
}

type World struct {
	Heights    MatrixInt
	Start, End Vector2i
}

func h(endPos Vector2i) func(Vector2i) int {
	return func(pos Vector2i) int {
		// manhattan distance
		return utils.Abs(pos.X-endPos.X) + utils.Abs(pos.Y-endPos.Y)
	}
}

func d(_ MatrixInt) func(Vector2i, Vector2i) int {
	return func(nodeFrom, nodeTo Vector2i) int {
		// distance is always 1 - going to the neighbour
		return 1
	}
}

func neighbours(heights MatrixInt) func(origin Vector2i) []Vector2i {
	return func(origin Vector2i) []Vector2i {
		var neighbours []Vector2i
		for _, dir := range dirs {
			nextPos := origin.Add(dir)

			// check validity
			currentHeight := heights.GetV(origin)
			neighbourHeight, ok := heights.GetVSafe(nextPos)
			if ok && (neighbourHeight-currentHeight <= 1) {
				neighbours = append(neighbours, nextPos)
			}
		}

		return neighbours
	}
}

func DoWithInput(world World) int {
	m := world.Heights
	endPos := world.End
	_, score, found := utils.AStar(world.Start, endPos, h(endPos), d(m), neighbours(m))
	if found {
		return score
	}

	panic("No path found!")
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rows [][]int

	start, end := Vector2i{}, Vector2i{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []int

		for x, digitAscii := range []rune(line) {
			if digitAscii == 'S' {
				digitAscii = 'a'
				start = Vector2i{x, y}
			}

			if digitAscii == 'E' {
				digitAscii = 'z'
				end = Vector2i{x, y}
			}

			digit := int(digitAscii) - int('a')
			row = append(row, digit)
		}

		rows = append(rows, row)
		y++
	}

	heights := utils.NewMatrixNumberRowNotation(rows)

	return World{
		Heights: heights,
		Start:   start,
		End:     end,
	}
}
