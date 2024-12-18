package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Rule struct {
	Left, Right int
}

type Update []int

type World struct {
	Rules   []Rule
	Updates []Update
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
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
