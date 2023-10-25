package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Item struct {
}

type World struct {
	Items []Item
	//Matrix utils.Matrix[Item]
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) Item {
		return Item{}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Items: items}
}

//func ParseInput(r io.Reader) World {
//	scanner := bufio.NewScanner(r)
//	scanner.Split(bufio.ScanLines)
//
//	var items []Item
//	for scanner.Scan() {
//		//parts := strings.Split(scanner.Text(), ",")
//		//ints := utils.ExtractInts(scanner.Text(), false)
//
//		item := Item{}
//
//		items = append(items, item)
//	}
//
//	return World{Items: items}
//}
//
//func ParseInput(r io.Reader) World {
//	parseItem := func(char rune) Item {
//		return Item{}
//	}
//
//	return World{Matrix: parsers.ParseToMatrix(r, parseItem)}
//}
