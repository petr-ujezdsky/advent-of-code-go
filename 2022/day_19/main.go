package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"io"
	"time"
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

func (m1 Materials) AddFromRobots(robotsCounts RobotsCounts) Materials {
	for i, robotsCount := range robotsCounts {
		materialType := MaterialType(i)
		m1[materialType] += robotsCount
	}
	return m1
}

func (m1 Materials) SubtractAndValidate(m2 Materials) (m Materials, valid bool) {
	valid = true
	for i := range m1 {
		m1[i] -= m2[i]
		if m1[i] < 0 {
			valid = false
		}
	}

	return m1, valid
}

type RobotsCounts [4]int

func (r RobotsCounts) AddRobot(materialType MaterialType) RobotsCounts {
	r[materialType]++
	return r
}

type Blueprint struct {
	Id          int
	RobotsCosts [4]Materials
}

type State struct {
	RemainingTime int
	Materials     Materials
	RobotsCounts  RobotsCounts
	PreviousState *State
}

func (s State) String() string {
	return fmt.Sprintf("%2d | Remaining time: %2d, mats: %v, robots: %v", 24-s.RemainingTime, s.RemainingTime, s.Materials, s.RobotsCounts)
}

func printState(state *State) {
	if state == nil {
		return
	}

	printState(state.PreviousState)
	fmt.Println(state)
}

func maxMaterials(m [4]Materials) (max Materials) {
	max = m[0]
	for _, materials := range m {
		for i, material := range materials {
			max[i] = utils.Max(max[i], material)
		}
	}

	return max
}
func maxGeodeCountInTime(blueprint Blueprint, remainingTime int) (int, State) {
	maxPrice := maxMaterials(blueprint.RobotsCosts)

	cost := func(state State) int {
		// maximizing geodes count
		return -state.Materials[Geode]
		//return -state.Materials[Clay]
	}

	lowerBound := func(state State) int {
		// start with current geodes count, then add all geodes collected by current robots and finally add all geodes
		// collected from by robots created each remaining minute
		return -(state.Materials[Geode] + state.RemainingTime*state.RobotsCounts[Geode] + utils.SumNtoM(0, state.RemainingTime-1))
		//return -(state.Materials[Clay] + state.RemainingTime*state.RobotsCounts[Clay] + utils.SumNtoM(0, state.RemainingTime-1))
		//return math.MinInt
	}

	nextStatesProvider := func(state State) ([]State, bool) {
		if state.RemainingTime <= 0 {
			return nil, true
		}

		var states []State

		// do not buy anything
		nextRemainingTime := state.RemainingTime - 1
		nextState := State{
			RemainingTime: nextRemainingTime,
			Materials:     state.Materials.AddFromRobots(state.RobotsCounts),
			RobotsCounts:  state.RobotsCounts,
			PreviousState: &state,
		}
		states = append(states, nextState)

		for _, materialType := range [4]MaterialType{Ore, Clay, Obsidian, Geode} {
			// no need to have more robots than the max price
			if state.RobotsCounts[materialType] == maxPrice[materialType] && maxPrice[materialType] > 0 {
				continue
			}

			// buy robot
			matsBuyedRobot, buyable := state.Materials.SubtractAndValidate(blueprint.RobotsCosts[materialType])
			//if buyable && (materialType != Ore || state.RobotsCounts[Ore] < 2) {
			if buyable {
				// collect materials, without new robot
				nextMaterials := matsBuyedRobot.AddFromRobots(state.RobotsCounts)

				// add robot
				nextRobotsCounts := state.RobotsCounts.AddRobot(materialType)

				// shift time
				nextRemainingTime := state.RemainingTime - 1

				// next state
				nextState := State{
					RemainingTime: nextRemainingTime,
					Materials:     nextMaterials,
					RobotsCounts:  nextRobotsCounts,
					PreviousState: &state,
				}

				states = append(states, nextState)
			}
		}

		return states, len(states) == 0
	}

	initialState := State{
		RemainingTime: remainingTime,
		Materials:     [4]int{0, 0, 0, 0},
		RobotsCounts:  RobotsCounts{1, 0, 0, 0},
	}

	min, minState := alg.BranchAndBoundDeepFirst(initialState, cost, lowerBound, nextStatesProvider)
	return -min, minState
}

func DoWithInput(world World) int {
	sum := 0

	for _, blueprint := range world.Blueprints {
		fmt.Printf("Computing blueprint #%v...\n", blueprint.Id)
		geodes, _ := maxGeodeCountInTime(blueprint, 24)
		sum += blueprint.Id * geodes
		fmt.Printf(" - produces max %v geodes\n\n", geodes)
	}

	return sum
}

func DoWithInputParallel(world World) int {
	qualityLevelChannels := make([]<-chan int, len(world.Blueprints))

	for i, blueprint := range world.Blueprints {
		channel := make(chan int)
		qualityLevelChannels[i] = channel
		go func(b Blueprint, ch chan int) {
			fmt.Printf("Computing blueprint #%v...\n", b.Id)
			geodes, _ := maxGeodeCountInTime(b, 24)
			fmt.Printf("%v produces max %v geodes\n", b.Id, geodes)
			ch <- b.Id * geodes
		}(blueprint, channel)
	}

	sum := 0
	for _, channel := range qualityLevelChannels {
		sum += <-channel
	}

	return sum
}

func DoWithInputParallelFirstN(world World, n int) int {
	qualityLevelChannels := make([]<-chan int, n)

	for i, blueprint := range world.Blueprints[0:n] {
		channel := make(chan int)
		qualityLevelChannels[i] = channel
		go func(b Blueprint, ch chan int) {
			start := time.Now()
			fmt.Printf("Computing blueprint #%v...\n", b.Id)
			geodes, _ := maxGeodeCountInTime(b, 32)
			elapsed := time.Since(start)
			fmt.Printf("%v produces max %v geodes (took %s)\n", b.Id, geodes, elapsed)
			ch <- geodes
		}(blueprint, channel)
	}

	product := 1
	for _, channel := range qualityLevelChannels {
		product *= <-channel
	}

	return product
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
