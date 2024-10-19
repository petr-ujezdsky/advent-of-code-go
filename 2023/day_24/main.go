package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/equations"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"math"
	"strings"
)

type Line struct {
	Position  utils.Vector3i
	Direction utils.Vector3i

	PositionF  utils.Vector3f
	DirectionF utils.Vector3f
}

func (line Line) ToStringMath3D() string {
	return fmt.Sprintf("[%d%+d*t, %d%+d*t, %d%+d*t]",
		line.Position.X, line.Direction.X,
		line.Position.Y, line.Direction.Y,
		line.Position.Z, line.Direction.Z)
}

func (line Line) ToStringInput() string {
	return fmt.Sprintf("%d, %d, %d @ %d, %d, %d",
		line.Position.X, line.Position.Y, line.Position.Z,
		line.Direction.X, line.Direction.Y, line.Direction.Z)
}

func (line Line) PositionAtTime(time int) utils.Vector3i {
	return line.Position.Add(line.Direction.Multiply(time))
}

type World struct {
	Lines []Line
}

func findCrossing2D(line1, line2 Line) (utils.Vector2f, bool) {
	x, t1, t2, ok := findCrossing2DNoTimeRestriction(line1, line2)
	if !ok {
		return utils.Vector2f{}, false
	}

	if t1 < 0 || t2 < 0 {
		// intersection in past -> do not use
		return utils.Vector2f{}, false
	}

	return x, true
}

func findCrossing2DNoTimeRestriction(line1, line2 Line) (utils.Vector2f, float64, float64, bool) {
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
		return utils.Vector2f{}, 0, 0, false
	}

	x := utils.Vector2f{
		X: float64(line1.Position.X) + ts.Items[0]*float64(line1.Direction.X),
		Y: float64(line1.Position.Y) + ts.Items[0]*float64(line1.Direction.Y),
	}

	return x, ts.Items[0], ts.Items[1], true
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

func extractVariables(index int, input utils.VectorNf) (p, v utils.Vector3f, t float64) {
	p = utils.Vector3f{
		X: input.Items[0],
		Y: input.Items[1],
		Z: input.Items[2],
	}

	v = utils.Vector3f{
		X: input.Items[3],
		Y: input.Items[4],
		Z: input.Items[5],
	}

	t = input.Items[6+index]

	return
}

func fLine(line Line, lineIndex int, input, output utils.VectorNf) (utils.VectorNf, int) {
	p0, v0, ti := extractVariables(lineIndex, input)

	// pi - p0 + ti * (vi - v0)
	result := line.PositionF.Add(p0.Multiply(-1)).Add(line.DirectionF.Add(v0.Multiply(-1)).Multiply(ti))

	output.Items[3*lineIndex+0] = result.X
	output.Items[3*lineIndex+1] = result.Y
	output.Items[3*lineIndex+2] = result.Z

	return output, lineIndex + 1
}

func jLine(line Line, lineIndex int, input utils.VectorNf, output [][]float64) ([][]float64, int) {
	_, v0, ti := extractVariables(lineIndex, input)

	derivationByTi := line.DirectionF.Add(v0.Multiply(-1))

	rowX := make([]float64, 9)
	rowX[0] = -1
	rowX[3] = -ti
	rowX[6+lineIndex] = derivationByTi.X

	rowY := make([]float64, 9)
	rowY[1] = -1
	rowY[4] = -ti
	rowY[6+lineIndex] = derivationByTi.Y

	rowZ := make([]float64, 9)
	rowZ[2] = -1
	rowZ[5] = -ti
	rowZ[6+lineIndex] = derivationByTi.Z

	output[3*lineIndex+0] = rowX
	output[3*lineIndex+1] = rowY
	output[3*lineIndex+2] = rowZ

	return output, lineIndex + 1
}

func DoWithInputPart02(world World) int {
	line1 := world.Lines[0]
	line2 := world.Lines[1]
	line3 := world.Lines[2]

	F := func(input utils.VectorNf) utils.VectorNf {
		// 9 = 3 (p) + 3 (v) + 3*1 (t)
		output := utils.VectorNf{Items: make([]float64, 9)}
		index := 0

		output, index = fLine(line1, index, input, output)
		output, index = fLine(line2, index, input, output)
		output, index = fLine(line3, index, input, output)

		return output
	}

	J := func(input utils.VectorNf) matrix.MatrixFloat {
		rows := make([][]float64, 9)
		index := 0

		rows, index = jLine(line1, index, input, rows)
		rows, index = jLine(line2, index, input, rows)
		rows, index = jLine(line3, index, input, rows)

		Jxi := matrix.NewMatrixNumberRowNotation[float64](rows)

		//fmt.Println(Jxi.StringFmt(matrix.FmtFmt[float64]("%9.6f")))
		//fmt.Println(Jxi.StringFmt(matrix.FmtFmt[float64]("%3.0f")))
		//fmt.Println()

		return Jxi
	}

	x0 := utils.VectorNf{Items: []float64{
		// p
		line1.PositionF.X + line1.DirectionF.X,
		line1.PositionF.Y + line1.DirectionF.Y,
		line1.PositionF.Z + line1.DirectionF.Z,

		// v
		-9,
		5,
		8,

		// t
		1,
		5,
		15,
	}}

	solution, ok := equations.SolveNonLinearEquationsThreshold(F, J, x0, 0.1, 10_000)
	if !ok {
		return -1
	}

	p, _, _ := extractVariables(0, solution)

	return int(math.Round(p.X)) + int(math.Round(p.Y)) + int(math.Round(p.Z))
}

func printLinesMath3D(lines []Line) {
	for _, line := range lines {
		fmt.Println(line.ToStringMath3D())
	}
}

func printLinesInput(lines []Line) {
	offset := lines[0].Position.Multiply(-1)

	for _, line := range lines {
		line.Position = line.Position.Add(offset)
		fmt.Println(line.ToStringInput())
	}
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

			PositionF: utils.Vector3f{
				X: float64(position[0]),
				Y: float64(position[1]),
				Z: float64(position[2]),
			},
			DirectionF: utils.Vector3f{
				X: float64(direction[0]),
				Y: float64(direction[1]),
				Z: float64(direction[2]),
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
