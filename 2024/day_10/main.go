package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type VecSet map[utils.Vector2i]struct{}

type World struct {
	Matrix     matrix.Matrix[int]
	TrailHeads []utils.Vector2i
}

func walk(pos, from utils.Vector2i, height int, m matrix.Matrix[int], peaks VecSet) {
	if height == 9 {
		peaks[pos] = struct{}{}
		return
	}

	for _, step := range utils.Direction4Steps {
		nextPos := pos.Add(step)
		if nextPos == from {
			continue
		}

		nextHeight, ok := m.GetVSafe(nextPos)
		if !ok {
			continue
		}

		if nextHeight != height+1 {
			continue
		}

		walk(nextPos, pos, nextHeight, m, peaks)
	}
}

func DoWithInputPart01(world World) int {
	score := 0

	for i, trailHead := range world.TrailHeads {
		fmt.Printf("Trying trailhead #%v: ", i)
		peaks := make(VecSet)
		walk(trailHead, utils.Vector2i{X: -1, Y: -1}, 0, world.Matrix, peaks)
		count := len(peaks)
		fmt.Printf("%vx\n", count)
		score += count
	}

	return score
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	var traileheads []utils.Vector2i

	parseItem := func(char rune, x, y int) int {
		if char == '0' {
			traileheads = append(traileheads, utils.Vector2i{X: x, Y: y})
		}
		return utils.ParseInt(string(char))
	}

	return World{
		Matrix:     parsers.ParseToMatrixIndexed(r, parseItem),
		TrailHeads: traileheads,
	}
}
