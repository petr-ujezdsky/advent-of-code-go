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
	trench, bounds := WalkOrders(world.Orders)

	return LagoonArea(trench, bounds)
}

func WalkOrders(orders []DigOrder) (map[utils.Vector2i]DigOrder, utils.BoundingRectangle) {
	trench := make(map[utils.Vector2i]DigOrder)

	position := utils.Vector2i{}

	bounds := utils.NewBoundingRectangleFromPoints(position, position)

	// origin
	trench[position] = DigOrder{
		Direction: -1,
		Amount:    -1,
		Color:     "",
	}

	for _, order := range orders {
		step := order.Direction.ToStep()

		for i := 0; i < order.Amount; i++ {
			nextPosition := position.Add(step)

			trench[nextPosition] = order
			bounds = bounds.Enlarge(nextPosition)

			position = nextPosition
		}
	}

	return trench, bounds
}

func LagoonArea(trench map[utils.Vector2i]DigOrder, bounds utils.BoundingRectangle) int {
	// use rendering algorithm
	area := 0

	for x := bounds.Horizontal.Low; x <= bounds.Horizontal.High; x++ {
		previous := '.'
		crossingsCount := 0
		for y := bounds.Vertical.Low; y <= bounds.Vertical.High; y++ {
			position := utils.Vector2i{X: x, Y: y}

			current := '.'
			if _, ok := trench[position]; ok {
				current = '#'
			}

			if previous != current {
				crossingsCount++
			}

			if crossingsCount%2 == 1 {
				// inside
				area++
			}

			previous = current
		}
	}

	return area
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
