package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"io"
	"math"
	"regexp"
	"sort"
)

var valvesRegex = regexp.MustCompile(`[A-Z]{2}`)

type World struct {
	RootNode    *ValveNode
	AllNodes    []*ValveNode
	AllNodesMap map[string]*ValveNode
}

type WorldState struct {
	World
	CurrentNode      *ValveNode
	ClosedValves     []*ValveNode
	RemainingTime    int
	PressureReleased int
}

type ValveNode struct {
	Id       int
	Name     string
	FlowRate int
	Closed   bool
	Children []*ValveNode
}

//func h(closedValves []*ValveNode, remainingTime int) func(*ValveNode) int {
//	return func(node *ValveNode) int {
//		// sort valves by flow rate
//		sort.Slice(closedValves, func(i, j int) bool { return closedValves[i].FlowRate > closedValves[j].FlowRate })
//
//		sum := 0
//		for i := remainingTime; i >= 0; i-- {
//			// step
//			i--
//			if i <= 0 {
//				break
//			}
//			// open valve
//			i--
//			sum += i *
//		}
//	}
//}

func h(_ *ValveNode) int {
	return 0
}

func d(_, _ *ValveNode) int {
	return 1
}

func neighbours(node *ValveNode) []*ValveNode {
	return node.Children
}

//func findClosedValves(valves []*ValveNode) []*ValveNode {
//	var closed []*ValveNode
//	for _, valve := range valves {
//		if valve.Closed {
//			closed = append(closed, valve)
//		}
//	}
//
//	return closed
//}

func estimateRemainingReleasedPressure(without *ValveNode, closedValves []*ValveNode, timeRemaining int) int {
	// sort valves by flow rate
	sort.Slice(closedValves, func(i, j int) bool { return closedValves[i].FlowRate > closedValves[j].FlowRate })
	maxReleasedPressure := 0
	for _, closedValve := range closedValves {
		if closedValve == without {
			continue
		}

		// not enough time to get here and open valve
		if timeRemaining < 2 {
			return maxReleasedPressure
		}

		maxReleasedPressure += (timeRemaining - 2) * closedValve.FlowRate
		timeRemaining -= 2
	}

	return maxReleasedPressure
}

func findNextNode(from *ValveNode, closedValves []*ValveNode, timeRemaining int) (*ValveNode, int, int) {
	_, allCosts, _, _ := AStar(from, nil, h, d, neighbours)

	maxTotalPressureReleased := math.MinInt
	//maxPressureReleased := math.MinInt
	var maxPressureReleaseNode *ValveNode
	maxPathCost := 0
	maxPressureReleaseNodeIndex := -1
	//maxRemainingEstimate := 0

	for i, closedValve := range closedValves {
		pathCost := allCosts[closedValve]
		// not enough time to get here and open valve
		if pathCost+1 >= timeRemaining {
			continue
		}

		pressureReleased := (timeRemaining - pathCost - 1) * closedValve.FlowRate
		totalPressureReleased := pressureReleased + estimateRemainingReleasedPressure(closedValve, closedValves, timeRemaining-pathCost-1)
		//if pressureReleased > maxPressureReleased {
		if totalPressureReleased > maxTotalPressureReleased {
			maxTotalPressureReleased = totalPressureReleased
			//maxPressureReleased = pressureReleased
			maxPressureReleaseNode = closedValve
			maxPathCost = pathCost
			maxPressureReleaseNodeIndex = i
			//maxRemainingEstimate
		}
	}

	return maxPressureReleaseNode, maxPathCost, maxPressureReleaseNodeIndex
}

type State struct {
	CurrentNode *ValveNode
	//ClosedValves     []*ValveNode
	//ClosedValvesInt  int
	ClosedValvesSet  utils.BitSet16
	RemainingTime    int
	PressureReleased int
}

func (s State) ClosedValvesSlice(nodes []*ValveNode) []*ValveNode {
	var closed []*ValveNode
	for i, node := range nodes {
		if s.ClosedValvesSet.Contains(i) && node.FlowRate > 0 {
			closed = append(closed, node)
		}
	}
	return closed
}

func computeDistances(world World) utils.MatrixInt {
	distances := utils.NewMatrixInt(len(world.AllNodes), len(world.AllNodes))

	for i, nodeFrom := range world.AllNodes {
		_, allCosts, _, _ := utils.AStar(nodeFrom, nil, h, d, neighbours)

		for j, nodeTo := range world.AllNodes {
			distances.Columns[i][j] = allCosts[nodeTo]
		}
	}

	return distances
}

