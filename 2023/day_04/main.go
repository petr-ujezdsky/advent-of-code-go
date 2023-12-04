package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strings"
)

type Card struct {
	Id             int
	Winning, Drawn map[int]struct{}
}

func (card Card) CommonNumbersCount() int {
	return len(maps.Intersection([]map[int]struct{}{card.Winning, card.Drawn}))
}

type World struct {
	Cards []Card
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, card := range world.Cards {
		sum += scoreCard(card)
	}

	return sum
}

func scoreCard(card Card) int {
	if matchedCount := card.CommonNumbersCount(); matchedCount > 0 {
		// 2^matchedCount
		return 1 << (matchedCount - 1)
	}

	return 0
}

func DoWithInputPart02(world World) int {
	cardCounts := make([]int, len(world.Cards))

	// initialize counts by 1
	for i := range world.Cards {
		cardCounts[i] = 1
	}

	for i, card := range world.Cards {
		if matchedCount := card.CommonNumbersCount(); matchedCount > 0 {
			cardInstancesCount := cardCounts[i]

			for j := 0; j < matchedCount; j++ {
				cardCounts[i+j+1] += cardInstancesCount
			}
		}
	}

	sum := 0
	for _, instancesCount := range cardCounts {
		sum += instancesCount
	}
	return sum
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) Card {
		tokens := strings.Split(str, "|")

		// left - id and winning numbers
		left := utils.ExtractInts(tokens[0], false)
		id := left[0]
		winning := make(map[int]struct{})
		for _, number := range left[1:] {
			winning[number] = struct{}{}
		}

		// right - drawn numbers
		right := utils.ExtractInts(tokens[1], false)
		drawn := make(map[int]struct{})
		for _, number := range right {
			drawn[number] = struct{}{}
		}

		return Card{
			Id:      id,
			Winning: winning,
			Drawn:   drawn,
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Cards: items}
}
