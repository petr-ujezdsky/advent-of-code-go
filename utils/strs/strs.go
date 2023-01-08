package strs

import "github.com/petr-ujezdsky/advent-of-code-go/utils/slices"

func ReverseString(str string) string {
	return string(slices.Reverse([]rune(str)))
}

func Substring(str string, from, to int) string {
	return string(([]rune(str))[from:to])
}
