package common

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
)

const debug = false

type IntCodeComputer struct {
	Name                  string
	memory                *unifiedMemory
	inputs, outputs, halt chan int
}

type unifiedMemory struct {
	program []int
	heap    map[int]int
}

func (m unifiedMemory) read(address int) int {
	if address < 0 {
		panic(fmt.Sprintf("Address must be positive, %v", address))
	}

	if address < len(m.program) {
		return m.program[address]
	}

	value, ok := m.heap[address]
	if !ok {
		//panic(fmt.Sprintf("Address not initialized, %v", address))
		return 0
	}

	return value
}

func (m unifiedMemory) readMany(address, count int) []int {
	if address < 0 {
		panic(fmt.Sprintf("Address must be positive, %v", address))
	}

	if address+count <= len(m.program) {
		return m.program[address : address+count]
	}

	values := make([]int, count)
	for i := address; i < address+count; i++ {
		values[i-address] = m.read(address)
	}

	return values
}

func (m unifiedMemory) write(address, value int) {
	if address < 0 {
		panic(fmt.Sprintf("Address must be positive, %v", address))
	}

	if address < len(m.program) {
		m.program[address] = value
	}

	m.heap[address] = value
}

type argumentMode int

const (
	modePosition  argumentMode = 0
	modeImmediate              = 1
	modeRelative               = 2
)

func NewIntCodeComputer(name string, program []int, inputs, outputs, halt chan int) IntCodeComputer {
	return IntCodeComputer{
		Name: name,
		memory: &unifiedMemory{
			program: program,
			heap:    make(map[int]int),
		},
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

func RunProgram(inputs []int, program []int) []int {
	program = slices.Clone(program)
	input := make(chan int)
	output := make(chan int)
	halt := make(chan int)

	defer close(input)
	defer close(halt)

	computer := NewIntCodeComputer("Unknown", program, input, output, halt)

	// input
	go InputSlice(inputs, input)()

	go Run(computer)

	// close output on halt
	go func() {
		<-halt
		close(output)
	}()

	// output
	var outputs []int
	for value := range output {
		outputs = append(outputs, value)
	}

	return outputs
}

func Run(computer IntCodeComputer) {
	memory := computer.memory
	inputs := computer.inputs
	outputs := computer.outputs
	halt := computer.halt

	lastOutput := 0
	relativeBase := 0
	index := 0

	for {
		opRaw := memory.read(index)

		op := opRaw % 100

		switch op {
		case 1:
			// add two numbers
			args := parseArguments(index, 4, memory, relativeBase, 2)

			destI := args[2]
			memory.write(destI, args[0]+args[1])

			if index != destI {
				index += 4
			}
		case 2:
			// multiply two numbers
			args := parseArguments(index, 4, memory, relativeBase, 2)

			destI := args[2]
			memory.write(destI, args[0]*args[1])

			if index != destI {
				index += 4
			}
		case 3:
			// read input
			args := parseArguments(index, 2, memory, relativeBase, 0)

			destI := args[0]
			inputValue := <-inputs
			memory.write(destI, inputValue)

			if index != destI {
				index += 2
			}
		case 4:
			// write output
			args := parseArguments(index, 2, memory, relativeBase, -1)

			lastOutput = args[0]
			if outputs != nil {
				outputs <- lastOutput
			}

			index += 2
		case 5:
			// jump-if-true
			args := parseArguments(index, 3, memory, relativeBase, -1)

			if args[0] != 0 {
				index = args[1]
			} else {
				index += 3
			}
		case 6:
			// jump-if-false
			args := parseArguments(index, 3, memory, relativeBase, -1)

			if args[0] == 0 {
				index = args[1]
			} else {
				index += 3
			}
		case 7:
			// less than
			args := parseArguments(index, 4, memory, relativeBase, 2)
			destI := args[2]

			if args[0] < args[1] {
				memory.write(destI, 1)
			} else {
				memory.write(destI, 0)
			}

			if index != destI {
				index += 4
			}
		case 8:
			// equals
			args := parseArguments(index, 4, memory, relativeBase, 2)
			destI := args[2]

			if args[0] == args[1] {
				memory.write(destI, 1)
			} else {
				memory.write(destI, 0)
			}

			if index != destI {
				index += 4
			}
		case 9:
			// set relative base
			args := parseArguments(index, 2, memory, relativeBase, -1)

			relativeBase += args[0]

			index += 2
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

func parseArguments(addressFrom, count int, memory *unifiedMemory, relativeBase, positionModeOnly int) []int {
	instruction := memory.readMany(addressFrom, count)

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
			parsed[i] = parseArgument(argumentMode(mode), value, memory, relativeBase)
		}

		opMode = opMode / 10
	}

	return parsed
}

func parseArgument(mode argumentMode, value int, memory *unifiedMemory, relativeBase int) int {
	// position mode
	if mode == modePosition {
		return memory.read(value)
	}

	// immediate mode
	if mode == modeImmediate {
		return value
	}

	// relative mode
	if mode == modeRelative {
		return memory.read(relativeBase + value)
	}

	panic(fmt.Sprintf("Unknown mode %v", mode))
}
