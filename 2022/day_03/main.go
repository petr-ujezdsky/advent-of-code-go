package main

import (
	"bufio"
	_ "embed"
	"io"
)

type Rucksack struct {
	Left, Right []rune
}

func NewRucksack(text string) Rucksack {
	line := []rune(text)

	return Rucksack{
		Left:  line[:len(line)/2],
		Right: line[len(line)/2:],
	}
}

func commonItem(rucksack Rucksack) rune {
	leftItems := make(map[rune]struct{})
	for _, leftItem := range rucksack.Left {
		leftItems[leftItem] = struct{}{}
	}

	for _, rightItem := range rucksack.Right {
		_, ok := leftItems[rightItem]
		if ok {
			return rightItem
		}
	}

	panic("No common item found")
}

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func itemScore(r rune) int {
	if 'a' <= r && r <= 'z' {
		return int(r - 'a' + 1)
	}

	return int(r - 'A' + 27)
}

func Score(rucksacks []Rucksack) int {
	score := 0

	for _, rucksack := range rucksacks {
		ci := commonItem(rucksack)
		score += itemScore(ci)
	}

	return score
}

func ParseInput(r io.Reader) []Rucksack {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rucksacks []Rucksack

	for scanner.Scan() {
		rucksacks = append(rucksacks, NewRucksack(scanner.Text()))
	}

	return rucksacks
}
