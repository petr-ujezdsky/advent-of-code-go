package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	Rows []string
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, row := range world.Rows {
		sum += extractNumber(row)
	}

	return sum
}

func extractNumber(row string) int {
	var first, last rune

	for _, char := range []rune(row) {
		if char >= '0' && char <= '9' {
			if first == 0 {
				first = char
			}
			last = char
		}
	}

	return utils.ParseInt(string([]rune{first, last}))
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	items := parsers.ParseToStrings(r)
	return World{Rows: items}
}
