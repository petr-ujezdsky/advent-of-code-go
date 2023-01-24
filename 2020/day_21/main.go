package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/strs"
	"io"
	"strings"
)

type StringSet = map[string]struct{}

type Ingredients = map[string]*Ingredient
type Ingredient struct {
	Name              string
	PossibleAllergens StringSet
}

type Allergens = map[string]*Allergen
type Allergen struct {
	Name                string
	PossibleIngredients StringSet
}

type Food struct {
	Ingredients Ingredients
	Allergens   Allergens
}

type World struct {
	Foods          []Food
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

func findByAllergen(allergen string, foods []Food) []Food {
	var containing []Food

	for _, food := range foods {
		if _, ok := food.Allergens[allergen]; ok {
			containing = append(containing, food)
		}
	}

	return containing
}

func findByIngredient(ingredient string, foods []Food) []Food {
	var containing []Food

	for _, food := range foods {
		if _, ok := food.Ingredients[ingredient]; ok {
			containing = append(containing, food)
		}
	}

	return containing
}

func findHavingOnePossibility(ingredients Ingredients) (*Ingredient, string) {
	for _, ingredient := range ingredients {
		if len(ingredient.PossibleAllergens) == 1 {
			allergen := maps.FirstKey(ingredient.PossibleAllergens)
			return ingredient, allergen
		}
	}

	return nil, ""
}
func DoWithInputPart01(world World) int {
	foods := world.Foods

	// fill all possibilities
	for _, food := range foods {
		for _, allergen := range food.Allergens {
			for _, ingredient := range food.Ingredients {
				ingredient.PossibleAllergens[allergen.Name] = struct{}{}
				allergen.PossibleIngredients[ingredient.Name] = struct{}{}
			}
		}
	}

	for _, allergen := range world.AllAllergens {
		fmt.Printf("%v:\n", allergen.Name)

		//allergen := maps.FirstKey(world.AllAllergens)
		foodsWithAllergen := findByAllergen(allergen.Name, foods)
		fmt.Printf("  * in %v foods\n", len(foodsWithAllergen))

		ingredients := slices.Map(foodsWithAllergen, func(f Food) Ingredients { return f.Ingredients })

		intersectingIngredients := maps.Intersection(ingredients)

		// remove allergen from ingredients that are not in intersection
		count := 0
		for _, ingredient := range world.AllIngredients {
			if _, ok := intersectingIngredients[ingredient.Name]; !ok {
				delete(ingredient.PossibleAllergens, allergen.Name)
				delete(allergen.PossibleIngredients, ingredient.Name)
				count++
				//fmt.Printf("  * removed from ingredient %v\n", name)
			}
		}
		fmt.Printf("  * removed from %v ingredients\n", count)
	}

	fmt.Println()
	fmt.Println("Resolved:")

	resolvedAllergens := make(map[string]string)

	for {
		ingredient, allergenName := findHavingOnePossibility(world.AllIngredients)
		if ingredient == nil {
			break
		}

		fmt.Printf("  * %10v -> %v\n", allergenName, ingredient.Name)

		resolvedAllergens[allergenName] = ingredient.Name

		for _, otherIngredient := range world.AllIngredients {
			delete(otherIngredient.PossibleAllergens, allergenName)
		}

		for _, otherAllergen := range world.AllAllergens {
			delete(otherAllergen.PossibleIngredients, ingredient.Name)
		}
		allergen := world.AllAllergens[allergenName]
		allergen.PossibleIngredients = make(StringSet)
	}

	fmt.Println()
	fmt.Println("Ingredients without allergens:")

	for _, ingredient := range world.AllIngredients {
		if len(ingredient.PossibleAllergens) == 0 {
			fmt.Printf("  * %v\n", ingredient.Name)
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

	return 0
}

func DoWithInputPart02(world World) int {
	return 0
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

	parseItem := func(str string) Food {
		parts := strings.Split(str, " (contains ")

		ingredientNames := strings.Split(parts[0], " ")
		allergenNames := strings.Split(strs.Substring(parts[1], 0, len(parts[1])-1), ", ")

		ingredients := make(Ingredients)
		for _, ingredientName := range ingredientNames {
			ingredient := getOrCreateIngredient(ingredientName, allIngredients)
			ingredients[ingredientName] = ingredient
		}

		allergens := make(Allergens)
		for _, allergenName := range allergenNames {
			allergen := getOrCreateAllergen(allergenName, allAllergens)
			allergens[allergenName] = allergen
		}

		return Food{
			Ingredients: ingredients,
			Allergens:   allergens,
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{
		Foods:          items,
		AllIngredients: allIngredients,
		AllAllergens:   allAllergens,
	}
}
