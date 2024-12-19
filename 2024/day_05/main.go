package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	slices2 "github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"slices"
)

type Rule struct {
	Left, Right int
}

type Update []int

type World struct {
	Rules   []Rule
	Updates []Update
}

func createComparator(rules []Rule) func(a, b int) int {
	rulesSet := slices2.ToSet(rules)

	return func(a, b int) int {
		if _, ok := rulesSet[Rule{Left: a, Right: b}]; ok {
			return -1
		}

		if _, ok := rulesSet[Rule{Left: b, Right: a}]; ok {
			return 1
		}

		panic("Not found")
	}
}

func DoWithInputPart01(world World) int {
	cmp := createComparator(world.Rules)

	middlesSum := 0
	for _, update := range world.Updates {
		sorted := slices.IsSortedFunc(update, cmp)
		if !sorted {
			fmt.Printf("❌ %v\n", update)
			continue
		}

		fmt.Printf("✅ %v\n", update)

		middle := update[len(update)/2]
		middlesSum += middle
	}

	return middlesSum
}

func DoWithInputPart02(world World) int {
	cmp := createComparator(world.Rules)

	middlesSum := 0
	for _, update := range world.Updates {
		sorted := slices.IsSortedFunc(update, cmp)
		if !sorted {
			fmt.Printf("❌ %v\n", update)

			sortedUpdate := slices.SortedFunc(slices.Values(update), cmp)

			middle := sortedUpdate[len(sortedUpdate)/2]
			middlesSum += middle
			continue
		}

		fmt.Printf("✅ %v\n", update)
	}

	return middlesSum
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// Rules
	var rules []Rule
	for scanner.Scan() && len(scanner.Text()) > 0 {
		ints := utils.ExtractInts(scanner.Text(), false)
		rules = append(rules, Rule{
			Left:  ints[0],
			Right: ints[1],
		})
	}

	// Updates
	var updates []Update
	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)
		updates = append(updates, ints)
	}

	return World{
		Rules:   rules,
		Updates: updates,
	}
}
