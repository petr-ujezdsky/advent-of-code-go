package day_10

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type InputRows []string

type Char struct {
	Value   rune
	Opening bool
	Group   int
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

func checkCorruption(row string) (Char, int, bool) {
	openersStack := utils.NewStack[Char]()

	for pos, charRune := range []rune(row) {
		char := NewChar(charRune)

		if char.Opening {
			openersStack.Push(char)
		} else {
			// closing at wrong time
			if openersStack.Empty() {
				return char, pos, true
			}

			previous := openersStack.Peek()
			// closing wrong char
			if previous.Group != char.Group {
				return char, pos, true
			}

			// remove opening char
			openersStack.Pop()
		}
	}

	var empty Char
	return empty, -1, false
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

func ParseInput(r io.Reader) (InputRows, error) {
	return utils.ParseToStrings(r)
}
