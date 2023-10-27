package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
	"strings"
)

type Wire struct {
	Steps  []utils.Vector2i
	Points []utils.Vector2i
	Lines  []utils.LineOrthogonal2i
}

type World struct {
	Wire1, Wire2 Wire
}

func DoWithInputPart01(world World) int {
	minDistance := math.MaxInt

	for _, line1 := range world.Wire1.Lines {
		for _, line2 := range world.Wire2.Lines {
			// check intersection
			if intersection, ok := line1.Intersection(line2); ok {
				// intersection is point
				if intersection.A == intersection.B {
					// skip origin
					if intersection.A == (utils.Vector2i{X: 0, Y: 0}) {
						continue
					}

					minDistance = utils.Min(minDistance, intersection.A.LengthManhattan())
				} else {
					panic("Intersection is line - not implemented")
				}
			}
		}
	}

	return minDistance
}

func DoWithInputPart02(world World) int {
	minDistance := math.MaxInt

	processedIntersections := make(map[utils.Vector2i]struct{})

	length1 := 0
	for _, line1 := range world.Wire1.Lines {
		length2 := 0
		for _, line2 := range world.Wire2.Lines {
			// check intersection
			if intersection, ok := line1.Intersection(line2); ok {
				// intersection is point
				if intersection.A == intersection.B {
					// skip origin
					if intersection.A == (utils.Vector2i{X: 0, Y: 0}) {
						continue
					}

					// skip already processed intersections
					if _, ok := processedIntersections[intersection.A]; ok {
						continue
					}

					//toIntersectionLength1 := line1.A.Subtract(intersection.A).LengthManhattan() + length1
					//toIntersectionLength2 := line2.A.Subtract(intersection.A).LengthManhattan() + length2

					toIntersectionLength1 := utils.NewLineOrthogonal2i(line1.A, intersection.A).Length() + length1 - 1
					toIntersectionLength2 := utils.NewLineOrthogonal2i(line2.A, intersection.A).Length() + length2 - 1

					minDistance = utils.Min(minDistance, toIntersectionLength1+toIntersectionLength2)

					processedIntersections[intersection.A] = struct{}{}
				} else {
					panic("Intersection is line - not implemented")
				}
			}
			length2 += line2.Length() - 1
		}
		length1 += line1.Length() - 1
	}

	return minDistance
}

func parseStep(step string) utils.Vector2i {
	// first char is direction, rest is amount
	dir := string(step[0])
	amount := utils.ParseInt(step[1:])

	var vec utils.Vector2i

	switch dir {
	case "R":
		vec = utils.Vector2i{X: amount, Y: 0}
	case "D":
		vec = utils.Vector2i{X: 0, Y: -amount}
	case "L":
		vec = utils.Vector2i{X: -amount, Y: 0}
	case "U":
		vec = utils.Vector2i{X: 0, Y: amount}
	default:
		panic("Unknown direction " + dir)
	}

	return vec
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var wires []Wire
	for scanner.Scan() {
		stepsStr := strings.Split(scanner.Text(), ",")

		steps := make([]utils.Vector2i, len(stepsStr))
		points := make([]utils.Vector2i, len(stepsStr)+1)
		lines := make([]utils.LineOrthogonal2i, len(stepsStr))

		points[0] = utils.Vector2i{X: 0, Y: 0}

		for i, stepStr := range stepsStr {
			step := parseStep(stepStr)

			steps[i] = step
			points[i+1] = points[i].Add(step)
			lines[i] = utils.NewLineOrthogonal2i(points[i], points[i+1])
		}

		wires = append(wires, Wire{
			Steps:  steps,
			Points: points,
			Lines:  lines})
	}

	return World{
		Wire1: wires[0],
		Wire2: wires[1],
	}
}
