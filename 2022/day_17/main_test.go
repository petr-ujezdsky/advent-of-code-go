package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	assert.Equal(t, 40, len(jetDirections))
	assert.Equal(t, "[1 1 1 -1 -1 1 -1 1 1 -1 -1 -1 1 1 -1 1 1 1 -1 -1 -1 1 1 1 -1 -1 -1 1 -1 -1 -1 1 1 -1 1 1 -1 -1 1 1]", fmt.Sprint(jetDirections))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	result := InspectFallingRocks(jetDirections)
	assert.Equal(t, 3068, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	result := InspectFallingRocks(jetDirections)
	assert.Equal(t, 3227, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	result := InspectFallingRocks(jetDirections)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	jetDirections := ParseInput(reader)

	result := InspectFallingRocks(jetDirections)
	assert.Equal(t, 0, result)
}
