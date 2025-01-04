package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
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
		//return strconv.Itoa(i)
		return "+"
	}

	fmt.Println(matrix.StringFmtSeparatorIndexed(m, false, "", formatter))
}

func findPeriod(robot *Robot, dimensions utils.Vector2i) int {
	initialPosition := robot.Position
	for i := 0; i < 5_000_000; i++ {
		play(1, dimensions, []*Robot{robot})
		if robot.Position == initialPosition {
			return i + 1
		}
	}

	return -1
}

func isSymmetric(max int, robots []*Robot, width, height int) bool {
	m := matrix.NewMatrix[int](width, height)

	for _, robot := range robots {
		pos := robot.Position

		//m.SetV(pos, m.GetV(pos)+1)
		m.SetV(pos, 1)
	}

	mFlipped := m.FlipHorizontal()

	// check equality
	mismatched := 0
	for x, column := range m.Columns {
		for y, value1 := range column {
			value2 := mFlipped.Get(x, y)

			if value1 != value2 {
				mismatched++
			}

			if mismatched > max {
				return false
			}
		}
	}

	return true
}

func longestPath(robots []*Robot) int {
	positions := make(map[utils.Vector2i]struct{})

	for _, robot := range robots {
		positions[robot.Position] = struct{}{}
	}

	mx := -1

	for {
		p := maps.FirstKey(positions)
		path := longestPathRecursive(p, positions)
		mx = utils.Max(path, mx)
		if len(positions) == 0 {
			break
		}
	}

	return mx
}

func longestPathRecursive(position utils.Vector2i, positions map[utils.Vector2i]struct{}) int {
	mx := -1
	for _, step := range utils.Direction4Steps {
		neighbour := position.Add(step)
		path := 0
		if _, ok := positions[neighbour]; ok {
			delete(positions, neighbour)
			path += 1 + longestPathRecursive(neighbour, positions)
		}

		neighbour = position.Subtract(step)

		if _, ok := positions[neighbour]; ok {
			delete(positions, neighbour)
			path += longestPathRecursive(neighbour, positions)
		}

		mx = utils.Max(path, mx)
	}

	return mx
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
	//for i, robot := range world.Robots {
	//	period := findPeriod(robot, world.Dimensions)
	//	fmt.Printf("Period #%3d: %v\n", i, period)
	//}

	width := world.Dimensions.X
	height := world.Dimensions.Y

	//qw := 10
	//qh := 15
	//bounds := []utils.BoundingRectangle{
	//	{
	//		Horizontal: utils.IntervalI{High: qw - 1},
	//		Vertical:   utils.IntervalI{High: qh - 1},
	//	},
	//	{
	//		Horizontal: utils.IntervalI{Low: width - qw, High: width - 1},
	//		Vertical:   utils.IntervalI{High: qh - 1},
	//	},
	//}

	//bounds := []utils.BoundingRectangle{
	//	{
	//		Horizontal: utils.IntervalI{High: width/2 - 1},
	//		Vertical:   utils.IntervalI{High: height - 1},
	//	},
	//	{
	//		Horizontal: utils.IntervalI{Low: width/2 + 1, High: width - 1},
	//		Vertical:   utils.IntervalI{High: height - 1},
	//	},
	//}

	//printRobots(world.Robots, corners, width, height)
	//play(10403, world.Dimensions, world.Robots)
	//fmt.Printf("After %3d seconds -----------------------------------------------------------------------------------\n", 10403)
	//printRobots(world.Robots, corners, width, height)

	//bounds := createQuadrants(width, height)

	for seconds := 0; seconds < width*height+1; seconds++ {
		//counts := countsInBounds(bounds, world.Robots)
		//if counts[0]+counts[1] == 0 {

		//diff := 0
		//if utils.Abs(counts[0]-counts[2]) == diff && utils.Abs(counts[1]-counts[3]) == diff && counts[0] < counts[1] {
		//	//if counts[0] == counts[1] {
		//	fmt.Printf("After %5d seconds ---------------------------------------------------------------------------------\n", seconds)
		//	printRobots(world.Robots, nil, width, height)
		//}

		//if isSymmetric(1000, world.Robots, width, height) {
		//	fmt.Printf("After %3d seconds -----------------------------------------------------------------------------------\n", seconds)
		//	printRobots(world.Robots, nil, width, height)
		//}

		if longestPath(world.Robots) > 2 {
			fmt.Printf("After %3d seconds -----------------------------------------------------------------------------------\n", seconds)
			printRobots(world.Robots, nil, width, height)
		}

		play(1, world.Dimensions, world.Robots)
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
