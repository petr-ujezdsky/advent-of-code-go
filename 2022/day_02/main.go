package main

import (
	"bufio"
	_ "embed"
	"io"
	"strings"
)

type Round = []byte

func decrypt(choice byte) byte {
	return choice - ('X' - 'A')
}

func outcomeScore(p1, p2 byte) int {
	diff := int(p1) - int(p2)

	// same choices - draw
	if diff == 0 {
		return 3
	}

	// p1 wins
	if diff == 1 || diff == -2 {
		return 6
	}

	// p1 looses
	if diff == -1 || diff == 2 {
		return 0
	}

	panic("Unknown outcome")
}

func choiceScore(choice byte) int {
	return int(choice - 'A' + 1)
}

func Score(rounds []Round) int {
	sum := 0
	for _, round := range rounds {
		sum += choiceScore(decrypt(round[1]))
		sum += outcomeScore(decrypt(round[1]), round[0])
	}

	return sum
}

func ParseInput(r io.Reader) []Round {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rounds []Round

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		round := Round{parts[0][0], parts[1][0]}
		rounds = append(rounds, round)
	}

	return rounds
}
