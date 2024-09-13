package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	Matrix matrix.Matrix[rune]
	Start  utils.Vector2i
}

func DoWithInputPart01(world World, steps int) int {
	positions := findPositions(world, steps)

	printPositions(world.Matrix, world.Start, positions)

	return len(positions)
}

func findPositions(world World, steps int) map[utils.Vector2i]struct{} {
	currentPositions := map[utils.Vector2i]struct{}{}

	currentPositions[world.Start] = struct{}{}

	for i := 0; i < steps; i++ {
		nextPositions := map[utils.Vector2i]struct{}{}

		for pos := range currentPositions {
			for _, step := range utils.Direction4Steps {
				nextPos := pos.Add(step)

				nextPosLooped := loopPosition(nextPos, world.Matrix)

				// check position on map
				ch := world.Matrix.GetV(nextPosLooped)

				if ch == '#' {
					// rock -> can not step on it
					continue
				}

				// save this position
				nextPositions[nextPos] = struct{}{}
			}
		}

		// swap positions
		currentPositions = nextPositions
	}

	return currentPositions
}

func printPositions(m matrix.Matrix[rune], start utils.Vector2i, positions map[utils.Vector2i]struct{}) {
	str := matrix.StringFmtSeparatorIndexed[rune](m, true, "", func(value rune, x, y int) string {
		pos := utils.Vector2i{X: x, Y: y}

		if _, ok := positions[pos]; ok {
			return "O"
		}

		if pos == start {
			return "S"
		}

		return string(value)
	})

	fmt.Println(str)
}

//func printCounts(m matrix.MatrixNumber[int]) {
//	str := matrix.StringFmtSeparatorIndexed[int](m, true, "", func(value int, x, y int) string {
//		return fmt.Sprintf("%10d", value)
//	})
//
//	fmt.Println(str)
//}

func DoWithInputPart02(world World, steps int) int {
	//samples := 9
	//width := samples * world.Matrix.Width
	//
	//positions := findPositions(world, width/2)
	//
	//counts := matrix.NewMatrixInt(samples, samples)
	//
	//for pos := range positions {
	//	x := samples/2 + int(math.Floor(float64(pos.X)/float64(world.Matrix.Width)))
	//	y := samples/2 + int(math.Floor(float64(pos.Y)/float64(world.Matrix.Height)))
	//
	//	count := counts.Get(x, y)
	//	counts.Set(x, y, count+1)
	//}
	//
	//printCounts(counts)

	// 0          0         0         0       958      5630       935         0         0         0
	// 1          0         0       958      6549      7521      6548       935         0         0
	// 2          0       958      6549      7521      7467      7521      6548       935         0
	// 3        958      6549      7521      7467      7521      7467      7521      6548       935
	// 4       5616      7521      7467      7521      7467      7521      7467      7521      5629
	// 5        966      6534      7521      7467      7521      7467      7521      6548       963
	// 6          0       966      6534      7521      7467      7521      6548       963         0
	// 7          0         0       966      6534      7521      6548       963         0         0
	// 8          0         0         0       966      5615       963         0         0         0

	totalCount := 0

	// left corner
	totalCount += 958 + 5616 + 966

	actualSamples := 1 + 2*((steps-world.Matrix.Width/2)/world.Matrix.Width) // 404601

	// left part
	// sum N=0..actualSamples/2-2  col2 + N*(7521+7467)
	col2 := 958 + 6549 + 7521 + 6534 + 966

	c := actualSamples/2 - 2
	coef := 7521 + 7467
	sumLeft := (c+1)*col2 + coef*c*(c+1)/2

	totalCount += sumLeft

	// middle column
	middle := 5630 + 5615 + 7521*(actualSamples/2) + 7467*(actualSamples/2-1)
	totalCount += middle

	// right part
	colSecondToLast := 935 + 6548 + 7521 + 6548 + 963
	// sum N=0..actualSamples/2-2  colSecondToLast + N*(7521+7467)
	sumRight := (c+1)*colSecondToLast + coef*c*(c+1)/2

	totalCount += sumRight

	// right corner
	totalCount += 935 + 5629 + 963

	return totalCount
}

func loopPosition(pos utils.Vector2i, m matrix.Matrix[rune]) utils.Vector2i {
	x := utils.ModFloor(pos.X, m.GetWidth())
	y := utils.ModFloor(pos.Y, m.GetHeight())

	return utils.Vector2i{X: x, Y: y}
}

func ParseInput(r io.Reader) World {
	start := utils.Vector2i{}

	parseItem := func(char rune, x, y int) rune {
		if char == 'S' {
			start = utils.Vector2i{X: x, Y: y}
			return '.'
		}

		return char
	}

	return World{
		Matrix: parsers.ParseToMatrixIndexed(r, parseItem),
		Start:  start,
	}
}
