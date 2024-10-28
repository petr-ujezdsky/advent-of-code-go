package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/2019/common"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type World struct {
	Program []int
}

func DoWithInputPart01(world World) int {
	outputs := common.RunProgram([]int{1}, world.Program)
	if len(outputs) > 1 {
		panic(fmt.Sprintf("Invalid implementation, errors: %v", outputs))
	}

	return outputs[0]
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
