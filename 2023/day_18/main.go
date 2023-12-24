package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"sort"
	"strconv"
	"strings"
)

type DigOrder struct {
	Raw            string
	Direction      utils.Direction4
	Amount         int
	Color          string
	Previous, Next *DigOrder
}

type TrenchSegment struct {
	Line     utils.LineOrthogonal2i
	DigOrder *DigOrder
}

type IntersectingSegment struct {
	IntersectingLine utils.LineOrthogonal2i
	TrenchSegment    TrenchSegment
}

type World struct {
	Orders  []*DigOrder
	Borders utils.BoundingRectangle
}

func DoWithInputPart01(world World) int {
	trench, bounds := WalkOrders(world.Orders)

	return LagoonArea(trench, bounds)
}

func WalkOrders(orders []*DigOrder) ([]TrenchSegment, utils.BoundingRectangle) {
	var trench []TrenchSegment

	position := utils.Vector2i{}

	bounds := utils.NewBoundingRectangleFromPoints(position, position)

	for _, order := range orders {
		step := order.Direction.ToStep()

		to := position.Add(step.Multiply(order.Amount))
		line := utils.NewLineOrthogonal2i(position, to)

		bounds = bounds.Enlarge(to)

		trench = append(trench, TrenchSegment{
			Line:     line,
			DigOrder: order,
		})

		position = to
	}

	return trench, bounds
}

func LagoonArea(trench []TrenchSegment, bounds utils.BoundingRectangle) int {
	// use rendering algorithm
	area := 0

	//m := matrix.NewMatrix[rune](bounds.Width(), bounds.Height())
	//mOrigin := utils.Vector2i{X: bounds.Horizontal.Low, Y: bounds.Vertical.Low}

	for x := bounds.Horizontal.Low; x <= bounds.Horizontal.High; x++ {
		from := utils.Vector2i{X: x, Y: bounds.Vertical.Low}
		to := utils.Vector2i{X: x, Y: bounds.Vertical.High}
		scanLine := utils.NewLineOrthogonal2i(from, to)

		var intersectingSegments []IntersectingSegment
		for _, segment := range trench {
			if intersection, ok := scanLine.Intersection(segment.Line); ok {
				// intersection is edge point
				if intersection.IsPoint() && (intersection.A == segment.Line.A || intersection.A == segment.Line.B) {
					// skip - solve by subsequent line intersection
					continue
				}

				intersectingSegments = append(intersectingSegments, IntersectingSegment{
					IntersectingLine: intersection,
					TrenchSegment:    segment,
				})
			}
		}

		// sort by Y axis
		sort.Slice(intersectingSegments, func(i, j int) bool {
			return intersectingSegments[i].IntersectingLine.A.Y < intersectingSegments[j].IntersectingLine.A.Y
		})

		yArea := 0
		inside := false
		lastInsideY := 0
		for _, segment := range intersectingSegments {

			aggregateArea := true
			if segment.IntersectingLine.IsPoint() {
				inside = !inside
			} else {
				// intersection is line, check the shape

				// "Z" shape - act as "point"
				digOrder := segment.TrenchSegment.DigOrder
				if digOrder.Previous.Direction == digOrder.Next.Direction {
					inside = !inside
				} else {
					// "C" shape - change nothing

					// aggregate segment
					if inside {
						// current subarea segment
						yArea += utils.Min(segment.IntersectingLine.A.Y, segment.IntersectingLine.B.Y) - lastInsideY
					} else {
						aggregateArea = false
					}
				}
			}

			// aggregate intersection segment
			yArea += segment.IntersectingLine.Length()

			currentInsideY := utils.Max(segment.IntersectingLine.A.Y, segment.IntersectingLine.B.Y) + 1

			if !inside && aggregateArea {
				// aggregate area
				yArea += currentInsideY - lastInsideY - segment.IntersectingLine.Length()
			}

			lastInsideY = currentInsideY
		}

		//
		//areaY := 0
		//for y := bounds.Vertical.Low; y <= bounds.Vertical.High; y++ {
		//	position := utils.Vector2i{X: x, Y: y}
		//
		//	current := '.'
		//	if t, ok := trench[position]; ok {
		//		current = '#'
		//
		//		// detect vertical
		//		if t.Direction == utils.Up || t.Direction == utils.Down {
		//
		//		}
		//	}
		//
		//	m.SetV(position.Subtract(mOrigin), current)
		//
		//	if current == '#' && previous != current {
		//		crossingsCount++
		//	}
		//
		//	if current == '#' || crossingsCount%2 == 1 {
		//		// inside
		//		areaY++
		//	}
		//
		//	previous = current
		//}
		//fmt.Printf("x=%d  area %d\n", x, yArea)

		area += yArea
	}

	//m = m.FlipVertical()

	//str := matrix.StringFmtSeparatorIndexedOrigin[rune](m, true, mOrigin, "", func(r rune, x, y int) string {
	//	return string(r)
	//})
	//
	//fmt.Printf("Lagoon:\n\n%v\n", str)

	return area
}

func DoWithInputPart02(world World) int {
	fixed := fixDigOrders(world.Orders)
	trench, bounds := WalkOrders(fixed)

	return LagoonArea(trench, bounds)
}

func fixDigOrders(orders []*DigOrder) []*DigOrder {
	fixed := slices.Map(orders, func(broken *DigOrder) *DigOrder {
		amount, _ := strconv.ParseInt(broken.Color[1:6], 16, 32)
		dir, _ := strconv.ParseInt(broken.Color[6:], 16, 32)
		direction := parseDirInt(int(dir))

		return &DigOrder{
			Raw:       broken.Raw,
			Direction: direction,
			Amount:    int(amount),
			Color:     broken.Color,
			Previous:  nil,
			Next:      nil,
		}
	})

	// link them
	previous := fixed[len(fixed)-1]
	for _, current := range fixed {
		current.Previous = previous
		previous.Next = current

		previous = current
	}

	return fixed
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

func parseDirInt(i int) utils.Direction4 {
	switch i {
	case 0:
		return utils.Right
	case 1:
		return utils.Down
	case 2:
		return utils.Left
	case 3:
		return utils.Up
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
			Raw:       str,
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
