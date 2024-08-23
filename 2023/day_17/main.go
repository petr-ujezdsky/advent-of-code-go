package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"strconv"
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
	Position         utils.Vector2i
	Direction        utils.Vector2i
	StepsInDirection int
}

func DoWithInputPart01(world World) int {
	tiles := world.Tiles
	path, _, totalHeatLoss, ok := FindMinHeatLossPath(tiles)
	if !ok {
		panic("No path found")
	}

	pathMap := slices.ToMap(path, func(s State) utils.Vector2i { return s.Position })

	dirStr := make(map[utils.Vector2i]string)
	dirStr[dirs[0]] = "<"
	dirStr[dirs[1]] = "^"
	dirStr[dirs[2]] = ">"
	dirStr[dirs[3]] = "v"

	str := matrix.StringFmtSeparatorIndexed[int](tiles, "", func(value int, x, y int) string {
		if s, ok := pathMap[utils.Vector2i{X: x, Y: y}]; ok {
			if char, ok := dirStr[s.Direction]; ok {
				return char
			}
			return "."
		}

		return strconv.Itoa(value)
	})

	fmt.Printf("Tiles:\n%v\n", str)

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
			if _, ok := m.GetVSafe(nextPos); !ok {
				continue
			}

			steps := 0
			if origin.Direction == dir {
				steps = origin.StepsInDirection + 1
			}

			if steps > 2 {
				continue
			}

			nextState := State{Position: nextPos, Direction: dir, StepsInDirection: steps}
			neighbours = append(neighbours, nextState)
		}

		return neighbours
	}
}

func isEnd(endPos utils.Vector2i) func(state State) bool {
	return func(state State) bool {
		return state.Position == endPos
	}
}

func FindMinHeatLossPath(m Matrix2i) ([]State, map[State]int, int, bool) {
	endPos := utils.Vector2i{X: m.Width - 1, Y: m.Height - 1}
	return alg.AStarEndFunc(State{}, isEnd(endPos), h(endPos), d(m), neighbours(m))
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	return World{Tiles: parsers.ParseToMatrixInt(r)}
}
