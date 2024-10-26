package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/2019/common"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/combi"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

type World struct {
	Program []int
}

func runSetting(phases [5]int, program []int) int {
	chEA := make(chan int, 1)
	chAB := make(chan int, 1)
	chBC := make(chan int, 1)
	chCD := make(chan int, 1)
	chDE := make(chan int, 1)
	end := make(chan int)

	// init computers
	computers := [5]common.IntCodeComputer{
		common.NewIntCodeComputer("A", slices.Clone(program), chEA, chAB, nil),
		common.NewIntCodeComputer("B", slices.Clone(program), chAB, chBC, nil),
		common.NewIntCodeComputer("C", slices.Clone(program), chBC, chCD, nil),
		common.NewIntCodeComputer("D", slices.Clone(program), chCD, chDE, nil),
		common.NewIntCodeComputer("E", slices.Clone(program), chDE, chEA, end),
	}

	// start computers
	for _, computer := range computers {
		go common.Run(computer)
	}

	// first input is phase
	go func() {
		chEA <- phases[0]
		chAB <- phases[1]
		chBC <- phases[2]
		chCD <- phases[3]
		chDE <- phases[4]

		// start off with 0
		chEA <- 0
	}()

	// wait for E to halt
	return <-end
}

func findMaxWithPhaseValues(initialPhases []int, world World) int {
	quit := make(chan interface{})
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

func DoWithInputPart01(world World) int {
	return findMaxWithPhaseValues([]int{0, 1, 2, 3, 4}, world)
}

func DoWithInputPart02(world World) int {
	return findMaxWithPhaseValues([]int{5, 6, 7, 8, 9}, world)
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
