package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"io"
)

type Vector3i = utils.Vector3i
type Cube = Vector3i
type Cubes map[Cube]*int

func SurfaceArea(cubes Cubes) int {
	totalFaceCount := 6 * len(cubes)

	for len(cubes) > 0 {
		// take 1 cube
		cube, faceCount := maps.FirstEntry(cubes)
		delete(cubes, cube)

		// inspect neighbours
		for _, direction := range utils.Direction3D6Arr {
			neighbour := cube.Add(direction)

			// neighbour exists
			if neighbourFaceCount, ok := cubes[neighbour]; ok {
				*neighbourFaceCount--
				*faceCount--
				totalFaceCount -= 2
			}
		}
	}

	return totalFaceCount
}

func ParseInput(r io.Reader) Cubes {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	cubes := make(Cubes)

	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)

		cube := Cube{X: ints[0], Y: ints[1], Z: ints[2]}

		fc := 6
		cubes[cube] = &fc
	}

	return cubes
}
