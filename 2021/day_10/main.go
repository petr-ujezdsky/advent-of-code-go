package day_10

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"sort"
)

type InputRows []string

type Char struct {
	Value       rune
	Counterpart rune
	Opening     bool
	Group       int
}

func NewChar(char rune) Char {
	return Char{
		Value:   char,
		Opening: charOpening(char),
		Group:   charGroup(char),
	}
}

func charOpening(char rune) bool {
	switch char {
	case '(', '[', '{', '<':
		return true
	case ')', ']', '}', '>':
		return false
	default:
		panic("Char '" + string(char) + "' is not valid")
	}
}

func charGroup(char rune) int {
	switch char {
	case '(', ')':
		return 1
	case '[', ']':
		return 2
	case '{', '}':
		return 3
	case '<', '>':
		return 4
	default:
		panic("Char '" + string(char) + "' is not valid")
	}
}

//func charCloseCounterpart(char rune) rune {
//	switch char {
//	case '(':
//		return ')'
//	case '[':
//		return ']'
//	case '{':
//		return '}'
//	case '<':
//		return '>'
//	default:
//		panic("Char '" + string(char) + "' is not valid")
//	}
//}

func charScore(char rune) int {
	switch char {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		panic("Char '" + string(char) + "' is not valid")
	}
}

func charIncompleteScore(char rune) int {
	switch char {
	case '(':
		return 1
	case '[':
		return 2
	case '{':
		return 3
	case '<':
		return 4
	default:
		panic("Char '" + string(char) + "' is not valid")
	}
}

func checkCorruption(row string) (Char, collections.Stack[Char], bool) {
	openersStack := collections.NewStack[Char]()

	for _, charRune := range []rune(row) {
		char := NewChar(charRune)

		if char.Opening {
			openersStack.Push(char)
		} else {
			// closing at wrong time
			if openersStack.Empty() {
				return char, openersStack, true
			}

			previous := openersStack.Peek()
			// closing wrong char
			if previous.Group != char.Group {
				return char, openersStack, true
			}

			// remove opening char
			openersStack.Pop()
		}
	}

	var empty Char
	return empty, openersStack, false
}

func CorruptionScore(rows InputRows) int {
	totalScore := 0

	for _, row := range rows {
		char, _, corrupted := checkCorruption(row)
		if corrupted {
			totalScore += charScore(char.Value)
		}
	}

	return totalScore
}

func IncompleteScore(rows InputRows) int {
	var totalScores []int

	for _, row := range rows {
		_, openersStack, corrupted := checkCorruption(row)

		if !corrupted {
			totalScore := 0
			for !openersStack.Empty() {
				char := openersStack.Pop()
				totalScore = totalScore*5 + charIncompleteScore(char.Value)
			}
			totalScores = append(totalScores, totalScore)
		}
	}

	sort.Ints(totalScores)

	return totalScores[len(totalScores)/2]
}

func ParseInput(r io.Reader) (InputRows, error) {
	return parsers.ParseToStrings(r)
}
