package day_22

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"regexp"
)

type Vector3i = utils.Vector3i

type Cube struct {
	Low, High Vector3i
	Value     bool
}

func NewCubeSymmetric(halfSideLength int, value bool) Cube {
	return Cube{
		Low:   Vector3i{-halfSideLength, -halfSideLength, -halfSideLength},
		High:  Vector3i{halfSideLength, halfSideLength, halfSideLength},
		Value: value,
	}
}

func (c Cube) Contains(p Vector3i) bool {
	return c.Low.X <= p.X && p.X <= c.High.X &&
		c.Low.Y <= p.Y && p.Y <= c.High.Y &&
		c.Low.Z <= p.Z && p.Z <= c.High.Z
}

var regexCube = regexp.MustCompile("(on|off) x=(-?\\d+)\\.\\.(-?\\d+),y=(-?\\d+)\\.\\.(-?\\d+),z=(-?\\d+)\\.\\.(-?\\d+)")

func resolveOnOff(point Vector3i, cubes []Cube) bool {
	for _, cube := range cubes {
		if cube.Contains(point) {
			return cube.Value
		}
	}

	panic("Don't know if on or off!")
}

func NaiveCount(world Cube, cubes []Cube) int {
	count := 0

	// start investigation with *last* added cube and so on
	cubes = utils.Reverse(cubes)

	// final cube is world itself
	cubes = append(cubes, world)

	for x := world.Low.X; x <= world.High.X; x++ {
		for y := world.Low.Y; y <= world.High.Y; y++ {
			for z := world.Low.Z; z <= world.High.Z; z++ {
				if resolveOnOff(Vector3i{x, y, z}, cubes) {
					count++
				}
			}
		}
	}

	return count
}

func ParseInput(r io.Reader) []Cube {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var cubes []Cube
	for scanner.Scan() {
		parts := regexCube.FindStringSubmatch(scanner.Text())
		x1, x2 := utils.ParseInt(parts[2]), utils.ParseInt(parts[3])
		y1, y2 := utils.ParseInt(parts[4]), utils.ParseInt(parts[5])
		z1, z2 := utils.ParseInt(parts[6]), utils.ParseInt(parts[7])

		cube := Cube{
			Low:   Vector3i{x1, y1, z1},
			High:  Vector3i{x2, y2, z2},
			Value: parts[1] == "on",
		}

		cubes = append(cubes, cube)
	}

	return cubes
}
