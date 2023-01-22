package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
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

func findIntersection(s1, s2 StringSet) StringSet {
	intersection := make(StringSet)
	for v1 := range s1 {
		if _, ok := s2[v1]; ok {
			intersection[v1] = struct{}{}
		}
	}

	return intersection
}

func removeIngredientAndAllergen(ingredient, allergen string, foods []Food) {
	for _, food := range foods {
		delete(food.Ingredients, ingredient)
		delete(food.Allergens, allergen)
	}
}

func findTranslation(foods []Food) (string, string, bool) {
	for i, food1 := range foods {
		// check direct translation
		if len(food1.Ingredients) == 1 && len(food1.Allergens) == 1 {
			ingredient := maps.FirstKey(food1.Ingredients)
			allergen := maps.FirstKey(food1.Allergens)

			return ingredient, allergen, true
		}

		// check intersections with other foods
		for j := i + 1; j < len(foods); j++ {
			food2 := foods[j]

			sameAllergens := findIntersection(food1.Allergens, food2.Allergens)
			if len(sameAllergens) == 1 {
				sameIngredients := findIntersection(food1.Ingredients, food2.Ingredients)
				if len(sameIngredients) == 1 {
					ingredient := maps.FirstKey(sameIngredients)
					allergen := maps.FirstKey(sameAllergens)

					return ingredient, allergen, true
				}
			}
		}
	}

	return "", "", false
}

func DoWithInputPart01(world World) int {
	foundTranslations := make(map[string]string)
	foods := world.Foods

	for {
		ingredient, allergen, ok := findTranslation(foods)
		if !ok {
			break
		}

		removeIngredientAndAllergen(ingredient, allergen, foods)
		foundTranslations[ingredient] = allergen
	}

	count := 0
	for _, food := range foods {
		count += len(food.Ingredients)
	}

	return count
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
