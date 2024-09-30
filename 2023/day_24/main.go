package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/equations"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strings"
)

type Line struct {
	Position   utils.Vector3i
	Direction  utils.Vector3i
	Position2  utils.Vector2i
	Direction2 utils.Vector2i
}

type World struct {
	Lines []Line
}

func findCrossing2D(line1, line2 Line) (utils.Vector2f, bool) {
	A := matrix.NewMatrixNumberRowNotation([][]float64{
		{float64(line1.Direction.X), float64(-line2.Direction.X)},
		{float64(line1.Direction.Y), float64(-line2.Direction.Y)},
	})

	b := utils.VectorNf{Items: []float64{
		float64(line2.Position.X - line1.Position.X),
		float64(line2.Position.Y - line1.Position.Y),
	}}

	ts, ok := equations.SolveLinearEquations(A, b)
	if !ok {
		return utils.Vector2f{}, false
	}

	if ts.Items[0] < 0 || ts.Items[1] < 0 {
		// intersection in past -> do not use
		return utils.Vector2f{}, false
	}

	x := utils.Vector2f{
		X: float64(line1.Position.X) + ts.Items[0]*float64(line1.Direction.X),
		Y: float64(line1.Position.Y) + ts.Items[0]*float64(line1.Direction.Y),
	}

	return x, true
}

func withinBounds(v utils.Vector2f, low, high float64) bool {
	if v.X > high || v.Y > high {
		return false
	}

	if v.X < low || v.Y < low {
		return false
	}

	return true
}

func DoWithInputPart01(world World, low, high float64) int {
	lines := world.Lines
	count := 0

	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			line1 := lines[i]
			line2 := lines[j]

			cross, ok := findCrossing2D(line1, line2)
			if !ok {
				continue
			}

			if withinBounds(cross, low, high) {
				count++
			}
		}
	}

	return count
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) Line {
		parts := strings.Split(str, " @ ")
		position := utils.ToIntsP(trim(strings.Split(parts[0], ", ")))
		direction := utils.ToIntsP(trim(strings.Split(parts[1], ", ")))

		return Line{
			Position: utils.Vector3i{
				X: position[0],
				Y: position[1],
				Z: position[2],
			},
			Direction: utils.Vector3i{
				X: direction[0],
				Y: direction[1],
				Z: direction[2],
			},
			Position2: utils.Vector2i{
				X: position[0],
				Y: position[1],
			},
			Direction2: utils.Vector2i{
				X: direction[0],
				Y: direction[1],
			},
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Lines: items}
}

func trim(rows []string) []string {
	trimmed := make([]string, len(rows))

	for i, row := range rows {
		trimmed[i] = strings.TrimSpace(row)
	}

	return trimmed
}
