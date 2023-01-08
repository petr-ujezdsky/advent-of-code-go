package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"math"
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

func FindRange(numbers []int, windowSize int) int {
	targetSum := FindInvalidNumber(numbers, windowSize)

	sum := 0
	iFirst, iLast := 0, 0

	for i, number := range numbers {
		nextSum := sum + number
		if nextSum <= targetSum {
			iLast = i
			sum = nextSum
			continue
		}

		nextIFirst := iFirst
		for nextSum > targetSum {
			nextSum -= numbers[nextIFirst]
			nextIFirst++

			if nextIFirst == iLast {
				break
			}
		}

		sum = nextSum
		iFirst = nextIFirst
		iLast = i

		if sum == targetSum {
			min := math.MaxInt
			max := math.MinInt
			for _, subNumber := range numbers[iFirst : iLast+1] {
				min = utils.Min(min, subNumber)
				max = utils.Max(max, subNumber)
			}
			return min + max
		}
	}

	panic("No solution found")
}

func ParseInput(r io.Reader) []int {
	return parsers.ParseToObjects(r, parsers.MapperIntegers)
}
