package main

import (
	"bufio"
	_ "embed"
	"io"
)

type World struct {
	DiskMap []int
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

	var numbers []int
	for scanner.Scan() {
		for _, numberChar := range scanner.Text() {
			numbers = append(numbers, int(numberChar-'0'))
		}
	}

	return World{DiskMap: numbers}
}
