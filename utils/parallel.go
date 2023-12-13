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
