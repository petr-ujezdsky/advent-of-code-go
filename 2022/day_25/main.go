package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

func char2decimal(char rune) int {
	switch char {
	case '=':
		return -2
	case '-':
		return -1
	}

	return int(char - '0')
}

func ParseSNAFU(str string) int {
	reversed := slices.Reverse([]rune(str))

	sum := 0
	k := 1
	for _, char := range reversed {
		sum += k * char2decimal(char)
		k *= 5
	}

	return sum
}

func DoWithInput(snafus []string) int {
	return len(snafus)
}

func ParseInput(r io.Reader) []string {
	return utils.ParseToStringsP(r)
}
