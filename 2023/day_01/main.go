package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/strs"
	"io"
	"strings"
)

type World struct {
	Rows []string
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, row := range world.Rows {
		sum += extractNumber(row)
	}

	return sum
}

func extractNumber(row string) int {
	var first, last rune

	for _, char := range []rune(row) {
		if char >= '0' && char <= '9' {
			if first == 0 {
				first = char
			}
			last = char
		}
	}

	return utils.ParseInt(string([]rune{first, last}))
}

func DoWithInputPart02(world World) int {
	sum := 0

	for _, row := range world.Rows {
		sum += extractNumberWords(row)
	}

	return sum
}

var dictionaryForward = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,

	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var dictionaryBackward = map[string]int{
	"eno":   1,
	"owt":   2,
	"eerht": 3,
	"ruof":  4,
	"evif":  5,
	"xis":   6,
	"neves": 7,
	"thgie": 8,
	"enin":  9,

	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

func extractNumberWords(row string) int {
	first := findFirstNumber(row, dictionaryForward)
	last := findFirstNumber(strs.ReverseString(row), dictionaryBackward)

	return first*10 + last
}

func findFirstNumber(str string, dic map[string]int) int {
	index := 0
	for {
		if index >= len(str) {
			break
		}

		for text, digit := range dic {
			if strings.Index(str[index:], text) == 0 {
				return digit
			}
		}

		index++
	}

	panic("No number found")
}

func ParseInput(r io.Reader) World {
	items := parsers.ParseToStrings(r)
	return World{Rows: items}
}
