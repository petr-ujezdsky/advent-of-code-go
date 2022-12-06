package main

import (
	"bufio"
	_ "embed"
	"io"
)

func FindPacketStart(str string) int {
	chars := []rune(str)
	window := make(map[rune]int)
	leftChar := rune(0)

	for i, ch := range chars {
		window[ch]++

		if len(window) == 4 {
			return i + 1
		}

		if i > 2 {
			leftChar = chars[i-3]
			window[leftChar]--
			if window[leftChar] == 0 {
				delete(window, leftChar)
			}
		}
	}

	panic("No marker found!")
}

func ParseInput(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		return scanner.Text()
	}

	panic("No lines")
}
