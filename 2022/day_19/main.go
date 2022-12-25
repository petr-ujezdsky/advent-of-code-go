package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
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

func (r Robot) GenerateMaterials() Materials {
	m := Materials{}
	m[r.Type] = 1
	return m
}

type State struct {
	RemainingTime int
	Materials     Materials
	Robots        []Robot
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

func maxGeodeCountInTime(blueprint Blueprint) int {

	cost := func(state State) int {
		// maximizing geodes count
		return -state.Materials[Geode]
	}

	lowerBound := func(state State) int {
		return -state.RemainingTime
	}

	nextStatesProvider := func(state State) []State {
		if state.RemainingTime <= 0 {
			return nil
		}

		var states []State

		//// buy ore robot
		//matsBuyedOreRobot, buyable := state.Materials.SubtractAndRemainder(blueprint.RobotsCosts[Ore])
		//if buyable {
		//	// add robot
		//	robot := Robot{
		//		Type:               Ore,
		//		SinceRemainingTime: state.RemainingTime,
		//	}
		//
		//	robots := append(utils.ShallowCopy(state.Robots), robot)
		//
		//	// calculate mats, including new robot
		//	nextState := State{
		//		RemainingTime: state.RemainingTime - 1,
		//		Materials:     matsBuyedOreRobot.Add(robot.GenerateMaterials()),
		//		Robots:        robots,
		//	}
		//
		//	states = append(states, nextState)
		//}
		for _, materialType := range [4]MaterialType{Geode, Obsidian, Clay, Ore} {
			// buy robot
			matsBuyedRobot, buyable := state.Materials.SubtractAndRemainder(blueprint.RobotsCosts[materialType])
			if buyable {
				// add robot
				robot := Robot{
					Type:               materialType,
					SinceRemainingTime: state.RemainingTime,
				}

				nextRobots := append(state.Robots, robot)

				// calculate mats, including new robot
				nextRemainingTime := state.RemainingTime - 1
				nextState := State{
					RemainingTime: nextRemainingTime,
					Materials:     matsBuyedRobot.Add(generateMaterials(nextRemainingTime, nextRobots)),
					Robots:        nextRobots,
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
		}
		states = append(states, nextState)

		return states
	}

	remainingTime := 24
	initialState := State{
		RemainingTime: remainingTime,
		Materials:     [4]int{},
		Robots:        []Robot{{Type: Ore, SinceRemainingTime: remainingTime}},
	}

	min, _ := alg.BranchAndBoundDeepFirst(initialState, cost, lowerBound, nextStatesProvider)
	return min
}

func DoWithInput(world World) int {
	sum := 0

	for _, blueprint := range world.Blueprints {
		geodes := maxGeodeCountInTime(blueprint)
		sum += blueprint.Id * geodes
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
