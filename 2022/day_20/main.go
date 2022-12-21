package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type PositionedNumber struct {
	Value, Index int
}

func MixNumbers(numbers []PositionedNumber) int {
	return len(numbers)
}

func ParseInput(r io.Reader) []PositionedNumber {
	ints := utils.ParseToIntsP(r)

	var numbers []PositionedNumber

	for i, v := range ints {
		numbers = append(numbers, PositionedNumber{
			Value: v,
			Index: i,
		})
	}

	return numbers
}
