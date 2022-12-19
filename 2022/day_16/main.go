package main

import (
	"bufio"
	_ "embed"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
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
	ClosedValvesSet  utils.BitSet128
	AllNodes         []*ValveNode
	AllNodesSorted   []*ValveNode
	RemainingTime    int
	PressureReleased int
}

func (s WorldState) ClosedValvesSlice() []*ValveNode {
	var closed []*ValveNode
	for i, node := range s.AllNodes {
		if s.ClosedValvesSet.Contains(i) && node.FlowRate > 0 {
			closed = append(closed, node)
		}
	}
	return closed
}

type WorldState2 struct {
	CurrentNode      *ValveNode
	ClosedValvesSet  utils.BitSet128
	RemainingTime    int
	PressureReleased int
}

func (s WorldState2) ClosedValvesSlice(allNodes []*ValveNode) []*ValveNode {
	var closed []*ValveNode
	for i, node := range allNodes {
		if s.ClosedValvesSet.Contains(i) && node.FlowRate > 0 {
			closed = append(closed, node)
		}
	}
	return closed
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

func findMaxPressureReleaseStateMinMax(state WorldState, distances utils.MatrixInt, best *int) int {
	max := state.PressureReleased

	closedValves := state.ClosedValvesSlice()
	for _, closedValve := range closedValves {
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
		ClosedValvesSet:  utils.NewFullBitSet128(),
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

	next := func(state *WorldState2) []*WorldState2 {
		var nextStates []*WorldState2

		closedValves := state.ClosedValvesSlice(world.AllNodes)
		for _, closedValve := range closedValves {
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

		return nextStates
	}

	initialState := &WorldState2{
		CurrentNode:      world.RootNode,
		ClosedValvesSet:  utils.NewFullBitSet128(),
		RemainingTime:    30,
		PressureReleased: 0,
	}

	maxPressureReleased, _ := utils.BranchAndBound(initialState, cost, lowerBound, next)

	return -maxPressureReleased
}

func computeDistances(world World) utils.MatrixInt {
	distances := utils.NewMatrixInt(len(world.AllNodes), len(world.AllNodes))

	h := func(_ *ValveNode) int { return 0 }
	d := func(_, _ *ValveNode) int { return 1 }
	neighbours := func(node *ValveNode) []*ValveNode { return node.Children }

	for i, nodeFrom := range world.AllNodes {
		_, allCosts, _, _ := utils.AStar(nodeFrom, nil, h, d, neighbours)

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

	allNodes := utils.MapValues(nodes)
	for i, node := range allNodes {
		node.Id = i
	}

	// sort valves by flow rate
	allNodesSorted := utils.ShallowCopy(allNodes)
	sort.Slice(allNodesSorted, func(i, j int) bool { return allNodesSorted[i].FlowRate > allNodesSorted[j].FlowRate })

	return World{
		RootNode:       nodes["AA"],
		AllNodes:       allNodes,
		AllNodesSorted: allNodesSorted,
	}
}
