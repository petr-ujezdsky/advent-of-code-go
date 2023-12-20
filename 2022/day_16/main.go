package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/alg"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/collections"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/maps"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/matrix"
	"github.com/petr-ujezdsky/advent-of-code-go/utils/slices"
	"io"
	"math"
	"regexp"
	"sort"
)

var valvesRegex = regexp.MustCompile(`[A-Z]{2}`)

type ValveNode struct {
	Id       int
	Name     string
	FlowRate int
	Closed   bool
	Children []*ValveNode
}

type World struct {
	RootNode       *ValveNode
	AllNodes       []*ValveNode
	AllNodesSorted []*ValveNode
}

type WorldState struct {
	CurrentNode      *ValveNode
	ClosedValvesSet  collections.BitSet128
	AllNodes         []*ValveNode
	AllNodesSorted   []*ValveNode
	RemainingTime    int
	PressureReleased int
}

type WorldState2 struct {
	CurrentNode      *ValveNode
	ClosedValvesSet  collections.BitSet128
	RemainingTime    int
	PressureReleased int
}

type PlayerState struct {
	CurrentNode   *ValveNode
	RemainingTime int
}

type WorldState3 struct {
	Player1State, Player2State PlayerState
	ValveOpenDuration          []int
	PressureReleased           int
}

func maxPossibleReleasedPressure(state WorldState) int {
	remainingTime := state.RemainingTime
	maxReleasedPressure := state.PressureReleased

	for _, closedValve := range state.AllNodesSorted {
		if !state.ClosedValvesSet.Contains(closedValve.Id) || closedValve.FlowRate == 0 {
			continue
		}

		// not enough time to get here and open valve
		if remainingTime < 2 {
			return maxReleasedPressure
		}

		maxReleasedPressure += (remainingTime - 2) * closedValve.FlowRate
		remainingTime -= 2
	}

	return maxReleasedPressure
}

func maxPossibleReleasedPressure2(state *WorldState2, allNodesSorted []*ValveNode) int {
	remainingTime := state.RemainingTime
	maxReleasedPressure := state.PressureReleased

	for _, closedValve := range allNodesSorted {
		if !state.ClosedValvesSet.Contains(closedValve.Id) || closedValve.FlowRate == 0 {
			continue
		}

		// not enough time to get here and open valve
		if remainingTime < 2 {
			return maxReleasedPressure
		}

		maxReleasedPressure += (remainingTime - 2) * closedValve.FlowRate
		remainingTime -= 2
	}

	return maxReleasedPressure
}

func maxPossibleRemainingReleasedPressurePlayer(state *WorldState3, player PlayerState, allNodesSorted []*ValveNode, playerId int) int {
	remainingReleasedPressure := 0
	remainingTime := player.RemainingTime

	i := 0
	for _, closedValve := range allNodesSorted {
		if state.ValveOpenDuration[closedValve.Id] > 0 || closedValve.FlowRate == 0 {
			continue
		}

		// not enough time to get here and open valve
		if remainingTime < 2 {
			break
		}

		if i%2 == playerId {
			remainingReleasedPressure += (remainingTime - 2) * closedValve.FlowRate
			remainingTime -= 2
		}

		i++
	}

	return remainingReleasedPressure
}

func maxPossibleReleasedPressure3(state *WorldState3, allNodesSorted []*ValveNode) int {
	maxReleasedPressure := state.PressureReleased

	// player 1
	maxReleasedPressure += maxPossibleRemainingReleasedPressurePlayer(state, state.Player1State, allNodesSorted, 0)

	// player 2
	maxReleasedPressure += maxPossibleRemainingReleasedPressurePlayer(state, state.Player2State, allNodesSorted, 1)

	return maxReleasedPressure
}

