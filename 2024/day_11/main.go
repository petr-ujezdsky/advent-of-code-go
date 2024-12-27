package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Stone struct {
	Number int
}

type World struct {
	Stones []Stone
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var stones []Stone
	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)

		for _, value := range ints {
			stones = append(stones, Stone{Number: value})
		}
	}

	return World{Stones: stones}
}
