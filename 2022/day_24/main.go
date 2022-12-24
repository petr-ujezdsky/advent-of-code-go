package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"io"
	"strings"
)

type Vector2i = utils.Vector2i

type World struct {
	ColumnBlizzards, RowBlizzards              [][]Blizzard
	Width, Height                              int
	BoundingRectangle                          utils.BoundingRectangle
	StartPosition, PreEndPosition, EndPosition Vector2i
	StartRemainingTime                         int
}

func (w World) IsBlizzardAt(remainingTime int, pos Vector2i) bool {
	elapsedTime := w.StartRemainingTime - remainingTime

	for _, blizzard := range w.ColumnBlizzards[pos.X] {
		if pos.Y == blizzard.Position(elapsedTime, w.Height) {
			return true
		}
	}

	for _, blizzard := range w.RowBlizzards[pos.Y] {
		if pos.X == blizzard.Position(elapsedTime, w.Width) {
			return true
		}
	}

	return false
}

type Blizzard struct {
	InitialPosition int
	Direction       int
}

func (b Blizzard) Position(time, size int) int {
	return utils.ModFloor(b.InitialPosition+time*b.Direction, size)
}

type State struct {
	Position      Vector2i
	RemainingTime int
}

func (s State) String(world World) string {
	sb := &strings.Builder{}

	for y := 0; y < world.BoundingRectangle.Height(); y++ {
		sb.WriteString("#")
		for x := 0; x < world.BoundingRectangle.Width(); x++ {
			pos := Vector2i{X: x, Y: y}

			if pos == s.Position {
				sb.WriteString("E")
				continue
			}

			if world.IsBlizzardAt(s.RemainingTime, pos) {
				sb.WriteString("@")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("#\n")
	}

	return sb.String()
}

func isEnd(world World) func(State) bool {
	return func(state State) bool {
		return state.Position == world.EndPosition
	}
}

func h(world World) func(State) int {
	return func(state State) int {
		return state.Position.Subtract(world.EndPosition).LengthManhattan()
	}
}

func d(_ World) func(State, State) int {
	return func(state1, state2 State) int {
		return state2.Position.Subtract(state1.Position).LengthManhattan()
	}
}

func moveStates(state State, world World, neighbours []State) []State {
	nextRemainingTime := state.RemainingTime - 1

	for _, step := range utils.Direction4Steps {
		nextPos := state.Position.Add(step)

		// ensure position is within world
		if !world.BoundingRectangle.Contains(nextPos) {
			continue
		}

		// ensure there is no blizzard
		if world.IsBlizzardAt(nextRemainingTime, nextPos) {
			continue
		}

		// all ok -> move there
		nextState := State{
			Position:      nextPos,
			RemainingTime: nextRemainingTime,
		}
		neighbours = append(neighbours, nextState)
	}
	return neighbours
}

func waitingStates(remainingTime int, position Vector2i, world World) []State {
	var states []State
	nextRemainingTime := remainingTime - 1

	for nextRemainingTime > 1 {
		if world.IsBlizzardAt(nextRemainingTime, position) {
			break
		}

		nextState := State{
			Position:      position,
			RemainingTime: nextRemainingTime,
		}
		states = append(states, nextState)

		nextRemainingTime--
	}

	return states
}

func neighbours(world World) func(state State) []State {
	return func(state State) []State {
		if state.RemainingTime <= 0 {
			return nil
		}

		// state just before end
		if state.Position == world.PreEndPosition {
			return []State{{
				Position:      world.EndPosition,
				RemainingTime: state.RemainingTime - 1,
			}}
		}

		var nextStates []State

		// try move to all directions
		nextStates = moveStates(state, world, nextStates)

		// wait
		nextStates = append(nextStates, waitingStates(state.RemainingTime, state.Position, world)...)

		return nextStates
	}
}

func DoWithInput(world World) int {
	start := State{
		Position:      world.StartPosition,
		RemainingTime: world.StartRemainingTime,
	}

	path, _, score, found := alg.AStarEndFunc[State](start, isEnd(world), h(world), d(world), neighbours(world))
	if !found {
		panic("Not found!")
	}

	for _, state := range path {
		fmt.Println(state.String(world))
		fmt.Println()
	}

	return score
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	columnBlizzards := make([][]Blizzard, 200)
	rowBlizzards := make([][]Blizzard, 200)
	y := -2
	width := 0

	for scanner.Scan() {
		y++
		width = len(scanner.Text()) - 2

		// skip first and last row
		if y == -1 || strings.HasPrefix(scanner.Text(), "##") {
			continue
		}

		for i, char := range scanner.Text() {
			// walls
			if i == 0 || i == len(scanner.Text())-1 {
				continue
			}

			// empty
			if char == '.' {
				continue
			}

			x := i - 1

			// save blizzard state
			switch char {
			case '>':
				blizzard := Blizzard{
					InitialPosition: x,
					Direction:       1,
				}
				rowBlizzards[y] = append(rowBlizzards[y], blizzard)
			case '<':
				blizzard := Blizzard{
					InitialPosition: x,
					Direction:       -1,
				}
				rowBlizzards[y] = append(rowBlizzards[y], blizzard)
			case '^':
				blizzard := Blizzard{
					InitialPosition: y,
					Direction:       -1,
				}
				columnBlizzards[x] = append(columnBlizzards[x], blizzard)
			case 'v':
				blizzard := Blizzard{
					InitialPosition: y,
					Direction:       1,
				}
				columnBlizzards[x] = append(columnBlizzards[x], blizzard)
			default:
				panic("Unknown char " + string(char))
			}
		}
	}

	height := y

	boundingRectangle := utils.BoundingRectangle{
		Horizontal: utils.IntervalI{Low: 0, High: width - 1},
		Vertical:   utils.IntervalI{Low: 0, High: height - 1},
	}

	return World{
		ColumnBlizzards:    columnBlizzards[0:width],
		RowBlizzards:       rowBlizzards[0:height],
		Width:              width,
		Height:             height,
		BoundingRectangle:  boundingRectangle,
		StartPosition:      Vector2i{X: 0, Y: -1},
		PreEndPosition:     Vector2i{X: width - 1, Y: height - 1},
		EndPosition:        Vector2i{X: width - 1, Y: height},
		StartRemainingTime: 30,
	}
}
