package utils

import (
	"bufio"
	"io"
	"strconv"
)

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

// Returns absolute integer value
func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

// Sums integers from N to M inclusive
func SumNtoM(n, m int) int {
	return (n + m) * (1 + m - n) / 2
}

// Restricts value to interval (low, high)
func Clamp(val, low, high int) int {
	if val < low {
		return low
	}

	if val > high {
		return high
	}

	return val
}
