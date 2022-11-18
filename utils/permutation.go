package utils

// see https://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/

func Permute[T any](values []T) chan []T {
	output := make(chan []T)

	go permuteAndClose(values, output)

	return output
}

func permuteAndClose[T any](values []T, output chan []T) {
	defer close(output)
	permute(values, 0, len(values)-1, output)
}

func permute[T any](values []T, left, right int, output chan []T) {
	if left == right {
		// output slice copy
		output <- ShallowCopy(values)
	} else {
		for i := left; i <= right; i++ {
			// swap #1
			values[left], values[i] = values[i], values[left]

			permute(values, left+1, right, output)

			// swap #2
			values[left], values[i] = values[i], values[left]
		}
	}
}
