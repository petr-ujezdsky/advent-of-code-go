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

type Mul struct {
	Left, Right int
}

var regexMul = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func extractMuls(str string) []Mul {
	var muls []Mul

	for _, mulRaw := range regexMul.FindAllStringSubmatch(str, -1) {
		left := utils.ParseInt(mulRaw[1])
		right := utils.ParseInt(mulRaw[2])

		muls = append(muls, Mul{
			Left:  left,
			Right: right,
		})
	}

	return muls
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, row := range world.Rows {
		for _, mul := range extractMuls(row) {
			sum += mul.Left * mul.Right
		}
	}

	return sum
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	return World{Rows: parsers.ParseToStrings(r)}
}
