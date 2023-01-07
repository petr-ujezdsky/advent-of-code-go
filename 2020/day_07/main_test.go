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
	assert.Equal(t, "vibrant plum", bagRules["vibrant plum"].Color)
	assert.Equal(t, 5, bagRules["vibrant plum"].NeededCounts["faded blue"])
	assert.Equal(t, 6, bagRules["vibrant plum"].NeededCounts["dotted black"])

	assert.Equal(t, 0, len(bagRules["faded blue"].NeededCounts))
}

func Test_01_ExpandRules(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	expandedNeededCounts := ExpandRules("faded blue", bagRules)
	assert.Equal(t, 0, len(expandedNeededCounts))

	expandedNeededCounts = ExpandRules("light red", bagRules)
	assert.Equal(t, 7, len(expandedNeededCounts))
	assert.Equal(t, 83, expandedNeededCounts["faded blue"])
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	result := ContainableByBagsCount(bagRules)
	assert.Equal(t, 4, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	result := ContainableByBagsCount(bagRules)
	assert.Equal(t, 326, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	result := TotalBagsCount(bagRules)
	assert.Equal(t, 32, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	bagRules := ParseInput(reader)

	result := TotalBagsCount(bagRules)
	assert.Equal(t, 5635, result)
}
