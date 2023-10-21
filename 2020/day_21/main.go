package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/strs"
	"io"
	"sort"
	"strings"
)

type StringSet = map[string]struct{}

type Ingredients = map[string]*Ingredient
type Ingredient struct {
	Name              string
	PossibleAllergens StringSet
	Foods             []*Food
}

type Allergens = map[string]*Allergen
type Allergen struct {
	Name                string
	PossibleIngredients StringSet
	Foods               []*Food
}

type Food struct {
	Ingredients Ingredients
	Allergens   Allergens
}

type World struct {
	Foods          []*Food
	AllIngredients Ingredients
	AllAllergens   Allergens
}

//func findIntersection(s1, s2 StringSet) StringSet {
//	intersection := make(StringSet)
//	for v1 := range s1 {
//		if _, ok := s2[v1]; ok {
//			intersection[v1] = struct{}{}
//		}
//	}
//
//	return intersection
//}

//
//func removeIngredientAndAllergen(ingredient, allergen string, foods []Food) {
//	for _, food := range foods {
//		delete(food.Ingredients, ingredient)
//		delete(food.Allergens, allergen)
//	}
//}
//
//func findTranslation(foods []Food) (string, string, bool) {
//	for i, food1 := range foods {
//		// check direct translation
//		if len(food1.Ingredients) == 1 && len(food1.Allergens) == 1 {
//			ingredient := maps.FirstKey(food1.Ingredients)
//			allergen := maps.FirstKey(food1.Allergens)
//
//			return ingredient, allergen, true
//		}
//
//		// check intersections with other foods
//		for j := i + 1; j < len(foods); j++ {
//			food2 := foods[j]
//
//			sameAllergens := findIntersection(food1.Allergens, food2.Allergens)
//			if len(sameAllergens) == 1 {
//				sameIngredients := findIntersection(food1.Ingredients, food2.Ingredients)
//				if len(sameIngredients) == 1 {
//					ingredient := maps.FirstKey(sameIngredients)
//					allergen := maps.FirstKey(sameAllergens)
//
//					return ingredient, allergen, true
//				}
//			}
//		}
//	}
//
//	return "", "", false
//}

//
//func naiveSolution(world World) int {
//	foundTranslations := make(map[string]string)
//	foods := world.Foods
//
//	for {
//		ingredient, allergen, ok := findTranslation(foods)
//		if !ok {
//			break
//		}
//
//		removeIngredientAndAllergen(ingredient, allergen, foods)
//		foundTranslations[ingredient] = allergen
//	}
//
//	count := 0
//	for _, food := range foods {
//		count += len(food.Ingredients)
//	}
//
//	return count
//}

//func findByAllergen(allergen string, foods []*Food) []*Food {
//	var containing []*Food
//
//	for _, food := range foods {
//		if _, ok := food.Allergens[allergen]; ok {
//			containing = append(containing, food)
//		}
//	}
//
//	return containing
//}

//func findByIngredient(ingredient string, foods []*Food) []*Food {
//	var containing []*Food
//
//	for _, food := range foods {
//		if _, ok := food.Ingredients[ingredient]; ok {
//			containing = append(containing, food)
//		}
//	}
//
//	return containing
//}

//func findHavingOnePossibility(ingredients Ingredients, resolvedAllergens map[string]string) (*Ingredient, string) {
//	for _, ingredient := range ingredients {
//		if len(ingredient.PossibleAllergens) == 1 {
//			allergen := maps.FirstKey(ingredient.PossibleAllergens)
//
//			if _, ok := resolvedAllergens[allergen]; ok {
//				continue
//			}
//
//			return ingredient, allergen
//		}
//	}
//
//	return nil, ""
//}

func findHavingOnePossibilityWithinFood(foods []*Food, resolvedAllergens map[string]string) (*Ingredient, *Allergen, int) {
	for i, food := range foods {
		for _, allergen := range food.Allergens {
			// skip already resolved allergens
			if _, ok := resolvedAllergens[allergen.Name]; ok {
				continue
			}

			var foundIngredient *Ingredient
			for _, ingredient := range food.Ingredients {
				// the allergen is not within possible allergens of the ingredient
				if _, ok := ingredient.PossibleAllergens[allergen.Name]; !ok {
					continue
				}

				if foundIngredient != nil {
					// found another ingredient with the same allergen -> skip and continue to next allergen
					foundIngredient = nil
					break
				}

				foundIngredient = ingredient
			}

			if foundIngredient != nil {
				return foundIngredient, allergen, i
			}
		}
	}

	return nil, nil, 0
}

