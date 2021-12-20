package day_03

import (
	"os"
	"strconv"
	"testing"

	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rows, err := utils.ParseToStrings(reader)
	assert.Nil(t, err)

	bits := mostCommonBits(rows, "1")
	assert.Equal(t, "10110", bits)

	gamma, epsilon := decodeGammaEpsilon(bits)

	assert.Equal(t, 22, gamma)
	assert.Equal(t, 9, epsilon)

	assert.Equal(t, 198, gamma*epsilon)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rows, err := utils.ParseToStrings(reader)
	assert.Nil(t, err)

	bits := mostCommonBits(rows, "1")
	assert.Equal(t, "001100100101", bits)

	gamma, epsilon := decodeGammaEpsilon(bits)

	assert.Equal(t, 805, gamma)
	assert.Equal(t, 3290, epsilon)

	assert.Equal(t, 2648450, gamma*epsilon)
}

func mostCommonBits(rows []string, equalityBit string) string {
	onesCount := make([]int, len(rows[0]))

	for _, row := range rows {
		for i, ch := range row {
			if ch == '1' {
				onesCount[i]++
			}
		}
	}

	var bits string
	for _, count1 := range onesCount {
		count0 := len(rows) - count1

		if count1 > count0 {
			bits += "1"
		} else if count1 == count0 {
			bits += equalityBit
		} else {
			bits += "0"
		}
	}

	return bits
}

func decodeGammaEpsilon(bits string) (int, int) {
	gamma, _ := strconv.ParseInt(bits, 2, 0)

	epsilon, _ := strconv.ParseInt(invertBits(bits), 2, 0)

	return int(gamma), int(epsilon)
}

func invertBits(bits string) string {
	inverted := ""

	for _, ch := range bits {
		if ch == '0' {
			inverted += "1"
		} else {
			inverted += "0"
		}
	}

	return inverted
}
