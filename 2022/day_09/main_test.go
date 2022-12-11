package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	steps := ParseInput(reader)

	assert.Equal(t, 8, len(steps))
	assert.Equal(t, 'R', steps[7].Dir)
	assert.Equal(t, 2, steps[7].Amount)
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	steps := ParseInput(reader)

	result := DoWithInput(steps, 2)
	assert.Equal(t, 13, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	steps := ParseInput(reader)

	result := DoWithInput(steps, 2)
	assert.Equal(t, 6745, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	steps := ParseInput(reader)

	result := DoWithInput(steps, 10)
	assert.Equal(t, 1, result)
}

func Test_02_example_2(t *testing.T) {
	reader, err := os.Open("data-00-example-2.txt")
	assert.Nil(t, err)

	steps := ParseInput(reader)

	result := DoWithInput(steps, 10)
	assert.Equal(t, 36, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	steps := ParseInput(reader)

	result := DoWithInput(steps, 10)
	assert.Equal(t, 2793, result)
}
