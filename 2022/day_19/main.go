package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
)

type World struct {
	Blueprints []Blueprint
}

type MaterialType int

const (
	Ore MaterialType = iota
	Clay
	Obsidian
	Geode
)

type Materials [4]int

func (m1 Materials) IsEnoughFor(m2 Materials) bool {
	for i, have := range m1 {
		if m2[i] > have {
			return false
		}
	}

	return true
}

func (m1 Materials) Add(m2 Materials) Materials {
	for i := range m1 {
		m1[i] += m2[i]
	}

	return m1
}

func (m1 Materials) SubtractAndRemainder(m2 Materials) (m Materials, valid bool) {
	valid = true
	for i := range m1 {
		m1[i] -= m2[i]
		if m1[i] < 0 {
			valid = false
		}
	}

	return m1, valid
}

type Blueprint struct {
	Id          int
	RobotsCosts [4]Materials
}

type Robot struct {
	Type               MaterialType
	SinceRemainingTime int
}

func (r Robot) String() string {
	return fmt.Sprintf("%d", r.Type)
}

func (r Robot) GenerateMaterials() Materials {
	m := Materials{}
	m[r.Type] = 1
	return m
}

type State struct {
	RemainingTime int
	Materials     Materials
	Robots        []Robot
	PreviousState *State
}

func (s State) String() string {
	return fmt.Sprintf("Remaining time: %2d, mats: %v, robots: %v", s.RemainingTime, s.Materials, s.Robots)
}

func printState(state *State) {
	if state == nil {
		return
	}

	printState(state.PreviousState)
	fmt.Println(state)
}

func generateMaterials(remainingTime int, robots []Robot) Materials {
	mats := Materials{}
	for _, robot := range robots {
		if robot.SinceRemainingTime > remainingTime {
			mats = mats.Add(robot.GenerateMaterials())
		}
	}
	return mats
}

func maxGeodeCountInTime(blueprint Blueprint) (int, State) {
	cost := func(state State) int {
		// maximizing geodes count
		return -state.Materials[Geode]
	}

	lowerBound := func(state State) int {
		return cost(state) - state.RemainingTime
	}

	nextStatesProvider := func(state State) []State {
		if state.RemainingTime <= 0 {
			return nil
		}

		var states []State

		for _, materialType := range [4]MaterialType{Geode, Obsidian, Clay, Ore} {
			// buy robot
			matsBuyedRobot, buyable := state.Materials.SubtractAndRemainder(blueprint.RobotsCosts[materialType])
			if buyable {
				// add robot
				robot := Robot{
					Type:               materialType,
					SinceRemainingTime: state.RemainingTime,
				}

				nextRobots := slices.CloneAndAdd(state.Robots, robot)

				// calculate mats, including new robot
				nextRemainingTime := state.RemainingTime - 1
				nextState := State{
					RemainingTime: nextRemainingTime,
					Materials:     matsBuyedRobot.Add(generateMaterials(nextRemainingTime, nextRobots)),
					Robots:        nextRobots,
					PreviousState: &state,
				}

				states = append(states, nextState)
			}
		}

		// do not buy anything
		nextRemainingTime := state.RemainingTime - 1
		nextState := State{
			RemainingTime: nextRemainingTime,
			Materials:     state.Materials.Add(generateMaterials(nextRemainingTime, state.Robots)),
			Robots:        state.Robots,
			PreviousState: &state,
		}
		states = append(states, nextState)

		return states
	}

	remainingTime := 24
	initialState := State{
		RemainingTime: remainingTime,
		Materials:     [4]int{0, 0, 0, 0},
		Robots:        []Robot{{Type: Ore, SinceRemainingTime: remainingTime}},
	}

	min, minState := alg.BranchAndBoundDeepFirst(initialState, cost, lowerBound, nextStatesProvider)
	return -min, minState
}

func DoWithInput(world World) int {
	sum := 0

	for _, blueprint := range world.Blueprints {
		fmt.Printf("Computing blueprint #%v...\n", blueprint.Id)
		geodes, _ := maxGeodeCountInTime(blueprint)
		sum += blueprint.Id * geodes
		fmt.Printf(" - produces max %v geodes\n\n", geodes)
	}

	return sum
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var blueprints []Blueprint
	for scanner.Scan() {
		ints := utils.ExtractInts(scanner.Text(), false)

		id := ints[0]
		robotsCosts := [4]Materials{
			{ints[1], 0, 0, 0},
			{ints[2], 0, 0, 0},
			{ints[3], ints[4], 0, 0},
			{ints[5], 0, ints[6], 0},
		}

		item := Blueprint{
			Id:          id,
			RobotsCosts: robotsCosts,
		}

		blueprints = append(blueprints, item)
	}

	return World{Blueprints: blueprints}
}
