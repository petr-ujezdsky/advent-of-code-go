package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/2019/common"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/combi"
	"io"
)

type World struct {
	Program []int
}

func runSetting(phases [5]int, program []int) int {
	input := 0

	for _, phase := range phases {
		input = common.RunProgram([]int{phase, input}, program)
	}

	return input
}

func DoWithInputPart01(world World) int {
	quit := make(chan interface{})
	initialPhases := []int{0, 1, 2, 3, 4}
	phases := combi.Permute(quit, initialPhases)

	maximum := -1
	var maximumPhases []int

	for phase := range phases {

		output := runSetting([5]int(phase), world.Program)
		if output > maximum {
			maximum = output
			maximumPhases = phase
		}
	}

	close(quit)

	fmt.Printf("Max: %v, phases: %v\n", maximum, maximumPhases)

	return maximum
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