func findMaxPressureReleaseStateMinMax(state WorldState, distances matrix.MatrixInt, best *int) int {
	max := state.PressureReleased

	for _, closedValve := range state.AllNodes {
		if !state.ClosedValvesSet.Contains(closedValve.Id) || closedValve.FlowRate == 0 {
			continue
		}

		pathCost := distances.Columns[state.CurrentNode.Id][closedValve.Id]
		moveAndOpenCost := pathCost + 1
		// not enough time to get here and open valve
		if moveAndOpenCost >= state.RemainingTime {
			continue
		}

		nextRemainingTime := state.RemainingTime - moveAndOpenCost
		currentPressureReleased := nextRemainingTime * closedValve.FlowRate

		nextClosedValvesSet := state.ClosedValvesSet.Clone()
		nextClosedValvesSet.Remove(closedValve.Id)

		nextState := WorldState{
			AllNodes:         state.AllNodes,
			AllNodesSorted:   state.AllNodesSorted,
			CurrentNode:      closedValve,
			ClosedValvesSet:  nextClosedValvesSet,
			RemainingTime:    nextRemainingTime,
			PressureReleased: state.PressureReleased + currentPressureReleased,
		}

		if maxPossibleReleasedPressure(nextState) < *best {
			continue
		}

		bestTotalNextPressureReleased := findMaxPressureReleaseStateMinMax(nextState, distances, best)

		if bestTotalNextPressureReleased > *best {
			*best = bestTotalNextPressureReleased
		}

		if bestTotalNextPressureReleased > max {
			max = bestTotalNextPressureReleased
		}
	}

	return max
}

func FindMaxPressureReleaseStateMinMax(world World) int {
	initialState := WorldState{
		AllNodes:         world.AllNodes,
		AllNodesSorted:   world.AllNodesSorted,
		CurrentNode:      world.RootNode,
		ClosedValvesSet:  collections.NewFullBitSet128(),
		RemainingTime:    30,
		PressureReleased: 0,
	}
	distances := computeDistances(world)
	best := math.MinInt

	maxPressureReleased := findMaxPressureReleaseStateMinMax(initialState, distances, &best)

	return maxPressureReleased
}

func FindMaxPressureReleaseStateMinMaxGeneralized(world World) int {
	distances := computeDistances(world)

	cost := func(state *WorldState2) int {
		return -state.PressureReleased
	}

	lowerBound := func(state *WorldState2) int {
		return -maxPossibleReleasedPressure2(state, world.AllNodesSorted)
	}

	next := func(state *WorldState2) ([]*WorldState2, bool) {
		var nextStates []*WorldState2

		for _, closedValve := range world.AllNodes {
			if !state.ClosedValvesSet.Contains(closedValve.Id) || closedValve.FlowRate == 0 {
				continue
			}

			pathCost := distances.Columns[state.CurrentNode.Id][closedValve.Id]
			moveAndOpenCost := pathCost + 1
			// not enough time to get here and open valve
			if moveAndOpenCost >= state.RemainingTime {
				continue
			}

			nextRemainingTime := state.RemainingTime - moveAndOpenCost
			currentPressureReleased := nextRemainingTime * closedValve.FlowRate

			nextClosedValvesSet := state.ClosedValvesSet.Clone()
			nextClosedValvesSet.Remove(closedValve.Id)

			nextState := &WorldState2{
				CurrentNode:      closedValve,
				ClosedValvesSet:  nextClosedValvesSet,
				RemainingTime:    nextRemainingTime,
				PressureReleased: state.PressureReleased + currentPressureReleased,
			}

			nextStates = append(nextStates, nextState)
		}

		return nextStates, len(nextStates) == 0
	}

	initialState := &WorldState2{
		CurrentNode:      world.RootNode,
		ClosedValvesSet:  collections.NewFullBitSet128(),
		RemainingTime:    30,
		PressureReleased: 0,
	}

	maxPressureReleased, _ := alg.BranchAndBoundDeepFirst(initialState, cost, lowerBound, next)

	return -maxPressureReleased
}

func nextPlayerState(player PlayerState, valveOpenDuration []int, allNodes []*ValveNode, distances matrix.MatrixInt) ([]PlayerState, [][]int) {
	var nextStates []PlayerState
	var nextValveOpenDurations [][]int
	for _, closedValve := range allNodes {
		if valveOpenDuration[closedValve.Id] > 0 || closedValve.FlowRate == 0 {
			continue
		}

		pathCost := distances.Columns[player.CurrentNode.Id][closedValve.Id]
		moveAndOpenCost := pathCost + 1
		// not enough time to get here and open valve
		if moveAndOpenCost >= player.RemainingTime {
			continue
		}

		nextRemainingTime := player.RemainingTime - moveAndOpenCost
		//currentPressureReleased := nextRemainingTime * closedValve.FlowRate

		nextValveOpenDuration := slices.Clone(valveOpenDuration)
		nextValveOpenDuration[closedValve.Id] = nextRemainingTime

		nextStates = append(nextStates, PlayerState{
			CurrentNode:   closedValve,
			RemainingTime: nextRemainingTime,
		})

		nextValveOpenDurations = append(nextValveOpenDurations, nextValveOpenDuration)
	}

	return nextStates, nextValveOpenDurations
}

