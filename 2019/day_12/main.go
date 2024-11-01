package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Moon struct {
	Position, Velocity utils.Vector3i
}

type World struct {
	Moons []*Moon
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) *Moon {
		numbers := utils.ExtractInts(str, true)

		return &Moon{
			Position: utils.Vector3i{X: numbers[0], Y: numbers[1], Z: numbers[2]},
			Velocity: utils.Vector3i{},
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Moons: items}
}
