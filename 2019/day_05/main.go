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
	return common.RunProgram([]int{input}, world.Program)
}

func DoWithInputPart02(input int, world World) int {
	return common.RunProgram([]int{input}, world.Program)
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
