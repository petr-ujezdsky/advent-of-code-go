package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
	"sort"
)

type Elf = int

func FindMax(elves []Elf) int {
	max := math.MinInt
	for _, elf := range elves {
		max = utils.Max(max, elf)
	}

	return max
}

func FindTopThree(elves []Elf) int {
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	return elves[0] + elves[1] + elves[2]
}

func ParseInput(r io.Reader) []Elf {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var elves []Elf

	for scanner.Scan() {
		currentSum := 0

		for scanner.Text() != "" {
			currentSum += utils.ParseInt(scanner.Text())
			scanner.Scan()
		}

		elves = append(elves, currentSum)
	}

	return elves
}
