package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type History []int

type World struct {
	Histories []History
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, history := range world.Histories {
		nextValue := Extrapolate(history)
		sum += nextValue
	}

	return sum
}

func Extrapolate(values []int) int {
	return extrapolateInner(values, 0)
}

func extrapolateInner(values []int, depth int) int {
	derivative, allZeros := derive(values)
	if allZeros {
		return values[len(values)-1]
	}

	added := extrapolateInner(derivative, depth+1)
	return values[len(values)-1] + added
}

func derive(values []int) ([]int, bool) {
	derivative := make([]int, len(values)-1)

	allZeros := true
	previous := values[0]
	for i, current := range values[1:] {
		diff := current - previous
		derivative[i] = diff

		allZeros = allZeros && diff == 0
		previous = current
	}

	return derivative, allZeros
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) History {
		return utils.ExtractInts(str, true)
	}

	histories := parsers.ParseToObjects(r, parseItem)
	return World{Histories: histories}
}
