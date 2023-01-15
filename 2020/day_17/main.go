package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

type Cubes = map[utils.Vector4i]struct{}

type World struct {
	ActiveCubes Cubes
}

func inspectNeighbours(activeCube utils.Vector4i, activeCubes Cubes, inactiveNeighbours map[utils.Vector4i]int, directions []utils.Vector4i) int {
	activeNeighboursCount := 0

	for _, step := range directions {
		neighbour := activeCube.Add(step)

		if _, ok := activeCubes[neighbour]; ok {
			activeNeighboursCount++
		} else {
			inactiveNeighbours[neighbour]++
		}
	}

	return activeNeighboursCount
}

func round(activeCubes Cubes, directions []utils.Vector4i) Cubes {
	nextActiveCubes := make(Cubes)
	inactiveNeighbours := make(map[utils.Vector4i]int)

	for activeCube := range activeCubes {
		activeNeighboursCount := inspectNeighbours(activeCube, activeCubes, inactiveNeighbours, directions)

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

	// use only 3D directions from 4D space
	directions := slices.Filter(utils.Direction4D80Arr[:], func(dir utils.Vector4i) bool { return dir.W == 0 })

	for i := 0; i < 6; i++ {
		activeCubes = round(activeCubes, directions)
	}

	return len(activeCubes)
}

func DoWithInputPart02(world World) int {
	activeCubes := world.ActiveCubes

	for i := 0; i < 6; i++ {
		activeCubes = round(activeCubes, utils.Direction4D80Arr[:])
	}

	return len(activeCubes)
}

func ParseInput(r io.Reader) World {
	activeCubes := make(Cubes)

	parseCube := func(char rune, x, y int) int {
		if char == '#' {
			position := utils.Vector4i{
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
