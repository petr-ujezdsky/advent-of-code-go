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
