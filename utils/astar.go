package utils

import "math"

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

// AStar algorithm as in https://en.wikipedia.org/wiki/A*_search_algorithm
// h(n) int - heuristic function to calculate expected cost from node n to the goal node
// d(from, to) int - cost function for step from node "from" to node "to
// neighbours(n) []T - function that returns all neighbours of node n
func AStar[T comparable](start, goal T, h func(T) int, d func(T, T) int, neighbours func(T) []T) ([]T, int, bool) {
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
			return reconstructPath(cameFrom, current), currentFScore, true
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
	return []T{}, 0, false
}
