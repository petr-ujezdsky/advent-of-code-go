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

func play(seconds int, dimensions utils.Vector2i, robots []*Robot) map[utils.Vector2i]int {
	positions := make(map[utils.Vector2i]int)

	for _, robot := range robots {
		position := robot.Position.Add(robot.Velocity.Multiply(seconds))

		position.X = utils.ModFloor(position.X, dimensions.X)
		position.Y = utils.ModFloor(position.Y, dimensions.Y)

		positions[position]++
	}

	return positions
}

func DoWithInputPart01(world World) int {
	width := world.Dimensions.X
	height := world.Dimensions.Y

	positions := play(100, world.Dimensions, world.Robots)

	quadrants := createQuadrants(width, height)

	counts := countsInBounds(quadrants, positions)
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

func longestPath(positions map[utils.Vector2i]int) (int, map[utils.Vector2i]struct{}) {
	var maximum map[utils.Vector2i]struct{}
	usedTotal := make(map[utils.Vector2i]struct{})

	for position := range positions {
		used := make(map[utils.Vector2i]struct{})

		path := longestPathRecursive(position, positions, used, usedTotal)
		if path > len(maximum) {
			maximum = used
		}
	}

	return len(maximum), maximum
}

func longestPathRecursive(position utils.Vector2i, positions map[utils.Vector2i]int, used, usedTotal map[utils.Vector2i]struct{}) int {
	if _, ok := used[position]; ok {
		return 0
	}

	if _, ok := usedTotal[position]; ok {
		return 0
	}

	used[position] = struct{}{}
	usedTotal[position] = struct{}{}

	length := 1
	for _, step := range utils.Direction4Steps {
		neighbour := position.Add(step)
		if _, ok := positions[neighbour]; !ok {
			continue
		}

		length += longestPathRecursive(neighbour, positions, used, usedTotal)
	}

	return length
}

func countsInBounds(bounds []utils.BoundingRectangle, positions map[utils.Vector2i]int) []int {
	counts := make([]int, len(bounds))

	for position, count := range positions {
		for i, bound := range bounds {
			if bound.Contains(position) {
				counts[i] += count
			}
		}
	}

	return counts
}

func DoWithInputPart02(world World) int {
	width := world.Dimensions.X
	height := world.Dimensions.Y

	result := make(chan int)
	for seconds := 0; seconds < width*height+1; seconds++ {
		go func() {
			positions := play(seconds, world.Dimensions, world.Robots)

			if l, path := longestPath(positions); l > 100 {
				fmt.Printf("After %3d seconds (%3v) -----------------------------------------------------------------------------    \n", seconds, l)
				printRobots(world.Robots, nil, path, width, height)

				result <- seconds
			}
		}()

	}

	return <-result
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
