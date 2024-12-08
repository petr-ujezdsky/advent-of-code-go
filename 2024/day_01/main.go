package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"sort"
)

type Item struct {
	Left, Right int
}

type World struct {
	Items       []Item
	Left, Right []int
}

func DoWithInputPart01(world World) int {
	left := slices.Clone(world.Left)
	sort.Ints(left)

	right := slices.Clone(world.Right)
	sort.Ints(right)

	totalDistance := 0
	for i, leftItem := range left {
		rightItem := right[i]

		totalDistance += utils.Abs(rightItem - leftItem)
	}

	return totalDistance
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var left []int
	var right []int

	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)

		left = append(left, ints[0])
		right = append(right, ints[1])
	}

	return World{Left: left, Right: right}
}
