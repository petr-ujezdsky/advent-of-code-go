package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse_BoardingPass(t *testing.T) {
	parsed := NewBoardingPass("FBFBBFFRLR")
	expected := BoardingPass{
		Row: 44,
		Col: 5,
	}

	assert.Equal(t, expected, parsed)
}
func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	boardingPasses := ParseInput(reader)

	assert.Equal(t, 4, len(boardingPasses))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	boardingPasses := ParseInput(reader)

	result := FindMaxSeatId(boardingPasses)
	assert.Equal(t, 820, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	boardingPasses := ParseInput(reader)

	result := FindMaxSeatId(boardingPasses)
	assert.Equal(t, 864, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	boardingPasses := ParseInput(reader)

	result := FindMissingSeatId(boardingPasses)
	assert.Equal(t, 0, result)
}
