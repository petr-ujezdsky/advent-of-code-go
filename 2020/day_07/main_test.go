package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	assert.Equal(t, 9, len(bagRules))
	assert.Equal(t, "vibrant plum", bagRules[6].Color)
	assert.Equal(t, 5, bagRules[6].NeededCounts["faded blue"])
	assert.Equal(t, 6, bagRules[6].NeededCounts["dotted black"])

	assert.Equal(t, 0, len(bagRules[7].NeededCounts))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	result := DoWithInput(bagRules)
	assert.Equal(t, 0, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	result := DoWithInput(bagRules)
	assert.Equal(t, 0, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	result := DoWithInput(bagRules)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	result := DoWithInput(bagRules)
	assert.Equal(t, 0, result)
}
