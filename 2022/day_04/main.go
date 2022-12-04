package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"regexp"
)

type Range struct {
	Min, Max int
}

type Pair struct {
	Left, Right Range
}

var regexPair = regexp.MustCompile("(\\d+)-(\\d+),(\\d+)-(\\d+)")

func CountContained(pairs []Pair) int {
	count := 0
	for _, pair := range pairs {
		intersectionType, _, _ := utils.IntervalIntersectionDetail(pair.Left.Min, pair.Left.Max, pair.Right.Min, pair.Right.Max)
		if intersectionType == utils.Inside || intersectionType == utils.Wraps || intersectionType == utils.Identical {
			count++
		}
	}

	return count
}
func CountOverlapped(pairs []Pair) int {
	count := 0
	for _, pair := range pairs {
		intersectionType, _, _ := utils.IntervalIntersectionDetail(pair.Left.Min, pair.Left.Max, pair.Right.Min, pair.Right.Max)
		if intersectionType != utils.None {
			count++
		}
	}

	return count
}

func ParseInput(r io.Reader) []Pair {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var pairs []Pair

	for scanner.Scan() {
		tokens := regexPair.FindStringSubmatch(scanner.Text())
		pair := Pair{
			Left:  Range{utils.ParseInt(tokens[1]), utils.ParseInt(tokens[2])},
			Right: Range{utils.ParseInt(tokens[3]), utils.ParseInt(tokens[4])},
		}

		pairs = append(pairs, pair)
	}

	return pairs
}
