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

		return 1001
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

			nextPos := origin.Position.Add(dir)

			if nn, ok := m.GetVSafe(nextPos); !ok || nn == "#" {
				continue
			}

			nextState := State{
				Position:  nextPos,
				Direction: dir,
			}

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

func DoWithInputPart02(world World) int {
	return 0
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
