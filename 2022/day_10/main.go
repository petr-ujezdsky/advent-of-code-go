package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type MatrixB = utils.Matrix[bool]

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

func drawPixel(cycle, x int, pixels MatrixB) {
	pos := cycle - 1
	px := pos % 40
	py := pos / 40

	if x-1 <= px && px <= x+1 {
		pixels.Columns[px][py] = true
	} else {
		pixels.Columns[px][py] = false
	}
}

func DoWithInput(instructions []Instruction) (int, MatrixB) {
	pixels := utils.NewMatrix[bool](40, 6)
	strengthSum := 0
	cycle := 0
	x := 1
	for i, instruction := range instructions {
		if instruction.Name == "noop" {
			cycle++
			strengthSum += readSignalStrength(cycle, x)
			drawPixel(cycle, x, pixels)
			continue
		}

		cycle++
		strengthSum += readSignalStrength(cycle, x)
		drawPixel(cycle, x, pixels)
		cycle++
		strengthSum += readSignalStrength(cycle, x)
		drawPixel(cycle, x, pixels)
		x += instruction.Amount
		_ = i
	}

	return strengthSum, pixels
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
