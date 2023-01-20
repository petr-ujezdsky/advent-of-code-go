package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World = matrix.Matrix[bool]

func treesCountForStep(world World, step utils.Vector2i) int {
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

func TreesCount(world World) int {
	step := utils.Vector2i{X: 3, Y: 1}
	return treesCountForStep(world, step)
}

func TreesCountManySteps(world World) int {
	steps := []utils.Vector2i{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}
	result := 1

	for _, step := range steps {
		result *= treesCountForStep(world, step)
	}

	return result
}

func ParseInput(r io.Reader) World {
	return parsers.ParseToMatrix(r, parsers.MapperBoolean('#', '.'))
}
