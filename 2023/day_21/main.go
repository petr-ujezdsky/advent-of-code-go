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
	currentPositions := map[utils.Vector2i]struct{}{}

	currentPositions[world.Start] = struct{}{}

	for i := 0; i < steps; i++ {
		nextPositions := map[utils.Vector2i]struct{}{}

		for pos := range currentPositions {
			for _, step := range utils.Direction4Steps {
				nextPos := pos.Add(step)

				// check position on map
				ch, ok := world.Matrix.GetVSafe(nextPos)
				if !ok {
					// out of bounds
					continue
				}

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

		printPositions(world.Matrix, world.Start, currentPositions)
	}

	return len(currentPositions)
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

func DoWithInputPart02(world World, steps int) int {
	currentPositions := map[utils.Vector2i]struct{}{}

	currentPositions[world.Start] = struct{}{}

	for i := 0; i < steps; i++ {

		nextPositions := map[utils.Vector2i]struct{}{}

		for pos := range currentPositions {
			for _, step := range utils.Direction4Steps {
				nextPos := pos.Add(step)

				nextPosLooped := loopPosition(nextPos, world.Matrix)

				// check position on map
				ch, ok := world.Matrix.GetVSafe(nextPosLooped)
				if !ok {
					// out of bounds
					panic("Out of bounds!")
				}

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

		//fmt.Printf("Step #%5d, size=%d\n", i+1, len(currentPositions))
		//printPositions(world.Matrix, world.Start, currentPositions)
	}

	return len(currentPositions)
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
