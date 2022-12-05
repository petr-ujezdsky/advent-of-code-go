package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"strings"
)

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
