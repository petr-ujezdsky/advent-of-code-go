package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
)

type Matrix2i = utils.Matrix2i
type Vector2i = utils.Vector2i

func processColumn(col []int, x int, points map[Vector2i]struct{}, transposed, reversed bool) int {
	// skip bottom edge - it is always visible
	col = col[:len(col)-1]

	max := math.MinInt
	count := 0

	for i, height := range col {
		if height > max {
			max = height

			// do not count top edge
			if i == 0 {
				continue
			}

			y := i
			if reversed {
				y = len(col) - 1 - i + 1
			}

			point := Vector2i{x, y}
			if transposed {
				point = point.Reverse()
			}

			points[point] = struct{}{}
		}
	}

	return count
}

func processVertically(heights Matrix2i, points map[Vector2i]struct{}, transposed bool) {
	for x, col := range heights.Columns {
		// skip vertical edges
		if x == 0 || x == heights.Width-1 {
			continue
		}

		// visible from top
		processColumn(col, x, points, transposed, false)
		// visible from bottom
		processColumn(utils.Reverse(col), x, points, transposed, true)
	}
}

func CountVisibleTrees(heights Matrix2i) int {
	// start with edges
	count := 2*heights.Height + 2*heights.Width - 4

	interiorPoints := make(map[Vector2i]struct{})
	// process vertically
	processVertically(heights, interiorPoints, false)
	// process horizontally
	processVertically(heights.Transpose(), interiorPoints, true)

	count += len(interiorPoints)
	return count
}

func ParseInput(r io.Reader) utils.Matrix2i {
	return utils.ParseToMatrixP(r)
}
