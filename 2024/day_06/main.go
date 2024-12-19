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

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	start := utils.Vector2i{}

	mapper := func(char rune, x, y int) rune {
		if char == '^' {
			start = utils.Vector2i{X: x, Y: y}
		}

		return char
	}

	return World{
		Matrix: parsers.ParseToMatrixIndexed(r, mapper),
		Start:  start,
	}
}
