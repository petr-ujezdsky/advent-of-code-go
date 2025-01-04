package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Robot struct {
	Position, Velocity utils.Vector2i
}

type World struct {
	Robots     []*Robot
	Dimensions utils.Vector2i
}

func play(seconds int, dimensions utils.Vector2i, robots []*Robot) {
	for _, robot := range robots {
		robot.Position = robot.Position.Add(robot.Velocity.Multiply(seconds))

		robot.Position.X = utils.ModFloor(robot.Position.X, dimensions.X)
		robot.Position.Y = utils.ModFloor(robot.Position.Y, dimensions.Y)
	}
}

func DoWithInputPart01(world World) int {
	width := world.Dimensions.X
	height := world.Dimensions.Y

	play(100, world.Dimensions, world.Robots)

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

func printRobots(robots []*Robot, bounds []utils.BoundingRectangle, path map[utils.Vector2i]struct{}, width int, height int) {
	m := matrix.NewMatrix[int](width, height)

	for _, robot := range robots {
		pos := robot.Position

		m.SetV(pos, m.GetV(pos)+1)
	}

	formatter := func(i, x, y int) string {
		if i > 9 {
			panic("Too much")
		}

		pos := utils.Vector2i{X: x, Y: y}

		for _, bound := range bounds {
			if bound.Contains(pos) {
				return "x"
			}
		}

		if _, ok := path[pos]; ok {
			return "#"
		}

		if i == 0 {
			return " "
		}
		//return strconv.Itoa(i)
		return "+"
	}

	fmt.Println(matrix.StringFmtSeparatorIndexed(m, false, "", formatter))
}

func longestPath(robots []*Robot) (int, map[utils.Vector2i]struct{}) {
	positions := make(map[utils.Vector2i]struct{})

	for _, robot := range robots {
		positions[robot.Position] = struct{}{}
	}

	var maximum map[utils.Vector2i]struct{}

	for position := range positions {
		used := make(map[utils.Vector2i]struct{})

		path := longestPathRecursive(position, positions, used)
		if path > len(maximum) {
			maximum = used
		}
	}

	return len(maximum), maximum
}

func longestPathRecursive(position utils.Vector2i, positions map[utils.Vector2i]struct{}, used map[utils.Vector2i]struct{}) int {
	if _, ok := used[position]; ok {
		return 0
	}

	used[position] = struct{}{}

	length := 1
	for _, step := range utils.Direction4Steps {
		neighbour := position.Add(step)
		if _, ok := positions[neighbour]; !ok {
			continue
		}

		length += longestPathRecursive(neighbour, positions, used)
	}

	return length
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

func DoWithInputPart02(world World) int {
	width := world.Dimensions.X
	height := world.Dimensions.Y

	for seconds := 0; seconds < width*height+1; seconds++ {
		if l, path := longestPath(world.Robots); l > 100 {
			fmt.Printf("After %3d seconds (%3v) -----------------------------------------------------------------------------    \n", seconds, l)
			printRobots(world.Robots, nil, path, width, height)

			return seconds
		}

		play(1, world.Dimensions, world.Robots)
	}

	panic("No tree found")
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

	var dimensions utils.Vector2i
	if len(robots) < 100 {
		dimensions = utils.Vector2i{X: 11, Y: 7}
	} else {
		dimensions = utils.Vector2i{X: 101, Y: 103}

	}

	return World{
		Robots:     robots,
		Dimensions: dimensions,
	}
}
