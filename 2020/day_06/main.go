package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type TrueAnswers = map[rune]int

type Group struct {
	TrueAnswers TrueAnswers
	Size        int
}

func CountTrueAnswersPerGroupAnyone(groups []Group) int {
	sum := 0

	for _, group := range groups {
		sum += len(group.TrueAnswers)
	}

	return sum
}

func CountTrueAnswersPerGroupEveryone(groups []Group) int {
	sum := 0

	for _, group := range groups {
		for _, count := range group.TrueAnswers {
			if count == group.Size {
				sum++
			}
		}
	}

	return sum
}

func parseVotes(str string, trueAnswers TrueAnswers) {
	for _, char := range str {
		trueAnswers[char]++
	}
}

func parseGroup(lines []string, _ int) Group {
	trueAnswers := make(TrueAnswers)

	for _, line := range lines {
		parseVotes(line, trueAnswers)
	}

	return Group{
		TrueAnswers: trueAnswers,
		Size:        len(lines),
	}
}

func ParseInput(r io.Reader) []Group {
	return parsers.ParseToGroups(r, parseGroup)
}
