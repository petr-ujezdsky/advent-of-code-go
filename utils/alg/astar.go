package alg

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
)

func reconstructPath[T comparable](cameFrom map[T]T, current T) []T {
	totalPath := []T{current}
	for current, ok := cameFrom[current]; ok; current, ok = cameFrom[current] {
		totalPath = append([]T{current}, totalPath...)
	}

	return totalPath
}

// AStar algorithm as in https://en.wikipedia.org/wiki/A%2A_search_algorithm
// h(n) int - heuristic function to calculate expected cost from node n to the goal node
// d(from, to) int - cost function for step from node "from" to node "to
// neighbours(n) []T - function that returns all neighbours of node n
func AStar[T comparable](start, goal T, h func(T) int, d func(T, T) int, neighbours func(T) []T) (path []T, scores map[T]int, score int, found bool) {
	return AStarEndFunc(start, func(state T) bool { return state == goal }, h, d, neighbours)
}

// AStarEndFunc algorithm as in https://en.wikipedia.org/wiki/A%2A_search_algorithm
// h(n) int - heuristic function to calculate expected cost from node n to the goal node
// d(from, to) int - cost function for step from node "from" to node "to
// neighbours(n) []T - function that returns all neighbours of node n
func AStarEndFunc[T comparable](start T, isEnd func(T) bool, h func(T) int, d func(T, T) int, neighbours func(T) []T) (path []T, scores map[T]int, score int, found bool) {
	hStart := h(start)

	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	// This is usually implemented as a min-heap or priority queue rather than a hash-set.
	openSet := utils.NewMinHeapInt[T]()
	openSet.Push(start, hStart)

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start
	// to n currently known.
	cameFrom := make(map[T]T)

	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	gScore := make(map[T]int) //map with default value of Infinity
	gScore[start] = 0

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how cheap a path could be from start to finish if it goes through n.
	fScore := make(map[T]int) // map with default value of Infinity
	fScore[start] = hStart

	for !openSet.Empty() {
		// This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
		current, currentFScore := openSet.PopWithValue() // the node in openSet having the lowest fScore[] value

		if isEnd(current) {
			return reconstructPath(cameFrom, current), gScore, currentFScore, true
		}

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

				neighbourFScore := tentativeGScore + h(neighbour)
				fScore[neighbour] = neighbourFScore

				if openSet.Contains(neighbour) {
					openSet.Fix(neighbour, neighbourFScore)
				} else {
					openSet.Push(neighbour, neighbourFScore)
				}
			}
		}
	}

	// Open set is empty but goal was never reached
	return []T{}, gScore, 0, false
}
