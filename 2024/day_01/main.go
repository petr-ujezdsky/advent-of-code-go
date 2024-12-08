package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"sort"
)

type Item struct {
	Left, Right int
}

type World struct {
	Items []Item
}

func sortByLeft(items []Item) []Item {
	sort.Slice(items, func(i, j int) bool { return items[i].Left < items[j].Left })
	return items
}

func sortByRight(items []Item) []Item {
	sort.Slice(items, func(i, j int) bool { return items[i].Right < items[j].Right })
	return items
}

func DoWithInputPart01(world World) int {
	left := sortByLeft(slices.Clone(world.Items))
	right := sortByRight(slices.Clone(world.Items))

	totalDistance := 0
	for i, leftItem := range left {
		rightItem := right[i]

		totalDistance += utils.Abs(rightItem.Right - leftItem.Left)
	}

	return totalDistance
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) Item {
		numbers := utils.ExtractInts(str, false)

		return Item{
			Left:  numbers[0],
			Right: numbers[1],
		}
	}

	items := parsers.ParseToObjects(r, parseItem)

	return World{Items: items}
}
