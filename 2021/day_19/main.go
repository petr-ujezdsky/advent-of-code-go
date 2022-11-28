package day_19

import (
	"bufio"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Vector3i = utils.Vector3i

var regexHeader = regexp.MustCompile("--- scanner (\\d+) ---")

type BeaconScanner struct {
	Id             int
	RotatedBeacons [][]Vector3i
}

func roll(positions []Vector3i) []Vector3i {
	transformed := make([]Vector3i, len(positions))
	for i, v := range positions {
		transformed[i] = Vector3i{v.X, v.Z, -v.Y}
	}
	return transformed
}

func turn(positions []Vector3i) []Vector3i {
	transformed := make([]Vector3i, len(positions))
	for i, v := range positions {
		transformed[i] = Vector3i{-v.Y, v.X, v.Z}
	}
	return transformed
}

// allRotations returns all 24 variants of rotated vectors. The original vectors are in first element
// see https://stackoverflow.com/a/16467849/1310733
func allRotations(positions []Vector3i) [][]Vector3i {
	length := 24
	rotations := make([][]Vector3i, length)
	i := 0
	for cycle := 0; cycle < 2; cycle++ {
		// 3x RTTT
		for step := 0; step < 3; step++ {
			// R
			positions = roll(positions)
			rotations[length-i-1] = positions
			i++

			// 3x T
			for t := 0; t < 3; t++ {
				positions = turn(positions)
				rotations[length-i-1] = positions
				i++
			}
		}
	}

	return rotations
}

func ParseInput(r io.Reader) ([]BeaconScanner, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var beaconScanners []BeaconScanner
	for scanner.Scan() {
		// header with ID
		id, err := strconv.Atoi(regexHeader.FindStringSubmatch(scanner.Text())[1])
		if err != nil {
			return nil, err
		}

		// beacon positions
		var beacons []Vector3i
		for scanner.Scan() && scanner.Text() != "" {
			coords := strings.Split(scanner.Text(), ",")

			x, err := strconv.Atoi(coords[0])
			if err != nil {
				return nil, err
			}

			y, err := strconv.Atoi(coords[1])
			if err != nil {
				return nil, err
			}

			z, err := strconv.Atoi(coords[2])
			if err != nil {
				return nil, err
			}

			beacons = append(beacons, Vector3i{x, y, z})
		}

		beaconScanners = append(beaconScanners, BeaconScanner{id, allRotations(beacons)})
	}

	return beaconScanners, nil
}
