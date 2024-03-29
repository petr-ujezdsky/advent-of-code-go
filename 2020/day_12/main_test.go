package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_01_parse(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	assert.Equal(t, 5, len(instructions))
}

func Test_01_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	result := DoWithInputPart01(instructions)
	assert.Equal(t, 25, result)
}

func Test_01(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	result := DoWithInputPart01(instructions)
	assert.Equal(t, 1496, result)
}

func Test_02_example(t *testing.T) {
	reader, err := os.Open("data-00-example.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	result := DoWithInputPart02(instructions)
	assert.Equal(t, 286, result)
}

func Test_02(t *testing.T) {
	reader, err := os.Open("data-01.txt")
	assert.Nil(t, err)

	instructions := ParseInput(reader)

	result := DoWithInputPart02(instructions)
	assert.Equal(t, 63843, result)
}
