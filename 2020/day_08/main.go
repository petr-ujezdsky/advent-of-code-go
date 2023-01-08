package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
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

func ValueBeforeCycle(operations []*Operation) int {
	i := 0
	global := 0
	visited := make([]bool, len(operations))

	for {
		op := operations[i]

		if visited[i] {
			return global
		}

		newGlobal, offset := op.Instruction(global, op.Arg)

		global = newGlobal
		visited[i] = true
		i += offset
	}
}

func parseInstruction(str string) *Operation {
	name := utils.Substring(str, 0, 3)
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
