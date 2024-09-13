package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strings"
)

type Cube struct {
	Box utils.BoundingBox
}

type World struct {
	Cubes []*Cube
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) *Cube {
		points := strings.Split(str, "~")

		pointA := parsePoint(points[0])
		pointB := parsePoint(points[1])

		return &Cube{Box: utils.NewBoundingBoxPoints(pointA, pointB)}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Cubes: items}
}

func parsePoint(str string) utils.Vector3i {
	coordinates := strings.Split(str, ",")

	return utils.Vector3i{
		X: utils.ParseInt(coordinates[0]),
		Y: utils.ParseInt(coordinates[1]),
		Z: utils.ParseInt(coordinates[2]),
	}
}
