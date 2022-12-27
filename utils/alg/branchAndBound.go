package alg

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"math"
)

type storage[T any] interface {
	Push(T, int)
	Pop() T
	Empty() bool
}

// BranchAndBoundBestFirst finds state with minimal cost. It skips states having lower bound greater than currently
// found minimum.
// Uses best-first search.
func BranchAndBoundBestFirst[T any](start T, cost func(T) int, lowerBound func(T) int, nextStatesProvider func(T) []T) (min int, minState T) {
	storage := utils.NewMinHeapInt[T]()
	return branchAndBound[T](&storage, start, cost, lowerBound, nextStatesProvider)
}

// BranchAndBoundDeepFirst finds state with minimal cost. It skips states having lower bound greater than currently
// found minimum.
// Uses deep-first search.
func BranchAndBoundDeepFirst[T any](start T, cost func(T) int, lowerBound func(T) int, nextStatesProvider func(T) []T) (min int, minState T) {
	storage := newStackStorage[T]()
	return branchAndBound[T](&storage, start, cost, lowerBound, nextStatesProvider)
}

// BranchAndBoundBreadthFirst finds state with minimal cost. It skips states having lower bound greater than currently
// found minimum.
// Uses breadth-first search.
func BranchAndBoundBreadthFirst[T any](start T, cost func(T) int, lowerBound func(T) int, nextStatesProvider func(T) []T) (min int, minState T) {
	storage := newQueueStorage[T]()
	return branchAndBound[T](&storage, start, cost, lowerBound, nextStatesProvider)
}

// branchAndBound finds state with minimal cost. It skips states having lower bound greater than currently found minimum.
func branchAndBound[T any](storage storage[T], start T, cost func(T) int, lowerBound func(T) int, nextStatesProvider func(T) []T) (min int, minState T) {
	openSet := storage
	openSet.Push(start, lowerBound(start))

	min = math.MaxInt
	minState = start

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
			bound := lowerBound(next)
			if bound > min {
				continue
			}

			openSet.Push(next, bound)
		}
	}

	return min, minState
}

// Adapters to storage interface

type basicStorage[T any] interface {
	Push(T)
	Pop() T
	Empty() bool
}

type stackStorage[T any] struct {
	storage basicStorage[T]
}

func newStackStorage[T any]() stackStorage[T] {
	stack := utils.NewStack[T]()
	return stackStorage[T]{storage: &stack}
}

func newQueueStorage[T any]() stackStorage[T] {
	queue := utils.NewQueue[T]()
	return stackStorage[T]{storage: &queue}
}

func (s *stackStorage[T]) Push(item T, _ int) {
	s.storage.Push(item)
}

func (s *stackStorage[T]) Pop() T {
	return s.storage.Pop()
}

func (s *stackStorage[T]) Empty() bool {
	return s.storage.Empty()
}
