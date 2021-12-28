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
	assert.Equal(t, [5]int{22, 13, 17, 11, 0}, bingo.Numbers[0])
	assert.Equal(t, 300, bingo.SumAll)
	assert.Equal(t, 0, bingo.SumMarked)
}
