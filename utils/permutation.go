package utils

import "github.com/petr-ujezdsky/advent-of-code-go/utils/slices"

// see https://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/

func Permute[T any](quit chan interface{}, values []T) chan []T {
	output := make(chan []T)

	go permuteAndClose(quit, values, output)

	return output
}

func permuteAndClose[T any](quit chan interface{}, values []T, output chan []T) {
	defer close(output)
	permute(quit, values, 0, len(values)-1, output)
}

func permute[T any](quit chan interface{}, values []T, left, right int, output chan []T) bool {
	if left == right {
		// output slice copy
		select {
		case output <- slices.ShallowCopy(values):
		case <-quit:
			return false
		}
	} else {
		for i := left; i <= right; i++ {
			// swap #1
			values[left], values[i] = values[i], values[left]

			if !permute(quit, values, left+1, right, output) {
				return false
			}

			// swap #2
			values[left], values[i] = values[i], values[left]
		}
	}

	return true
}
