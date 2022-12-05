package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type CratesStack = utils.Stack[rune]

type Operation struct {
	From, To, Count int
}

func move(from, to *CratesStack, count int) {
	for i := 0; i < count; i++ {
		to.Push(from.Pop())
	}
}

func MoveCratesByOps(stacks []*CratesStack, ops []Operation) string {
	for _, op := range ops {
		move(stacks[op.From-1], stacks[op.To-1], op.Count)
	}

	topCrates := ""
	for _, stack := range stacks {
		if !stack.Empty() {
			topCrates += string(stack.Peek())
		}
	}

	return topCrates
}

func ParseInput(r io.Reader) ([]*CratesStack, []Operation) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	// stacks status
	var lines [][]rune
	for scanner.Scan() && scanner.Text() != "" {
		lines = append(lines, []rune(scanner.Text()))
	}

	// last char in last line is last stack number -> count
	stacksCount := lines[len(lines)-1][len(lines[len(lines)-1])-1] - '0'
	stacks := make([]*utils.Stack[rune], stacksCount)
	for _, line := range utils.Reverse(lines[:len(lines)-1]) {
		for i, stack := range stacks {
			index := i*4 + 1

			if index >= len(line) {
				continue
			}

			if stack == nil {
				s := utils.NewStack[rune]()
				stack = &s
				stacks[i] = &s
			}
			crate := line[index]
			if crate != ' ' {
				stack.Push(line[index])
			}
		}
	}

	// operations
	var operations []Operation
	for scanner.Scan() {
		line := scanner.Text()
		ints := utils.ExtractInts(line, true)

		op := Operation{
			From:  ints[1],
			To:    ints[2],
			Count: ints[0],
		}

		operations = append(operations, op)
	}

	return stacks, operations
}
