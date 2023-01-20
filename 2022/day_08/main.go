package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"io"
	"math"
)

type Matrix2i = matrix.MatrixInt
type Vector2i = utils.Vector2i

func processColumn(col []int, x int, points map[Vector2i]struct{}, transposed bool, from, to int) {
	max := math.MinInt

	step := utils.Signum(to - from)
	for y := from; y != to; y += step {
		height := col[y]

		if height > max {
			max = height

			// do not count beginning (edge)
			if y == from {
				continue
			}

			point := Vector2i{x, y}
			if transposed {
				point = point.Reverse()
			}

			points[point] = struct{}{}
		}
	}
}

func processVertically(heights Matrix2i, points map[Vector2i]struct{}, transposed bool) {
	for x, col := range heights.Columns {
		// skip vertical edges
		if x == 0 || x == heights.Width-1 {
			continue
		}

		// visible from top
		processColumn(col, x, points, transposed, 0, len(col)-1)
		// visible from bottom
		processColumn(col, x, points, transposed, len(col)-1, 0)
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

func countTreesFromHouseView(col []int, houseHeight, from, to int) int {
	count := 0

	step := utils.Signum(to - from)
	for i := from; i != to+step || step == 0; i += step {
		treeHeight := col[i]

		count++
		if treeHeight >= houseHeight {
			break
		}

		if step == 0 {
			break
		}
	}

	return count
}

func processTreeHouseVertically(heights Matrix2i, scores *Matrix2i) {
	for x, col := range heights.Columns {
		// skip vertical edges
		if x == 0 || x == heights.Width-1 {
			continue
		}

		for y, houseHeight := range col {
			// skip horizontal edges
			if y == 0 || y == heights.Height-1 {
				continue
			}

			score := 1

			// view up
			score *= countTreesFromHouseView(col, houseHeight, y-1, 0)

			// view down
			score *= countTreesFromHouseView(col, houseHeight, y+1, len(col)-1)

			scores.Columns[x][y] *= score
		}
	}
}

func FindBestTreeHouseLocationScore(heights Matrix2i) int {
	// default value is 1 because the score is multiplied
	scores := matrix.NewMatrixInt(heights.Width, heights.Height).SetAll(1)

	// process vertically
	processTreeHouseVertically(heights, &scores)

	// process horizontally
	heights = heights.Transpose()
	scores = scores.Transpose()
	processTreeHouseVertically(heights, &scores)

	_, max := scores.ArgMax()

	return max
}

func ParseInput(r io.Reader) matrix.MatrixInt {
	return utils.ParseToMatrixP(r)
}
