package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"sort"
)

func chain(adapter int, adapters map[int]struct{}, device int) (stepCounts map[int]int, steps []int, ok bool) {
	if adapter == device {
		stepCounts = make(map[int]int)
		steps = make([]int, 0, len(adapters)+2)
		return stepCounts, steps, true
	}

	for step := 1; step <= 3; step++ {
		nextAdapter := adapter + step
		if _, ok = adapters[nextAdapter]; ok || nextAdapter == device {
			stepCounts, steps, ok = chain(nextAdapter, adapters, device)
			if ok {
				stepCounts[step]++
				steps = append(steps, step)
				return stepCounts, steps, true
			}
		}
	}

	return nil, nil, false
}

func DoWithInput3(adapters []int) int {
	device := slices.Max(adapters) + 3
	set := slices.ToSet(adapters)

	stepCounts, _, ok := chain(0, set, device)
	if !ok {
		panic("No solution found")
	}

	return stepCounts[1] * stepCounts[3]
}

func DoWithInput(adapters []int) int {
	device := slices.Max(adapters) + 3

	adapters = append(adapters, device, 0)
	sort.Ints(adapters)

	differentials := slices.Differentials(adapters)
	stepCounts := make(map[int]int)

	for _, step := range differentials {
		stepCounts[step]++
	}

	//fmt.Println(slices.Reverse(steps))

	return stepCounts[1] * stepCounts[3]
}

func DoWithInput2(adapters []int) int {
	device := slices.Max(adapters) + 3

	adapters = append(adapters, device, 0)
	sort.Ints(adapters)

	differentials := slices.Differentials(adapters)
	fmt.Printf("Adapters:      %v\n", slices.Sprintf(adapters, "%2d"))
	fmt.Printf("Differentials:    %v\n", slices.Sprintf(differentials, "%2d"))

	arrangementsCount := 1
	skippable1s := 0
	sums2 := make([]int, len(differentials)-1)
	for i := range differentials[0 : len(differentials)-1] {
		sum := differentials[i] + differentials[i+1]
		if sum <= 3 {
			// skippable
			skippable1s++
			arrangementsCount *= 2
		}
		sums2[i] = sum
	}
	fmt.Printf("Sums 2:           %v\n", slices.Sprintf(sums2, "%2d"))

	skippable2s := 0
	sums3 := make([]int, len(differentials)-2)
	for i := range differentials[0 : len(differentials)-2] {
		sum := differentials[i] + differentials[i+1] + differentials[i+2]
		if sum == 3 {
			// skippable
			skippable2s++
			arrangementsCount *= 2
		}
		sums3[i] = sum
	}
	fmt.Printf("Sums 3:           %v\n", slices.Sprintf(sums3, "%2d"))

	return arrangementsCount
}

func ParseInput(r io.Reader) []int {
	return parsers.ParseToObjects(r, parsers.MapperIntegers)
}
