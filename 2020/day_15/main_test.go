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

	assert.Equal(t, 3, len(numbers))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInputPart01(numbers)
	assert.Equal(t, 436, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInputPart01(numbers)
	assert.Equal(t, 475, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInputPart02(numbers)
	assert.Equal(t, 0, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	numbers := ParseInput(reader)

	result := DoWithInputPart02(numbers)
	assert.Equal(t, 0, result)
}
