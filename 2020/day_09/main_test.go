package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	assert.Equal(t, 20, len(numbers))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInput(numbers, 5)
	assert.Equal(t, 127, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInput(numbers, 25)
	assert.Equal(t, 20874512, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInput(numbers, 25)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInput(numbers, 25)
	assert.Equal(t, 0, result)
}
