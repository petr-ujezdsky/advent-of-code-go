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
		_, right := ExtrapolateLeftRight(history)
		sum += right
	}

	return sum
}

func ExtrapolateLeftRight(values []int) (int, int) {
	derivative, allZeros := derive(values)
	if allZeros {
		return values[0], values[len(values)-1]
	}

	left, right := ExtrapolateLeftRight(derivative)
	return values[0] - left, values[len(values)-1] + right
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
	sum := 0

	for _, history := range world.Histories {
		left, _ := ExtrapolateLeftRight(history)
		sum += left
	}

	return sum
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) History {
		return utils.ExtractInts(str, true)
	}

	histories := parsers.ParseToObjects(r, parseItem)
	return World{Histories: histories}
}
