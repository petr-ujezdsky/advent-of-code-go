package utils

import (
	"sync"
)

type Result[T any] struct {
	Index int
	Value T
}

func ProcessParallel[T any, R any](workUnits []T, processor func(T, int) R) chan Result[R] {
	resultsChan := make(chan Result[R])
	var wg sync.WaitGroup

	for i, workUnit := range workUnits {
		wg.Add(1)
		go func(index int, data T, ch chan Result[R], wg2 *sync.WaitGroup) {
			defer wg2.Done()
			result := processor(data, index)

			ch <- Result[R]{
				Index: index,
				Value: result,
			}
		}(i, workUnit, resultsChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	return resultsChan
}

func ProcessSerial[T any, R any](workUnits []T, processor func(T, int) R) chan Result[R] {
	resultsChan := make(chan Result[R])

	go func(data []T, ch chan Result[R]) {
		defer close(ch)

		for i, workUnit := range data {
			result := processor(workUnit, i)

			ch <- Result[R]{
				Index: i,
				Value: result,
			}
		}
	}(workUnits, resultsChan)

	return resultsChan
}
