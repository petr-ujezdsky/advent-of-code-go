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
	Cost     int
}

type World struct {
	Matrix     matrix.Matrix[Item]
	Start, End utils.Vector2i
}

func neighbours(m matrix.Matrix[Item], slopesConstraint bool) func(origin State, path iterators.Iterator[State]) []State {
	return func(origin State, path iterators.Iterator[State]) []State {
		var neighbours []State

		currentTile := m.GetV(origin.Position)
		steps := utils.Direction4Steps[:]

		if slopesConstraint {
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
				Cost:     origin.Cost + 1,
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

var count = 0

func MaximizePathLength(world World, slopesConstraint bool) (int, State) {
	endPos := world.End

	visited := [20_000]bool{}
	visited[toVisitedIndex(world.Start, world.Matrix)] = true

	startState := State{
		Position: world.Start,
		Visited:  visited,
		Cost:     0,
	}

	cost := func(state State) int {
		count++
		fmt.Printf("Costs: %v (%v)\n", count, state.Cost)
		return -state.Cost
	}

	lowerBound := func(state State) int { return -1_000_000 }

	n := neighbours(world.Matrix, slopesConstraint)
	end := isEnd(endPos)
	nextStatesProvider := func(state State) ([]State, bool) {
		if end(state) {
			return nil, true
		}

		return n(state, nil), false
	}

	min, minState := alg.BranchAndBoundDeepFirst(startState, cost, lowerBound, nextStatesProvider)

	return -min, minState
}

func printSteps(world World, lastState State) {
	visited := lastState.Visited

	str := matrix.StringFmtSeparatorIndexed[Item](world.Matrix, true, "", func(value Item, x, y int) string {
		pos := utils.Vector2i{X: x, Y: y}
		visitedIndex := toVisitedIndex(pos, world.Matrix)

		if visited[visitedIndex] {
			return "O"
		}

		return string(world.Matrix.GetV(pos))
	})

	fmt.Println(str)
}

func DoWithInputPart01(world World) int {
	length, lastState := MaximizePathLength(world, true)

	printSteps(world, lastState)

	return length
}

func DoWithInputPart02(world World) int {
	length, lastState := MaximizePathLength(world, false)

	printSteps(world, lastState)

	return length
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
