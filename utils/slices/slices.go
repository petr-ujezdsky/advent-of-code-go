package slices

// Clone creates shallow copy of the given slice
func Clone[T any](slice []T) []T {
	// prepare destination slice
	cloned := make([]T, len(slice))

	// copy elements
	copy(cloned, slice)

	// return
	return cloned
}

// CloneAndAdd creates shallow copy of the given slice and adds new item
func CloneAndAdd[T any](slice []T, item T) []T {
	// prepare destination slice
	cloned := make([]T, len(slice)+1)

	// copy elements
	copy(cloned, slice)

	// add item at the end
	cloned[len(cloned)-1] = item

	// return
	return cloned
}

// Copy copies all values from source slice into target slice
func Copy[T any](source []T, target []T) {
	for i, v := range source {
		target[i] = v
	}
}

func Reverse[T any](slice []T) []T {
	length := len(slice)
	reversed := make([]T, length)

	for i, v := range slice {
		reversed[length-i-1] = v
	}

	return reversed
}

// RemoveUnordered removes element at index i and returns slice without this element. Changes items order in slice.
func RemoveUnordered[T any](s []T, i int) []T {
	// swap i-th and last element
	s[i] = s[len(s)-1]

	// return len-1 elements
	return s[:len(s)-1]
}

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, av := range a {
		if av != b[i] {
			return false
		}
	}
	return true
}

func Filled[T any](v T, length int) []T {
	s := make([]T, length)
	for i := 0; i < len(s); i++ {
		s[i] = v
	}
	return s
}

func Fill[T any](slice []T, value T) {
	for i := 0; i < len(slice); i++ {
		slice[i] = value
	}
}
