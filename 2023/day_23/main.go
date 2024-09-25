package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Item rune

type World struct {
	Matrix     matrix.Matrix[Item]
	Start, End utils.Vector2i
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(char rune) Item {
		return Item(char)
	}

	m := parsers.ParseToMatrix(r, parseItem)

	return World{
		Matrix: m,
		Start:  utils.Vector2i{X: 1},
		End:    utils.Vector2i{X: m.Width - 2, Y: m.Height - 1},
	}
}
