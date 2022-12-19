package utils

import "math"

func BranchAndBound[T comparable](start T, cost func(T) int, lowerBound func(T) int, nextStatesProvider func(T) []T) (int, T) {
	openSet := make(map[T]struct{})
	openSet[start] = struct{}{}

	min := math.MaxInt
	minState := start

	for len(openSet) > 0 {

		current := FirstMapKey(openSet)
		delete(openSet, current)

		nextStates := nextStatesProvider(current)

		// current is terminal state
		if len(nextStates) == 0 {
			currentCost := cost(current)

			if currentCost < min {
				min = currentCost
				minState = current
			}

			continue
		}

		for _, next := range nextStates {
			if lowerBound(next) <= min {
				openSet[next] = struct{}{}
			}
		}
	}

	return min, minState
}
