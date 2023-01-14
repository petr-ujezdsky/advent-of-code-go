package slices

import (
	"fmt"
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"strings"
)

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

func Map[S, T any](slice []S, mapper func(s S) T) []T {
	mapped := make([]T, len(slice))

	for i, s := range slice {
		mapped[i] = mapper(s)
	}

	return mapped
}

func Filter[T any](slice []T, take func(s T) bool) []T {
	var filtered []T

	for _, s := range slice {
		if take(s) {
			filtered = append(filtered, s)
		}
	}

	return filtered
}

func ToMap[T any, K comparable](slice []T, keyMapper func(v T) K) map[K]T {
	m := make(map[K]T, len(slice))

	for _, value := range slice {
		m[keyMapper(value)] = value
	}

	return m
}

func ToSet[T comparable](slice []T) map[T]struct{} {
	m := make(map[T]struct{})

	for _, value := range slice {
		m[value] = struct{}{}
	}

	return m
}

// Max returns maximum value in the slice
func Max[T utils.Number](slice []T) T {
	max := slice[0]
	for _, value := range slice {
		max = utils.Max(max, value)
	}
	return max
}

// Differentials returns slice of differentials, result[i] = source[i+1] - source[i]
func Differentials[T utils.Number](slice []T) []T {
	diffs := make([]T, len(slice)-1)
	for i, value := range slice[0 : len(slice)-1] {
		diffs[i] = slice[i+1] - value
	}
	return diffs
}

// Sprintf converts slice into string using custom item formatter
func Sprintf[T any](slice []T, format string) string {
	sb := &strings.Builder{}
	sb.WriteRune('[')

	for i, value := range slice {
		if i != 0 {
			sb.WriteRune(' ')
		}
		sb.WriteString(fmt.Sprintf(format, value))
	}

	sb.WriteRune(']')
	return sb.String()
}
