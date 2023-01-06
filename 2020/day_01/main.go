package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

func DoWithInput(numbers []int) int {
	for i, number1 := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			number2 := numbers[j]
			if number1+number2 == 2020 {
				return number1 * number2
			}
		}
	}

	panic("Not found")
}

func ParseInput(r io.Reader) []int {
	return utils.ParseToIntsP(r)
}
