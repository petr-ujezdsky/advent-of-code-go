package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World = utils.Matrix[bool]

func DoWithInput(world World) int {
	step := utils.Vector2i{X: 3, Y: 1}
	pos := utils.Vector2i{}

	treesCount := 0
	for pos.Y < world.Height {
		if world.GetV(pos) {
			treesCount++
		}

		pos = pos.Add(step)
		pos.X = pos.X % world.Width
	}

	return treesCount
}

func ParseInput(r io.Reader) World {
	return parsers.ParseToMatrix(r, parsers.MapperBoolean('#', '.'))
}
