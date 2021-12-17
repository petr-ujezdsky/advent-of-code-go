package day_01_test

import (
	"os"
	"testing"

	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	data, err := utils.ParseToInts(reader)
	assert.Nil(t, err)

	count := getIncreasesCount(data)

	assert.Equal(t, 7, count)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	data, err := utils.ParseToInts(reader)
	assert.Nil(t, err)

	count := getIncreasesCount(data)

	assert.Equal(t, 1791, count)
}

func getIncreasesCount(data []int) int {
	var increasesCount = 0

	for i := 1; i < len(data); i++ {
		diff := data[i] - data[i-1]
		if diff > 0 {
			increasesCount++
		}
	}

	return increasesCount
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	data, err := utils.ParseToInts(reader)
	assert.Nil(t, err)

	count := getGroupIncreasesCount(data)

	assert.Equal(t, 5, count)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	data, err := utils.ParseToInts(reader)
	assert.Nil(t, err)

	count := getGroupIncreasesCount(data)

	assert.Equal(t, 1822, count)
}

// The middle values (200, 208) are irrelevant, they would cancel out. So only #1 and #4 values are diffed.
// 199  A
// 200  A B
// 208  A B
// 210    B
func getGroupIncreasesCount(data []int) int {
	var increasesCount = 0

	for i := 0; i < len(data)-3; i++ {
		diff := data[i+3] - data[i]

		if diff > 0 {
			increasesCount++
		}
	}

	return increasesCount
}
