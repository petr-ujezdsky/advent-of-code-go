package utils

import "math"

type storage[T any] interface {
	Push(T)
	Pop() T
	Empty() bool
}

// BranchAndBoundDeepFirst finds state with minimal cost. It skips states having lower bound greater than currently
// found minimum.
// Uses deep-first search.
func BranchAndBoundDeepFirst[T comparable](start T, cost func(T) int, lowerBound func(T) int, nextStatesProvider func(T) []T) (int, T) {
	storage := NewStack[T]()
	return branchAndBound[T](&storage, start, cost, lowerBound, nextStatesProvider)
}

// BranchAndBoundBreadthFirst finds state with minimal cost. It skips states having lower bound greater than currently
// found minimum.
// Uses breadth-first search.
func BranchAndBoundBreadthFirst[T comparable](start T, cost func(T) int, lowerBound func(T) int, nextStatesProvider func(T) []T) (int, T) {
	storage := NewQueue[T]()
	return branchAndBound[T](&storage, start, cost, lowerBound, nextStatesProvider)
}

// branchAndBound finds state with minimal cost. It skips states having lower bound greater than currently found minimum.
func branchAndBound[T any](storage storage[T], start T, cost func(T) int, lowerBound func(T) int, nextStatesProvider func(T) []T) (int, T) {
	openSet := storage
	openSet.Push(start)

	min := math.MaxInt
	minState := start

	for !openSet.Empty() {

		current := openSet.Pop()

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
			if lowerBound(next) > min {
				continue
			}

			openSet.Push(next)
		}
	}

	return min, minState
}
