package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	ActiveCubes map[utils.Vector3i]struct{}
}

func DoWithInputPart01(world World) int {
	return len(world.ActiveCubes)
}

func DoWithInputPart02(world World) int {
	return len(world.ActiveCubes)
}

func ParseInput(r io.Reader) World {
	activeCubes := make(map[utils.Vector3i]struct{})

	parseCube := func(char rune, x, y int) int {
		if char == '#' {
			position := utils.Vector3i{
				X: x,
				Y: y,
				Z: 0,
			}
			activeCubes[position] = struct{}{}
		}
		return 0
	}

	parsers.ParseToMatrixIndexed(r, parseCube)

	return World{ActiveCubes: activeCubes}
}
