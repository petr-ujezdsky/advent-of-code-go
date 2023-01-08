package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

func chain(adapter int, adapters map[int]struct{}, device int) (stepCounts map[int]int, ok bool) {
	if adapter == device {
		stepCounts = make(map[int]int)
		return stepCounts, true
	}

	for step := 1; step <= 3; step++ {
		nextAdapter := adapter + step
		if _, ok = adapters[nextAdapter]; ok || nextAdapter == device {
			stepCounts, ok = chain(nextAdapter, adapters, device)
			if ok {
				stepCounts[step]++
				return stepCounts, true
			}
		}
	}

	return nil, false
}

func DoWithInput(adapters []int) int {
	device := slices.Max(adapters) + 3
	set := slices.ToSet(adapters)

	stepCounts, ok := chain(0, set, device)
	if !ok {
		panic("No solution found")
	}

	return stepCounts[1] * stepCounts[3]
}

func ParseInput(r io.Reader) []int {
	return parsers.ParseToObjects(r, parsers.MapperIntegers)
}
