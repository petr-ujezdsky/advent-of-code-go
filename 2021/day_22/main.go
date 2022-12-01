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

var regexCube = regexp.MustCompile("(on|off) x=(\\d+)\\.\\.(\\d+),y=(\\d+)\\.\\.(\\d+),z=(\\d+)\\.\\.(\\d+)")

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
