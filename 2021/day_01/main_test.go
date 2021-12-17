package day_01_test

import (
	"os"
	"testing"

	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
)

func Test_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	data, err := utils.ParseToInts(reader)
	assert.Nil(t, err)

	count := getIncreasesCount(data)

	assert.Equal(t, 7, count)
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
