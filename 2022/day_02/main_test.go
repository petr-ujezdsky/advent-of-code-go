package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_decrypt(t *testing.T) {
	assert.Equal(t, 'A', decrypt('X'))
	assert.Equal(t, 'B', decrypt('Y'))
	assert.Equal(t, 'C', decrypt('Z'))
}

func Test_01_outcomeScore(t *testing.T) {
	assert.Equal(t, 3, outcomeScore("AA"))
	assert.Equal(t, 6, outcomeScore("BA"))
	assert.Equal(t, 0, outcomeScore("CA"))

	assert.Equal(t, 0, outcomeScore("AB"))
	assert.Equal(t, 3, outcomeScore("BB"))
	assert.Equal(t, 6, outcomeScore("CB"))

	assert.Equal(t, 6, outcomeScore("AC"))
	assert.Equal(t, 0, outcomeScore("BC"))
	assert.Equal(t, 3, outcomeScore("CC"))
}

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rounds := ParseInput(reader)

	assert.Equal(t, 3, len(rounds))
	assert.Equal(t, []rune("CC"), rounds[2])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rounds := ParseInput(reader)

	score := Score(rounds)
	assert.Equal(t, 15, score)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rounds := ParseInput(reader)

	score := Score(rounds)
	assert.Equal(t, 9177, score)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	rounds := ParseInput(reader)

	score := Score02(rounds)
	assert.Equal(t, 12, score)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	rounds := ParseInput(reader)

	score := Score02(rounds)
	assert.Equal(t, 12111, score)
}
