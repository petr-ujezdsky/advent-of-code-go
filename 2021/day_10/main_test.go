package day_10

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_01_example_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rows := ParseInput(reader)

	// first row
	assert.Equal(t, "[({(<(())[]>[[{[]{<()<>>", rows[0])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rows := ParseInput(reader)

	score := CorruptionScore(rows)

	assert.Equal(t, 26397, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rows := ParseInput(reader)

	score := CorruptionScore(rows)

	assert.Equal(t, 296535, score)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rows := ParseInput(reader)

	score := IncompleteScore(rows)

	assert.Equal(t, 288957, score)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rows := ParseInput(reader)

	score := IncompleteScore(rows)

	assert.Equal(t, 4245130838, score)
}
