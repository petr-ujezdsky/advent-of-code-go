package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Antenna struct {
	Position utils.Vector2i
	Name     string
}

type World struct {
	Antennae []Antenna
	Matrix   matrix.Matrix[rune]
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	var antennae []Antenna

	parseItem := func(char rune, x, y int) rune {
		if char != '.' {
			antennae = append(antennae, Antenna{
				Position: utils.Vector2i{X: x, Y: y},
				Name:     string(char),
			})
		}

		return char
	}

	return World{
		Antennae: antennae,
		Matrix:   parsers.ParseToMatrixIndexed(r, parseItem),
	}
}
