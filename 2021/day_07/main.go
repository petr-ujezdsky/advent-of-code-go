package day_07

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type CostFunc func(target, position int) int

func CostSimple(target, position int) int {
	return utils.Abs(target - position)
}

func CostSteppingUp(target, position int) int {
	steps := utils.Abs(target - position)
	return utils.SumNtoM(1, steps)
}

func CalculateFuelCosts(positions []int, costFunc CostFunc) []int {
	costs := make([]int, len(positions))

	for i, _ := range positions {
		target := i + 1
		totalCost := 0

		for _, position := range positions {
			totalCost += costFunc(target, position)
		}

		costs[i] = totalCost
	}

	return costs
}

func LowestAlignment(positions []int, costFunc CostFunc) (int, int) {
	costs := CalculateFuelCosts(positions, costFunc)

	index, cost := utils.ArgMin(costs...)

	return index + 1, cost
}

func ParseInput(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	return utils.ToInts(strings.Split(scanner.Text(), ","))
}