func worldStateCost(valveOpenDuration []int, allNodes []*ValveNode) int {
	sum := 0

	for i, openDuration := range valveOpenDuration {
		sum += openDuration * allNodes[i].FlowRate
	}

	return sum
}

func FindMaxPressureReleasedWithElephant(world World) int {
	distances := computeDistances(world)

	cost := func(state *WorldState3) int {
		return -state.PressureReleased
	}

	lowerBound := func(state *WorldState3) int {
		return -maxPossibleReleasedPressure3(state, world.AllNodesSorted)
	}

	next := func(state *WorldState3) ([]*WorldState3, bool) {
		var nextStates []*WorldState3

		// player 1
		nextP1States, nextP1ValveOpenDurations := nextPlayerState(state.Player1State, state.ValveOpenDuration, world.AllNodes, distances)

		for ip1, p1 := range nextP1States {
			p1ValveOpenDurations := nextP1ValveOpenDurations[ip1]

			// player 2
			nextP2States, nextP2ValveOpenDurations := nextPlayerState(state.Player2State, p1ValveOpenDurations, world.AllNodes, distances)

			for ip2, p2 := range nextP2States {
				p2ValveOpenDurations := nextP2ValveOpenDurations[ip2]

				nextState := &WorldState3{
					Player1State:      p1,
					Player2State:      p2,
					ValveOpenDuration: p2ValveOpenDurations,
					PressureReleased:  worldStateCost(p2ValveOpenDurations, world.AllNodes),
				}

				nextStates = append(nextStates, nextState)
			}
		}

		return nextStates, len(nextStates) == 0
	}

	initialState := &WorldState3{
		Player1State: PlayerState{
			CurrentNode:   world.RootNode,
			RemainingTime: 26,
		},
		Player2State: PlayerState{
			CurrentNode:   world.RootNode,
			RemainingTime: 26,
		},
		ValveOpenDuration: make([]int, len(world.AllNodes)),
		PressureReleased:  0,
	}

	maxPressureReleased, _ := alg.BranchAndBoundDeepFirst(initialState, cost, lowerBound, next)

	return -maxPressureReleased
}

func computeDistances(world World) matrix.MatrixInt {
	distances := matrix.NewMatrixInt(len(world.AllNodes), len(world.AllNodes))

	h := func(_ *ValveNode) int { return 0 }
	d := func(_, _ *ValveNode) int { return 1 }
	neighbours := func(node *ValveNode) []*ValveNode { return node.Children }

	for i, nodeFrom := range world.AllNodes {
		_, allCosts, _, _ := alg.AStar(nodeFrom, nil, h, d, neighbours)

		for j, nodeTo := range world.AllNodes {
			distances.Columns[i][j] = allCosts[nodeTo]
		}
	}

	return distances
}

func getOrCreateNodes(names []string, nodes map[string]*ValveNode) []*ValveNode {
	resultNodes := make([]*ValveNode, len(names))

	for i, name := range names {
		name = name[0:2]
		if node, ok := nodes[name]; ok {
			resultNodes[i] = node
		} else {
			node = &ValveNode{Name: name, Closed: true}
			resultNodes[i] = node
			nodes[name] = node
		}
	}

	return resultNodes
}

func ParseInput(r io.Reader) World {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	nodes := make(map[string]*ValveNode)

	for scanner.Scan() {
		flowRate := utils.ExtractInts(scanner.Text(), false)[0]
		names := valvesRegex.FindAllString(scanner.Text(), -1)

		name := names[0]

		node := getOrCreateNodes([]string{name}, nodes)[0]

		node.FlowRate = flowRate
		node.Children = append(node.Children, getOrCreateNodes(names[1:], nodes)...)
	}

	allNodes := maps.Values(nodes)
	for i, node := range allNodes {
		node.Id = i
	}

	// sort valves by flow rate
	allNodesSorted := slices.Clone(allNodes)
	sort.Slice(allNodesSorted, func(i, j int) bool { return allNodesSorted[i].FlowRate > allNodesSorted[j].FlowRate })

	return World{
		RootNode:       nodes["AA"],
		AllNodes:       allNodes,
		AllNodesSorted: allNodesSorted,
	}
}
