package parsers

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

func MapperBoolean(trueChar, falseChar rune) func(ch rune, i, j int) bool {
	return func(ch rune, i, j int) bool {
		if ch == trueChar {
			return true
		}

		if ch == falseChar {
			return false
		}

		panic("Unknown char " + string(ch))
	}
}

// ParseToMatrix returns the matrix of objects
func ParseToMatrix[T any](r io.Reader, mapper func(ch rune) T) utils.Matrix[T] {
	indexedMapper := func(line rune, i, j int) T { return mapper(line) }
	return ParseToMatrixIndexed(r, indexedMapper)

}

// ParseToMatrixIndexed returns the matrix of objects, uses row and column index
func ParseToMatrixIndexed[T any](r io.Reader, mapper func(ch rune, i, j int) T) utils.Matrix[T] {
	lineMapper := func(line string, i int) []T {
		var row []T
		j := 0
		for _, char := range []rune(line) {
			item := mapper(char, i, j)
			row = append(row, item)
			j++
		}

		return row
	}

	rows := ParseToObjectsIndexed(r, lineMapper)

	return utils.NewMatrixRowNotation[T](rows)
}

// ParseToObjects returns slice of objects mapped from rows
func ParseToObjects[T any](r io.Reader, mapper func(line string) T) []T {
	indexedMapper := func(line string, i int) T { return mapper(line) }
	return ParseToObjectsIndexed(r, indexedMapper)
}

// ParseToObjectsIndexed returns slice of objects mapped from rows, uses row index
func ParseToObjectsIndexed[T any](r io.Reader, mapper func(line string, i int) T) []T {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var objects []T

	i := 0
	for scanner.Scan() {
		object := mapper(scanner.Text(), i)

		objects = append(objects, object)
	}

	if scanner.Err() != nil {
		panic("Error parsing input")
	}

	return objects
}
