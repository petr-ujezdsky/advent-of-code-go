package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type Instruction struct {
	Name   string
	Amount int
}

func readSignalStrength(cycle, x int) int {
	if cycle%40 == 20 {
		return cycle * x
	}

	return 0
}

func DoWithInput(instructions []Instruction) int {
	strengthSum := 0
	cycle := 0
	x := 1
	for i, instruction := range instructions {
		if instruction.Name == "noop" {
			cycle++
			strengthSum += readSignalStrength(cycle, x)
			continue
		}

		cycle++
		strengthSum += readSignalStrength(cycle, x)
		cycle++
		strengthSum += readSignalStrength(cycle, x)
		x += instruction.Amount
		_ = i
	}
	return strengthSum
}

func ParseInput(r io.Reader) []Instruction {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var instructions []Instruction
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		var instruction Instruction

		name := parts[0]
		if name == "noop" {
			instruction = Instruction{Name: parts[0]}
		} else {
			instruction = Instruction{
				Name:   parts[0],
				Amount: utils.ParseInt(parts[1]),
			}
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}
