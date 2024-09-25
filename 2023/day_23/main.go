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

type Item rune

type State struct {
	Position utils.Vector2i
	Visited  [20_000]bool
}

type World struct {
	Matrix     matrix.Matrix[Item]
	Start, End utils.Vector2i
}

func h(endPos utils.Vector2i) func(state State) int {
	return func(state State) int {
		// manhattan distance
		return -utils.ManhattanDistance(state.Position, endPos)
	}
}

func d() func(State, State) int {
	return func(nodeFrom, nodeTo State) int {
		return -utils.ManhattanDistance(nodeFrom.Position, nodeTo.Position)
	}
}

func neighbours(m matrix.Matrix[Item]) func(origin State, path iterators.Iterator[State]) []State {
	return func(origin State, path iterators.Iterator[State]) []State {
		var neighbours []State

		currentTile := m.GetV(origin.Position)
		steps := utils.Direction4Steps[:]

		switch currentTile {
		case '>':
			steps = []utils.Vector2i{utils.Right.ToStep()}
		case '<':
			steps = []utils.Vector2i{utils.Left.ToStep()}
		case '^':
			steps = []utils.Vector2i{utils.Down.ToStep()}
		case 'v':
			steps = []utils.Vector2i{utils.Up.ToStep()}
		case '.':
			steps = utils.Direction4Steps[:]
		default:
			panic(fmt.Sprintf("Unknown current tile %v", string(currentTile)))
		}

		for _, dir := range steps {
			nextPos := origin.Position.Add(dir)

			nextTile, ok := m.GetVSafe(nextPos)
			if !ok {
				// out of bounds of map
				continue
			}

			if nextTile == '#' {
				// can not step on forest
				continue
			}

			nextVisitedIndex := toVisitedIndex(nextPos, m)
			if origin.Visited[nextVisitedIndex] {
				// already visited
				continue
			}

			nextState := State{
				Position: nextPos,
				Visited:  origin.Visited,
			}

			nextState.Visited[nextVisitedIndex] = true

			neighbours = append(neighbours, nextState)
		}

		return neighbours
	}
}

func toVisitedIndex(pos utils.Vector2i, m matrix.Matrix[Item]) int {
	return pos.Y*m.Width + pos.X
}

func isEnd(endPos utils.Vector2i) func(state State) bool {
	return func(state State) bool {
		return state.Position == endPos
	}
}

func MaximizePathLength(world World) ([]State, map[State]int, int, bool) {
	endPos := world.End

	visited := [20000]bool{}
	visited[toVisitedIndex(world.Start, world.Matrix)] = true

	startState := State{
		Position: world.Start,
		Visited:  visited,
	}

	return alg.AStarEndFunc(startState, isEnd(endPos), h(endPos), d(), neighbours(world.Matrix))
}

func DoWithInputPart01(world World) int {
	_, _, length, ok := MaximizePathLength(world)
	if !ok {
		panic("No path found")
	}

	return -length
}

func DoWithInputPart02(world World) int {
	return 0
}

func ParseInput(r io.Reader) World {
	parseItem := func(char rune) Item {
		return Item(char)
	}

	m := parsers.ParseToMatrix(r, parseItem)

	return World{
		Matrix: m,
		Start:  utils.Vector2i{X: 1},
		End:    utils.Vector2i{X: m.Width - 2, Y: m.Height - 1},
	}
}
