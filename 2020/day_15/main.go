package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type History = map[int]int

func DoWithInputPart01(numbers []int) int {
	history := make(History)

	for i, number := range numbers[0 : len(numbers)-1] {
		history[number] = i + 1
	}

	currentNumber := numbers[len(numbers)-1]

	for i := len(numbers); i < 2020; i++ {
		var nextNumber int
		// lookup in history
		if previousPos, ok := history[currentNumber]; ok {
			nextNumber = i - previousPos
		} else {
			nextNumber = 0
		}

		// store last number
		history[currentNumber] = i

		currentNumber = nextNumber
	}

	return currentNumber
}

func DoWithInputPart02(numbers []int) int {
	return len(numbers)
}

func ParseInput(r io.Reader) []int {
	parseItem := func(str string) []int {
		return utils.ExtractInts(str, false)
	}

	return parsers.ParseToObjects(r, parseItem)[0]
}
