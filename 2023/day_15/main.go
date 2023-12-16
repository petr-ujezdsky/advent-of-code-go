package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strings"
)

type Step struct {
	Raw string
}

type World struct {
	Steps []Step
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, step := range world.Steps {
		sum += Hash(step)
	}

	return sum
}

func Hash(step Step) int {
	hash := 0

	for _, char := range step.Raw {
		hash = ((hash + int(char)) * 17) % 256
	}

	return hash
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	tokens := strings.Split(scanner.Text(), ",")

	steps := slices.Map(tokens, func(token string) Step { return Step{Raw: token} })

	return World{Steps: steps}
}
