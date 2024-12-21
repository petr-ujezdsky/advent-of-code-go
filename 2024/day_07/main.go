package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strconv"
)

type Equation struct {
	Result  int
	Numbers []int
}

func (e Equation) String() string {
	return fmt.Sprintf("%v: %v", e.Result, e.Numbers)
}

type World struct {
	Equations []Equation
}

func equationSolvable(equation Equation, index, result int, concatEnabled bool) bool {
	if index == len(equation.Numbers) {
		return result == equation.Result
	}

	number := equation.Numbers[index]

	maxAdd := equation.Result - result
	if number <= maxAdd {
		// try add operation
		nextResult := result + number
		if equationSolvable(equation, index+1, nextResult, concatEnabled) {
			return true
		}
	}

	if result != 0 {
		maxMul := equation.Result / result
		if number <= maxMul {
			// try multiply operation
			nextResult := result * number
			if equationSolvable(equation, index+1, nextResult, concatEnabled) {
				return true
			}
		}

		if concatEnabled {
			nextResult := utils.ParseInt(strconv.Itoa(result) + strconv.Itoa(number))

			if equationSolvable(equation, index+1, nextResult, concatEnabled) {
				return true
			}
		}
	}

	return false
}

func DoWithInputPart01(world World) int {
	sum := 0
	for _, equation := range world.Equations {
		if equationSolvable(equation, 0, 0, false) {
			sum += equation.Result
			fmt.Printf("✅ %v\n", equation)
		} else {
			fmt.Printf("❌ %v\n", equation)
		}
	}
	return sum
}

func DoWithInputPart02(world World) int {
	sum := 0
	for _, equation := range world.Equations {
		if equationSolvable(equation, 0, 0, true) {
			sum += equation.Result
			fmt.Printf("✅ %v\n", equation)
		} else {
			fmt.Printf("❌ %v\n", equation)
		}
	}
	return sum
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
