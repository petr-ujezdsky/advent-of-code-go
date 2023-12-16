package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strings"
)

type Step struct {
	Raw         string
	Label       string
	Operation   rune
	FocalLength int
}

type Lens struct {
	Label       string
	FocalLength int
}

type World struct {
	Steps []Step
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, step := range world.Steps {
		sum += Hash(step.Raw)
	}

	return sum
}

func Hash(s string) int {
	hash := 0

	for _, char := range s {
		hash = ((hash + int(char)) * 17) % 256
	}

	return hash
}

func DoWithInputPart02(world World) int {
	table := make([]*linkedhashmap.Map, 256)

	for i := range table {
		table[i] = linkedhashmap.New()
	}

	for _, step := range world.Steps {
		box := table[Hash(step.Label)]

		switch step.Operation {
		case '-':
			box.Remove(step.Label)
		case '=':
			box.Put(step.Label, Lens{
				Label:       step.Label,
				FocalLength: step.FocalLength,
			})
		}

	}

	PrintTable(table)

	return CalculateFocusingPower(table)
}

func CalculateFocusingPower(table []*linkedhashmap.Map) int {
	sum := 0

	for iBox, box := range table {
		for iLens, lens := range box.Values() {
			sum += (iBox + 1) * (iLens + 1) * lens.(Lens).FocalLength
		}
	}

	return sum
}

func PrintTable(table []*linkedhashmap.Map) {
	for i, box := range table {
		if box.Empty() {
			continue
		}
		fmt.Printf("Box %d: %v\n", i, box.Values())
	}
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	tokens := strings.Split(scanner.Text(), ",")

	steps := slices.Map(tokens, func(stepStr string) Step {
		parts := strings.Split(stepStr, "=")

		if len(parts) == 2 {
			return Step{
				Raw:         stepStr,
				Label:       parts[0],
				Operation:   '=',
				FocalLength: utils.ParseInt(parts[1]),
			}
		}

		return Step{
			Raw:         stepStr,
			Label:       stepStr[:len(stepStr)-1],
			Operation:   '-',
			FocalLength: 0,
		}
	})

	return World{Steps: steps}
}
