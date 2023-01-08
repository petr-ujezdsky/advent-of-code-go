package day_24

import (
	"bufio"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strconv"
	"strings"
)

type InputStack = collections.Stack[int]
type Registers = [4]int

type Instruction struct {
	Name                  string
	Evaluator             Evaluator
	ILeft, IRight, VRight int
}

type Evaluator = func(left, right int, input *InputStack) int

var evaluators = map[string]Evaluator{
	"inp": inp,
	"add": add,
	"mul": mul,
	"div": div,
	"mod": mod,
	"eql": eql,
}

func inp(_, _ int, input *InputStack) int {
	return input.Pop()
}

func add(left, right int, _ *InputStack) int {
	return left + right
}

func mul(left, right int, _ *InputStack) int {
	return left * right
}

func div(left, right int, _ *InputStack) int {
	if right == 0 {
		panic("Division by zero")
	}
	return left / right
}

func mod(left, right int, _ *InputStack) int {
	if left < 0 {
		panic("Invalid modulo input - left < 0, " + strconv.Itoa(left))
	}
	if right <= 0 {
		panic("Invalid modulo input - right <= 0, " + strconv.Itoa(right))
	}
	return left % right
}

func eql(left, right int, _ *InputStack) int {
	if left == right {
		return 1
	}
	return 0
}

func reg(str string) int {
	return int(str[0] - 'w')
}

func regOrVal(str string) (int, int) {
	switch str[0] {
	case 'w', 'x', 'y', 'z':
		return reg(str), 0
	}

	return -1, utils.ParseInt(str)
}

func NewInputStack(input string) InputStack {
	reversed := slices.Reverse([]rune(input))

	digits := make([]int, len(input))
	for i, d := range reversed {
		digits[i] = int(d - '0')
	}

	return collections.NewStackFilled(digits)
}

func Run(instructions []Instruction, input string) Registers {
	return RunRegisters(Registers{}, instructions, input)
}

func RunRegisters(registers Registers, instructions []Instruction, input string) Registers {
	inputStack := NewInputStack(input)

	debuggingGroupIndex := 0
	iGroup := 0
	for j, i := range instructions {
		left := registers[i.ILeft]
		right := i.VRight
		if i.IRight != -1 {
			right = registers[i.IRight]
		}

		if iGroup == debuggingGroupIndex {
			//fmt.Printf("%v\n%v %2d %2d (%v,%v,%v)\n", registers, i.Name, left, right, i.ILeft, i.IRight, i.VRight)
			//fmt.Printf("%v\n%v %v %2d\n", registers, i.Name, string(rune('w'+i.ILeft)), right)
		}

		if j%18 == 6 {
			//fmt.Printf("input %2d -> %v\n", iGroup, left)
		}

		registers[i.ILeft] = i.Evaluator(left, right, &inputStack)

		if j%18 == 17 {
			//fmt.Printf("Group #%2d: z = %v\n", iGroup, registers[3])
			iGroup++
			if iGroup == 11 {
				//fmt.Println()
			}
			if iGroup == 12 {
				//fmt.Printf("%v, %v\n", i.Name, registers)
			}
		}
	}

	return registers
}

func tryRegisterGroup(registers Registers, input int, iGroup int, instructions []Instruction) (Registers, bool) {
	inputStack := InputStack{}
	inputStack.Push(input)

	for j, i := range instructions {
		left := registers[i.ILeft]
		right := i.VRight
		if i.IRight != -1 {
			right = registers[i.IRight]
		}

		registers[i.ILeft] = i.Evaluator(left, right, &inputStack)

		// second eql instruction in groups with potential to small the z register
		if j == 7 && (iGroup == 3 || iGroup == 5 || iGroup >= 9 && iGroup <= 13) {
			// x needs to be 0 to perform division on z in the end (lowers te z value)
			if registers[1] != 0 {
				return Registers{}, false
			}
		}
	}

	return registers, true
}

func tryRegisterGroupRecursive(registers Registers, iGroup int, from, to int, groups [][]Instruction) (Registers, bool, *collections.Stack[int]) {
	if iGroup == len(groups) {
		return registers, true, &collections.Stack[int]{}
	}

	step := utils.Signum(to - from)

	for input := from; input != to+step; input += step {
		localRegisters, ok := tryRegisterGroup(registers, input, iGroup, groups[iGroup])

		// not feasible input
		if !ok {
			continue
		}

		// verification is OK, continue
		nextRegisters, ok, resultInputs := tryRegisterGroupRecursive(localRegisters, iGroup+1, from, to, groups)
		if ok {
			resultInputs.Push(input)
			return nextRegisters, true, resultInputs
		}
	}

	return Registers{}, false, nil
}

func BruteForcePossibleValues(instructions []Instruction, from, to int) string {
	groups := groupInstructions(instructions)

	registers, ok, resultInputs := tryRegisterGroupRecursive(Registers{}, 0, from, to, groups)
	if !ok {
		panic("Not found any")
	}

	if registers[3] != 0 {
		panic(fmt.Sprintf("Z is not 0, registers: %v", registers))
	}

	digits := slices.Reverse(resultInputs.PeekAll())
	sb := &strings.Builder{}
	for _, digit := range digits {
		sb.WriteString(strconv.Itoa(digit))
	}

	return sb.String()
}

func groupInstructions(instructions []Instruction) [][]Instruction {
	groups := make([][]Instruction, 14)

	for i := 0; len(instructions) > 0; i++ {
		groups[i], instructions = instructions[:18], instructions[18:]
	}

	return groups
}

func extractABC(instructions []Instruction) []utils.Vector3i {
	var abcs []utils.Vector3i
	A, B, C := 0, 0, 0
	for i, instruction := range instructions {
		if i%18 == 4 {
			A = instruction.VRight
		}

		if i%18 == 5 {
			B = instruction.VRight
		}

		if i%18 == 15 {
			C = instruction.VRight

			abcs = append(abcs, utils.Vector3i{A, B, C})
		}
	}

	return abcs
}

func ParseInput(r io.Reader) []Instruction {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var instructions []Instruction
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		var instruction Instruction
		name := parts[0]
		evaluator := evaluators[name]

		if name == "inp" {
			instruction = Instruction{
				Name:      name,
				Evaluator: evaluator,
				ILeft:     reg(parts[1]),
			}
		} else {
			iRight, vRight := regOrVal(parts[2])

			instruction = Instruction{
				Name:      name,
				Evaluator: evaluator,
				ILeft:     reg(parts[1]),
				IRight:    iRight,
				VRight:    vRight,
			}
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}
