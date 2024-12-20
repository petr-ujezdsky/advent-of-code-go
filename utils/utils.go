package utils

import (
	"bufio"
	"io"
	"math"
	"regexp"
	"strconv"
)

// ParseToInts parses each line as integer and returns the list
func ParseToInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var ints []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return ints, err
		}
		ints = append(ints, x)
	}

	return ints, scanner.Err()
}

// ParseToIntsP parses each line as integer and returns the list, panics in case of an error
func ParseToIntsP(r io.Reader) []int {
	ints, err := ParseToInts(r)
	if err != nil {
		panic(err)
	}

	return ints
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

// ToIntsP parses each line into integer and returns the list, panics in case of an error
func ToIntsP(intsStr []string) []int {
	ints, err := ToInts(intsStr)
	if err != nil {
		panic(err)
	}
	return ints
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
// Most significant bit is on the left side of the string
func ParseBinary8(onesAndZeros string) uint8 {
	v, err := strconv.ParseUint(onesAndZeros, 2, 8)
	if err != nil {
		panic("Can not convert binary string " + onesAndZeros + " to number")
	}
	return uint8(v)
}

// ParseBinary16 parses string with zeros and ones to 16-bit number
// Most significant bit is on the left side of the string
func ParseBinary16(onesAndZeros string) uint16 {
	v, err := strconv.ParseUint(onesAndZeros, 2, 16)
	if err != nil {
		panic("Can not convert binary string " + onesAndZeros + " to number")
	}
	return uint16(v)
}

// ParseBinary32 parses string with zeros and ones to 32-bit number
// Most significant bit is on the left side of the string
func ParseBinary32(onesAndZeros string) uint32 {
	v, err := strconv.ParseUint(onesAndZeros, 2, 32)
	if err != nil {
		panic("Can not convert binary string " + onesAndZeros + " to number")
	}
	return uint32(v)
}

// ParseBinary64 parses string with zeros and ones to 64-bit number
// Most significant bit is on the left side of the string
func ParseBinary64(onesAndZeros string) uint64 {
	v, err := strconv.ParseUint(onesAndZeros, 2, 64)
	if err != nil {
		panic("Can not convert binary string " + onesAndZeros + " to number")
	}
	return v
}

// ParseBinaryBool16 parses boolean slice to 16-bit number
// Most significant bit is on the left side of the slice
func ParseBinaryBool16(bits []bool) uint16 {
	if len(bits) > 16 {
		panic("Too many bits " + strconv.Itoa(len(bits)))
	}

	sum := uint16(0)
	for i, bit := range bits {
		if bit {
			sum += 1 << (len(bits) - 1 - i)
		}
	}

	return sum
}

// ParseBinaryBool64 parses boolean slice to 64-bit number
// Most significant bit is on the left side of the slice
func ParseBinaryBool64(bits []bool) uint64 {
	if len(bits) > 64 {
		panic("Too many bits " + strconv.Itoa(len(bits)))
	}

	sum := uint64(0)
	for i, bit := range bits {
		if bit {
			sum += 1 << (len(bits) - 1 - i)
		}
	}

	return sum
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

func Copysign[T Number](a, b T) T {
	return T(math.Copysign(float64(a), float64(b)))
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

func Sum[T AnyNumber](numbers []T) T {
	sum := T(0)
	for _, number := range numbers {
		sum += number
	}
	return sum
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

// ModFloor modifies modulo operator to work with negative values
// -2 % 10          = -2
// ModFloor(-2, 10) = 8
// see https://stackoverflow.com/a/43827557/1310733
func ModFloor(value, size int) int {
	return (((value) % size) + size) % size
}

func Msg(str string) string {
	return str[1:]
}

func FloorTowardsZero(value float64) float64 {
	if value > 0 {
		return math.Floor(value)
	}

	return -math.Floor(-value)
}

func CeilAwayFromZero(value float64) float64 {
	if value > 0 {
		return math.Ceil(value)
	}

	return -math.Ceil(-value)
}