func FindMaxPressureReleaseState(world World) int {
	distances := computeDistances(world)

	fmt.Println(distances)

	h2 := func(state State) int {
		return -estimateRemainingReleasedPressure(state.CurrentNode, state.ClosedValvesSlice(world.AllNodes), state.RemainingTime)
	}

	d2 := func(s1, s2 State) int {
		return -(s2.PressureReleased - s1.PressureReleased)
	}

	neighbours2 := func(state State) []State {
		closedValves := state.ClosedValvesSlice(world.AllNodes)
		nextStates := make([]State, 0, len(closedValves))

		for _, closedValve := range closedValves {
			pathCost := distances.Columns[state.CurrentNode.Id][closedValve.Id]
			moveAndOpenCost := pathCost + 1
			// not enough time to get here and open valve
			if moveAndOpenCost >= state.RemainingTime {
				continue
			}

			remainingTime := state.RemainingTime - moveAndOpenCost

			pressureReleased := remainingTime * closedValve.FlowRate

			closedValvesSetNew := state.ClosedValvesSet.Clone()
			closedValvesSetNew.Remove(closedValve.Id)

			nextState := State{
				CurrentNode:      closedValve,
				ClosedValvesSet:  closedValvesSetNew,
				RemainingTime:    remainingTime,
				PressureReleased: state.PressureReleased + pressureReleased,
			}

			nextStates = append(nextStates, nextState)
		}

		return nextStates
	}

	//var s map[State]int 3:36

	initialState := State{
		CurrentNode:      world.RootNode,
		ClosedValvesSet:  utils.NewFullBitSet16(),
		RemainingTime:    30,
		PressureReleased: 0,
	}

	path, allCosts, cost, found := AStar(initialState, State{}, h2, d2, neighbours2)

	i, min := utils.ArgMin(utils.MapValues(allCosts)...)

	fmt.Printf("path %v, allCosts %v, cost %v, found %v, min %v, i %v\n", path, allCosts[initialState], cost, found, min, i)
	return -min
}

func FindMaxPressureRelease(world World) int {
	closedValves := world.AllNodes
	currentNode := world.RootNode
	timeRemaining := 30
	pressureReleased := 0

	for timeRemaining > 1 {
		fmt.Println(currentNode.Name)
		nextNode, pathCost, nextNodeIndex := findNextNode(currentNode, closedValves, timeRemaining)

		// move to the node and open vent
		timeRemaining -= pathCost + 1

		currentNode = nextNode
		closedValves = utils.RemoveUnordered(closedValves, nextNodeIndex)

		if timeRemaining > 0 {
			pressureReleased += timeRemaining * nextNode.FlowRate
		}

		if len(closedValves) == 0 {
			break
		}
	}

	return pressureReleased
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

	return World{
		RootNode:    nodes["AA"],
		AllNodes:    allNodes,
		AllNodesMap: nodes,
	}
}

func reconstructPath[T comparable](cameFrom map[T]T, current T) []T {
	totalPath := []T{current}
	for current, ok := cameFrom[current]; ok; current, ok = cameFrom[current] {
		totalPath = append([]T{current}, totalPath...)
	}

	return totalPath
}

func nodeWithLowestFScore[T comparable](openSet map[T]struct{}, fScore map[T]int) (T, int) {
	minScore := math.MaxInt
	var minNode T

	for node := range openSet {
		score, ok := fScore[node]
		if ok && score < minScore {
			minScore = score
			minNode = node
		}
	}

	return minNode, minScore
}

// AStar algorithm as in https://en.wikipedia.org/wiki/A%2A_search_algorithm
// h(n) int - heuristic function to calculate expected cost from node n to the goal node
// d(from, to) int - cost function for step from node "from" to node "to
// neighbours(n) []T - function that returns all neighbours of node n
func AStar[T comparable](start, goal T, h func(T) int, d func(T, T) int, neighbours func(T) []T) (path []T, scores map[T]int, score int, found bool) {
	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	// This is usually implemented as a min-heap or priority queue rather than a hash-set.
	openSet := make(map[T]struct{})
	openSet[start] = struct{}{}

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start
	// to n currently known.
	cameFrom := make(map[T]T)

	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := make(map[T]int) //map with default value of Infinity
	gScore[start] = 0

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how cheap a path could be from start to finish if it goes through n.
	fScore := make(map[T]int) // map with default value of Infinity
	fScore[start] = h(start)

	for len(openSet) > 0 {
		// This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
		current, currentFScore := nodeWithLowestFScore(openSet, fScore) // the node in openSet having the lowest fScore[] value
		if current == goal {
			return reconstructPath(cameFrom, current), gScore, currentFScore, true
		}

		delete(openSet, current)

		for _, neighbour := range neighbours(current) {

			if _, ok := gScore[current]; !ok {
				panic("Not known")
			}

			// d(current,neighbor) is the weight of the edge from current to neighbor
			// tentative_gScore is the distance from start to the neighbor through current
			tentativeGScore := gScore[current] + d(current, neighbour)

			if neighbourGScore, ok := gScore[neighbour]; !ok || tentativeGScore < neighbourGScore {
				// This path to neighbor is better than any previous one. Record it!
				cameFrom[neighbour] = current
				gScore[neighbour] = tentativeGScore
				fScore[neighbour] = tentativeGScore + h(neighbour)

				if _, ok := openSet[neighbour]; !ok {
					openSet[neighbour] = struct{}{}
				}
			}
		}
	}

	// Open set is empty but goal was never reached
	return []T{}, gScore, 0, false
}
