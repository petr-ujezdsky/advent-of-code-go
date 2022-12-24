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

func OuterSurfaceArea(world World) int {
	// enlarge boundingBox by 1 in each direction
	boundingBox := world.BoundingBox
	boundingBox = boundingBox.Enlarge(utils.Vector3i{X: boundingBox.XInterval.Low - 1, Y: boundingBox.YInterval.Low - 1, Z: boundingBox.ZInterval.Low - 1})
	boundingBox = boundingBox.Enlarge(utils.Vector3i{X: boundingBox.XInterval.High + 1, Y: boundingBox.YInterval.High + 1, Z: boundingBox.ZInterval.High + 1})

	cubes := world.Cubes

	outerFaceCount := 0
	inspected := make(Cubes)
	toInspect := make(Cubes)
	// pick definitely outer cube and start search with it
	startProbe := utils.Vector3i{X: boundingBox.XInterval.Low, Y: boundingBox.YInterval.Low, Z: boundingBox.ZInterval.Low}
	toInspect[startProbe] = struct{}{}

	for len(toInspect) > 0 {
		probe, _ := maps.FirstEntry(toInspect)
		delete(toInspect, probe)

		// inspect probe neighbours
		for _, direction := range utils.Direction3D6Arr {
			neighbour := probe.Add(direction)

			// neighbour is outside the area
			if !boundingBox.Contains(neighbour) {
				continue
			}

			// already inspected -> skip
			if _, ok := inspected[neighbour]; ok {
				continue
			}

			if _, ok := cubes[neighbour]; ok {
				// neighbour is rock -> count 1 face
				outerFaceCount++
			} else {
				// continue search
				toInspect[neighbour] = struct{}{}
			}
		}

		inspected[probe] = struct{}{}
	}

	return outerFaceCount
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
