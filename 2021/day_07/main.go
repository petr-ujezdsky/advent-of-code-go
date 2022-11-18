package day_07

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

func CalculateFuelCosts(positions []int) []int {
	costs := make([]int, len(positions))

	for i, _ := range positions {
		target := i + 1
		totalCost := 0

		for _, position := range positions {
			totalCost += utils.Abs(target - position)
		}

		costs[i] = totalCost
	}

	return costs
}

func LowestAlignment(positions []int) (int, int) {
	costs := CalculateFuelCosts(positions)

	index, cost := utils.ArgMin(costs...)

	return index + 1, cost
}

func ParseInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	return utils.ToInts(strings.Split(scanner.Text(), ","))
}
