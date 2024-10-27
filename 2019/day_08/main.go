package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"math"
)

type World struct {
	Numbers []byte
}

func computeHistogram(layer []byte) [10]int {
	histogram := [10]int{}

	for _, number := range layer {
		histogram[number]++
	}

	return histogram
}

func DoWithInputPart01(world World, width int, height int) int {
	layers := world.Numbers
	layerSize := width * height

	minZerosHistogram := [10]int{}
	minZerosHistogram[0] = math.MaxInt

	for index := 0; index < len(layers); index += layerSize {
		layer := layers[index : index+layerSize]
		histogram := computeHistogram(layer)

		if histogram[0] < minZerosHistogram[0] {
			minZerosHistogram = histogram
		}
	}

	return minZerosHistogram[1] * minZerosHistogram[2]
}

func DoWithInputPart02(world World, width int, height int) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(str string) []byte {
		numbers := make([]byte, len(str))
		for i, char := range str {
			numbers[i] = byte(char - '0')
		}
		return numbers
	}

	items := parsers.ParseToObjects(r, parseItem)
	return World{Numbers: items[0]}
}
