package main

import (
	"bufio"
	_ "embed"
	"io"
)

type Item struct {
}

func DoWithInputPart01(items []Item) int {
	return len(items)
}

func DoWithInputPart02(items []Item) int {
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

//func ParseInput(r io.Reader) []Item {
//	parseItem := func(str string) Item {
//		return Item{}
//	}
//
//	return parsers.ParseToObjects(r, parseItem)
//}

//func ParseInput(r io.Reader) utils.Matrix[Item] {
//	parseItem := func(char rune) Item {
//		return Item{}
//	}
//
//	return parsers.ParseToMatrix(r, parseItem)
//}
