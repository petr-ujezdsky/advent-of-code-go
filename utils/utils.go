package utils

import (
	"bufio"
	"io"
	"regexp"
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
func ParseToMatrix(r io.Reader) (MatrixInt, error) {
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

	return NewMatrixNumberRowNotation(rows), scanner.Err()
}

// ParseToMatrixP returns the matrix of integers (panics in case of an error)
func ParseToMatrixP(r io.Reader) MatrixInt {
	m, err := ParseToMatrix(r)
	if err != nil {
		panic("Problem parsing integer matrix")
	}
	return m
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

// ParseInt parses string to number or panics
func ParseInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic("Can not convert " + str + " to number")
	}
	return v
}

// ParseBinary8 parses string with zeros and ones to 8-bit number
func ParseBinary8(onesAndZeros string) uint8 {
	v, err := strconv.ParseUint(onesAndZeros, 2, 8)
	if err != nil {
		panic("Can not convert binary string " + onesAndZeros + " to number")
	}
	return uint8(v)
}

// ParseBinary16 parses string with zeros and ones to 16-bit number
func ParseBinary16(onesAndZeros string) uint16 {
	v, err := strconv.ParseUint(onesAndZeros, 2, 16)
	if err != nil {
		panic("Can not convert binary string " + onesAndZeros + " to number")
	}
	return uint16(v)
}

var regexIntNegative = regexp.MustCompile(`-?\d+`)
var regexIntPositive = regexp.MustCompile(`\d+`)

// ExtractInts extracts all integers in given string
func ExtractInts(str string, allowNegative bool) []int {
	var regex *regexp.Regexp

	if allowNegative {
		regex = regexIntNegative
	} else {
		regex = regexIntPositive
	}

	stringValues := regex.FindAllString(str, -1)

	ints := make([]int, len(stringValues))
	for i, stringValue := range stringValues {
		ints[i] = ParseInt(stringValue)
	}

	return ints
}

// Abs returns absolute value
func Abs[T Number](i T) T {
	if i < 0 {
		return -i
	}

	return i
}

// Signum returns 1 for positive number, -1 for negative and 0 for 0
func Signum[T Number](i T) T {
	if i < 0 {
		return -1
	}

	if i > 0 {
		return 1
	}

	return 0
}

// Max returns maximum of two numbers
func Max[T Number](i, j T) T {
	if i > j {
		return i
	}

	return j
}

// Min returns minimum of two numbers
func Min[T Number](i, j T) T {
	if i > j {
		return j
	}

	return i
}

// ArgMin finds index and value of minimum
func ArgMin[T Number](values ...T) (int, T) {
	min := values[0]
	index := -1

	for i, v := range values {
		if v <= min {
			min = v
			index = i
		}
	}

	return index, min
}

// ArgMax finds index and value of maximum
func ArgMax[T Number](values ...T) (int, T) {
	max := values[0]
	index := -1

	for i, v := range values {
		if v <= max {
			max = v
			index = i
		}
	}

	return index, max
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

func NextPowOf2(n int) int {
	k := 1
	for k < n {
		k = k << 1
	}
	return k
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

func EqualSlice[T comparable](a, b []T) bool {
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

func FilledSlice[T any](v T, length int) []T {
	s := make([]T, length)
	for i := 0; i < len(s); i++ {
		s[i] = v
	}
	return s
}

func FillSlice[T any](slice []T, value T) {
	for i := 0; i < len(slice); i++ {
		slice[i] = value
	}
}

// MapValues returns slice of map values
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	i := 0

	for _, value := range m {
		values[i] = value
		i++
	}

	return values
}

func FirstMapKey[T comparable](m map[T]struct{}) T {
	for key := range m {
		return key
	}
	panic("Empty set")
}

func Msg(str string) string {
	return str[1:]
}
