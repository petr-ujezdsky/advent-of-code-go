package day_24

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strconv"
	"strings"
)

type InputStack = utils.Stack[int]
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
	reversed := utils.Reverse([]rune(input))

	digits := make([]int, len(input))
	for i, d := range reversed {
		digits[i] = int(d - '0')
	}

	return utils.NewStackFilled(digits)
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

func groupInstructions(instructions []Instruction) [][]Instruction {
	groups := make([][]Instruction, 14)

	for i := 0; len(instructions) > 0; i++ {
		groups[i], instructions = instructions[:18], instructions[18:]
	}

	return groups
}

func deGroupInstructions(groups [][]Instruction, indexes ...int) []Instruction {
	var instructions []Instruction

	for _, iGroup := range indexes {
		instructions = append(instructions, groups[iGroup]...)
	}

	return instructions
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

func RunDecompiled(instructions []Instruction, input string) int {
	return RunDecompiledRegister(0, instructions, input)
}

func RunDecompiledRegister(z int, instructions []Instruction, input string) int {
	//fmt.Printf("Input %v\n", input)

	abcs := extractABC(instructions)

	//prevABC := utils.Vector3i{}
	//prevIn := 10000
	for i, abc := range abcs {
		in := int(input[i] - '0')
		A := abc.X
		B := abc.Y
		C := abc.Z

		//fmt.Printf("Group #%2d: z = (%v", i, z)

		if z%26+B == in {
			z = z / A
			//fmt.Printf(" / %v) = %v\n", A, z)
		} else {
			z = (z/A)*26 + in + C
			//fmt.Printf(" / %v) * 26 + %v + %v = %v\n", A, in, C, z)
			//fmt.Printf("input %2d -> %d\n", i, z%26+B)
		}
		//if prevABC.Z+B == in-prevIn {
		//	z = z / A
		//	prevDivided = true
		//	fmt.Printf(" / %v) = %v\n", A, z)
		//} else {
		//	z = (z/A)*26 + in + C
		//	prevDivided = false
		//	fmt.Printf(" / %v) * 26 + %v + %v = %v\n", A, in, C, z)
		//}

		//
		//z /= A
		//z *= 26
		//z += in + C

		//fmt.Printf("input %2d - %2d = %2d (prev=%v, current=%v)\n", i, i-1, prevABC.Z+B, prevABC, abc)
		//fmt.Printf("input %2d - %2d = %2d\n", i, i-1, prevABC.Z+B)

		//prevABC = abc
		//prevIn = in
	}

	return z
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
