package main

import (
	"bufio"
	_ "embed"
	"io"
)

type Round = []rune

func decrypt(choice rune) rune {
	return choice - ('X' - 'A')
}

// A/X - Rock
// B/Y - Paper
// C/Z - Scissors
func outcomeScore(round string) int {
	switch round {
	case "AA", "BB", "CC":
		// same choices - draw
		return 3
	case "CA", "AB", "BC":
		// p2 wins
		return 6
	case "BA", "CB", "AC":
		// p2 looses
		return 0
	}

	panic("Unknown outcome")
}

func choiceScore(choice rune) int {
	return int(choice - 'A' + 1)
}

func Score(rounds []Round) int {
	sum := 0
	for _, round := range rounds {
		sum += choiceScore(round[1])
		sum += outcomeScore(string(round))
	}

	return sum
}

func ParseInput(r io.Reader) []Round {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rounds []Round

	for scanner.Scan() {
		line := []rune(scanner.Text())

		round := Round{line[0], decrypt(line[2])}
		rounds = append(rounds, round)
	}

	return rounds
}
