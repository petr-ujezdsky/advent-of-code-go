package parsers

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

func MapperBoolean(trueChar, falseChar rune) func(ch rune) bool {
	return func(ch rune) bool {
		if ch == trueChar {
			return true
		}

		if ch == falseChar {
			return false
		}

		panic("Unknown char " + string(ch))
	}
}

// ParseToMatrix returns the matrix of integers
func ParseToMatrix[T any](r io.Reader, mapper func(ch rune) T) utils.Matrix[T] {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rows [][]T

	for scanner.Scan() {
		line := scanner.Text()
		var row []T

		for _, char := range []rune(line) {
			item := mapper(char)
			row = append(row, item)
		}

		rows = append(rows, row)
	}

	if scanner.Err() != nil {
		panic("Error parsing matrix")
	}

	return utils.NewMatrixRowNotation[T](rows)
}
