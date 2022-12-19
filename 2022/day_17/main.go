package main

import (
	"bufio"
	_ "embed"
	"io"
)

type JetDirection = int

func DoWithInput(items []JetDirection) int {
	return len(items)
}

func ParseInput(r io.Reader) []JetDirection {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var directions []JetDirection
	for scanner.Scan() {
		directions = make([]JetDirection, len(scanner.Text()))

		for i, char := range scanner.Text() {
			if char == '<' {
				directions[i] = -1
			} else {
				directions[i] = 1
			}
		}
	}

	return directions
}
