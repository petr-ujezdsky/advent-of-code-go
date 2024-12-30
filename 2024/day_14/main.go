package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Robot struct {
	Position, Velocity utils.Vector2i
}

type World struct {
	Robots []*Robot
}

func play(int, []*Robot) {

}

func DoWithInputPart01(world World, width int, height int) int {
	play(100, world.Robots)
	return 0
}

func DoWithInputPart02(world World, width int, height int) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) *Robot {
		ints := utils.ExtractInts(str, true)

		return &Robot{
			Position: utils.Vector2i{X: ints[0], Y: ints[1]},
			Velocity: utils.Vector2i{X: ints[2], Y: ints[3]},
		}
	}

	robots := parsers.ParseToObjects(r, parseItem)
	return World{Robots: robots}
}
