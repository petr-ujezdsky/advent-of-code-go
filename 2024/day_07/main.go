package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Equation struct {
	Result  int
	Numbers []int
}

type World struct {
	Equations []Equation
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) Equation {
		ints := utils.ExtractInts(str, false)
		return Equation{
			Result:  ints[0],
			Numbers: ints[1:],
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Equations: items}
}
