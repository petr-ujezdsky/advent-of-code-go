package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"sort"
)

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

func countArrangements(targetValue1, targetValue2 int, sums2, sums3, sums4 []int, arrangementsCount int) int {
	i := 0
	for i < len(sums2) {
		for i < len(sums2) && sums2[i] != targetValue1 && sums2[i] != targetValue2 {
			i++
		}
		groupSize := 0
		iGroupStart := i
		for i < len(sums2) && (sums2[i] == targetValue1 || sums2[i] == targetValue2) {
			i++
			groupSize++
		}
		groupSize3 := 0
		j := iGroupStart
		for j < len(sums3) && (sums3[j] == 3) {
			j++
			groupSize3++
		}
		groupSize4 := 0
		j = iGroupStart
		for j < len(sums4) && (sums4[j] == 4) {
			j++
			groupSize4++
		}

		arrangementsCount *= (groupSize + groupSize3 + groupSize4) + 1
	}

	return arrangementsCount
}

func DoWithInput2(adapters []int) int {
	device := slices.Max(adapters) + 3

	adapters = append(adapters, device, 0)
	sort.Ints(adapters)

	differentials := slices.Differentials(adapters)
	fmt.Printf("Adapters:      %v\n", slices.Sprintf(adapters, "%2d"))
	fmt.Printf("Differentials:    %v\n", slices.Sprintf(differentials, "%2d"))

	arrangementsCount := 1

	sums2 := make([]int, len(differentials)-1)
	sums3 := make([]int, len(differentials)-2)
	sums4 := make([]int, len(differentials)-3)

	for i := 0; i < len(differentials)-1; i++ {
		sums2[i] = differentials[i] + differentials[i+1]

		if i < len(differentials)-2 {
			sums3[i] = differentials[i] + differentials[i+1] + differentials[i+2]
		}

		if i < len(differentials)-3 {
			sums4[i] = differentials[i] + differentials[i+1] + differentials[i+2] + differentials[i+3]
		}
	}
	fmt.Printf("Sums 2:           %v\n", slices.Sprintf(sums2, "%2d"))
	fmt.Printf("Sums 3:           %v\n", slices.Sprintf(sums3, "%2d"))
	fmt.Printf("Sums 4:           %v\n", slices.Sprintf(sums4, "%2d"))

	arrangementsCount = countArrangements(2, 3, sums2, sums3, sums4, arrangementsCount)

	return arrangementsCount
}

func ParseInput(r io.Reader) []int {
	return parsers.ParseToObjects(r, parsers.MapperIntegers)
}
