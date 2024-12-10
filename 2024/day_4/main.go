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
}

var patterns01 = []matrix.Matrix[rune]{
	// horizontal
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'X', 'M', 'A', 'S'},
	}),
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'S', 'A', 'M', 'X'},
	}),

	// vertical
	matrix.NewMatrixColumnNotation[rune]([][]rune{
		{'X', 'M', 'A', 'S'},
	}),
	matrix.NewMatrixColumnNotation[rune]([][]rune{
		{'S', 'A', 'M', 'X'},
	}),

	// diag 1
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'X', '.', '.', '.'},
		{'.', 'M', '.', '.'},
		{'.', '.', 'A', '.'},
		{'.', '.', '.', 'S'},
	}),
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'S', '.', '.', '.'},
		{'.', 'A', '.', '.'},
		{'.', '.', 'M', '.'},
		{'.', '.', '.', 'X'},
	}),

	// diag 2
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'.', '.', '.', 'S'},
		{'.', '.', 'A', '.'},
		{'.', 'M', '.', '.'},
		{'X', '.', '.', '.'},
	}),
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'.', '.', '.', 'X'},
		{'.', '.', 'M', '.'},
		{'.', 'A', '.', '.'},
		{'S', '.', '.', '.'},
	}),
}

var patterns02 = []matrix.Matrix[rune]{
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'M', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'S'},
	}),
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'M', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'S'},
	}),
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'S', '.', 'M'},
		{'.', 'A', '.'},
		{'S', '.', 'M'},
	}),
	matrix.NewMatrixRowNotation[rune]([][]rune{
		{'S', '.', 'S'},
		{'.', 'A', '.'},
		{'M', '.', 'M'},
	}),
}

func matchesXmas(m, pattern matrix.Matrix[rune], pos utils.Vector2i) bool {
	for x, column := range pattern.Columns {
		for y, valueExpected := range column {
			if valueExpected == '.' {
				continue
			}

			valueActual, ok := m.GetVSafe(pos.Add(utils.Vector2i{X: x, Y: y}))
			if !ok || valueActual != valueExpected {
				return false
			}
		}
	}

	return true
}

func countXmas(m matrix.Matrix[rune], patterns []matrix.Matrix[rune]) int {
	count := 0

	for x, column := range m.Columns {
		for y := range column {
			pos := utils.Vector2i{X: x, Y: y}

			for _, pattern := range patterns {

				if matchesXmas(m, pattern, pos) {
					count++
				}
			}
		}
	}

	return count
}

func DoWithInputPart01(world World) int {
	return countXmas(world.Matrix, patterns01)
}

func DoWithInputPart02(world World) int {
	return countXmas(world.Matrix, patterns02)
}

func ParseInput(r io.Reader) World {
	parseItem := func(char rune) rune {
		return char
	}

	return World{Matrix: parsers.ParseToMatrix(r, parseItem)}
}
