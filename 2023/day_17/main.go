package main

import (
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type Matrix2i = matrix.MatrixInt

type World struct {
	Tiles Matrix2i
}

func DoWithInputPart01(world World) int {
	tiles := world.Tiles
	_, _, totalHeatLoss, ok := FindMinHeatLossPath(tiles)
	if !ok {
		panic("No path found")
	}

	return totalHeatLoss
}

func h(endPos utils.Vector2i) func(utils.Vector2i) int {
	return func(pos utils.Vector2i) int {
		// manhattan distance
		return utils.Abs(pos.X-endPos.X) + utils.Abs(pos.Y-endPos.Y)
	}
}

func d(m Matrix2i) func(utils.Vector2i, utils.Vector2i) int {
	return func(nodeFrom, nodeTo utils.Vector2i) int {
		// step heat loss is the heat loss of target node
		return m.GetV(nodeTo)
	}
}

func neighbours(m Matrix2i) func(origin utils.Vector2i) []utils.Vector2i {
	return func(origin utils.Vector2i) []utils.Vector2i {
		var neighbours []utils.Vector2i
		for _, dir := range dirs {
			nextPos := origin.Add(dir)

			// check validity
			if _, ok := m.GetVSafe(nextPos); ok {
				neighbours = append(neighbours, nextPos)
			}
		}

		return neighbours
	}
}

func FindMinHeatLossPath(m Matrix2i) ([]utils.Vector2i, map[utils.Vector2i]int, int, bool) {
	endPos := utils.Vector2i{m.Width - 1, 0}
	return alg.AStar(utils.Vector2i{Y: m.Height - 1}, endPos, h(endPos), d(m), neighbours(m))
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	return World{Tiles: parsers.ParseToMatrixInt(r)}
}
