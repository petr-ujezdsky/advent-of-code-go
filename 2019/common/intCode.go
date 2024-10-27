package common

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
)

const debug = false

type IntCodeComputer struct {
	Name                  string
	program               []int
	inputs, outputs, halt chan int
}

type argumentMode int

const (
	modePosition  argumentMode = 0
	modeImmediate              = 1
)

func NewIntCodeComputer(name string, program []int, inputs, outputs, halt chan int) IntCodeComputer {
	return IntCodeComputer{
		Name:    name,
		program: program,
		inputs:  inputs,
		outputs: outputs,
		halt:    halt,
	}
}

func InputSlice(inputs []int, input chan int) func() {
	return func() {
		for _, inputValue := range inputs {
			input <- inputValue
		}
	}
}

func RunProgram(inputs []int, program []int) int {
	program = slices.Clone(program)
	input := make(chan int)
	halt := make(chan int)

	defer close(input)
	defer close(halt)

	computer := NewIntCodeComputer("Unknown", program, input, nil, halt)

	// input
	go InputSlice(inputs, input)()

	go Run(computer)

	return <-halt
}

func Run(computer IntCodeComputer) {
	program := computer.program
	inputs := computer.inputs
	outputs := computer.outputs
	halt := computer.halt

	lastOutput := 0
	index := 0

	for {
		opRaw := program[index]

		op := opRaw % 100

		switch op {
		case 1:
			// add two numbers
			args := parseArguments(program[index:index+4], program, 2)

			destI := args[2]
			program[destI] = args[0] + args[1]

			if index != destI {
				index += 4
			}
		case 2:
			// multiply two numbers
			args := parseArguments(program[index:index+4], program, 2)

			destI := args[2]
			program[destI] = args[0] * args[1]

			if index != destI {
				index += 4
			}
		case 3:
			// read input
			args := parseArguments(program[index:index+2], program, 0)

			destI := args[0]
			inputValue := <-inputs
			program[destI] = inputValue

			if index != destI {
				index += 2
			}
		case 4:
			// write output
			args := parseArguments(program[index:index+2], program, -1)

			lastOutput = args[0]
			if outputs != nil {
				outputs <- lastOutput
			}

			index += 2
		case 5:
			// jump-if-true
			args := parseArguments(program[index:index+3], program, -1)

			if args[0] != 0 {
				index = args[1]
			} else {
				index += 3
			}
		case 6:
			// jump-if-false
			args := parseArguments(program[index:index+3], program, -1)

			if args[0] == 0 {
				index = args[1]
			} else {
				index += 3
			}
		case 7:
			// less than
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
			// equals
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
		case 99:
			// halt
			if halt != nil {
				halt <- lastOutput
			}

			return
		default:
			panic(fmt.Sprintf("Unknown op %v at index %v", opRaw, index))
		}
	}
}

func parseArguments(instruction, program []int, positionModeOnly int) []int {
	if debug {
		fmt.Printf("Parsing instruction %v\n", instruction)
	}

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
			parsed[i] = parseArgument(argumentMode(mode), value, program)
		}

		opMode = opMode / 10
	}

	return parsed
}

func parseArgument(mode argumentMode, value int, program []int) int {
	// position mode
	if mode == modePosition {
		return program[value]
	}

	// immediate mode
	if mode == modeImmediate {
		return value
	}

	panic(fmt.Sprintf("Unknown mode %v", mode))
}
