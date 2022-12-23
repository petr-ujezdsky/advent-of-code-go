package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Elf utils.Vector2i

type World map[Elf]struct{}

func DoWithInput(elves World) int {
	return len(elves)
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	elves := make(World)

	y := 0
	for scanner.Scan() {
		for x, char := range scanner.Text() {
			if char == '.' {
				continue
			}

			elf := Elf{X: x, Y: y}
			elves[elf] = struct{}{}
		}
		y++
	}

	return elves
}
