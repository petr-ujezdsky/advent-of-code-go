package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Vector2i = utils.Vector2i
type IntervalI = utils.IntervalI

type Scanner struct {
	Readouts []Readout
	Beacons  []Vector2i
}

type Readout struct {
	Sensor, NearestBeacon Vector2i
	Distance              int
}

func beaconsAtY(beacons []Vector2i, y int) []Vector2i {
	var beaconsFiltered []Vector2i

	for _, beacon := range beacons {
		if beacon.Y == y {
			beaconsFiltered = append(beaconsFiltered, beacon)
		}
	}

	return beaconsFiltered
}

func unionSizeWithoutBeacons(union []IntervalI, beacons []Vector2i, y int) int {
	// filter beacons to given y
	beacons = beaconsAtY(beacons, y)

	// calculate union size
	totalSize := 0
	for _, i := range union {
		size := i.Size()
		for _, beacon := range beacons {
			if i.Contains(beacon.X) {
				size--
			}
		}
		if size > 0 {
			totalSize += size
		}
	}

	return totalSize
}

func noBeaconPositions(scanner Scanner, y int) []IntervalI {
	readouts := scanner.Readouts

	var intervals []utils.IntervalI
	for _, readout := range readouts {
		sensor := readout.Sensor

		// check intersection possibility
		yDistance := utils.Abs(sensor.Y - y)
		if yDistance <= readout.Distance {
			xHalf := readout.Distance - yDistance
			from := sensor.X - xHalf
			to := sensor.X + xHalf

			span := IntervalI{from, to}
			intervals = append(intervals, span)
		}
	}

	// find intervals union
	return utils.Union(intervals)
}

func NoBeaconPositionsCount(scanner Scanner, y int) int {
	union := noBeaconPositions(scanner, y)

	return unionSizeWithoutBeacons(union, scanner.Beacons, y)
}

func BeaconPositionFrequency(scanner Scanner, yMax int) int {
	for y := 0; y <= yMax; y++ {
		union := noBeaconPositions(scanner, y)
		if len(union) > 1 {
			// found a spot

			x := union[0].High + 1

			return 4_000_000*x + y
		}
	}

	panic("Found nothing")
}

func ParseInput(r io.Reader) Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var readouts []Readout
	beaconsMap := make(map[Vector2i]struct{})

	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), true)

		sensor := Vector2i{ints[0], ints[1]}
		nearestBeacon := Vector2i{ints[2], ints[3]}

		item := Readout{
			Sensor:        sensor,
			NearestBeacon: nearestBeacon,
			Distance:      sensor.Subtract(nearestBeacon).LengthManhattan(),
		}

		readouts = append(readouts, item)
		beaconsMap[nearestBeacon] = struct{}{}
	}

	beacons := make([]Vector2i, len(beaconsMap))
	for beacon := range beaconsMap {
		beacons = append(beacons, beacon)
	}

	return Scanner{readouts, beacons}
}
