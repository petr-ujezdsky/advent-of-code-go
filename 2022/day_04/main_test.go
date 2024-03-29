package main

import (
	"github.com/petr-ujezdsky/advent-of-code-go/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	assert.Equal(t, 6, len(pairs))
	assert.Equal(t, utils.NewInterval(2, 6), pairs[5].Left)
	assert.Equal(t, utils.NewInterval(4, 8), pairs[5].Right)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	count := CountContained(pairs)
	assert.Equal(t, 2, count)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	count := CountContained(pairs)
	assert.Equal(t, 605, count)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	count := CountOverlapped(pairs)
	assert.Equal(t, 4, count)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	pairs := ParseInput(reader)

	count := CountOverlapped(pairs)
	assert.Equal(t, 914, count)
}
