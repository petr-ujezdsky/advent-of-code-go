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

			// do not reverse
			if dir == origin.Direction.Multiply(-1) {
				continue
			}

			if nextState, ok := dirStep(origin, dir, m, 1, 1, 3); ok {
				neighbours = append(neighbours, nextState)
			}
		}

		return neighbours
	}
}

func d2(m Matrix2i) func(State, State) int {
	return func(nodeFrom, nodeTo State) int {
		//dir := nodeTo.Position.Subtract(nodeFrom.Position).Signum()
		revDir := nodeFrom.Position.Subtract(nodeTo.Position).Signum()

		heatLoss := 0
		for pos := nodeTo.Position; pos != nodeFrom.Position; pos = pos.Add(revDir) {
			heatLoss += m.GetV(pos)
		}

		return heatLoss
	}
}

func neighbours2(m Matrix2i) func(origin State) []State {
	return func(origin State) []State {
		var neighbours []State
		for _, dir := range dirs {

			// do not reverse
			if dir == origin.Direction.Multiply(-1) {
				continue
			}

			// step 1
			if nextState, ok := dirStep(origin, dir, m, 1, 4, 10); ok {
				neighbours = append(neighbours, nextState)
			}

			// step 4
			if nextState, ok := dirStep(origin, dir, m, 4, 4, 10); ok {
				neighbours = append(neighbours, nextState)
			}
		}

		return neighbours
	}
}

func dirStep(origin State, dir utils.Vector2i, m Matrix2i, stepSize, minStraight, maxStraight int) (State, bool) {
	step := dir.Multiply(stepSize)
	nextPos := origin.Position.Add(step)

	// check validity
	if _, ok := m.GetVSafe(nextPos); !ok {
		return State{}, false
	}

	steps := stepSize
	if origin.Direction == dir {
		steps = origin.StepsInDirection + stepSize
	}

	if steps > maxStraight {
		return State{}, false
	}

	if steps < minStraight {
		return State{}, false
	}

	return State{Position: nextPos, Direction: dir, StepsInDirection: steps}, true
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

func FindMinHeatLossPath2(m Matrix2i) ([]State, map[State]int, int, bool) {
	endPos := utils.Vector2i{X: m.Width - 1, Y: m.Height - 1}
	return alg.AStarEndFunc(State{}, isEnd(endPos), h(endPos), d2(m), neighbours2(m))
}

func DoWithInputPart02(world World) int {
	tiles := world.Tiles
	path, _, totalHeatLoss, ok := FindMinHeatLossPath2(tiles)
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

func ParseInput(r io.Reader) World {
	return World{Tiles: parsers.ParseToMatrixInt(r)}
}
