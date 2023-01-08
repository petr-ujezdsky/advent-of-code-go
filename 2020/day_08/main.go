package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/strs"
	"io"
)

type Instruction func(global, arg int) (result, jump int)

type Operation struct {
	Name        string
	Arg         int
	Instruction Instruction
}

var instructions = map[string]Instruction{
	"acc": func(global, arg int) (int, int) { return global + arg, 1 },
	"jmp": func(global, arg int) (int, int) { return global, arg },
	"nop": func(global, arg int) (int, int) { return global, 1 },
}

func runOperations(operations []*Operation) (int, bool) {
	i := 0
	global := 0
	visited := make([]bool, len(operations))

	for {
		if i == len(operations) {
			return global, true
		}

		op := operations[i]

		if visited[i] {
			return global, false
		}

		newGlobal, offset := op.Instruction(global, op.Arg)

		global = newGlobal
		visited[i] = true
		i += offset
	}
}

func ValueBeforeCycle(operations []*Operation) int {
	value, ok := runOperations(operations)
	if ok {
		panic("Did not fail")
	}

	return value
}

func FixTheCode(operations []*Operation) int {
	// try without change
	global, ok := runOperations(operations)
	if ok {
		return global
	}

	for i, op := range operations {
		if op.Name == "jmp" {
			// switch to nop
			operationsNew := slices.Clone(operations)
			operationsNew[i] = &Operation{
				Name:        "nop",
				Arg:         op.Arg,
				Instruction: instructions["nop"],
			}

			global, ok = runOperations(operationsNew)
			if ok {
				return global
			}
		}

		if op.Name == "nop" {
			// switch to jmp
			operationsNew := slices.Clone(operations)
			operationsNew[i] = &Operation{
				Name:        "jmp",
				Arg:         op.Arg,
				Instruction: instructions["jmp"],
			}

			global, ok = runOperations(operationsNew)
			if ok {
				return global
			}
		}
	}

	panic("No solution found")
}

func parseInstruction(str string) *Operation {
	name := strs.Substring(str, 0, 3)
	arg := utils.ExtractInts(str, true)[0]
	instruction := instructions[name]

	return &Operation{
		Name:        name,
		Arg:         arg,
		Instruction: instruction,
	}
}

func ParseInput(r io.Reader) []*Operation {
	return parsers.ParseToObjects(r, parseInstruction)
}
