package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Antenna struct {
	Position utils.Vector2i
	Name     string
}

type World struct {
	Antennae []Antenna
	Matrix   matrix.Matrix[rune]
}

func DoWithInputPart01(world World) int {
	bounds := world.Matrix.Bounds()

	antinodes := make(map[utils.Vector2i]struct{})

	for i, antenna1 := range world.Antennae {
		for _, antenna2 := range world.Antennae[i+1:] {
			if antenna1.Name != antenna2.Name {
				continue
			}

			step := antenna2.Position.Subtract(antenna1.Position)

			antinode1 := antenna1.Position.Subtract(step)
			if bounds.Contains(antinode1) {
				antinodes[antinode1] = struct{}{}
			}

			antinode2 := antenna1.Position.Add(step.Multiply(2))
			if bounds.Contains(antinode2) {
				antinodes[antinode2] = struct{}{}
			}
		}
	}

	return len(antinodes)
}

func DoWithInputPart02(world World) int {
	bounds := world.Matrix.Bounds()

	antinodes := make(map[utils.Vector2i]struct{})

	for i, antenna1 := range world.Antennae {
		for _, antenna2 := range world.Antennae[i+1:] {
			if antenna1.Name != antenna2.Name {
				continue
			}

			antinodes[antenna1.Position] = struct{}{}
			antinodes[antenna2.Position] = struct{}{}

			step := antenna2.Position.Subtract(antenna1.Position)

			antinode1 := antenna1.Position.Subtract(step)
			for bounds.Contains(antinode1) {
				antinodes[antinode1] = struct{}{}
				antinode1 = antinode1.Subtract(step)
			}

			antinode2 := antenna2.Position.Add(step)
			for bounds.Contains(antinode2) {
				antinodes[antinode2] = struct{}{}
				antinode2 = antinode2.Add(step)
			}
		}
	}

	formatter := func(char rune, x, y int) string {
		pos := utils.Vector2i{X: x, Y: y}
		_, antinode := antinodes[pos]
		antenna := char != '.'

		if antinode && antenna {
			return "+"
		}

		if antinode {
			return "#"
		}

		return string(char)
	}

	str := matrix.StringFmtSeparatorIndexed(world.Matrix, true, "", formatter)
	fmt.Println(str)

	return len(antinodes)
}

func ParseInput(r io.Reader) World {
	var antennae []Antenna

	parseItem := func(char rune, x, y int) rune {
		if char != '.' {
			antennae = append(antennae, Antenna{
				Position: utils.Vector2i{X: x, Y: y},
				Name:     string(char),
			})
		}

		return char
	}

	return World{
		Antennae: antennae,
		Matrix:   parsers.ParseToMatrixIndexed(r, parseItem),
	}
}
