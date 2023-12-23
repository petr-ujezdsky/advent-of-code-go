package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strings"
)

type DigOrder struct {
	Direction utils.Direction4
	Amount    int
	Color     string
}

type TrenchCube struct {
}

type World struct {
	Orders  []DigOrder
	Borders utils.BoundingRectangle
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func parseDir(s string) utils.Direction4 {
	switch s {
	case "U":
		return utils.Up
	case "R":
		return utils.Right
	case "D":
		return utils.Down
	case "L":
		return utils.Left
	}

	panic("Unknown direction")
}
func ParseInput(r io.Reader) World {
	parseItem := func(str string) DigOrder {
		parts := strings.Split(str, " ")

		direction := parseDir(parts[0])
		amount := utils.ParseInt(parts[1])
		color := parts[2][1 : len(parts[2])-1]

		return DigOrder{
			Direction: direction,
			Amount:    amount,
			Color:     color,
		}
	}

	orders := parsers.ParseToObjects(r, parseItem)
	return World{Orders: orders}
}
