package day_22

import (
	"bufio"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

type IntervalI = utils.IntervalI

type Cube struct {
	X, Y, Z IntervalI
	On      bool
}

func (c Cube) Volume() int {
	return c.X.Size() * c.Y.Size() * c.Z.Size()
}

func (c Cube) Subtract(c2 Cube) []Cube {
	ix, ok := c.X.Intersection(c2.X)
	if !ok {
		return []Cube{c}
	}

	iy, ok := c.Y.Intersection(c2.Y)
	if !ok {
		return []Cube{c}
	}

	iz, ok := c.Z.Intersection(c2.Z)
	if !ok {
		return []Cube{c}
	}

	if c.X == ix && c.Y == iy && c.Z == iz {
		// whole c is inside c2 -> result is empty
		// the algorithm below will work, this is just a shortcut
		return nil
	}

	xs := []IntervalI{{c.X.Low, ix.Low - 1}, ix, {ix.High + 1, c.X.High}}
	ys := []IntervalI{{c.Y.Low, iy.Low - 1}, iy, {iy.High + 1, c.Y.High}}
	zs := []IntervalI{{c.Z.Low, iz.Low - 1}, iz, {iz.High + 1, c.Z.High}}

	// generate all 27 sub-cubes
	subs := make([]Cube, 0, 27)
	for xj, xi := range xs {
		if xi.IsInversed() {
			// invalid interval - intersection is at the beginning (or end) of the cube
			continue
		}
		for yj, yi := range ys {
			if yi.IsInversed() {
				// invalid interval - intersection is at the beginning (or end) of the cube
				continue
			}
			for zj, zi := range zs {
				// intersection cube is not part of the output
				if xj == 1 && yj == 1 && zj == 1 {
					continue
				}

				if zi.IsInversed() {
					// invalid interval - intersection is at the beginning (or end) of the cube
					continue
				}

				sub := Cube{
					X:  xi,
					Y:  yi,
					Z:  zi,
					On: c.On,
				}

				subs = append(subs, sub)
			}
		}
	}

	return subs
}

func CountOnCubes(cubes []Cube) int {
	var onCubes []Cube
	for i, cube := range cubes {
		if cube.On {
			subs := []Cube{cube}

			// subtract all current on-cubes from cube
			for _, onCube := range onCubes {
				var newSubs []Cube
				for _, sub := range subs {
					newSubs = append(newSubs, sub.Subtract(onCube)...)
				}
				subs = newSubs
			}
			onCubes = append(onCubes, subs...)
		} else {
			var newOnCubes []Cube

			// subtract cube from all current on-cubes
			for _, onCube := range onCubes {
				newOnCubes = append(newOnCubes, onCube.Subtract(cube)...)
			}
			onCubes = newOnCubes
		}
		fmt.Printf("Cube #%v (%v), onCubes count: %v\n", i+1, cube.On, len(onCubes))
	}

	count := 0
	for _, onCube := range onCubes {
		count += onCube.Volume()
	}

	return count
}

func ParseInput(r io.Reader) []Cube {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var cubes []Cube
	for scanner.Scan() {
		line := scanner.Text()
		ints := utils.ExtractInts(line, true)

		on := strings.HasPrefix(line, "on")

		cube := Cube{
			X:  utils.NewInterval(ints[0], ints[1]),
			Y:  utils.NewInterval(ints[2], ints[3]),
			Z:  utils.NewInterval(ints[4], ints[5]),
			On: on,
		}

		cubes = append(cubes, cube)
	}

	return cubes
}
