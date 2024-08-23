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

var dirs = []utils.Vector2i{
	// left
	{-1, 0},
	// up
	{0, -1},
	// right
	{1, 0},
	// down
	{0, 1},
}

type State struct {
	Position utils.Vector2i
}

func DoWithInputPart01(world World) int {
	tiles := world.Tiles
	_, _, totalHeatLoss, ok := FindMinHeatLossPath(tiles)
	if !ok {
		panic("No path found")
	}

	return totalHeatLoss
}

func h(endPos utils.Vector2i) func(state State) int {
	return func(state State) int {
		// manhattan distance
		return utils.Abs(state.Position.X-endPos.X) + utils.Abs(state.Position.Y-endPos.Y)
	}
}

func d(m Matrix2i) func(State, State) int {
	return func(nodeFrom, nodeTo State) int {
		// step heat loss is the heat loss of target node
		return m.GetV(nodeTo.Position)
	}
}

func neighbours(m Matrix2i) func(origin State) []State {
	return func(origin State) []State {
		var neighbours []State
		for _, dir := range dirs {
			nextPos := origin.Position.Add(dir)

			// check validity
			if _, ok := m.GetVSafe(nextPos); ok {
				nextState := State{nextPos}
				neighbours = append(neighbours, nextState)
			}
		}

		return neighbours
	}
}

func FindMinHeatLossPath(m Matrix2i) ([]State, map[State]int, int, bool) {
	endPos := utils.Vector2i{X: m.Width - 1, Y: m.Height - 1}
	return alg.AStar(State{}, State{endPos}, h(endPos), d(m), neighbours(m))
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	return World{Tiles: parsers.ParseToMatrixInt(r)}
}
