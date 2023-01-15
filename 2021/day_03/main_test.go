package day_03

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils/parsers"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rows := parsers.ParseToStrings(reader)

	bits := mostCommonBits(rows, "1")
	assert.Equal(t, "10110", bits)

	gamma, epsilon, power := decodeGammaEpsilon(bits)

	assert.Equal(t, 22, gamma)
	assert.Equal(t, 9, epsilon)

	assert.Equal(t, 198, power)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rows := parsers.ParseToStrings(reader)

	bits := mostCommonBits(rows, "1")
	assert.Equal(t, "001100100101", bits)

	gamma, epsilon, power := decodeGammaEpsilon(bits)

	assert.Equal(t, 805, gamma)
	assert.Equal(t, 3290, epsilon)

	assert.Equal(t, 2648450, power)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rows := parsers.ParseToStrings(reader)

	oxygen, co2, lifeSupport := decodeOxygenCo2(rows)

	assert.Equal(t, 23, oxygen)
	assert.Equal(t, 10, co2)
	assert.Equal(t, 230, lifeSupport)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rows := parsers.ParseToStrings(reader)

	oxygen, co2, lifeSupport := decodeOxygenCo2(rows)

	assert.Equal(t, 841, oxygen)
	assert.Equal(t, 3384, co2)
	assert.Equal(t, 2845944, lifeSupport)
}

func decodeOxygenCo2(rows []string) (int, int, int) {
	oxygen := filterByBitCriteria(rows, false)
	oxygenI, _ := strconv.ParseInt(oxygen, 2, 0)

	co2 := filterByBitCriteria(rows, true)
	co2I, _ := strconv.ParseInt(co2, 2, 0)

	return int(oxygenI), int(co2I), int(oxygenI * co2I)
}

func filterByBitCriteria(rows []string, co2 bool) string {
	for bitIndex := 0; bitIndex < len(rows[0]); bitIndex++ {
		var bits string
		if co2 {
			bits = mostCommonBits(invertBitsRows(rows), "0")
		} else {
			bits = mostCommonBits(rows, "1")
		}

		var filtered []string

		for _, row := range rows {
			if row[bitIndex] == bits[bitIndex] {
				filtered = append(filtered, row)
			}
		}

		if len(filtered) == 1 {
			return filtered[0]
		}

		rows = filtered
	}

	return ""
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

func decodeGammaEpsilon(bits string) (int, int, int) {
	gamma, _ := strconv.ParseInt(bits, 2, 0)

	epsilon, _ := strconv.ParseInt(invertBits(bits), 2, 0)

	return int(gamma), int(epsilon), int(gamma * epsilon)
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

func invertBitsRows(rows []string) []string {
	var inverted []string

	for _, row := range rows {
		inverted = append(inverted, invertBits(row))
	}

	return inverted
}
