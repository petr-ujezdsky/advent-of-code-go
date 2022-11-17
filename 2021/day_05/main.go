package day_05

import (
	"bufio"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strconv"
	"strings"
)

type Matrix2 [][]int

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

func NewMatrix2(x, y int) Matrix2 {
	matrixCols := make([][]int, x)
	cells := make([]int, x*y)

	for col := range matrixCols {
		matrixCols[col], cells = cells[:y], cells[y:]
	}

	return matrixCols
}

func (matrix Matrix2) String() string {
	m := len(matrix)
	n := len(matrix[0])

	var sb strings.Builder

	for y := 0; y < n; y++ {
		for x := 0; x < m; x++ {
			val := matrix[x][y]

			if val == 0 {
				sb.WriteString(" .")
			} else {
				sb.WriteString(" ")
				sb.WriteString(strconv.Itoa(val))
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (matrix Matrix2) Print() {
	fmt.Println(matrix.String())
}

func CountIntersections(lines []Line, includeDiagonals bool) (int, Matrix2) {
	// find max values
	xMax := 0
	yMax := 0

	for _, line := range lines {
		xMax = utils.Max(utils.Max(xMax, line.A.X), line.B.X)
		yMax = utils.Max(utils.Max(yMax, line.A.Y), line.B.Y)
	}

	// create area matrix
	area := NewMatrix2(xMax+1, yMax+1)

	// draw in lines
	for _, line := range lines {
		if includeDiagonals || (line.A.X == line.B.X || line.A.Y == line.B.Y) {
			dirX := utils.Signum(line.B.X - line.A.X)
			dirY := utils.Signum(line.B.Y - line.A.Y)

			x := line.A.X
			y := line.A.Y

			for {
				area[x][y]++
				x += dirX
				y += dirY

				if dirX != 0 && x == line.B.X+dirX || dirY != 0 && y == line.B.Y+dirY {
					break
				}
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

	return overlaps, area
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
