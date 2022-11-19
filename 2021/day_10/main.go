package day_10

import (
	"bufio"
	"io"
)

type InputRows [][]rune

func ParseInput(r io.Reader) (InputRows, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rows InputRows

	for scanner.Scan() {
		line := scanner.Text()

		rows = append(rows, []rune(line))
	}

	return rows, scanner.Err()
}
