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
	common := maps.Intersection([]map[int]struct{}{card.Winning, card.Drawn})
	if matchedCount := len(common); matchedCount > 0 {
		// 2^matchedCount
		return 1 << (matchedCount - 1)
	}

	return 0
}

func DoWithInputPart02(world World) int {
	return 0
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
