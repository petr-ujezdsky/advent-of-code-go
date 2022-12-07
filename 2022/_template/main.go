package main

import (
	"bufio"
	_ "embed"
	"io"
)

type Item struct {
}

func DoWithInput(_ []Item) int {
	return 0
}

func ParseInput(r io.Reader) []Item {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var items []Item
	for scanner.Scan() {
		//parts := strings.Split(scanner.Text(), ",")
		//ints := utils.ExtractInts(scanner.Text(), false)

		item := Item{}

		items = append(items, item)
	}

	return items
}
