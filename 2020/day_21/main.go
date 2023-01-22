package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/strs"
	"io"
	"strings"
)

type StringSet = map[string]struct{}

type Food struct {
	Ingredients, Allergens StringSet
}

type World struct {
	Foods []Food
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) Food {
		parts := strings.Split(str, " (contains ")

		ingredients := strings.Split(parts[0], " ")
		allergens := strings.Split(strs.Substring(parts[1], 0, len(parts[1])-1), ", ")

		return Food{
			Ingredients: slices.ToSet(ingredients),
			Allergens:   slices.ToSet(allergens),
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Foods: items}
}
