package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"io"
	"strconv"
	"strings"
)

type Vector2i = utils.Vector2i

type World struct {
	ColumnBlizzards, RowBlizzards              [][]Blizzard
	Width, Height                              int
	BoundingRectangle                          utils.BoundingRectangle
	StartPosition, PreEndPosition, EndPosition Vector2i
}

func (w World) IsBlizzardAt(elapsedTime int, pos Vector2i) bool {
	// start / end state
	if !w.BoundingRectangle.Contains(pos) {
		return false
	}

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

func (w World) BlizzardsAt(elapsedTime int, pos Vector2i) []Blizzard {
	// start state
	if pos.Y < 0 {
		return []Blizzard{}
	}

	var blizzards []Blizzard

	for _, blizzard := range w.ColumnBlizzards[pos.X] {
		if pos.Y == blizzard.Position(elapsedTime, w.Height) {
			blizzards = append(blizzards, blizzard)
		}
	}

	for _, blizzard := range w.RowBlizzards[pos.Y] {
		if pos.X == blizzard.Position(elapsedTime, w.Width) {
			blizzards = append(blizzards, blizzard)
		}
	}

	return blizzards
}

type Blizzard struct {
	InitialPosition int
	Direction       int
	Char            rune
}

func (b Blizzard) Position(time, size int) int {
	return utils.ModFloor(b.InitialPosition+time*b.Direction, size)
}

type State struct {
	Position      Vector2i
	ElapsedTime   int
	PreviousState *State
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

			blizzards := world.BlizzardsAt(s.ElapsedTime, pos)
			switch len(blizzards) {
			case 0:
				sb.WriteString(".")
			case 1:
				sb.WriteRune(blizzards[0].Char)

			default:
				sb.WriteString(strconv.Itoa(len(blizzards)))
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
		return state2.ElapsedTime - state1.ElapsedTime
	}
}

func moveStates(state State, world World, neighbours []State) []State {
	nextElapsedTime := state.ElapsedTime + 1

	for _, step := range utils.Direction4Steps {
		nextPos := state.Position.Add(step)

		// ensure position is within world
		if !world.BoundingRectangle.Contains(nextPos) {
			continue
		}

		// ensure there is no blizzard
		if world.IsBlizzardAt(nextElapsedTime, nextPos) {
			continue
		}

		// all ok -> move there
		nextState := State{
			Position:      nextPos,
			ElapsedTime:   nextElapsedTime,
			PreviousState: &state,
		}
		neighbours = append(neighbours, nextState)
	}
	return neighbours
}

func waitingStates(state State, world World) []State {
	var states []State
	position := state.Position

	nextElapsedTime := state.ElapsedTime + 1
	waitingDuration := 1
	nextState := &state
	for !world.IsBlizzardAt(nextElapsedTime, position) && waitingDuration < world.Width*world.Height {
		nextState = &State{
			Position:      position,
			ElapsedTime:   nextElapsedTime,
			PreviousState: nextState,
		}
		states = append(states, *nextState)

		nextElapsedTime++
		waitingDuration++
	}

	return states
}

func neighbours(world World) func(state State) []State {
	return func(state State) []State {
		// state just before end
		if state.Position == world.PreEndPosition {
			return []State{{
				Position:      world.EndPosition,
				ElapsedTime:   state.ElapsedTime + 1,
				PreviousState: &state,
			}}
		}

		var nextStates []State

		// try move to all directions
		nextStates = moveStates(state, world, nextStates)

		// wait
		nextStates = append(nextStates, waitingStates(state, world)...)

		return nextStates
	}
}

func ShortestPathAStar(world World) int {
	start := State{
		Position:    world.StartPosition,
		ElapsedTime: 0,
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

func ShortestPathBranchAndBound(world World) int {
	cost := func(state State) int {
		return state.ElapsedTime
	}

	lowerBound := func(state State) int {
		return cost(state) + state.Position.Subtract(world.EndPosition).LengthManhattan()
	}

	start := State{
		Position:    world.StartPosition,
		ElapsedTime: 0,
	}

	min, minState := alg.BranchAndBoundDeepFirst(start, cost, lowerBound, neighbours(world))

	state := &minState
	for state != nil {
		fmt.Println(state.String(world))
		fmt.Println()
		state = state.PreviousState
	}

	return min
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
					Char:            char,
				}
				rowBlizzards[y] = append(rowBlizzards[y], blizzard)
			case '<':
				blizzard := Blizzard{
					InitialPosition: x,
					Direction:       -1,
					Char:            char,
				}
				rowBlizzards[y] = append(rowBlizzards[y], blizzard)
			case '^':
				blizzard := Blizzard{
					InitialPosition: y,
					Direction:       -1,
					Char:            char,
				}
				columnBlizzards[x] = append(columnBlizzards[x], blizzard)
			case 'v':
				blizzard := Blizzard{
					InitialPosition: y,
					Direction:       1,
					Char:            char,
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
		ColumnBlizzards:   columnBlizzards[0:width],
		RowBlizzards:      rowBlizzards[0:height],
		Width:             width,
		Height:            height,
		BoundingRectangle: boundingRectangle,
		StartPosition:     Vector2i{X: 0, Y: -1},
		PreEndPosition:    Vector2i{X: width - 1, Y: height - 1},
		EndPosition:       Vector2i{X: width - 1, Y: height},
	}
}
