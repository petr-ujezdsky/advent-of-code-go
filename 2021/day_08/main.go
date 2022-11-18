package day_07

import (
	"bufio"
	"io"
	"strings"
)

// Entry
// List of digits and their segments count, * marks unique
// 0 ~ 6 segments
// 1 ~ 2 segments *
// 2 ~ 5 segments
// 3 ~ 5 segments
// 4 ~ 4 segments *
// 5 ~ 5 segments
// 6 ~ 6 segments
// 7 ~ 3 segments *
// 8 ~ 7 segments *
// 9 ~ 6 segments
type Entry struct {
	Digits  []string
	Outputs []string
}

func CountEasyOutputs(entries []Entry) int {
	// easy digit lengths:
	// 1 ~ 2 segments
	// 4 ~ 4 segments
	// 7 ~ 3 segments
	// 8 ~ 7 segments
	count := 0

	for _, entry := range entries {
		for _, output := range entry.Outputs {
			length := len(output)

			if length == 2 || length == 4 || length == 3 || length == 7 {
				count++
			}
		}
	}

	return count
}

func ParseInput(r io.Reader) ([]Entry, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var entries []Entry

	// example line
	// be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " | ")

		digits := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		entries = append(entries, Entry{digits, output})
	}

	return entries, scanner.Err()
}
