package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/equations"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type ClawMachine struct {
	DirA, DirB    utils.Vector2i
	PrizePosition utils.Vector2i
}

type World struct {
	ClawMachines []ClawMachine
}

func quickWin(m ClawMachine) int {
	A := matrix.NewMatrixInt(2, 2)
	b := utils.NewVectorNn[int](2)

	A.Columns[0][0] = m.DirA.X
	A.Columns[1][0] = m.DirB.X
	b.Items[0] = m.PrizePosition.X

	A.Columns[0][1] = m.DirA.Y
	A.Columns[1][1] = m.DirB.Y
	b.Items[1] = m.PrizePosition.Y

	result, ok := equations.SolveLinearEquationsInt(A, b)
	if !ok {
		return 0
	}

	return result.Items[0]*3 + result.Items[1]*1
}

func DoWithInputPart01(world World) int {
	sum := 0

	for _, clawMachine := range world.ClawMachines {
		sum += quickWin(clawMachine)
	}

	return sum
}

func DoWithInputPart02(world World) int {
	shift := utils.Vector2i{X: 10000000000000, Y: 10000000000000}

	sum := 0

	for _, clawMachine := range world.ClawMachines {
		clawMachine.PrizePosition = clawMachine.PrizePosition.Add(shift)
		sum += quickWin(clawMachine)
	}

	return sum
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
