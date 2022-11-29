package day_19

import (
	"bufio"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Vector3i = utils.Vector3i

var regexHeader = regexp.MustCompile("--- scanner (\\d+) ---")

type BeaconScanner struct {
	Id             int
	RotatedBeacons [][]Vector3i
	UniqueBeacons  map[Vector3i]struct{}
}

func NewBeaconScanner(id int, beacons []Vector3i) BeaconScanner {
	// find all rotations
	rotations := allRotations(beacons)

	// add unique beacons
	uniqueBeacons := make(map[Vector3i]struct{})
	for _, beacon := range rotations[0] {
		uniqueBeacons[beacon] = struct{}{}
	}

	return BeaconScanner{id, rotations, uniqueBeacons}
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
	rotations := make([][]Vector3i, 24)
	i := 0
	for cycle := 0; cycle < 2; cycle++ {
		// 3x RTTT
		for step := 0; step < 3; step++ {
			// R
			positions = roll(positions)
			rotations[i] = positions
			i++

			// 3x T
			for t := 0; t < 3; t++ {
				positions = turn(positions)
				rotations[i] = positions
				i++
			}
		}

		// RTR for next cycle
		positions = roll(turn(roll(positions)))
	}

	// original positions as first "rotation"
	rotations[0], rotations[11] = rotations[11], rotations[0]

	return rotations
}

func FindOverlap(s1unique map[Vector3i]struct{}, s2 []Vector3i) (bool, Vector3i) {
	// choose b1
	for b1 := range s1unique {
		// choose b2
		for _, b2 := range s2 {
			count := 0

			// try to align all s2 beacons using vec b2 -> b1
			step := b1.Subtract(b2)

			// find overlapped beacons
			for _, bb2 := range s2 {
				// align using step
				bb2 = bb2.Add(step)

				// check overlap
				if _, ok := s1unique[bb2]; ok {
					count++

					if count >= 12 {
						return true, step
					}
				}
			}
		}
	}

	return false, Vector3i{}
}

func FindOverlapRotations(s1, s2 BeaconScanner) (bool, int, Vector3i) {
	// over all s2 rotations
	for irot, s2beacons := range s2.RotatedBeacons {
		overlap, step := FindOverlap(s1.UniqueBeacons, s2beacons)
		if overlap {
			return true, irot, step
		}
	}

	return false, 0, Vector3i{}
}

func consume(mainScanner BeaconScanner, beacons []Vector3i, step Vector3i) BeaconScanner {
	allBeacons := mainScanner.RotatedBeacons[0]

	for _, beacon := range beacons {
		beacon = beacon.Add(step)

		// beacon is new
		if _, ok := mainScanner.UniqueBeacons[beacon]; !ok {
			allBeacons = append(allBeacons, beacon)
		}
	}

	return NewBeaconScanner(mainScanner.Id, allBeacons)
}

func largestManhattanDistance(origins []Vector3i) int {
	max := math.MinInt

	for o1i := 0; o1i < len(origins); o1i++ {
		for o2i := o1i + 1; o2i < len(origins); o2i++ {
			o1 := origins[o1i]
			o2 := origins[o2i]

			max = utils.Max(max, o2.Subtract(o1).ManhattanLength())
		}
	}

	return max
}

func SearchAndConsume(scanners []BeaconScanner) (int, int) {
	// all scanners will be rotated, translated and merged into this one
	mainScanner := scanners[0]
	scanners = utils.RemoveUnordered(scanners, 0)

	// all scanner origins, start with 0,0,0 for the main scanner (#0)
	scannerOrigins := []Vector3i{{0, 0, 0}}

	r := 0
	for len(scanners) > 0 {
		fmt.Printf("Round #%d\n", r)

		for i, scanner := range scanners {
			fmt.Printf("Finding overlap in %d / %d\n", i, len(scanners))

			ok, irot, step := FindOverlapRotations(mainScanner, scanner)
			if ok {
				// consume scanner
				mainScanner = consume(mainScanner, scanner.RotatedBeacons[irot], step)

				// remove scanner
				scanners = utils.RemoveUnordered(scanners, i)

				// save origin
				scannerOrigins = append(scannerOrigins, step)

				fmt.Printf("Found! Consuming %d\n", i)
				break
			}
		}
		r++
	}

	return len(mainScanner.UniqueBeacons), largestManhattanDistance(scannerOrigins)
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

		// find all rotations
		rotations := allRotations(beacons)

		// add unique beacons
		uniqueBeacons := make(map[Vector3i]struct{})
		for _, beacon := range rotations[0] {
			uniqueBeacons[beacon] = struct{}{}
		}

		beaconScanners = append(beaconScanners, NewBeaconScanner(id, beacons))
	}

	return beaconScanners, nil
}
