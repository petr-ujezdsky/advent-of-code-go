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

func play(seconds, width, height int, robots []*Robot) {
	for _, robot := range robots {
		robot.Position = robot.Position.Add(robot.Velocity.Multiply(seconds))

		robot.Position.X = utils.ModFloor(robot.Position.X, width)
		robot.Position.Y = utils.ModFloor(robot.Position.Y, height)
	}
}

func DoWithInputPart01(world World, width int, height int) int {
	play(100, width, height, world.Robots)

	quadrants := createQuadrants(width, height)

	counts := [4]int{}
	for _, robot := range world.Robots {
		for i, quadrant := range quadrants {
			if quadrant.Contains(robot.Position) {
				counts[i]++
			}
		}
	}
	return counts[0] * counts[1] * counts[2] * counts[3]
}

func createQuadrants(width, height int) []utils.BoundingRectangle {
	qw := width / 2
	qh := height / 2

	var quadrants []utils.BoundingRectangle

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			quadrant := utils.BoundingRectangle{
				Horizontal: utils.IntervalI{
					Low:  i * (qw + 1),
					High: i*(qw+1) + qw - 1,
				},
				Vertical: utils.IntervalI{
					Low:  j * (qh + 1),
					High: j*(qh+1) + qh - 1,
				},
			}

			quadrants = append(quadrants, quadrant)
		}
	}
	return quadrants
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
