package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type World struct {
	Program []int
}

func DoWithInputPart01(input int, world World) int {
	program := world.Program
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

			dest := &program[args[2]]
			*dest = args[0] + args[1]

			index += 4
		case 2:
			args := parseArguments(program[index:index+4], program, 2)

			dest := &program[args[2]]
			*dest = args[0] * args[1]

			index += 4
		case 3:
			args := parseArguments(program[index:index+2], program, 0)

			dest := &program[args[0]]
			*dest = input

			index += 2
		case 4:
			args := parseArguments(program[index:index+2], program, -1)

			output = args[0]

			index += 2
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
func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	return World{Program: utils.ExtractInts(scanner.Text(), true)}
}