func DoWithInput(world World) (int, string) {
	foods := world.Foods

	// fill all possibilities (Allergens aren't always marked so everything is possible)
	for _, allergen := range world.AllAllergens {
		for _, ingredient := range world.AllIngredients {
			ingredient.PossibleAllergens[allergen.Name] = struct{}{}
			allergen.PossibleIngredients[ingredient.Name] = struct{}{}
		}
	}

	fmt.Printf("Stats:\n")
	fmt.Printf("  * %v foods\n", len(foods))
	fmt.Printf("  * %v ingredients\n", len(world.AllIngredients))
	fmt.Printf("  * %v allergens\n", len(world.AllAllergens))
	fmt.Println()

	for _, allergen := range world.AllAllergens {
		fmt.Printf("%v:\n", allergen.Name)

		//allergen := maps.FirstKey(world.AllAllergens)
		foodsWithAllergen := allergen.Foods
		fmt.Printf("  * in %v foods\n", len(foodsWithAllergen))

		ingredients := slices.Map(foodsWithAllergen, func(f *Food) Ingredients { return f.Ingredients })

		intersectingIngredients := maps.Intersection(ingredients)

		possibleIngredientsCountBefore := len(allergen.PossibleIngredients)

		// remove allergen from ingredients that are not in intersection
		for _, ingredient := range world.AllIngredients {
			if _, ok := intersectingIngredients[ingredient.Name]; !ok {
				delete(ingredient.PossibleAllergens, allergen.Name)
				delete(allergen.PossibleIngredients, ingredient.Name)
				//fmt.Printf("  * removed from ingredient %v\n", name)
			}
		}
		fmt.Printf("  * ingredients reduced %v -> %v\n", possibleIngredientsCountBefore, len(allergen.PossibleIngredients))
	}

	fmt.Println()
	fmt.Println("Resolved:")

	resolvedAllergens := make(map[string]string)

	for {
		ingredient, allergen, foodIndex := findHavingOnePossibilityWithinFood(world.Foods, resolvedAllergens)
		if ingredient == nil {
			break
		}

		fmt.Printf("  * %10v -> %v by food #%v\n", allergen.Name, ingredient.Name, foodIndex+1)

		resolvedAllergens[allergen.Name] = ingredient.Name

		for _, otherIngredient := range world.AllIngredients {
			delete(otherIngredient.PossibleAllergens, allergen.Name)
		}
		ingredient.PossibleAllergens = make(StringSet)
		ingredient.PossibleAllergens[allergen.Name] = struct{}{}

		for _, otherAllergen := range world.AllAllergens {
			delete(otherAllergen.PossibleIngredients, ingredient.Name)
		}
		allergen.PossibleIngredients = make(StringSet)
		allergen.PossibleIngredients[ingredient.Name] = struct{}{}
	}

	fmt.Println()
	fmt.Println("Ingredients without allergens:")

	count := 0
	for _, ingredient := range world.AllIngredients {
		if len(ingredient.PossibleAllergens) == 0 {
			fmt.Printf("  * %v\n", ingredient.Name)

			count += len(ingredient.Foods)
		}
	}

	fmt.Println()
	fmt.Println("Ingredients possible allergens:")

	for _, ingredient := range world.AllIngredients {
		fmt.Printf("%10v: (%v) %v\n", ingredient.Name, len(ingredient.PossibleAllergens), maps.Keys(ingredient.PossibleAllergens))
	}

	fmt.Println()
	fmt.Println("Allergens possible ingredients:")

	for _, allergen := range world.AllAllergens {
		fmt.Printf("%10v: (%v) %v\n\n", allergen.Name, len(allergen.PossibleIngredients), maps.Keys(allergen.PossibleIngredients))
	}

	// part 2
	sortedAllergenNames := maps.Keys(world.AllAllergens)
	sort.Strings(sortedAllergenNames)

	var ingredientNames []string
	for _, allergenName := range sortedAllergenNames {
		allergen := world.AllAllergens[allergenName]

		if len(allergen.PossibleIngredients) > 1 {
			panic("Allergen has more than one possible ingredient")
		}

		ingredientNames = append(ingredientNames, maps.FirstKey(allergen.PossibleIngredients))
	}

	joined := strings.Join(ingredientNames, ",")

	return count, joined
}

func getOrCreateIngredient(name string, ingredients Ingredients) *Ingredient {
	if ingredient, ok := ingredients[name]; ok {
		return ingredient
	}

	ingredient := &Ingredient{
		Name:              name,
		PossibleAllergens: make(StringSet),
	}

	ingredients[name] = ingredient

	return ingredient
}

func getOrCreateAllergen(name string, allergens Allergens) *Allergen {
	if allergen, ok := allergens[name]; ok {
		return allergen
	}

	allergen := &Allergen{
		Name:                name,
		PossibleIngredients: make(StringSet),
	}

	allergens[name] = allergen

	return allergen
}

func ParseInput(r io.Reader) World {
	allIngredients := make(Ingredients)
	allAllergens := make(Allergens)

	parseItem := func(str string) *Food {
		parts := strings.Split(str, " (contains ")

		ingredientNames := strings.Split(parts[0], " ")
		allergenNames := strings.Split(strs.Substring(parts[1], 0, len(parts[1])-1), ", ")

		food := &Food{}

		ingredients := make(Ingredients)
		food.Ingredients = ingredients
		for _, ingredientName := range ingredientNames {
			ingredient := getOrCreateIngredient(ingredientName, allIngredients)
			ingredient.Foods = append(ingredient.Foods, food)
			ingredients[ingredientName] = ingredient
		}

		allergens := make(Allergens)
		food.Allergens = allergens
		for _, allergenName := range allergenNames {
			allergen := getOrCreateAllergen(allergenName, allAllergens)
			allergen.Foods = append(allergen.Foods, food)
			allergens[allergenName] = allergen
		}

		return food
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{
		Foods:          items,
		AllIngredients: allIngredients,
		AllAllergens:   allAllergens,
	}
}
