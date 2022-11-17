package day_05

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type Point struct {
	X, Y int
}

type Line struct {
	A, B Point
}

func NewLine(x1, y1, x2, y2 int) Line {
	return Line{
		A: Point{x1, y1},
		B: Point{x2, y2},
	}
}

func Create2DSlice(x, y int) [][]int {
	matrixCols := make([][]int, x)
	cells := make([]int, x*y)

	for col := range matrixCols {
		matrixCols[col], cells = cells[:y], cells[y:]
	}

	return matrixCols
}

func CountIntersections(lines []Line) int {
	// find max values
	xMax := 0
	yMax := 0

	for _, line := range lines {
		xMax = utils.Max(utils.Max(xMax, line.A.X), line.B.X)
		yMax = utils.Max(utils.Max(yMax, line.A.Y), line.B.Y)
	}

	// create area matrix
	area := Create2DSlice(xMax+1, yMax+1)

	// draw in lines
	for _, line := range lines {
		if line.A.X == line.B.X {
			// vertical line
			for y := utils.Min(line.A.Y, line.B.Y); y <= utils.Max(line.A.Y, line.B.Y); y++ {
				area[line.A.X][y]++
			}
		} else if line.A.Y == line.B.Y {
			// horizontal line
			for x := utils.Min(line.A.X, line.B.X); x <= utils.Max(line.A.X, line.B.X); x++ {
				area[x][line.A.Y]++
			}
		}
	}

	overlaps := 0

	// count overlaps
	for _, col := range area {
		for _, count := range col {
			if count > 1 {
				overlaps++
			}
		}
	}

	return overlaps
}

func ParseInput(r io.Reader) ([]Line, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var lines []Line

	// each line has format
	// 964,133 -> 596,133
	for iRow := 0; scanner.Scan(); iRow++ {

		// split to two points
		points := strings.Split(scanner.Text(), " -> ")

		// split to first point coordinates
		startPoint, err := utils.ToInts(strings.Split(points[0], ","))
		if err != nil {
			return nil, err
		}

		// split to second point coordinates
		endPoint, err := utils.ToInts(strings.Split(points[1], ","))
		if err != nil {
			return nil, err
		}

		// create new line and append it to list
		line := NewLine(startPoint[0], startPoint[1], endPoint[0], endPoint[1])
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}
