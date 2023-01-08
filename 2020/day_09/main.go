package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

func DoWithInput(numbers []int) int {
	return len(numbers)
}

func ParseInput(r io.Reader) []int {
	return parsers.ParseToObjects(r, parsers.MapperIntegers)
}
