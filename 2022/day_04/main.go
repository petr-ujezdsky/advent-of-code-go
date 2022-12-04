package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"regexp"
)

type Pair struct {
	Left, Right utils.IntervalI
}

var regexPair = regexp.MustCompile("(\\d+)-(\\d+),(\\d+)-(\\d+)")

func CountContained(pairs []Pair) int {
	count := 0
	for _, pair := range pairs {
		intersectionType, _ := pair.Left.IntersectionDetail(pair.Right)
		if intersectionType == utils.Inside || intersectionType == utils.Wraps || intersectionType == utils.Identical {
			count++
		}
	}

	return count
}
func CountOverlapped(pairs []Pair) int {
	count := 0
	for _, pair := range pairs {
		intersectionType, _ := pair.Left.IntersectionDetail(pair.Right)
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
			Left:  utils.NewInterval(utils.ParseInt(tokens[1]), utils.ParseInt(tokens[2])),
			Right: utils.NewInterval(utils.ParseInt(tokens[3]), utils.ParseInt(tokens[4])),
		}

		pairs = append(pairs, pair)
	}

	return pairs
}
