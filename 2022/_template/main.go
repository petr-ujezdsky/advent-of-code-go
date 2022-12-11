package main

import (
	"bufio"
	_ "embed"
	"io"
)

type Item struct {
}

func DoWithInput(items []Item) int {
	return len(items)
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
