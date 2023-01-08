package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Window = map[int]struct{}

func isValid(i int, window Window) bool {
	for windowNumber := range window {
		remainder := i - windowNumber

		if remainder == windowNumber {
			// pair consists of 2 different numbers
			continue
		}

		// check remainder existence
		if _, ok := window[remainder]; ok {
			return true
		}
	}

	return false
}

func FindInvalidNumber(numbers []int, windowSize int) int {
	// init window
	window := make(Window, windowSize)
	for i := 0; i < windowSize; i++ {
		window[numbers[i]] = struct{}{}
	}

	for i := windowSize; i < len(numbers); i++ {
		number := numbers[i]
		if !isValid(number, window) {
			return number
		}

		// remove the left-most number from window
		delete(window, numbers[i-windowSize])
		// add current number
		window[number] = struct{}{}
	}

	return len(numbers)
}

func ParseInput(r io.Reader) []int {
	return parsers.ParseToObjects(r, parsers.MapperIntegers)
}
