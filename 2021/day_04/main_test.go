package day_04

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	drawn, bingos, err := ParseInput(reader)
	assert.Nil(t, err)

	assert.Equal(t, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}, drawn)

	bingo := bingos[0]
	assert.Equal(t, [5][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19}},
		bingo.Numbers)

	assert.Equal(t, 300, bingo.SumAll)
	assert.Equal(t, 0, bingo.SumMarked)

	bingo = bingos[2]
	assert.Equal(t, [5][5]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7}},
		bingo.Numbers)

	assert.Equal(t, 325, bingo.SumAll)
	assert.Equal(t, 0, bingo.SumMarked)

	i := 0
	for ; i < 11; i++ {
		winning, _ := bingo.Mark(drawn[i])
		assert.False(t, winning)
	}

	winning, score := bingo.Mark(drawn[i])
	assert.True(t, winning)
	assert.Equal(t, 4512, score)
}
