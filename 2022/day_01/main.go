package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
)

type Elf = int

func FindMax(elves []Elf) int {
	max := math.MinInt
	for _, elf := range elves {
		max = utils.Max(max, elf)
	}

	return max
}

func ParseInput(r io.Reader) []Elf {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var elves []Elf

	currentSum := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, currentSum)
			currentSum = 0
		} else {
			currentSum += utils.ParseInt(scanner.Text())
		}
	}

	return elves
}
