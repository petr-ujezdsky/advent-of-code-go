package main

import (
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/iterators"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"io"
)

type World struct {
	Matrix     matrix.Matrix[string]
	Start, End utils.Vector2i
}

type State struct {
	Position  utils.Vector2i
	Direction utils.Vector2i
}

func h(endPos utils.Vector2i) func(state State) int {
	return func(state State) int {
		return utils.ManhattanDistance(state.Position, endPos)
	}
}

func d() func(State, State) int {
	return func(nodeFrom, nodeTo State) int {
		if nodeFrom.Direction == nodeTo.Direction {
			return 1
		}

		return 1000
	}
}

func n(m matrix.Matrix[string]) func(origin State, path iterators.Iterator[State]) []State {
	return func(origin State, path iterators.Iterator[State]) []State {
		var neighbours []State
		for _, dir := range utils.Direction4Steps {

			// do not reverse
			if dir == origin.Direction.Multiply(-1) {
				continue
			}

			if dir == origin.Direction {
				// move forward
				nextPos := origin.Position.Add(dir)

				if nn, ok := m.GetVSafe(nextPos); !ok || nn == "#" {
					continue
				}

				nextState := State{
					Position:  nextPos,
					Direction: dir,
				}

				neighbours = append(neighbours, nextState)
			} else {
				// rotate
				nextState := State{
					Position:  origin.Position,
					Direction: dir,
				}

				neighbours = append(neighbours, nextState)
			}
		}

		return neighbours
	}
}

func isEnd(endPos utils.Vector2i) func(state State) bool {
	return func(state State) bool {
		return state.Position == endPos
	}
}

func DoWithInputPart01(world World) int {
	start := State{
		Position:  world.Start,
		Direction: utils.Right.ToStep(),
	}

	path, _, score, _ := alg.AStarEndFunc(start, isEnd(world.End), h(world.End), d(), n(world.Matrix))

	for _, state := range path {
		world.Matrix.SetV(state.Position, "x")
	}
	formatter := matrix.NonIndexedAdapter(matrix.FmtFmt[string]("%s"))
	str := matrix.StringFmtSeparatorIndexed(world.Matrix, true, "", formatter)
	fmt.Println(str)

	return score
}

func connectBothDirections(forward, backward map[State]int, shortestPaths map[utils.Vector2i]struct{}, targetScore int) {
	for state, scoreFromStart := range forward {
		state.Direction = state.Direction.Multiply(-1)
		scoreFromEnd, ok := backward[state]
		if !ok {
			continue
		}

		if scoreFromEnd+scoreFromStart == targetScore {
			shortestPaths[state.Position] = struct{}{}
		}
	}
}

func DoWithInputPart02(world World) int {
	// forward
	startFw := State{
		Position:  world.Start,
		Direction: utils.Right.ToStep(),
	}

	// find the shortest path
	_, _, score, _ := alg.AStarEndFunc(startFw, isEnd(world.End), h(world.End), d(), n(world.Matrix))

	// find all paths
	_, gScoreFw, _, _ := alg.AStarEndFunc(startFw, func(state State) bool { return false }, h(world.End), d(), n(world.Matrix))

	// backward south
	startBwS := State{
		Position:  world.End,
		Direction: utils.Up.ToStep(),
	}

	// find all paths
	_, gScoreBwS, _, _ := alg.AStarEndFunc(startBwS, func(state State) bool { return false }, h(world.End), d(), n(world.Matrix))

	// backward west
	startBwW := State{
		Position:  world.End,
		Direction: utils.Left.ToStep(),
	}

	// find all paths
	_, gScoreBwW, _, _ := alg.AStarEndFunc(startBwW, func(state State) bool { return false }, h(world.End), d(), n(world.Matrix))

	// connect from both sides
	shortestPaths := make(map[utils.Vector2i]struct{})

	connectBothDirections(gScoreFw, gScoreBwS, shortestPaths, score)
	connectBothDirections(gScoreFw, gScoreBwW, shortestPaths, score)

	for path := range shortestPaths {
		world.Matrix.SetV(path, "O")
	}
	formatter := matrix.NonIndexedAdapter(matrix.FmtFmt[string]("%s"))
	str := matrix.StringFmtSeparatorIndexed(world.Matrix, true, "", formatter)
	fmt.Println(str)

	return len(shortestPaths)
}

func ParseInput(r io.Reader) World {
	var start utils.Vector2i
	var end utils.Vector2i

	parseItem := func(char rune, x, y int) string {
		if char == 'S' {
			start = utils.Vector2i{X: x, Y: y}
		} else if char == 'E' {
			end = utils.Vector2i{X: x, Y: y}
		}

		return string(char)
	}

	return World{
		Matrix: parsers.ParseToMatrixIndexed(r, parseItem),
		Start:  start,
		End:    end,
	}
}
