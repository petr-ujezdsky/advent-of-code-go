package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
)

func DoWithInputPart01(program []int) int {
	index := 0

	for {
		op := program[index]

		if op == 99 {
			return program[0]
		}

		arg1 := program[program[index+1]]
		arg2 := program[program[index+2]]
		dest := &program[program[index+3]]

		switch op {
		case 1:
			*dest = arg1 + arg2
		case 2:
			*dest = arg1 * arg2
		default:
			panic(fmt.Sprintf("Unknown op %v at index %v", op, index))
		}

		index += 4
	}
}

func PatchProgram(program []int) {
	program[1] = 12
	program[2] = 2
}

func DoWithInputPart02(program []int) int {
	return 0
}

func ParseInput(input string) []int {
	return utils.ExtractInts(input, false)
}
