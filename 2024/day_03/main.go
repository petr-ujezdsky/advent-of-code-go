package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"regexp"
)

type World struct {
	Rows []string
}

var regexMulti = regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)|(don't)\(\)|(do)\(\)`)

func DoWithInputPart01(world World) int {
	sum := 0

	for _, row := range world.Rows {
		for _, parsed := range regexMulti.FindAllStringSubmatch(row, -1) {
			op := parsed[1] + parsed[4] + parsed[5]

			switch op {
			case "mul":
				sum += utils.ParseInt(parsed[2]) * utils.ParseInt(parsed[3])
			}
		}
	}

	return sum
}

func DoWithInputPart02(world World) int {
	enabled := true
	sum := 0

	for _, row := range world.Rows {
		for _, parsed := range regexMulti.FindAllStringSubmatch(row, -1) {
			op := parsed[1] + parsed[4] + parsed[5]

			switch op {
			case "mul":
				if enabled {
					sum += utils.ParseInt(parsed[2]) * utils.ParseInt(parsed[3])
				}
			case "don't":
				enabled = false
			case "do":
				enabled = true
			}
		}
	}

	return sum
}

func ParseInput(r io.Reader) World {
	return World{Rows: parsers.ParseToStrings(r)}
}
