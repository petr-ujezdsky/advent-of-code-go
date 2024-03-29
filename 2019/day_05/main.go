package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

type World struct {
	Program []int
}

func DoWithInputPart01(input int, world World) int {
	return RunProgram(input, world)
}

func RunProgram(input int, world World) int {
	program := slices.Clone(world.Program)
	output := 0
	index := 0

	for {
		opRaw := program[index]

		op := opRaw % 100

		if op == 99 {
			return output
		}

		switch op {
		case 1:
			args := parseArguments(program[index:index+4], program, 2)

			destI := args[2]
			program[destI] = args[0] + args[1]

			if index != destI {
				index += 4
			}
		case 2:
			args := parseArguments(program[index:index+4], program, 2)

			destI := args[2]
			program[destI] = args[0] * args[1]

			if index != destI {
				index += 4
			}
		case 3:
			args := parseArguments(program[index:index+2], program, 0)

			destI := args[0]
			program[destI] = input

			if index != destI {
				index += 2
			}
		case 4:
			args := parseArguments(program[index:index+2], program, -1)

			output = args[0]

			index += 2
		case 5:
			args := parseArguments(program[index:index+3], program, -1)

			if args[0] != 0 {
				index = args[1]
			} else {
				index += 3
			}
		case 6:
			args := parseArguments(program[index:index+3], program, -1)

			if args[0] == 0 {
				index = args[1]
			} else {
				index += 3
			}
		case 7:
			args := parseArguments(program[index:index+4], program, 2)
			destI := args[2]

			if args[0] < args[1] {
				program[destI] = 1
			} else {
				program[destI] = 0
			}

			if index != destI {
				index += 4
			}
		case 8:
			args := parseArguments(program[index:index+4], program, 2)
			destI := args[2]

			if args[0] == args[1] {
				program[destI] = 1
			} else {
				program[destI] = 0
			}

			if index != destI {
				index += 4
			}
		default:
			panic(fmt.Sprintf("Unknown op %v at index %v", opRaw, index))
		}
	}
}

func parseArguments(instruction, program []int, positionModeOnly int) []int {
	fmt.Printf("Parsing instruction %v\n", instruction)

	opMode := instruction[0] / 100

	parsed := make([]int, len(instruction)-1)

	for i, value := range instruction[1:] {
		mode := opMode % 10

		if i == positionModeOnly {
			if mode != 0 {
				panic(fmt.Sprintf("Unexpected immediate mode"))
			}

			parsed[i] = value
		} else {
			parsed[i] = parseArgument(mode, value, program)
		}

		opMode = opMode / 10
	}

	return parsed
}

func parseArgument(mode, value int, program []int) int {
	// position mode
	if mode == 0 {
		return program[value]
	}

	// immediate mode
	if mode == 1 {
		return value
	}

	panic(fmt.Sprintf("Unknown mode %v", mode))
}
func DoWithInputPart02(input int, world World) int {
	return RunProgram(input, world)
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
