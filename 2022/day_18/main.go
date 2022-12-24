package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Vector3i = utils.Vector3i
type Cube Vector3i
type Cubes map[Cube]struct{}

func SurfaceArea(cubes Cubes) int {
	return len(cubes)
}

func ParseInput(r io.Reader) Cubes {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	cubes := make(Cubes)

	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)

		cube := Cube{X: ints[0], Y: ints[1], Z: ints[2]}

		cubes[cube] = struct{}{}
	}

	return cubes
}
