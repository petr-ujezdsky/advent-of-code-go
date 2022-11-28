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
	Id      int
	Beacons []Vector3i
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

		beaconScanners = append(beaconScanners, BeaconScanner{id, beacons})
	}

	return beaconScanners, nil
}
