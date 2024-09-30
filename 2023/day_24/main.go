package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"strings"
)

type Line struct {
	Position   utils.Vector3i
	Direction  utils.Vector3i
	Position2  utils.Vector2i
	Direction2 utils.Vector2i
}

type World struct {
	Lines []Line
}

func DoWithInputPart01(world World) int {
	return 0
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) Line {
		parts := strings.Split(str, " @ ")
		position := utils.ToIntsP(trim(strings.Split(parts[0], ", ")))
		direction := utils.ToIntsP(trim(strings.Split(parts[1], ", ")))

		return Line{
			Position: utils.Vector3i{
				X: position[0],
				Y: position[1],
				Z: position[2],
			},
			Direction: utils.Vector3i{
				X: direction[0],
				Y: direction[1],
				Z: direction[2],
			},
			Position2: utils.Vector2i{
				X: position[0],
				Y: position[1],
			},
			Direction2: utils.Vector2i{
				X: direction[0],
				Y: direction[1],
			},
		}
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Lines: items}
}

func trim(rows []string) []string {
	trimmed := make([]string, len(rows))

	for i, row := range rows {
		trimmed[i] = strings.TrimSpace(row)
	}

	return trimmed
}
