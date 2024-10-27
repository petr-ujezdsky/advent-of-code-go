package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
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

func DoWithInputPart02(world World, width int, height int) string {
	layers := world.Numbers
	layerSize := width * height

	image := matrix.NewMatrix[byte](width, height)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for i := 0; i < len(layers); i += layerSize {
				index := y*width + x + i
				value := layers[index]

				if value == 2 {
					// transparent -> skip
					continue
				}

				// black or white -> store
				image.Set(x, y, value)
				break
			}
		}
	}

	str := image.StringFmtSeparator("", func(value byte) string {
		if value == 0 {
			return " "
		}

		return "#"
	})

	fmt.Println(str)

	return str
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
