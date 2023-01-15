package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Cubes = map[utils.Vector3i]struct{}

type World struct {
	ActiveCubes Cubes
}

func inspectNeighbours(activeCube utils.Vector3i, activeCubes Cubes, inactiveNeighbours map[utils.Vector3i]int) int {
	activeNeighboursCount := 0

	for _, step := range utils.Direction3D26Arr {
		neighbour := activeCube.Add(step)

		if _, ok := activeCubes[neighbour]; ok {
			activeNeighboursCount++
		} else {
			inactiveNeighbours[neighbour]++
		}
	}

	return activeNeighboursCount
}

func round(activeCubes Cubes) Cubes {
	nextActiveCubes := make(Cubes)
	inactiveNeighbours := make(map[utils.Vector3i]int)

	for activeCube := range activeCubes {
		activeNeighboursCount := inspectNeighbours(activeCube, activeCubes, inactiveNeighbours)

		if activeNeighboursCount == 2 || activeNeighboursCount == 3 {
			// will remain active
			nextActiveCubes[activeCube] = struct{}{}
		}
	}

	for inactiveNeighbour, activesCount := range inactiveNeighbours {
		if activesCount == 3 {
			// inactive has 3 active neighbours -> activate
			nextActiveCubes[inactiveNeighbour] = struct{}{}
		}
	}

	return nextActiveCubes
}

func DoWithInputPart01(world World) int {
	activeCubes := world.ActiveCubes

	for i := 0; i < 6; i++ {
		activeCubes = round(activeCubes)
	}

	return len(activeCubes)
}

func DoWithInputPart02(world World) int {
	return len(world.ActiveCubes)
}

func ParseInput(r io.Reader) World {
	activeCubes := make(Cubes)

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
