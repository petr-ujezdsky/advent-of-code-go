package main

import (
	"bufio"
	_ "embed"
	"io"
	"strings"
)

type Rucksack struct {
	Left, Right []rune
	All         string
}

func NewRucksack(text string) Rucksack {
	line := []rune(text)

	return Rucksack{
		Left:  line[:len(line)/2],
		Right: line[len(line)/2:],
		All:   text,
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

func commonCharsTwo(chars map[rune]struct{}, s string) map[rune]struct{} {
	common := make(map[rune]struct{})

	for ch := range chars {
		if strings.ContainsRune(s, ch) {
			common[ch] = struct{}{}
		}
	}

	return common
}

func commonChar(strings []string) rune {
	common := make(map[rune]struct{})
	for _, ch := range strings[0] {
		common[ch] = struct{}{}
	}

	for i := 1; i < len(strings); i++ {
		s := strings[i]
		common = commonCharsTwo(common, s)
	}

	if len(common) > 1 {
		panic("Too many common chars!")
	}

	for ch := range common {
		return ch
	}

	panic("No common chars!")
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

func GroupsScore(rucksacks []Rucksack) int {
	score := 0

	for i := 0; i < len(rucksacks)-2; i += 3 {
		r1 := rucksacks[i]
		r2 := rucksacks[i+1]
		r3 := rucksacks[i+2]

		common := commonChar([]string{r1.All, r2.All, r3.All})
		score += itemScore(common)
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
