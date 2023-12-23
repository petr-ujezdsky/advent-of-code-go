package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strings"
)

type DigOrder struct {
	Direction      utils.Direction4
	Amount         int
	Color          string
	Previous, Next *DigOrder
}

type World struct {
	Orders  []*DigOrder
	Borders utils.BoundingRectangle
}

func DoWithInputPart01(world World) int {
	trench, bounds := WalkOrders(world.Orders)

	return LagoonArea(trench, bounds)
}

func WalkOrders(orders []*DigOrder) (map[utils.Vector2i]*DigOrder, utils.BoundingRectangle) {
	trench := make(map[utils.Vector2i]*DigOrder)

	position := utils.Vector2i{}

	bounds := utils.NewBoundingRectangleFromPoints(position, position)

	// origin
	trench[position] = &DigOrder{
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

func LagoonArea(trench map[utils.Vector2i]*DigOrder, bounds utils.BoundingRectangle) int {
	// use rendering algorithm
	area := 0

	m := matrix.NewMatrix[rune](bounds.Width(), bounds.Height())
	mOrigin := utils.Vector2i{X: bounds.Horizontal.Low, Y: bounds.Vertical.Low}

	for x := bounds.Horizontal.Low; x <= bounds.Horizontal.High; x++ {
		previous := '.'
		crossingsCount := 0
		areaY := 0
		for y := bounds.Vertical.Low; y <= bounds.Vertical.High; y++ {
			position := utils.Vector2i{X: x, Y: y}

			current := '.'
			if t, ok := trench[position]; ok {
				current = '#'

				// detect vertical
				if t.Direction == utils.Up || t.Direction == utils.Down {

				}
			}

			m.SetV(position.Subtract(mOrigin), current)

			if current == '#' && previous != current {
				crossingsCount++
			}

			if current == '#' || crossingsCount%2 == 1 {
				// inside
				areaY++
			}

			previous = current
		}
		fmt.Printf("x=%d  area %d\n", x, areaY)

		area += areaY
	}

	//m = m.FlipVertical()

	str := matrix.StringFmtSeparatorIndexedOrigin[rune](m, true, mOrigin, "", func(r rune, x, y int) string {
		return string(r)
	})

	fmt.Printf("Lagoon:\n\n%v\n", str)

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
	parseItem := func(str string) *DigOrder {
		parts := strings.Split(str, " ")

		direction := parseDir(parts[0])
		amount := utils.ParseInt(parts[1])
		color := parts[2][1 : len(parts[2])-1]

		return &DigOrder{
			Direction: direction,
			Amount:    amount,
			Color:     color,
		}
	}

	orders := parsers.ParseToObjects(r, parseItem)

	// link them
	previous := orders[len(orders)-1]
	for _, current := range orders {
		current.Previous = previous
		previous.Next = current

		previous = current
	}

	return World{Orders: orders}
}
