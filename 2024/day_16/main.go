package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	Matrix     matrix.Matrix[string]
	Start, End utils.Vector2i
}

func DoWithInputPart01(world World) int {

	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	var start utils.Vector2i
	var end utils.Vector2i

	parseItem := func(char rune, x, y int) string {
		if char == 'S' {
			start = utils.Vector2i{X: x, Y: y}
		} else if char == 'E' {
			end = utils.Vector2i{X: x, Y: y}
		}

		return string(char)
	}

	return World{
		Matrix: parsers.ParseToMatrixIndexed(r, parseItem),
		Start:  start,
		End:    end,
	}
}
