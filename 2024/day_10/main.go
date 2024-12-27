package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	Matrix     matrix.Matrix[int]
	TrailHeads []utils.Vector2i
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	var traileheads []utils.Vector2i

	parseItem := func(char rune, x, y int) int {
		if char == '0' {
			traileheads = append(traileheads, utils.Vector2i{X: x, Y: y})
		}
		return utils.ParseInt(string(char))
	}

	return World{
		Matrix:     parsers.ParseToMatrixIndexed(r, parseItem),
		TrailHeads: traileheads,
	}
}
