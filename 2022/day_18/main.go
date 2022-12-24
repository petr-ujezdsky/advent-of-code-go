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
type World struct {
	Cubes       Cubes
	BoundingBox utils.BoundingBox
}

type Cubes map[Cube]struct{}

func SurfaceArea(world World) int {
	cubes := world.Cubes
	totalFaceCount := 6 * len(cubes)

	for len(cubes) > 0 {
		// take 1 cube
		cube, _ := maps.FirstEntry(cubes)
		delete(cubes, cube)

		// inspect neighbours
		for _, direction := range utils.Direction3D6Arr {
			neighbour := cube.Add(direction)

			// neighbour exists
			if _, ok := cubes[neighbour]; ok {
				totalFaceCount -= 2
			}
		}
	}

	return totalFaceCount
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	cubes := make(Cubes)
	var boundingBox utils.BoundingBox

	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)

		cube := Cube{X: ints[0], Y: ints[1], Z: ints[2]}

		if len(cubes) == 0 {
			boundingBox = utils.NewBoundingBox(cube)
		} else {
			boundingBox = boundingBox.Enlarge(cube)
		}

		cubes[cube] = struct{}{}
	}

	return World{
		Cubes:       cubes,
		BoundingBox: boundingBox,
	}
}
