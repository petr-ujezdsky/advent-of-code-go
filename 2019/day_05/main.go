package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/2019/common"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type World struct {
	Program []int
}

func DoWithInputPart01(input int, world World) int {
	outputs := common.RunProgram([]int{input}, world.Program)
	return outputs[len(outputs)-1]
}

func DoWithInputPart02(input int, world World) int {
	outputs := common.RunProgram([]int{input}, world.Program)
	return outputs[len(outputs)-1]
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
