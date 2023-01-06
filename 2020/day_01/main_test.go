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

	assert.Equal(t, 6, len(numbers))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInputTwo(numbers)
	assert.Equal(t, 514579, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInputTwo(numbers)
	assert.Equal(t, 955584, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInputThree(numbers)
	assert.Equal(t, 241861950, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInputThree(numbers)
	assert.Equal(t, 287503934, result)
}
