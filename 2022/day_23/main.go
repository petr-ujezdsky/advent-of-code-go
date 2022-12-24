package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
)

type Vector2i = utils.Vector2i

type Elf Vector2i

type World map[Elf]struct{}

type Proposition struct {
	DirectionsToCheck [3]Vector2i
	Direction         utils.Direction4
}

var propositions = [4]Proposition{
	{
		DirectionsToCheck: [3]Vector2i{utils.North.ToStep(), utils.NorthEast.ToStep(), utils.NorthWest.ToStep()},
		Direction:         utils.Up,
	},
	{
		DirectionsToCheck: [3]Vector2i{utils.South.ToStep(), utils.SouthEast.ToStep(), utils.SouthWest.ToStep()},
		Direction:         utils.Down,
	},
	{
		DirectionsToCheck: [3]Vector2i{utils.West.ToStep(), utils.NorthWest.ToStep(), utils.SouthWest.ToStep()},
		Direction:         utils.Left,
	},
	{
		DirectionsToCheck: [3]Vector2i{utils.East.ToStep(), utils.NorthEast.ToStep(), utils.SouthEast.ToStep()},
		Direction:         utils.Right,
	},
}

func DoWithInput(elves World) int {
	iProposition := 0
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	elves := make(World)

	y := 0
	for scanner.Scan() {
		for x, char := range scanner.Text() {
			if char == '.' {
				continue
			}

			elf := Elf{X: x, Y: y}
			elves[elf] = struct{}{}
		}
		y--
	}

	return elves
}
