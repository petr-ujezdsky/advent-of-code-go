package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Vector2i = utils.Vector2i

type Readout struct {
	Sensor, NearestBeacon Vector2i
}

func DoWithInput(items []Readout) int {
	return len(items)
}

func ParseInput(r io.Reader) []Readout {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var readouts []Readout
	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), true)

		item := Readout{
			Sensor:        Vector2i{ints[0], ints[1]},
			NearestBeacon: Vector2i{ints[2], ints[3]},
		}

		readouts = append(readouts, item)
	}

	return readouts
}
