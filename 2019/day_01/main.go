package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	Weights []int
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, weight := range world.Weights {
		sum += weight/3 - 2
	}

	return sum
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	weights := parsers.ParseToObjects(r, parsers.MapperIntegers)
	return World{Weights: weights}
}
