package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strconv"
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

	counts := countsInBounds(quadrants, world.Robots)
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

func printRobots(robots []*Robot, bounds []utils.BoundingRectangle, width int, height int) {
	m := matrix.NewMatrix[int](width, height)

	for _, robot := range robots {
		pos := robot.Position

		m.SetV(pos, m.GetV(pos)+1)
	}

	formatter := func(i, x, y int) string {
		if i > 9 {
			panic("Too much")
		}

		for _, bound := range bounds {
			if bound.Contains(utils.Vector2i{X: x, Y: y}) {
				return "x"
			}
		}

		if i == 0 {
			return " "
		}
		return strconv.Itoa(i)
	}

	fmt.Println(matrix.StringFmtSeparatorIndexed(m, false, "", formatter))
}

func countsInBounds(bounds []utils.BoundingRectangle, robots []*Robot) []int {
	counts := make([]int, len(bounds))

	for _, robot := range robots {
		for i, bound := range bounds {
			if bound.Contains(robot.Position) {
				counts[i]++
			}
		}
	}

	return counts
}

func DoWithInputPart02(world World, width int, height int) int {
	qw := 21
	qh := 7
	corners := []utils.BoundingRectangle{
		{
			Horizontal: utils.IntervalI{High: qw - 1},
			Vertical:   utils.IntervalI{High: qh - 1},
		},
		{
			Horizontal: utils.IntervalI{Low: width - qw, High: width - 1},
			Vertical:   utils.IntervalI{High: qh - 1},
		},
	}

	seconds := 0
	for count := 0; count < 3*10; {
		counts := countsInBounds(corners, world.Robots)
		if counts[0]+counts[1] == 0 {
			count++
			fmt.Printf("After %3d seconds -----------------------------------------------------------------------------------\n", seconds)
			printRobots(world.Robots, corners, width, height)
		}
		play(1, width, height, world.Robots)
		seconds++
		if seconds%1_000_000 == 0 {
			fmt.Printf("%v. second...\n", seconds)
		}
	}

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
