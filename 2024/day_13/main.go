package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
	"math"
)

type ClawMachine struct {
	DirA, DirB    utils.Vector2i
	PrizePosition utils.Vector2i
}

type World struct {
	ClawMachines []ClawMachine
}

func bruteForceWin(m ClawMachine) int {
	cost := math.MaxInt

	for a := 0; a < 100; a++ {
		posA := m.DirA.Multiply(a)
		for b := 0; b < 100; b++ {
			posB := m.DirB.Multiply(b)

			pos := posA.Add(posB)

			if pos == m.PrizePosition {
				cost = utils.Min(cost, a*3+b*1)
			}
		}
	}

	if cost != math.MaxInt {
		return cost
	}

	return 0
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, clawMachine := range world.ClawMachines {
		sum += bruteForceWin(clawMachine)
	}

	return sum
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(lines []string, i int) ClawMachine {
		ints := utils.ExtractInts(lines[0], true)
		dirA := utils.Vector2i{X: ints[0], Y: ints[1]}

		ints = utils.ExtractInts(lines[1], true)
		dirB := utils.Vector2i{X: ints[0], Y: ints[1]}

		ints = utils.ExtractInts(lines[2], true)
		prizePosition := utils.Vector2i{X: ints[0], Y: ints[1]}

		return ClawMachine{
			DirA:          dirA,
			DirB:          dirB,
			PrizePosition: prizePosition,
		}
	}

	items := parsers.ParseToGroups(r, parseItem)
	return World{ClawMachines: items}
}
