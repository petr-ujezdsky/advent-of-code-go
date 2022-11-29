package utils

import (
	"bufio"
	"io"
	"math"
	"strconv"
)

// ParseToInts parses each line as integer and returns the list
func ParseToInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}

	return result, scanner.Err()
}

// ParseToStrings returns the list of lines
func ParseToStrings(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result []string

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result, scanner.Err()
}

// ParseToMatrix returns the matrix of integers
func ParseToMatrix(r io.Reader) (Matrix2i, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var rows [][]int

	for scanner.Scan() {
		line := scanner.Text()
		var row []int

		for _, digitAscii := range []rune(line) {
			digit := int(digitAscii) - int('0')
			row = append(row, digit)
		}

		rows = append(rows, row)
	}

	return NewMatrix2RowNotation(rows), scanner.Err()
}

// ToInts parses each line into integer and returns the list
func ToInts(intsStr []string) ([]int, error) {
	var result []int

	for _, s := range intsStr {
		i, err := strconv.Atoi(s)
		if err != nil {
			return result, err
		}

		result = append(result, i)
	}

	return result, nil
}

// Abs returns absolute integer value
func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

// Signum returns 1 for positive number, -1 for negative and 0 for 0
func Signum(i int) int {
	if i < 0 {
		return -1
	}

	if i > 0 {
		return 1
	}

	return 0
}

// Max returns maximum of two integers
func Max(i, j int) int {
	if i > j {
		return i
	}

	return j
}

// Min returns minimum of two integers
func Min(i, j int) int {
	if i > j {
		return j
	}

	return i
}

// ArgMin finds index and value of minimum
func ArgMin(values ...int) (int, int) {
	min := math.MaxInt
	index := -1

	for i, v := range values {
		if v <= min {
			min = v
			index = i
		}
	}

	return index, min
}

// SumNtoM sums integers from N to M inclusive
func SumNtoM(n, m int) int {
	return (n + m) * (1 + m - n) / 2
}

// Clamp restricts value to interval (low, high)
func Clamp(val, low, high int) int {
	if val < low {
		return low
	}

	if val > high {
		return high
	}

	return val
}

// ShallowCopy creates shallow copy of the given slice
func ShallowCopy[T any](slice []T) []T {
	// prepare destination slice
	cloned := make([]T, len(slice))

	// copy elements
	copy(cloned, slice)

	// return
	return cloned
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
