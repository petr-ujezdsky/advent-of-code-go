package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strconv"
	"strings"
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

func toBase5Reversed(n int) []int {
	var base5 []int

	for {
		digit := n % 5
		base5 = append(base5, digit)
		n /= 5

		if n == 0 {
			break
		}
	}

	return base5
}

func base5ReversedToSNAFU(base5 []int) string {
	snafu := &strings.Builder{}

	transfer := 0
	for _, digit := range base5 {
		digit += transfer
		transfer = 0

		if digit > 4 {
			digit = 0
			transfer = 1
		}

		switch digit {
		case 0, 1, 2:
			snafu.WriteString(strconv.Itoa(digit))
		case 3:
			transfer = 1
			snafu.WriteRune('=')
		case 4:
			transfer = 1
			snafu.WriteRune('-')
		}
	}

	if transfer > 0 {
		snafu.WriteString(strconv.Itoa(transfer))
	}

	return utils.ReverseString(snafu.String())
}

func CreateSNAFU(n int) string {
	return base5ReversedToSNAFU(toBase5Reversed(n))
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
