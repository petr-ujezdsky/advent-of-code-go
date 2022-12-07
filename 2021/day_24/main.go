package day_24

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type InputStack = utils.Stack[int]
type Registers = [4]int

type Instruction struct {
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
	if left < 0 || right <= 0 {
		panic("Invalid modulo input")
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
	inputStack := NewInputStack(input)
	registers := Registers{}

	for _, i := range instructions {
		left := registers[i.ILeft]
		right := i.VRight
		if i.IRight != -1 {
			right = registers[i.IRight]
		}

		registers[i.ILeft] = i.Evaluator(left, right, &inputStack)
	}

	return registers
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
				Evaluator: evaluator,
				ILeft:     reg(parts[1]),
			}
		} else {
			iRight, vRight := regOrVal(parts[2])

			instruction = Instruction{
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
