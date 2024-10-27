package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	Chars []rune
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) []rune {
		return []rune(str)
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Chars: items[0]}
}
