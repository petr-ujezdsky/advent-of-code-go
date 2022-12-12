package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Vector2i = utils.Vector2i
type MatrixInt = utils.MatrixInt

type World struct {
	Heights    MatrixInt
	Start, End Vector2i
}

func DoWithInput(items World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rows [][]int

	start, end := Vector2i{}, Vector2i{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []int

		for x, digitAscii := range []rune(line) {
			if digitAscii == 'S' {
				digitAscii = 'a'
				start = Vector2i{x, y}
			}

			if digitAscii == 'E' {
				digitAscii = 'z'
				end = Vector2i{x, y}
			}

			digit := int(digitAscii) - int('a')
			row = append(row, digit)
		}

		rows = append(rows, row)
		y++
	}

	heights := utils.NewMatrixNumberRowNotation(rows)

	return World{
		Heights: heights,
		Start:   start,
		End:     end,
	}
}
