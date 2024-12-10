package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"regexp"
)

type World struct {
	Matrix matrix.Matrix[rune]
}

var regex1 = regexp.MustCompile(`XMAS`)
var regex2 = regexp.MustCompile(`SAMX`)

func toColumns(m matrix.Matrix[rune]) []string {
	columns := make([]string, m.Width)

	for i, column := range m.Columns {
		columns[i] = string(column)
	}

	return columns
}

func rotate45(m matrix.Matrix[rune]) []string {
	var rows []string
	step := utils.Vector2i{X: 1, Y: 1}

	startPosition := utils.Vector2i{X: m.Width - 1, Y: 0}

	for index := 0; index < m.Width+m.Height-1; index++ {
		var row []rune

		pos := startPosition
		for {
			char, ok := m.GetVSafe(pos)
			if !ok {
				break
			}

			row = append(row, char)
			pos = pos.Add(step)
		}

		if startPosition.X > 0 {
			startPosition.X--
		} else {
			startPosition.Y++
		}

		rows = append(rows, string(row))
	}

	return rows
}

func countXmas(rows []string) int {
	count := 0

	for _, row := range rows {
		count += len(regex1.FindAllString(row, -1))
		count += len(regex2.FindAllString(row, -1))
	}

	return count
}

func DoWithInputPart01(world World) int {
	totalCount := 0
	m := world.Matrix

	// columns
	totalCount += countXmas(toColumns(m))

	// diag
	totalCount += countXmas(rotate45(m))

	// rotate 90
	m = m.Rotate90CounterClockwise(1)

	// rows
	totalCount += countXmas(toColumns(m))

	// diag
	totalCount += countXmas(rotate45(m))

	return totalCount
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(char rune) rune {
		return char
	}

	return World{Matrix: parsers.ParseToMatrix(r, parseItem)}
}
