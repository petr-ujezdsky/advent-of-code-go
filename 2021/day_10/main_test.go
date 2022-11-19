package day_10

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rows, err := ParseInput(reader)
	assert.Nil(t, err)

	// first row
	assert.Equal(t, "[({(<(())[]>[[{[]{<()<>>", rows[0])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rows, err := ParseInput(reader)
	assert.Nil(t, err)

	score := CorruptionScore(rows)

	assert.Equal(t, 26397, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rows, err := ParseInput(reader)
	assert.Nil(t, err)

	score := CorruptionScore(rows)

	assert.Equal(t, 296535, score)
}
