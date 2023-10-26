package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
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
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result, ok := RunProgram(program, noun, verb)
			if ok && result == 19690720 {
				return 100*noun + verb
			}
		}
	}

	panic("No result found")
}

func RunProgram(program []int, noun, verb int) (int, bool) {
	// create copy
	program = slices.Clone(program)

	// patch program
	program[1] = noun
	program[2] = verb

	instructionPointer := 0

	for {
		instruction := program[instructionPointer]

		if instruction == 99 {
			return program[0], true
		}

		param1 := program[program[instructionPointer+1]]
		param2 := program[program[instructionPointer+2]]
		dest := &program[program[instructionPointer+3]]

		switch instruction {
		case 1:
			*dest = param1 + param2
		case 2:
			*dest = param1 * param2
		default:
			return 0, false
		}

		instructionPointer += 4
	}
}

func ParseInput(input string) []int {
	return utils.ExtractInts(input, false)
}
